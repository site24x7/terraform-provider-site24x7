package endpoints

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type LocationTemplate interface {
	Get() (*api.LocationTemplate, error)
}

type locationTemplate struct {
	client rest.Client
}

func NewLocationTemplate(client rest.Client) LocationTemplate {
	return &locationTemplate{
		client: client,
	}
}

func (c *locationTemplate) Get() (*api.LocationTemplate, error) {
	template := &api.LocationTemplate{}
	err := c.client.
		Get().
		Resource("location_template").
		Do().
		Parse(&template)

	return template, err
}
