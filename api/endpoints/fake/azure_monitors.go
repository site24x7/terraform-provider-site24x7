package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/monitors"
	"github.com/stretchr/testify/mock"
)

var _ monitors.AzureMonitors = &AzureMonitors{}

type AzureMonitors struct {
	mock.Mock
}

func (e *AzureMonitors) Get(monitorID string) (*api.AzureMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.AzureMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *AzureMonitors) Create(monitor *api.AzureMonitor) (*api.AzureMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.AzureMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *AzureMonitors) Update(monitor *api.AzureMonitor) (*api.AzureMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.AzureMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *AzureMonitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *AzureMonitors) List() ([]*api.AzureMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.AzureMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *AzureMonitors) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *AzureMonitors) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
