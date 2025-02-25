/*
Copyright 2019 AppsCode Inc.

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
	v1alpha1 "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
)

// FakeAppBindings implements AppBindingInterface
type FakeAppBindings struct {
	Fake *FakeAppcatalogV1alpha1
	ns   string
}

var appbindingsResource = schema.GroupVersionResource{Group: "appcatalog.appscode.com", Version: "v1alpha1", Resource: "appbindings"}

var appbindingsKind = schema.GroupVersionKind{Group: "appcatalog.appscode.com", Version: "v1alpha1", Kind: "AppBinding"}

// Get takes name of the appBinding, and returns the corresponding appBinding object, and an error if there is any.
func (c *FakeAppBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.AppBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(appbindingsResource, c.ns, name), &v1alpha1.AppBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppBinding), err
}

// List takes label and field selectors, and returns the list of AppBindings that match those selectors.
func (c *FakeAppBindings) List(opts v1.ListOptions) (result *v1alpha1.AppBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(appbindingsResource, appbindingsKind, c.ns, opts), &v1alpha1.AppBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AppBindingList{ListMeta: obj.(*v1alpha1.AppBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.AppBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appBindings.
func (c *FakeAppBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(appbindingsResource, c.ns, opts))

}

// Create takes the representation of a appBinding and creates it.  Returns the server's representation of the appBinding, and an error, if there is any.
func (c *FakeAppBindings) Create(appBinding *v1alpha1.AppBinding) (result *v1alpha1.AppBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(appbindingsResource, c.ns, appBinding), &v1alpha1.AppBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppBinding), err
}

// Update takes the representation of a appBinding and updates it. Returns the server's representation of the appBinding, and an error, if there is any.
func (c *FakeAppBindings) Update(appBinding *v1alpha1.AppBinding) (result *v1alpha1.AppBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(appbindingsResource, c.ns, appBinding), &v1alpha1.AppBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppBinding), err
}

// Delete takes name of the appBinding and deletes it. Returns an error if one occurs.
func (c *FakeAppBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(appbindingsResource, c.ns, name), &v1alpha1.AppBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(appbindingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AppBindingList{})
	return err
}

// Patch applies the patch and returns the patched appBinding.
func (c *FakeAppBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(appbindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AppBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppBinding), err
}
