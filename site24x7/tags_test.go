package site24x7

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTagCreate(t *testing.T) {
	d := tagTestResourceData(t)

	c := fake.NewClient()

	a := &api.Tag{
		TagName:  "foobar",
		TagValue: "baz",
		TagColor: "#B7DA9E",
		TagType:  1,
	}

	c.FakeTags.On("Create", a).Return(a, nil).Once()

	require.NoError(t, tagCreate(d, c))

	c.FakeTags.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := tagCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestTagUpdate(t *testing.T) {
	d := tagTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.Tag{
		TagID:    "123",
		TagName:  "foobar",
		TagValue: "baz",
		TagColor: "#B7DA9E",
		TagType:  1,
	}

	c.FakeTags.On("Update", a).Return(a, nil).Once()

	require.NoError(t, tagUpdate(d, c))

	c.FakeTags.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := tagUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestTagRead(t *testing.T) {
	d := tagTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeTags.On("Get", "123").Return(&api.Tag{}, nil).Once()

	require.NoError(t, tagRead(d, c))

	c.FakeTags.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := tagRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestTagDelete(t *testing.T) {
	d := tagTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeTags.On("Delete", "123").Return(nil).Once()

	require.NoError(t, tagDelete(d, c))

	c.FakeTags.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, tagDelete(d, c))
}

func TestTagExists(t *testing.T) {
	d := tagTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeTags.On("Get", "123").Return(&api.Tag{}, nil).Once()

	exists, err := tagExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeTags.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = tagExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeTags.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = tagExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func tagTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, TagSchema, map[string]interface{}{
		"tag_name":  "foobar",
		"tag_value": "baz",
		"tag_color": "#B7DA9E",
	})
}
