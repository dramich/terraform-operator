// Code generated by main. DO NOT EDIT.

package fake

import (
	terraformcontrollercattleiov1 "github.com/rancher/terraform-controller/pkg/apis/terraformcontroller.cattle.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeExecutions implements ExecutionInterface
type FakeExecutions struct {
	Fake *FakeTerraformcontrollerV1
	ns   string
}

var executionsResource = schema.GroupVersionResource{Group: "terraformcontroller.cattle.io", Version: "v1", Resource: "executions"}

var executionsKind = schema.GroupVersionKind{Group: "terraformcontroller.cattle.io", Version: "v1", Kind: "Execution"}

// Get takes name of the execution, and returns the corresponding execution object, and an error if there is any.
func (c *FakeExecutions) Get(name string, options v1.GetOptions) (result *terraformcontrollercattleiov1.Execution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(executionsResource, c.ns, name), &terraformcontrollercattleiov1.Execution{})

	if obj == nil {
		return nil, err
	}
	return obj.(*terraformcontrollercattleiov1.Execution), err
}

// List takes label and field selectors, and returns the list of Executions that match those selectors.
func (c *FakeExecutions) List(opts v1.ListOptions) (result *terraformcontrollercattleiov1.ExecutionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(executionsResource, executionsKind, c.ns, opts), &terraformcontrollercattleiov1.ExecutionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &terraformcontrollercattleiov1.ExecutionList{ListMeta: obj.(*terraformcontrollercattleiov1.ExecutionList).ListMeta}
	for _, item := range obj.(*terraformcontrollercattleiov1.ExecutionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested executions.
func (c *FakeExecutions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(executionsResource, c.ns, opts))

}

// Create takes the representation of a execution and creates it.  Returns the server's representation of the execution, and an error, if there is any.
func (c *FakeExecutions) Create(execution *terraformcontrollercattleiov1.Execution) (result *terraformcontrollercattleiov1.Execution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(executionsResource, c.ns, execution), &terraformcontrollercattleiov1.Execution{})

	if obj == nil {
		return nil, err
	}
	return obj.(*terraformcontrollercattleiov1.Execution), err
}

// Update takes the representation of a execution and updates it. Returns the server's representation of the execution, and an error, if there is any.
func (c *FakeExecutions) Update(execution *terraformcontrollercattleiov1.Execution) (result *terraformcontrollercattleiov1.Execution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(executionsResource, c.ns, execution), &terraformcontrollercattleiov1.Execution{})

	if obj == nil {
		return nil, err
	}
	return obj.(*terraformcontrollercattleiov1.Execution), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeExecutions) UpdateStatus(execution *terraformcontrollercattleiov1.Execution) (*terraformcontrollercattleiov1.Execution, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(executionsResource, "status", c.ns, execution), &terraformcontrollercattleiov1.Execution{})

	if obj == nil {
		return nil, err
	}
	return obj.(*terraformcontrollercattleiov1.Execution), err
}

// Delete takes name of the execution and deletes it. Returns an error if one occurs.
func (c *FakeExecutions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(executionsResource, c.ns, name), &terraformcontrollercattleiov1.Execution{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeExecutions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(executionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &terraformcontrollercattleiov1.ExecutionList{})
	return err
}

// Patch applies the patch and returns the patched execution.
func (c *FakeExecutions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *terraformcontrollercattleiov1.Execution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(executionsResource, c.ns, name, pt, data, subresources...), &terraformcontrollercattleiov1.Execution{})

	if obj == nil {
		return nil, err
	}
	return obj.(*terraformcontrollercattleiov1.Execution), err
}