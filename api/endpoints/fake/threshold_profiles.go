package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.ThresholdProfiles = &ThresholdProfiles{}

type ThresholdProfiles struct {
	mock.Mock
}

func (e *ThresholdProfiles) Get(profileID string) (*api.ThresholdProfile, error) {
	args := e.Called(profileID)
	if obj, ok := args.Get(0).(*api.ThresholdProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *ThresholdProfiles) Create(profile *api.ThresholdProfile) (*api.ThresholdProfile, error) {
	args := e.Called(profile)
	if obj, ok := args.Get(0).(*api.ThresholdProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *ThresholdProfiles) Update(profile *api.ThresholdProfile) (*api.ThresholdProfile, error) {
	args := e.Called(profile)
	if obj, ok := args.Get(0).(*api.ThresholdProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *ThresholdProfiles) Delete(profileID string) error {
	args := e.Called(profileID)
	return args.Error(0)
}

func (e *ThresholdProfiles) List() ([]*api.ThresholdProfile, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.ThresholdProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
