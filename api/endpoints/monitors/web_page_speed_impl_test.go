package monitors

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebPageSpeedMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create monitor",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitors",
			ExpectedBody: validation.Fixture(t, "requests/create_web_page_speed_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				monitor := &api.WebPageSpeedMonitor{
					Website:        "http://www.example.com",
					Type:           "HOMEPAGE",
					BrowserType:    1,
					BrowserVersion: 83,
					WebsiteType:    1,
					DeviceType:     "1",
					WPAResolution:  "1024,768",
					IPType:         2,
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
					ThresholdProfileID:    "123412341234123414",
					LocationProfileID:     "123412341234123412",
					NotificationProfileID: "123412341234123413",
					UserGroupIDs: []string{
						"123412341234123415",
					},
					ThirdPartyServiceIDs: []string{
						"456987654321012",
						"456987654321013",
					},
					UserAgent: "Mozilla Firefox",
					Timeout:   30,
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
					MatchCase: true,

					HTTPMethod: "P",
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
				}

				_, err := NewWebPageSpeedMonitors(c).Create(monitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "create monitor error",
			StatusCode:   500,
			ResponseBody: []byte("whoops"),
			Fn: func(t *testing.T, c rest.Client) {
				_, err := NewWebPageSpeedMonitors(c).Create(&api.WebPageSpeedMonitor{})
				assert.True(t, apierrors.HasStatusCode(err, 500))
			},
		},
		{
			Name:         "get monitor",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors/123412341234123411",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_web_page_speed_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				monitor, err := NewWebPageSpeedMonitors(c).Get("123412341234123411")
				require.NoError(t, err)

				expected := &api.WebPageSpeedMonitor{
					MonitorID:      "123412341234123411",
					Website:        "http://www.example.com",
					Type:           "HOMEPAGE",
					BrowserType:    1,
					BrowserVersion: 83,
					WebsiteType:    1,
					DeviceType:     "1",
					WPAResolution:  "1024,768",
					IPType:         2,
					UpStatusCodes:  "200",
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
					TagIDs: []string{
						"123456987654321012",
						"123456987654321013",
					},
					ThirdPartyServiceIDs: []string{
						"456987654321012",
						"456987654321013",
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
				}

				assert.Equal(t, expected, monitor)
			},
		},
		{
			Name:         "list monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_web_page_speed_monitors.json"),
			Fn: func(t *testing.T, c rest.Client) {
				monitor, err := NewWebPageSpeedMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.WebPageSpeedMonitor{
					{
						MonitorID:      "12340000016033021",
						Website:        "https://foo.bar/",
						Type:           "HOMEPAGE",
						BrowserType:    1,
						IPType:         2,
						BrowserVersion: 83,
						WebsiteType:    1,
						DeviceType:     "1",
						WPAResolution:  "1024,768",
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
					},
					{
						MonitorID:      "12340000016108026",
						Website:        "https://some.api.tld/api/v1/status",
						Type:           "HOMEPAGE",
						BrowserType:    1,
						IPType:         2,
						BrowserVersion: 83,
						WebsiteType:    1,
						DeviceType:     "1",
						WPAResolution:  "1024,768",
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
					},
				}

				assert.Equal(t, expected, monitor)
			},
		},
		{
			Name:         "update monitor",
			ExpectedVerb: "PUT",
			ExpectedPath: "/monitors/456",
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, map[string]interface{}{
				"monitor_id":   "456",
				"display_name": "bar",
			}),
			Fn: func(t *testing.T, c rest.Client) {
				monitor := &api.WebPageSpeedMonitor{MonitorID: "456", DisplayName: "bar"}

				monitor, err := NewWebPageSpeedMonitors(c).Update(monitor)
				require.NoError(t, err)

				expected := &api.WebPageSpeedMonitor{
					MonitorID:   "456",
					DisplayName: "bar",
				}

				assert.Equal(t, expected, monitor)
			},
		},
		{
			Name:       "update monitor error",
			StatusCode: 400,
			ResponseBody: validation.JsonBody(t, &api.ErrorResponse{
				ErrorCode: 123,
				Message:   "bad request",
				ErrorInfo: map[string]interface{}{"foo": "bar"},
			}),
			Fn: func(t *testing.T, c rest.Client) {
				_, err := NewWebPageSpeedMonitors(c).Update(&api.WebPageSpeedMonitor{})
				assert.True(t, apierrors.HasStatusCode(err, 400))
			},
		},
		{
			Name:         "delete monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewWebPageSpeedMonitors(c).Delete("123"))
			},
		},
		{
			Name:       "delete monitor not found",
			StatusCode: 404,
			Fn: func(t *testing.T, c rest.Client) {
				err := NewWebPageSpeedMonitors(c).Delete("123")
				assert.True(t, apierrors.IsNotFound(err))
			},
		},
	})
}
