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

func TestAzureMonitorCreate(t *testing.T) {
	d := azureTestResourceData(t)
	c := fake.NewClient()

	a := &api.AzureMonitor{
		DisplayName:           "Test Azure Monitor",
		TenantID:              "tenant-123",
		ClientID:              "client-abc",
		ClientSecret:          "secret-key",
		Type:                  "AZURE",
		Services:              []string{"vm", "sql"},
		ManagementGroupReg:    0,
		NotificationProfileID: "notif-789",
		UserGroupIDs:          []string{"user1", "user2"},
		ThresholdProfileID:    "threshold-456",
		DiscoveryInterval:     "30",
		AutoAddSubscription:   1,
	}

	c.FakeAzureMonitors.On("Create", a).Return(a, nil).Once()

	require.NoError(t, azureMonitorCreate(d, c))

	c.FakeAzureMonitors.On("Create", a).Return(nil, apierrors.NewStatusError(500, "error")).Once()
	err := azureMonitorCreate(d, c)
	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestAzureMonitorUpdate(t *testing.T) {
	d := azureTestResourceData(t)
	d.SetId("monitor-123")

	c := fake.NewClient()
	a := &api.AzureMonitor{
		MonitorID:             "monitor-123",
		DisplayName:           "Test Azure Monitor",
		TenantID:              "tenant-123",
		ClientID:              "client-abc",
		ClientSecret:          "secret-key",
		Type:                  "AZURE",
		Services:              []string{"vm", "sql"},
		ManagementGroupReg:    0,
		NotificationProfileID: "notif-789",
		UserGroupIDs:          []string{"user1", "user2"},
		ThresholdProfileID:    "threshold-456",
		DiscoveryInterval:     "30",
		AutoAddSubscription:   1,
	}

	c.FakeAzureMonitors.On("Update", a).Return(a, nil).Once()
	require.NoError(t, azureMonitorUpdate(d, c))

	c.FakeAzureMonitors.On("Update", a).Return(nil, apierrors.NewStatusError(500, "error")).Once()
	err := azureMonitorUpdate(d, c)
	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestAzureMonitorRead(t *testing.T) {
	d := azureTestResourceData(t)
	d.SetId("monitor-123")

	c := fake.NewClient()
	c.FakeAzureMonitors.On("Get", "monitor-123").Return(&api.AzureMonitor{}, nil).Once()
	require.NoError(t, azureMonitorRead(d, c))

	c.FakeAzureMonitors.On("Get", "monitor-123").Return(nil, apierrors.NewStatusError(500, "error")).Once()
	err := azureMonitorRead(d, c)
	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestAzureMonitorDelete(t *testing.T) {
	d := azureTestResourceData(t)
	d.SetId("monitor-123")

	c := fake.NewClient()
	c.FakeAzureMonitors.On("Delete", "monitor-123").Return(nil).Once()
	require.NoError(t, azureMonitorDelete(d, c))

	c.FakeAzureMonitors.On("Delete", "monitor-123").Return(apierrors.NewStatusError(404, "not found")).Once()
	require.NoError(t, azureMonitorDelete(d, c))
}

func TestAzureMonitorExists(t *testing.T) {
	d := azureTestResourceData(t)
	d.SetId("monitor-123")

	c := fake.NewClient()
	c.FakeAzureMonitors.On("Get", "monitor-123").Return(&api.AzureMonitor{}, nil).Once()

	exists, err := azureMonitorExists(d, c)
	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeAzureMonitors.On("Get", "monitor-123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()
	exists, err = azureMonitorExists(d, c)
	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeAzureMonitors.On("Get", "monitor-123").Return(nil, apierrors.NewStatusError(500, "error")).Once()
	exists, err = azureMonitorExists(d, c)
	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func azureTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, AzureMonitorSchema, map[string]interface{}{
		"display_name":            "Test Azure Monitor",
		"tenant_id":               "tenant-123",
		"client_id":               "client-abc",
		"client_secret":           "secret-key",
		"type":                    "AZURE",
		"services":                []interface{}{"vm", "sql"},
		"management_group_reg":    0,
		"notification_profile_id": "notif-789",
		"user_group_ids":          []interface{}{"user1", "user2"},
		"threshold_profile_id":    "threshold-456",
		"discovery_interval":      "30",
		"auto_add_subscription":   1,
	})
}
