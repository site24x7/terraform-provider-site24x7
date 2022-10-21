package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.URLActions = &URLActions{}

type URLActions struct {
	mock.Mock
}

func (e *URLActions) Get(actionID string) (*api.URLAction, error) {
	args := e.Called(actionID)
	if obj, ok := args.Get(0).(*api.URLAction); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *URLActions) Create(automation *api.URLAction) (*api.URLAction, error) {
	args := e.Called(automation)
	if obj, ok := args.Get(0).(*api.URLAction); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *URLActions) Update(automation *api.URLAction) (*api.URLAction, error) {
	args := e.Called(automation)
	if obj, ok := args.Get(0).(*api.URLAction); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *URLActions) Delete(actionID string) error {
	args := e.Called(actionID)
	return args.Error(0)
}

func (e *URLActions) List() ([]*api.URLAction, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.URLAction); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
