package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/filecoin-project/go-address"
	marketapi "github.com/filecoin-project/venus/venus-shared/api/market/v1"
	mkTypes "github.com/filecoin-project/venus/venus-shared/types/market"
	"github.com/hunjixin/brightbird/env"
	"github.com/hunjixin/brightbird/types"
	"github.com/hunjixin/brightbird/version"
	"go.uber.org/fx"
	"text/tabwriter"
)

var Info = types.PluginInfo{
	Name:        "actor-upsert",
	Version:     version.Version(),
	Category:    types.TestExec,
	Description: "actor upsert",
}

type TestCaseParams struct {
	fx.In
	AdminToken  types.AdminToken
	K8sEnv      *env.K8sEnvDeployer      `json:"-"`
	VenusMarket env.IVenusMarketDeployer `json:"-"`
}

func Exec(ctx context.Context, params TestCaseParams) error {

	mAddr, err := actorUpsert(ctx, params)
	if err != nil {
		fmt.Printf("market net listen err: %v\n", err)
		return err
	}

	err = actorDelete(ctx, params, mAddr)
	if err != nil {
		fmt.Printf("market net listen err: %v\n", err)
		return err
	}

	id, err := actorList(ctx, params, mAddr)
	if id == "" {
		fmt.Printf("actor delete err: %v\n", err)
		return err
	}
	return nil
}

func actorUpsert(ctx context.Context, params TestCaseParams) (address.Address, error) {
	endpoint := params.VenusMarket.SvcEndpoint()
	if env.Debug {
		var err error
		endpoint, err = params.K8sEnv.PortForwardPod(ctx, params.VenusMarket.Pods()[0].GetName(), int(params.VenusMarket.Svc().Spec.Ports[0].Port))
		if err != nil {
			return address.Undef, err
		}
	}
	client, closer, err := marketapi.NewIMarketRPC(ctx, endpoint.ToHttp(), nil)
	if err != nil {
		return address.Undef, err
	}
	defer closer()

	miner := "t01999"
	mAddr, err := address.NewFromString(miner)
	if err != nil {
		return address.Undef, err
	}

	bAdd, err := client.ActorUpsert(ctx, mkTypes.User{Addr: mAddr})
	if err != nil {
		return address.Undef, nil
	}

	opr := "Add"
	if !bAdd {
		opr = "Update"
	}

	fmt.Printf("%s miner %s success\n", opr, mAddr)

	return mAddr, err
}

func actorDelete(ctx context.Context, params TestCaseParams, mAddr address.Address) error {
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

	err = client.ActorDelete(ctx, mAddr)
	if err != nil {
		return err
	}

	fmt.Printf("delete miner %s success\n", mAddr)

	return err
}

func actorList(ctx context.Context, params TestCaseParams, mAddr address.Address) (string, error) {
	endpoint := params.VenusMarket.SvcEndpoint()
	if env.Debug {
		var err error
		endpoint, err = params.K8sEnv.PortForwardPod(ctx, params.VenusMarket.Pods()[0].GetName(), int(params.VenusMarket.Svc().Spec.Ports[0].Port))
		if err != nil {
			return "", err
		}
	}
	client, closer, err := marketapi.NewIMarketRPC(ctx, endpoint.ToHttp(), nil)
	if err != nil {
		return "", err
	}
	defer closer()

	miners, err := client.ActorList(ctx)
	if err != nil {
		return "", nil
	}

	buf := &bytes.Buffer{}
	tw := tabwriter.NewWriter(buf, 2, 4, 2, ' ', 0)
	_, _ = fmt.Fprintln(tw, "miner\taccount")
	for _, miner := range miners {
		_, _ = fmt.Fprintf(tw, "%s\t%s\n", miner.Addr.String(), miner.Account)
		if miner.Addr == mAddr {
			return miner.Addr.String(), nil
		}
	}

	return "", err
}
