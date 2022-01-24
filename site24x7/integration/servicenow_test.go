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

func TestServiceNowIntegrationCreate(t *testing.T) {
	d := serviceNowIntegrationTestResourceData(t)
	c := fake.NewClient()

	a := &api.ServiceNowIntegration{
		Name:          "foo",
		InstanceURL:   "https://www.example.com",
		SelectionType: 0,
		SenderName:    "Site24x7",
		Title:         "test-title",
		UserName:      "username",
		Password:      "",
		TroubleAlert:  true,
		CriticalAlert: true,
		DownAlert:     true,
		AlertTagIDs:   []string{"123", "456"},
	}

	c.FakeServiceNowIntegration.On("Create", a).Return(a, nil).Once()

	require.NoError(t, serviceNowIntegrationCreate(d, c))

	c.FakeServiceNowIntegration.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := serviceNowIntegrationCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestServiceNowIntegrationUpdate(t *testing.T) {
	d := serviceNowIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.ServiceNowIntegration{
		ServiceID:     "123",
		Name:          "foo",
		InstanceURL:   "https://www.example.com",
		SelectionType: 0,
		UserName:      "username",
		Password:      "",
		SenderName:    "Site24x7",
		Title:         "test-title",
		TroubleAlert:  true,
		CriticalAlert: true,
		DownAlert:     true,
		AlertTagIDs:   []string{"123", "456"},
	}

	c.FakeServiceNowIntegration.On("Update", a).Return(a, nil).Once()

	require.NoError(t, serviceNowIntegrationUpdate(d, c))

	c.FakeServiceNowIntegration.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := serviceNowIntegrationUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestServiceNowIntegrationRead(t *testing.T) {
	d := serviceNowIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeServiceNowIntegration.On("Get", "123").Return(&api.ServiceNowIntegration{}, nil).Once()

	require.NoError(t, serviceNowIntegrationRead(d, c))

	c.FakeServiceNowIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := serviceNowIntegrationRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestServiceNowIntegrationDelete(t *testing.T) {
	d := serviceNowIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeThirdPartyIntegrations.On("Delete", "123").Return(nil).Once()

	require.NoError(t, serviceNowIntegrationDelete(d, c))

	c.FakeThirdPartyIntegrations.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, serviceNowIntegrationDelete(d, c))
}

func TestServiceNowIntegrationExists(t *testing.T) {
	d := serviceNowIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeServiceNowIntegration.On("Get", "123").Return(&api.ServiceNowIntegration{}, nil).Once()

	exists, err := serviceNowIntegrationExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeServiceNowIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = serviceNowIntegrationExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeServiceNowIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = serviceNowIntegrationExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func serviceNowIntegrationTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, serviceNowIntegrationSchema, map[string]interface{}{
		"name":           "foo",
		"instance_url":   "https://www.example.com",
		"user_name":      "username",
		"password":       "",
		"selection_type": 0,
		"title":          "test-title",
		"sender_name":    "Site24x7",
		"trouble_alert":  true,
		"critical_alert": true,
		"down_alert":     true,
		"alert_tags_id": []interface{}{
			"123",
			"456",
		},
	})
}
