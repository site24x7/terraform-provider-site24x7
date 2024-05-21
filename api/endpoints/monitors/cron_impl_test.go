package monitors

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCronMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create cron monitor",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitors",
			ExpectedBody: validation.Fixture(t, "requests/create_cron_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				cronMonitor := &api.CronMonitor{
					DisplayName:           "foo",
					CronExpression:        "* * * * *",
					CronTz:                "IST",
					WaitTime:              30,
					Type:                  "CRON",
					ThresholdProfileID:    "012",
					NotificationProfileID: "789",
					MonitorGroups:         []string{"234", "567"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
					ThirdPartyServiceIDs:  []string{"123", "456"},
					OnCallScheduleID:      "1244",
				}

				_, err := NewCronMonitors(c).Create(cronMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get cron monitor",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors/897654345678",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_cron_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				cronMonitor, err := NewCronMonitors(c).Get("897654345678")
				require.NoError(t, err)

				expected := &api.CronMonitor{
					MonitorID:             "897654345678",
					DisplayName:           "foo",
					CronExpression:        "* * * * *",
					CronTz:                "IST",
					WaitTime:              30,
					Type:                  "CRON",
					ThresholdProfileID:    "012",
					NotificationProfileID: "789",
					MonitorGroups:         []string{"234", "567"},
					UserGroupIDs:          []string{"123", "456"},
					ThirdPartyServiceIDs:  []string{"123", "456"},
					OnCallScheduleID:      "1244",
				}

				assert.Equal(t, expected, cronMonitor)
			},
		},
		{
			Name:         "list cron monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_cron_monitors.json"),
			Fn: func(t *testing.T, c rest.Client) {
				cronMonitors, err := NewCronMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.CronMonitor{
					{
						MonitorID:             "897654345678",
						DisplayName:           "foo",
						CronExpression:        "* * * * *",
						CronTz:                "IST",
						WaitTime:              30,
						Type:                  "CRON",
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
						CronExpression:        "* * * * *",
						CronTz:                "IST",
						WaitTime:              30,
						Type:                  "CRON",
						ThresholdProfileID:    "012",
						NotificationProfileID: "789",
						MonitorGroups:         []string{"234", "567"},
						UserGroupIDs:          []string{"123", "456"},
						ThirdPartyServiceIDs:  []string{"123", "456"},
						OnCallScheduleID:      "1244",
					},
				}

				assert.Equal(t, expected, cronMonitors)
			},
		},
		{
			Name:         "update cron monitor",
			ExpectedVerb: "PUT",
			ExpectedPath: "/monitors/123",
			ExpectedBody: validation.Fixture(t, "requests/update_cron_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				cronMonitor := &api.CronMonitor{
					MonitorID:             "123",
					DisplayName:           "foo",
					CronExpression:        "* * * * *",
					CronTz:                "IST",
					WaitTime:              30,
					Type:                  "CRON",
					ThresholdProfileID:    "012",
					NotificationProfileID: "789",
					MonitorGroups:         []string{"234", "567"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
					ThirdPartyServiceIDs:  []string{"123", "456"},
					OnCallScheduleID:      "1244",
				}

				_, err := NewCronMonitors(c).Update(cronMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete cron monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewCronMonitors(c).Delete("123"))
			},
		},
	})
}
