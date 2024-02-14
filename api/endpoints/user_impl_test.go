package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUsers(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create user",
			ExpectedVerb: "POST",
			ExpectedPath: "/users",
			ExpectedBody: validation.Fixture(t, "requests/create_user.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				group := &api.User{
					DisplayName:  "FooUser",
					Email:        "admin@test.com",
					UserGroupIDs: []string{"123", "987654", "987"},
					AlertSettings: map[string]interface{}{
						"email_format": 1,
						"critical":     []int{1},
						"down":         []int{1},
						"trouble":      []int{1},
						"up":           []int{1},
					},
					NotificationMedium:    []int{1},
					SelectionType:         0,
					UserRole:              1,
					ConsentForNonEUAlerts: false,
				}

				_, err := NewUsers(c).Create(group)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get user",
			ExpectedVerb: "GET",
			ExpectedPath: "/users/897654345678",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_user.json"),
			Fn: func(t *testing.T, c rest.Client) {
				group, err := NewUsers(c).Get("897654345678")
				require.NoError(t, err)

				expected := &api.User{
					ID:                 "897654345678",
					DisplayName:        "Admin User",
					Email:              "admin@test.com",
					NotificationMedium: []int{1},
					SelectionType:      0,
					UserRole:           1,
					UserGroupIDs:       []string{"87654643", "32434567890"},
				}

				assert.Equal(t, expected, group)
			},
		},
		{
			Name:         "list users",
			ExpectedVerb: "GET",
			ExpectedPath: "/users",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_users.json"),
			Fn: func(t *testing.T, c rest.Client) {
				groups, err := NewUsers(c).List()
				require.NoError(t, err)

				expected := []*api.User{
					{
						ID:                 "765432344567878",
						DisplayName:        "Tester",
						Email:              "admin@test.com",
						NotificationMedium: []int{1},
						SelectionType:      0,
						UserRole:           1,
						UserGroupIDs: []string{
							"654367756756",
							"232423435455",
						},
					},
					{
						ID:                 "797300000123437",
						DisplayName:        "misc",
						Email:              "admin@test.com",
						NotificationMedium: []int{1},
						SelectionType:      0,
						UserRole:           1,
						UserGroupIDs: []string{
							"79730000001337083",
							"12340000005000031",
						},
					},
				}

				assert.Equal(t, expected, groups)
			},
		},
		{
			Name:         "update user",
			ExpectedVerb: "PUT",
			ExpectedPath: "/users/123",
			ExpectedBody: validation.Fixture(t, "requests/update_user.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				group := &api.User{
					ID:                 "123",
					DisplayName:        "Admin User",
					Email:              "admin@test.com",
					NotificationMedium: []int{1},
					SelectionType:      0,
					AlertSettings: map[string]interface{}{
						"email_format": 1,
						"critical":     []int{1},
						"down":         []int{1},
						"trouble":      []int{1},
						"up":           []int{1},
					},
					UserRole:              1,
					UserGroupIDs:          []string{"87654643", "32434567890"},
					ConsentForNonEUAlerts: false,
				}

				_, err := NewUsers(c).Update(group)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete user",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/users/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewUsers(c).Delete("123"))
			},
		},
	})
}
