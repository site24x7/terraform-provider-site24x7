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

func TestNotificationProfileCreate(t *testing.T) {
	d := notificationProfileTestResourceData(t)

	c := fake.NewClient()

	a := &api.NotificationProfile{
		RcaNeeded:                   true,
		NotifyAfterExecutingActions: true,
		ProfileName:                 "Notifi Profile",
		EscalationWaitTime:          60,
		TemplateID:                  "0",
		DowntimeNotificationDelay:   1,
	}

	c.FakeNotificationProfiles.On("Create", a).Return(a, nil).Once()

	require.NoError(t, notificationProfileCreate(d, c))

	c.FakeNotificationProfiles.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := notificationProfileCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestNotificationProfileUpdate(t *testing.T) {
	d := notificationProfileTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.NotificationProfile{
		ProfileID:                   "123",
		ProfileName:                 "Notifi Profile",
		RcaNeeded:                   true,
		EscalationWaitTime:          60,
		NotifyAfterExecutingActions: true,
		TemplateID:                  "0",
		DowntimeNotificationDelay:   1,
	}

	c.FakeNotificationProfiles.On("Update", a).Return(a, nil).Once()

	require.NoError(t, notificationProfileUpdate(d, c))

	c.FakeNotificationProfiles.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := notificationProfileUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestNotificationProfileRead(t *testing.T) {
	d := notificationProfileTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeNotificationProfiles.On("Get", "123").Return(&api.NotificationProfile{}, nil).Once()

	require.NoError(t, notificationProfileRead(d, c))

	c.FakeNotificationProfiles.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := notificationProfileRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestNotificationProfileDelete(t *testing.T) {
	d := notificationProfileTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeNotificationProfiles.On("Delete", "123").Return(nil).Once()

	require.NoError(t, notificationProfileDelete(d, c))

	c.FakeNotificationProfiles.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, notificationProfileDelete(d, c))
}

func TestNotificationProfileExists(t *testing.T) {
	d := notificationProfileTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeNotificationProfiles.On("Get", "123").Return(&api.NotificationProfile{}, nil).Once()

	exists, err := notificationProfileExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeNotificationProfiles.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = notificationProfileExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeNotificationProfiles.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = notificationProfileExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func notificationProfileTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, NotificationProfileSchema, map[string]interface{}{
		"rca_needed":                     true,
		"notify_after_executing_actions": true,
		"profile_name":                   "Notifi Profile",
		"escalation_wait_time":           60,
		"suppress_automation":            false,
	})
}
