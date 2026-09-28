package common

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// This file was scaffolded from the REST API monitor tests: it was named
// TestRestApiMonitors, pointed at the monitor fixtures, and expected
// /credential_profiles for every operation. The client uses the singular
// /credential_profile for get, create, update and delete, and the plural only
// for listing - the same split the Site24x7 API uses elsewhere, for example
// apminsight/agent_config_profile versus apminsight/agent_config_profiles.
func TestCredentialProfiles(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "Create credential profile",
			ExpectedVerb: "POST",
			ExpectedPath: "/credential_profile",
			ExpectedBody: validation.Fixture(t, "requests/create_credential_profile.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				credentialProfile := &api.CredentialProfile{
					CredentialType: 3,
					CredentialName: "Credential profile",
					UserName:       "UserName",
					Password:       "password",
				}

				_, err := NewCredentialProfile(c).Create(credentialProfile)
				require.NoError(t, err)
			},
		},
		{
			Name:         "Get Credential profile",
			ExpectedVerb: "GET",
			ExpectedPath: "/credential_profile/123",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_credential_profiles.json"),
			Fn: func(t *testing.T, c rest.Client) {
				credentialProfile, err := NewCredentialProfile(c).Get("123")
				require.NoError(t, err)

				expected := &api.CredentialProfile{
					ID:             "123",
					CredentialType: 3,
					CredentialName: "Credential profile",
					UserName:       "UserName",
					Password:       "password",
				}

				assert.Equal(t, expected, credentialProfile)
			},
		},
		{
			Name:         "Update Credential profile",
			ExpectedVerb: "PUT",
			ExpectedPath: "/credential_profile/123",
			ExpectedBody: validation.Fixture(t, "requests/update_credential_profile.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				credentialProfile := &api.CredentialProfile{
					ID:             "123",
					CredentialType: 3,
					CredentialName: "Credential profile",
					UserName:       "UserName",
					Password:       "password",
				}

				_, err := NewCredentialProfile(c).Update(credentialProfile)
				require.NoError(t, err)
			},
		},
		{
			Name:         "Delete Credential profile",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/credential_profile/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewCredentialProfile(c).Delete("123"))
			},
		},
	})
}
