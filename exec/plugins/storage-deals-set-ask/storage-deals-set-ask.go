package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/docker/go-units"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/venus/pkg/constants"
	marketapi "github.com/filecoin-project/venus/venus-shared/api/market/v1"
	"github.com/filecoin-project/venus/venus-shared/api/wallet"
	vTypes "github.com/filecoin-project/venus/venus-shared/types"
	"github.com/filecoin-project/venus/venus-shared/types/market"
	"github.com/hunjixin/brightbird/env"
	"github.com/hunjixin/brightbird/types"
	"github.com/hunjixin/brightbird/version"
	"go.uber.org/fx"
	"math/rand"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

var Info = types.PluginInfo{
	Name:        "actor-list",
	Version:     version.Version(),
	Category:    types.TestExec,
	Description: "actor list",
}

type TestCaseParams struct {
	fx.In
	AdminToken                 types.AdminToken
	K8sEnv                     *env.K8sEnvDeployer             `json:"-"`
	VenusMarket                env.IVenusMarketDeployer        `json:"-"`
	VenusWallet                env.IVenusWalletDeployer        `json:"-" svcname:"Wallet"`
	VenusSectorManagerDeployer env.IVenusSectorManagerDeployer `json:"-"`
}

func Exec(ctx context.Context, params TestCaseParams) error {

	walletAddr, err := CreateWallet(ctx, params)
	if err != nil {
		fmt.Printf("create wallet failed: %v\n", err)
		return err
	}

	minerAddr, err := CreateMiner(ctx, params, walletAddr)
	if err != nil {
		fmt.Printf("create miner failed: %v\n", err)
		return err
	}

	minerInfo, err := GetMinerInfo(ctx, params, minerAddr)
	if err != nil {
		fmt.Printf("get miner info failed: %v\n", err)
		return err
	}
	fmt.Println("miner info: %v", minerInfo)

	err = StorageAskSet(ctx, params, minerAddr)
	if err != nil {
		fmt.Printf("market net listen err: %v\n", err)
		return err
	}

	err = StorageAskGet(ctx, params, minerAddr)
	if err != nil {
		fmt.Printf("market net listen err: %v\n", err)
		return err
	}

	return nil
}

func StorageAskSet(ctx context.Context, params TestCaseParams, mAddr address.Address) error {
	endpoint := params.VenusMarket.SvcEndpoint()
	if env.Debug {
		var err error
		endpoint, err = params.K8sEnv.PortForwardPod(ctx, params.VenusMarket.Pods()[0].GetName(), int(params.VenusMarket.Svc().Spec.Ports[0].Port))
		if err != nil {
			return err
		}
	}
	client, closer, err := marketapi.NewIMarketRPC(ctx, endpoint.ToHttp(), nil)
	if err != nil {
		return err
	}
	defer closer()

	isUpdate := true
	storageAsk, err := client.MarketGetAsk(ctx, mAddr)
	if err != nil {
		if !strings.Contains(err.Error(), "record not found") {
			return err
		}
		storageAsk = &market.SignedStorageAsk{}
		isUpdate = false
	}

	pri, err := vTypes.ParseFIL("0.000000001")
	if err != nil {
		return err
	}

	vpri, err := vTypes.ParseFIL("0")
	if err != nil {
		return err
	}

	dur, err := time.ParseDuration("720h0m0s")
	if err != nil {
		return fmt.Errorf("cannot parse duration: %w", err)
	}

	qty := dur.Seconds() / float64(constants.MainNetBlockDelaySecs)

	min, err := units.RAMInBytes("1KB")
	if err != nil {
		return fmt.Errorf("cannot parse min-piece-size to quantity of bytes: %w", err)
	}

	if min < 256 {
		return errors.New("minimum piece size (w/bit-padding) is 256B")
	}

	max, err := units.RAMInBytes("32GiB")
	if err != nil {
		return fmt.Errorf("cannot parse max-piece-size to quantity of bytes: %w", err)
	}

	ssize, err := client.ActorSectorSize(ctx, mAddr)
	if err != nil {
		return fmt.Errorf("get miner's size %w", err)
	}

	smax := int64(ssize)

	if max == 0 {
		max = smax
	}

	if max > smax {
		return fmt.Errorf("max piece size (w/bit-padding) %s cannot exceed miner sector size %s", vTypes.SizeStr(vTypes.NewInt(uint64(max))), vTypes.SizeStr(vTypes.NewInt(uint64(smax))))
	}

	if isUpdate {
		storageAsk.Ask.Price = vTypes.BigInt(pri)
		storageAsk.Ask.VerifiedPrice = vTypes.BigInt(vpri)
		storageAsk.Ask.MinPieceSize = abi.PaddedPieceSize(min)
		storageAsk.Ask.MaxPieceSize = abi.PaddedPieceSize(max)
		return client.MarketSetAsk(ctx, mAddr, storageAsk.Ask.Price, storageAsk.Ask.VerifiedPrice, abi.ChainEpoch(qty), storageAsk.Ask.MinPieceSize, storageAsk.Ask.MaxPieceSize)
	}

	err = client.MarketSetAsk(ctx, mAddr, vTypes.BigInt(pri), vTypes.BigInt(vpri), abi.ChainEpoch(qty), abi.PaddedPieceSize(min), abi.PaddedPieceSize(max))
	if err != nil {
		return fmt.Errorf("market set ask err %w", err)
	}

	return err
}

func StorageAskGet(ctx context.Context, params TestCaseParams, mAddr address.Address) error {
	endpoint := params.VenusMarket.SvcEndpoint()
	if env.Debug {
		var err error
		endpoint, err = params.K8sEnv.PortForwardPod(ctx, params.VenusMarket.Pods()[0].GetName(), int(params.VenusMarket.Svc().Spec.Ports[0].Port))
		if err != nil {
			return err
		}
	}
	client, closer, err := marketapi.NewIMarketRPC(ctx, endpoint.ToHttp(), nil)
	if err != nil {
		return err
	}
	defer closer()

	sask, err := client.MarketGetAsk(ctx, mAddr)
	if err != nil {
		return err
	}

	var ask *storagemarket.StorageAsk
	if sask != nil && sask.Ask != nil {
		ask = sask.Ask
	}

	w := tabwriter.NewWriter(os.Stdout, 2, 4, 2, ' ', 0)
	fmt.Fprintf(w, "Price per GiB/Epoch\tVerified\tMin. Piece Size (padded)\tMax. Piece Size (padded)\tExpiry (Epoch)\tExpiry (Appx. Rem. Time)\tSeq. No.\n")
	if ask == nil {
		fmt.Fprintf(w, "<miner does not have an ask>\n")
		return w.Flush()
	}
	return nil
}

func CreateWallet(ctx context.Context, params TestCaseParams) (address.Address, error) {
	walletToken, err := env.ReadWalletToken(ctx, params.K8sEnv, params.VenusWallet.Pods()[0].GetName())
	if err != nil {
		return address.Undef, fmt.Errorf("read wallet token failed: %w\n", err)
	}

	endpoint := params.VenusWallet.SvcEndpoint()
	if env.Debug {
		var err error
		endpoint, err = params.K8sEnv.PortForwardPod(ctx, params.VenusWallet.Pods()[0].GetName(), int(params.VenusWallet.Svc().Spec.Ports[0].Port))
		if err != nil {
			return address.Undef, fmt.Errorf("port forward failed: %w\n", err)
		}
	}

	walletRpc, closer, err := wallet.DialIFullAPIRPC(ctx, endpoint.ToMultiAddr(), walletToken, nil)
	if err != nil {
		return address.Undef, fmt.Errorf("dial iFullAPI rpc failed: %w\n", err)
	}
	defer closer()

	password := "123456"
	err = walletRpc.SetPassword(ctx, password)
	if err != nil {
		return address.Undef, fmt.Errorf("set password failed: %w\n", err)
	}

	walletAddr, err := walletRpc.WalletNew(ctx, vTypes.KTBLS)
	if err != nil {
		return address.Undef, fmt.Errorf("create wallet failed: %w\n", err)
	}
	fmt.Printf("wallet: %v\n", walletAddr)

	return walletAddr, nil
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

	minerAddr, err := params.K8sEnv.ExecRemoteCmd(ctx, params.VenusSectorManagerDeployer.Pods()[0].GetName(), cmd)
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
	minerInfo, err := params.K8sEnv.ExecRemoteCmd(ctx, params.VenusSectorManagerDeployer.Pods()[0].GetName(), getMinerCmd)
	if err != nil {
		return "", fmt.Errorf("exec remote cmd failed: %w\n", err)
	}

	return string(minerInfo), nil
}