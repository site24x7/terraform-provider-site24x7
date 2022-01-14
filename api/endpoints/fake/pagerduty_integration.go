package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/integration"
	"github.com/stretchr/testify/mock"
)

var _ integration.PagerDutyIntegration = &PagerDutyIntegration{}

type PagerDutyIntegration struct {
	mock.Mock
}

func (s *PagerDutyIntegration) Get(pagerDutyIntegrationID string) (*api.PagerDutyIntegration, error) {
	args := s.Called(pagerDutyIntegrationID)
	if obj, ok := args.Get(0).(*api.PagerDutyIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (s *PagerDutyIntegration) Create(pagerDutyIntegration *api.PagerDutyIntegration) (*api.PagerDutyIntegration, error) {
	args := s.Called(pagerDutyIntegration)
	if obj, ok := args.Get(0).(*api.PagerDutyIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (s *PagerDutyIntegration) Update(pagerDutyIntegration *api.PagerDutyIntegration) (*api.PagerDutyIntegration, error) {
	args := s.Called(pagerDutyIntegration)
	if obj, ok := args.Get(0).(*api.PagerDutyIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
