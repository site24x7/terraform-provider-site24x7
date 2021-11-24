package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.Tags = &Tags{}

type Tags struct {
	mock.Mock
}

func (e *Tags) Get(groupID string) (*api.Tag, error) {
	args := e.Called(groupID)
	if obj, ok := args.Get(0).(*api.Tag); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *Tags) Create(group *api.Tag) (*api.Tag, error) {
	args := e.Called(group)
	if obj, ok := args.Get(0).(*api.Tag); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *Tags) Update(group *api.Tag) (*api.Tag, error) {
	args := e.Called(group)
	if obj, ok := args.Get(0).(*api.Tag); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *Tags) Delete(groupID string) error {
	args := e.Called(groupID)
	return args.Error(0)
}

func (e *Tags) List() ([]*api.Tag, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.Tag); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
