package site24x7

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSSLMonitorCreate(t *testing.T) {
	d := sslTestResourceData(t)

	c := fake.NewClient()

	a := &api.SSLMonitor{
		DisplayName:           "foo",
		DomainName:            "www.example.com",
		Type:                  "SSL_CERT",
		Timeout:               30,
		Protocol:              "HTTPS",
		Port:                  443,
		ExpireDays:            30,
		HTTPProtocolVersion:   "H1.1",
		LocationProfileID:     "456",
		NotificationProfileID: "789",
		ThresholdProfileID:    "012",
		MonitorGroups:         []string{"234", "567"},
		UserGroupIDs:          []string{"123", "456"},
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
	c.FakeSSLMonitors.On("Create", a).Return(a, nil).Once()

	require.NoError(t, sslMonitorCreate(d, c))

	c.FakeSSLMonitors.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := sslMonitorCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestSSLMonitorUpdate(t *testing.T) {
	d := sslTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.SSLMonitor{
		MonitorID:             "123",
		DisplayName:           "foo",
		DomainName:            "www.example.com",
		Type:                  "SSL_CERT",
		Timeout:               30,
		Protocol:              "HTTPS",
		Port:                  443,
		ExpireDays:            30,
		HTTPProtocolVersion:   "H1.1",
		IgnoreDomainMismatch:  false,
		IgnoreTrust:           false,
		LocationProfileID:     "456",
		NotificationProfileID: "789",
		ThresholdProfileID:    "012",
		MonitorGroups:         []string{"234", "567"},
		UserGroupIDs:          []string{"123", "456"},
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

	c.FakeSSLMonitors.On("Update", a).Return(a, nil).Once()

	require.NoError(t, sslMonitorUpdate(d, c))

	c.FakeSSLMonitors.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := sslMonitorUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestSSLMonitorRead(t *testing.T) {
	d := sslTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeSSLMonitors.On("Get", "123").Return(&api.SSLMonitor{}, nil).Once()

	require.NoError(t, sslMonitorRead(d, c))

	c.FakeSSLMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := sslMonitorRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestSSLMonitorDelete(t *testing.T) {
	d := sslTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeSSLMonitors.On("Delete", "123").Return(nil).Once()

	require.NoError(t, sslMonitorDelete(d, c))

	c.FakeSSLMonitors.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, sslMonitorDelete(d, c))
}

func TestSSLMonitorExists(t *testing.T) {
	d := sslTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeSSLMonitors.On("Get", "123").Return(&api.SSLMonitor{}, nil).Once()

	exists, err := sslMonitorExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeSSLMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = sslMonitorExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeSSLMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = sslMonitorExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func sslTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, SSLMonitorSchema, map[string]interface{}{
		"display_name":            "foo",
		"type":                    "SSL_CERT",
		"domain_name":             "www.example.com",
		"timeout":                 30,
		"protocol":                "HTTPS",
		"port":                    443,
		"expire_days":             30,
		"http_protocol_version":   "H1.1",
		"ignore_domain_mismatch":  false,
		"ignore_trust":            false,
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
	})
}
