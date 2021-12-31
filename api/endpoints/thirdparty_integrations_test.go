package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestThirdPartyIntegrations(t *testing.T) {
	runTests(t, []*endpointTest{
		{
			name:         "list third party integrations",
			expectedVerb: "GET",
			expectedPath: "/third_party_services",
			statusCode:   200,
			responseBody: fixture(t, "responses/list_third_party_integrations.json"),
			fn: func(t *testing.T, c rest.Client) {
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
			name:         "delete third party integration",
			expectedVerb: "DELETE",
			expectedPath: "/integration/thirdparty_service/123",
			statusCode:   200,
			fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewThirdpartyIntegrations(c).Delete("123"))
			},
		},
	})
}
