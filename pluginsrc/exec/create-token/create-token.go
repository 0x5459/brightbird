package main

import (
	"context"

	"github.com/hunjixin/brightbird/env"
	"github.com/hunjixin/brightbird/env/plugin"
	sophonauth "github.com/hunjixin/brightbird/pluginsrc/deploy/sophon-auth"
	"github.com/hunjixin/brightbird/types"
	"github.com/hunjixin/brightbird/version"
	"github.com/ipfs-force-community/sophon-auth/jwtclient"
)

func main() {
	plugin.SetupPluginFromStdin(Info, Exec)
}

var Info = types.PluginInfo{
	Name:        "create_token",
	Version:     version.Version(),
	PluginType:  types.TestExec,
	Description: "create token",
}

type CreateTokenReturn struct {
	Token string `json:"token" description:"generated token"`
}

type TestCaseParams struct {
	Auth sophonauth.SophonAuthDeployReturn `json:"SophonAuth" jsonschema:"SophonAuth" title:"Sophon Auth" require:"true" description:"sophon auth return"`

	UserName string `json:"userName" jsonschema:"userName" title:"UserName" require:"true" description:"token user name"`
	Perm     string `json:"perm" jsonschema:"perm" title:"Perm" require:"true" description:"custom string in JWT payload" enum:"read,write,sign,admin"`
	Extra    string `json:"extra" jsonschema:"extra" title:"Extra" description:"addition information"`
}

func Exec(ctx context.Context, k8sEnv *env.K8sEnvDeployer, params TestCaseParams) (*CreateTokenReturn, error) {
	authAPIClient, err := jwtclient.NewAuthClient(params.Auth.SvcEndpoint.ToHTTP(), params.Auth.AdminToken)
	if err != nil {
		return nil, err
	}

	token, err := authAPIClient.GenerateToken(ctx, params.UserName, params.Perm, params.Extra)
	if err != nil {
		return nil, err
	}

	return &CreateTokenReturn{
		Token: token,
	}, nil
}
