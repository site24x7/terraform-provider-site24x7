package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLocationProfiles(t *testing.T) {
	runTests(t, []*endpointTest{
		{
			name:         "create location_profile",
			expectedVerb: "POST",
			expectedPath: "/location_profiles",
			expectedBody: fixture(t, "requests/create_location_profile.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
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
			name:         "create location_profile error",
			statusCode:   500,
			responseBody: []byte("whoops"),
			fn: func(t *testing.T, c rest.Client) {
				_, err := NewLocationProfiles(c).Create(&api.LocationProfile{})
				assert.True(t, apierrors.HasStatusCode(err, 500))
			},
		},
		{
			name:         "get location_profile",
			expectedVerb: "GET",
			expectedPath: "/location_profiles/12341234",
			statusCode:   200,
			responseBody: fixture(t, "responses/get_location_profile.json"),
			fn: func(t *testing.T, c rest.Client) {
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
			name:         "list location_profiles",
			expectedVerb: "GET",
			expectedPath: "/location_profiles",
			statusCode:   200,
			responseBody: fixture(t, "responses/list_location_profiles.json"),
			fn: func(t *testing.T, c rest.Client) {
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
			name:         "update location_profile",
			expectedVerb: "PUT",
			expectedPath: "/location_profiles/456",
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, map[string]interface{}{
				"profile_id":   "456",
				"profile_name": "TEST_Profile_update",
			}),
			fn: func(t *testing.T, c rest.Client) {
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
			name:       "update create_location_profile error",
			statusCode: 400,
			responseBody: jsonBody(t, &api.ErrorResponse{
				ErrorCode: 123,
				Message:   "bad request",
				ErrorInfo: map[string]interface{}{"foo": "bar"},
			}),
			fn: func(t *testing.T, c rest.Client) {
				_, err := NewLocationProfiles(c).Update(&api.LocationProfile{})
				assert.True(t, apierrors.HasStatusCode(err, 400))
			},
		},
		{
			name:         "delete location_profile",
			expectedVerb: "DELETE",
			expectedPath: "/location_profiles/123",
			statusCode:   200,
			fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewLocationProfiles(c).Delete("123"))
			},
		},
		{
			name:       "delete location_profile not found",
			statusCode: 404,
			fn: func(t *testing.T, c rest.Client) {
				err := NewLocationProfiles(c).Delete("123")
				assert.True(t, apierrors.IsNotFound(err))
			},
		},
	})
}
