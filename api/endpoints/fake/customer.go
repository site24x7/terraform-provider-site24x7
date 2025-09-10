package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/msp"
	"github.com/stretchr/testify/mock"
)

// Ensuring the mock struct implements the interface
var _ msp.Customers = &Customer{}

// Customer is the mock implementation of the CustomerService interface
type Customer struct {
	mock.Mock
}

// Get retrieves a customer by ID
func (c *Customer) Get(customerID string) (*api.Customer, error) {
	args := c.Called(customerID)
	if obj, ok := args.Get(0).(*api.Customer); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

// Create creates a new customer
func (c *Customer) Create(customer *api.Customer) (*api.Customer, error) {
	args := c.Called(customer)
	if obj, ok := args.Get(0).(*api.Customer); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

// Update modifies an existing customer
func (c *Customer) Update(customer *api.Customer) (*api.Customer, error) {
	args := c.Called(customer)
	if obj, ok := args.Get(0).(*api.Customer); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

// Delete removes a customer by ID
func (c *Customer) Delete(customerID string) error {
	return nil
}

// List retrieves all customers
func (c *Customer) List() ([]*api.Customer, error) {
	args := c.Called()
	if obj, ok := args.Get(0).([]*api.Customer); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
