package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMonitorGroups(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create monitor group",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitor_groups",
			ExpectedBody: validation.Fixture(t, "requests/create_monitor_group.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
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
			Name:         "get monitor group",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitor_groups/113770000041271035",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_monitor_group.json"),
			Fn: func(t *testing.T, c rest.Client) {
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
			Name:         "list monitor groups",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitor_groups",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_monitor_groups.json"),
			Fn: func(t *testing.T, c rest.Client) {
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
		// 	Name:         "update monitor group",
		// 	ExpectedVerb: "PUT",
		// 	ExpectedPath: "/monitor_groups/123",
		// 	ExpectedBody: validation.Fixture(t, "requests/update_monitor_groups.json"),
		// 	StatusCode:   200,
		// 	ResponseBody: validation.JsonAPIResponseBody(t, nil),
		// 	Fn: func(t *testing.T, c rest.Client) {
		// 		group := &api.MonitorGroup{
		// 			GroupID:     "123",
		// 			DisplayName: "foo",
		// 		}

		// 		_, err := NewMonitorGroups(c).Update(group)
		// 		require.NoError(t, err)
		// 	},
		// },
		{
			Name:         "delete monitor group",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitor_groups/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewMonitorGroups(c).Delete("123"))
			},
		},
	})
}
