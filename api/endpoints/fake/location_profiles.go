package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.LocationProfiles = &LocationProfiles{}

type LocationProfiles struct {
	mock.Mock
}

func (e *LocationProfiles) Get(profileID string) (*api.LocationProfile, error) {
	args := e.Called(profileID)
	if obj, ok := args.Get(0).(*api.LocationProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *LocationProfiles) Create(profile *api.LocationProfile) (*api.LocationProfile, error) {
	args := e.Called(profile)
	if obj, ok := args.Get(0).(*api.LocationProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *LocationProfiles) Update(profile *api.LocationProfile) (*api.LocationProfile, error) {
	args := e.Called(profile)
	if obj, ok := args.Get(0).(*api.LocationProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *LocationProfiles) Delete(profileID string) error {
	args := e.Called(profileID)
	return args.Error(0)
}

func (e *LocationProfiles) List() ([]*api.LocationProfile, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.LocationProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
