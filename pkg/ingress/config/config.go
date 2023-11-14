package config

import "github.com/cuongpiger/vcontainer-ingress-controller/pkg/client"

type (
	Config struct {
		ClusterName string          `mapstructure:"cluster-name"`
		Kubernetes  kubeConfig      `mapstructure:"kubernetes"`
		IAM         client.AuthOpts `mapstructure:"iam"`
		vLB         vLBConfig       `mapstructure:"vlb"`
	}

	// Configuration for connecting to Kubernetes API server, either api_host or kubeconfig should be configured.
	kubeConfig struct {
		// (Optional)Kubernetes API server host address.
		ApiserverHost string `mapstructure:"api-host"`

		// (Optional)Kubeconfig file used to connect to Kubernetes cluster.
		KubeConfig string `mapstructure:"kubeconfig"`
	}

	vLBConfig struct {
		// (Optional) Provider name for the load balancer. Default: amphora.
		Provider string `mapstructure:"provider"`

		// (Required) Subnet ID to create the load balancer.
		SubnetID string `mapstructure:"subnet-id"`

		// (Optional) Public network ID to create floating IP.
		// If empty, no floating IP will be allocated to the load balancer vip.
		FloatingIPNetwork string `mapstructure:"floating-network-id"`

		// (Optional) If the ingress controller should manage the security groups attached to the cluster nodes.
		// Default is false.
		ManageSecurityGroups bool `mapstructure:"manage-security-groups"`

		// (Optional) Flavor ID to create the load balancer.
		// If empty, the default flavor will be used.
		FlavorID string `mapstructure:"flavor-id"`

		// (Optional) If the ingress controller should use serial API calls when creating and updating
		// the load balancer instead of the bulk update API call.
		// Default is false.
		ProviderRequiresSerialAPICalls bool `mapstructure:"provider-requires-serial-api-calls"`
	}
)
