package site24x7

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/fake"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserGroupCreate(t *testing.T) {
	d := userGroupTestResourceData(t)

	c := fake.NewClient()

	a := &api.UserGroup{
		DisplayName:      "test_user_group",
		Users:            []string{"123", "456"},
		AttributeGroupID: "789",
		ProductID:        0,
	}

	c.FakeUserGroups.On("Create", a).Return(a, nil).Once()

	require.NoError(t, userGroupCreate(d, c))

	c.FakeUserGroups.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := userGroupCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestUserGroupUpdate(t *testing.T) {
	d := userGroupTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.UserGroup{
		UserGroupID:      "123",
		DisplayName:      "test_user_group",
		Users:            []string{"123", "456"},
		AttributeGroupID: "789",
		ProductID:        0,
	}

	c.FakeUserGroups.On("Update", a).Return(a, nil).Once()

	require.NoError(t, userGroupUpdate(d, c))

	c.FakeUserGroups.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := userGroupUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestUserGroupRead(t *testing.T) {
	d := userGroupTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeUserGroups.On("Get", "123").Return(&api.UserGroup{}, nil).Once()

	require.NoError(t, userGroupRead(d, c))

	c.FakeUserGroups.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := userGroupRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestUserGroupDelete(t *testing.T) {
	d := userGroupTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeUserGroups.On("Delete", "123").Return(nil).Once()

	require.NoError(t, userGroupDelete(d, c))

	c.FakeUserGroups.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, userGroupDelete(d, c))
}

func TestUserGroupExists(t *testing.T) {
	d := userGroupTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeUserGroups.On("Get", "123").Return(&api.UserGroup{}, nil).Once()

	exists, err := userGroupExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeUserGroups.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = userGroupExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeUserGroups.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = userGroupExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func userGroupTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, UserGroupSchema, map[string]interface{}{
		"display_name": "test_user_group",
		"users": []interface{}{
			"123",
			"456",
		},
		"attribute_group_id": "789",
		"product_id":         0,
	})
}
