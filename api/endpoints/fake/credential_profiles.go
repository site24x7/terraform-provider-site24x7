package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/common"
	"github.com/stretchr/testify/mock"
)

var _ common.CredentialProfile = &CredentialProfile{}

type CredentialProfile struct {
	mock.Mock
}

func (e *CredentialProfile) Get(credentialProfileID string) (*api.CredentialProfile, error) {
	args := e.Called(credentialProfileID)
	if obj, ok := args.Get(0).(*api.CredentialProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *CredentialProfile) Create(credentialProfile *api.CredentialProfile) (*api.CredentialProfile, error) {
	args := e.Called(credentialProfile)
	if obj, ok := args.Get(0).(*api.CredentialProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *CredentialProfile) Update(credentialProfile *api.CredentialProfile) (*api.CredentialProfile, error) {
	args := e.Called(credentialProfile)
	if obj, ok := args.Get(0).(*api.CredentialProfile); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *CredentialProfile) Delete(credentialProfileID string) error {
	args := e.Called(credentialProfileID)
	return args.Error(0)
}
