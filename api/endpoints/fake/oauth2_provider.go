package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/common"
	"github.com/stretchr/testify/mock"
)

var _ common.OAuth2Provider = &OAuth2Provider{}

type OAuth2Provider struct {
	mock.Mock
}

func (e *OAuth2Provider) Get(providerID string) (*api.OAuth2Provider, error) {
	args := e.Called(providerID)
	if obj, ok := args.Get(0).(*api.OAuth2Provider); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *OAuth2Provider) Create(provider *api.OAuth2Provider) (*api.OAuth2Provider, error) {
	args := e.Called(provider)
	if obj, ok := args.Get(0).(*api.OAuth2Provider); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *OAuth2Provider) Update(provider *api.OAuth2Provider) (*api.OAuth2Provider, error) {
	args := e.Called(provider)
	if obj, ok := args.Get(0).(*api.OAuth2Provider); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *OAuth2Provider) Delete(providerID string) error {
	args := e.Called(providerID)
	return args.Error(0)
}

func (e *OAuth2Provider) List() ([]*api.OAuth2Provider, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.OAuth2Provider); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
