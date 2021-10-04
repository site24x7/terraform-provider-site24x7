package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestURLAutomations(t *testing.T) {
	runTests(t, []*endpointTest{
		{
			name:         "create it_automation",
			expectedVerb: "POST",
			expectedPath: "/it_automation",
			expectedBody: fixture(t, "requests/create_it_automation.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
				automation := &api.URLAutomation{
					ActionType:    2,
					ActionTimeout: 30,
					ActionMethod:  "P",
					ActionName:    "takeaction",
					ActionUrl:     "testing.tld",
				}
				_, err := NewURLAutomations(c).Create(automation)
				require.NoError(t, err)
			},
		},
		{
			name:         "create it_automation error",
			statusCode:   500,
			responseBody: []byte("whoops"),
			fn: func(t *testing.T, c rest.Client) {
				_, err := NewURLAutomations(c).Create(&api.URLAutomation{})
				assert.True(t, apierrors.HasStatusCode(err, 500))
			},
		},
		{
			name:         "get it_automation",
			expectedVerb: "GET",
			expectedPath: "/it_automation/123",
			statusCode:   200,
			responseBody: fixture(t, "responses/get_it_automation.json"),
			fn: func(t *testing.T, c rest.Client) {
				urlAutomation, err := NewURLAutomations(c).Get("123")
				require.NoError(t, err)

				expected := &api.URLAutomation{
					ActionID:               "123",
					ActionName:             "takeaction",
					ActionUrl:              "testing.tld",
					ActionTimeout:          30,
					ActionType:             2,
					ActionMethod:           "P",
					SendInJsonFormat:       true,
					SendCustomParameters:   true,
					CustomParameters:       "{\"message_type\":\"TEST\"}",
					SendIncidentParameters: true,
				}
				assert.Equal(t, expected, urlAutomation)
			},
		},
		{
			name:         "list it_automations",
			expectedVerb: "GET",
			expectedPath: "/it_automation",
			statusCode:   200,
			responseBody: fixture(t, "responses/list_it_automations.json"),
			fn: func(t *testing.T, c rest.Client) {
				urlAutomations, err := NewURLAutomations(c).List()
				require.NoError(t, err)

				expected := []*api.URLAutomation{
					{
						ActionID:               "123",
						ActionType:             2,
						ActionMethod:           "P",
						ActionName:             "takeaction",
						CustomParameters:       "{\"message_type\":\"TEST\"}",
						SendInJsonFormat:       true,
						SendCustomParameters:   true,
						ActionUrl:              "testing.tld",
						ActionTimeout:          30,
						SendIncidentParameters: true,
					},
					{
						ActionID:         "321",
						ActionType:       4,
						ActionMethod:     "PP",
						ActionName:       "action",
						SendInJsonFormat: true,
						ActionUrl:        "testing.tld",
						ActionTimeout:    30,
					},
				}

				assert.Equal(t, expected, urlAutomations)
			},
		},
		{
			name:         "update it_automation",
			expectedVerb: "PUT",
			expectedPath: "/it_automation/123",
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
				urlAutomation := &api.URLAutomation{
					ActionID:               "123",
					ActionType:             1,
					ActionMethod:           "P",
					ActionName:             "takeaction",
					SendInJsonFormat:       true,
					SendCustomParameters:   true,
					ActionUrl:              "https://alert.generic.tld",
					ActionTimeout:          30,
					SendIncidentParameters: true,
				}

				_, err := NewURLAutomations(c).Update(urlAutomation)
				require.NoError(t, err)
			},
		},
		{
			name:       "update create_it_automation error",
			statusCode: 400,
			responseBody: jsonBody(t, &api.ErrorResponse{
				ErrorCode: 123,
				Message:   "bad request",
				ErrorInfo: map[string]interface{}{"foo": "bar"},
			}),
			fn: func(t *testing.T, c rest.Client) {
				_, err := NewURLAutomations(c).Update(&api.URLAutomation{})
				assert.True(t, apierrors.HasStatusCode(err, 400))
			},
		},
		{
			name:         "delete it_automation",
			expectedVerb: "DELETE",
			expectedPath: "/it_automation/123",
			statusCode:   200,
			fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewURLAutomations(c).Delete("123"))
			},
		},
		{
			name:       "delete it_automation not found",
			statusCode: 404,
			fn: func(t *testing.T, c rest.Client) {
				err := NewURLAutomations(c).Delete("123")
				assert.True(t, apierrors.IsNotFound(err))
			},
		},
	})
}
