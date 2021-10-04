package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.WebsiteMonitors = &WebsiteMonitors{}

type WebsiteMonitors struct {
	mock.Mock
}

func (e *WebsiteMonitors) Get(monitorID string) (*api.WebsiteMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.WebsiteMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *WebsiteMonitors) Create(monitor *api.WebsiteMonitor) (*api.WebsiteMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.WebsiteMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *WebsiteMonitors) Update(monitor *api.WebsiteMonitor) (*api.WebsiteMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.WebsiteMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *WebsiteMonitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *WebsiteMonitors) List() ([]*api.WebsiteMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.WebsiteMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *WebsiteMonitors) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *WebsiteMonitors) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
