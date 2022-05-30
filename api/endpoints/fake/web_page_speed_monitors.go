package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/monitors"
	"github.com/stretchr/testify/mock"
)

var _ monitors.WebPageSpeedMonitors = &WebPageSpeedMonitors{}

type WebPageSpeedMonitors struct {
	mock.Mock
}

func (e *WebPageSpeedMonitors) Get(monitorID string) (*api.WebPageSpeedMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.WebPageSpeedMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *WebPageSpeedMonitors) Create(monitor *api.WebPageSpeedMonitor) (*api.WebPageSpeedMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.WebPageSpeedMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *WebPageSpeedMonitors) Update(monitor *api.WebPageSpeedMonitor) (*api.WebPageSpeedMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.WebPageSpeedMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *WebPageSpeedMonitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *WebPageSpeedMonitors) List() ([]*api.WebPageSpeedMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.WebPageSpeedMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *WebPageSpeedMonitors) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *WebPageSpeedMonitors) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
