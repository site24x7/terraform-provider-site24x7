package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/integration"
	"github.com/stretchr/testify/mock"
)

var _ integration.ServiceNowIntegration = &ServiceNowIntegration{}

type ServiceNowIntegration struct {
	mock.Mock
}

func (s *ServiceNowIntegration) Get(serviceNowIntegrationID string) (*api.ServiceNowIntegration, error) {
	args := s.Called(serviceNowIntegrationID)
	if obj, ok := args.Get(0).(*api.ServiceNowIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (s *ServiceNowIntegration) Create(serviceNowIntegration *api.ServiceNowIntegration) (*api.ServiceNowIntegration, error) {
	args := s.Called(serviceNowIntegration)
	if obj, ok := args.Get(0).(*api.ServiceNowIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (s *ServiceNowIntegration) Update(serviceNowIntegration *api.ServiceNowIntegration) (*api.ServiceNowIntegration, error) {
	args := s.Called(serviceNowIntegration)
	if obj, ok := args.Get(0).(*api.ServiceNowIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
