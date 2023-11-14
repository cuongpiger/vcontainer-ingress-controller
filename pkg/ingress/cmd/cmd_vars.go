package cmd

import (
	"github.com/cuongpiger/vcontainer-ingress-controller/pkg/ingress/config"
	"github.com/cuongpiger/vcontainer-ingress-controller/pkg/ingress/controller"
	"github.com/spf13/cobra"
)

var (
	conf config.Config

	rootCmd = &cobra.Command{
		Use:   "vcontainer-ingress-controller",
		Short: "vContainer Ingress Controller",
		Long:  "vContainer integrated with vLB to provide Ingress functionality for VNG-CLOUD Kubernetes clusters",

		Run: func(cmd *cobra.Command, args []string) {
			vconIngress := controller.NewController(conf)
		},
	}
)
