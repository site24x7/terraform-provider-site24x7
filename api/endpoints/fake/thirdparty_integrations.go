package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.ThirdpartyIntegrations = &ThirdPartyIntegrations{}

type ThirdPartyIntegrations struct {
	mock.Mock
}

func (t *ThirdPartyIntegrations) Delete(integrationID string) error {
	args := t.Called(integrationID)
	return args.Error(0)
}

func (t *ThirdPartyIntegrations) List() ([]*api.ThirdPartyIntegrations, error) {
	args := t.Called()
	if obj, ok := args.Get(0).([]*api.ThirdPartyIntegrations); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (t *ThirdPartyIntegrations) Activate(integrationID string) error {
	args := t.Called(integrationID)
	return args.Error(0)
}

func (t *ThirdPartyIntegrations) Suspend(integrationID string) error {
	args := t.Called(integrationID)
	return args.Error(0)
}
