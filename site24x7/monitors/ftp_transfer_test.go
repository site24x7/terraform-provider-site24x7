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

func TestFTPTransferMonitorCreate(t *testing.T) {
	d := ftpTransferTestResourceData(t)

	c := fake.NewClient()

	a := &api.FTPTransferMonitor{
		DisplayName:           "FTP Transfer Monitor",
		HostName:              "www.example.com",
		Protocol:              "FTP",
		Type:                  "FTP",
		Port:                  443,
		CheckFrequency:        "5",
		Timeout:               30,
		CheckUpload:           true,
		CheckDownload:         true,
		Username:              "sas",
		Password:              "sas",
		Destination:           "/Home/sas/",
		PerformAutomation:     true,
		CredentialProfileID:   "2345536536",
		OnCallScheduleID:      "8687567555",
		LocationProfileID:     "456",
		NotificationProfileID: "789",
		ThresholdProfileID:    "012",
		MonitorGroups:         []string{"234", "567"},
		DependencyResourceIDs: []string{"234", "567"},
		UserGroupIDs:          []string{"123", "456"},
		TagIDs:                []string{"123"},
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

	c.FakeFTPTransferMonitors.On("Create", a).Return(a, nil).Once()

	require.NoError(t, sslMonitorCreate(d, c))

	c.FakeFTPTransferMonitors.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := ftpTransferMonitorCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestFTPTransferMonitorUpdate(t *testing.T) {
	d := ftpTransferTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.FTPTransferMonitor{
		DisplayName:           "FTP Transfer Monitor",
		HostName:              "www.example.com",
		Protocol:              "FTP",
		Type:                  "FTP",
		Port:                  443,
		CheckFrequency:        "5",
		Timeout:               30,
		CheckUpload:           true,
		CheckDownload:         true,
		Username:              "sas",
		Password:              "sas",
		Destination:           "/Home/sas/",
		PerformAutomation:     true,
		CredentialProfileID:   "2345536536",
		OnCallScheduleID:      "8687567555",
		LocationProfileID:     "456",
		NotificationProfileID: "789",
		ThresholdProfileID:    "012",
		MonitorGroups:         []string{"234", "567"},
		DependencyResourceIDs: []string{"234", "567"},
		UserGroupIDs:          []string{"123", "456"},
		TagIDs:                []string{"123"},
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

	c.FakeFTPTransferMonitors.On("Update", a).Return(a, nil).Once()

	require.NoError(t, sslMonitorUpdate(d, c))

	c.FakeFTPTransferMonitors.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := sslMonitorUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestFTPTransferMonitorRead(t *testing.T) {
	d := ftpTransferTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeFTPTransferMonitors.On("Get", "123").Return(&api.FTPTransferMonitor{}, nil).Once()

	require.NoError(t, ftpTransferMonitorRead(d, c))

	c.FakeFTPTransferMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := ftpTransferMonitorRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestFTPTransferMonitorDelete(t *testing.T) {
	d := ftpTransferTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeFTPTransferMonitors.On("Delete", "123").Return(nil).Once()

	require.NoError(t, ftpTransferMonitorDelete(d, c))

	c.FakeFTPTransferMonitors.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, ftpTransferMonitorDelete(d, c))
}

func TestFTPTransferMonitorExists(t *testing.T) {
	d := ftpTransferTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeFTPTransferMonitors.On("Get", "123").Return(&api.FTPTransferMonitor{}, nil).Once()

	exists, err := ftpTransferMonitorExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeFTPTransferMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = ftpTransferMonitorExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeFTPTransferMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = ftpTransferMonitorExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func ftpTransferTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, FTPTransferMonitorSchema, map[string]interface{}{
		"display_name":            "FTP Monitor",
		"type":                    "FTP",
		"host_name":               "www.example.com",
		"timeout":                 30,
		"protocol":                "HTTPS",
		"port":                    443,
		"check_frequency":         "5",
		"check_upload":            true,
		"check_download":          true,
		"username":                "sas",
		"password":                "sas",
		"destination":             "/home/sas",
		"perform_automation":      true,
		"credential_profile_id":   "234354543523",
		"on_call_schedule_id":     "232432423",
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
	})
}
