package integration

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebhookIntegrationCreate(t *testing.T) {
	d := webhookIntegrationTestResourceData(t)

	c := fake.NewClient()

	a := &api.WebhookIntegration{
		Name:                         "webhook_test",
		URL:                          "http://example.com",
		Timeout:                      30,
		Method:                       "P",
		SelectionType:                0,
		TroubleAlert:                 true,
		CriticalAlert:                false,
		DownAlert:                    false,
		IsPollerWebhook:              false,
		UserAgent:                    "Mozilla",
		ManageTickets:                false,
		SendInJsonFormat:             true,
		SendIncidentParameters:       false,
		SendCustomParameters:         true,
		UpdateSendCustomParameters:   false,
		UpdateSendIncidentParameters: false,
		CloseSendCustomParameters:    false,
		CloseSendIncidentParameters:  false,
		CustomParameters:             "{\"test\":\"abcd\"}",
		Monitors:                     []string{"234", "567"},
		AlertTagIDs:                  []string{"123", "456"},
		CustomHeaders: []api.Header{
			{
				Name:  "Accept",
				Value: "application/json",
			},
			{
				Name:  "Cache-Control",
				Value: "nocache",
			},
		},
	}

	c.FakeWebhookIntegration.On("Create", a).Return(a, nil).Once()

	require.NoError(t, webhookIntegrationCreate(d, c))

	c.FakeWebhookIntegration.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := webhookIntegrationCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestWebhookIntegrationUpdate(t *testing.T) {
	d := webhookIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.WebhookIntegration{
		ServiceID:                    "123",
		Name:                         "webhook_test",
		URL:                          "http://example.com",
		Timeout:                      30,
		Method:                       "P",
		SelectionType:                0,
		TroubleAlert:                 true,
		CriticalAlert:                false,
		DownAlert:                    false,
		IsPollerWebhook:              false,
		UserAgent:                    "Mozilla",
		ManageTickets:                false,
		SendInJsonFormat:             true,
		SendIncidentParameters:       false,
		SendCustomParameters:         true,
		UpdateSendCustomParameters:   false,
		UpdateSendIncidentParameters: false,
		CloseSendCustomParameters:    false,
		CloseSendIncidentParameters:  false,
		CustomParameters:             "{\"test\":\"abcd\"}",
		Monitors:                     []string{"234", "567"},
		AlertTagIDs:                  []string{"123", "456"},
		CustomHeaders: []api.Header{
			{
				Name:  "Accept",
				Value: "application/json",
			},
			{
				Name:  "Cache-Control",
				Value: "nocache",
			},
		},
	}

	c.FakeWebhookIntegration.On("Update", a).Return(a, nil).Once()

	require.NoError(t, webhookIntegrationUpdate(d, c))

	c.FakeWebhookIntegration.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := webhookIntegrationUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestWebhookIntegrationRead(t *testing.T) {
	d := webhookIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeWebhookIntegration.On("Get", "123").Return(&api.WebhookIntegration{}, nil).Once()

	require.NoError(t, webhookIntegrationRead(d, c))

	c.FakeWebhookIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := webhookIntegrationRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestWebhookIntegrationDelete(t *testing.T) {
	d := webhookIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeThirdPartyIntegrations.On("Delete", "123").Return(nil).Once()

	require.NoError(t, webhookIntegrationDelete(d, c))

	c.FakeThirdPartyIntegrations.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, webhookIntegrationDelete(d, c))
}

func TestWebhookIntegrationExists(t *testing.T) {
	d := webhookIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeWebhookIntegration.On("Get", "123").Return(&api.WebhookIntegration{}, nil).Once()

	exists, err := webhookIntegrationExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeWebhookIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = webhookIntegrationExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeWebhookIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = webhookIntegrationExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func webhookIntegrationTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, WebhookIntegrationSchema, map[string]interface{}{
		"name":                            "webhook_test",
		"url":                             "http://example.com",
		"timeout":                         30,
		"method":                          "P",
		"selection_type":                  0,
		"trouble_alert":                   true,
		"critical_alert":                  false,
		"down_alert":                      false,
		"is_poller_webhook":               false,
		"send_incident_parameters":        false,
		"send_in_json_format":             true,
		"update_send_incident_parameters": false,
		"update_send_custom_parameters":   false,
		"close_send_incident_parameters":  false,
		"close_send_custom_parameters":    false,
		"manage_tickets":                  false,
		"send_custom_parameters":          true,
		"custom_parameters":               "{\"test\":\"abcd\"}",
		"user_agent":                      "Mozilla",
		"custom_headers": map[string]interface{}{
			"Accept":        "application/json",
			"Cache-Control": "nocache",
		},
		"monitors": []interface{}{
			"234",
			"567",
		},
		"alert_tags_id": []interface{}{
			"123",
			"456",
		},
	})
}
