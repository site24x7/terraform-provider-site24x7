package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLocationProfiles(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create location_profile",
			ExpectedVerb: "POST",
			ExpectedPath: "/location_profiles",
			ExpectedBody: validation.Fixture(t, "api/endpoints/testdata/fixtures/requests/create_location_profile.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				profile := &api.LocationProfile{
					ProfileName:                      "Europe_TEST_Profile",
					PrimaryLocation:                  "16",
					RestrictAlternateLocationPolling: false,
					SecondaryLocations: []string{
						"3",
						"2",
						"1",
					},
				}
				_, err := NewLocationProfiles(c).Create(profile)
				require.NoError(t, err)
			},
		},
		{
			Name:         "create location_profile error",
			StatusCode:   500,
			ResponseBody: []byte("whoops"),
			Fn: func(t *testing.T, c rest.Client) {
				_, err := NewLocationProfiles(c).Create(&api.LocationProfile{})
				assert.True(t, apierrors.HasStatusCode(err, 500))
			},
		},
		{
			Name:         "get location_profile",
			ExpectedVerb: "GET",
			ExpectedPath: "/location_profiles/12341234",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "api/endpoints/testdata/fixtures/responses/get_location_profile.json"),
			Fn: func(t *testing.T, c rest.Client) {
				locationProfile, err := NewLocationProfiles(c).Get("12341234")
				require.NoError(t, err)

				expected := &api.LocationProfile{
					ProfileID:                        "7262465",
					ProfileName:                      "Europe_TEST_Profile",
					PrimaryLocation:                  "16",
					RestrictAlternateLocationPolling: false,
					SecondaryLocations: []string{
						"3",
						"2",
						"1",
					},
				}
				assert.Equal(t, expected, locationProfile)
			},
		},
		{
			Name:         "list location_profiles",
			ExpectedVerb: "GET",
			ExpectedPath: "/location_profiles",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "api/endpoints/testdata/fixtures/responses/list_location_profiles.json"),
			Fn: func(t *testing.T, c rest.Client) {
				locationProfiles, err := NewLocationProfiles(c).List()
				require.NoError(t, err)

				expected := []*api.LocationProfile{
					{
						ProfileID:                        "7262465",
						ProfileName:                      "Europe_TEST_Profile",
						PrimaryLocation:                  "16",
						RestrictAlternateLocationPolling: false,
						SecondaryLocations: []string{
							"3",
							"2",
							"1",
						},
					},
					{
						ProfileID:   "123",
						ProfileName: "TEST",
					},
				}

				assert.Equal(t, expected, locationProfiles)
			},
		},
		{
			Name:         "update location_profile",
			ExpectedVerb: "PUT",
			ExpectedPath: "/location_profiles/456",
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, map[string]interface{}{
				"profile_id":   "456",
				"profile_name": "TEST_Profile_update",
			}),
			Fn: func(t *testing.T, c rest.Client) {
				locationProfile := &api.LocationProfile{ProfileID: "456", ProfileName: "TEST_Profile_update"}

				locationProfile, err := NewLocationProfiles(c).Update(locationProfile)
				require.NoError(t, err)

				expected := &api.LocationProfile{
					ProfileID:   "456",
					ProfileName: "TEST_Profile_update",
				}

				assert.Equal(t, expected, locationProfile)
			},
		},
		{
			Name:       "update create_location_profile error",
			StatusCode: 400,
			ResponseBody: validation.JsonBody(t, &api.ErrorResponse{
				ErrorCode: 123,
				Message:   "bad request",
				ErrorInfo: map[string]interface{}{"foo": "bar"},
			}),
			Fn: func(t *testing.T, c rest.Client) {
				_, err := NewLocationProfiles(c).Update(&api.LocationProfile{})
				assert.True(t, apierrors.HasStatusCode(err, 400))
			},
		},
		{
			Name:         "delete location_profile",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/location_profiles/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewLocationProfiles(c).Delete("123"))
			},
		},
		{
			Name:       "delete location_profile not found",
			StatusCode: 404,
			Fn: func(t *testing.T, c rest.Client) {
				err := NewLocationProfiles(c).Delete("123")
				assert.True(t, apierrors.IsNotFound(err))
			},
		},
	})
}
