package endpoints

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type MSP interface {
	List() ([]*api.MSPCustomer, error)
}

type msp struct {
	client rest.Client
}

func NewMSP(client rest.Client) MSP {
	return &msp{
		client: client,
	}
}

func (c *msp) List() ([]*api.MSPCustomer, error) {
	mspCustomers := []*api.MSPCustomer{}
	err := c.client.
		Get().
		Resource("short/msp/customers").
		Do().
		Parse(&mspCustomers)
	return mspCustomers, err
}
