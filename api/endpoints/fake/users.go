package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.Users = &Users{}

type Users struct {
	mock.Mock
}

func (e *Users) Get(groupID string) (*api.User, error) {
	args := e.Called(groupID)
	if obj, ok := args.Get(0).(*api.User); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *Users) Create(group *api.User) (*api.User, error) {
	args := e.Called(group)
	if obj, ok := args.Get(0).(*api.User); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *Users) Update(group *api.User) (*api.User, error) {
	args := e.Called(group)
	if obj, ok := args.Get(0).(*api.User); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *Users) Delete(groupID string) error {
	args := e.Called(groupID)
	return args.Error(0)
}

func (e *Users) List() ([]*api.User, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.User); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
