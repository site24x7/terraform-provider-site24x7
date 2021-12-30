package endpoints

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type SlackIntegration interface {
	Get(slackIntegrationID string) (*api.SlackIntegration, error)
	Create(slackIntegration *api.SlackIntegration) (*api.SlackIntegration, error)
	Update(slackIntegration *api.SlackIntegration) (*api.SlackIntegration, error)
}

type slack struct {
	client rest.Client
}

func NewSlack(client rest.Client) SlackIntegration {
	return &slack{
		client: client,
	}
}

func (s *slack) Get(slackIntegrationID string) (*api.SlackIntegration, error) {
	slackIntegration := &api.SlackIntegration{}
	err := s.client.
		Get().
		Resource("integration/slack").
		ResourceID(slackIntegrationID).
		Do().
		Parse(slackIntegration)

	return slackIntegration, err
}

func (s *slack) Create(slackIntegration *api.SlackIntegration) (*api.SlackIntegration, error) {
	newSlackIntegration := &api.SlackIntegration{}
	err := s.client.
		Post().
		Resource("integration/slack").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(slackIntegration).
		Do().
		Parse(newSlackIntegration)

	return newSlackIntegration, err
}

func (s *slack) Update(slackIntegration *api.SlackIntegration) (*api.SlackIntegration, error) {
	updatedSlackIntegration := &api.SlackIntegration{}
	err := s.client.
		Put().
		Resource("integration/slack").
		ResourceID(slackIntegration.ServiceID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(slackIntegration).
		Do().
		Parse(updatedSlackIntegration)

	return updatedSlackIntegration, err
}
