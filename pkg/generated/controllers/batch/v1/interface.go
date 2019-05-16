// Code generated by main. DO NOT EDIT.

package v1

import (
	"github.com/rancher/wrangler/pkg/generic"
	v1 "k8s.io/api/batch/v1"
	informers "k8s.io/client-go/informers/batch/v1"
	clientset "k8s.io/client-go/kubernetes/typed/batch/v1"
)

type Interface interface {
	Job() JobController
}

func New(controllerManager *generic.ControllerManager, client clientset.BatchV1Interface,
	informers informers.Interface) Interface {
	return &version{
		controllerManager: controllerManager,
		client:            client,
		informers:         informers,
	}
}

type version struct {
	controllerManager *generic.ControllerManager
	informers         informers.Interface
	client            clientset.BatchV1Interface
}

func (c *version) Job() JobController {
	return NewJobController(v1.SchemeGroupVersion.WithKind("Job"), c.controllerManager, c.client, c.informers.Jobs())
}