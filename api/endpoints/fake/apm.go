package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/apm"
	"github.com/stretchr/testify/mock"
)

var _ apm.APMApplications = &APMApplications{}

type APMApplications struct {
	mock.Mock
}

func (e *APMApplications) List(timeWindow string) ([]*api.APMApplication, error) {
	args := e.Called(timeWindow)
	if obj, ok := args.Get(0).([]*api.APMApplication); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *APMApplications) Get(applicationID, timeWindow string) (*api.APMApplication, error) {
	args := e.Called(applicationID, timeWindow)
	if obj, ok := args.Get(0).(*api.APMApplication); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *APMApplications) Manage(applicationID string) error {
	args := e.Called(applicationID)
	return args.Error(0)
}

func (e *APMApplications) Unmanage(applicationID string) error {
	args := e.Called(applicationID)
	return args.Error(0)
}

func (e *APMApplications) Delete(applicationID string) error {
	args := e.Called(applicationID)
	return args.Error(0)
}

var _ apm.APMInstances = &APMInstances{}

type APMInstances struct {
	mock.Mock
}

func (e *APMInstances) List(timeWindow string) ([]*api.APMInstance, error) {
	args := e.Called(timeWindow)
	if obj, ok := args.Get(0).([]*api.APMInstance); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *APMInstances) Get(instanceID, timeWindow string) (*api.APMInstance, error) {
	args := e.Called(instanceID, timeWindow)
	if obj, ok := args.Get(0).(*api.APMInstance); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *APMInstances) Manage(instanceID string) error {
	args := e.Called(instanceID)
	return args.Error(0)
}

func (e *APMInstances) Unmanage(instanceID string) error {
	args := e.Called(instanceID)
	return args.Error(0)
}

func (e *APMInstances) Delete(instanceID string) error {
	args := e.Called(instanceID)
	return args.Error(0)
}

var _ apm.APMAgentConfigProfiles = &APMAgentConfigProfiles{}

type APMAgentConfigProfiles struct {
	mock.Mock
}

func (e *APMAgentConfigProfiles) Get(profileID string) (*api.APMAgentConfigProfile, error) {
	args := e.Called(profileID)
	if obj, ok := args.Get(0).(*api.APMAgentConfigProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *APMAgentConfigProfiles) GetDefault(agentType string) (*api.APMAgentConfigProfile, error) {
	args := e.Called(agentType)
	if obj, ok := args.Get(0).(*api.APMAgentConfigProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *APMAgentConfigProfiles) Create(profile *api.APMAgentConfigProfile) (*api.APMAgentConfigProfile, error) {
	args := e.Called(profile)
	if obj, ok := args.Get(0).(*api.APMAgentConfigProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *APMAgentConfigProfiles) Update(profile *api.APMAgentConfigProfile) (*api.APMAgentConfigProfile, error) {
	args := e.Called(profile)
	if obj, ok := args.Get(0).(*api.APMAgentConfigProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *APMAgentConfigProfiles) Delete(profileID string) error {
	args := e.Called(profileID)
	return args.Error(0)
}

func (e *APMAgentConfigProfiles) List() ([]*api.APMAgentConfigProfile, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.APMAgentConfigProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *APMAgentConfigProfiles) ListByAgentType(agentType string) ([]*api.APMAgentConfigProfile, error) {
	args := e.Called(agentType)
	if obj, ok := args.Get(0).([]*api.APMAgentConfigProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
