package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/monitors"
	"github.com/stretchr/testify/mock"
)

var _ monitors.DatabricksMonitors = &DatabricksMonitors{}

type DatabricksMonitors struct {
	mock.Mock
}

func (e *DatabricksMonitors) Get(monitorID string) (*api.DatabricksMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.DatabricksMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *DatabricksMonitors) Create(monitor *api.DatabricksMonitor) (*api.DatabricksMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.DatabricksMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *DatabricksMonitors) Update(monitor *api.DatabricksMonitor) (*api.DatabricksMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.DatabricksMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *DatabricksMonitors) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *DatabricksMonitors) List() ([]*api.DatabricksMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.DatabricksMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
