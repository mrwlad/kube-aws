package root

import (
	"github.com/kubernetes-incubator/kube-aws/core/controlplane/config"
	"github.com/kubernetes-incubator/kube-aws/core/root/render"
)

func StackAssetsRendererFromFile(configPath string) (render.StackRenderer, error) {
	cluster, err := config.ClusterFromFile(configPath)
	if err != nil {
		return nil, err
	}
	c, err := cluster.Config()
	if err != nil {
		return nil, err
	}
	return render.NewStackRenderer(c), nil
}

func CredentialsRendererFromFile(configPath string) (render.CredentialsRenderer, error) {
	cluster, err := config.ClusterFromFile(configPath)
	if err != nil {
		return nil, err
	}
	return render.NewCredentialsRenderer(cluster), nil
}
