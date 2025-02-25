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

// FakeAWSRoles implements AWSRoleInterface
type FakeAWSRoles struct {
	Fake *FakeEngineV1alpha1
	ns   string
}

var awsrolesResource = schema.GroupVersionResource{Group: "engine.kubevault.com", Version: "v1alpha1", Resource: "awsroles"}

var awsrolesKind = schema.GroupVersionKind{Group: "engine.kubevault.com", Version: "v1alpha1", Kind: "AWSRole"}

// Get takes name of the aWSRole, and returns the corresponding aWSRole object, and an error if there is any.
func (c *FakeAWSRoles) Get(name string, options v1.GetOptions) (result *v1alpha1.AWSRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsrolesResource, c.ns, name), &v1alpha1.AWSRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSRole), err
}

// List takes label and field selectors, and returns the list of AWSRoles that match those selectors.
func (c *FakeAWSRoles) List(opts v1.ListOptions) (result *v1alpha1.AWSRoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsrolesResource, awsrolesKind, c.ns, opts), &v1alpha1.AWSRoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AWSRoleList{ListMeta: obj.(*v1alpha1.AWSRoleList).ListMeta}
	for _, item := range obj.(*v1alpha1.AWSRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aWSRoles.
func (c *FakeAWSRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsrolesResource, c.ns, opts))

}

// Create takes the representation of a aWSRole and creates it.  Returns the server's representation of the aWSRole, and an error, if there is any.
func (c *FakeAWSRoles) Create(aWSRole *v1alpha1.AWSRole) (result *v1alpha1.AWSRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsrolesResource, c.ns, aWSRole), &v1alpha1.AWSRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSRole), err
}

// Update takes the representation of a aWSRole and updates it. Returns the server's representation of the aWSRole, and an error, if there is any.
func (c *FakeAWSRoles) Update(aWSRole *v1alpha1.AWSRole) (result *v1alpha1.AWSRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsrolesResource, c.ns, aWSRole), &v1alpha1.AWSRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSRole), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAWSRoles) UpdateStatus(aWSRole *v1alpha1.AWSRole) (*v1alpha1.AWSRole, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(awsrolesResource, "status", c.ns, aWSRole), &v1alpha1.AWSRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSRole), err
}

// Delete takes name of the aWSRole and deletes it. Returns an error if one occurs.
func (c *FakeAWSRoles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsrolesResource, c.ns, name), &v1alpha1.AWSRole{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAWSRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsrolesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AWSRoleList{})
	return err
}

// Patch applies the patch and returns the patched aWSRole.
func (c *FakeAWSRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AWSRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsrolesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AWSRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSRole), err
}
