// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	examplecomv1 "github.com/xigang/helloworld-operator/pkg/apis/example.com/v1"
	versioned "github.com/xigang/helloworld-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/xigang/helloworld-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/xigang/helloworld-operator/pkg/client/listers/example.com/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DatabaseInformer provides access to a shared informer and lister for
// Databases.
type DatabaseInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DatabaseLister
}

type databaseInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDatabaseInformer constructs a new informer for Database type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDatabaseInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDatabaseInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDatabaseInformer constructs a new informer for Database type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDatabaseInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1().Databases(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1().Databases(namespace).Watch(options)
			},
		},
		&examplecomv1.Database{},
		resyncPeriod,
		indexers,
	)
}

func (f *databaseInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDatabaseInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *databaseInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&examplecomv1.Database{}, f.defaultInformer)
}

func (f *databaseInformer) Lister() v1.DatabaseLister {
	return v1.NewDatabaseLister(f.Informer().GetIndexer())
}
