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

func TestDNSServerMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "Create DNS server monitor",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitors",
			ExpectedBody: validation.Fixture(t, "requests/create_dns_server_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				monitor := &api.DNSServerMonitor{
					UseIPV6:               false,
					Type:                  "DNS",
					LookupType:            1,
					DNSSEC:                false,
					DisplayName:           "DNS Server Impl Test",
					DNSHost:               "8.8.8.8",
					DNSPort:               "53",
					DomainName:            "www.example.com",
					CheckFrequency:        "5",
					Timeout:               10,
					DeepDiscovery:         false,
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
					MonitorGroups: []string{
						"123412341234123416",
						"123412341234123417",
					},
					ActionIDs: []api.ActionRef{
						{
							ActionID:  "123412341234123418",
							AlertType: 20,
						},
					},
					SearchConfig: []api.SearchConfig{
						{
							Addr: "1.2.3.4",
							TTLO: 60,
							TTL:  60,
						},
					},
				}

				_, err := NewDNSServerMonitors(c).Create(monitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "create monitor error",
			StatusCode:   500,
			ResponseBody: []byte("whoops"),
			Fn: func(t *testing.T, c rest.Client) {
				_, err := NewDNSServerMonitors(c).Create(&api.DNSServerMonitor{})
				assert.True(t, apierrors.HasStatusCode(err, 500))
			},
		},
		{
			Name:         "get monitor",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors/123412341234123411",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_dns_server_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				monitor, err := NewDNSServerMonitors(c).Get("123412341234123411")
				require.NoError(t, err)

				expected := &api.DNSServerMonitor{
					MonitorID:      "123412341234123411",
					DisplayName:    "DNS Server Impl Test",
					DNSHost:        "8.8.8.8",
					DNSPort:        "53",
					Type:           "DNS",
					UseIPV6:        false,
					LookupType:     1,
					DomainName:     "www.example.com",
					CheckFrequency: "5",
					Timeout:        10,
					DeepDiscovery:  false,
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
					MonitorGroups: []string{
						"123412341234123416",
						"123412341234123417",
					},
					ThresholdProfileID:    "123412341234123414",
					NotificationProfileID: "123412341234123413",
					ActionIDs: []api.ActionRef{
						{
							ActionID:  "123412341234123418",
							AlertType: 20,
						},
					},
					SearchConfig: []api.SearchConfig{
						{
							Addr: "5.6.7.8",
							TTLO: 59,
							TTL:  59,
						},
					},
				}

				assert.Equal(t, expected, monitor)
			},
		},
		{
			Name:         "list monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_dns_server_monitors.json"),
			Fn: func(t *testing.T, c rest.Client) {
				monitor, err := NewDNSServerMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.DNSServerMonitor{
					{
						MonitorID:      "12340000016033021",
						DisplayName:    "foo.bar",
						LookupType:     1,
						DNSHost:        "8.8.8.8",
						DNSPort:        "53",
						UseIPV6:        false,
						DomainName:     "www.example.com",
						CheckFrequency: "5",
						Timeout:        10,
						DeepDiscovery:  false,
						Type:           "DNS",
						UserGroupIDs: []string{
							"12340000000018013",
						},
						LocationProfileID: "12340000001806001",
						MonitorGroups: []string{
							"12340000005749001",
						},
						ThresholdProfileID:    "12340000001812001",
						NotificationProfileID: "12340000003579003",
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
						SearchConfig: []api.SearchConfig{
							{
								Addr: "1.1.1.1",
								TTLO: 58,
								TTL:  58,
							},
						},
					},
					{
						MonitorID:      "12340000016108026",
						DisplayName:    "some.api.tld",
						DNSHost:        "8.8.8.8",
						DNSPort:        "53",
						UseIPV6:        false,
						DomainName:     "www.example.com",
						CheckFrequency: "5",
						Timeout:        10,
						DeepDiscovery:  false,
						Type:           "DNS",
						UserGroupIDs: []string{
							"12340000015652005",
						},
						LocationProfileID: "12340000001806001",
						MonitorGroups: []string{
							"12340000002807001",
						},
						ThresholdProfileID:    "12340000001812001",
						NotificationProfileID: "12340000003579003",
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
				monitor := &api.DNSServerMonitor{MonitorID: "456", DisplayName: "bar"}

				monitor, err := NewDNSServerMonitors(c).Update(monitor)
				require.NoError(t, err)

				expected := &api.DNSServerMonitor{
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
				_, err := NewDNSServerMonitors(c).Update(&api.DNSServerMonitor{})
				assert.True(t, apierrors.HasStatusCode(err, 400))
			},
		},
		{
			Name:         "delete monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewDNSServerMonitors(c).Delete("123"))
			},
		},
		{
			Name:       "delete monitor not found",
			StatusCode: 404,
			Fn: func(t *testing.T, c rest.Client) {
				err := NewDNSServerMonitors(c).Delete("123")
				assert.True(t, apierrors.IsNotFound(err))
			},
		},
	})
}
