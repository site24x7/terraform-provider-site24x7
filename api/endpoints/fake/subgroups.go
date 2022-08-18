package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.Subgroups = &Subgroups{}

type Subgroups struct {
	mock.Mock
}

func (e *Subgroups) Get(groupID string) (*api.Subgroup, error) {
	args := e.Called(groupID)
	if obj, ok := args.Get(0).(*api.Subgroup); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *Subgroups) Create(group *api.Subgroup) (*api.Subgroup, error) {
	args := e.Called(group)
	if obj, ok := args.Get(0).(*api.Subgroup); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *Subgroups) Update(group *api.Subgroup) (*api.Subgroup, error) {
	args := e.Called(group)
	if obj, ok := args.Get(0).(*api.Subgroup); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *Subgroups) Delete(groupID string) error {
	args := e.Called(groupID)
	return args.Error(0)
}

func (e *Subgroups) List() ([]*api.Subgroup, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.Subgroup); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
