package monitors

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServerMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "get server monitor",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors/897654345678",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_server_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				serverMonitor, err := NewServerMonitors(c).Get("897654345678")
				require.NoError(t, err)

				expected := &api.ServerMonitor{
					MonitorID:             "897654345678",
					DisplayName:           "foo",
					Type:                  "SERVER",
					NotificationProfileID: "789",
					ThresholdProfileID:    "012",
					MonitorGroups:         []string{"234", "567"},
					UserGroupIDs:          []string{"123", "456"},
				}

				assert.Equal(t, expected, serverMonitor)
			},
		},
		{
			Name:         "list server monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_server_monitors.json"),
			Fn: func(t *testing.T, c rest.Client) {
				serverMonitors, err := NewServerMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.ServerMonitor{
					{
						MonitorID:             "897654345678",
						DisplayName:           "foo",
						Type:                  "SERVER",
						NotificationProfileID: "789",
						ThresholdProfileID:    "012",
						MonitorGroups:         []string{"234", "567"},
						UserGroupIDs:          []string{"123", "456"},
					},
					{
						MonitorID:             "933654345678",
						DisplayName:           "foo",
						Type:                  "SERVER",
						NotificationProfileID: "789",
						ThresholdProfileID:    "012",
						MonitorGroups:         []string{"234", "567"},
						UserGroupIDs:          []string{"123", "456"},
					},
				}

				assert.Equal(t, expected, serverMonitors)
			},
		},
		{
			Name:         "update server monitor",
			ExpectedVerb: "PUT",
			ExpectedPath: "/monitors/123",
			ExpectedBody: validation.Fixture(t, "requests/update_server_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				serverMonitor := &api.ServerMonitor{
					MonitorID:             "123",
					DisplayName:           "foo",
					Type:                  "SERVER",
					HostName:              "abc",
					IPAddress:             "192.12.23.334",
					PerformAutomation:     false,
					LogNeeded:             false,
					PluginModule:          false,
					ITAutomationModule:    false,
					PollInterval:          1,
					TemplateID:            "456",
					ResourceProfileID:     "123",
					NotificationProfileID: "789",
					ThresholdProfileID:    "012",
					MonitorGroups:         []string{"234", "567"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
				}

				_, err := NewServerMonitors(c).Update(serverMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete server monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewServerMonitors(c).Delete("123"))
			},
		},
	})
}
