package integration

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type ServiceNowIntegration interface {
	Get(serviceNowIntegrationID string) (*api.ServiceNowIntegration, error)
	Create(serviceNowIntegration *api.ServiceNowIntegration) (*api.ServiceNowIntegration, error)
	Update(serviceNowIntegration *api.ServiceNowIntegration) (*api.ServiceNowIntegration, error)
}

type serviceNow struct {
	client rest.Client
}

func NewServiceNow(client rest.Client) ServiceNowIntegration {
	return &serviceNow{
		client: client,
	}
}

func (s *serviceNow) Get(serviceNowIntegrationID string) (*api.ServiceNowIntegration, error) {
	serviceNowIntegration := &api.ServiceNowIntegration{}
	err := s.client.
		Get().
		Resource("integration/service_now").
		ResourceID(serviceNowIntegrationID).
		Do().
		Parse(serviceNowIntegration)

	return serviceNowIntegration, err
}

func (s *serviceNow) Create(serviceNowIntegration *api.ServiceNowIntegration) (*api.ServiceNowIntegration, error) {
	newServiceNowIntegration := &api.ServiceNowIntegration{}
	err := s.client.
		Post().
		Resource("integration/service_now").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(serviceNowIntegration).
		Do().
		Parse(newServiceNowIntegration)

	return newServiceNowIntegration, err
}

func (s *serviceNow) Update(serviceNowIntegration *api.ServiceNowIntegration) (*api.ServiceNowIntegration, error) {
	updatedServiceNowIntegration := &api.ServiceNowIntegration{}
	err := s.client.
		Put().
		Resource("integration/service_now").
		ResourceID(serviceNowIntegration.ServiceID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(serviceNowIntegration).
		Do().
		Parse(updatedServiceNowIntegration)

	return updatedServiceNowIntegration, err
}
