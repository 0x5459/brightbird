package main

import (
	"context"
	"fmt"

	"github.com/filecoin-project/venus/venus-shared/api/messager"
	"github.com/hunjixin/brightbird/env"
	"github.com/hunjixin/brightbird/env/types"
	"github.com/hunjixin/brightbird/version"
	"go.uber.org/fx"
)

var Info = types.PluginInfo{
	Name:        "verity_message",
	Version:     version.Version(),
	Category:    types.TestExec,
	Description: "verity message if normal",
}

type TestCaseParams struct {
	fx.In
	K8sEnv       *env.K8sEnvDeployer `json:"-"`
	VenusMessage env.IDeployer       `json:"-" svcname:"VenusMessage"`
	VenusAuth    env.IDeployer       `json:"-" svcname:"VenusAuth"`
}

func Exec(ctx context.Context, params TestCaseParams) (env.IExec, error) {

	adminTokenV, err := params.VenusAuth.Param("AdminToken")
	if err != nil {
		return nil, err
	}

	err = CreateMessage(ctx, params, adminTokenV.(string))
	if err != nil {
		fmt.Printf("create message rpc failed: %v\n", err)
		return nil, err
	}

	return env.NewSimpleExec(), nil
}

func CreateMessage(ctx context.Context, params TestCaseParams, authToken string) error {
	endpoint := params.VenusMessage.SvcEndpoint()
	if env.Debug {
		pods, err := params.VenusMessage.Pods(ctx)
		if err != nil {
			return err
		}

		svc, err := params.VenusMessage.Svc(ctx)
		if err != nil {
			return err
		}
		endpoint, err = params.K8sEnv.PortForwardPod(ctx, pods[0].GetName(), int(svc.Spec.Ports[0].Port))
		if err != nil {
			return err
		}
	}

	client, closer, err := messager.DialIMessagerRPC(ctx, endpoint.ToHttp(), authToken, nil)
	if err != nil {
		return err
	}
	defer closer()

	messageVersion, err := client.Version(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Message: %v\n", messageVersion)

	return nil
}