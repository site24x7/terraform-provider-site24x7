package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.MonitorGroups = &MonitorGroups{}

type MonitorGroups struct {
	mock.Mock
}

func (e *MonitorGroups) Get(groupID string) (*api.MonitorGroup, error) {
	args := e.Called(groupID)
	if obj, ok := args.Get(0).(*api.MonitorGroup); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *MonitorGroups) Create(group *api.MonitorGroup) (*api.MonitorGroup, error) {
	args := e.Called(group)
	if obj, ok := args.Get(0).(*api.MonitorGroup); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *MonitorGroups) Update(group *api.MonitorGroup) (*api.MonitorGroup, error) {
	args := e.Called(group)
	if obj, ok := args.Get(0).(*api.MonitorGroup); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *MonitorGroups) Delete(groupID string) error {
	args := e.Called(groupID)
	return args.Error(0)
}

func (e *MonitorGroups) List() ([]*api.MonitorGroup, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.MonitorGroup); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
