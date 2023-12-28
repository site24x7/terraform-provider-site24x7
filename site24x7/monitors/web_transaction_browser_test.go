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

func TestWebTransactionBrowserMonitorCreate(t *testing.T) {
	d := webTransactionBrowserMonitorTestResourceData(t)

	c := fake.NewClient()

	a := &api.WebTransactionBrowserMonitor{
		DisplayName: "RBM-Terraform",
		Type:        string(api.REALBROWSER),
		BaseURL:     "https://www.example.com/",
		//AsyncDCEnabled:     false,
		CheckFrequency:     "15",
		IgnoreCertError:    false,
		IPType:             0,
		SeleniumScript:     "Script for the monitor",
		ScriptType:         "txt",
		ThresholdProfileID: "789",
		PageLoadTime:       0,
		PerformAutomation:  false,
		Resolution:         "1600,900",
		MonitorGroups:      []string{"234", "567"},
		UserGroupIDs:       []string{"123", "456"},
		TagIDs:             []string{"123"},
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

	c.FakeWebTransactionBrowserMonitors.On("Create", a).Return(a, nil).Once()

	require.NoError(t, webTransactionBrowserMonitorCreate(d, c))

	c.FakeWebTransactionBrowserMonitors.On("Create	", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := webTransactionBrowserMonitorCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestWebTransactionBrowserMonitorUpdate(t *testing.T) {
	d := webTransactionBrowserMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.WebTransactionBrowserMonitor{
		MonitorID:          "123",
		DisplayName:        "RBM-Terraform",
		Type:               string(api.REALBROWSER),
		BaseURL:            "https://www.example.com/",
		AsyncDCEnabled:     false,
		CheckFrequency:     "15",
		IgnoreCertError:    false,
		IPType:             0,
		SeleniumScript:     "Script for the monitor",
		ScriptType:         "txt",
		ThresholdProfileID: "789",
		PageLoadTime:       0,
		PerformAutomation:  false,
		Resolution:         "1600,900",
		MonitorGroups:      []string{"234", "567"},
		UserGroupIDs:       []string{"123", "456"},
		TagIDs:             []string{"123"},
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

	c.FakeWebTransactionBrowserMonitors.On("Update", a).Return(a, nil).Once()

	require.NoError(t, webTransactionBrowserMonitorUpdate(d, c))

	c.FakeWebTransactionBrowserMonitors.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := webTransactionBrowserMonitorUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestWebTransactionBrowserMonitorRead(t *testing.T) {
	d := webTransactionBrowserMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeWebTransactionBrowserMonitors.On("Get", "123").Return(&api.DomainExpiryMonitor{}, nil).Once()

	require.NoError(t, webTransactionBrowserMonitorRead(d, c))

	c.FakeWebTransactionBrowserMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := webTransactionBrowserMonitorRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestWebTransactionBrowserMonitorDelete(t *testing.T) {
	d := webTransactionBrowserMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeWebTransactionBrowserMonitors.On("Delete", "123").Return(nil).Once()

	require.NoError(t, webTransactionBrowserMonitorDelete(d, c))

	c.FakeWebTransactionBrowserMonitors.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, webTransactionBrowserMonitorDelete(d, c))
}

func TestWebTransactionBrowserMonitorExists(t *testing.T) {
	d := webTransactionBrowserMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeWebTransactionBrowserMonitors.On("Get", "123").Return(&api.WebTransactionBrowserMonitor{}, nil).Once()

	exists, err := webTransactionBrowserMonitorExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeWebTransactionBrowserMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = webTransactionBrowserMonitorExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeWebTransactionBrowserMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = webTransactionBrowserMonitorExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func webTransactionBrowserMonitorTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, WebTransactionBrowserMonitorSchema, map[string]interface{}{
		"display_name":         "RBM-Terraform",
		"type":                 string(api.REALBROWSER),
		"base_url":             "https://www.example.com/",
		"async_dc_enabled":     false,
		"browser_version":      0,
		"check_frequency":      "15",
		"ignore_cert_err":      false,
		"ip_type":              0,
		"threshold_profile_id": "789",
		"page_load_time":       0,
		"perform_automation":   false,
		"resolution":           "1600,900",
		"location_profile_id":  "456",
		"monitor_groups":       []string{"234", "567"},
		"user_group_ids":       []string{"123", "456"},
		"tag_ids":              []string{"123"},
	},
	)
}
