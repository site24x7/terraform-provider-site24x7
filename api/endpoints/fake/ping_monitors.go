package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/monitors"
	"github.com/stretchr/testify/mock"
)

var _ monitors.PINGMonitors = &PINGMonitors{}

type PINGMonitors struct {
	mock.Mock
}

func (e *PINGMonitors) Get(monitorID string) (*api.PINGMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.PINGMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *PINGMonitors) Create(monitor *api.PINGMonitor) (*api.PINGMonitor, error) {

	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.PINGMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *PINGMonitors) Update(monitor *api.PINGMonitor) (*api.PINGMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.PINGMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *PINGMonitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *PINGMonitors) List() ([]*api.PINGMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.PINGMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *PINGMonitors) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *PINGMonitors) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
