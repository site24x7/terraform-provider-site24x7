package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserGroups(t *testing.T) {
	runTests(t, []*endpointTest{
		{
			name:         "create user group",
			expectedVerb: "POST",
			expectedPath: "/user_groups",
			expectedBody: fixture(t, "requests/create_user_group.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
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
			name:         "get user group",
			expectedVerb: "GET",
			expectedPath: "/user_groups/897654345678",
			statusCode:   200,
			responseBody: fixture(t, "responses/get_user_group.json"),
			fn: func(t *testing.T, c rest.Client) {
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
			name:         "list user groups",
			expectedVerb: "GET",
			expectedPath: "/user_groups",
			statusCode:   200,
			responseBody: fixture(t, "responses/list_user_groups.json"),
			fn: func(t *testing.T, c rest.Client) {
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
		{
			name:         "update user group",
			expectedVerb: "PUT",
			expectedPath: "/user_groups/123",
			expectedBody: fixture(t, "requests/update_user_group.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
				group := &api.UserGroup{
					UserGroupID:      "123",
					DisplayName:      "Lead",
					AttributeGroupID: "1234",
					Users: []string{
						"43255",
						"76543",
					},
				}

				_, err := NewUserGroups(c).Update(group)
				require.NoError(t, err)
			},
		},
		{
			name:         "delete user group",
			expectedVerb: "DELETE",
			expectedPath: "/user_groups/123",
			statusCode:   200,
			fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewUserGroups(c).Delete("123"))
			},
		},
	})
}
