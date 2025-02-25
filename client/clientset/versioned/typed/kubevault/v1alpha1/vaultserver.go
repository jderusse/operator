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
	v1alpha1 "kubevault.dev/operator/apis/kubevault/v1alpha1"
	scheme "kubevault.dev/operator/client/clientset/versioned/scheme"
)

// VaultServersGetter has a method to return a VaultServerInterface.
// A group's client should implement this interface.
type VaultServersGetter interface {
	VaultServers(namespace string) VaultServerInterface
}

// VaultServerInterface has methods to work with VaultServer resources.
type VaultServerInterface interface {
	Create(*v1alpha1.VaultServer) (*v1alpha1.VaultServer, error)
	Update(*v1alpha1.VaultServer) (*v1alpha1.VaultServer, error)
	UpdateStatus(*v1alpha1.VaultServer) (*v1alpha1.VaultServer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VaultServer, error)
	List(opts v1.ListOptions) (*v1alpha1.VaultServerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VaultServer, err error)
	VaultServerExpansion
}

// vaultServers implements VaultServerInterface
type vaultServers struct {
	client rest.Interface
	ns     string
}

// newVaultServers returns a VaultServers
func newVaultServers(c *KubevaultV1alpha1Client, namespace string) *vaultServers {
	return &vaultServers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vaultServer, and returns the corresponding vaultServer object, and an error if there is any.
func (c *vaultServers) Get(name string, options v1.GetOptions) (result *v1alpha1.VaultServer, err error) {
	result = &v1alpha1.VaultServer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vaultservers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VaultServers that match those selectors.
func (c *vaultServers) List(opts v1.ListOptions) (result *v1alpha1.VaultServerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VaultServerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vaultservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vaultServers.
func (c *vaultServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vaultservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a vaultServer and creates it.  Returns the server's representation of the vaultServer, and an error, if there is any.
func (c *vaultServers) Create(vaultServer *v1alpha1.VaultServer) (result *v1alpha1.VaultServer, err error) {
	result = &v1alpha1.VaultServer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vaultservers").
		Body(vaultServer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a vaultServer and updates it. Returns the server's representation of the vaultServer, and an error, if there is any.
func (c *vaultServers) Update(vaultServer *v1alpha1.VaultServer) (result *v1alpha1.VaultServer, err error) {
	result = &v1alpha1.VaultServer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vaultservers").
		Name(vaultServer.Name).
		Body(vaultServer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *vaultServers) UpdateStatus(vaultServer *v1alpha1.VaultServer) (result *v1alpha1.VaultServer, err error) {
	result = &v1alpha1.VaultServer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vaultservers").
		Name(vaultServer.Name).
		SubResource("status").
		Body(vaultServer).
		Do().
		Into(result)
	return
}

// Delete takes name of the vaultServer and deletes it. Returns an error if one occurs.
func (c *vaultServers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vaultservers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vaultServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vaultservers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched vaultServer.
func (c *vaultServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VaultServer, err error) {
	result = &v1alpha1.VaultServer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vaultservers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
