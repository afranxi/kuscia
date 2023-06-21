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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDataObjects implements DataObjectInterface
type FakeDataObjects struct {
	Fake *FakeKusciaV1alpha1
	ns   string
}

var dataobjectsResource = schema.GroupVersionResource{Group: "kuscia.secretflow", Version: "v1alpha1", Resource: "dataobjects"}

var dataobjectsKind = schema.GroupVersionKind{Group: "kuscia.secretflow", Version: "v1alpha1", Kind: "DataObject"}

// Get takes name of the dataObject, and returns the corresponding dataObject object, and an error if there is any.
func (c *FakeDataObjects) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DataObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dataobjectsResource, c.ns, name), &v1alpha1.DataObject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataObject), err
}

// List takes label and field selectors, and returns the list of DataObjects that match those selectors.
func (c *FakeDataObjects) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DataObjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dataobjectsResource, dataobjectsKind, c.ns, opts), &v1alpha1.DataObjectList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataObjectList{ListMeta: obj.(*v1alpha1.DataObjectList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataObjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataObjects.
func (c *FakeDataObjects) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dataobjectsResource, c.ns, opts))

}

// Create takes the representation of a dataObject and creates it.  Returns the server's representation of the dataObject, and an error, if there is any.
func (c *FakeDataObjects) Create(ctx context.Context, dataObject *v1alpha1.DataObject, opts v1.CreateOptions) (result *v1alpha1.DataObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dataobjectsResource, c.ns, dataObject), &v1alpha1.DataObject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataObject), err
}

// Update takes the representation of a dataObject and updates it. Returns the server's representation of the dataObject, and an error, if there is any.
func (c *FakeDataObjects) Update(ctx context.Context, dataObject *v1alpha1.DataObject, opts v1.UpdateOptions) (result *v1alpha1.DataObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dataobjectsResource, c.ns, dataObject), &v1alpha1.DataObject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataObject), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataObjects) UpdateStatus(ctx context.Context, dataObject *v1alpha1.DataObject, opts v1.UpdateOptions) (*v1alpha1.DataObject, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dataobjectsResource, "status", c.ns, dataObject), &v1alpha1.DataObject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataObject), err
}

// Delete takes name of the dataObject and deletes it. Returns an error if one occurs.
func (c *FakeDataObjects) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(dataobjectsResource, c.ns, name, opts), &v1alpha1.DataObject{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataObjects) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dataobjectsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataObjectList{})
	return err
}

// Patch applies the patch and returns the patched dataObject.
func (c *FakeDataObjects) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DataObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dataobjectsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DataObject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataObject), err
}