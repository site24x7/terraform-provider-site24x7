package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/monitors"
	"github.com/stretchr/testify/mock"
)

var _ monitors.ServerMonitors = &ServerMonitors{}

type ServerMonitors struct {
	mock.Mock
}

func (e *ServerMonitors) Get(monitorID string) (*api.ServerMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.ServerMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *ServerMonitors) Create(monitor *api.ServerMonitor) (*api.ServerMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.ServerMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *ServerMonitors) Update(monitor *api.ServerMonitor) (*api.ServerMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.ServerMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *ServerMonitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *ServerMonitors) List() ([]*api.ServerMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.ServerMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *ServerMonitors) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *ServerMonitors) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
