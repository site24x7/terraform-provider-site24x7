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

func TestSOAPMonitorCreate(t *testing.T) {
	d := soapMonitorTestResourceData(t)

	c := fake.NewClient()

	a := &api.SOAPMonitor{
		DisplayName:    "SOAP Monitor",
		Website:        "www.example.com",
		RequestParam:   "",
		Type:           "SOAP",
		UseIPV6:        true,
		SSLProtocol:    "",
		Timeout:        10,
		HTTPMethod:     "",
		CheckFrequency: "5",
		ResponseHeaders: api.HTTPResponseHeader{
			Severity: api.Trouble,
			Value: []api.Header{
				{
					Name:  "Accept-Encoding",
					Value: "gzip",
				},
				{
					Name:  "Cache-Control",
					Value: "nocache",
				},
			},
		},
		OnCallScheduleID:      "23524543545245",
		LocationProfileID:     "123412341234123412",
		NotificationProfileID: "123412341234123412",
		MonitorGroups:         []string{"234", "567"},
		DependencyResourceIDs: []string{"123", "456"},
		UserGroupIDs:          []string{"123", "456"},
		PerformAutomation:     true,
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

	c.FakeSOAPMonitors.On("Create", a).Return(a, nil).Once()

	require.NoError(t, soapMonitorCreate(d, c))

	c.FakeSOAPMonitors.On("Create	", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := soapMonitorCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestSOAPMonitorUpdate(t *testing.T) {
	d := soapMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.SOAPMonitor{
		MonitorID:             "123",
		DisplayName:           "foo",
		Type:                  string(api.SOAP),
		LocationProfileID:     "456",
		NotificationProfileID: "789",
		MonitorGroups:         []string{"234", "567"},
		UserGroupIDs:          []string{"123", "456"},
		TagIDs:                []string{"123"},
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

	c.FakeSOAPMonitors.On("Update", a).Return(a, nil).Once()

	require.NoError(t, soapMonitorUpdate(d, c))

	c.FakeSOAPMonitors.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := soapMonitorUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestSOAPMonitorRead(t *testing.T) {
	d := soapMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeSOAPMonitors.On("Get", "123").Return(&api.SOAPMonitor{}, nil).Once()

	require.NoError(t, soapMonitorRead(d, c))

	c.FakeSOAPMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := soapMonitorRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestSOAPMonitorDelete(t *testing.T) {
	d := soapMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeSOAPMonitors.On("Delete", "123").Return(nil).Once()

	require.NoError(t, soapMonitorDelete(d, c))

	c.FakeSOAPMonitors.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, soapMonitorDelete(d, c))
}

func TestSOAPMonitorExists(t *testing.T) {
	d := soapMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeSOAPMonitors.On("Get", "123").Return(&api.SOAPMonitor{}, nil).Once()

	exists, err := soapMonitorExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeSOAPMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = soapMonitorExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeSOAPMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = soapMonitorExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func soapMonitorTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, SOAPMonitorSchema, map[string]interface{}{
		"display_name":            "SOAP Monitor",
		"website":                 "www.example.com",
		"timeout":                 0,
		"on_call_schedule_id":     "234",
		"ignore_registry_date":    false,
		"location_profile_id":     "456",
		"notification_profile_id": "789",
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
