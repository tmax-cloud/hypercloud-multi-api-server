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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/tmax-cloud/efk-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// FluentBitConfigurationGetter has a method to return a FluentBitConfigurationInterface.
// A group's client should implement this interface.
type FluentBitConfigurationGetter interface {
	FluentBitConfigurations(namespace string) FluentBitConfigurationInterface
}

// FluentBitConfigurationInterface has methods to work with fluentBitConfiguration resources.
type FluentBitConfigurationInterface interface {
	Create(ctx context.Context, fluentBitConfiguration *v1alpha1.FluentBitConfiguration, opts metav1.CreateOptions) (*v1alpha1.FluentBitConfiguration, error)
	Update(ctx context.Context, fluentBitConfiguration *v1alpha1.FluentBitConfiguration, opts metav1.UpdateOptions) (*v1alpha1.FluentBitConfiguration, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1alpha1.FluentBitConfiguration, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1alpha1.FluentBitConfigurationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.FluentBitConfiguration, err error)
	fluentBitConfigurationExpansion
}

// fluentBitConfigurations implements FluentBitConfigurationInterface
type fluentBitConfigurations struct {
	client rest.Interface
	ns     string
}

// newfluentBitConfigurations returns a fluentBitConfigurations
func newFluentBitConfigurations(c *ConfigV1alpha1Client, namespace string) *fluentBitConfigurations {
	return &fluentBitConfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the fluentBitConfiguration, and returns the corresponding fluentBitConfiguration object, and an error if there is any.
func (c *fluentBitConfigurations) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1alpha1.FluentBitConfiguration, err error) {
	result = &v1alpha1.FluentBitConfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("fluentBitConfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of fluentBitConfigurations that match those selectors.
func (c *fluentBitConfigurations) List(ctx context.Context, opts metav1.ListOptions) (result *v1alpha1.FluentBitConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FluentBitConfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("fluentBitConfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested fluentBitConfigurations.
func (c *fluentBitConfigurations) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("fluentBitConfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a fluentBitConfiguration and creates it.  Returns the server's representation of the fluentBitConfiguration, and an error, if there is any.
func (c *fluentBitConfigurations) Create(ctx context.Context, fluentBitConfiguration *v1alpha1.FluentBitConfiguration, opts metav1.CreateOptions) (result *v1alpha1.FluentBitConfiguration, err error) {
	result = &v1alpha1.FluentBitConfiguration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("fluentBitConfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fluentBitConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a fluentBitConfiguration and updates it. Returns the server's representation of the fluentBitConfiguration, and an error, if there is any.
func (c *fluentBitConfigurations) Update(ctx context.Context, fluentBitConfiguration *v1alpha1.FluentBitConfiguration, opts metav1.UpdateOptions) (result *v1alpha1.FluentBitConfiguration, err error) {
	result = &v1alpha1.FluentBitConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("fluentBitConfigurations").
		Name(fluentBitConfiguration.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fluentBitConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the fluentBitConfiguration and deletes it. Returns an error if one occurs.
func (c *fluentBitConfigurations) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("fluentBitConfigurations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *fluentBitConfigurations) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("fluentBitConfigurations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched fluentBitConfiguration.
func (c *fluentBitConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.FluentBitConfiguration, err error) {
	result = &v1alpha1.FluentBitConfiguration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("fluentBitConfigurations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
