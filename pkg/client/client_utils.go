package client

import (
	"github.com/cuongpiger/vcontainer-sdk/client"
	"github.com/cuongpiger/vcontainer-sdk/vcontainer"
)

func NewVContainerClient(authOpts *AuthOpts) (*client.ProviderClient, error) {
	provider, _ := vcontainer.NewClient(authOpts.IdentityURL)
	err := vcontainer.Authenticate(provider, authOpts.ToOAuth2Options())

	return provider, err
}
