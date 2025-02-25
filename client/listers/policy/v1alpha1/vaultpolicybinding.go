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
	v1alpha1 "kubevault.dev/operator/apis/policy/v1alpha1"
)

// VaultPolicyBindingLister helps list VaultPolicyBindings.
type VaultPolicyBindingLister interface {
	// List lists all VaultPolicyBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.VaultPolicyBinding, err error)
	// VaultPolicyBindings returns an object that can list and get VaultPolicyBindings.
	VaultPolicyBindings(namespace string) VaultPolicyBindingNamespaceLister
	VaultPolicyBindingListerExpansion
}

// vaultPolicyBindingLister implements the VaultPolicyBindingLister interface.
type vaultPolicyBindingLister struct {
	indexer cache.Indexer
}

// NewVaultPolicyBindingLister returns a new VaultPolicyBindingLister.
func NewVaultPolicyBindingLister(indexer cache.Indexer) VaultPolicyBindingLister {
	return &vaultPolicyBindingLister{indexer: indexer}
}

// List lists all VaultPolicyBindings in the indexer.
func (s *vaultPolicyBindingLister) List(selector labels.Selector) (ret []*v1alpha1.VaultPolicyBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VaultPolicyBinding))
	})
	return ret, err
}

// VaultPolicyBindings returns an object that can list and get VaultPolicyBindings.
func (s *vaultPolicyBindingLister) VaultPolicyBindings(namespace string) VaultPolicyBindingNamespaceLister {
	return vaultPolicyBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VaultPolicyBindingNamespaceLister helps list and get VaultPolicyBindings.
type VaultPolicyBindingNamespaceLister interface {
	// List lists all VaultPolicyBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.VaultPolicyBinding, err error)
	// Get retrieves the VaultPolicyBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.VaultPolicyBinding, error)
	VaultPolicyBindingNamespaceListerExpansion
}

// vaultPolicyBindingNamespaceLister implements the VaultPolicyBindingNamespaceLister
// interface.
type vaultPolicyBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VaultPolicyBindings in the indexer for a given namespace.
func (s vaultPolicyBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VaultPolicyBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VaultPolicyBinding))
	})
	return ret, err
}

// Get retrieves the VaultPolicyBinding from the indexer for a given namespace and name.
func (s vaultPolicyBindingNamespaceLister) Get(name string) (*v1alpha1.VaultPolicyBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vaultpolicybinding"), name)
	}
	return obj.(*v1alpha1.VaultPolicyBinding), nil
}
