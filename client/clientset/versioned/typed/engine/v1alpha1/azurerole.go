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

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubevault.dev/operator/apis/engine/v1alpha1"
	scheme "kubevault.dev/operator/client/clientset/versioned/scheme"
)

// AzureRolesGetter has a method to return a AzureRoleInterface.
// A group's client should implement this interface.
type AzureRolesGetter interface {
	AzureRoles(namespace string) AzureRoleInterface
}

// AzureRoleInterface has methods to work with AzureRole resources.
type AzureRoleInterface interface {
	Create(*v1alpha1.AzureRole) (*v1alpha1.AzureRole, error)
	Update(*v1alpha1.AzureRole) (*v1alpha1.AzureRole, error)
	UpdateStatus(*v1alpha1.AzureRole) (*v1alpha1.AzureRole, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzureRole, error)
	List(opts v1.ListOptions) (*v1alpha1.AzureRoleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureRole, err error)
	AzureRoleExpansion
}

// azureRoles implements AzureRoleInterface
type azureRoles struct {
	client rest.Interface
	ns     string
}

// newAzureRoles returns a AzureRoles
func newAzureRoles(c *EngineV1alpha1Client, namespace string) *azureRoles {
	return &azureRoles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the azureRole, and returns the corresponding azureRole object, and an error if there is any.
func (c *azureRoles) Get(name string, options v1.GetOptions) (result *v1alpha1.AzureRole, err error) {
	result = &v1alpha1.AzureRole{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azureroles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzureRoles that match those selectors.
func (c *azureRoles) List(opts v1.ListOptions) (result *v1alpha1.AzureRoleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzureRoleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azureroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azureRoles.
func (c *azureRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("azureroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azureRole and creates it.  Returns the server's representation of the azureRole, and an error, if there is any.
func (c *azureRoles) Create(azureRole *v1alpha1.AzureRole) (result *v1alpha1.AzureRole, err error) {
	result = &v1alpha1.AzureRole{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("azureroles").
		Body(azureRole).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azureRole and updates it. Returns the server's representation of the azureRole, and an error, if there is any.
func (c *azureRoles) Update(azureRole *v1alpha1.AzureRole) (result *v1alpha1.AzureRole, err error) {
	result = &v1alpha1.AzureRole{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("azureroles").
		Name(azureRole.Name).
		Body(azureRole).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azureRoles) UpdateStatus(azureRole *v1alpha1.AzureRole) (result *v1alpha1.AzureRole, err error) {
	result = &v1alpha1.AzureRole{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("azureroles").
		Name(azureRole.Name).
		SubResource("status").
		Body(azureRole).
		Do().
		Into(result)
	return
}

// Delete takes name of the azureRole and deletes it. Returns an error if one occurs.
func (c *azureRoles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azureroles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azureRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azureroles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azureRole.
func (c *azureRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureRole, err error) {
	result = &v1alpha1.AzureRole{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("azureroles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
