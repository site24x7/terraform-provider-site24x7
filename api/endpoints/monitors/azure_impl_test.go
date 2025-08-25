package monitors

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAzureMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create azure monitor",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitors",
			ExpectedBody: validation.Fixture(t, "requests/create_azure_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				azureMonitor := &api.AzureMonitor{
					DisplayName:           "Test Azure Monitor",
					TenantID:              "fake-tenant-id",
					ClientID:              "fake-client-id",
					ClientSecret:          "fake-client-secret",
					Type:                  "AZURE",
					Services:              []string{"Microsoft.Compute/virtualMachines", "Microsoft.Storage/storageAccounts"},
					ManagementGroupReg:    1,
					NotificationProfileID: "123456000000057005",
					UserGroupIDs:          []string{"123456000000057000", "123456000000057001"},
					ThresholdProfileID:    "123456000000057003",
					DiscoveryInterval:     "30",
					AutoAddSubscription:   1,
					AzureIncludeTags: &api.AzureTagCondition{
						Type: 1,
						Tags: map[string][]string{
							"Environment": {"Production"},
						},
					},
					AzureExcludeTags: &api.AzureTagCondition{
						Type: 1,
						Tags: map[string][]string{
							"Environment": {"Development"},
						},
					},
				}

				_, err := NewAzureMonitors(c).Create(azureMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "list azure monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_azure_monitors.json"),
			Fn: func(t *testing.T, c rest.Client) {
				monitors, err := NewAzureMonitors(c).List()
				require.NoError(t, err)
				assert.Len(t, monitors, 2)
			},
		},
		{
			Name:         "update azure monitor",
			ExpectedVerb: "PUT",
			ExpectedPath: "/monitors/azure-monitor-id",
			ExpectedBody: validation.Fixture(t, "requests/update_azure_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				azureMonitor := &api.AzureMonitor{
					MonitorID:             "azure-monitor-id",
					DisplayName:           "Updated Azure Monitor",
					TenantID:              "fake-tenant-id",
					ClientID:              "fake-client-id",
					ClientSecret:          "fake-client-secret",
					Type:                  "AZURE",
					Services:              []string{"Microsoft.Compute/virtualMachines"},
					ManagementGroupReg:    1,
					NotificationProfileID: "123456000000057005",
					UserGroupIDs:          []string{"123456000000057000"},
					ThresholdProfileID:    "123456000000057003",
					DiscoveryInterval:     "60",
					AutoAddSubscription:   1,
					AzureIncludeTags: &api.AzureTagCondition{
						Type: 1,
						Tags: map[string][]string{
							"Team": {"Ops"},
						},
					},
					AzureExcludeTags: &api.AzureTagCondition{
						Type: 1,
						Tags: map[string][]string{
							"Region": {"EastUs"},
						},
					},
				}

				_, err := NewAzureMonitors(c).Update(azureMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete azure monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/azure-monitor-id",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewAzureMonitors(c).Delete("azure-monitor-id"))
			},
		},
	})
}
