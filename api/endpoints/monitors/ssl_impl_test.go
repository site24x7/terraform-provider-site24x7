package monitors

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSSLMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create ssl monitor",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitors",
			ExpectedBody: validation.Fixture(t, "api/endpoints/testdata/fixtures/requests/create_ssl_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				sslMonitor := &api.SSLMonitor{
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
					TagIDs:                []string{"123"},
				}

				_, err := NewSSLMonitors(c).Create(sslMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get ssl monitor",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors/897654345678",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "api/endpoints/testdata/fixtures/responses/get_ssl_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				sslMonitor, err := NewSSLMonitors(c).Get("897654345678")
				require.NoError(t, err)

				expected := &api.SSLMonitor{
					MonitorID:             "897654345678",
					DisplayName:           "foo",
					DomainName:            "www.example.com",
					Type:                  "SSL_CERT",
					Timeout:               30,
					Protocol:              "HTTPS",
					ExpireDays:            30,
					HTTPProtocolVersion:   "H1.1",
					LocationProfileID:     "456",
					NotificationProfileID: "789",
					ThresholdProfileID:    "012",
					MonitorGroups:         []string{"234", "567"},
					UserGroupIDs:          []string{"123", "456"},
				}

				assert.Equal(t, expected, sslMonitor)
			},
		},
		{
			Name:         "list ssl monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "api/endpoints/testdata/fixtures/responses/list_ssl_monitors.json"),
			Fn: func(t *testing.T, c rest.Client) {
				sslMonitors, err := NewSSLMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.SSLMonitor{
					{
						MonitorID:   "897654345678",
						DisplayName: "foo",
						DomainName:  "www.example.com",
						Type:        "SSL_CERT",
						Timeout:     30,
						Protocol:    "HTTPS",

						ExpireDays:            30,
						HTTPProtocolVersion:   "H1.1",
						LocationProfileID:     "456",
						NotificationProfileID: "789",
						ThresholdProfileID:    "012",
						MonitorGroups:         []string{"234", "567"},
						UserGroupIDs:          []string{"123", "456"},
					},
					{
						MonitorID:   "933654345678",
						DisplayName: "foo",
						DomainName:  "www.example.com",
						Type:        "SSL_CERT",
						Timeout:     30,
						Protocol:    "HTTPS",

						ExpireDays:            30,
						HTTPProtocolVersion:   "H1.1",
						LocationProfileID:     "456",
						NotificationProfileID: "789",
						ThresholdProfileID:    "012",
						MonitorGroups:         []string{"234", "567"},
						UserGroupIDs:          []string{"123", "456"},
					},
				}

				assert.Equal(t, expected, sslMonitors)
			},
		},
		{
			Name:         "update ssl monitor",
			ExpectedVerb: "PUT",
			ExpectedPath: "/monitors/123",
			ExpectedBody: validation.Fixture(t, "api/endpoints/testdata/fixtures/requests/update_ssl_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				sslMonitor := &api.SSLMonitor{
					MonitorID:             "123",
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
					TagIDs:                []string{"123"},
				}

				_, err := NewSSLMonitors(c).Update(sslMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete ssl monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewSSLMonitors(c).Delete("123"))
			},
		},
	})
}
