package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserGroups(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create user group",
			ExpectedVerb: "POST",
			ExpectedPath: "/user_groups",
			ExpectedBody: validation.Fixture(t, "requests/create_user_group.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				group := &api.UserGroup{
					DisplayName:      "FooUserGroup",
					Users:            []string{"123", "987654", "987"},
					AttributeGroupID: "9876",
				}

				_, err := NewUserGroups(c).Create(group)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get user group",
			ExpectedVerb: "GET",
			ExpectedPath: "/user_groups/897654345678",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_user_group.json"),
			Fn: func(t *testing.T, c rest.Client) {
				group, err := NewUserGroups(c).Get("897654345678")
				require.NoError(t, err)

				expected := &api.UserGroup{
					UserGroupID:      "897654345678",
					DisplayName:      "Admin Group",
					Users:            []string{"87654643", "32434567890"},
					AttributeGroupID: "89765467854",
				}

				assert.Equal(t, expected, group)
			},
		},
		{
			Name:         "list user groups",
			ExpectedVerb: "GET",
			ExpectedPath: "/user_groups",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_user_groups.json"),
			Fn: func(t *testing.T, c rest.Client) {
				groups, err := NewUserGroups(c).List()
				require.NoError(t, err)

				expected := []*api.UserGroup{
					{
						UserGroupID:      "765432344567878",
						DisplayName:      "Tester",
						AttributeGroupID: "098577445678",
						Users: []string{
							"654367756756",
							"232423435455",
						},
					},
					{
						UserGroupID: "797300000123437",
						DisplayName: "misc",
						Users: []string{
							"79730000001337083",
							"12340000005000031",
						},
					},
					{
						UserGroupID: "2434322",
						DisplayName: "foo",
					},
				}

				assert.Equal(t, expected, groups)
			},
		},
		// {
		// 	Name:         "update user group",
		// 	ExpectedVerb: "PUT",
		// 	ExpectedPath: "/user_groups/123",
		// 	ExpectedBody: validation.Fixture(t, "requests/update_user_group.json"),
		// 	StatusCode:   200,
		// 	ResponseBody: validation.JsonAPIResponseBody(t, nil),
		// 	Fn: func(t *testing.T, c rest.Client) {
		// 		group := &api.UserGroup{
		// 			UserGroupID:      "123",
		// 			DisplayName:      "Lead",
		// 			AttributeGroupID: "1234",
		// 			Users: []string{
		// 				"43255",
		// 				"76543",
		// 			},
		// 		}

		// 		_, err := NewUserGroups(c).Update(group)
		// 		require.NoError(t, err)
		// 	},
		// },
		{
			Name:         "delete user group",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/user_groups/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewUserGroups(c).Delete("123"))
			},
		},
	})
}
