/*
Copyright The Kubernetes Authors.

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

package v1beta1

import (
	v1beta1 "intel/isecl/k8s-custom-controller/v2/crdSchema/api/hostattribute/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HostAttributesCrdLister helps list HostAttributesCrds.
type HostAttributesCrdLister interface {
	// List lists all HostAttributesCrds in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.HostAttributesCrd, err error)
	// HostAttributesCrds returns an object that can list and get HostAttributesCrds.
	HostAttributesCrds(namespace string) HostAttributesCrdNamespaceLister
	HostAttributesCrdListerExpansion
}

// hostAttributesCrdLister implements the HostAttributesCrdLister interface.
type hostAttributesCrdLister struct {
	indexer cache.Indexer
}

// NewHostAttributesCrdLister returns a new HostAttributesCrdLister.
func NewHostAttributesCrdLister(indexer cache.Indexer) HostAttributesCrdLister {
	return &hostAttributesCrdLister{indexer: indexer}
}

// List lists all HostAttributesCrds in the indexer.
func (s *hostAttributesCrdLister) List(selector labels.Selector) (ret []*v1beta1.HostAttributesCrd, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.HostAttributesCrd))
	})
	return ret, err
}

// HostAttributesCrds returns an object that can list and get HostAttributesCrds.
func (s *hostAttributesCrdLister) HostAttributesCrds(namespace string) HostAttributesCrdNamespaceLister {
	return hostAttributesCrdNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HostAttributesCrdNamespaceLister helps list and get HostAttributesCrds.
type HostAttributesCrdNamespaceLister interface {
	// List lists all HostAttributesCrds in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.HostAttributesCrd, err error)
	// Get retrieves the HostAttributesCrd from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.HostAttributesCrd, error)
	HostAttributesCrdNamespaceListerExpansion
}

// hostAttributesCrdNamespaceLister implements the HostAttributesCrdNamespaceLister
// interface.
type hostAttributesCrdNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HostAttributesCrds in the indexer for a given namespace.
func (s hostAttributesCrdNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.HostAttributesCrd, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.HostAttributesCrd))
	})
	return ret, err
}

// Get retrieves the HostAttributesCrd from the indexer for a given namespace and name.
func (s hostAttributesCrdNamespaceLister) Get(name string) (*v1beta1.HostAttributesCrd, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("hostattributescrd"), name)
	}
	return obj.(*v1beta1.HostAttributesCrd), nil
}
