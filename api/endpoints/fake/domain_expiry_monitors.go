package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/monitors"
	"github.com/stretchr/testify/mock"
)

var _ monitors.DomainExpiryMonitors = &DomainExpiryMonitors{}

type DomainExpiryMonitors struct {
	mock.Mock
}

func (e *DomainExpiryMonitors) Get(monitorID string) (*api.DomainExpiryMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.DomainExpiryMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *DomainExpiryMonitors) Create(monitor *api.DomainExpiryMonitor) (*api.DomainExpiryMonitor, error) {

	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.DomainExpiryMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *DomainExpiryMonitors) Update(monitor *api.DomainExpiryMonitor) (*api.DomainExpiryMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.DomainExpiryMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *DomainExpiryMonitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *DomainExpiryMonitors) List() ([]*api.DomainExpiryMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.DomainExpiryMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *DomainExpiryMonitors) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *DomainExpiryMonitors) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
