package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSlackIntegration(t *testing.T) {
	runTests(t, []*endpointTest{
		{
			name:         "create slack integration",
			expectedVerb: "POST",
			expectedPath: "/integration/slack",
			expectedBody: fixture(t, "requests/create_slack_integration.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
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
			name:         "get slack integration",
			expectedVerb: "GET",
			expectedPath: "/integration/slack/113770000023231022",
			statusCode:   200,
			responseBody: fixture(t, "responses/get_slack_integration.json"),
			fn: func(t *testing.T, c rest.Client) {
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
			name:         "update slack integration",
			expectedVerb: "PUT",
			expectedPath: "/integration/slack/123",
			expectedBody: fixture(t, "requests/update_slack_integration.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
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
