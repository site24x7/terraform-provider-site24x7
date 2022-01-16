package integration

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpsgenieIntegration(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create opsgenie integration",
			ExpectedVerb: "POST",
			ExpectedPath: "/integration/opsgenie",
			ExpectedBody: validation.Fixture(t, "requests/create_opsgenie_integration.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				opsgenieIntegration := &api.OpsgenieIntegration{
					Name:          "OpsGenie Integration With Site24x7",
					URL:           "https://api.opsgenie.com/v1/json/site24x7?apiKey=a19y1cdd-bz7a-455a-z4b1-c1528323502s",
					SelectionType: 0,
					Monitors:      []string{"6111000000000068", "6111000000000130", "6111000000015045", "6111000000015057", "6111000000015069", "6111000000015083"},
					TroubleAlert:  false,
					ManualResolve: true,
					AlertTagIDs:   []string{"113770000023231001"},
				}

				_, err := NewOpsgenie(c).Create(opsgenieIntegration)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get opsgenie integration",
			ExpectedVerb: "GET",
			ExpectedPath: "/integration/opsgenie/113770000041271035",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_opsgenie_integration.json"),
			Fn: func(t *testing.T, c rest.Client) {
				opsgenie_integration, err := NewOpsgenie(c).Get("113770000041271035")
				require.NoError(t, err)

				expected := &api.OpsgenieIntegration{
					Name:          "OpsGenie Integration With Site24x7",
					ServiceID:     "113770000023231022",
					ServiceStatus: 0,
					URL:           "https://api.opsgenie.com/v1/json/site24x7?apiKey=a19y1cdd-bz7a-455a-z4b1-c1528323502s",
					SelectionType: 0,
					Monitors:      []string{"6111000000000068", "6111000000000130", "6111000000015045", "6111000000015057", "6111000000015069", "6111000000015083"},
					TroubleAlert:  false,
					ManualResolve: true,
					AlertTagIDs:   []string{"113770000023231001"},
				}

				assert.Equal(t, expected, opsgenie_integration)
			},
		},
		{
			Name:         "update opsgenie integration",
			ExpectedVerb: "PUT",
			ExpectedPath: "/integration/opsgenie/123",
			ExpectedBody: validation.Fixture(t, "requests/update_opsgenie_integration.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				opsgenie_integration := &api.OpsgenieIntegration{
					Name:          "Update OpsGenie Integration With Site24x7",
					URL:           "https://api.opsgenie.com/v1/json/site24x7?apiKey=a19y1cdd-bz7a-455a-z4b1-c1528323502s",
					ServiceID:     "123",
					SelectionType: 2,
					Monitors:      []string{"6111000000000068", "6111000000000130", "6111000000015045"},
					TroubleAlert:  false,
					ManualResolve: true,
					AlertTagIDs:   []string{"113770000023231001", "113770000023231002"},
				}

				_, err := NewOpsgenie(c).Update(opsgenie_integration)
				require.NoError(t, err)
			},
		},
	})
}
