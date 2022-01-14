package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/integration"
	"github.com/stretchr/testify/mock"
)

var _ integration.OpsgenieIntegration = &OpsgenieIntegration{}

type OpsgenieIntegration struct {
	mock.Mock
}

func (o *OpsgenieIntegration) Get(opsgenieIntegrationID string) (*api.OpsgenieIntegration, error) {
	args := o.Called(opsgenieIntegrationID)
	if obj, ok := args.Get(0).(*api.OpsgenieIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (o *OpsgenieIntegration) Create(opsgenieIntegration *api.OpsgenieIntegration) (*api.OpsgenieIntegration, error) {
	args := o.Called(opsgenieIntegration)
	if obj, ok := args.Get(0).(*api.OpsgenieIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (o *OpsgenieIntegration) Update(opsgenieIntegration *api.OpsgenieIntegration) (*api.OpsgenieIntegration, error) {
	args := o.Called(opsgenieIntegration)
	if obj, ok := args.Get(0).(*api.OpsgenieIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
