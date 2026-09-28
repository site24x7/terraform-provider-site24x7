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

func TestPINGMonitorCreate(t *testing.T) {
	d := pingMonitorTestResourceData(t)

	c := fake.NewClient()

	a := &api.PINGMonitor{
		DisplayName:       "PING Monitor",
		Type:              string(api.PING),
		HostName:          "www.example.com",
		Timeout:           10,
		UseIPV6:           true,
		OnCallScheduleID:  "234",
		PerformAutomation: false,
		// check_frequency is not in the test's ResourceData; "5" is the schema default.
		CheckFrequency:        "5",
		LocationProfileID:     "456",
		NotificationProfileID: "789",
		ThresholdProfileID:    "012",
		MonitorGroups:         []string{"234", "567"},
		DependencyResourceIDs: []string{},
		UserGroupIDs:          []string{"123", "456"},
		TagIDs:                []string{"123"},
		ActionIDs:             []api.ActionRef{},
	}

	locationProfiles := []*api.LocationProfile{
		{
			ProfileID:   "123",
			ProfileName: "Location Profile",
		},
		{
			ProfileID:   "456",
			ProfileName: "TEST",
		},
	}
	c.FakeLocationProfiles.On("List").Return(locationProfiles, nil)

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

	c.FakePINGMonitors.On("Create", a).Return(a, nil).Once()

	require.NoError(t, pingMonitorCreate(d, c))

	c.FakePINGMonitors.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := pingMonitorCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestPINGMonitorUpdate(t *testing.T) {
	d := pingMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.PINGMonitor{
		MonitorID:             "123",
		DisplayName:           "PING Monitor",
		Type:                  string(api.PING),
		HostName:              "www.example.com",
		Timeout:               10,
		UseIPV6:               true,
		OnCallScheduleID:      "234",
		CheckFrequency:        "5",
		LocationProfileID:     "456",
		NotificationProfileID: "789",
		ThresholdProfileID:    "012",
		MonitorGroups:         []string{"234", "567"},
		DependencyResourceIDs: []string{},
		UserGroupIDs:          []string{"123", "456"},
		TagIDs:                []string{"123"},
		ActionIDs:             []api.ActionRef{},
		// ActionIDs: []api.ActionRef{
		// 	{
		// 		ActionID:  "123action",
		// 		AlertType: 1,
		// 	},
		// 	{
		// 		ActionID:  "234action",
		// 		AlertType: 5,
		// 	},
		// },
		// MatchingKeyword: map[string]interface{}{
		// 	"severity": "2",
		// 	"value":    "aaa",
		// },
		// UnmatchingKeyword: map[string]interface{}{
		// 	"severity": "2",
		// 	"value":    "bbb",
		// },
		// MatchRegex: map[string]interface{}{
		// 	"severity": "0",
		// 	"value":    "*.a.*",
		// },
	}

	locationProfiles := []*api.LocationProfile{
		{
			ProfileID:   "123",
			ProfileName: "Location Profile",
		},
		{
			ProfileID:   "456",
			ProfileName: "TEST",
		},
	}
	c.FakeLocationProfiles.On("List").Return(locationProfiles, nil)

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

	c.FakePINGMonitors.On("Update", a).Return(a, nil).Once()

	require.NoError(t, pingMonitorUpdate(d, c))

	c.FakePINGMonitors.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := pingMonitorUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestPINGMonitorRead(t *testing.T) {
	d := pingMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakePINGMonitors.On("Get", "123").Return(&api.PINGMonitor{}, nil).Once()

	require.NoError(t, pingMonitorRead(d, c))

	c.FakePINGMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := pingMonitorRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestPINGMonitorDelete(t *testing.T) {
	d := pingMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakePINGMonitors.On("Delete", "123").Return(nil).Once()

	require.NoError(t, pingMonitorDelete(d, c))

	c.FakePINGMonitors.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, pingMonitorDelete(d, c))
}

func TestPINGMonitorExists(t *testing.T) {
	d := pingMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakePINGMonitors.On("Get", "123").Return(&api.PINGMonitor{}, nil).Once()

	exists, err := pingMonitorExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakePINGMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = pingMonitorExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakePINGMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = pingMonitorExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func pingMonitorTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, PINGMonitorSchema, map[string]interface{}{
		// expire_days and ignore_registry_date are domain-expiry keys that
		// PINGMonitorSchema does not declare. threshold_profile_id is set so the
		// resource does not fall back to looking one up, which needs a mock the
		// test never registered.
		"display_name":            "PING Monitor",
		"host_name":               "www.example.com",
		"timeout":                 10,
		"use_ipv6":                true,
		"on_call_schedule_id":     "234",
		"location_profile_id":     "456",
		"notification_profile_id": "789",
		"threshold_profile_id":    "012",
		"monitor_groups": []interface{}{
			"234",
			"567",
		},
		"user_group_ids": []interface{}{
			"123",
			"456",
		},
		"tag_ids": []interface{}{
			"123",
		},
	},
	)
}
