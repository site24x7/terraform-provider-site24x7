package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestThresholdProfiles(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create threshold profile",
			ExpectedVerb: "POST",
			ExpectedPath: "/threshold_profiles",
			ExpectedBody: validation.Fixture(t, "requests/create_threshold_profile.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				thresholdProfile := &api.ThresholdProfile{
					ProfileID:             "123",
					ProfileName:           "URL profile",
					Type:                  "URL",
					ProfileType:           0,
					DownLocationThreshold: 8,
				}

				_, err := NewThresholdProfiles(c).Create(thresholdProfile)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get threshold profile",
			ExpectedVerb: "GET",
			ExpectedPath: "/threshold_profiles/123",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_threshold_profile.json"),
			Fn: func(t *testing.T, c rest.Client) {
				thresholdProfile, err := NewThresholdProfiles(c).Get("123")
				require.NoError(t, err)

				expected := &api.ThresholdProfile{
					ProfileID:             "123",
					Type:                  "URL",
					ProfileName:           "URL profile",
					DownLocationThreshold: 8,
				}

				assert.Equal(t, expected, thresholdProfile)
			},
		},
		{
			Name:         "list threshold profiles",
			ExpectedVerb: "GET",
			ExpectedPath: "/threshold_profiles",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_threshold_profiles.json"),
			Fn: func(t *testing.T, c rest.Client) {
				thresholdProfiles, err := NewThresholdProfiles(c).List()
				require.NoError(t, err)

				expected := []*api.ThresholdProfile{
					{
						ProfileID:             "123",
						ProfileName:           "Threshold Profile",
						Type:                  "DNS",
						DownLocationThreshold: 8,
					},
					{
						ProfileID:             "876",
						ProfileName:           "Default",
						Type:                  "URL",
						DownLocationThreshold: 4,
					},
				}

				assert.Equal(t, expected, thresholdProfiles)
			},
		},
		{
			Name:         "update threshold profile",
			ExpectedVerb: "PUT",
			ExpectedPath: "/threshold_profiles/123",
			ExpectedBody: validation.Fixture(t, "requests/update_threshold_profile.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				thresholdProfile := &api.ThresholdProfile{
					ProfileID:             "123",
					ProfileName:           "URL profile",
					Type:                  "URL",
					ProfileType:           0,
					DownLocationThreshold: 8,
				}

				_, err := NewThresholdProfiles(c).Update(thresholdProfile)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete threshold profile",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/threshold_profiles/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewThresholdProfiles(c).Delete("123"))
			},
		},
	})
}
