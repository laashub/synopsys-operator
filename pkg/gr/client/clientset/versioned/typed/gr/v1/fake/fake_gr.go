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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	grv1 "github.com/blackducksoftware/synopsys-operator/pkg/api/gr/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGrs implements GrInterface
type FakeGrs struct {
	Fake *FakeSynopsysV1
	ns   string
}

var grsResource = schema.GroupVersionResource{Group: "synopsys", Version: "v1", Resource: "grs"}

var grsKind = schema.GroupVersionKind{Group: "synopsys", Version: "v1", Kind: "Gr"}

// Get takes name of the gr, and returns the corresponding gr object, and an error if there is any.
func (c *FakeGrs) Get(name string, options v1.GetOptions) (result *grv1.Gr, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(grsResource, c.ns, name), &grv1.Gr{})

	if obj == nil {
		return nil, err
	}
	return obj.(*grv1.Gr), err
}

// List takes label and field selectors, and returns the list of Grs that match those selectors.
func (c *FakeGrs) List(opts v1.ListOptions) (result *grv1.GrList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(grsResource, grsKind, c.ns, opts), &grv1.GrList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &grv1.GrList{ListMeta: obj.(*grv1.GrList).ListMeta}
	for _, item := range obj.(*grv1.GrList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested grs.
func (c *FakeGrs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(grsResource, c.ns, opts))

}

// Create takes the representation of a gr and creates it.  Returns the server's representation of the gr, and an error, if there is any.
func (c *FakeGrs) Create(gr *grv1.Gr) (result *grv1.Gr, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(grsResource, c.ns, gr), &grv1.Gr{})

	if obj == nil {
		return nil, err
	}
	return obj.(*grv1.Gr), err
}

// Update takes the representation of a gr and updates it. Returns the server's representation of the gr, and an error, if there is any.
func (c *FakeGrs) Update(gr *grv1.Gr) (result *grv1.Gr, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(grsResource, c.ns, gr), &grv1.Gr{})

	if obj == nil {
		return nil, err
	}
	return obj.(*grv1.Gr), err
}

// Delete takes name of the gr and deletes it. Returns an error if one occurs.
func (c *FakeGrs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(grsResource, c.ns, name), &grv1.Gr{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGrs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(grsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &grv1.GrList{})
	return err
}

// Patch applies the patch and returns the patched gr.
func (c *FakeGrs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *grv1.Gr, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(grsResource, c.ns, name, data, subresources...), &grv1.Gr{})

	if obj == nil {
		return nil, err
	}
	return obj.(*grv1.Gr), err
}