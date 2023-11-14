package controller

import (
	apiv1 "k8s.io/api/core/v1"
)

type Controller struct {
	stopChannel chan struct{}
	knownNodes  []*apiv1.Node
}
