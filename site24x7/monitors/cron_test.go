package monitors

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCronMonitorCreate(t *testing.T) {
	d := cronTestResourceData(t)

	c := fake.NewClient()

	a := &api.CronMonitor{
		DisplayName:           "foo",
		CronExpression:        "* * * * *",
		CronTz:                "IST",
		WaitTime:              30,
		Type:                  "CRON",
		ThresholdProfileID:    "012",
		NotificationProfileID: "789",
		MonitorGroups:         []string{"234", "567"},
		UserGroupIDs:          []string{"456", "123"},
		TagIDs:                []string{"123"},
		ThirdPartyServiceIDs:  []string{"456", "123"},
		OnCallScheduleID:      "1244",
	}

	notificationProfiles := []*api.NotificationProfile{
		{
			ProfileID:   "123",
			ProfileName: "Notifi Profile",
			RcaNeeded:   true,
		},
		{
			ProfileID:   "456",
			ProfileName: "TEST",
			RcaNeeded:   false,
		},
	}
	c.FakeNotificationProfiles.On("List").Return(notificationProfiles, nil)

	userGroups := []*api.UserGroup{
		{
			DisplayName:      "Admin Group",
			Users:            []string{"123", "456"},
			AttributeGroupID: "789",
			ProductID:        0,
		},
		{
			DisplayName:      "Network Group",
			Users:            []string{"123", "456"},
			AttributeGroupID: "345",
			ProductID:        0,
		},
	}
	c.FakeUserGroups.On("List").Return(userGroups, nil)

	tags := []*api.Tag{
		{
			TagID:    "123",
			TagName:  "aws tag",
			TagValue: "baz",
			TagColor: "#B7DA9E",
		},
		{
			TagID:    "456",
			TagName:  "website tag",
			TagValue: "baz 1",
			TagColor: "#B7DA9E",
		},
	}
	c.FakeTags.On("List").Return(tags, nil)

	c.FakeCronMonitors.On("Create", a).Return(a, nil).Once()

	require.NoError(t, cronMonitorCreate(d, c))

	c.FakeCronMonitors.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := cronMonitorCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestCronMonitorUpdate(t *testing.T) {
	d := cronTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.CronMonitor{
		MonitorID:             "123",
		DisplayName:           "foo",
		CronExpression:        "* * * * *",
		CronTz:                "IST",
		WaitTime:              30,
		Type:                  "CRON",
		ThresholdProfileID:    "012",
		NotificationProfileID: "789",
		MonitorGroups:         []string{"234", "567"},
		UserGroupIDs:          []string{"456", "123"},
		TagIDs:                []string{"123"},
		ThirdPartyServiceIDs:  []string{"456", "123"},
		OnCallScheduleID:      "1244",
	}

	notificationProfiles := []*api.NotificationProfile{
		{
			ProfileID:   "123",
			ProfileName: "Notifi Profile",
			RcaNeeded:   true,
		},
		{
			ProfileID:   "456",
			ProfileName: "TEST",
			RcaNeeded:   false,
		},
	}
	c.FakeNotificationProfiles.On("List").Return(notificationProfiles, nil)

	userGroups := []*api.UserGroup{
		{
			DisplayName:      "Admin Group",
			Users:            []string{"123", "456"},
			AttributeGroupID: "789",
			ProductID:        0,
		},
		{
			DisplayName:      "Network Group",
			Users:            []string{"123", "456"},
			AttributeGroupID: "345",
			ProductID:        0,
		},
	}
	c.FakeUserGroups.On("List").Return(userGroups, nil)

	tags := []*api.Tag{
		{
			TagID:    "123",
			TagName:  "aws tag",
			TagValue: "baz",
			TagColor: "#B7DA9E",
		},
		{
			TagID:    "456",
			TagName:  "website tag",
			TagValue: "baz 1",
			TagColor: "#B7DA9E",
		},
	}
	c.FakeTags.On("List").Return(tags, nil)

	c.FakeCronMonitors.On("Update", a).Return(a, nil).Once()

	require.NoError(t, cronMonitorUpdate(d, c))

	c.FakeCronMonitors.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := cronMonitorUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestCronMonitorRead(t *testing.T) {
	d := cronTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeCronMonitors.On("Get", "123").Return(&api.CronMonitor{}, nil).Once()

	require.NoError(t, cronMonitorRead(d, c))

	c.FakeCronMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := cronMonitorRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestCronMonitorDelete(t *testing.T) {
	d := cronTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeCronMonitors.On("Delete", "123").Return(nil).Once()

	require.NoError(t, cronMonitorDelete(d, c))

	c.FakeCronMonitors.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, cronMonitorDelete(d, c))
}

func TestCronMonitorExists(t *testing.T) {
	d := cronTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeCronMonitors.On("Get", "123").Return(&api.CronMonitor{}, nil).Once()

	exists, err := cronMonitorExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeCronMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = cronMonitorExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeCronMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = cronMonitorExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func cronTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, CronMonitorSchema, map[string]interface{}{
		"display_name":            "foo",
		"type":                    "CRON",
		"cron_expression":         "* * * * *",
		"cron_tz":                 "IST",
		"wait_time":               30,
		"threshold_profile_id":    "012",
		"notification_profile_id": "789",
		"monitor_groups": []interface{}{
			"234",
			"567",
		},
		"dependency_resource_ids": []interface{}{
			"234",
			"567",
		},
		"user_group_ids": []interface{}{
			"123",
			"456",
		},
		"third_party_service_ids": []interface{}{
			"123",
			"456",
		},
		"on_call_schedule_id": "1244",
	})
}
