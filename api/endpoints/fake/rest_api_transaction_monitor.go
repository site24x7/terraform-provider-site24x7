package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/monitors"
	"github.com/stretchr/testify/mock"
)

var _ monitors.RestApiTransactionMonitors = &RestApiTransactionMonitor{}

type RestApiTransactionMonitor struct {
	mock.Mock
}

func (e *RestApiTransactionMonitor) Get(monitorID string) (*api.RestApiTransactionMonitor, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.RestApiTransactionMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *RestApiTransactionMonitor) GetSteps(monitorID string) (*[]api.Steps, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*[]api.Steps); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *RestApiTransactionMonitor) Create(monitor *api.RestApiTransactionMonitor) (*api.RestApiTransactionMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.RestApiTransactionMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *RestApiTransactionMonitor) Update(monitor *api.RestApiTransactionMonitor) (*api.RestApiTransactionMonitor, error) {
	args := e.Called(monitor)
	if obj, ok := args.Get(0).(*api.RestApiTransactionMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *RestApiTransactionMonitor) Delete(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *RestApiTransactionMonitor) List() ([]*api.RestApiTransactionMonitor, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.RestApiTransactionMonitor); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *RestApiTransactionMonitor) Activate(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}

func (e *RestApiTransactionMonitor) Suspend(monitorID string) error {
	args := e.Called(monitorID)
	return args.Error(0)
}
