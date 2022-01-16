package integration

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type PagerDutyIntegration interface {
	Get(pagerDutyIntegrationID string) (*api.PagerDutyIntegration, error)
	Create(pagerDutyIntegration *api.PagerDutyIntegration) (*api.PagerDutyIntegration, error)
	Update(pagerDutyIntegration *api.PagerDutyIntegration) (*api.PagerDutyIntegration, error)
}

type pagerDuty struct {
	client rest.Client
}

func NewPagerDuty(client rest.Client) PagerDutyIntegration {
	return &pagerDuty{
		client: client,
	}
}

func (s *pagerDuty) Get(pagerDutyIntegrationID string) (*api.PagerDutyIntegration, error) {
	pagerDutyIntegration := &api.PagerDutyIntegration{}
	err := s.client.
		Get().
		Resource("integration/pager_duty").
		ResourceID(pagerDutyIntegrationID).
		Do().
		Parse(pagerDutyIntegration)

	return pagerDutyIntegration, err
}

func (s *pagerDuty) Create(pagerDutyIntegration *api.PagerDutyIntegration) (*api.PagerDutyIntegration, error) {
	newPagerDutyIntegration := &api.PagerDutyIntegration{}
	err := s.client.
		Post().
		Resource("integration/pager_duty").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(pagerDutyIntegration).
		Do().
		Parse(newPagerDutyIntegration)

	return newPagerDutyIntegration, err
}

func (s *pagerDuty) Update(pagerDutyIntegration *api.PagerDutyIntegration) (*api.PagerDutyIntegration, error) {
	updatedPagerDutyIntegration := &api.PagerDutyIntegration{}
	err := s.client.
		Put().
		Resource("integration/pager_duty").
		ResourceID(pagerDutyIntegration.ServiceID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(pagerDutyIntegration).
		Do().
		Parse(updatedPagerDutyIntegration)

	return updatedPagerDutyIntegration, err
}
