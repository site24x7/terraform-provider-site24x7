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

func TestDomainExpiryMonitorCreate(t *testing.T) {
	d := domainExpiryMonitorTestResourceData(t)

	c := fake.NewClient()

	a := &api.DomainExpiryMonitor{
		DisplayName:           "foo",
		Type:                  string(api.DOMAINEXPIRY),
		HostName:              "www.example.com",
		Timeout:               10,
		LocationProfileID:     "456",
		NotificationProfileID: "789",
		MonitorGroups:         []string{"234", "567"},
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

	c.FakeDomainExpiryMonitors.On("Create", a).Return(a, nil).Once()

	require.NoError(t, domainExpiryMonitorCreate(d, c))

	c.FakeDomainExpiryMonitors.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := domainExpiryMonitorCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestDomainExpiryMonitorUpdate(t *testing.T) {
	d := domainExpiryMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.DomainExpiryMonitor{
		MonitorID:             "123",
		DisplayName:           "foo",
		Type:                  string(api.DOMAINEXPIRY),
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

	c.FakeDomainExpiryMonitors.On("Update", a).Return(a, nil).Once()

	require.NoError(t, domainExpiryMonitorUpdate(d, c))

	c.FakeDomainExpiryMonitors.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := domainExpiryMonitorUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestDomainExpiryMonitorRead(t *testing.T) {
	d := domainExpiryMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeDomainExpiryMonitors.On("Get", "123").Return(&api.DomainExpiryMonitor{}, nil).Once()

	require.NoError(t, domainExpiryMonitorRead(d, c))

	c.FakeDomainExpiryMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := domainExpiryMonitorRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestDomainExpiryMonitorDelete(t *testing.T) {
	d := domainExpiryMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeDomainExpiryMonitors.On("Delete", "123").Return(nil).Once()

	require.NoError(t, domainExpiryMonitorDelete(d, c))

	c.FakeDomainExpiryMonitors.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, domainExpiryMonitorDelete(d, c))
}

func TestDomainExpiryMonitorExists(t *testing.T) {
	d := domainExpiryMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeDomainExpiryMonitors.On("Get", "123").Return(&api.DomainExpiryMonitor{}, nil).Once()

	exists, err := domainExpiryMonitorExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeDomainExpiryMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = domainExpiryMonitorExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeDomainExpiryMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = domainExpiryMonitorExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func domainExpiryMonitorTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, DomainExpiryMonitorSchema, map[string]interface{}{
		"display_name":            "foo",
		"type":                    "RESTAPI",
		"host_name":               "www.example.com",
		"port":                    43,
		"timeout":                 30,
		"expire_days":             30,
		"location_profile_id":     "456",
		"notification_profile_id": "789",
		"on_call_schedule_id":     "234",
		"ignore_registry_date":    false,
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
