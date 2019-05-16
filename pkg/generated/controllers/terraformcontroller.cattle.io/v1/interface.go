// Code generated by main. DO NOT EDIT.

package v1

import (
	v1 "github.com/rancher/terraform-controller/pkg/apis/terraformcontroller.cattle.io/v1"
	clientset "github.com/rancher/terraform-controller/pkg/generated/clientset/versioned/typed/terraformcontroller.cattle.io/v1"
	informers "github.com/rancher/terraform-controller/pkg/generated/informers/externalversions/terraformcontroller.cattle.io/v1"
	"github.com/rancher/wrangler/pkg/generic"
)

type Interface interface {
	Execution() ExecutionController
	ExecutionRun() ExecutionRunController
	Module() ModuleController
}

func New(controllerManager *generic.ControllerManager, client clientset.TerraformcontrollerV1Interface,
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
	client            clientset.TerraformcontrollerV1Interface
}

func (c *version) Execution() ExecutionController {
	return NewExecutionController(v1.SchemeGroupVersion.WithKind("Execution"), c.controllerManager, c.client, c.informers.Executions())
}
func (c *version) ExecutionRun() ExecutionRunController {
	return NewExecutionRunController(v1.SchemeGroupVersion.WithKind("ExecutionRun"), c.controllerManager, c.client, c.informers.ExecutionRuns())
}
func (c *version) Module() ModuleController {
	return NewModuleController(v1.SchemeGroupVersion.WithKind("Module"), c.controllerManager, c.client, c.informers.Modules())
}