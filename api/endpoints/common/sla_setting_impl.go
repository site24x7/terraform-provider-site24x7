package common

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type SLASetting interface {
	Get(slaID string) (*api.SLASetting, error)
	Create(sla *api.SLASetting) (*api.SLASetting, error)
	Update(sla *api.SLASetting) (*api.SLASetting, error)
	Delete(slaID string) error
	List() ([]*api.SLASetting, error)
}

type slasetting struct {
	client rest.Client
}

func NewSLASetting(client rest.Client) SLASetting {
	return &slasetting{
		client: client,
	}
}

func (c *slasetting) Get(slaID string) (*api.SLASetting, error) {
	sla := &api.SLASetting{}
	err := c.client.
		Get().
		Resource("sla_settings").
		ResourceID(slaID).
		Do().
		Parse(sla)

	return sla, err
}

func (c *slasetting) Create(sla *api.SLASetting) (*api.SLASetting, error) {
	newSLA := &api.SLASetting{}
	err := c.client.
		Post().
		Resource("sla_settings").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(sla).
		Do().
		Parse(newSLA)

	return newSLA, err
}

func (c *slasetting) Update(sla *api.SLASetting) (*api.SLASetting, error) {
	updatedSLA := &api.SLASetting{}
	err := c.client.
		Put().
		Resource("sla_settings").
		ResourceID(sla.SLAID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(sla).
		Do().
		Parse(updatedSLA)

	return updatedSLA, err
}

func (c *slasetting) Delete(slaID string) error {
	return c.client.
		Delete().
		Resource("sla_settings").
		ResourceID(slaID).
		Do().
		Err()
}

func (c *slasetting) List() ([]*api.SLASetting, error) {
	slaSettings := []*api.SLASetting{}
	err := c.client.
		Get().
		Resource("sla_settings").
		Do().
		Parse(&slaSettings)

	return slaSettings, err
}
