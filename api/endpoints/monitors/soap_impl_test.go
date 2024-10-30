package monitors

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSOAPMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create soap monitor",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitors",
			ExpectedBody: validation.Fixture(t, "requests/create_soap_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				soapMonitors := &api.SOAPMonitor{
					DisplayName:    "SOAP Monitor",
					Website:        "www.example.com",
					RequestParam:   "",
					Type:           "SOAP",
					UseIPV6:        true,
					SSLProtocol:    "",
					Timeout:        10,
					HTTPMethod:     "",
					CheckFrequency: "5",
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
					OnCallScheduleID:      "23524543545245",
					LocationProfileID:     "123412341234123412",
					NotificationProfileID: "123412341234123412",
					MonitorGroups:         []string{"234", "567"},
					DependencyResourceIDs: []string{"123", "456"},
					UserGroupIDs:          []string{"123", "456"},
					PerformAutomation:     true,
				}

				_, err := NewSOAPMonitors(c).Create(soapMonitors)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get soap monitor",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors/897654345678",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_soap_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				soapMonitors, err := NewSOAPMonitors(c).Get("897654345678")
				require.NoError(t, err)

				expected := &api.SOAPMonitor{
					MonitorID:             "897654345678",
					DisplayName:           "SOAP Monitor",
					Website:               "www.example.com",
					RequestParam:          "",
					Type:                  "SOAP",
					UseIPV6:               true,
					Timeout:               10,
					CheckFrequency:        "5",
					OnCallScheduleID:      "23524543545245",
					LocationProfileID:     "123412341234123412",
					NotificationProfileID: "123412341234123412",
					ThresholdProfileID:    "123412341234123414",
					MonitorGroups:         []string{"234", "567"},
					DependencyResourceIDs: []string{"123", "456"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
					PerformAutomation:     false,
				}

				assert.Equal(t, expected, soapMonitors)
			},
		},
		{
			Name:         "list soap monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_soap_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				soapMonitors, err := NewSOAPMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.SOAPMonitor{
					{
						MonitorID:             "897654345678",
						DisplayName:           "SOAP Monitor",
						Website:               "www.example.com",
						Type:                  "SOAP",
						UseIPV6:               true,
						Timeout:               10,
						CheckFrequency:        "5",
						OnCallScheduleID:      "23524543545245",
						LocationProfileID:     "123412341234123412",
						NotificationProfileID: "123412341234123412",
						ThresholdProfileID:    "123412341234123414",
						MonitorGroups:         []string{"234", "567"},
						DependencyResourceIDs: []string{"123", "456"},
						UserGroupIDs:          []string{"123", "456"},
						TagIDs:                []string{"123"},
						PerformAutomation:     true,
					},
					{
						MonitorID:             "654568778999889",
						DisplayName:           "SOAP Monitor",
						Website:               "www.example.com",
						Type:                  "SOAP",
						UseIPV6:               true,
						Timeout:               10,
						CheckFrequency:        "5",
						OnCallScheduleID:      "23524543545245",
						LocationProfileID:     "123412341234123412",
						NotificationProfileID: "123412341234123412",
						ThresholdProfileID:    "123412341234123414",
						MonitorGroups:         []string{"234", "567"},
						DependencyResourceIDs: []string{"123", "456"},
						UserGroupIDs:          []string{"123", "456"},
						TagIDs:                []string{"123"},
						PerformAutomation:     true,
					},
				}

				assert.Equal(t, expected, soapMonitors)
			},
		},
		{
			Name:         "update soap monitor",
			ExpectedVerb: "PUT",
			ExpectedPath: "/monitors/897654345678",
			ExpectedBody: validation.Fixture(t, "requests/update_soap_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				soapMonitors := &api.SOAPMonitor{
					MonitorID:   "897654345678",
					DisplayName: "SOAP Monitor",
					Website:     "www.example.com",
					UseIPV6:     false,
					Type:        "SOAP",

					Timeout:        10,
					CheckFrequency: "5",
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
					OnCallScheduleID:      "23524543545245",
					LocationProfileID:     "123412341234123412",
					NotificationProfileID: "123412341234123412",
					ThresholdProfileID:    "123412341234123414",
					MonitorGroups:         []string{"234", "567"},
					DependencyResourceIDs: []string{"123", "456"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
					PerformAutomation:     true,
				}

				_, err := NewSOAPMonitors(c).Update(soapMonitors)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete soap monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/897654345678",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewSOAPMonitors(c).Delete("897654345678"))
			},
		},
	})
}
