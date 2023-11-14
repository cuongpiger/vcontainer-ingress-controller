package controller

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/cuongpiger/vcontainer-ingress-controller/pkg/ingress/config"
)

// NewController create a new vLB ingress controller
func NewController(conf config.Config) *Controller {
	kubeClient, err := createApiserverClient(conf.Kubernetes.ApiserverHost, conf.Kubernetes.KubeConfig)
	if err != nil {
		log.WithFields(log.Fields{
			"api_server":  conf.Kubernetes.ApiserverHost,
			"kuberconfig": conf.Kubernetes.KubeConfig,
			"error":       err,
		}).Fatal("failed to initialize kubernetes client")
	}

	provider, err := NewVCon
}

// ************************************************* PRIVATE FUNCTIONS *************************************************

func createApiserverClient(apiServerHost string, kubeConfig string) (*kubernetes.Clientset, error) {
	cfg, err := clientcmd.BuildConfigFromFlags(apiServerHost, kubeConfig)
	if err != nil {
		return nil, err
	}

	cfg.QPS = defaultQPS
	cfg.Burst = defaultBurst
	cfg.ContentType = "application/vnd.kubernetes.protobuf"

	log.Debug("creating kubernetes API client")

	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	v, err := client.Discovery().ServerVersion()
	if err != nil {
		return nil, err
	}
	log.WithFields(log.Fields{
		"version": fmt.Sprintf("v%v.%v", v.Major, v.Minor),
	}).Debug("kubernetes API client created")

	return client, nil
}
