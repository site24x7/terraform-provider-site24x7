package monitors

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPortMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create port monitor",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitors",
			ExpectedBody: validation.Fixture(t, "requests/create_port_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				portMonitor := &api.PortMonitor{
					DisplayName:           "Port Monitor",
					HostName:              "www.example.com",
					Type:                  "PORT",
					ApplicationType:       "FTP",
					Port:                  80,
					UseIPV6:               true,
					Timeout:               10,
					Command:               "new command",
					UseSSL:                true,
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

				_, err := NewPortMonitors(c).Create(portMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get port monitor",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors/897654345678",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_port_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				portMonitor, err := NewPortMonitors(c).Get("897654345678")
				require.NoError(t, err)

				expected := &api.PortMonitor{
					MonitorID:             "897654345678",
					DisplayName:           "Port Monitor",
					HostName:              "www.example.com",
					Type:                  "PORT",
					ApplicationType:       "FTP",
					Port:                  80,
					UseIPV6:               true,
					Timeout:               10,
					Command:               "new command",
					UseSSL:                true,
					CheckFrequency:        "5",
					InvertPortCheck:       false,
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

				assert.Equal(t, expected, portMonitor)
			},
		},
		{
			Name:         "list port monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_port_monitors.json"),
			Fn: func(t *testing.T, c rest.Client) {
				portMonitor, err := NewPortMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.PortMonitor{
					{
						MonitorID:             "897654345678",
						DisplayName:           "Port Monitor",
						HostName:              "www.example.com",
						ApplicationType:       "FTP",
						Type:                  "PORT",
						Port:                  80,
						UseIPV6:               true,
						Timeout:               10,
						Command:               "new command",
						UseSSL:                true,
						CheckFrequency:        "5",
						InvertPortCheck:       false,
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
						DisplayName:           "Port Monitor",
						HostName:              "www.example.com",
						ApplicationType:       "FTP",
						Type:                  "PORT",
						Port:                  80,
						UseIPV6:               true,
						Timeout:               10,
						Command:               "new command",
						UseSSL:                true,
						CheckFrequency:        "5",
						InvertPortCheck:       false,
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

				assert.Equal(t, expected, portMonitor)
			},
		},
		{
			Name:         "update port monitor",
			ExpectedVerb: "PUT",
			ExpectedPath: "/monitors/897654345678",
			ExpectedBody: validation.Fixture(t, "requests/update_port_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				portMonitor := &api.PortMonitor{
					MonitorID:             "897654345678",
					DisplayName:           "Port Monitor",
					HostName:              "www.example.com",
					ApplicationType:       "FTP",
					Type:                  "PORT",
					Port:                  80,
					UseIPV6:               true,
					Timeout:               10,
					Command:               "new command",
					UseSSL:                true,
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

				_, err := NewPortMonitors(c).Update(portMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete port monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/897654345678",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewPortMonitors(c).Delete("897654345678"))
			},
		},
	})
}
