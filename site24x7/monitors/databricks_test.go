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

func databricksIntPtr(i int) *int    { return &i }
func databricksBoolPtr(b bool) *bool { return &b }

func expectedDatabricksMonitor(id string, updateAuth bool) *api.DatabricksMonitor {
	return &api.DatabricksMonitor{
		MonitorID:             id,
		DisplayName:           "Databricks Workspace",
		Type:                  "DBRICKS_WS",
		CheckFrequency:        "30",
		AuthType:              databricksIntPtr(4),
		TokenProfile:          "databricks_oidc",
		WorkspaceURI:          "https://adb-1234567890123456.7.azuredatabricks.net/",
		AccessKey:             "client-id",
		SecretKey:             "client-secret",
		MonitoredServices:     []int{5002},
		DisableDiscovery:      databricksIntPtr(0),
		ThresholdProfileID:    "threshold-456",
		NotificationProfileID: "notif-789",
		UserGroupIDs:          []string{"ug-1"},
		TagIDs:                []string{},
		UpdateAuthDetails:     databricksBoolPtr(updateAuth),
	}
}

// What GET /monitors/{id} returns: no credentials, matching the sample create response.
func databricksAPIResponse(id string) *api.DatabricksMonitor {
	return &api.DatabricksMonitor{
		MonitorID:             id,
		DisplayName:           "Databricks Workspace",
		Type:                  "DBRICKS_WS",
		CheckFrequency:        "30",
		ThresholdProfileID:    "threshold-456",
		NotificationProfileID: "notif-789",
		UserGroupIDs:          []string{"ug-1"},
	}
}

func mockDatabricksProfileLookups(c *fake.Client) {
	c.FakeNotificationProfiles.On("List").Return([]*api.NotificationProfile{{ProfileID: "notif-789", ProfileName: "Default"}}, nil)
	c.FakeUserGroups.On("List").Return([]*api.UserGroup{{UserGroupID: "ug-1", DisplayName: "Admins"}}, nil)
	c.FakeTags.On("List").Return([]*api.Tag{}, nil)
}

func TestDatabricksMonitorCreate(t *testing.T) {
	d := databricksTestResourceData(t)
	c := fake.NewClient()
	mockDatabricksProfileLookups(c)

	c.FakeDatabricksMonitors.On("Create", expectedDatabricksMonitor("", false)).Return(databricksAPIResponse("123"), nil).Once()
	c.FakeDatabricksMonitors.On("Get", "123").Return(databricksAPIResponse("123"), nil).Once()

	require.NoError(t, databricksMonitorCreate(d, c))
	assert.Equal(t, "123", d.Id())
	// Values the API does not echo must survive the post-create refresh.
	assert.Equal(t, "client-secret", d.Get("secret_key"))
	assert.Equal(t, "client-id", d.Get("access_key"))
	assert.Equal(t, "https://adb-1234567890123456.7.azuredatabricks.net/", d.Get("workspace_uri"))
}

func TestDatabricksMonitorCreateError(t *testing.T) {
	d := databricksTestResourceData(t)
	c := fake.NewClient()
	mockDatabricksProfileLookups(c)

	c.FakeDatabricksMonitors.On("Create", expectedDatabricksMonitor("", false)).Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := databricksMonitorCreate(d, c)
	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.Equal(t, "", d.Id())
}

func TestDatabricksMonitorUpdate(t *testing.T) {
	d := databricksTestResourceData(t)
	d.SetId("123")
	c := fake.NewClient()
	mockDatabricksProfileLookups(c)

	// TestResourceDataRaw has no prior state, so every auth attribute counts as changed.
	expected := expectedDatabricksMonitor("123", true)
	c.FakeDatabricksMonitors.On("Update", expected).Return(databricksAPIResponse("123"), nil).Once()
	c.FakeDatabricksMonitors.On("Get", "123").Return(databricksAPIResponse("123"), nil).Once()
	require.NoError(t, databricksMonitorUpdate(d, c))

	c.FakeDatabricksMonitors.On("Update", expected).Return(nil, apierrors.NewStatusError(500, "error")).Once()
	err := databricksMonitorUpdate(d, c)
	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestDatabricksMonitorRead(t *testing.T) {
	d := databricksTestResourceData(t)
	d.SetId("123")
	c := fake.NewClient()

	resp := databricksAPIResponse("123")
	resp.DisplayName = "Renamed in UI"
	resp.MonitoredServices = []int{5002, 5003}
	resp.DisableDiscovery = databricksIntPtr(1)
	c.FakeDatabricksMonitors.On("Get", "123").Return(resp, nil).Once()

	require.NoError(t, databricksMonitorRead(d, c))
	assert.Equal(t, "123", d.Id())
	assert.Equal(t, "Renamed in UI", d.Get("display_name"))
	assert.Equal(t, 2, d.Get("monitored_services").(*schema.Set).Len())
	assert.Equal(t, true, d.Get("disable_discovery"))
	assert.Equal(t, "client-secret", d.Get("secret_key"), "Read must not overwrite the write-only secret")

	c.FakeDatabricksMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()
	err := databricksMonitorRead(d, c)
	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestDatabricksMonitorReadNotFound(t *testing.T) {
	d := databricksTestResourceData(t)
	d.SetId("123")
	c := fake.NewClient()

	c.FakeDatabricksMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, databricksMonitorRead(d, c))
	assert.Equal(t, "", d.Id())
}

func TestDatabricksMonitorDelete(t *testing.T) {
	d := databricksTestResourceData(t)
	d.SetId("123")
	c := fake.NewClient()

	c.FakeDatabricksMonitors.On("Delete", "123").Return(nil).Once()
	require.NoError(t, databricksMonitorDelete(d, c))

	c.FakeDatabricksMonitors.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()
	require.NoError(t, databricksMonitorDelete(d, c))

	c.FakeDatabricksMonitors.On("Delete", "123").Return(apierrors.NewStatusError(500, "error")).Once()
	assert.Equal(t, apierrors.NewStatusError(500, "error"), databricksMonitorDelete(d, c))
}

func databricksTestResourceData(t *testing.T) *schema.ResourceData {
	// TestResourceDataRaw does not apply schema defaults, so every defaulted
	// attribute is set explicitly.
	return schema.TestResourceDataRaw(t, DatabricksMonitorSchema, map[string]interface{}{
		"display_name":            "Databricks Workspace",
		"workspace_uri":           "https://adb-1234567890123456.7.azuredatabricks.net/",
		"access_key":              "client-id",
		"secret_key":              "client-secret",
		"auth_type":               4,
		"token_profile":           "databricks_oidc",
		"check_frequency":         "30",
		"disable_discovery":       false,
		"threshold_profile_id":    "threshold-456",
		"notification_profile_id": "notif-789",
		"user_group_ids":          []interface{}{"ug-1"},
	})
}
