package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/monitors"
	"github.com/stretchr/testify/mock"
)

var _ monitors.FTPTransferMonitors = &FTPTransferMonitors{}

type FTPTransferMonitors struct {
	mock.Mock
}

func (e *FTPTransferMonitors) Get(monitorID string) (*api.FTPTransferMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.FTPTransferMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *FTPTransferMonitors) Create(monitor *api.FTPTransferMonitor) (*api.FTPTransferMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.FTPTransferMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *FTPTransferMonitors) Update(monitor *api.FTPTransferMonitor) (*api.FTPTransferMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.FTPTransferMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *FTPTransferMonitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *FTPTransferMonitors) List() ([]*api.FTPTransferMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.FTPTransferMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *FTPTransferMonitors) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *FTPTransferMonitors) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
