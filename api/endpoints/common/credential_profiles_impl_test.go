package common

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRestApiMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "Create credential profile",
			ExpectedVerb: "POST",
			ExpectedPath: "/credential_profiles",
			ExpectedBody: validation.Fixture(t, "requests/create_rest_api_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				credentialProfile := &api.CredentialProfile{
					CredentialType: 3,
					CredentialName: "Creditial profile",
					UserName:       "postman",
					Password:       "test",
				}

				_, err := NewCredentialProfile(c).Create(credentialProfile)
				require.NoError(t, err)
			},
		},
		{
			Name:         "Get Credential profile",
			ExpectedVerb: "GET",
			ExpectedPath: "/credential_profiles/123",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_rest_api_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				credentialProfile, err := NewCredentialProfile(c).Get("123")
				require.NoError(t, err)

				expected := &api.CredentialProfile{
					CredentialType: 3,
					CredentialName: "Creditial profile",
					UserName:       "postman",
					Password:       "test",
				}

				assert.Equal(t, expected, credentialProfile)
			},
		},
		{
			Name:         "Update Credential profile",
			ExpectedVerb: "PUT",
			ExpectedPath: "/credential_profiles/123",
			ExpectedBody: validation.Fixture(t, "requests/update_rest_api_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				credentialProfile := &api.CredentialProfile{
					ID:             "123",
					CredentialType: 3,
					CredentialName: "Creditial profile",
					UserName:       "postman",
					Password:       "test",
				}

				_, err := NewCredentialProfile(c).Update(credentialProfile)
				require.NoError(t, err)
			},
		},
		{
			Name:         "Delete Credential profile",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/credential_profiles/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewCredentialProfile(c).Delete("123"))
			},
		},
	})
}
