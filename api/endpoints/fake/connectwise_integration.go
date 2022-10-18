package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/integration"
	"github.com/stretchr/testify/mock"
)

var _ integration.ConnectwiseIntegration = &ConnectwiseIntegration{}

type ConnectwiseIntegration struct {
	mock.Mock
}

func (s *ConnectwiseIntegration) Get(connectwiseIntegrationID string) (*api.ConnectwiseIntegration, error) {
	args := s.Called(connectwiseIntegrationID)
	if obj, ok := args.Get(0).(*api.ConnectwiseIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (s *ConnectwiseIntegration) Create(connectwiseIntegration *api.ConnectwiseIntegration) (*api.ConnectwiseIntegration, error) {
	args := s.Called(connectwiseIntegration)
	if obj, ok := args.Get(0).(*api.ConnectwiseIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (s *ConnectwiseIntegration) Update(connectwiseIntegration *api.ConnectwiseIntegration) (*api.ConnectwiseIntegration, error) {
	args := s.Called(connectwiseIntegration)
	if obj, ok := args.Get(0).(*api.ConnectwiseIntegration); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
