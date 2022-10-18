package integration

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConnectwiseIntegration(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create connectwise integration",
			ExpectedVerb: "POST",
			ExpectedPath: "/integration/connectwise",
			ExpectedBody: validation.Fixture(t, "requests/create_connectwise_integration.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				connectwiseIntegration := &api.ConnectwiseIntegration{
					Name:          "Site24x7-Connectwise Integration",
					URL:           "https://wefvsefv.connectwisedev.com/",
					Company:       "zylker_c",
					PublicKey:	   "KefwvwfrmAb",
					PrivateKey:    "wegraaeagt",
					CompanyId:     "GreenInc",
					SelectionType: 0,
					CloseStatus:    "Closed (resolved)",
					TroubleAlert:  true,
					CriticalAlert: false,
					DownAlert:     false,
					AlertTagIDs:   []string{"123450023231001"},
				}

				_, err := NewConnectwise(c).Create(connectwiseIntegration)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get connectwise integration",
			ExpectedVerb: "GET",
			ExpectedPath: "/integration/connectwise/113770000023231022",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_connectwise_integration.json"),
			Fn: func(t *testing.T, c rest.Client) {
				connectwise_integration, err := NewConnectwise(c).Get("113770000023231022")
				require.NoError(t, err)

				expected := &api.ConnectwiseIntegration{
					Name:          "Site24x7-Connectwise Integration",
					URL:           "https://wefvsefv.connectwisedev.com/",
					ServiceID:     "123456306001",
					ServiceStatus: 0,
					SelectionType: 0,
					CloseStatus:   "Closed (resolved)",
					Company:       "zylker_c",
					PublicKey:	   "KefwvwfrmAb",
					PrivateKey:    "wegraaeagt",
					CompanyId:     "GreenInc",
					AlertTagIDs:   []string{"123450023231001"},
				}

				assert.Equal(t, expected, connectwise_integration)
			},
		},
		{
			Name:         "update connetwise integration",
			ExpectedVerb: "PUT",
			ExpectedPath: "/integration/connectwise/123",
			ExpectedBody: validation.Fixture(t, "requests/update_connectwise_integration.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				connectwise_integration := &api.ConnectwiseIntegration{
					Name:          "Site24x7-Connectwise Integration",
					URL:           "https://wefvsefv.connectwisedev.com/",
					ServiceID:     "123",
					SelectionType: 2,
					Company:       "zylker_c",
					PublicKey:	   "KefwvwfrmAb",
					PrivateKey:    "wegraaeagt",
					CompanyId:     "GreenInc",
					CloseStatus:   "Closed (resolved)",
					TroubleAlert:  true,
					CriticalAlert: false,
					DownAlert:     false,
					Monitors:      []string{"12345623231032", "1234568231043"},
					AlertTagIDs:   []string{"12345623231001", "1234563231002"},
				}

				_, err := NewConnectwise(c).Update(connectwise_integration)
				require.NoError(t, err)
			},
		},
	})
}
