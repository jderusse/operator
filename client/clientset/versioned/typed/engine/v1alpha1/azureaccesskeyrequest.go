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

// AzureAccessKeyRequestsGetter has a method to return a AzureAccessKeyRequestInterface.
// A group's client should implement this interface.
type AzureAccessKeyRequestsGetter interface {
	AzureAccessKeyRequests(namespace string) AzureAccessKeyRequestInterface
}

// AzureAccessKeyRequestInterface has methods to work with AzureAccessKeyRequest resources.
type AzureAccessKeyRequestInterface interface {
	Create(*v1alpha1.AzureAccessKeyRequest) (*v1alpha1.AzureAccessKeyRequest, error)
	Update(*v1alpha1.AzureAccessKeyRequest) (*v1alpha1.AzureAccessKeyRequest, error)
	UpdateStatus(*v1alpha1.AzureAccessKeyRequest) (*v1alpha1.AzureAccessKeyRequest, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzureAccessKeyRequest, error)
	List(opts v1.ListOptions) (*v1alpha1.AzureAccessKeyRequestList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureAccessKeyRequest, err error)
	AzureAccessKeyRequestExpansion
}

// azureAccessKeyRequests implements AzureAccessKeyRequestInterface
type azureAccessKeyRequests struct {
	client rest.Interface
	ns     string
}

// newAzureAccessKeyRequests returns a AzureAccessKeyRequests
func newAzureAccessKeyRequests(c *EngineV1alpha1Client, namespace string) *azureAccessKeyRequests {
	return &azureAccessKeyRequests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the azureAccessKeyRequest, and returns the corresponding azureAccessKeyRequest object, and an error if there is any.
func (c *azureAccessKeyRequests) Get(name string, options v1.GetOptions) (result *v1alpha1.AzureAccessKeyRequest, err error) {
	result = &v1alpha1.AzureAccessKeyRequest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azureaccesskeyrequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzureAccessKeyRequests that match those selectors.
func (c *azureAccessKeyRequests) List(opts v1.ListOptions) (result *v1alpha1.AzureAccessKeyRequestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzureAccessKeyRequestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azureaccesskeyrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azureAccessKeyRequests.
func (c *azureAccessKeyRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("azureaccesskeyrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azureAccessKeyRequest and creates it.  Returns the server's representation of the azureAccessKeyRequest, and an error, if there is any.
func (c *azureAccessKeyRequests) Create(azureAccessKeyRequest *v1alpha1.AzureAccessKeyRequest) (result *v1alpha1.AzureAccessKeyRequest, err error) {
	result = &v1alpha1.AzureAccessKeyRequest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("azureaccesskeyrequests").
		Body(azureAccessKeyRequest).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azureAccessKeyRequest and updates it. Returns the server's representation of the azureAccessKeyRequest, and an error, if there is any.
func (c *azureAccessKeyRequests) Update(azureAccessKeyRequest *v1alpha1.AzureAccessKeyRequest) (result *v1alpha1.AzureAccessKeyRequest, err error) {
	result = &v1alpha1.AzureAccessKeyRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("azureaccesskeyrequests").
		Name(azureAccessKeyRequest.Name).
		Body(azureAccessKeyRequest).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azureAccessKeyRequests) UpdateStatus(azureAccessKeyRequest *v1alpha1.AzureAccessKeyRequest) (result *v1alpha1.AzureAccessKeyRequest, err error) {
	result = &v1alpha1.AzureAccessKeyRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("azureaccesskeyrequests").
		Name(azureAccessKeyRequest.Name).
		SubResource("status").
		Body(azureAccessKeyRequest).
		Do().
		Into(result)
	return
}

// Delete takes name of the azureAccessKeyRequest and deletes it. Returns an error if one occurs.
func (c *azureAccessKeyRequests) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azureaccesskeyrequests").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azureAccessKeyRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azureaccesskeyrequests").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azureAccessKeyRequest.
func (c *azureAccessKeyRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureAccessKeyRequest, err error) {
	result = &v1alpha1.AzureAccessKeyRequest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("azureaccesskeyrequests").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
