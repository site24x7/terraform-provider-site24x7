package common

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type BusinessHourService interface {
	Get(businessHourID string) (*api.BusinessHour, error)
	Create(businessHour *api.BusinessHour) (*api.BusinessHour, error)
	Update(businessHour *api.BusinessHour) (*api.BusinessHour, error)
	Delete(businessHourID string) error
	List() ([]*api.BusinessHour, error)
}

type BusinessHour struct {
	client rest.Client
}

func NewBusinessHour(client rest.Client) BusinessHourService {
	return &BusinessHour{
		client: client,
	}
}

func (b *BusinessHour) Get(businessHourID string) (*api.BusinessHour, error) {
	businessHour := &api.BusinessHour{}
	err := b.client.
		Get().
		Resource("business_hours").
		ResourceID(businessHourID).
		AddHeader("Accept", "application/json; version=2.0"). // Added Accept header
		Do().
		Parse(businessHour)

	return businessHour, err
}

func (b *BusinessHour) Create(businessHour *api.BusinessHour) (*api.BusinessHour, error) {
	newBusinessHour := &api.BusinessHour{}
	err := b.client.
		Post().
		Resource("business_hours").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		AddHeader("Accept", "application/json; version=2.0"). // Added Accept header
		Body(businessHour).
		Do().
		Parse(newBusinessHour)

	return newBusinessHour, err
}

func (b *BusinessHour) Update(businessHour *api.BusinessHour) (*api.BusinessHour, error) {
	updatedBusinessHour := &api.BusinessHour{}
	err := b.client.
		Put().
		Resource("business_hours").
		ResourceID(businessHour.ID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		AddHeader("Accept", "application/json; version=2.0"). // Added Accept header
		Body(businessHour).
		Do().
		Parse(updatedBusinessHour)

	return updatedBusinessHour, err
}

func (b *BusinessHour) Delete(businessHourID string) error {
	return b.client.
		Delete().
		Resource("business_hours").
		ResourceID(businessHourID).
		Do().
		Err()
}

func (b *BusinessHour) List() ([]*api.BusinessHour, error) {
	businessHourList := []*api.BusinessHour{}
	err := b.client.
		Get().
		Resource("business_hours").
		Do().
		Parse(&businessHourList)

	return businessHourList, err
}
