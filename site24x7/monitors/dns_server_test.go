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

func TestDNSServerMonitorCreate(t *testing.T) {
	d := dnsServerMonitorTestResourceData(t)

	c := fake.NewClient()

	a := &api.DNSServerMonitor{
		DisplayName:           "foo",
		Type:                  "DNS",
		DNSHost:               "8.8.8.8",
		DNSPort:               "53",
		UseIPV6:               false,
		DomainName:            "global.realtime.primary.cartography.cluster.ably-nonprod.net",
		CheckFrequency:        "5",
		Timeout:               10,
		LocationProfileID:     "456",
		NotificationProfileID: "789",
		ThresholdProfileID:    "012",
		UserGroupIDs:          []string{"123", "456"},
		MonitorGroups:         []string{"234", "567"},
		DependencyResourceIDs: []string{"234", "567"},
		LookupType:            1,
		DNSSEC:                false,
		TagIDs:                []string{"123"},
		DeepDiscovery:         false,
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
		SearchConfig: []api.SearchConfig{},
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
			TagName:  "DNSServer tag",
			TagValue: "baz 1",
			TagColor: "#B7DA9E",
		},
	}
	c.FakeTags.On("List").Return(tags, nil)

	c.FakeDNSServerMonitors.On("Create", a).Return(a, nil).Once()

	require.NoError(t, dnsServerMonitorCreate(d, c))

	c.FakeDNSServerMonitors.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := dnsServerMonitorCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestDNSServerMonitorUpdate(t *testing.T) {
	d := dnsServerMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.DNSServerMonitor{
		MonitorID:             "123",
		DisplayName:           "foo",
		Type:                  "DNS",
		DNSHost:               "8.8.8.8",
		DNSPort:               "53",
		UseIPV6:               false,
		DomainName:            "global.realtime.primary.cartography.cluster.ably-nonprod.net",
		CheckFrequency:        "5",
		Timeout:               10,
		LocationProfileID:     "456",
		NotificationProfileID: "789",
		ThresholdProfileID:    "012",
		UserGroupIDs:          []string{"123", "456"},
		MonitorGroups:         []string{"234", "567"},
		DependencyResourceIDs: []string{"234", "567"},
		LookupType:            1,
		DNSSEC:                false,
		TagIDs:                []string{"123"},
		DeepDiscovery:         false,
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
		SearchConfig: []api.SearchConfig{},
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
			TagName:  "DNSServer tag",
			TagValue: "baz 1",
			TagColor: "#B7DA9E",
		},
	}
	c.FakeTags.On("List").Return(tags, nil)

	c.FakeDNSServerMonitors.On("Update", a).Return(a, nil).Once()

	require.NoError(t, dnsServerMonitorUpdate(d, c))

	c.FakeDNSServerMonitors.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := dnsServerMonitorUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestDNSServerMonitorRead(t *testing.T) {
	d := dnsServerMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeDNSServerMonitors.On("Get", "123").Return(&api.DNSServerMonitor{}, nil).Once()

	require.NoError(t, dnsServerMonitorRead(d, c))

	c.FakeDNSServerMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := dnsServerMonitorRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestDNSServerMonitorDelete(t *testing.T) {
	d := dnsServerMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeDNSServerMonitors.On("Delete", "123").Return(nil).Once()

	require.NoError(t, dnsServerMonitorDelete(d, c))

	c.FakeDNSServerMonitors.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, dnsServerMonitorDelete(d, c))
}

func TestDNSServerMonitorExists(t *testing.T) {
	d := dnsServerMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeDNSServerMonitors.On("Get", "123").Return(&api.DNSServerMonitor{}, nil).Once()

	exists, err := dnsServerMonitorExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeDNSServerMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = dnsServerMonitorExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeDNSServerMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = dnsServerMonitorExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func dnsServerMonitorTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, dnsServerMonitorSchema, map[string]interface{}{
		"display_name":            "foo",
		"type":                    "DNS",
		"dns_host":                "8.8.8.8",
		"dns_port":                53,
		"use_ipv6":                false,
		"domain_name":             "global.realtime.primary.cartography.cluster.ably-nonprod.net",
		"check_frequency":         "5",
		"timeout":                 10,
		"lookup_type":             1,
		"dnssec":                  false,
		"deep_discovery":          false,
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
		"search_config": map[string]interface{}{
			"addr": "1.2.3.4",
			"ttl": 60,
			"ttlo": 60,
		},
	})
}
