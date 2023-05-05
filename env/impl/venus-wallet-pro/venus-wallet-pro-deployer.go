package venus_wallet_pro

import (
	"context"
	"embed"
	"fmt"

	wConfig "github.com/filecoin-project/venus-wallet/config"

	"github.com/hunjixin/brightbird/env"
	"github.com/hunjixin/brightbird/types"
	"github.com/hunjixin/brightbird/utils"
	"github.com/hunjixin/brightbird/version"
	"github.com/pelletier/go-toml"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type Config struct {
	env.BaseConfig

	GatewayUrl string `json:"-"`
	UserToken  string `json:"-"`

	UserName        string `json:"userName"`
	CreateIfNotExit bool   `json:"createIfNotExit"`
}

type RenderParams struct {
	Config

	PrivateRegistry string
	Args            []string

	UniqueId string
}

func DefaultConfig() Config {
	return Config{}
}

var PluginInfo = types.PluginInfo{
	Name:        string(env.VenusWalletPro),
	Version:     version.Version(),
	Category:    types.Deploy,
	Repo:        "",
	ImageTarget: "venus-wallet",
	Description: "",
}

var _ env.IVenusWalletProDeployer = (*VenusWalletProDeployer)(nil)

type VenusWalletProDeployer struct {
	env *env.K8sEnvDeployer
	cfg *Config

	svcEndpoint types.Endpoint

	configMapName   string
	statefulSetName string
	svcName         string
}

func NewVenusWalletProDeployer(env *env.K8sEnvDeployer, gatewayUrl, userToken string, supportAccounts ...string) *VenusWalletProDeployer {
	return &VenusWalletProDeployer{
		env: env,
		cfg: &Config{
			GatewayUrl: gatewayUrl,
			UserToken:  userToken,
		},
	}
}

func DeployerFromConfig(env *env.K8sEnvDeployer, cfg Config, params Config) (env.IVenusWalletProDeployer, error) {
	cfg, err := utils.MergeStructAndInterface(DefaultConfig(), cfg, params)
	if err != nil {
		return nil, err
	}
	return &VenusWalletProDeployer{
		env: env,
		cfg: &cfg,
	}, nil
}

func (deployer *VenusWalletProDeployer) Name() string {
	return PluginInfo.Name
}

func (deployer *VenusWalletProDeployer) Pods(ctx context.Context) ([]corev1.Pod, error) {
	return deployer.env.GetPodsByLabel(ctx, fmt.Sprintf("venus-wallet-%s-pod", deployer.env.UniqueId(deployer.cfg.SvcMap[types.OutLabel])))
}

func (deployer *VenusWalletProDeployer) StatefulSet(ctx context.Context) (*appv1.StatefulSet, error) {
	return deployer.env.GetStatefulSet(ctx, deployer.statefulSetName)
}

func (deployer *VenusWalletProDeployer) Svc(ctx context.Context) (*corev1.Service, error) {
	return deployer.env.GetSvc(ctx, deployer.svcName)
}

func (deployer *VenusWalletProDeployer) SvcEndpoint() types.Endpoint {
	return deployer.svcEndpoint
}

var f embed.FS

func (deployer *VenusWalletProDeployer) Deploy(ctx context.Context) (err error) {
	renderParams := RenderParams{
		PrivateRegistry: deployer.env.PrivateRegistry(),
		UniqueId:        deployer.env.UniqueId(deployer.cfg.SvcMap[types.OutLabel]),
		Config:          *deployer.cfg,
	}
	//create configmap
	configMapCfg, err := f.Open("venus-wallet-pro/venus-wallet-configmap.yaml")
	if err != nil {
		return err
	}
	configMap, err := deployer.env.RunConfigMap(ctx, configMapCfg, renderParams)
	if err != nil {
		return err
	}
	deployer.configMapName = configMap.GetName()

	//create deployment
	deployCfg, err := f.Open("venus-wallet-pro/venus-wallet-statefulset.yaml")
	if err != nil {
		return err
	}
	statefulSet, err := deployer.env.RunStatefulSets(ctx, deployCfg, renderParams)
	if err != nil {
		return err
	}
	deployer.statefulSetName = statefulSet.GetName()

	//create service
	svcCfg, err := f.Open("venus-wallet-pro/venus-wallet-headless.yaml")
	if err != nil {
		return err
	}
	svc, err := deployer.env.RunService(ctx, svcCfg, renderParams)
	if err != nil {
		return err
	}
	deployer.svcName = svc.GetName()

	deployer.svcEndpoint, err = deployer.env.WaitForServiceReady(ctx, deployer)
	if err != nil {
		return err
	}
	return nil
}

func (deployer *VenusWalletProDeployer) GetConfig(ctx context.Context) (interface{}, error) {
	cfgData, err := deployer.env.GetConfigMap(ctx, deployer.configMapName, "config.toml")
	if err != nil {
		return nil, err
	}

	cfg := &wConfig.Config{}
	err = toml.Unmarshal(cfgData, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func (deployer *VenusWalletProDeployer) Update(ctx context.Context, updateCfg interface{}) error {
	if updateCfg != nil {
		cfgData, err := toml.Marshal(updateCfg)
		if err != nil {
			return err
		}
		err = deployer.env.SetConfigMap(ctx, deployer.configMapName, "config.toml", cfgData)
		if err != nil {
			return err
		}

		pods, err := deployer.Pods(ctx)
		if err != nil {
			return nil
		}
		for _, pod := range pods {
			_, err = deployer.env.ExecRemoteCmd(ctx, pod.GetName(), "echo", "'"+string(cfgData)+"'", ">", "/root/.venus_wallet_pro/config.toml")
			if err != nil {
				return err
			}
		}
	}

	err := deployer.env.UpdateStatefulSets(ctx, deployer.statefulSetName)
	if err != nil {
		return err
	}
	return nil
}
