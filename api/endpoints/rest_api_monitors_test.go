package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRestApiMonitors(t *testing.T) {
	runTests(t, []*endpointTest{
		{
			name:         "create rest api monitor",
			expectedVerb: "POST",
			expectedPath: "/monitors",
			expectedBody: fixture(t, "requests/create_rest_api_monitor.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
				restApiMonitor := &api.RestApiMonitor{
					DisplayName:               "foo",
					Type:                      string(api.RESTAPI),
					Website:                   "www.test.tld",
					CheckFrequency:            "5",
					Timeout:                   10,
					HttpMethod:                "G",
					HttpProtocol:              "H1.1",
					SslProtocol:               "Auto",
					UseAlpn:                   false,
					UseIPV6:                   false,
					RequestParam:              "req_param",
					RequestContentType:        "JSON",
					ResponseContentType:       "T",
					OAuth2Provider:            "provider",
					ClientCertificatePassword: "pass",
					JwtID:                     "111",
					LocationProfileID:         "456",
					NotificationProfileID:     "789",
					ThresholdProfileID:        "012",
					UseNameServer:             true,
					MatchCase:                 true,
					JSONSchemaCheck:           false,
					UserAgent:                 "firefox",
					MonitorGroups:             []string{"234", "567"},
					UserGroupIDs:              []string{"123", "456"},
					AuthUser:                  "username",
					AuthPass:                  "password",
					ResponseHeaders: api.HTTPResponseHeader{
						Severity: api.Trouble,
						Value: []api.Header{
							{
								Name:  "Accept-Encoding",
								Value: "gzip",
							},
							{
								Name:  "Cache-Control",
								Value: "nocache",
							},
						},
					},
				}

				_, err := NewRestApiMonitors(c).Create(restApiMonitor)
				require.NoError(t, err)
			},
		},
		{
			name:         "get rest api monitor",
			expectedVerb: "GET",
			expectedPath: "/monitors/897654345678",
			statusCode:   200,
			responseBody: fixture(t, "responses/get_rest_api_monitor.json"),
			fn: func(t *testing.T, c rest.Client) {
				restApiMonitor, err := NewRestApiMonitors(c).Get("897654345678")
				require.NoError(t, err)

				expected := &api.RestApiMonitor{
					MonitorID:                 "897654345678",
					DisplayName:               "foo",
					Type:                      string(api.RESTAPI),
					Website:                   "www.test.tld",
					CheckFrequency:            "5",
					Timeout:                   10,
					HttpMethod:                "G",
					HttpProtocol:              "H1.1",
					SslProtocol:               "Auto",
					UseAlpn:                   false,
					UseIPV6:                   false,
					RequestParam:              "req_param",
					RequestContentType:        "JSON",
					ResponseContentType:       "T",
					OAuth2Provider:            "provider",
					ClientCertificatePassword: "pass",
					JwtID:                     "111",
					LocationProfileID:         "456",
					NotificationProfileID:     "789",
					ThresholdProfileID:        "012",
					UseNameServer:             true,
					MatchCase:                 true,
					JSONSchemaCheck:           false,
					UserAgent:                 "firefox",
					MonitorGroups:             []string{"234", "567"},
					UserGroupIDs:              []string{"123", "456"},
					AuthUser:                  "username",
					AuthPass:                  "password",
				}

				assert.Equal(t, expected, restApiMonitor)
			},
		},
		{
			name:         "list rest api monitors",
			expectedVerb: "GET",
			expectedPath: "/monitors",
			statusCode:   200,
			responseBody: fixture(t, "responses/list_rest_api_monitors.json"),
			fn: func(t *testing.T, c rest.Client) {
				restApiMonitors, err := NewRestApiMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.RestApiMonitor{
					{
						MonitorID:                 "897654345678",
						DisplayName:               "foo",
						Type:                      string(api.RESTAPI),
						Website:                   "www.test.tld",
						CheckFrequency:            "5",
						Timeout:                   10,
						HttpMethod:                "G",
						HttpProtocol:              "H1.1",
						SslProtocol:               "Auto",
						UseAlpn:                   false,
						UseIPV6:                   false,
						RequestParam:              "req_param",
						RequestContentType:        "JSON",
						ResponseContentType:       "T",
						OAuth2Provider:            "provider",
						ClientCertificatePassword: "pass",
						JwtID:                     "111",
						LocationProfileID:         "456",
						NotificationProfileID:     "789",
						ThresholdProfileID:        "012",
						UseNameServer:             true,
						MatchCase:                 true,
						JSONSchemaCheck:           false,
						UserAgent:                 "firefox",
						MonitorGroups:             []string{"234", "567"},
						UserGroupIDs:              []string{"123", "456"},
						AuthUser:                  "username",
						AuthPass:                  "password",
					},
					{
						MonitorID:                 "933654345678",
						DisplayName:               "foo",
						Type:                      string(api.RESTAPI),
						Website:                   "www.test.tld",
						CheckFrequency:            "5",
						Timeout:                   10,
						HttpMethod:                "G",
						HttpProtocol:              "H1.1",
						SslProtocol:               "Auto",
						UseAlpn:                   false,
						UseIPV6:                   false,
						RequestParam:              "req_param",
						RequestContentType:        "JSON",
						ResponseContentType:       "T",
						OAuth2Provider:            "provider",
						ClientCertificatePassword: "pass",
						JwtID:                     "111",
						LocationProfileID:         "456",
						NotificationProfileID:     "789",
						ThresholdProfileID:        "012",
						UseNameServer:             true,
						MatchCase:                 true,
						JSONSchemaCheck:           false,
						UserAgent:                 "firefox",
						MonitorGroups:             []string{"234", "567"},
						UserGroupIDs:              []string{"123", "456"},
						AuthUser:                  "username",
						AuthPass:                  "password",
					},
				}

				assert.Equal(t, expected, restApiMonitors)
			},
		},
		{
			name:         "update rest api monitor",
			expectedVerb: "PUT",
			expectedPath: "/monitors/123",
			expectedBody: fixture(t, "requests/update_rest_api_monitor.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
				restApiMonitor := &api.RestApiMonitor{
					MonitorID:                 "123",
					DisplayName:               "foo",
					Type:                      string(api.RESTAPI),
					Website:                   "www.test.tld",
					CheckFrequency:            "5",
					Timeout:                   10,
					HttpMethod:                "G",
					HttpProtocol:              "H1.1",
					SslProtocol:               "Auto",
					UseAlpn:                   false,
					UseIPV6:                   false,
					RequestParam:              "req_param",
					RequestContentType:        "JSON",
					ResponseContentType:       "T",
					OAuth2Provider:            "provider",
					ClientCertificatePassword: "pass",
					JwtID:                     "111",
					LocationProfileID:         "456",
					NotificationProfileID:     "789",
					ThresholdProfileID:        "012",
					UseNameServer:             true,
					MatchCase:                 true,
					JSONSchemaCheck:           false,
					UserAgent:                 "firefox",
					MonitorGroups:             []string{"234", "567"},
					UserGroupIDs:              []string{"123", "456"},
					AuthUser:                  "username",
					AuthPass:                  "password",
					ResponseHeaders: api.HTTPResponseHeader{
						Severity: api.Trouble,
						Value: []api.Header{
							{
								Name:  "Accept-Encoding",
								Value: "gzip",
							},
							{
								Name:  "Cache-Control",
								Value: "nocache",
							},
						},
					},
				}

				_, err := NewRestApiMonitors(c).Update(restApiMonitor)
				require.NoError(t, err)
			},
		},
		{
			name:         "delete rest api monitor",
			expectedVerb: "DELETE",
			expectedPath: "/monitors/123",
			statusCode:   200,
			fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewRestApiMonitors(c).Delete("123"))
			},
		},
	})
}
