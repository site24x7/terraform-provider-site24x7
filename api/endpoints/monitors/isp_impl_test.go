package monitors

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestISPMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create isp monitor",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitors",
			ExpectedBody: validation.Fixture(t, "requests/create_isp_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				ispMonitor := &api.ISPMonitor{
					DisplayName:           "ISP Monitor",
					Hostname:              "www.example.com",
					UseIPV6:               true,
					Type:                  "ISP",
					Timeout:               30,
					Protocol:              "1",
					Port:                  443,
					CheckFrequency:        "5",
					LocationProfileID:     "123412341234123412",
					NotificationProfileID: "123412341234123412",
					ThresholdProfileID:    "123412341234123414",
					MonitorGroups:         []string{"234", "567"},
					DependencyResourceIDs: []string{"123", "456"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
				}

				_, err := NewISPMonitors(c).Create(ispMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get isp monitor",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors/897654345678",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_isp_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				ispMonitor, err := NewISPMonitors(c).Get("897654345678")
				require.NoError(t, err)

				expected := &api.ISPMonitor{
					MonitorID:             "897654345678",
					DisplayName:           "ISP Monitor",
					Hostname:              "www.example.com",
					UseIPV6:               true,
					Type:                  "ISP",
					Timeout:               30,
					Protocol:              "1",
					Port:                  443,
					CheckFrequency:        "5",
					LocationProfileID:     "123412341234123412",
					NotificationProfileID: "123412341234123412",
					ThresholdProfileID:    "123412341234123414",
					MonitorGroups:         []string{"234", "567"},
					DependencyResourceIDs: []string{"123", "456"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
				}

				assert.Equal(t, expected, ispMonitor)
			},
		},
		{
			Name:         "list isp monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_isp_monitors.json"),
			Fn: func(t *testing.T, c rest.Client) {
				ispMonitors, err := NewISPMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.ISPMonitor{
					{
						MonitorID:             "897654345678",
						DisplayName:           "ISP Monitor",
						Hostname:              "www.example.com",
						UseIPV6:               true,
						Type:                  "ISP",
						Timeout:               30,
						Protocol:              "1",
						Port:                  443,
						CheckFrequency:        "5",
						LocationProfileID:     "123412341234123412",
						NotificationProfileID: "123412341234123412",
						ThresholdProfileID:    "123412341234123414",
						MonitorGroups:         []string{"234", "567"},
						DependencyResourceIDs: []string{"123", "456"},
						UserGroupIDs:          []string{"123", "456"},
						TagIDs:                []string{"123"},
					},
					{
						MonitorID:             "933654345678",
						DisplayName:           "ISP Monitor",
						Hostname:              "www.example.com",
						UseIPV6:               true,
						Type:                  "ISP",
						Timeout:               30,
						Protocol:              "1",
						Port:                  443,
						CheckFrequency:        "5",
						LocationProfileID:     "123412341234123412",
						NotificationProfileID: "123412341234123412",
						ThresholdProfileID:    "123412341234123414",
						MonitorGroups:         []string{"234", "567"},
						DependencyResourceIDs: []string{"123", "456"},
						UserGroupIDs:          []string{"123", "456"},
						TagIDs:                []string{"123"},
					},
				}

				assert.Equal(t, expected, ispMonitors)
			},
		},
		{
			Name:         "update isp monitor",
			ExpectedVerb: "PUT",
			ExpectedPath: "/monitors/897654345678",
			ExpectedBody: validation.Fixture(t, "requests/update_isp_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				ispMonitor := &api.ISPMonitor{
					MonitorID:             "897654345678",
					DisplayName:           "ISP Monitor",
					Hostname:              "www.example.com",
					UseIPV6:               true,
					Type:                  "ISP",
					Timeout:               30,
					Protocol:              "1",
					Port:                  443,
					CheckFrequency:        "5",
					LocationProfileID:     "123412341234123412",
					NotificationProfileID: "123412341234123412",
					ThresholdProfileID:    "123412341234123414",
					MonitorGroups:         []string{"234", "567"},
					DependencyResourceIDs: []string{"123", "456"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
				}

				_, err := NewISPMonitors(c).Update(ispMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete isp monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/897654345678",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewISPMonitors(c).Delete("897654345678"))
			},
		},
	})
}
