package common

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type CredentialProfile interface {
	Get(credentialProfileID string) (*api.CredentialProfile, error)
	Create(credentialProfile *api.CredentialProfile) (*api.CredentialProfile, error)
	Update(credentialProfile *api.CredentialProfile) (*api.CredentialProfile, error)
	Delete(credentialProfileID string) error
	ListWebCredentials() ([]*api.CredentialProfile, error)
}

type credentialprofile struct {
	client rest.Client
}

func NewCredentialProfile(client rest.Client) CredentialProfile {
	return &credentialprofile{
		client: client,
	}
}

func (c *credentialprofile) Get(credentialProfileID string) (*api.CredentialProfile, error) {
	credentialProfile := &api.CredentialProfile{}
	err := c.client.
		Get().
		Resource("credential_profile").
		ResourceID(credentialProfileID).
		Do().
		Parse(credentialProfile)

	return credentialProfile, err
}

func (c *credentialprofile) Create(credentialProfile *api.CredentialProfile) (*api.CredentialProfile, error) {
	newCredentialProfile := &api.CredentialProfile{}
	err := c.client.
		Post().
		Resource("credential_profile").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(credentialProfile).
		Do().
		Parse(newCredentialProfile)

	return newCredentialProfile, err
}

func (c *credentialprofile) Update(credentialProfile *api.CredentialProfile) (*api.CredentialProfile, error) {
	updatedCredentialProfile := &api.CredentialProfile{}
	err := c.client.
		Put().
		Resource("credential_profile").
		ResourceID(credentialProfile.ID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(credentialProfile).
		Do().
		Parse(updatedCredentialProfile)

	return updatedCredentialProfile, err
}

func (c *credentialprofile) Delete(credentialProfileID string) error {
	return c.client.
		Delete().
		Resource("credential_profile").
		ResourceID(credentialProfileID).
		Do().
		Err()
}

func (c *credentialprofile) ListWebCredentials() ([]*api.CredentialProfile, error) {
	credentialProfiles := []*api.CredentialProfile{}
	err := c.client.
		Get().
		Resource("credential_profiles").
		Do().
		Parse(&credentialProfiles)

	return credentialProfiles, err
}
