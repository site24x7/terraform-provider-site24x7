package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSSLMonitors(t *testing.T) {
	runTests(t, []*endpointTest{
		{
			name:         "create ssl monitor",
			expectedVerb: "POST",
			expectedPath: "/monitors",
			expectedBody: fixture(t, "requests/create_ssl_monitor.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
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
			name:         "get ssl monitor",
			expectedVerb: "GET",
			expectedPath: "/monitors/897654345678",
			statusCode:   200,
			responseBody: fixture(t, "responses/get_ssl_monitor.json"),
			fn: func(t *testing.T, c rest.Client) {
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
			name:         "list ssl monitors",
			expectedVerb: "GET",
			expectedPath: "/monitors",
			statusCode:   200,
			responseBody: fixture(t, "responses/list_ssl_monitors.json"),
			fn: func(t *testing.T, c rest.Client) {
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
			name:         "update ssl monitor",
			expectedVerb: "PUT",
			expectedPath: "/monitors/123",
			expectedBody: fixture(t, "requests/update_ssl_monitor.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
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
			name:         "delete ssl monitor",
			expectedVerb: "DELETE",
			expectedPath: "/monitors/123",
			statusCode:   200,
			fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewSSLMonitors(c).Delete("123"))
			},
		},
	})
}
