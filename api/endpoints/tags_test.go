package endpoints

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTags(t *testing.T) {
	runTests(t, []*endpointTest{
		{
			name:         "create tags",
			expectedVerb: "POST",
			expectedPath: "/tags",
			expectedBody: fixture(t, "requests/create_tag.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
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
			name:         "get tags",
			expectedVerb: "GET",
			expectedPath: "/tags/113770000041271035",
			statusCode:   200,
			responseBody: fixture(t, "responses/get_tag.json"),
			fn: func(t *testing.T, c rest.Client) {
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
			name:         "list tags",
			expectedVerb: "GET",
			expectedPath: "/tags",
			statusCode:   200,
			responseBody: fixture(t, "responses/list_tags.json"),
			fn: func(t *testing.T, c rest.Client) {
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
			name:         "update tags",
			expectedVerb: "PUT",
			expectedPath: "/tags/123",
			expectedBody: fixture(t, "requests/update_tag.json"),
			statusCode:   200,
			responseBody: jsonAPIResponseBody(t, nil),
			fn: func(t *testing.T, c rest.Client) {
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
			name:         "delete tags",
			expectedVerb: "DELETE",
			expectedPath: "/tags/123",
			statusCode:   200,
			fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewTags(c).Delete("123"))
			},
		},
	})
}
