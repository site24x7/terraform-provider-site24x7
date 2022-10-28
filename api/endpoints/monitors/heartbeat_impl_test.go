package monitors

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeartbeatMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create heartbeat monitor",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitors",
			ExpectedBody: validation.Fixture(t, "requests/create_heartbeat_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				heartbeatMonitor := &api.HeartbeatMonitor{
					DisplayName:           "foo",
					NameInPingURL:         "status_check",
					Type:                  "HEARTBEAT",
					ThresholdProfileID:    "012",
					NotificationProfileID: "789",
					MonitorGroups:         []string{"234", "567"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
					ThirdPartyServiceIDs:  []string{"123", "456"},
					OnCallScheduleID:      "1244",
				}

				_, err := NewHeartbeatMonitors(c).Create(heartbeatMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get heartbeat monitor",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors/897654345678",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_heartbeat_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				heartbeatMonitor, err := NewHeartbeatMonitors(c).Get("897654345678")
				require.NoError(t, err)

				expected := &api.HeartbeatMonitor{
					MonitorID:             "897654345678",
					DisplayName:           "foo",
					NameInPingURL:         "status_check",
					Type:                  "HEARTBEAT",
					ThresholdProfileID:    "012",
					NotificationProfileID: "789",
					MonitorGroups:         []string{"234", "567"},
					UserGroupIDs:          []string{"123", "456"},
					ThirdPartyServiceIDs:  []string{"123", "456"},
					OnCallScheduleID:      "1244",
				}

				assert.Equal(t, expected, heartbeatMonitor)
			},
		},
		{
			Name:         "list heartbeat monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_heartbeat_monitors.json"),
			Fn: func(t *testing.T, c rest.Client) {
				heartbeatMonitors, err := NewHeartbeatMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.HeartbeatMonitor{
					{
						MonitorID:             "897654345678",
						DisplayName:           "foo",
						NameInPingURL:         "status_check",
						Type:                  "HEARTBEAT",
						ThresholdProfileID:    "012",
						NotificationProfileID: "789",
						MonitorGroups:         []string{"234", "567"},
						UserGroupIDs:          []string{"123", "456"},
						ThirdPartyServiceIDs:  []string{"123", "456"},
						OnCallScheduleID:      "1244",
					},
					{
						MonitorID:             "933654345678",
						DisplayName:           "foo",
						NameInPingURL:         "status_check",
						Type:                  "HEARTBEAT",
						ThresholdProfileID:    "012",
						NotificationProfileID: "789",
						MonitorGroups:         []string{"234", "567"},
						UserGroupIDs:          []string{"123", "456"},
						ThirdPartyServiceIDs:  []string{"123", "456"},
						OnCallScheduleID:      "1244",
					},
				}

				assert.Equal(t, expected, heartbeatMonitors)
			},
		},
		{
			Name:         "update heartbeat monitor",
			ExpectedVerb: "PUT",
			ExpectedPath: "/monitors/123",
			ExpectedBody: validation.Fixture(t, "requests/update_heartbeat_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				heartbeatMonitor := &api.HeartbeatMonitor{
					MonitorID:             "123",
					DisplayName:           "foo",
					NameInPingURL:         "status_check",
					Type:                  "HEARTBEAT",
					ThresholdProfileID:    "012",
					NotificationProfileID: "789",
					MonitorGroups:         []string{"234", "567"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
					ThirdPartyServiceIDs:  []string{"123", "456"},
					OnCallScheduleID:      "1244",
				}

				_, err := NewHeartbeatMonitors(c).Update(heartbeatMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete heartbeat monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewHeartbeatMonitors(c).Delete("123"))
			},
		},
	})
}
