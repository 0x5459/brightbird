package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/filecoin-project/go-address"
	v1api "github.com/filecoin-project/venus/venus-shared/api/chain/v1"
	clientapi "github.com/filecoin-project/venus/venus-shared/api/market/client"
	vtypes "github.com/filecoin-project/venus/venus-shared/types"
	"github.com/hunjixin/brightbird/env"
	"github.com/hunjixin/brightbird/env/types"
	"github.com/hunjixin/brightbird/version"
	"github.com/libp2p/go-libp2p/core/peer"
	"go.uber.org/fx"
)

var Info = types.PluginInfo{
	Name:        "storage-deals",
	Version:     version.Version(),
	Category:    types.TestExec,
	Description: "storage-deals",
}

type TestCaseParams struct {
	fx.In
	K8sEnv                     *env.K8sEnvDeployer `json:"-"`
	VenusAuth                  env.IDeployer       `json:"-" svcname:"VenusAuth"`
	MarketClient               env.IDeployer       `json:"-" svcname:"MarketClient"`
	Venus                      env.IDeployer       `json:"-" svcname:"Venus"`
	VenusSectorManagerDeployer env.IDeployer       `json:"-" svcname:"VenusSectorManager"`
	CreateWallet               env.IExec           `json:"-" svcname:"CreateWallet"`
}

func Exec(ctx context.Context, params TestCaseParams) (env.IExec, error) {

	walletAddr, err := params.CreateWallet.Param("CreateWallet")
	if err != nil {
		return nil, err
	}

	minerAddr, err := CreateMiner(ctx, params, walletAddr.(address.Address))
	if err != nil {
		fmt.Printf("create miner failed: %v\n", err)
		return nil, err
	}

	minerInfo, err := GetMinerInfo(ctx, params, minerAddr)
	if err != nil {
		fmt.Printf("get miner info failed: %v\n", err)
		return nil, err
	}
	fmt.Println("miner info: %v", minerInfo)

	err = StorageAsksQuery(ctx, params, minerAddr)
	if err != nil {
		fmt.Printf("storage asks query failed: %v\n", err)
		return nil, err
	}
	return env.NewSimpleExec(), nil
}

func CreateMiner(ctx context.Context, params TestCaseParams, walletAddr address.Address) (address.Address, error) {
	cmd := []string{
		"./venus-sector-manager",
		"util",
		"miner",
		"create",
		"--sector-size=8MiB",
		"--exid=" + string(rune(rand.Intn(100000))),
	}
	cmd = append(cmd, "--from="+walletAddr.String())

	pods, err := params.VenusSectorManagerDeployer.Pods(ctx)
	if err != nil {
		return address.Undef, err
	}

	minerAddr, err := params.K8sEnv.ExecRemoteCmd(ctx, pods[0].GetName(), cmd...)
	if err != nil {
		return address.Undef, fmt.Errorf("exec remote cmd failed: %w\n", err)
	}

	addr, err := address.NewFromBytes(minerAddr)
	if err != nil {
		return address.Undef, err
	}
	return addr, nil
}

func GetMinerInfo(ctx context.Context, params TestCaseParams, minerAddr address.Address) (string, error) {
	getMinerCmd := []string{
		"./venus-sector-manager",
		"util",
		"miner",
		"info",
		minerAddr.String(),
	}

	pods, err := params.VenusSectorManagerDeployer.Pods(ctx)
	if err != nil {
		return "", err
	}

	minerInfo, err := params.K8sEnv.ExecRemoteCmd(ctx, pods[0].GetName(), getMinerCmd...)
	if err != nil {
		return "", fmt.Errorf("exec remote cmd failed: %w\n", err)
	}

	return string(minerInfo), nil
}

func StorageAsksQuery(ctx context.Context, params TestCaseParams, maddr address.Address) error {
	endpoint := params.MarketClient.SvcEndpoint()
	if env.Debug {
		pods, err := params.MarketClient.Pods(ctx)
		if err != nil {
			return err
		}

		svc, err := params.MarketClient.Svc(ctx)
		if err != nil {
			return err
		}

		endpoint, err = params.K8sEnv.PortForwardPod(ctx, pods[0].GetName(), int(svc.Spec.Ports[0].Port))
		if err != nil {
			return err
		}
	}
	api, closer, err := clientapi.NewIMarketClientRPC(ctx, endpoint.ToHttp(), nil)
	if err != nil {
		return err
	}
	defer closer()

	venusEndpoint := params.Venus.SvcEndpoint()
	if env.Debug {
		pods, err := params.Venus.Pods(ctx)
		if err != nil {
			return err
		}

		svc, err := params.Venus.Svc(ctx)
		if err != nil {
			return err
		}

		endpoint, err = params.K8sEnv.PortForwardPod(ctx, pods[0].GetName(), int(svc.Spec.Ports[0].Port))
		if err != nil {
			return err
		}
	}
	fnapi, closer, err := v1api.NewFullNodeRPC(ctx, venusEndpoint.ToHttp(), nil)
	if err != nil {
		return err
	}
	defer closer()

	var pid peer.ID

	mi, err := fnapi.StateMinerInfo(ctx, maddr, vtypes.EmptyTSK)
	if err != nil {
		return fmt.Errorf("failed to get peerID for miner: %w", err)
	}

	if mi.PeerId == nil || *mi.PeerId == peer.ID("SETME") {
		return fmt.Errorf("the miner hasn't initialized yet")
	}

	pid = *mi.PeerId

	ask, err := api.ClientQueryAsk(ctx, pid, maddr)
	if err != nil {
		return err
	}

	fmt.Printf("Ask: %s\n", maddr)
	fmt.Printf("Price per GiB: %s\n", vtypes.FIL(ask.Price))
	fmt.Printf("Verified Price per GiB: %s\n", vtypes.FIL(ask.VerifiedPrice))
	fmt.Printf("Max Piece size: %s\n", vtypes.SizeStr(vtypes.NewInt(uint64(ask.MaxPieceSize))))
	fmt.Printf("Min Piece size: %s\n", vtypes.SizeStr(vtypes.NewInt(uint64(ask.MinPieceSize))))

	return nil
}