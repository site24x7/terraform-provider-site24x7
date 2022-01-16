package integration

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestThirdPartyIntegrations(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "list third party integrations",
			ExpectedVerb: "GET",
			ExpectedPath: "/third_party_services",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "api/endpoints/testdata/fixtures/responses/list_third_party_integrations.json"),
			Fn: func(t *testing.T, c rest.Client) {
				thirdPartyIntegrations, err := NewThirdpartyIntegrations(c).List()
				require.NoError(t, err)

				expected := []*api.ThirdPartyIntegrations{
					{
						Name:          "OpsGenie Integration With Site24x7",
						ServiceID:     "113770000023231022",
						ServiceStatus: 0,
						SelectionType: 0,
						TroubleAlert:  false,
						Type:          10,
					},
					{
						Name:          "Site24x7-Slack Integration",
						ServiceID:     "113770000023231023",
						ServiceStatus: 0,
						SelectionType: 0,
						SenderName:    "Site24x7",
						Title:         "$MONITORNAME is $STATUS",
						TroubleAlert:  false,
						Type:          5,
					},
				}

				assert.Equal(t, expected, thirdPartyIntegrations)
			},
		},
		{
			Name:         "delete third party integration",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/integration/thirdparty_service/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewThirdpartyIntegrations(c).Delete("123"))
			},
		},
	})
}
