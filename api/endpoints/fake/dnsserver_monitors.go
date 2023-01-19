package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/monitors"
	"github.com/stretchr/testify/mock"
)

var _ monitors.WebsiteMonitors = &WebsiteMonitors{}

type DNSServerMonitors struct {
	mock.Mock
}

func (e *DNSServerMonitors) Get(monitorID string) (*api.DNSServerMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.DNSServerMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *DNSServerMonitors) Create(monitor *api.DNSServerMonitor) (*api.DNSServerMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.DNSServerMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *DNSServerMonitors) Update(monitor *api.DNSServerMonitor) (*api.DNSServerMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.DNSServerMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *DNSServerMonitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *DNSServerMonitors) List() ([]*api.DNSServerMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.DNSServerMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *DNSServerMonitors) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *DNSServerMonitors) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
