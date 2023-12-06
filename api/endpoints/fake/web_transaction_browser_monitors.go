package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/monitors"
	"github.com/stretchr/testify/mock"
)

var _ monitors.WebTransactionBrowserMonitors = &WebTransactionBrowserMonitors{}

//var _ monitors.WebTransactionBrowserMonitorsUpdate = &WebTransactionBrowserMonitorsUpdate{}

type WebTransactionBrowserMonitors struct {
	mock.Mock
}

func (e *WebTransactionBrowserMonitors) Get(monitorID string) (*api.WebTransactionBrowserMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.WebTransactionBrowserMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *WebTransactionBrowserMonitors) Create(monitor *api.WebTransactionBrowserMonitor) (*api.WebTransactionBrowserMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.WebTransactionBrowserMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *WebTransactionBrowserMonitors) Update(monitor *api.WebTransactionBrowserMonitor) (*api.WebTransactionBrowserMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.WebTransactionBrowserMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *WebTransactionBrowserMonitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *WebTransactionBrowserMonitors) List() ([]*api.WebTransactionBrowserMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.WebTransactionBrowserMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *WebTransactionBrowserMonitors) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *WebTransactionBrowserMonitors) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
