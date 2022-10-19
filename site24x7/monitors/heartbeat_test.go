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

func TestHeartbeatMonitorCreate(t *testing.T) {
	d := heartbeatTestResourceData(t)

	c := fake.NewClient()

	a := &api.HeartbeatMonitor{
		DisplayName:           "foo",
		NameInPingURL:         "status_check",
		Type:                  "HEARTBEAT",
		ThresholdProfileID:    "012",
		NotificationProfileID: "789",
		MonitorGroups:         []string{"234", "567"},
		UserGroupIDs:          []string{"123", "456"},
		TagIDs:                []string{"123"},
		ThirdPartyServiceIDs:  []string{"123", "456"},
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

	c.FakeHeartbeatMonitors.On("Create", a).Return(a, nil).Once()

	require.NoError(t, heartbeatMonitorCreate(d, c))

	c.FakeHeartbeatMonitors.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := heartbeatMonitorCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestHeartbeatMonitorUpdate(t *testing.T) {
	d := heartbeatTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.HeartbeatMonitor{
		MonitorID:             "123",
		DisplayName:           "foo",
		NameInPingURL:         "status_check",
		Type:                  "HEARTBEAT",
		ThresholdProfileID:    "012",
		NotificationProfileID: "789",
		MonitorGroups:         []string{"234", "567"},
		UserGroupIDs:          []string{"123", "456"},
		TagIDs:                []string{"123"},
		ThirdPartyServiceIDs:  []string{"123", "456"},
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

	c.FakeHeartbeatMonitors.On("Update", a).Return(a, nil).Once()

	require.NoError(t, heartbeatMonitorUpdate(d, c))

	c.FakeHeartbeatMonitors.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := heartbeatMonitorUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestHeartbeatMonitorRead(t *testing.T) {
	d := heartbeatTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeHeartbeatMonitors.On("Get", "123").Return(&api.HeartbeatMonitor{}, nil).Once()

	require.NoError(t, heartbeatMonitorRead(d, c))

	c.FakeHeartbeatMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := heartbeatMonitorRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestHeartbeatMonitorDelete(t *testing.T) {
	d := heartbeatTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeHeartbeatMonitors.On("Delete", "123").Return(nil).Once()

	require.NoError(t, heartbeatMonitorDelete(d, c))

	c.FakeHeartbeatMonitors.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, heartbeatMonitorDelete(d, c))
}

func TestHeartbeatMonitorExists(t *testing.T) {
	d := heartbeatTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeHeartbeatMonitors.On("Get", "123").Return(&api.HeartbeatMonitor{}, nil).Once()

	exists, err := heartbeatMonitorExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeHeartbeatMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = heartbeatMonitorExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeHeartbeatMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = heartbeatMonitorExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func heartbeatTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, HeartbeatMonitorSchema, map[string]interface{}{
		"display_name":            "foo",
		"type":                    "HEARTBEAT",
		"name_in_ping_url":        "status_check",
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
