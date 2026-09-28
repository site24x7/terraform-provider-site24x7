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

func TestISPMonitorCreate(t *testing.T) {
	d := ispTestResourceData(t)

	c := fake.NewClient()

	a := &api.ISPMonitor{
		DisplayName:           "ISP Monitor",
		Hostname:              "www.example.com",
		UseIPV6:               true,
		Type:                  "ISP",
		Timeout:               30,
		Protocol:              "1",
		Port:                  443,
		CheckFrequency:        "5",
		LocationProfileID:     "456",
		NotificationProfileID: "789",
		ThresholdProfileID:    "012",
		MonitorGroups:         []string{"234", "567"},
		DependencyResourceIDs: []string{"234", "567"},
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

	c.FakeISPMonitors.On("Create", a).Return(a, nil).Once()

	require.NoError(t, ispMonitorCreate(d, c))

	c.FakeISPMonitors.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := ispMonitorCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestISPMonitorUpdate(t *testing.T) {
	d := ispTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.ISPMonitor{
		MonitorID:             "123",
		DisplayName:           "ISP Monitor",
		Hostname:              "www.example.com",
		UseIPV6:               true,
		Type:                  "ISP",
		Timeout:               30,
		Protocol:              "1",
		Port:                  443,
		CheckFrequency:        "5",
		LocationProfileID:     "456",
		NotificationProfileID: "789",
		ThresholdProfileID:    "012",
		MonitorGroups:         []string{"234", "567"},
		DependencyResourceIDs: []string{"234", "567"},
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

	c.FakeISPMonitors.On("Update", a).Return(a, nil).Once()

	require.NoError(t, ispMonitorUpdate(d, c))

	c.FakeISPMonitors.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := ispMonitorUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestISPMonitorRead(t *testing.T) {
	d := ispTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeISPMonitors.On("Get", "123").Return(&api.ISPMonitor{}, nil).Once()

	require.NoError(t, ispMonitorRead(d, c))

	c.FakeISPMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := ispMonitorRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestISPMonitorDelete(t *testing.T) {
	d := ispTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeISPMonitors.On("Delete", "123").Return(nil).Once()

	require.NoError(t, ispMonitorDelete(d, c))

	c.FakeISPMonitors.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, ispMonitorDelete(d, c))
}

func TestISPMonitorExists(t *testing.T) {
	d := ispTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeISPMonitors.On("Get", "123").Return(&api.ISPMonitor{}, nil).Once()

	exists, err := ispMonitorExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeISPMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = ispMonitorExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeISPMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = ispMonitorExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func ispTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, ISPMonitorSchema, map[string]interface{}{
		// These were SSL monitor keys (domain_name, expire_days,
		// http_protocol_version, ignore_domain_mismatch, ignore_trust), copied
		// from ssl_test.go along with the sslMonitorCreate calls. ISPMonitorSchema
		// declares none of them, so every one of those values was dropped.
		"display_name":            "ISP Monitor",
		"type":                    "ISP",
		"hostname":                "www.example.com",
		"use_ipv6":                true,
		"timeout":                 30,
		"protocol":                "1",
		"port":                    443,
		"check_frequency":         "5",
		"location_profile_id":     "456",
		"notification_profile_id": "789",
		"threshold_profile_id":    "012",
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
		"tag_ids": []interface{}{
			"123",
		},
	})
}
