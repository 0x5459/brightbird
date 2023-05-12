package main

import (
	"context"
	chain "github.com/filecoin-project/venus/venus-shared/api/chain/v1"
	"github.com/hunjixin/brightbird/env"
	"github.com/hunjixin/brightbird/env/types"
	"github.com/hunjixin/brightbird/version"
	"go.uber.org/fx"
)

var Info = types.PluginInfo{
	Name:        "auth_request_venus",
	Version:     version.Version(),
	Category:    types.TestExec,
	Description: "auth request venus",
}

type TestCaseParams struct {
	fx.In
	Params struct {
		//Permission string `json:"permission"`
	} `optional:"true"`

	K8sEnv    *env.K8sEnvDeployer `json:"-"`
	VenusAuth env.IDeployer       `json:"-" svcname:"VenusAuth"`
	Venus     env.IDeployer       `json:"-" svcname:"Venus"`
}

func Exec(ctx context.Context, params TestCaseParams) (env.IExec, error) {

	adminToken, err := params.VenusAuth.Param("AdminToken")
	if err != nil {
		return nil, err
	}

	err = checkPermission(ctx, adminToken.(string), params)
	if err != nil {
		return nil, err
	}
	return env.NewSimpleExec(), nil
}

func checkPermission(ctx context.Context, token string, params TestCaseParams) error {
	endpoint := params.Venus.SvcEndpoint()
	if env.Debug {
		venusPods, err := params.Venus.Pods(ctx)
		if err != nil {
			return err
		}

		svc, err := params.Venus.Svc(ctx)
		if err != nil {
			return err
		}
		endpoint, err = params.K8sEnv.PortForwardPod(ctx, venusPods[0].GetName(), int(svc.Spec.Ports[0].Port))
		if err != nil {
			return err
		}
	}
	chainRpc, closer, err := chain.DialFullNodeRPC(ctx, endpoint.ToMultiAddr(), token, nil)
	if err != nil {
		return err
	}
	defer closer()

	_, err = chainRpc.ChainHead(ctx)
	if err != nil {
		return err
	}

	return nil
}