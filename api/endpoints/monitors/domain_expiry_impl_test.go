package monitors

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDomainExpiryMonitors(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create rest api monitor",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitors",
			ExpectedBody: validation.Fixture(t, "requests/create_rest_api_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				domainExpiryMonitor := &api.DomainExpiryMonitor{
					DisplayName:           "foo",
					Type:                  string(api.DOMAINEXPIRY),
					LocationProfileID:     "456",
					NotificationProfileID: "789",
					MonitorGroups:         []string{"234", "567"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
				}

				_, err := NewDomainExpiryMonitors(c).Create(domainExpiryMonitor)
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
				domainExpiryMonitor, err := NewDomainExpiryMonitors(c).Get("897654345678")
				require.NoError(t, err)

				expected := &api.DomainExpiryMonitor{
					MonitorID:             "897654345678",
					DisplayName:           "foo",
					Type:                  string(api.RESTAPI),
					LocationProfileID:     "456",
					NotificationProfileID: "789",
					MonitorGroups:         []string{"234", "567"},
					UserGroupIDs:          []string{"123", "456"},
				}

				assert.Equal(t, expected, domainExpiryMonitor)
			},
		},
		{
			Name:         "list rest api monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_rest_api_monitors.json"),
			Fn: func(t *testing.T, c rest.Client) {
				domainExpiryMonitor, err := NewDomainExpiryMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.DomainExpiryMonitor{
					{
						MonitorID:             "897654345678",
						DisplayName:           "foo",
						Type:                  string(api.RESTAPI),
						Timeout:               10,
						LocationProfileID:     "456",
						NotificationProfileID: "789",
						MonitorGroups:         []string{"234", "567"},
						UserGroupIDs:          []string{"123", "456"},
					},
					{
						MonitorID:             "933654345678",
						DisplayName:           "foo",
						Type:                  string(api.RESTAPI),
						Timeout:               10,
						LocationProfileID:     "456",
						NotificationProfileID: "789",
						MonitorGroups:         []string{"234", "567"},
						UserGroupIDs:          []string{"123", "456"},
					},
				}

				assert.Equal(t, expected, domainExpiryMonitor)
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
				domainExpiryMonitor := &api.DomainExpiryMonitor{
					MonitorID:             "123",
					DisplayName:           "foo",
					Type:                  string(api.RESTAPI),
					Timeout:               10,
					LocationProfileID:     "456",
					NotificationProfileID: "789",
					MonitorGroups:         []string{"234", "567"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
				}

				_, err := NewDomainExpiryMonitors(c).Update(domainExpiryMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete rest api monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewDomainExpiryMonitors(c).Delete("123"))
			},
		},
	})
}
