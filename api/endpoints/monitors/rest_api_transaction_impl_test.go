package monitors

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRestApiTransactionMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create rest api monitor",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitors",
			ExpectedBody: validation.Fixture(t, "requests/create_rest_api_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				restApiTransactionMonitor := &api.RestApiTransactionMonitor{
					DisplayName:               "foo",
					Type:                      string(api.RESTAPISEQ),
					CheckFrequency:            "5",
					LocationProfileID:         "456",
					NotificationProfileID:     "789",
					ThresholdProfileID:        "012",
					Steps: [] api.Steps{
						{
							DisplayName: "Step 1",
							StepsDetails: [] api.StepDetails{
								{
									StepUrl:                   "www.test.tld",
									Timeout:                   10,
									HTTPMethod:                "G",
									HTTPProtocol:              "H1.1",
									SSLProtocol:               "Auto",
									UseAlpn:                   false,
									RequestBody:               "req_param",
									RequestContentType:        "JSON",
									ResponseContentType:       "T",
									OAuth2Provider:            "provider",
									ClientCertificatePassword: "pass",
									JwtID:                     "111",
									UseNameServer:             true,
									MatchCase:                 true,
									UserAgent:                 "firefox",
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

								},
							},
						},
					},
					MonitorGroups:             []string{"234", "567"},
					UserGroupIDs:              []string{"123", "456"},
					TagIDs:                    []string{"123"},
				}

				_, err := NewRestApiTransactionMonitors(c).Create(restApiTransactionMonitor)
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
				restApiTransactionMonitor, err := NewRestApiTransactionMonitors(c).Get("897654345678")
				require.NoError(t, err)

				expected := &api.RestApiTransactionMonitor{
					MonitorID:                 "897654345678",
					DisplayName:               "foo",
					Type:                      string(api.RESTAPISEQ),
					CheckFrequency:            "5",
					Steps: [] api.Steps{
						{
							DisplayName: "Step 1",
							StepsDetails: [] api.StepDetails{
								{
									StepUrl:                   "www.test.tld",
									Timeout:                   10,
									HTTPMethod:                "G",
									HTTPProtocol:              "H1.1",
									SSLProtocol:               "Auto",
									UseAlpn:                   false,
									RequestBody:               "req_param",
									RequestContentType:        "JSON",
									ResponseContentType:       "T",
									OAuth2Provider:            "provider",
									ClientCertificatePassword: "pass",
									JwtID:                     "111",
									UseNameServer:             true,
									MatchCase:                 true,
									JSONSchemaCheck:           false,
									UserAgent:                 "firefox",
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

								},
							},
						},
						{
							DisplayName: "Step 2",
							StepsDetails: [] api.StepDetails{
								{
									StepUrl:                   "www.test.tld",
									Timeout:                   10,
									HTTPMethod:                "G",
									HTTPProtocol:              "H1.1",
									SSLProtocol:               "Auto",
									UseAlpn:                   false,
									RequestBody:               "req_param",
									RequestContentType:        "JSON",
									ResponseContentType:       "T",
									OAuth2Provider:            "provider",
									ClientCertificatePassword: "pass",
									JwtID:                     "111",
									UseNameServer:             true,
									MatchCase:                 true,
									JSONSchemaCheck:           false,
									UserAgent:                 "firefox",
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

								},
							},
						},
					},
					LocationProfileID:         "456",
					NotificationProfileID:     "789",
					ThresholdProfileID:        "012",
					MonitorGroups:             []string{"234", "567"},
					UserGroupIDs:              []string{"123", "456"},
				}

				assert.Equal(t, expected, restApiTransactionMonitor)
			},
		},
		{
			Name:         "list rest api monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_rest_api_monitors.json"),
			Fn: func(t *testing.T, c rest.Client) {
				restApiTransactionMonitor, err := NewRestApiTransactionMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.RestApiTransactionMonitor{
					{
						MonitorID:                 "897654345678",
						DisplayName:               "foo",
						Type:                      string(api.RESTAPISEQ),
						CheckFrequency:            "5",
						LocationProfileID:         "456",
						NotificationProfileID:     "789",
						ThresholdProfileID:        "012",
						Steps: [] api.Steps{
							{
								DisplayName: "Step 1",
								StepsDetails: [] api.StepDetails{
									{
										StepUrl:                   "www.test.tld",
										Timeout:                   10,
										HTTPMethod:                "G",
										HTTPProtocol:              "H1.1",
										SSLProtocol:               "Auto",
										UseAlpn:                   false,
										RequestBody:               "req_param",
										RequestContentType:        "JSON",
										ResponseContentType:       "T",
										OAuth2Provider:            "provider",
										ClientCertificatePassword: "pass",
										JwtID:                     "111",
										UseNameServer:             true,
										MatchCase:                 true,
										JSONSchemaCheck:           false,
										UserAgent:                 "firefox",
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

									},
								},
							},
							{
								DisplayName: "Step 2",
								StepsDetails: [] api.StepDetails{
									{
										StepUrl:                   "www.test.tld",
										Timeout:                   10,
										HTTPMethod:                "G",
										HTTPProtocol:              "H1.1",
										SSLProtocol:               "Auto",
										UseAlpn:                   false,
										RequestBody:               "req_param",
										RequestContentType:        "JSON",
										ResponseContentType:       "T",
										OAuth2Provider:            "provider",
										ClientCertificatePassword: "pass",
										JwtID:                     "111",
										UseNameServer:             true,
										MatchCase:                 true,
										JSONSchemaCheck:           false,
										UserAgent:                 "firefox",
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

									},
								},
							},
						},
						MonitorGroups:             []string{"234", "567"},
						UserGroupIDs:              []string{"123", "456"},
					},
					{
						MonitorID:                 "933654345678",
						DisplayName:               "foo",
						Type:                      string(api.RESTAPISEQ),
						Steps: [] api.Steps{
							{
								DisplayName: "Step 1",
								StepsDetails: [] api.StepDetails{
									{
										StepUrl:                   "www.test.tld",
										Timeout:                   10,
										HTTPMethod:                "G",
										HTTPProtocol:              "H1.1",
										SSLProtocol:               "Auto",
										UseAlpn:                   false,
										RequestBody:               "req_param",
										RequestContentType:        "JSON",
										ResponseContentType:       "T",
										OAuth2Provider:            "provider",
										ClientCertificatePassword: "pass",
										JwtID:                     "111",
										UseNameServer:             true,
										MatchCase:                 true,
										JSONSchemaCheck:           false,
										UserAgent:                 "firefox",
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

									},
								},
							},
							{
								DisplayName: "Step 2",
								StepsDetails: [] api.StepDetails{
									{
										StepUrl:                   "www.test.tld",
										Timeout:                   10,
										HTTPMethod:                "G",
										HTTPProtocol:              "H1.1",
										SSLProtocol:               "Auto",
										UseAlpn:                   false,
										RequestBody:               "req_param",
										RequestContentType:        "JSON",
										ResponseContentType:       "T",
										OAuth2Provider:            "provider",
										ClientCertificatePassword: "pass",
										JwtID:                     "111",
										UseNameServer:             true,
										MatchCase:                 true,
										JSONSchemaCheck:           false,
										UserAgent:                 "firefox",
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

									},
								},
							},
						},
						CheckFrequency:            "5",
						LocationProfileID:         "456",
						NotificationProfileID:     "789",
						ThresholdProfileID:        "012",
						MonitorGroups:             []string{"234", "567"},
						UserGroupIDs:              []string{"123", "456"},
					},
				}

				assert.Equal(t, expected, restApiTransactionMonitor)
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
				restApiTransactionMonitor := &api.RestApiTransactionMonitor{
					MonitorID:                 "123",
					DisplayName:               "foo",
					Type:                      string(api.RESTAPISEQ),
					CheckFrequency:            "5",
					LocationProfileID:         "456",
					NotificationProfileID:     "789",
					ThresholdProfileID:        "012",
					Steps: [] api.Steps{
						{
							DisplayName: "Step 1",
							StepsDetails: [] api.StepDetails{
								{
									StepUrl:                   "www.test.tld",
									Timeout:                   10,
									HTTPMethod:                "G",
									HTTPProtocol:              "H1.1",
									SSLProtocol:               "Auto",
									UseAlpn:                   false,
									RequestBody:               "req_param",
									RequestContentType:        "JSON",
									ResponseContentType:       "T",
									OAuth2Provider:            "provider",
									ClientCertificatePassword: "pass",
									JwtID:                     "111",
									UseNameServer:             true,
									MatchCase:                 true,
									UserAgent:                 "firefox",
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

								},
							},
						},
						{
							DisplayName: "Step 2",
							StepsDetails: [] api.StepDetails{
								{
									StepUrl:                   "www.test.tld",
									Timeout:                   10,
									HTTPMethod:                "G",
									HTTPProtocol:              "H1.1",
									SSLProtocol:               "Auto",
									UseAlpn:                   false,
									RequestBody:               "req_param",
									RequestContentType:        "JSON",
									ResponseContentType:       "T",
									OAuth2Provider:            "provider",
									ClientCertificatePassword: "pass",
									JwtID:                     "111",
									UseNameServer:             true,
									MatchCase:                 true,
									UserAgent:                 "firefox",
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

								},
							},
						},
					},
					MonitorGroups:             []string{"234", "567"},
					UserGroupIDs:              []string{"123", "456"},
					TagIDs:                    []string{"123"},
				}

				_, err := NewRestApiTransactionMonitors(c).Update(restApiTransactionMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete rest api monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewRestApiTransactionMonitors(c).Delete("123"))
			},
		},
	})
}
