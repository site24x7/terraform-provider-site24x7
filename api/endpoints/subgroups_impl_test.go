package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubgroups(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create subgroup",
			ExpectedVerb: "POST",
			ExpectedPath: "/subgroups",
			ExpectedBody: validation.Fixture(t, "requests/create_subgroup.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				group := &api.Subgroup{
					DisplayName: "foo group",
					Description: "This is foo group",
					Monitors: []string{
						"726000000002460",
						"726000000002464",
					},
					TopGroupID:    "123",
					ParentGroupID: "456",
					Type:          2,
				}

				_, err := NewSubgroups(c).Create(group)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get subgroup",
			ExpectedVerb: "GET",
			ExpectedPath: "/subgroups/113770000041271035",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_subgroup.json"),
			Fn: func(t *testing.T, c rest.Client) {
				group, err := NewSubgroups(c).Get("113770000041271035")
				require.NoError(t, err)

				expected := &api.Subgroup{
					ID:          "113770000041271035",
					DisplayName: "foo group",
					Description: "This is foo group",
					Monitors: []string{
						"726000000002460",
						"726000000002464",
					},
					TopGroupID:    "123",
					ParentGroupID: "456",
					Type:          2,
				}

				assert.Equal(t, expected, group)
			},
		},
		{
			Name:         "list subgroups",
			ExpectedVerb: "GET",
			ExpectedPath: "/subgroups",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_subgroups.json"),
			Fn: func(t *testing.T, c rest.Client) {
				groups, err := NewSubgroups(c).List()
				require.NoError(t, err)

				expected := []*api.Subgroup{
					{
						ID:          "797300000123437",
						DisplayName: "misc",
						Description: "checks for misc sites",
						Monitors: []string{
							"726000000002460",
							"726000000002464",
						},
						ParentGroupID: "123",
						TopGroupID:    "456",
						Type:          2,
					},
					{
						ID:            "79123400003075053",
						DisplayName:   "api",
						ParentGroupID: "123",
						TopGroupID:    "456",
						Type:          2,
					},
					{
						ID:            "79730456703075223",
						DisplayName:   "web",
						ParentGroupID: "123",
						TopGroupID:    "456",
						Type:          2,
					},
				}

				assert.Equal(t, expected, groups)
			},
		},
		{
			Name:         "update subgroup",
			ExpectedVerb: "PUT",
			ExpectedPath: "/subgroups/123",
			ExpectedBody: validation.Fixture(t, "requests/update_subgroups.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				group := &api.Subgroup{
					ID:          "123",
					DisplayName: "foo group",
					Description: "This is foo group",
					Monitors: []string{
						"726000000002460",
						"726000000002464",
					},
					TopGroupID:    "123",
					ParentGroupID: "456",
					Type:          2,
				}

				_, err := NewSubgroups(c).Update(group)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete subgroup",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/subgroups/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewSubgroups(c).Delete("123"))
			},
		},
	})
}
