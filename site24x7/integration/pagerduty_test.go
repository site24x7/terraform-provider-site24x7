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

func TestPagerDutyIntegrationCreate(t *testing.T) {
	d := pagerDutyIntegrationTestResourceData(t)

	c := fake.NewClient()

	a := &api.PagerDutyIntegration{
		Name:          "foo",
		ServiceKey:    "service_key",
		SelectionType: 0,
		SenderName:    "Site24x7",
		Title:         "test-title",
		TroubleAlert:  true,
		CriticalAlert: true,
		DownAlert:     true,
		AlertTagIDs:   []string{"123", "456"},
	}

	c.FakePagerDutyIntegration.On("Create", a).Return(a, nil).Once()

	require.NoError(t, pagerDutyIntegrationCreate(d, c))

	c.FakePagerDutyIntegration.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := pagerDutyIntegrationCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestPagerDutyIntegrationUpdate(t *testing.T) {
	d := pagerDutyIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.PagerDutyIntegration{
		ServiceID:     "123",
		Name:          "foo",
		ServiceKey:    "service_key",
		SelectionType: 0,
		SenderName:    "Site24x7",
		Title:         "test-title",
		TroubleAlert:  true,
		CriticalAlert: true,
		DownAlert:     true,
		AlertTagIDs:   []string{"123", "456"},
	}

	c.FakePagerDutyIntegration.On("Update", a).Return(a, nil).Once()

	require.NoError(t, pagerDutyIntegrationUpdate(d, c))

	c.FakePagerDutyIntegration.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := pagerDutyIntegrationUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestPagerDutyIntegrationRead(t *testing.T) {
	d := pagerDutyIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakePagerDutyIntegration.On("Get", "123").Return(&api.PagerDutyIntegration{}, nil).Once()

	require.NoError(t, pagerDutyIntegrationRead(d, c))

	c.FakePagerDutyIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := pagerDutyIntegrationRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestPagerDutyIntegrationDelete(t *testing.T) {
	d := pagerDutyIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeThirdPartyIntegrations.On("Delete", "123").Return(nil).Once()

	require.NoError(t, pagerDutyIntegrationDelete(d, c))

	c.FakeThirdPartyIntegrations.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, pagerDutyIntegrationDelete(d, c))
}

func TestPagerDutyIntegrationExists(t *testing.T) {
	d := pagerDutyIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakePagerDutyIntegration.On("Get", "123").Return(&api.PagerDutyIntegration{}, nil).Once()

	exists, err := pagerDutyIntegrationExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakePagerDutyIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = pagerDutyIntegrationExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakePagerDutyIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = pagerDutyIntegrationExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func pagerDutyIntegrationTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, pagerDutyIntegrationSchema, map[string]interface{}{
		"name":           "foo",
		"service_key":    "service_key",
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
