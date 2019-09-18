// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/xigang/helloworld-operator/pkg/client/clientset/versioned/typed/example.com/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeExampleV1 struct {
	*testing.Fake
}

func (c *FakeExampleV1) Databases(namespace string) v1.DatabaseInterface {
	return &FakeDatabases{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeExampleV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
