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
					DisplayName:           "Azure Monitor Display Name",
					TenantID:              "tenant-id",
					ClientID:              "client-id",
					ClientSecret:          "client-secret",
					Type:                  "AZURE",
					Services:              []string{"VirtualMachines"},
					ManagementGroupReg:    0,
					NotificationProfileID: "np-id",
					UserGroupIDs:          []string{"ug-id"},
					ThresholdProfileID:    "tp-id",
					DiscoveryInterval:     "60",
					AutoAddSubscription:   1,
				}

				_, err := NewAzureMonitors(c).Create(azureMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get azure monitor",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors/azure-monitor-id",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_azure_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				azureMonitor, err := NewAzureMonitors(c).Get("azure-monitor-id")
				require.NoError(t, err)

				expected := &api.AzureMonitor{
					DisplayName:           "Azure Monitor Display Name",
					TenantID:              "tenant-id",
					ClientID:              "client-id",
					ClientSecret:          "client-secret",
					Type:                  "AZURE",
					Services:              []string{"VirtualMachines"},
					ManagementGroupReg:    0,
					NotificationProfileID: "np-id",
					UserGroupIDs:          []string{"ug-id"},
					ThresholdProfileID:    "tp-id",
					DiscoveryInterval:     "60",
					AutoAddSubscription:   1,
				}

				assert.Equal(t, expected, azureMonitor)
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
					TenantID:              "tenant-id",
					ClientID:              "client-id",
					ClientSecret:          "client-secret",
					Type:                  "AZURE",
					Services:              []string{"VirtualMachines"},
					ManagementGroupReg:    1,
					NotificationProfileID: "np-id",
					UserGroupIDs:          []string{"ug-id"},
					ThresholdProfileID:    "tp-id",
					DiscoveryInterval:     "360",
					AutoAddSubscription:   0,
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
