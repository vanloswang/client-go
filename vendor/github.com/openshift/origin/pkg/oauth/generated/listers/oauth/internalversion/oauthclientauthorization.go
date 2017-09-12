// This file was automatically generated by lister-gen

package internalversion

import (
	oauth "github.com/openshift/origin/pkg/oauth/apis/oauth"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OAuthClientAuthorizationLister helps list OAuthClientAuthorizations.
type OAuthClientAuthorizationLister interface {
	// List lists all OAuthClientAuthorizations in the indexer.
	List(selector labels.Selector) (ret []*oauth.OAuthClientAuthorization, err error)
	// Get retrieves the OAuthClientAuthorization from the index for a given name.
	Get(name string) (*oauth.OAuthClientAuthorization, error)
	OAuthClientAuthorizationListerExpansion
}

// oAuthClientAuthorizationLister implements the OAuthClientAuthorizationLister interface.
type oAuthClientAuthorizationLister struct {
	indexer cache.Indexer
}

// NewOAuthClientAuthorizationLister returns a new OAuthClientAuthorizationLister.
func NewOAuthClientAuthorizationLister(indexer cache.Indexer) OAuthClientAuthorizationLister {
	return &oAuthClientAuthorizationLister{indexer: indexer}
}

// List lists all OAuthClientAuthorizations in the indexer.
func (s *oAuthClientAuthorizationLister) List(selector labels.Selector) (ret []*oauth.OAuthClientAuthorization, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*oauth.OAuthClientAuthorization))
	})
	return ret, err
}

// Get retrieves the OAuthClientAuthorization from the index for a given name.
func (s *oAuthClientAuthorizationLister) Get(name string) (*oauth.OAuthClientAuthorization, error) {
	key := &oauth.OAuthClientAuthorization{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(oauth.Resource("oauthclientauthorization"), name)
	}
	return obj.(*oauth.OAuthClientAuthorization), nil
}
