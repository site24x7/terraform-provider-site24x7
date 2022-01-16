package integration

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSlackIntegration(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create slack integration",
			ExpectedVerb: "POST",
			ExpectedPath: "/integration/slack",
			ExpectedBody: validation.Fixture(t, "requests/create_slack_integration.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				slackIntegration := &api.SlackIntegration{
					Name:          "Site24x7-Slack Integration",
					URL:           "https://hooks.slack.com/services/B27AG46BW/W27JLYuDE/acc3vmmJIGrNuBG9CVRwiBxU",
					SelectionType: 0,
					SenderName:    "Site24x7",
					Title:         "$MONITOR_NAME is $STATUS",
					AlertTagIDs:   []string{"113770000023231001"},
				}

				_, err := NewSlack(c).Create(slackIntegration)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get slack integration",
			ExpectedVerb: "GET",
			ExpectedPath: "/integration/slack/113770000023231022",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_slack_integration.json"),
			Fn: func(t *testing.T, c rest.Client) {
				slack_integration, err := NewSlack(c).Get("113770000023231022")
				require.NoError(t, err)

				expected := &api.SlackIntegration{
					Name:          "Site24x7-Slack Integration",
					URL:           "https://hooks.slack.com/services/B27AG46BW/W27JLYuDE/acc3vmmJIGrNuBG9CVRwiBxU",
					ServiceID:     "113770000023231022",
					ServiceStatus: 0,
					SelectionType: 0,
					SenderName:    "Site24x7",
					Title:         "$MONITOR_NAME is $STATUS",
					AlertTagIDs:   []string{"113770000023231001"},
				}

				assert.Equal(t, expected, slack_integration)
			},
		},
		{
			Name:         "update slack integration",
			ExpectedVerb: "PUT",
			ExpectedPath: "/integration/slack/123",
			ExpectedBody: validation.Fixture(t, "requests/update_slack_integration.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				slack_integration := &api.SlackIntegration{
					Name:          "Site24x7-Slack Integration",
					URL:           "https://hooks.slack.com/services/B27AG46BW/W27JLYuDE/acc3vmmJIGrNuBG9CVRwiBxU",
					ServiceID:     "123",
					SelectionType: 2,
					SenderName:    "Site24x7",
					Monitors:      []string{"113770000023231032", "113770000023231043"},
					Title:         "$MONITOR_NAME is $STATUS",
					AlertTagIDs:   []string{"113770000023231001", "113770000023231002"},
				}

				_, err := NewSlack(c).Update(slack_integration)
				require.NoError(t, err)
			},
		},
	})
}
