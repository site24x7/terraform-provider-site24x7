package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMonitors(t *testing.T) {
	runTests(t, []*endpointTest{
		{
			name:         "create monitor",
			expectedVerb: "POST",
			expectedPath: "/monitors",
			expectedBody: fixture(t, "requests/create_monitor.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
				monitor := &api.WebsiteMonitor{
					Website: "http://www.example.com",
					Type:    "URL",
					CustomHeaders: []api.Header{
						{
							Name:  "Accept-Encoding",
							Value: "gzip",
						},
						{
							Name:  "Cache-Control",
							Value: "nocache",
						},
					},
					UserGroupIDs: []string{
						"123412341234123415",
					},
					LocationProfileID: "123412341234123412",
					UserAgent:         "Mozilla Firefox",
					Timeout:           30,
					MatchRegex: &api.ValueAndSeverity{
						Severity: api.Down,
						Value:    "^reg*",
					},
					AuthUser: "username",
					AuthPass: "password",
					MonitorGroups: []string{
						"123412341234123416",
						"123412341234123417",
					},
					ThresholdProfileID:    "123412341234123414",
					MatchCase:             true,
					NotificationProfileID: "123412341234123413",
					HTTPMethod:            "P",
					MatchingKeyword: &api.ValueAndSeverity{
						Severity: api.Down,
						Value:    "Title",
					},
					ActionIDs: []api.ActionRef{
						{
							ActionID:  "123412341234123418",
							AlertType: 20,
						},
					},
					UnmatchingKeyword: &api.ValueAndSeverity{
						Severity: api.Trouble,
						Value:    "Exception",
					},
					CheckFrequency: "1440",
					DisplayName:    "Display name for the monitor",
					UseNameServer:  true,
				}

				_, err := NewMonitors(c).Create(monitor)
				require.NoError(t, err)
			},
		},
		{
			name:         "create monitor error",
			statusCode:   500,
			responseBody: []byte("whoops"),
			fn: func(t *testing.T, c rest.Client) {
				_, err := NewMonitors(c).Create(&api.WebsiteMonitor{})
				assert.True(t, apierrors.HasStatusCode(err, 500))
			},
		},
		{
			name:         "get monitor",
			expectedVerb: "GET",
			expectedPath: "/monitors/123412341234123411",
			statusCode:   200,
			responseBody: fixture(t, "responses/get_monitor.json"),
			fn: func(t *testing.T, c rest.Client) {
				monitor, err := NewMonitors(c).Get("123412341234123411")
				require.NoError(t, err)

				expected := &api.WebsiteMonitor{
					MonitorID: "123412341234123411",
					Website:   "http://www.example.com",
					Type:      "URL",
					CustomHeaders: []api.Header{
						{
							Name:  "Accept-Encoding",
							Value: "gzip",
						},
						{
							Name:  "Cache-Control",
							Value: "nocache",
						},
					},
					UserGroupIDs: []string{
						"123412341234123415",
					},
					LocationProfileID: "123412341234123412",
					UserAgent:         "Mozilla Firefox",
					Timeout:           30,
					MatchRegex: &api.ValueAndSeverity{
						Severity: api.Down,
						Value:    "^reg*",
					},
					AuthUser: "username",
					AuthPass: "password",
					MonitorGroups: []string{
						"123412341234123416",
						"123412341234123417",
					},
					ThresholdProfileID:    "123412341234123414",
					MatchCase:             true,
					NotificationProfileID: "123412341234123413",
					HTTPMethod:            "P",
					MatchingKeyword: &api.ValueAndSeverity{
						Severity: api.Down,
						Value:    "Title",
					},
					ActionIDs: []api.ActionRef{
						{
							ActionID:  "123412341234123418",
							AlertType: 20,
						},
					},
					UnmatchingKeyword: &api.ValueAndSeverity{
						Severity: api.Trouble,
						Value:    "Exception",
					},
					CheckFrequency: "1440",
					DisplayName:    "Display name for the monitor",
					UseNameServer:  true,
					UpStatusCodes:  "200",
				}

				assert.Equal(t, expected, monitor)
			},
		},
		{
			name:         "list monitors",
			expectedVerb: "GET",
			expectedPath: "/monitors",
			statusCode:   200,
			responseBody: fixture(t, "responses/list_monitors.json"),
			fn: func(t *testing.T, c rest.Client) {
				monitor, err := NewMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.WebsiteMonitor{
					{
						MonitorID: "12340000016033021",
						Website:   "https://foo.bar/",
						Type:      "URL",
						CustomHeaders: []api.Header{
							{
								Name:  "Accept-Encoding",
								Value: "gzip",
							},
							{
								Name:  "Cache-Control",
								Value: "nocache",
							},
						},
						UserGroupIDs: []string{
							"12340000000018013",
						},
						LocationProfileID: "12340000001806001",
						Timeout:           10,
						MonitorGroups: []string{
							"12340000005749001",
						},
						ThresholdProfileID:    "12340000001812001",
						NotificationProfileID: "12340000003579003",
						HTTPMethod:            "G",
						ActionIDs: []api.ActionRef{
							{
								ActionID:  "12340000019175145",
								AlertType: 0,
							},
							{
								ActionID:  "12340000019181133",
								AlertType: 1,
							},
						},
						CheckFrequency: "1",
						DisplayName:    "foo.bar",
						UseNameServer:  true,
					},
					{
						MonitorID: "12340000016108026",
						Website:   "https://some.api.tld/api/v1/status",
						Type:      "URL",
						UserGroupIDs: []string{
							"12340000015652005",
						},
						LocationProfileID: "12340000001806001",
						Timeout:           30,
						MonitorGroups: []string{
							"12340000002807001",
						},
						ThresholdProfileID:    "12340000001812001",
						NotificationProfileID: "12340000003579003",
						HTTPMethod:            "P",
						ActionIDs: []api.ActionRef{
							{
								ActionID:  "12340000019180203",
								AlertType: 0,
							},
							{
								ActionID:  "12340000019181125",
								AlertType: 1,
							},
						},
						CheckFrequency: "5",
						DisplayName:    "some.api.tld",
						AuthUser:       "username",
						AuthPass:       "password",
						UseNameServer:  true,
					},
				}

				assert.Equal(t, expected, monitor)
			},
		},
		{
			name:         "update monitor",
			expectedVerb: "PUT",
			expectedPath: "/monitors/456",
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, map[string]interface{}{
				"monitor_id":   "456",
				"display_name": "bar",
			}),
			fn: func(t *testing.T, c rest.Client) {
				monitor := &api.WebsiteMonitor{MonitorID: "456", DisplayName: "bar"}

				monitor, err := NewMonitors(c).Update(monitor)
				require.NoError(t, err)

				expected := &api.WebsiteMonitor{
					MonitorID:   "456",
					DisplayName: "bar",
				}

				assert.Equal(t, expected, monitor)
			},
		},
		{
			name:       "update monitor error",
			statusCode: 400,
			responseBody: jsonBody(t, &api.ErrorResponse{
				ErrorCode: 123,
				Message:   "bad request",
				ErrorInfo: map[string]interface{}{"foo": "bar"},
			}),
			fn: func(t *testing.T, c rest.Client) {
				_, err := NewMonitors(c).Update(&api.WebsiteMonitor{})
				assert.True(t, apierrors.HasStatusCode(err, 400))
			},
		},
		{
			name:         "delete monitor",
			expectedVerb: "DELETE",
			expectedPath: "/monitors/123",
			statusCode:   200,
			fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewMonitors(c).Delete("123"))
			},
		},
		{
			name:       "delete monitor not found",
			statusCode: 404,
			fn: func(t *testing.T, c rest.Client) {
				err := NewMonitors(c).Delete("123")
				assert.True(t, apierrors.IsNotFound(err))
			},
		},
	})
}
