package fake

import (
	v1 "github.com/openshift/client-go/quota/clientset/typed/quota/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeQuotaV1 struct {
	*testing.Fake
}

func (c *FakeQuotaV1) AppliedClusterResourceQuotas(namespace string) v1.AppliedClusterResourceQuotaInterface {
	return &FakeAppliedClusterResourceQuotas{c, namespace}
}

func (c *FakeQuotaV1) ClusterResourceQuotas() v1.ClusterResourceQuotaInterface {
	return &FakeClusterResourceQuotas{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeQuotaV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}