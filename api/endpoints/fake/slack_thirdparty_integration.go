package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.SlackIntegration = &SlackIntegration{}

type SlackIntegration struct {
	mock.Mock
}

func (s *SlackIntegration) Get(slackIntegrationID string) (*api.SlackIntegration, error) {
	args := s.Called(slackIntegrationID)
	if obj, ok := args.Get(0).(*api.SlackIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (s *SlackIntegration) Create(slackIntegration *api.SlackIntegration) (*api.SlackIntegration, error) {
	args := s.Called(slackIntegration)
	if obj, ok := args.Get(0).(*api.SlackIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (s *SlackIntegration) Update(slackIntegration *api.SlackIntegration) (*api.SlackIntegration, error) {
	args := s.Called(slackIntegration)
	if obj, ok := args.Get(0).(*api.SlackIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
