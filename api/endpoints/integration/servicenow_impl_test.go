package integration

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServiceNow(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create servicenow integration",
			ExpectedVerb: "POST",
			ExpectedPath: "/integration/service_now",
			ExpectedBody: validation.Fixture(t, "requests/create_servicenow_integration.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				serviceNowIntegration := &api.ServiceNowIntegration{
					Name:          "Site24x7-ServiceNow Integration",
					InstanceURL:   "https://www.example.com",
					UserName:      "username",
					Password:      "password",
					TroubleAlert:  true,
					CriticalAlert: false,
					DownAlert:     false,
					SelectionType: 0,
					SenderName:    "Site24x7",
					Title:         "$MONITOR_NAME is $STATUS",
					AlertTagIDs:   []string{"113770000023231001"},
				}

				_, err := NewServiceNow(c).Create(serviceNowIntegration)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get servicenow integration",
			ExpectedVerb: "GET",
			ExpectedPath: "/integration/service_now/113770000023231022",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_servicenow_integration.json"),
			Fn: func(t *testing.T, c rest.Client) {
				servicenow_integration, err := NewServiceNow(c).Get("113770000023231022")
				require.NoError(t, err)

				expected := &api.ServiceNowIntegration{
					Name:          "Site24x7-ServiceNow Integration",
					InstanceURL:   "https://www.example.com",
					ServiceID:     "113770000023231022",
					UserName:      "username",
					Password:      "password",
					ServiceStatus: 0,
					SelectionType: 0,
					SenderName:    "Site24x7",
					Title:         "$MONITOR_NAME is $STATUS",
					AlertTagIDs:   []string{"113770000023231001"},
				}

				assert.Equal(t, expected, servicenow_integration)
			},
		},
		{
			Name:         "update servicenow integration",
			ExpectedVerb: "PUT",
			ExpectedPath: "/integration/service_now/123",
			ExpectedBody: validation.Fixture(t, "requests/update_servicenow_integration.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				servicenow_integration := &api.ServiceNowIntegration{
					Name:          "Site24x7-ServiceNow Integration",
					ServiceID:     "123",
					InstanceURL:   "https://www.example.com",
					UserName:      "username",
					Password:      "password",
					TroubleAlert:  true,
					CriticalAlert: false,
					DownAlert:     false,
					SelectionType: 2,
					SenderName:    "Site24x7",
					Monitors:      []string{"113770000023231032", "113770000023231043"},
					Title:         "$MONITOR_NAME is $STATUS",
					AlertTagIDs:   []string{"113770000023231001", "113770000023231002"},
				}

				_, err := NewServiceNow(c).Update(servicenow_integration)
				require.NoError(t, err)
			},
		},
	})
}
