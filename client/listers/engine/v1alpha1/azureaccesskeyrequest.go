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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubevault.dev/operator/apis/engine/v1alpha1"
)

// AzureAccessKeyRequestLister helps list AzureAccessKeyRequests.
type AzureAccessKeyRequestLister interface {
	// List lists all AzureAccessKeyRequests in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AzureAccessKeyRequest, err error)
	// AzureAccessKeyRequests returns an object that can list and get AzureAccessKeyRequests.
	AzureAccessKeyRequests(namespace string) AzureAccessKeyRequestNamespaceLister
	AzureAccessKeyRequestListerExpansion
}

// azureAccessKeyRequestLister implements the AzureAccessKeyRequestLister interface.
type azureAccessKeyRequestLister struct {
	indexer cache.Indexer
}

// NewAzureAccessKeyRequestLister returns a new AzureAccessKeyRequestLister.
func NewAzureAccessKeyRequestLister(indexer cache.Indexer) AzureAccessKeyRequestLister {
	return &azureAccessKeyRequestLister{indexer: indexer}
}

// List lists all AzureAccessKeyRequests in the indexer.
func (s *azureAccessKeyRequestLister) List(selector labels.Selector) (ret []*v1alpha1.AzureAccessKeyRequest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AzureAccessKeyRequest))
	})
	return ret, err
}

// AzureAccessKeyRequests returns an object that can list and get AzureAccessKeyRequests.
func (s *azureAccessKeyRequestLister) AzureAccessKeyRequests(namespace string) AzureAccessKeyRequestNamespaceLister {
	return azureAccessKeyRequestNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AzureAccessKeyRequestNamespaceLister helps list and get AzureAccessKeyRequests.
type AzureAccessKeyRequestNamespaceLister interface {
	// List lists all AzureAccessKeyRequests in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AzureAccessKeyRequest, err error)
	// Get retrieves the AzureAccessKeyRequest from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AzureAccessKeyRequest, error)
	AzureAccessKeyRequestNamespaceListerExpansion
}

// azureAccessKeyRequestNamespaceLister implements the AzureAccessKeyRequestNamespaceLister
// interface.
type azureAccessKeyRequestNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AzureAccessKeyRequests in the indexer for a given namespace.
func (s azureAccessKeyRequestNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AzureAccessKeyRequest, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AzureAccessKeyRequest))
	})
	return ret, err
}

// Get retrieves the AzureAccessKeyRequest from the indexer for a given namespace and name.
func (s azureAccessKeyRequestNamespaceLister) Get(name string) (*v1alpha1.AzureAccessKeyRequest, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("azureaccesskeyrequest"), name)
	}
	return obj.(*v1alpha1.AzureAccessKeyRequest), nil
}
