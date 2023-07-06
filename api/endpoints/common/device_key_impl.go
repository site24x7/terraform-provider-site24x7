package common

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type DeviceKey interface {
	Get() (*api.DeviceKey, error)
}

type devicekey struct {
	client rest.Client
}

func NewDeviceKey(client rest.Client) DeviceKey {
	return &devicekey{
		client: client,
	}
}

func (c *devicekey) Get() (*api.DeviceKey, error) {
	externalID := &api.DeviceKey{}
	err := c.client.
		Get().
		Resource("device_key").
		Do().
		Parse(externalID)
	return externalID, err
}
