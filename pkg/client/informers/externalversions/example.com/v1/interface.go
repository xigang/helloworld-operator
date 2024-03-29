// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/xigang/helloworld-operator/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Databases returns a DatabaseInformer.
	Databases() DatabaseInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Databases returns a DatabaseInformer.
func (v *version) Databases() DatabaseInformer {
	return &databaseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
