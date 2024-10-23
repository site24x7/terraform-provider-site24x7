package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/monitors"
	"github.com/stretchr/testify/mock"
)

var _ monitors.SOAPMonitors = &SOAPMonitors{}

type SOAPMonitors struct {
	mock.Mock
}

func (e *SOAPMonitors) Get(monitorID string) (*api.SOAPMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.SOAPMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *SOAPMonitors) Create(monitor *api.SOAPMonitor) (*api.SOAPMonitor, error) {

	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.SOAPMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *SOAPMonitors) Update(monitor *api.SOAPMonitor) (*api.SOAPMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.SOAPMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *SOAPMonitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *SOAPMonitors) List() ([]*api.SOAPMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.SOAPMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *SOAPMonitors) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *SOAPMonitors) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
