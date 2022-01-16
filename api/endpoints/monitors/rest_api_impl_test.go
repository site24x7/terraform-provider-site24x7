package monitors

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRestApiMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create rest api monitor",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitors",
			ExpectedBody: validation.Fixture(t, "requests/create_rest_api_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
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
					TagIDs:                    []string{"123"},
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
			Name:         "get rest api monitor",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors/897654345678",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_rest_api_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
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
			Name:         "list rest api monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_rest_api_monitors.json"),
			Fn: func(t *testing.T, c rest.Client) {
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
			Name:         "update rest api monitor",
			ExpectedVerb: "PUT",
			ExpectedPath: "/monitors/123",
			ExpectedBody: validation.Fixture(t, "requests/update_rest_api_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
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
					TagIDs:                    []string{"123"},
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
			Name:         "delete rest api monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewRestApiMonitors(c).Delete("123"))
			},
		},
	})
}
