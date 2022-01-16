package integration

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type WebhookIntegration interface {
	Get(webhookIntegrationID string) (*api.WebhookIntegration, error)
	Create(webhookIntegration *api.WebhookIntegration) (*api.WebhookIntegration, error)
	Update(webhookIntegration *api.WebhookIntegration) (*api.WebhookIntegration, error)
}

type webhook struct {
	client rest.Client
}

func NewWebhook(client rest.Client) WebhookIntegration {
	return &webhook{
		client: client,
	}
}

func (w *webhook) Get(webhookIntegrationID string) (*api.WebhookIntegration, error) {
	webhookIntegration := &api.WebhookIntegration{}
	err := w.client.
		Get().
		Resource("integration/webhooks").
		ResourceID(webhookIntegrationID).
		Do().
		Parse(webhookIntegration)

	return webhookIntegration, err
}

func (w *webhook) Create(webhookIntegration *api.WebhookIntegration) (*api.WebhookIntegration, error) {
	newWebhookIntegration := &api.WebhookIntegration{}
	err := w.client.
		Post().
		Resource("integration/webhooks").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(webhookIntegration).
		Do().
		Parse(newWebhookIntegration)

	return newWebhookIntegration, err
}

func (w *webhook) Update(webhookIntegration *api.WebhookIntegration) (*api.WebhookIntegration, error) {
	updatedWebhookIntegration := &api.WebhookIntegration{}
	err := w.client.
		Put().
		Resource("integration/webhooks").
		ResourceID(webhookIntegration.ServiceID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(webhookIntegration).
		Do().
		Parse(updatedWebhookIntegration)

	return updatedWebhookIntegration, err
}
