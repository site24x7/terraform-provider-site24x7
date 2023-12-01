package common

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCredentialProfileCreate(t *testing.T) {
	d := credentialProfileTestResourceData(t)

	c := fake.NewClient()

	a := &api.CredentialProfile{
		CredentialType: 3,
		CredentialName: "Creditial profile",
		UserName:       "postman",
		Password:       "test",
	}

	c.FakeCredentialProfile.On("Create", a).Return(a, nil).Once()

	require.NoError(t, resourceSite24x7CredentialProfileCreate(d, c))

	c.FakeCredentialProfile.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := resourceSite24x7CredentialProfileCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestCredentialProfileUpdate(t *testing.T) {
	d := credentialProfileTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.CredentialProfile{
		ID:             "123",
		CredentialType: 3,
		CredentialName: "Creditial profile",
		UserName:       "postman",
		Password:       "test",
	}

	c.FakeRestApiMonitors.On("Update", a).Return(a, nil).Once()

	require.NoError(t, resourceSite24x7CredentialProfileUpdate(d, c))

	c.FakeRestApiMonitors.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := resourceSite24x7CredentialProfileUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestRestApiMonitorRead(t *testing.T) {
	d := credentialProfileTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeRestApiMonitors.On("Get", "123").Return(&api.RestApiMonitor{}, nil).Once()

	require.NoError(t, resourceSite24x7CredentialProfileRead(d, c))

	c.FakeRestApiMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := resourceSite24x7CredentialProfileRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestRestApiMonitorDelete(t *testing.T) {
	d := credentialProfileTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeRestApiMonitors.On("Delete", "123").Return(nil).Once()

	require.NoError(t, resourceSite24x7CredentialProfileDelete(d, c))

	c.FakeRestApiMonitors.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, resourceSite24x7CredentialProfileDelete(d, c))
}

func credentialProfileTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, credentialProfileSchema, map[string]interface{}{
		"credential_name": "Credential profile",
		"credential_type": 3,
		"password":        "password",
		"username":        "UserName",
	})
}
