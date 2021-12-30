package site24x7

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSlackIntegrationCreate(t *testing.T) {
	d := slackIntegrationTestResourceData(t)

	c := fake.NewClient()

	a := &api.SlackIntegration{
		Name:          "foo",
		URL:           "www.test.tld",
		SelectionType: 0,
		Title:         "test-title",
		SenderName:    "test-sender",
		Monitors:      []string{"234", "567"},
		AlertTagIDs:   []string{"123", "456"},
	}

	c.FakeSlackIntegration.On("Create", a).Return(a, nil).Once()

	require.NoError(t, slackIntegrationCreate(d, c))

	c.FakeSlackIntegration.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := slackIntegrationCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestSlackIntegrationUpdate(t *testing.T) {
	d := slackIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.SlackIntegration{
		ServiceID:     "123",
		ServiceStatus: 0,
		Name:          "foo",
		URL:           "www.test.tld",
		SelectionType: 0,
		Title:         "test-title",
		SenderName:    "test-sender",
		Monitors:      []string{"234", "567"},
		AlertTagIDs:   []string{"123", "456"},
	}

	c.FakeSlackIntegration.On("Update", a).Return(a, nil).Once()

	require.NoError(t, slackIntegrationUpdate(d, c))

	c.FakeSlackIntegration.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := slackIntegrationUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestSlackIntegrationRead(t *testing.T) {
	d := slackIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeSlackIntegration.On("Get", "123").Return(&api.SlackIntegration{}, nil).Once()

	require.NoError(t, slackIntegrationRead(d, c))

	c.FakeSlackIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := slackIntegrationRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestSlackIntegrationDelete(t *testing.T) {
	d := slackIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeThirdPartyIntegrations.On("Delete", "123").Return(nil).Once()

	require.NoError(t, slackIntegrationDelete(d, c))

	c.FakeThirdPartyIntegrations.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, slackIntegrationDelete(d, c))
}

func TestSlackIntegrationExists(t *testing.T) {
	d := slackIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeSlackIntegration.On("Get", "123").Return(&api.SlackIntegration{}, nil).Once()

	exists, err := slackIntegrationExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeSlackIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = slackIntegrationExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeSlackIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = slackIntegrationExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func slackIntegrationTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, SlackIntegrationSchema, map[string]interface{}{
		"name":           "foo",
		"url":            "www.test.tld",
		"selection_type": 0,
		"title":          "test-title",
		"sender_name":    "test-sender",
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
