package monitors

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func intPtr(i int) *int    { return &i }
func boolPtr(b bool) *bool { return &b }

func TestDatabricksMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create databricks monitor",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitors",
			ExpectedBody: validation.Fixture(t, "requests/create_databricks_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_databricks_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				monitor := &api.DatabricksMonitor{
					DisplayName:           "Databricks Workspace",
					Type:                  "DBRICKS_WS",
					CheckFrequency:        "30",
					AuthType:              intPtr(4),
					TokenProfile:          "databricks_oidc",
					WorkspaceURI:          "https://adb-1234567890123456.7.azuredatabricks.net/",
					AccessKey:             "fake-client-id",
					SecretKey:             "fake-client-secret",
					MonitoredServices:     []int{5002},
					DisableDiscovery:      intPtr(0),
					ThresholdProfileID:    "123456000000057003",
					NotificationProfileID: "123456000000057005",
					UserGroupIDs:          []string{"123456000000057000"},
					TagIDs:                []string{},
					UpdateAuthDetails:     boolPtr(false),
				}

				created, err := NewDatabricksMonitors(c).Create(monitor)
				require.NoError(t, err)
				assert.Equal(t, "123456000028010003", created.MonitorID)
			},
		},
		{
			Name:         "get databricks monitor",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors/123456000028010003",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_databricks_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				monitor, err := NewDatabricksMonitors(c).Get("123456000028010003")
				require.NoError(t, err)

				expected := &api.DatabricksMonitor{
					MonitorID:             "123456000028010003",
					DisplayName:           "Databricks Workspace",
					Type:                  "DBRICKS_WS",
					CheckFrequency:        "30",
					AuthType:              intPtr(4),
					WorkspaceURI:          "https://adb-1234567890123456.7.azuredatabricks.net/",
					MonitoredServices:     api.IntList{5002},
					DisableDiscovery:      intPtr(0),
					ThresholdProfileID:    "123456000000057003",
					NotificationProfileID: "123456000000057005",
					UserGroupIDs:          []string{"123456000000057000"},
				}
				assert.Equal(t, expected, monitor)
			},
		},
		{
			Name:         "list databricks monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_databricks_monitors.json"),
			Fn: func(t *testing.T, c rest.Client) {
				monitors, err := NewDatabricksMonitors(c).List()
				require.NoError(t, err)
				assert.Len(t, monitors, 2)
			},
		},
		{
			Name:         "update databricks monitor",
			ExpectedVerb: "PUT",
			ExpectedPath: "/monitors/123456000028010003",
			ExpectedBody: validation.Fixture(t, "requests/update_databricks_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				monitor := &api.DatabricksMonitor{
					MonitorID:             "123456000028010003",
					DisplayName:           "Databricks Workspace Renamed",
					Type:                  "DBRICKS_WS",
					CheckFrequency:        "60",
					AuthType:              intPtr(4),
					TokenProfile:          "databricks_oidc",
					WorkspaceURI:          "https://adb-1234567890123456.7.azuredatabricks.net/",
					AccessKey:             "fake-client-id",
					SecretKey:             "rotated-client-secret",
					MonitoredServices:     []int{5002},
					DisableDiscovery:      intPtr(1),
					ThresholdProfileID:    "123456000000057003",
					NotificationProfileID: "123456000000057005",
					UserGroupIDs:          []string{"123456000000057000"},
					TagIDs:                []string{"123456000000057010"},
					UpdateAuthDetails:     boolPtr(true),
				}

				_, err := NewDatabricksMonitors(c).Update(monitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete databricks monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/123456000028010003",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewDatabricksMonitors(c).Delete("123456000028010003"))
			},
		},
	})
}
