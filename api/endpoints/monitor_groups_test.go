package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMonitorGroups(t *testing.T) {
	runTests(t, []*endpointTest{
		{
			name:         "create monitor group",
			expectedVerb: "POST",
			expectedPath: "/monitor_groups",
			expectedBody: fixture(t, "requests/create_monitor_group.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
				group := &api.MonitorGroup{
					DisplayName: "foo group",
					Description: "This is foo group",
					Monitors: []string{
						"726000000002460",
						"726000000002464",
					},
					// DependencyResourceID: []string{"123"},
					// SuppressAlert:        true,
					HealthThresholdCount: 10,
				}

				_, err := NewMonitorGroups(c).Create(group)
				require.NoError(t, err)
			},
		},
		{
			name:         "get monitor group",
			expectedVerb: "GET",
			expectedPath: "/monitor_groups/113770000041271035",
			statusCode:   200,
			responseBody: fixture(t, "responses/get_monitor_group.json"),
			fn: func(t *testing.T, c rest.Client) {
				group, err := NewMonitorGroups(c).Get("113770000041271035")
				require.NoError(t, err)

				expected := &api.MonitorGroup{
					GroupID:     "113770000041271035",
					DisplayName: "Group1",
					Description: "Group all IDC monitors.",
					Monitors: []string{
						"726000000002460",
						"726000000002464",
					},
					// DependencyResourceID: []string{"123"},
					// SuppressAlert:        true,
					HealthThresholdCount: 1,
				}

				assert.Equal(t, expected, group)
			},
		},
		{
			name:         "list monitor groups",
			expectedVerb: "GET",
			expectedPath: "/monitor_groups",
			statusCode:   200,
			responseBody: fixture(t, "responses/list_monitor_groups.json"),
			fn: func(t *testing.T, c rest.Client) {
				groups, err := NewMonitorGroups(c).List()
				require.NoError(t, err)

				expected := []*api.MonitorGroup{
					{
						GroupID:     "797300000123437",
						DisplayName: "misc",
						Description: "checks for misc sites",
						Monitors: []string{
							"13370000004999063",
							"79730133704999073",
							"79730000001337083",
							"12340000005000031",
						},
						HealthThresholdCount: 1,
					},
					{
						GroupID:              "79123400003075053",
						DisplayName:          "api",
						HealthThresholdCount: 1,
					},
					{
						GroupID:              "79730456703075223",
						DisplayName:          "web",
						HealthThresholdCount: 1,
					},
				}

				assert.Equal(t, expected, groups)
			},
		},
		// {
		// 	name:         "update monitor group",
		// 	expectedVerb: "PUT",
		// 	expectedPath: "/monitor_groups/123",
		// 	expectedBody: fixture(t, "requests/update_monitor_groups.json"),
		// 	statusCode:   200,
		// 	responseBody: jsonAPIResponseBody(t, nil),
		// 	fn: func(t *testing.T, c rest.Client) {
		// 		group := &api.MonitorGroup{
		// 			GroupID:     "123",
		// 			DisplayName: "foo",
		// 		}

		// 		_, err := NewMonitorGroups(c).Update(group)
		// 		require.NoError(t, err)
		// 	},
		// },
		{
			name:         "delete monitor group",
			expectedVerb: "DELETE",
			expectedPath: "/monitor_groups/123",
			statusCode:   200,
			fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewMonitorGroups(c).Delete("123"))
			},
		},
	})
}
