package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestURLActions(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create it_automation",
			ExpectedVerb: "POST",
			ExpectedPath: "/it_automation",
			ExpectedBody: validation.Fixture(t, "requests/create_it_automation.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				automation := &api.URLAction{
					ActionType:    2,
					ActionTimeout: 30,
					ActionMethod:  "P",
					ActionName:    "takeaction",
					ActionUrl:     "testing.tld",
				}
				_, err := NewURLActions(c).Create(automation)
				require.NoError(t, err)
			},
		},
		{
			Name:         "create it_automation error",
			StatusCode:   500,
			ResponseBody: []byte("whoops"),
			Fn: func(t *testing.T, c rest.Client) {
				_, err := NewURLActions(c).Create(&api.URLAction{})
				assert.True(t, apierrors.HasStatusCode(err, 500))
			},
		},
		{
			Name:         "get it_automation",
			ExpectedVerb: "GET",
			ExpectedPath: "/it_automation/123",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_it_automation.json"),
			Fn: func(t *testing.T, c rest.Client) {
				urlAction, err := NewURLActions(c).Get("123")
				require.NoError(t, err)

				expected := &api.URLAction{
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
				assert.Equal(t, expected, urlAction)
			},
		},
		{
			Name:         "list it_automations",
			ExpectedVerb: "GET",
			ExpectedPath: "/it_automation",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_it_automations.json"),
			Fn: func(t *testing.T, c rest.Client) {
				urlActions, err := NewURLActions(c).List()
				require.NoError(t, err)

				expected := []*api.URLAction{
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

				assert.Equal(t, expected, urlActions)
			},
		},
		{
			Name:         "update it_automation",
			ExpectedVerb: "PUT",
			ExpectedPath: "/it_automation/123",
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				urlAction := &api.URLAction{
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

				_, err := NewURLActions(c).Update(urlAction)
				require.NoError(t, err)
			},
		},
		{
			Name:       "update create_it_automation error",
			StatusCode: 400,
			ResponseBody: validation.JsonBody(t, &api.ErrorResponse{
				ErrorCode: 123,
				Message:   "bad request",
				ErrorInfo: map[string]interface{}{"foo": "bar"},
			}),
			Fn: func(t *testing.T, c rest.Client) {
				_, err := NewURLActions(c).Update(&api.URLAction{})
				assert.True(t, apierrors.HasStatusCode(err, 400))
			},
		},
		{
			Name:         "delete it_automation",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/it_automation/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewURLActions(c).Delete("123"))
			},
		},
		{
			Name:       "delete it_automation not found",
			StatusCode: 404,
			Fn: func(t *testing.T, c rest.Client) {
				err := NewURLActions(c).Delete("123")
				assert.True(t, apierrors.IsNotFound(err))
			},
		},
	})
}
