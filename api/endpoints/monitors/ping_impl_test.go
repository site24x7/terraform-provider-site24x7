package monitors

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPINGMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create ping monitor",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitors",
			ExpectedBody: validation.Fixture(t, "requests/create_ping_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				pingMonitor := &api.PINGMonitor{
					DisplayName:           "PING Monitor",
					HostName:              "www.example.com",
					Type:                  "PING",
					UseIPV6:               true,
					Timeout:               10,
					CheckFrequency:        "5",
					OnCallScheduleID:      "23524543545245",
					LocationProfileID:     "123412341234123412",
					NotificationProfileID: "123412341234123412",
					ThresholdProfileID:    "123412341234123414",
					MonitorGroups:         []string{"234", "567"},
					DependencyResourceIDs: []string{"123", "456"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
					PerformAutomation:     true,
				}

				_, err := NewPINGMonitors(c).Create(pingMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get ping monitor",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors/897654345678",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_ping_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				pingMonitor, err := NewPINGMonitors(c).Get("897654345678")
				require.NoError(t, err)

				expected := &api.PINGMonitor{
					MonitorID:             "897654345678",
					DisplayName:           "PING Monitor",
					HostName:              "www.example.com",
					Type:                  "PING",
					UseIPV6:               true,
					Timeout:               10,
					CheckFrequency:        "5",
					OnCallScheduleID:      "23524543545245",
					LocationProfileID:     "123412341234123412",
					NotificationProfileID: "123412341234123412",
					ThresholdProfileID:    "123412341234123414",
					MonitorGroups:         []string{"234", "567"},
					DependencyResourceIDs: []string{"123", "456"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
					PerformAutomation:     false,
				}

				assert.Equal(t, expected, pingMonitor)
			},
		},
		{
			Name:         "list ping monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_ping_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				pingMonitor, err := NewPINGMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.PINGMonitor{
					{
						MonitorID:             "897654345678",
						DisplayName:           "PING Monitor",
						HostName:              "www.example.com",
						Type:                  "PING",
						UseIPV6:               true,
						Timeout:               10,
						CheckFrequency:        "5",
						OnCallScheduleID:      "23524543545245",
						LocationProfileID:     "123412341234123412",
						NotificationProfileID: "123412341234123412",
						ThresholdProfileID:    "123412341234123414",
						MonitorGroups:         []string{"234", "567"},
						DependencyResourceIDs: []string{"123", "456"},
						UserGroupIDs:          []string{"123", "456"},
						TagIDs:                []string{"123"},
						PerformAutomation:     true,
					},
					{
						MonitorID:             "654568778999889",
						DisplayName:           "PING Monitor",
						HostName:              "www.example.com",
						Type:                  "PING",
						UseIPV6:               true,
						Timeout:               10,
						CheckFrequency:        "5",
						OnCallScheduleID:      "23524543545245",
						LocationProfileID:     "123412341234123412",
						NotificationProfileID: "123412341234123412",
						ThresholdProfileID:    "123412341234123414",
						MonitorGroups:         []string{"234", "567"},
						DependencyResourceIDs: []string{"123", "456"},
						UserGroupIDs:          []string{"123", "456"},
						TagIDs:                []string{"123"},
						PerformAutomation:     true,
					},
				}

				assert.Equal(t, expected, pingMonitor)
			},
		},
		{
			Name:         "update ping monitor",
			ExpectedVerb: "PUT",
			ExpectedPath: "/monitors/897654345678",
			ExpectedBody: validation.Fixture(t, "requests/update_ping_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				pingMonitor := &api.PINGMonitor{
					MonitorID:             "897654345678",
					DisplayName:           "PING Monitor",
					HostName:              "www.example.com",
					Type:                  "PING",
					UseIPV6:               true,
					Timeout:               10,
					CheckFrequency:        "5",
					OnCallScheduleID:      "23524543545245",
					LocationProfileID:     "123412341234123412",
					NotificationProfileID: "123412341234123412",
					ThresholdProfileID:    "123412341234123414",
					MonitorGroups:         []string{"234", "567"},
					DependencyResourceIDs: []string{"123", "456"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
					PerformAutomation:     true,
				}

				_, err := NewPINGMonitors(c).Update(pingMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete ping monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/897654345678",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewPINGMonitors(c).Delete("897654345678"))
			},
		},
	})
}
