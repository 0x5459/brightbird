package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/filecoin-project/venus-auth/jwtclient"
	"github.com/hunjixin/brightbird/env"
	"github.com/hunjixin/brightbird/env/plugin"
	"github.com/hunjixin/brightbird/types"
	"github.com/hunjixin/brightbird/version"
	"go.uber.org/fx"
)

func main() {
	plugin.SetupPluginFromStdin(Info, Exec)
}

var Info = types.PluginInfo{
	Name:        "list_token",
	Version:     version.Version(),
	PluginType:  types.TestExec,
	Description: "list token",
}

type TestCaseParams struct {
	fx.In
	Params struct {
		Skip  string `json:"skip"`
		Limit string `json:"limit"`
	} `optional:"true"`

	K8sEnv    *env.K8sEnvDeployer `json:"-"`
	VenusAuth env.IDeployer       `json:"-" svcname:"VenusAuth"`
}

func Exec(ctx context.Context, params TestCaseParams) (env.IExec, error) {
	endpoint, err := params.VenusAuth.SvcEndpoint()
	if err != nil {
		return nil, err
	}

	if env.Debug {
		pods, err := params.VenusAuth.Pods(ctx)
		if err != nil {
			return nil, err
		}

		svc, err := params.VenusAuth.Svc(ctx)
		if err != nil {
			return nil, err
		}
		endpoint, err = params.K8sEnv.PortForwardPod(ctx, pods[0].GetName(), int(svc.Spec.Ports[0].Port))
		if err != nil {
			return nil, err
		}
	}

	adminTokenV, err := params.VenusAuth.Param("AdminToken")
	if err != nil {
		return nil, err
	}

	authAPIClient, err := jwtclient.NewAuthClient(endpoint.ToHTTP(), adminTokenV.MustString())
	if err != nil {
		return nil, err
	}

	skip, err := strconv.ParseInt(params.Params.Skip, 10, 64)
	if err != nil {
		return nil, err
	}

	limit, err := strconv.ParseInt(params.Params.Limit, 10, 64)
	if err != nil {
		return nil, err
	}
	tokenList, err := authAPIClient.Tokens(ctx, skip, limit)
	if err != nil {
		return nil, err
	}
	for _, token := range tokenList {
		fmt.Println(token)
	}
	return env.NewSimpleExec(), nil
}
