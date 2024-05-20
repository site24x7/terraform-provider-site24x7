package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/monitors"
	"github.com/stretchr/testify/mock"
)

var _ monitors.CronMonitors = &CronMonitors{}

type CronMonitors struct {
	mock.Mock
}

func (e *CronMonitors) Get(monitorID string) (*api.CronMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.CronMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *CronMonitors) Create(monitor *api.CronMonitor) (*api.CronMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.CronMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *CronMonitors) Update(monitor *api.CronMonitor) (*api.CronMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.CronMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *CronMonitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *CronMonitors) List() ([]*api.CronMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.CronMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *CronMonitors) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *CronMonitors) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
