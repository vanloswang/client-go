package v1

import (
	"github.com/openshift/client-go/authorization/clientset/scheme"
	v1 "github.com/openshift/origin/pkg/authorization/apis/authorization/v1"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type AuthorizationV1Interface interface {
	RESTClient() rest.Interface
	ClusterPoliciesGetter
	ClusterPolicyBindingsGetter
	ClusterRolesGetter
	ClusterRoleBindingsGetter
	LocalResourceAccessReviewsGetter
	LocalSubjectAccessReviewsGetter
	PoliciesGetter
	PolicyBindingsGetter
	ResourceAccessReviewsGetter
	RolesGetter
	RoleBindingsGetter
	RoleBindingRestrictionsGetter
	SelfSubjectRulesReviewsGetter
	SubjectAccessReviewsGetter
	SubjectRulesReviewsGetter
}

// AuthorizationV1Client is used to interact with features provided by the authorization.openshift.io group.
type AuthorizationV1Client struct {
	restClient rest.Interface
}

func (c *AuthorizationV1Client) ClusterPolicies() ClusterPolicyInterface {
	return newClusterPolicies(c)
}

func (c *AuthorizationV1Client) ClusterPolicyBindings() ClusterPolicyBindingInterface {
	return newClusterPolicyBindings(c)
}

func (c *AuthorizationV1Client) ClusterRoles() ClusterRoleInterface {
	return newClusterRoles(c)
}

func (c *AuthorizationV1Client) ClusterRoleBindings() ClusterRoleBindingInterface {
	return newClusterRoleBindings(c)
}

func (c *AuthorizationV1Client) LocalResourceAccessReviews(namespace string) LocalResourceAccessReviewInterface {
	return newLocalResourceAccessReviews(c, namespace)
}

func (c *AuthorizationV1Client) LocalSubjectAccessReviews(namespace string) LocalSubjectAccessReviewInterface {
	return newLocalSubjectAccessReviews(c, namespace)
}

func (c *AuthorizationV1Client) Policies(namespace string) PolicyInterface {
	return newPolicies(c, namespace)
}

func (c *AuthorizationV1Client) PolicyBindings(namespace string) PolicyBindingInterface {
	return newPolicyBindings(c, namespace)
}

func (c *AuthorizationV1Client) ResourceAccessReviews() ResourceAccessReviewInterface {
	return newResourceAccessReviews(c)
}

func (c *AuthorizationV1Client) Roles(namespace string) RoleInterface {
	return newRoles(c, namespace)
}

func (c *AuthorizationV1Client) RoleBindings(namespace string) RoleBindingInterface {
	return newRoleBindings(c, namespace)
}

func (c *AuthorizationV1Client) RoleBindingRestrictions(namespace string) RoleBindingRestrictionInterface {
	return newRoleBindingRestrictions(c, namespace)
}

func (c *AuthorizationV1Client) SelfSubjectRulesReviews(namespace string) SelfSubjectRulesReviewInterface {
	return newSelfSubjectRulesReviews(c, namespace)
}

func (c *AuthorizationV1Client) SubjectAccessReviews() SubjectAccessReviewInterface {
	return newSubjectAccessReviews(c)
}

func (c *AuthorizationV1Client) SubjectRulesReviews(namespace string) SubjectRulesReviewInterface {
	return newSubjectRulesReviews(c, namespace)
}

// NewForConfig creates a new AuthorizationV1Client for the given config.
func NewForConfig(c *rest.Config) (*AuthorizationV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &AuthorizationV1Client{client}, nil
}

// NewForConfigOrDie creates a new AuthorizationV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *AuthorizationV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new AuthorizationV1Client for the given RESTClient.
func New(c rest.Interface) *AuthorizationV1Client {
	return &AuthorizationV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *AuthorizationV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
