package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/integration"
	"github.com/stretchr/testify/mock"
)

var _ integration.TelegramIntegration = &TelegramIntegration{}

type TelegramIntegration struct {
	mock.Mock
}

func (s *TelegramIntegration) Get(telegramIntegrationID string) (*api.TelegramIntegration, error) {
	args := s.Called(telegramIntegrationID)
	if obj, ok := args.Get(0).(*api.TelegramIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (s *TelegramIntegration) Create(telegramIntegration *api.TelegramIntegration) (*api.TelegramIntegration, error) {
	args := s.Called(telegramIntegration)
	if obj, ok := args.Get(0).(*api.TelegramIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (s *TelegramIntegration) Update(telegramIntegration *api.TelegramIntegration) (*api.TelegramIntegration, error) {
	args := s.Called(telegramIntegration)
	if obj, ok := args.Get(0).(*api.TelegramIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
