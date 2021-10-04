package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.SSLMonitors = &SSLMonitors{}

type SSLMonitors struct {
	mock.Mock
}

func (e *SSLMonitors) Get(monitorID string) (*api.SSLMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.SSLMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *SSLMonitors) Create(monitor *api.SSLMonitor) (*api.SSLMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.SSLMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *SSLMonitors) Update(monitor *api.SSLMonitor) (*api.SSLMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.SSLMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *SSLMonitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *SSLMonitors) List() ([]*api.SSLMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.SSLMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *SSLMonitors) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *SSLMonitors) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
