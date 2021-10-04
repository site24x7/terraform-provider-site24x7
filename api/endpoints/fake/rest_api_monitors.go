package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.RestApiMonitors = &RestApiMonitors{}

type RestApiMonitors struct {
	mock.Mock
}

func (e *RestApiMonitors) Get(monitorID string) (*api.RestApiMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.RestApiMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *RestApiMonitors) Create(monitor *api.RestApiMonitor) (*api.RestApiMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.RestApiMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *RestApiMonitors) Update(monitor *api.RestApiMonitor) (*api.RestApiMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.RestApiMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *RestApiMonitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *RestApiMonitors) List() ([]*api.RestApiMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.RestApiMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *RestApiMonitors) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *RestApiMonitors) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
