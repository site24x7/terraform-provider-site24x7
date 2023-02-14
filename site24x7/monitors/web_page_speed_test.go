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

func TestWebPageSpeedMonitorCreate(t *testing.T) {
	d := webPageSpeedMonitorTestResourceData(t)

	c := fake.NewClient()

	a := &api.WebPageSpeedMonitor{
		DisplayName:           "foo",
		Type:                  "HOMEPAGE",
		Website:               "www.test.tld",
		BrowserType:           1,
		BrowserVersion:		   83,
		WebsiteType:           1,
		DeviceType:            "1",
		WPAResolution:         "1024,768",
		CheckFrequency:        "5",
		HTTPMethod:            "G",
		Timeout:               10,
		LocationProfileID:     "456",
		NotificationProfileID: "789",
		ThresholdProfileID:    "012",
		MatchCase:             true,
		UserAgent:             "firefox",
		MonitorGroups:         []string{"234", "567"},
		DependencyResourceIDs: []string{"234", "567"},
		UserGroupIDs:          []string{"123", "456"},
		TagIDs:                []string{"123"},
		AuthUser:              "username",
		AuthPass:              "password",
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
		ActionIDs: []api.ActionRef{
			{
				ActionID:  "123action",
				AlertType: 1,
			},
			{
				ActionID:  "234action",
				AlertType: 5,
			},
		},
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

	c.FakeWebPageSpeedMonitors.On("Create", a).Return(a, nil).Once()

	require.NoError(t, webPageSpeedMonitorCreate(d, c))

	c.FakeWebPageSpeedMonitors.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := webPageSpeedMonitorCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestWebPageSpeedMonitorUpdate(t *testing.T) {
	d := webPageSpeedMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.WebPageSpeedMonitor{
		MonitorID:             "123",
		DisplayName:           "foo",
		Type:                  "HOMEPAGE",
		BrowserType:           1,
		BrowserVersion:		   83,
		WebsiteType:           1,
		DeviceType:            "1",
		WPAResolution:         "1024,768",
		Website:               "www.test.tld",
		CheckFrequency:        "5",
		HTTPMethod:            "G",
		Timeout:               10,
		LocationProfileID:     "456",
		NotificationProfileID: "789",
		ThresholdProfileID:    "012",
		MatchCase:             true,
		UserAgent:             "firefox",
		MonitorGroups:         []string{"234", "567"},
		DependencyResourceIDs: []string{"234", "567"},
		UserGroupIDs:          []string{"123", "456"},
		TagIDs:                []string{"123"},
		AuthUser:              "username",
		AuthPass:              "password",
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
		ActionIDs: []api.ActionRef{
			{
				ActionID:  "123action",
				AlertType: 1,
			},
			{
				ActionID:  "234action",
				AlertType: 5,
			},
		},
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

	c.FakeWebPageSpeedMonitors.On("Update", a).Return(a, nil).Once()

	require.NoError(t, webPageSpeedMonitorUpdate(d, c))

	c.FakeWebPageSpeedMonitors.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := webPageSpeedMonitorUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestWebPageSpeedMonitorRead(t *testing.T) {
	d := webPageSpeedMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeWebPageSpeedMonitors.On("Get", "123").Return(&api.WebPageSpeedMonitor{}, nil).Once()

	require.NoError(t, webPageSpeedMonitorRead(d, c))

	c.FakeWebPageSpeedMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := webPageSpeedMonitorRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestWebPageSpeedMonitorDelete(t *testing.T) {
	d := webPageSpeedMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeWebPageSpeedMonitors.On("Delete", "123").Return(nil).Once()

	require.NoError(t, webPageSpeedMonitorDelete(d, c))

	c.FakeWebPageSpeedMonitors.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, webPageSpeedMonitorDelete(d, c))
}

func TestWebPageSpeedMonitorExists(t *testing.T) {
	d := webPageSpeedMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeWebPageSpeedMonitors.On("Get", "123").Return(&api.WebPageSpeedMonitor{}, nil).Once()

	exists, err := webPageSpeedMonitorExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeWebPageSpeedMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = webPageSpeedMonitorExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeWebPageSpeedMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = webPageSpeedMonitorExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func webPageSpeedMonitorTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, webPageSpeedMonitorSchema, map[string]interface{}{
		"display_name":    "foo",
		"type":            "HOMEPAGE",
		"website":         "www.test.tld",
		"check_frequency": "5",
		"http_method":     "G",
		"website_type":    1,
		"browser_type":    1,
		"browser_version": 83,
		"device_type":     "1",
		"wpa_resolution":  "1024,768",
		"auth_user":       "username",
		"auth_pass":       "password",
		"match_case":      true,
		"user_agent":      "firefox",
		"custom_headers": map[string]interface{}{
			"Accept":        "application/json",
			"Cache-Control": "nocache",
		},
		"timeout":                 10,
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
		"actions": map[string]interface{}{
			"1": "123action",
			"5": "234action",
		},
	})
}
