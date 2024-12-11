package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/common"
	"github.com/stretchr/testify/mock"
)

var _ common.BusinessHourService = &BusinessHour{}

// BusinessHour is the mock for the BusinessHourService
type BusinessHour struct {
	mock.Mock
}

// Get retrieves the business hour by ID
func (b *BusinessHour) Get(businessHourID string) (*api.BusinessHour, error) {
	args := b.Called(businessHourID)
	if obj, ok := args.Get(0).(*api.BusinessHour); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

// Create creates a new business hour
func (b *BusinessHour) Create(businessHour *api.BusinessHour) (*api.BusinessHour, error) {
	args := b.Called(businessHour)
	if obj, ok := args.Get(0).(*api.BusinessHour); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

// Update updates an existing business hour
func (b *BusinessHour) Update(businessHour *api.BusinessHour) (*api.BusinessHour, error) {
	args := b.Called(businessHour)
	if obj, ok := args.Get(0).(*api.BusinessHour); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

// Delete deletes a business hour by ID
func (b *BusinessHour) Delete(businessHourID string) error {
	args := b.Called(businessHourID)
	return args.Error(0)
}

// List retrieves all business hours
func (b *BusinessHour) List() ([]*api.BusinessHour, error) {
	args := b.Called()
	if obj, ok := args.Get(0).([]*api.BusinessHour); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
