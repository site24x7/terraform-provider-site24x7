package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.NotificationProfiles = &NotificationProfiles{}

type NotificationProfiles struct {
	mock.Mock
}

func (e *NotificationProfiles) Get(profileID string) (*api.NotificationProfile, error) {
	args := e.Called(profileID)
	if obj, ok := args.Get(0).(*api.NotificationProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *NotificationProfiles) Create(profile *api.NotificationProfile) (*api.NotificationProfile, error) {
	args := e.Called(profile)
	if obj, ok := args.Get(0).(*api.NotificationProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *NotificationProfiles) Update(profile *api.NotificationProfile) (*api.NotificationProfile, error) {
	args := e.Called(profile)
	if obj, ok := args.Get(0).(*api.NotificationProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *NotificationProfiles) Delete(profileID string) error {
	args := e.Called(profileID)
	return args.Error(0)
}

func (e *NotificationProfiles) List() ([]*api.NotificationProfile, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.NotificationProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
