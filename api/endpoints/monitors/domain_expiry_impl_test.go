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
			Name:         "create domain-expiry-monitor",
			ExpectedVerb: "POST",
			ExpectedPath: "/monitors",
			ExpectedBody: validation.Fixture(t, "requests/create_domain_expiry_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				domainExpiryMonitor := &api.DomainExpiryMonitor{
					DisplayName:           "Domain Expiry Monitor",
					HostName:              "www.example.com",
					DomainName:            "www.example.com",
					Port:                  443,
					UseIPV6:               true,
					Timeout:               10,
					ExpireDays:            30,
					OnCallScheduleID:      "234",
					IgnoreRegistryDate:    false,
					LocationProfileID:     "456",
					NotificationProfileID: "789",
					MonitorGroups:         []string{"234", "567"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
					PerformAutomation:     true,
					MatchCase:             true,
				}

				_, err := NewDomainExpiryMonitors(c).Create(domainExpiryMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get domain expiry monitor",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors/897654345678",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_domain_expiry_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				domainExpiryMonitor, err := NewDomainExpiryMonitors(c).Get("897654345678")
				require.NoError(t, err)

				expected := &api.DomainExpiryMonitor{
					MonitorID:             "897654345678",
					DisplayName:           "Domain Expiry Monitor",
					HostName:              "www.example.com",
					DomainName:            "www.example.com",
					Port:                  443,
					UseIPV6:               true,
					Timeout:               10,
					ExpireDays:            30,
					OnCallScheduleID:      "234",
					IgnoreRegistryDate:    false,
					LocationProfileID:     "456",
					NotificationProfileID: "789",
					MonitorGroups:         []string{"234", "567"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
					PerformAutomation:     false,
					MatchCase:             true,
				}

				assert.Equal(t, expected, domainExpiryMonitor)
			},
		},
		{
			Name:         "list domain expiry monitors",
			ExpectedVerb: "GET",
			ExpectedPath: "/monitors",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_domain_expiry_monitor.json"),
			Fn: func(t *testing.T, c rest.Client) {
				domainExpiryMonitor, err := NewDomainExpiryMonitors(c).List()
				require.NoError(t, err)

				expected := []*api.DomainExpiryMonitor{
					{
						MonitorID:             "897654345678",
						DisplayName:           "Domain Expiry Monitor",
						HostName:              "www.example.com",
						DomainName:            "www.example.com",
						Port:                  443,
						UseIPV6:               true,
						Timeout:               10,
						ExpireDays:            30,
						LocationProfileID:     "456",
						NotificationProfileID: "789",
						OnCallScheduleID:      "234",
						IgnoreRegistryDate:    false,
						MonitorGroups:         []string{"234", "567"},
						UserGroupIDs:          []string{"123", "456"},
						TagIDs:                []string{"123"},
						PerformAutomation:     true,
						MatchCase:             true,
					},
					{
						MonitorID:             "654568778999889",
						DisplayName:           "Domain Expiry Monitor",
						HostName:              "www.example.com",
						DomainName:            "www.example.com",
						Port:                  443,
						UseIPV6:               true,
						Timeout:               10,
						ExpireDays:            30,
						LocationProfileID:     "456",
						NotificationProfileID: "789",
						OnCallScheduleID:      "234",
						IgnoreRegistryDate:    false,
						MonitorGroups:         []string{"234", "567"},
						UserGroupIDs:          []string{"123", "456"},
						TagIDs:                []string{"123"},
						PerformAutomation:     true,
						MatchCase:             true,
					},
				}

				assert.Equal(t, expected, domainExpiryMonitor)
			},
		},
		{
			Name:         "update domain expiry monitor",
			ExpectedVerb: "PUT",
			ExpectedPath: "/monitors/897654345678",
			ExpectedBody: validation.Fixture(t, "requests/update_domain_expiry_monitor.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				domainExpiryMonitor := &api.DomainExpiryMonitor{
					MonitorID:             "897654345678",
					DisplayName:           "Domain Expiry Monitor",
					HostName:              "www.example.com",
					DomainName:            "www.example.com",
					Port:                  443,
					UseIPV6:               true,
					Timeout:               10,
					ExpireDays:            30,
					LocationProfileID:     "456",
					NotificationProfileID: "789",
					OnCallScheduleID:      "234",
					IgnoreRegistryDate:    false,
					MonitorGroups:         []string{"234", "567"},
					UserGroupIDs:          []string{"123", "456"},
					TagIDs:                []string{"123"},
					PerformAutomation:     true,
					MatchCase:             true,
				}

				_, err := NewDomainExpiryMonitors(c).Update(domainExpiryMonitor)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete domain expiry monitor",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/monitors/897654345678",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewDomainExpiryMonitors(c).Delete("123"))
			},
		},
	})
}
