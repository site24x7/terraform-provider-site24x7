package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/monitors"
	"github.com/stretchr/testify/mock"
)

var _ monitors.ISPMonitors = &ISPMonitors{}

type ISPMonitors struct {
	mock.Mock
}

func (e *ISPMonitors) Get(monitorID string) (*api.ISPMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.ISPMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *ISPMonitors) Create(monitor *api.ISPMonitor) (*api.ISPMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.ISPMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *ISPMonitors) Update(monitor *api.ISPMonitor) (*api.ISPMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.ISPMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *ISPMonitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *ISPMonitors) List() ([]*api.ISPMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.ISPMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *ISPMonitors) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *ISPMonitors) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
