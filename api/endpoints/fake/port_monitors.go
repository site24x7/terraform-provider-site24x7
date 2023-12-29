package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/monitors"
	"github.com/stretchr/testify/mock"
)

var _ monitors.PortMonitors = &PortMonitors{}

type PortMonitors struct {
	mock.Mock
}

func (e *PortMonitors) Get(monitorID string) (*api.PortMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.PortMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *PortMonitors) Create(monitor *api.PortMonitor) (*api.PortMonitor, error) {

	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.PortMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *PortMonitors) Update(monitor *api.PortMonitor) (*api.PortMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.PortMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *PortMonitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *PortMonitors) List() ([]*api.PortMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.PortMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *PortMonitors) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *PortMonitors) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
