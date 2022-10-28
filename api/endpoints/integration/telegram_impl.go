package integration

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type TelegramIntegration interface {
	Get(telegramIntegrationID string) (*api.TelegramIntegration, error)
	Create(telegramIntegration *api.TelegramIntegration) (*api.TelegramIntegration, error)
	Update(telegramIntegration *api.TelegramIntegration) (*api.TelegramIntegration, error)
}

type telegram struct {
	client rest.Client
}

func NewTelegram(client rest.Client) TelegramIntegration {
	return &telegram{
		client: client,
	}
}

func (s *telegram) Get(telegramIntegrationID string) (*api.TelegramIntegration, error) {
	telegramIntegration := &api.TelegramIntegration{}
	err := s.client.
		Get().
		Resource("integration/telegram").
		ResourceID(telegramIntegrationID).
		Do().
		Parse(telegramIntegration)

	return telegramIntegration, err
}

func (s *telegram) Create(telegramIntegration *api.TelegramIntegration) (*api.TelegramIntegration, error) {
	newTelegramIntegration := &api.TelegramIntegration{}
	err := s.client.
		Post().
		Resource("integration/telegram").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(telegramIntegration).
		Do().
		Parse(newTelegramIntegration)

	return newTelegramIntegration, err
}

func (s *telegram) Update(telegramIntegration *api.TelegramIntegration) (*api.TelegramIntegration, error) {
	updatedTelegramIntegration := &api.TelegramIntegration{}
	err := s.client.
		Put().
		Resource("integration/telegram").
		ResourceID(telegramIntegration.ServiceID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(telegramIntegration).
		Do().
		Parse(updatedTelegramIntegration)

	return updatedTelegramIntegration, err
}
