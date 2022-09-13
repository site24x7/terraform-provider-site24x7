package monitors

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAmazonMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create amazon monitor",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitors",
			ExpectedBody: validation.Fixture(t, "requests/create_amazon_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				amazonMonitor := &api.AmazonMonitor{
					DisplayName:           "Amazon Monitor Display Name",
					Type:                  "AMAZON",
					RoleARN:               "secretkey",
					AWSExternalID:         "accesskey",
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
			Name:         "get amazon monitor",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors/113770000041271035",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_amazon_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				amazon_monitor, err := NewAmazonMonitors(c).Get("113770000041271035")
				require.NoError(t, err)

				expected := &api.AmazonMonitor{
					DisplayName:           "Amazon Monitor Display Name",
					Type:                  "AMAZON",
					RoleARN:               "secretkey",
					AWSExternalID:         "accesskey",
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
			Name:         "list amazon monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_amazon_monitors.json"),
			Fn: func(t *testing.T, c rest.Client) {
				amazonMonitor, err := NewAmazonMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.AmazonMonitor{
					{
						MonitorID:             "123447000020646003",
						DisplayName:           "Amazon Monitor Display Name",
						Type:                  "AMAZON",
						RoleARN:               "secretkey",
						AWSExternalID:         "accesskey",
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
						RoleARN:               "secretkey",
						AWSExternalID:         "accesskey",
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
			Name:         "update amazon monitor",
			ExpectedVerb: "PUT",
			ExpectedPath: "/monitors/123",
			ExpectedBody: validation.Fixture(t, "requests/update_amazon_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				amazon_monitor := &api.AmazonMonitor{
					MonitorID:             "123",
					DisplayName:           "foo",
					Type:                  "AMAZON",
					RoleARN:               "secretkey",
					AWSExternalID:         "accesskey",
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
			Name:         "delete amazon monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewAmazonMonitors(c).Delete("123"))
			},
		},
	})
}
