package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestThresholdProfiles(t *testing.T) {
	runTests(t, []*endpointTest{
		{
			name:         "create threshold profile",
			expectedVerb: "POST",
			expectedPath: "/threshold_profiles",
			expectedBody: fixture(t, "requests/create_threshold_profile.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
				thresholdProfile := &api.ThresholdProfile{
					ProfileID:              "123",
					ProfileName:            "URL profile",
					Type:                   "URL",
					ProfileType:            0,
					DownLocationThreshold:  8,
					WebsiteContentModified: false,
				}

				_, err := NewThresholdProfiles(c).Create(thresholdProfile)
				require.NoError(t, err)
			},
		},
		{
			name:         "get threshold profile",
			expectedVerb: "GET",
			expectedPath: "/threshold_profiles/123",
			statusCode:   200,
			responseBody: fixture(t, "responses/get_threshold_profile.json"),
			fn: func(t *testing.T, c rest.Client) {
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
			name:         "list threshold profiles",
			expectedVerb: "GET",
			expectedPath: "/threshold_profiles",
			statusCode:   200,
			responseBody: fixture(t, "responses/list_threshold_profiles.json"),
			fn: func(t *testing.T, c rest.Client) {
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
			name:         "update threshold profile",
			expectedVerb: "PUT",
			expectedPath: "/threshold_profiles/123",
			expectedBody: fixture(t, "requests/update_threshold_profile.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
				thresholdProfile := &api.ThresholdProfile{
					ProfileID:              "123",
					ProfileName:            "URL profile",
					Type:                   "URL",
					ProfileType:            0,
					DownLocationThreshold:  8,
					WebsiteContentModified: false,
				}

				_, err := NewThresholdProfiles(c).Update(thresholdProfile)
				require.NoError(t, err)
			},
		},
		{
			name:         "delete threshold profile",
			expectedVerb: "DELETE",
			expectedPath: "/threshold_profiles/123",
			statusCode:   200,
			fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewThresholdProfiles(c).Delete("123"))
			},
		},
	})
}
