/*
Copyright 2019 The Kube Vault Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubevault.dev/operator/apis/engine/v1alpha1"
)

// FakeAzureRoles implements AzureRoleInterface
type FakeAzureRoles struct {
	Fake *FakeEngineV1alpha1
	ns   string
}

var azurerolesResource = schema.GroupVersionResource{Group: "engine.kubevault.com", Version: "v1alpha1", Resource: "azureroles"}

var azurerolesKind = schema.GroupVersionKind{Group: "engine.kubevault.com", Version: "v1alpha1", Kind: "AzureRole"}

// Get takes name of the azureRole, and returns the corresponding azureRole object, and an error if there is any.
func (c *FakeAzureRoles) Get(name string, options v1.GetOptions) (result *v1alpha1.AzureRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(azurerolesResource, c.ns, name), &v1alpha1.AzureRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureRole), err
}

// List takes label and field selectors, and returns the list of AzureRoles that match those selectors.
func (c *FakeAzureRoles) List(opts v1.ListOptions) (result *v1alpha1.AzureRoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(azurerolesResource, azurerolesKind, c.ns, opts), &v1alpha1.AzureRoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzureRoleList{ListMeta: obj.(*v1alpha1.AzureRoleList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzureRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azureRoles.
func (c *FakeAzureRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(azurerolesResource, c.ns, opts))

}

// Create takes the representation of a azureRole and creates it.  Returns the server's representation of the azureRole, and an error, if there is any.
func (c *FakeAzureRoles) Create(azureRole *v1alpha1.AzureRole) (result *v1alpha1.AzureRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(azurerolesResource, c.ns, azureRole), &v1alpha1.AzureRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureRole), err
}

// Update takes the representation of a azureRole and updates it. Returns the server's representation of the azureRole, and an error, if there is any.
func (c *FakeAzureRoles) Update(azureRole *v1alpha1.AzureRole) (result *v1alpha1.AzureRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(azurerolesResource, c.ns, azureRole), &v1alpha1.AzureRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureRole), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzureRoles) UpdateStatus(azureRole *v1alpha1.AzureRole) (*v1alpha1.AzureRole, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(azurerolesResource, "status", c.ns, azureRole), &v1alpha1.AzureRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureRole), err
}

// Delete takes name of the azureRole and deletes it. Returns an error if one occurs.
func (c *FakeAzureRoles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(azurerolesResource, c.ns, name), &v1alpha1.AzureRole{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(azurerolesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzureRoleList{})
	return err
}

// Patch applies the patch and returns the patched azureRole.
func (c *FakeAzureRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(azurerolesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AzureRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureRole), err
}
