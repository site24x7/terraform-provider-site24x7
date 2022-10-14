package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/monitors"
	"github.com/stretchr/testify/mock"
)

var _ monitors.HeartbeatMonitors = &HeartbeatMonitors{}

type HeartbeatMonitors struct {
	mock.Mock
}

func (e *HeartbeatMonitors) Get(monitorID string) (*api.HeartbeatMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.HeartbeatMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *HeartbeatMonitors) Create(monitor *api.HeartbeatMonitor) (*api.HeartbeatMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.HeartbeatMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *HeartbeatMonitors) Update(monitor *api.HeartbeatMonitor) (*api.HeartbeatMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.HeartbeatMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *HeartbeatMonitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *HeartbeatMonitors) List() ([]*api.HeartbeatMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.HeartbeatMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *HeartbeatMonitors) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *HeartbeatMonitors) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
