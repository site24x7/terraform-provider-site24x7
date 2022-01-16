package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTags(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create tags",
			ExpectedVerb: "POST",
			ExpectedPath: "/tags",
			ExpectedBody: validation.Fixture(t, "requests/create_tag.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				tagCreate := &api.Tag{
					TagName:  "foobar",
					TagValue: "baz",
					TagColor: "#B7DA9E",
				}

				_, err := NewTags(c).Create(tagCreate)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get tags",
			ExpectedVerb: "GET",
			ExpectedPath: "/tags/113770000041271035",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_tag.json"),
			Fn: func(t *testing.T, c rest.Client) {
				group, err := NewTags(c).Get("113770000041271035")
				require.NoError(t, err)

				expected := &api.Tag{
					TagID:    "123",
					TagName:  "foobar",
					TagValue: "baz",
					TagColor: "#B7DA9E",
				}

				assert.Equal(t, expected, group)
			},
		},
		{
			Name:         "list tags",
			ExpectedVerb: "GET",
			ExpectedPath: "/tags",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_tags.json"),
			Fn: func(t *testing.T, c rest.Client) {
				groups, err := NewTags(c).List()
				require.NoError(t, err)

				expected := []*api.Tag{
					{
						TagID:    "123",
						TagName:  "foobar",
						TagValue: "baz",
						TagColor: "#B7DA9E",
					},
					{
						TagID:    "79123400003075053",
						TagName:  "foobar 1",
						TagValue: "baz 1",
						TagColor: "#B7DA9E",
					},
					{
						TagID:    "79730456703075223",
						TagName:  "foobar 2",
						TagValue: "baz 2",
						TagColor: "#B7DA9E",
					},
				}

				assert.Equal(t, expected, groups)
			},
		},
		{
			Name:         "update tags",
			ExpectedVerb: "PUT",
			ExpectedPath: "/tags/123",
			ExpectedBody: validation.Fixture(t, "requests/update_tag.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				tagUpdate := &api.Tag{
					TagID:    "123",
					TagName:  "foobar",
					TagValue: "baz",
					TagColor: "#B7DA9E",
				}

				_, err := NewTags(c).Update(tagUpdate)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete tags",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/tags/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewTags(c).Delete("123"))
			},
		},
	})
}
