/*
Copyright 2016 The Fission Authors.

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

package tpr

import (
	"k8s.io/client-go/1.5/pkg/api"
	"k8s.io/client-go/1.5/pkg/watch"
	"k8s.io/client-go/1.5/rest"
)

type (
	KuberneteswatchtriggerInterface interface {
		Create(*Kuberneteswatchtrigger) (*Kuberneteswatchtrigger, error)
		Get(name string) (*Kuberneteswatchtrigger, error)
		Update(*Kuberneteswatchtrigger) (*Kuberneteswatchtrigger, error)
		Delete(name string, options *api.DeleteOptions) error
		List(opts api.ListOptions) (*KuberneteswatchtriggerList, error)
		Watch(opts api.ListOptions) (watch.Interface, error)
	}

	kubernetesWatchTriggerClient struct {
		client    *rest.RESTClient
		namespace string
	}
)

func MakeKuberneteswatchtriggerInterface(tprClient *rest.RESTClient, namespace string) KuberneteswatchtriggerInterface {
	return &kubernetesWatchTriggerClient{
		client:    tprClient,
		namespace: namespace,
	}
}

func (c *kubernetesWatchTriggerClient) Create(obj *Kuberneteswatchtrigger) (*Kuberneteswatchtrigger, error) {
	var result Kuberneteswatchtrigger
	err := c.client.Post().
		Resource("kuberneteswatchtriggers").
		Namespace(c.namespace).
		Body(obj).
		Do().Into(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *kubernetesWatchTriggerClient) Get(name string) (*Kuberneteswatchtrigger, error) {
	var result Kuberneteswatchtrigger
	err := c.client.Get().
		Resource("kuberneteswatchtriggers").
		Namespace(c.namespace).
		Name(name).
		Do().Into(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *kubernetesWatchTriggerClient) Update(obj *Kuberneteswatchtrigger) (*Kuberneteswatchtrigger, error) {
	var result Kuberneteswatchtrigger
	err := c.client.Put().
		Resource("kuberneteswatchtriggers").
		Namespace(c.namespace).
		Name(obj.Metadata.Name).
		Body(obj).
		Do().Into(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *kubernetesWatchTriggerClient) Delete(name string, opts *api.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.namespace).
		Resource("kuberneteswatchtriggers").
		Name(name).
		Body(opts).
		Do().
		Error()
}

func (c *kubernetesWatchTriggerClient) List(opts api.ListOptions) (*KuberneteswatchtriggerList, error) {
	var result KuberneteswatchtriggerList
	err := c.client.Get().
		Namespace(c.namespace).
		Resource("kuberneteswatchtriggers").
		VersionedParams(&opts, api.ParameterCodec).
		Do().
		Into(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *kubernetesWatchTriggerClient) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.client.Get().
		Prefix("watch").
		Namespace(c.namespace).
		Resource("kuberneteswatchtriggers").
		VersionedParams(&opts, api.ParameterCodec).
		Watch()
}
