package integration

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPagerDuty(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create pagerduty integration",
			ExpectedVerb: "POST",
			ExpectedPath: "/integration/pager_duty",
			ExpectedBody: validation.Fixture(t, "requests/create_pagerduty_integration.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				pagerDutyIntegration := &api.PagerDutyIntegration{
					Name:          "Site24x7-PagerDuty Integration",
					ServiceKey:    "service_key",
					TroubleAlert:  true,
					CriticalAlert: false,
					DownAlert:     false,
					ManualResolve: false,
					SelectionType: 0,
					SenderName:    "Site24x7",
					Title:         "$MONITOR_NAME is $STATUS",
					AlertTagIDs:   []string{"113770000023231001"},
				}

				_, err := NewPagerDuty(c).Create(pagerDutyIntegration)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get pagerduty integration",
			ExpectedVerb: "GET",
			ExpectedPath: "/integration/pager_duty/113770000023231022",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_pagerduty_integration.json"),
			Fn: func(t *testing.T, c rest.Client) {
				pagerduty_integration, err := NewPagerDuty(c).Get("113770000023231022")
				require.NoError(t, err)

				expected := &api.PagerDutyIntegration{
					Name:          "Site24x7-PagerDuty Integration",
					ServiceID:     "113770000023231022",
					ServiceStatus: 0,
					SelectionType: 0,
					SenderName:    "Site24x7",
					Title:         "$MONITOR_NAME is $STATUS",
					AlertTagIDs:   []string{"113770000023231001"},
				}

				assert.Equal(t, expected, pagerduty_integration)
			},
		},
		{
			Name:         "update pagerduty integration",
			ExpectedVerb: "PUT",
			ExpectedPath: "/integration/pager_duty/123",
			ExpectedBody: validation.Fixture(t, "requests/update_pagerduty_integration.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				pagerduty_integration := &api.PagerDutyIntegration{
					Name:          "Site24x7-PagerDuty Integration",
					ServiceID:     "123",
					ServiceKey:    "service_key",
					TroubleAlert:  true,
					CriticalAlert: false,
					DownAlert:     false,
					ManualResolve: false,
					SelectionType: 2,
					SenderName:    "Site24x7",
					Monitors:      []string{"113770000023231032", "113770000023231043"},
					Title:         "$MONITOR_NAME is $STATUS",
					AlertTagIDs:   []string{"113770000023231001", "113770000023231002"},
				}

				_, err := NewPagerDuty(c).Update(pagerduty_integration)
				require.NoError(t, err)
			},
		},
	})
}
