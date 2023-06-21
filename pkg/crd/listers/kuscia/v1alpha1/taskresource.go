// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TaskResourceLister helps list TaskResources.
// All objects returned here must be treated as read-only.
type TaskResourceLister interface {
	// List lists all TaskResources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TaskResource, err error)
	// TaskResources returns an object that can list and get TaskResources.
	TaskResources(namespace string) TaskResourceNamespaceLister
	TaskResourceListerExpansion
}

// taskResourceLister implements the TaskResourceLister interface.
type taskResourceLister struct {
	indexer cache.Indexer
}

// NewTaskResourceLister returns a new TaskResourceLister.
func NewTaskResourceLister(indexer cache.Indexer) TaskResourceLister {
	return &taskResourceLister{indexer: indexer}
}

// List lists all TaskResources in the indexer.
func (s *taskResourceLister) List(selector labels.Selector) (ret []*v1alpha1.TaskResource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TaskResource))
	})
	return ret, err
}

// TaskResources returns an object that can list and get TaskResources.
func (s *taskResourceLister) TaskResources(namespace string) TaskResourceNamespaceLister {
	return taskResourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TaskResourceNamespaceLister helps list and get TaskResources.
// All objects returned here must be treated as read-only.
type TaskResourceNamespaceLister interface {
	// List lists all TaskResources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TaskResource, err error)
	// Get retrieves the TaskResource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TaskResource, error)
	TaskResourceNamespaceListerExpansion
}

// taskResourceNamespaceLister implements the TaskResourceNamespaceLister
// interface.
type taskResourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TaskResources in the indexer for a given namespace.
func (s taskResourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TaskResource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TaskResource))
	})
	return ret, err
}

// Get retrieves the TaskResource from the indexer for a given namespace and name.
func (s taskResourceNamespaceLister) Get(name string) (*v1alpha1.TaskResource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("taskresource"), name)
	}
	return obj.(*v1alpha1.TaskResource), nil
}