package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAmazonMonitors(t *testing.T) {
	runTests(t, []*endpointTest{
		{
			name:         "create amazon monitor",
			expectedVerb: "POST",
			expectedPath: "/monitors",
			expectedBody: fixture(t, "requests/create_amazon_monitor.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
				amazonMonitor := &api.AmazonMonitor{
					DisplayName:           "Amazon Monitor Display Name",
					Type:                  "AMAZON",
					SecretKey:             "secretkey",
					AccessKey:             "accesskey",
					DiscoverFrequency:     1,
					DiscoverServices:      []string{"1"},
					NotificationProfileID: "123412341234123413",
					UserGroupIDs: []string{
						"123412341234123415",
					},
				}

				_, err := NewAmazonMonitors(c).Create(amazonMonitor)
				require.NoError(t, err)
			},
		},
		{
			name:         "get amazon monitor",
			expectedVerb: "GET",
			expectedPath: "/monitors/113770000041271035",
			statusCode:   200,
			responseBody: fixture(t, "responses/get_amazon_monitor.json"),
			fn: func(t *testing.T, c rest.Client) {
				amazon_monitor, err := NewAmazonMonitors(c).Get("113770000041271035")
				require.NoError(t, err)

				expected := &api.AmazonMonitor{
					DisplayName:           "Amazon Monitor Display Name",
					Type:                  "AMAZON",
					SecretKey:             "secretkey",
					AccessKey:             "accesskey",
					DiscoverFrequency:     1,
					DiscoverServices:      []string{"1"},
					NotificationProfileID: "123412341234123413",
					UserGroupIDs: []string{
						"123412341234123415",
					},
				}

				assert.Equal(t, expected, amazon_monitor)
			},
		},
		{
			name:         "list amazon monitors",
			expectedVerb: "GET",
			expectedPath: "/monitors",
			statusCode:   200,
			responseBody: fixture(t, "responses/list_amazon_monitors.json"),
			fn: func(t *testing.T, c rest.Client) {
				amazonMonitor, err := NewAmazonMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.AmazonMonitor{
					{
						MonitorID:             "123447000020646003",
						DisplayName:           "Amazon Monitor Display Name",
						Type:                  "AMAZON",
						SecretKey:             "secretkey",
						AccessKey:             "accesskey",
						DiscoverFrequency:     5,
						DiscoverServices:      []string{"1"},
						NotificationProfileID: "123412341234123413",
						UserGroupIDs: []string{
							"123412341234123415",
						},
					},
					{
						MonitorID:             "123447000020646008",
						DisplayName:           "Amazon Monitor Display Name",
						Type:                  "AMAZON",
						SecretKey:             "secretkey",
						AccessKey:             "accesskey",
						DiscoverFrequency:     5,
						DiscoverServices:      []string{"1"},
						NotificationProfileID: "123412341234123413",
						UserGroupIDs: []string{
							"123412341234123415",
						},
					},
				}

				assert.Equal(t, expected, amazonMonitor)
			},
		},
		{
			name:         "update amazon monitor",
			expectedVerb: "PUT",
			expectedPath: "/monitors/123",
			expectedBody: fixture(t, "requests/update_amazon_monitor.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
				amazon_monitor := &api.AmazonMonitor{
					MonitorID:             "123",
					DisplayName:           "foo",
					Type:                  "AMAZON",
					SecretKey:             "secretkey",
					AccessKey:             "accesskey",
					DiscoverFrequency:     5,
					DiscoverServices:      []string{"1"},
					NotificationProfileID: "123412341234123413",
					UserGroupIDs: []string{
						"123412341234123415",
					},
				}

				_, err := NewAmazonMonitors(c).Update(amazon_monitor)
				require.NoError(t, err)
			},
		},
		{
			name:         "delete amazon monitor",
			expectedVerb: "DELETE",
			expectedPath: "/monitors/123",
			statusCode:   200,
			fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewAmazonMonitors(c).Delete("123"))
			},
		},
	})
}
