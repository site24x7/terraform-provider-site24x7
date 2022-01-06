package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.WebhookIntegration = &WebhookIntegration{}

type WebhookIntegration struct {
	mock.Mock
}

func (w *WebhookIntegration) Get(webhookIntegrationID string) (*api.WebhookIntegration, error) {
	args := w.Called(webhookIntegrationID)
	if obj, ok := args.Get(0).(*api.WebhookIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (w *WebhookIntegration) Create(webhookIntegration *api.WebhookIntegration) (*api.WebhookIntegration, error) {
	args := w.Called(webhookIntegration)
	if obj, ok := args.Get(0).(*api.WebhookIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (w *WebhookIntegration) Update(webhookIntegration *api.WebhookIntegration) (*api.WebhookIntegration, error) {
	args := w.Called(webhookIntegration)
	if obj, ok := args.Get(0).(*api.WebhookIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
