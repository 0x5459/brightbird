package main

import (
	"context"
	"fmt"
	"github.com/hunjixin/brightbird/env"
	"github.com/hunjixin/brightbird/types"
	"github.com/hunjixin/brightbird/version"
	"go.uber.org/fx"
)

var Info = types.PluginInfo{
	Name:        "test_deploy",
	Version:     version.Version(),
	Category:    types.TestExec,
	Description: "generate admin token",
}

type TestCaseParams struct {
	fx.In
	MarketClient   env.IMarketClientDeployer `json:"-"`
	VenusWallet    env.IVenusWalletDeployer  `json:"-" svcname:"Wallet"`
	VenusWalletNew env.IVenusWalletDeployer  `json:"-" svcname:"WalletNew"`
}

func Exec(ctx context.Context, params TestCaseParams) error {
	fmt.Println("Wallet", params.VenusWallet.SvcEndpoint())
	fmt.Println("WalletNew", params.VenusWalletNew.SvcEndpoint())
	return nil
}