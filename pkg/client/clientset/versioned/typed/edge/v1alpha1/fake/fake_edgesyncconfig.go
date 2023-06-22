/*
Copyright The KubeStellar Authors.

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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/kubestellar/kubestellar/pkg/apis/edge/v1alpha1"
)

// FakeEdgeSyncConfigs implements EdgeSyncConfigInterface
type FakeEdgeSyncConfigs struct {
	Fake *FakeEdgeV1alpha1
}

var edgesyncconfigsResource = schema.GroupVersionResource{Group: "edge.kcp.io", Version: "v1alpha1", Resource: "edgesyncconfigs"}

var edgesyncconfigsKind = schema.GroupVersionKind{Group: "edge.kcp.io", Version: "v1alpha1", Kind: "EdgeSyncConfig"}

// Get takes name of the edgeSyncConfig, and returns the corresponding edgeSyncConfig object, and an error if there is any.
func (c *FakeEdgeSyncConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EdgeSyncConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(edgesyncconfigsResource, name), &v1alpha1.EdgeSyncConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeSyncConfig), err
}

// List takes label and field selectors, and returns the list of EdgeSyncConfigs that match those selectors.
func (c *FakeEdgeSyncConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EdgeSyncConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(edgesyncconfigsResource, edgesyncconfigsKind, opts), &v1alpha1.EdgeSyncConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EdgeSyncConfigList{ListMeta: obj.(*v1alpha1.EdgeSyncConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.EdgeSyncConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested edgeSyncConfigs.
func (c *FakeEdgeSyncConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(edgesyncconfigsResource, opts))
}

// Create takes the representation of a edgeSyncConfig and creates it.  Returns the server's representation of the edgeSyncConfig, and an error, if there is any.
func (c *FakeEdgeSyncConfigs) Create(ctx context.Context, edgeSyncConfig *v1alpha1.EdgeSyncConfig, opts v1.CreateOptions) (result *v1alpha1.EdgeSyncConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(edgesyncconfigsResource, edgeSyncConfig), &v1alpha1.EdgeSyncConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeSyncConfig), err
}

// Update takes the representation of a edgeSyncConfig and updates it. Returns the server's representation of the edgeSyncConfig, and an error, if there is any.
func (c *FakeEdgeSyncConfigs) Update(ctx context.Context, edgeSyncConfig *v1alpha1.EdgeSyncConfig, opts v1.UpdateOptions) (result *v1alpha1.EdgeSyncConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(edgesyncconfigsResource, edgeSyncConfig), &v1alpha1.EdgeSyncConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeSyncConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEdgeSyncConfigs) UpdateStatus(ctx context.Context, edgeSyncConfig *v1alpha1.EdgeSyncConfig, opts v1.UpdateOptions) (*v1alpha1.EdgeSyncConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(edgesyncconfigsResource, "status", edgeSyncConfig), &v1alpha1.EdgeSyncConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeSyncConfig), err
}

// Delete takes name of the edgeSyncConfig and deletes it. Returns an error if one occurs.
func (c *FakeEdgeSyncConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(edgesyncconfigsResource, name, opts), &v1alpha1.EdgeSyncConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEdgeSyncConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(edgesyncconfigsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EdgeSyncConfigList{})
	return err
}

// Patch applies the patch and returns the patched edgeSyncConfig.
func (c *FakeEdgeSyncConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EdgeSyncConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(edgesyncconfigsResource, name, pt, data, subresources...), &v1alpha1.EdgeSyncConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeSyncConfig), err
}
