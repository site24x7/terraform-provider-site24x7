package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/common"
	"github.com/stretchr/testify/mock"
)

var _ common.SLASetting = &SLASetting{}

type SLASetting struct {
	mock.Mock
}

func (e *SLASetting) Get(slaID string) (*api.SLASetting, error) {
	args := e.Called(slaID)
	if obj, ok := args.Get(0).(*api.SLASetting); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *SLASetting) Create(sla *api.SLASetting) (*api.SLASetting, error) {
	args := e.Called(sla)
	if obj, ok := args.Get(0).(*api.SLASetting); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *SLASetting) Update(sla *api.SLASetting) (*api.SLASetting, error) {
	args := e.Called(sla)
	if obj, ok := args.Get(0).(*api.SLASetting); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *SLASetting) Delete(slaID string) error {
	args := e.Called(slaID)
	return args.Error(0)
}

func (e *SLASetting) List() ([]*api.SLASetting, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.SLASetting); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
