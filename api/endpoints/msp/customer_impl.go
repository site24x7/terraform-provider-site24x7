package msp

import (
	"github.com/jinzhu/copier"
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type Customers interface {
	Get(customerID string) (*api.Customer, error)
	Create(customer *api.Customer) (*api.Customer, error)
	Update(customer *api.Customer) (*api.Customer, error)
	Delete(customerID string) error
	List() ([]*api.Customer, error)
}

type customers struct {
	client rest.Client
}

func NewCustomers(client rest.Client) Customers {
	return &customers{
		client: client,
	}
}

func (c *customers) Get(customerID string) (*api.Customer, error) {
	customer := &api.Customer{}
	err := c.client.
		Get().
		Resource("msp/customers").
		ResourceID(customerID).
		Do().
		Parse(customer)

	return customer, err
}

func (c *customers) Create(customer *api.Customer) (*api.Customer, error) {
	newCustomer := &api.Customer{}
	err := c.client.
		Post().
		Resource("msp/customers").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		AddHeader("terraform_skip", "true").
		Body(customer).
		Do().
		Parse(newCustomer)

	return newCustomer, err
}

func (c *customers) Update(customer *api.Customer) (*api.Customer, error) {
	updatedCustomer := &api.Customer{}
	customerData := &api.Customer{}
	copier.Copy(customerData, customer)
	customerData.UserID = "" // don’t send ID in payload

	err := c.client.
		Put().
		Resource("msp/customers").
		ResourceID(customer.UserID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(customerData).
		Do().
		Parse(updatedCustomer)

	return updatedCustomer, err
}

func (c *customers) Delete(customerID string) error {
	return nil
}

func (c *customers) List() ([]*api.Customer, error) {
	customers := []*api.Customer{}
	err := c.client.
		Get().
		Resource("msp/customers").
		Do().
		Parse(&customers)

	return customers, err
}
