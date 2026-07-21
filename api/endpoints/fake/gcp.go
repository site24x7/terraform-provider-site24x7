package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/monitors"
	"github.com/stretchr/testify/mock"
)

var _ monitors.GCPMonitors = &GCPMonitors{}

type GCPMonitors struct {
	mock.Mock
}

func (e *GCPMonitors) Get(monitorID string) (*api.GCPMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.GCPMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *GCPMonitors) Create(monitor *api.GCPMonitor) (*api.GCPMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.GCPMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *GCPMonitors) Update(monitor *api.GCPMonitor) (*api.GCPMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.GCPMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *GCPMonitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *GCPMonitors) List() ([]*api.GCPMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.GCPMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *GCPMonitors) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *GCPMonitors) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
