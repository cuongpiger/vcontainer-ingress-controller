package vcontainer

import (
	"github.com/cuongpiger/vcontainer-ingress-controller/pkg/client"
	"github.com/cuongpiger/vcontainer-ingress-controller/pkg/ingress/config"
	sdkClient "github.com/cuongpiger/vcontainer-sdk/client"
)

func NewProvider(cfg config.Config) (*Provider, error) {
	provider, err := client.NewVContainerClient(&cfg.IAM)
	if err != nil {
		return nil, err
	}

	client2.ServiceClient{}
	vLB, err := sdkClient.ServiceClient{}
}
