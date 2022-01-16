package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/monitors"
	"github.com/stretchr/testify/mock"
)

var _ monitors.AmazonMonitors = &AmazonMonitors{}

type AmazonMonitors struct {
	mock.Mock
}

func (e *AmazonMonitors) Get(monitorID string) (*api.AmazonMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.AmazonMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *AmazonMonitors) Create(monitor *api.AmazonMonitor) (*api.AmazonMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.AmazonMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *AmazonMonitors) Update(monitor *api.AmazonMonitor) (*api.AmazonMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.AmazonMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *AmazonMonitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *AmazonMonitors) List() ([]*api.AmazonMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.AmazonMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *AmazonMonitors) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *AmazonMonitors) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
