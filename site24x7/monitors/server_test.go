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

func TestServerMonitorUpdate(t *testing.T) {
	d := sslTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.ServerMonitor{
		MonitorID:             "123",
		DisplayName:           "foo",
		Type:                  "SERVER",
		NotificationProfileID: "789",
		ThresholdProfileID:    "012",
		MonitorGroups:         []string{"234", "567"},
		UserGroupIDs:          []string{"123", "456"},
		TagIDs:                []string{"123"},
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

	c.FakeServerMonitors.On("Update", a).Return(a, nil).Once()

	require.NoError(t, serverMonitorUpdate(d, c))

	c.FakeServerMonitors.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := serverMonitorUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestServerMonitorRead(t *testing.T) {
	d := sslTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeServerMonitors.On("Get", "123").Return(&api.ServerMonitor{}, nil).Once()

	require.NoError(t, serverMonitorRead(d, c))

	c.FakeServerMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := serverMonitorRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestServerMonitorDelete(t *testing.T) {
	d := sslTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeServerMonitors.On("Delete", "123").Return(nil).Once()

	require.NoError(t, serverMonitorDelete(d, c))

	c.FakeServerMonitors.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, serverMonitorDelete(d, c))
}

func TestServerMonitorExists(t *testing.T) {
	d := sslTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeServerMonitors.On("Get", "123").Return(&api.ServerMonitor{}, nil).Once()

	exists, err := serverMonitorExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeServerMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = serverMonitorExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeServerMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = serverMonitorExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func serverTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, ServerMonitorSchema, map[string]interface{}{
		"display_name":            "foo",
		"type":                    "SERVER",
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
