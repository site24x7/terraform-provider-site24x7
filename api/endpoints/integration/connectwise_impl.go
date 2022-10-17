package integration

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type ConnectwiseIntegration interface {
	Get(connectwiseIntegrationID string) (*api.ConnectwiseIntegration, error)
	Create(connectwiseIntegration *api.ConnectwiseIntegration) (*api.ConnectwiseIntegration, error)
	Update(connectwiseIntegration *api.ConnectwiseIntegration) (*api.ConnectwiseIntegration, error)
}

type connectwise struct {
	client rest.Client
}

func NewConnectwise(client rest.Client) ConnectwiseIntegration {
	return &connectwise{
		client: client,
	}
}

func (s *connectwise) Get(connectwiseIntegrationID string) (*api.ConnectwiseIntegration, error) {
	connectwiseIntegration := &api.ConnectwiseIntegration{}
	err := s.client.
		Get().
		Resource("integration/connectwise").
		ResourceID(connectwiseIntegrationID).
		Do().
		Parse(connectwiseIntegration)

	return connectwiseIntegration, err
}

func (s *connectwise) Create(connectwiseIntegration *api.ConnectwiseIntegration) (*api.ConnectwiseIntegration, error) {
	newConnectwiseIntegration := &api.ConnectwiseIntegration{}
	err := s.client.
		Post().
		Resource("integration/connectwise").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(connectwiseIntegration).
		Do().
		Parse(newConnectwiseIntegration)

	return newConnectwiseIntegration, err
}

func (s *connectwise) Update(connectwiseIntegration *api.ConnectwiseIntegration) (*api.ConnectwiseIntegration, error) {
	updatedConnectwiseIntegration := &api.ConnectwiseIntegration{}
	err := s.client.
		Put().
		Resource("integration/connectwise").
		ResourceID(connectwiseIntegration.ServiceID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(connectwiseIntegration).
		Do().
		Parse(updatedConnectwiseIntegration)

	return updatedConnectwiseIntegration, err
}
