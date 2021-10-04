package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.URLAutomations = &URLAutomations{}

type URLAutomations struct {
	mock.Mock
}

func (e *URLAutomations) Get(actionID string) (*api.URLAutomation, error) {
	args := e.Called(actionID)
	if obj, ok := args.Get(0).(*api.URLAutomation); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *URLAutomations) Create(automation *api.URLAutomation) (*api.URLAutomation, error) {
	args := e.Called(automation)
	if obj, ok := args.Get(0).(*api.URLAutomation); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *URLAutomations) Update(automation *api.URLAutomation) (*api.URLAutomation, error) {
	args := e.Called(automation)
	if obj, ok := args.Get(0).(*api.URLAutomation); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *URLAutomations) Delete(actionID string) error {
	args := e.Called(actionID)
	return args.Error(0)
}

func (e *URLAutomations) List() ([]*api.URLAutomation, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.URLAutomation); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
