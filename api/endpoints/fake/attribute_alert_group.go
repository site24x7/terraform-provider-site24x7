package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/common"
	"github.com/stretchr/testify/mock"
)

var _ common.AttributeAlertGroup = &AttributeAlertGroup{}

type AttributeAlertGroup struct {
	mock.Mock
}

func (e *AttributeAlertGroup) Get(groupID string) (*api.AttributeAlertGroup, error) {
	args := e.Called(groupID)
	if obj, ok := args.Get(0).(*api.AttributeAlertGroup); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *AttributeAlertGroup) Create(group *api.AttributeAlertGroup) (*api.AttributeAlertGroup, error) {
	args := e.Called(group)
	if obj, ok := args.Get(0).(*api.AttributeAlertGroup); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *AttributeAlertGroup) Update(group *api.AttributeAlertGroup) (*api.AttributeAlertGroup, error) {
	args := e.Called(group)
	if obj, ok := args.Get(0).(*api.AttributeAlertGroup); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *AttributeAlertGroup) Delete(groupID string) error {
	args := e.Called(groupID)
	return args.Error(0)
}

func (e *AttributeAlertGroup) List() ([]*api.AttributeAlertGroup, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.AttributeAlertGroup); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
