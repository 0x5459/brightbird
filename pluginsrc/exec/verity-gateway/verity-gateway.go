package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/fx"

	"github.com/filecoin-project/go-address"
	v2API "github.com/filecoin-project/venus/venus-shared/api/gateway/v2"
	"github.com/hunjixin/brightbird/env"
	"github.com/hunjixin/brightbird/env/plugin"
	"github.com/hunjixin/brightbird/types"
	"github.com/hunjixin/brightbird/version"
)

func main() {
	plugin.SetupPluginFromStdin(Info, Exec)
}

var Info = types.PluginInfo{
	Name:        "verity_gateway",
	Version:     version.Version(),
	PluginType:  types.TestExec,
	Description: "verity gateway if normal",
}

type TestCaseParams struct {
	fx.In
	K8sEnv       *env.K8sEnvDeployer `json:"-"`
	VenusGateway env.IDeployer       `json:"-" svcname:"VenusGateway"`
	VenusWallet  env.IDeployer       `json:"-" svcname:"VenusWallet"`
	VenusAuth    env.IDeployer       `json:"-" svcname:"VenusAuth"`
	CreateWallet env.IExec           `json:"-" svcname:"CreateWallet"`
}

func Exec(ctx context.Context, params TestCaseParams) (env.IExec, error) {

	walletAddr, err := params.CreateWallet.Param("Wallet")
	if err != nil {
		return nil, err
	}

	adminTokenV, err := params.VenusAuth.Param("AdminToken")
	if err != nil {
		return nil, err
	}

	paramsBytes, err := json.Marshal(walletAddr)
	if err != nil {
		panic(err)
	}

	// 将字节数组转换为 address.Address 类型
	addr, err := address.NewFromBytes(paramsBytes)
	if err != nil {
		panic(err)
	}

	err = GetWalletInfo(ctx, params, adminTokenV.String(), addr)
	if err != nil {
		fmt.Printf("get wallet info failed: %v\n", err)
		return nil, err
	}

	return env.NewSimpleExec(), nil

}

func GetWalletInfo(ctx context.Context, params TestCaseParams, authToken string, walletAddr address.Address) error {
	endpoint, err := params.VenusWallet.SvcEndpoint()
	if err != nil {
		return err
	}

	if env.Debug {
		pods, err := params.VenusWallet.Pods(ctx)
		if err != nil {
			return err
		}

		svc, err := params.VenusWallet.Svc(ctx)
		if err != nil {
			return err
		}
		endpoint, err = params.K8sEnv.PortForwardPod(ctx, pods[0].GetName(), int(svc.Spec.Ports[0].Port))
		if err != nil {
			return err
		}
	}

	api, closer, err := v2API.DialIGatewayRPC(ctx, endpoint.ToHttp(), authToken, nil)
	if err != nil {
		return err
	}
	defer closer()

	wallets, err := api.ListWalletInfo(ctx)
	if err != nil {
		return err
	}
	for _, wallet := range wallets {
		if wallet.Account == walletAddr.String() {
			return nil
		}
	}
	return err
}
