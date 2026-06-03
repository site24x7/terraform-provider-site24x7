package common

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type OAuth2Provider interface {
	Get(providerID string) (*api.OAuth2Provider, error)
	Create(provider *api.OAuth2Provider) (*api.OAuth2Provider, error)
	Update(provider *api.OAuth2Provider) (*api.OAuth2Provider, error)
	Delete(providerID string) error
	List() ([]*api.OAuth2Provider, error)
}

type oauth2provider struct {
	client rest.Client
}

func NewOAuth2Provider(client rest.Client) OAuth2Provider {
	return &oauth2provider{
		client: client,
	}
}

func (c *oauth2provider) Get(providerID string) (*api.OAuth2Provider, error) {
	provider := &api.OAuth2Provider{}
	err := c.client.
		Get().
		Resource("oauth2_providers").
		ResourceID(providerID).
		Do().
		Parse(provider)

	return provider, err
}

func (c *oauth2provider) Create(provider *api.OAuth2Provider) (*api.OAuth2Provider, error) {
	newProvider := &api.OAuth2Provider{}
	err := c.client.
		Post().
		Resource("oauth2_providers").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(provider).
		Do().
		Parse(newProvider)

	return newProvider, err
}

func (c *oauth2provider) Update(provider *api.OAuth2Provider) (*api.OAuth2Provider, error) {
	updatedProvider := &api.OAuth2Provider{}
	err := c.client.
		Put().
		Resource("oauth2_providers").
		ResourceID(provider.ProviderID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(provider).
		Do().
		Parse(updatedProvider)

	return updatedProvider, err
}

func (c *oauth2provider) Delete(providerID string) error {
	return c.client.
		Delete().
		Resource("oauth2_providers").
		ResourceID(providerID).
		Do().
		Err()
}

func (c *oauth2provider) List() ([]*api.OAuth2Provider, error) {
	providers := []*api.OAuth2Provider{}
	err := c.client.
		Get().
		Resource("oauth2_providers").
		Do().
		Parse(&providers)

	return providers, err
}
