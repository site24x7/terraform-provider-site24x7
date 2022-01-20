package integration

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebhookIntegration(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create webhook integration",
			ExpectedVerb: "POST",
			ExpectedPath: "/integration/webhooks",
			ExpectedBody: validation.Fixture(t, "requests/create_webhook_integration.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				webhookIntegration := &api.WebhookIntegration{
					Name:                         "webhook_cloud",
					URL:                          "http://example.com",
					Method:                       "P",
					SendIncidentParameters:       false,
					Timeout:                      30,
					IsPollerWebhook:              false,
					AuthMethod:                   "B",
					ManageTickets:                false,
					Username:                     "username",
					Password:                     "password",
					SendCustomParameters:         true,
					SendInJsonFormat:             true,
					CloseSendCustomParameters:    false,
					CloseSendIncidentParameters:  false,
					UpdateSendIncidentParameters: false,
					UpdateSendCustomParameters:   false,
					CustomParameters:             "{\"test\":\"abcd\"}",
					SelectionType:                0,
					AlertTagIDs:                  []string{"113770000023231001"},
				}

				_, err := NewWebhook(c).Create(webhookIntegration)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get webhook integration",
			ExpectedVerb: "GET",
			ExpectedPath: "/integration/webhooks/113770000023231022",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_webhook_integration.json"),
			Fn: func(t *testing.T, c rest.Client) {
				webhook_integration, err := NewWebhook(c).Get("113770000023231022")
				require.NoError(t, err)

				expected := &api.WebhookIntegration{
					Name:                         "Test Webhook",
					ServiceID:                    "113770000023231022",
					ServiceStatus:                0,
					URL:                          "http://requestb.in",
					SelectionType:                0,
					IsPollerWebhook:              false,
					Timeout:                      30,
					Method:                       "P",
					AuthMethod:                   "B",
					Username:                     "username",
					Password:                     "password",
					SendCustomParameters:         true,
					CustomParameters:             "{\"test\":\"abcd\"}",
					SendIncidentParameters:       true,
					SendInJsonFormat:             true,
					CloseSendCustomParameters:    false,
					CloseSendIncidentParameters:  false,
					UpdateSendIncidentParameters: false,
					UpdateSendCustomParameters:   false,
					UserAgent:                    "Mozilla",
					AlertTagIDs:                  []string{"113770000023231001"},
				}

				assert.Equal(t, expected, webhook_integration)
			},
		},
		{
			Name:         "update webhook integration",
			ExpectedVerb: "PUT",
			ExpectedPath: "/integration/webhooks/113770000023231022",
			ExpectedBody: validation.Fixture(t, "requests/update_webhook_integration.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				webhook_integration := &api.WebhookIntegration{
					Name:                         "Test Webhook",
					URL:                          "http://requestb.in",
					IsPollerWebhook:              false,
					Timeout:                      30,
					Method:                       "P",
					AuthMethod:                   "B",
					Username:                     "username",
					Password:                     "password",
					SendCustomParameters:         true,
					ManageTickets:                false,
					CustomParameters:             "{\"test\":\"abcd\"}",
					SendIncidentParameters:       true,
					SendInJsonFormat:             true,
					CloseSendCustomParameters:    false,
					CloseSendIncidentParameters:  false,
					UpdateSendIncidentParameters: false,
					UpdateSendCustomParameters:   false,
					UserAgent:                    "Mozilla",
					ServiceID:                    "113770000023231022",
					SelectionType:                2,
					Monitors:                     []string{"113770000023231032", "113770000023231043"},
					AlertTagIDs:                  []string{"113770000023231001", "113770000023231002"},
				}

				_, err := NewWebhook(c).Update(webhook_integration)
				require.NoError(t, err)
			},
		},
	})
}
