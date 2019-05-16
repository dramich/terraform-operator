// Code generated by main. DO NOT EDIT.

package versioned

import (
	terraformcontrollerv1 "github.com/rancher/terraform-controller/pkg/generated/clientset/versioned/typed/terraformcontroller.cattle.io/v1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	TerraformcontrollerV1() terraformcontrollerv1.TerraformcontrollerV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Terraformcontroller() terraformcontrollerv1.TerraformcontrollerV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	terraformcontrollerV1 *terraformcontrollerv1.TerraformcontrollerV1Client
}

// TerraformcontrollerV1 retrieves the TerraformcontrollerV1Client
func (c *Clientset) TerraformcontrollerV1() terraformcontrollerv1.TerraformcontrollerV1Interface {
	return c.terraformcontrollerV1
}

// Deprecated: Terraformcontroller retrieves the default version of TerraformcontrollerClient.
// Please explicitly pick a version.
func (c *Clientset) Terraformcontroller() terraformcontrollerv1.TerraformcontrollerV1Interface {
	return c.terraformcontrollerV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.terraformcontrollerV1, err = terraformcontrollerv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.terraformcontrollerV1 = terraformcontrollerv1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.terraformcontrollerV1 = terraformcontrollerv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}