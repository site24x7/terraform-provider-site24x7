package monitors

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGCPMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create gcp monitor",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitors",
			ExpectedBody: validation.Fixture(t, "requests/create_gcp_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				gcpMonitor := &api.GCPMonitor{
					DisplayName:          "GCP Monitor Display Name",
					Type:                 "GCP",
					ProjectID:            "project-id",
					DiscoverServices:     []int{1},
					CheckFrequency:       "5",
					StopRediscoverOption: 1,
					GCPSAContent: struct {
						PrivateKey  string `json:"private_key"`
						ClientEmail string `json:"client_email"`
					}{
						PrivateKey:  "private-key",
						ClientEmail: "client-email",
					},
					NotificationProfileID: "123412341234123413",
					UserGroupIDs: []string{
						"123412341234123415",
					},
				}

				_, err := NewGCPMonitors(c).Create(gcpMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get gcp monitor",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors/113770000041271035",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_gcp_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				gcpMonitor, err := NewGCPMonitors(c).Get("113770000041271035")
				require.NoError(t, err)

				expected := &api.GCPMonitor{
					DisplayName:          "GCP Monitor Display Name",
					Type:                 "GCP",
					ProjectID:            "project-id",
					DiscoverServices:     []int{1},
					CheckFrequency:       "5",
					StopRediscoverOption: 1,
					GCPSAContent: struct {
						PrivateKey  string `json:"private_key"`
						ClientEmail string `json:"client_email"`
					}{
						PrivateKey:  "private-key",
						ClientEmail: "client-email",
					},
					NotificationProfileID: "123412341234123413",
					UserGroupIDs: []string{
						"123412341234123415",
					},
				}

				assert.Equal(t, expected, gcpMonitor)
			},
		},
		{
			Name:         "list gcp monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_gcp_monitors.json"),
			Fn: func(t *testing.T, c rest.Client) {
				gcpMonitors, err := NewGCPMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.GCPMonitor{
					{
						MonitorID:            "123447000020646003",
						DisplayName:          "GCP Monitor Display Name",
						Type:                 "GCP",
						ProjectID:            "project-id",
						DiscoverServices:     []int{1},
						CheckFrequency:       "5",
						StopRediscoverOption: 1,
						GCPSAContent: struct {
							PrivateKey  string `json:"private_key"`
							ClientEmail string `json:"client_email"`
						}{
							PrivateKey:  "private-key",
							ClientEmail: "client-email",
						},
						NotificationProfileID: "123412341234123413",
						UserGroupIDs: []string{
							"123412341234123415",
						},
					},
					{
						MonitorID:            "123447000020646008",
						DisplayName:          "GCP Monitor Display Name",
						Type:                 "GCP",
						ProjectID:            "project-id",
						DiscoverServices:     []int{1},
						CheckFrequency:       "5",
						StopRediscoverOption: 1,
						GCPSAContent: struct {
							PrivateKey  string `json:"private_key"`
							ClientEmail string `json:"client_email"`
						}{
							PrivateKey:  "private-key",
							ClientEmail: "client-email",
						},
						NotificationProfileID: "123412341234123413",
						UserGroupIDs: []string{
							"123412341234123415",
						},
					},
				}

				assert.Equal(t, expected, gcpMonitors)
			},
		},
		{
			Name:         "update gcp monitor",
			ExpectedVerb: "PUT",
			ExpectedPath: "/monitors/123",
			ExpectedBody: validation.Fixture(t, "requests/update_gcp_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				gcpMonitor := &api.GCPMonitor{
					MonitorID:            "123",
					DisplayName:          "foo",
					Type:                 "GCP",
					ProjectID:            "project-id",
					DiscoverServices:     []int{1},
					CheckFrequency:       "5",
					StopRediscoverOption: 1,
					GCPSAContent: struct {
						PrivateKey  string `json:"private_key"`
						ClientEmail string `json:"client_email"`
					}{
						PrivateKey:  "private-key",
						ClientEmail: "client-email",
					},
					NotificationProfileID: "123412341234123413",
					UserGroupIDs: []string{
						"123412341234123415",
					},
				}

				_, err := NewGCPMonitors(c).Update(gcpMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete gcp monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewGCPMonitors(c).Delete("123"))
			},
		},
	})
}
