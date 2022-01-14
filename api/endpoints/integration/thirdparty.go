package integration

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type ThirdpartyIntegrations interface {
	List() ([]*api.ThirdPartyIntegrations, error)
	Delete(integrationID string) error
	Activate(integrationID string) error
	Suspend(integrationID string) error
}

type thirdpartyIntegrations struct {
	client rest.Client
}

func NewThirdpartyIntegrations(client rest.Client) ThirdpartyIntegrations {
	return &thirdpartyIntegrations{
		client: client,
	}
}

func (t *thirdpartyIntegrations) Delete(integrationID string) error {
	return t.client.
		Delete().
		Resource("integration/thirdparty_service").
		ResourceID(integrationID).
		Do().
		Err()
}

func (t *thirdpartyIntegrations) List() ([]*api.ThirdPartyIntegrations, error) {
	thirdPartyIntegrations := []*api.ThirdPartyIntegrations{}
	err := t.client.
		Get().
		Resource("third_party_services").
		Do().
		Parse(&thirdPartyIntegrations)

	return thirdPartyIntegrations, err
}

func (t *thirdpartyIntegrations) Activate(integrationID string) error {
	return t.client.
		Put().
		Resource("integration/thirdparty_service/activate").
		ResourceID(integrationID).
		Do().
		Err()
}

func (t *thirdpartyIntegrations) Suspend(integrationID string) error {
	return t.client.
		Put().
		Resource("integration/thirdparty_service/suspend").
		ResourceID(integrationID).
		Do().
		Err()
}
