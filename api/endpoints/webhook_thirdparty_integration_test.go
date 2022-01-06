package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebhookIntegration(t *testing.T) {
	runTests(t, []*endpointTest{
		{
			name:         "create webhook integration",
			expectedVerb: "POST",
			expectedPath: "/integration/webhooks",
			expectedBody: fixture(t, "requests/create_webhook_integration.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
				webhookIntegration := &api.WebhookIntegration{
					Name:                   "webhook_cloud",
					URL:                    "http://example.com",
					Method:                 "P",
					SendIncidentParameters: false,
					Timeout:                30,
					IsPollerWebhook:        false,
					AuthMethod:             "B",
					ManageTickets:          false,
					Username:               "username",
					Password:               "password",
					SendCustomParameters:   true,
					SendInJsonFormat:       true,
					CustomParameters:       "{\"test\":\"abcd\"}",
					SelectionType:          0,
					AlertTagIDs:            []string{"113770000023231001"},
				}

				_, err := NewWebhook(c).Create(webhookIntegration)
				require.NoError(t, err)
			},
		},
		{
			name:         "get webhook integration",
			expectedVerb: "GET",
			expectedPath: "/integration/webhooks/113770000023231022",
			statusCode:   200,
			responseBody: fixture(t, "responses/get_webhook_integration.json"),
			fn: func(t *testing.T, c rest.Client) {
				webhook_integration, err := NewWebhook(c).Get("113770000023231022")
				require.NoError(t, err)

				expected := &api.WebhookIntegration{
					Name:                   "Test WebHook",
					ServiceID:              "113770000023231022",
					ServiceStatus:          0,
					URL:                    "http://requestb.in",
					SelectionType:          0,
					IsPollerWebhook:        false,
					Timeout:                30,
					Method:                 "P",
					AuthMethod:             "B",
					Username:               "username",
					Password:               "password",
					SendCustomParameters:   true,
					CustomParameters:       "param=value",
					SendIncidentParameters: true,
					SendInJsonFormat:       true,
					UserAgent:              "Mozilla",
					AlertTagIDs:            []string{"113770000023231001"},
				}

				assert.Equal(t, expected, webhook_integration)
			},
		},
		{
			name:         "update webhook integration",
			expectedVerb: "PUT",
			expectedPath: "/integration/webhooks/113770000023231022",
			expectedBody: fixture(t, "requests/update_webhook_integration.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
				webhook_integration := &api.WebhookIntegration{
					Name:                   "Test WebHook",
					URL:                    "http://requestb.in",
					IsPollerWebhook:        false,
					Timeout:                30,
					Method:                 "P",
					AuthMethod:             "B",
					Username:               "username",
					Password:               "password",
					SendCustomParameters:   true,
					ManageTickets:          false,
					CustomParameters:       "param=value",
					SendIncidentParameters: true,
					SendInJsonFormat:       true,
					UserAgent:              "Mozilla",
					ServiceID:              "113770000023231022",
					SelectionType:          2,
					Monitors:               []string{"113770000023231032", "113770000023231043"},
					AlertTagIDs:            []string{"113770000023231001", "113770000023231002"},
				}

				_, err := NewWebhook(c).Update(webhook_integration)
				require.NoError(t, err)
			},
		},
	})
}
