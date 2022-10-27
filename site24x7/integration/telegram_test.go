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

func TestTelegramIntegrationCreate(t *testing.T) {
	d := telegramIntegrationTestResourceData(t)

	c := fake.NewClient()

	a := &api.TelegramIntegration{
		Name:          "foo",
		URL:           "www.test.tld",
		BotToken:      "uojvsdoijsodijdsioj",
		SelectionType: 0,
		TroubleAlert:  true,
		CriticalAlert: false,
		DownAlert:     false,
		Title:         "test-title",
		Monitors:      []string{"234", "567"},
		AlertTagIDs:   []string{"123", "456"},
	}

	c.FakeTelegramIntegration.On("Create", a).Return(a, nil).Once()

	require.NoError(t, telegramIntegrationCreate(d, c))

	c.FakeTelegramIntegration.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := telegramIntegrationCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestTelegramIntegrationUpdate(t *testing.T) {
	d := telegramIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.TelegramIntegration{
		ServiceID:     "123",
		ServiceStatus: 0,
		Name:          "foo",
		URL:           "www.test.tld",
		BotToken:      "uojvsdoijsodijdsioj",
		SelectionType: 0,
		TroubleAlert:  true,
		CriticalAlert: false,
		DownAlert:     false,
		Title:         "test-title",
		Monitors:      []string{"234", "567"},
		AlertTagIDs:   []string{"123", "456"},
	}

	c.FakeTelegramIntegration.On("Update", a).Return(a, nil).Once()

	require.NoError(t, telegramIntegrationUpdate(d, c))

	c.FakeTelegramIntegration.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := telegramIntegrationUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestTelegramIntegrationRead(t *testing.T) {
	d := telegramIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeTelegramIntegration.On("Get", "123").Return(&api.TelegramIntegration{}, nil).Once()

	require.NoError(t, telegramIntegrationRead(d, c))

	c.FakeTelegramIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := telegramIntegrationRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestTelegramIntegrationDelete(t *testing.T) {
	d := telegramIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeThirdPartyIntegrations.On("Delete", "123").Return(nil).Once()

	require.NoError(t, telegramIntegrationDelete(d, c))

	c.FakeThirdPartyIntegrations.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, telegramIntegrationDelete(d, c))
}

func TestTelegramIntegrationExists(t *testing.T) {
	d := telegramIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeTelegramIntegration.On("Get", "123").Return(&api.TelegramIntegration{}, nil).Once()

	exists, err := telegramIntegrationExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeTelegramIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = telegramIntegrationExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeTelegramIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = telegramIntegrationExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func telegramIntegrationTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, TelegramIntegrationSchema, map[string]interface{}{
		"name":           "foo",
		"channel_url":            "www.test.tld",
		"token":		  "uojvsdoijsodijdsioj",
		"selection_type": 0,
		"title":          "test-title",
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
