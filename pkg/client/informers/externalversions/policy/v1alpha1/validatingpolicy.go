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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	policyv1alpha1 "github.com/kyverno/kyverno-json/pkg/apis/policy/v1alpha1"
	versioned "github.com/kyverno/kyverno-json/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kyverno/kyverno-json/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kyverno/kyverno-json/pkg/client/listers/policy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ValidatingPolicyInformer provides access to a shared informer and lister for
// ValidatingPolicies.
type ValidatingPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ValidatingPolicyLister
}

type validatingPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewValidatingPolicyInformer constructs a new informer for ValidatingPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewValidatingPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredValidatingPolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredValidatingPolicyInformer constructs a new informer for ValidatingPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredValidatingPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JsonV1alpha1().ValidatingPolicies().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JsonV1alpha1().ValidatingPolicies().Watch(context.TODO(), options)
			},
		},
		&policyv1alpha1.ValidatingPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *validatingPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredValidatingPolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *validatingPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&policyv1alpha1.ValidatingPolicy{}, f.defaultInformer)
}

func (f *validatingPolicyInformer) Lister() v1alpha1.ValidatingPolicyLister {
	return v1alpha1.NewValidatingPolicyLister(f.Informer().GetIndexer())
}