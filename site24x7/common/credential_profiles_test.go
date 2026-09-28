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
		CredentialName: "Credential profile",
		UserName:       "UserName",
		Password:       "password",
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
		CredentialName: "Credential profile",
		UserName:       "UserName",
		Password:       "password",
	}

	c.FakeCredentialProfile.On("Update", a).Return(a, nil).Once()

	require.NoError(t, resourceSite24x7CredentialProfileUpdate(d, c))

	c.FakeCredentialProfile.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := resourceSite24x7CredentialProfileUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

// Named for the resource it actually exercises; it was TestRestApiMonitorRead,
// copied from the REST API monitor tests along with the wrong mock and type.
func TestCredentialProfileRead(t *testing.T) {
	d := credentialProfileTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeCredentialProfile.On("Get", "123").Return(&api.CredentialProfile{ID: "123"}, nil).Once()

	require.NoError(t, resourceSite24x7CredentialProfileRead(d, c))

	c.FakeCredentialProfile.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := resourceSite24x7CredentialProfileRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestCredentialProfileDelete(t *testing.T) {
	d := credentialProfileTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeCredentialProfile.On("Delete", "123").Return(nil).Once()

	require.NoError(t, resourceSite24x7CredentialProfileDelete(d, c))

	// A successful delete clears the resource ID, so it has to be restored
	// before exercising the already-deleted case.
	d.SetId("123")

	c.FakeCredentialProfile.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	// This asserts current behaviour, not desired behaviour:
	// resourceSite24x7CredentialProfileDelete returns every error from the API,
	// so destroying a profile that was already removed outside Terraform fails.
	// site24x7_tag, site24x7_monitor_group and the APM resources all treat a 404
	// on delete as success instead. Worth aligning in the resource - a change
	// outside the scope of this test-only pass.
	assert.Equal(t, apierrors.NewStatusError(404, "not found"),
		resourceSite24x7CredentialProfileDelete(d, c))
}

func credentialProfileTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, credentialProfileSchema, map[string]interface{}{
		"credential_name": "Credential profile",
		"credential_type": 3,
		"password":        "password",
		"username":        "UserName",
	})
}
