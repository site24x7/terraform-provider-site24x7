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

func TestSubgroupCreate(t *testing.T) {
	d := subgroupTestResourceData(t)

	c := fake.NewClient()

	a := &api.Subgroup{
		DisplayName:          "foobar",
		Description:          "baz",
		TopGroupID:           "123",
		ParentGroupID:        "456",
		Type:                 2,
		HealthThresholdCount: 1,
		Monitors: []string{
			"726000000002460",
			"726000000002464",
		},
	}

	c.FakeSubgroups.On("Create", a).Return(a, nil).Once()

	require.NoError(t, subgroupCreate(d, c))

	c.FakeSubgroups.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := subgroupCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestSubgroupUpdate(t *testing.T) {
	d := subgroupTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.Subgroup{
		ID:                   "123",
		DisplayName:          "foobar",
		Description:          "baz",
		TopGroupID:           "123",
		ParentGroupID:        "456",
		Type:                 2,
		HealthThresholdCount: 1,
		Monitors: []string{
			"726000000002460",
			"726000000002464",
		},
	}

	c.FakeSubgroups.On("Update", a).Return(a, nil).Once()

	require.NoError(t, subgroupUpdate(d, c))

	c.FakeSubgroups.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := subgroupUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestSubgroupRead(t *testing.T) {
	d := subgroupTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeSubgroups.On("Get", "123").Return(&api.Subgroup{}, nil).Once()

	require.NoError(t, subgroupRead(d, c))

	c.FakeSubgroups.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := subgroupRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestSubgroupDelete(t *testing.T) {
	d := subgroupTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeSubgroups.On("Delete", "123").Return(nil).Once()

	require.NoError(t, subgroupDelete(d, c))

	c.FakeSubgroups.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, subgroupDelete(d, c))
}

func TestSubgroupExists(t *testing.T) {
	d := subgroupTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeSubgroups.On("Get", "123").Return(&api.Subgroup{}, nil).Once()

	exists, err := subgroupExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeSubgroups.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = subgroupExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeSubgroups.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = subgroupExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func subgroupTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, SubgroupSchema, map[string]interface{}{
		"display_name":           "foobar",
		"description":            "baz",
		"group_type":             2,
		"top_group_id":           "123",
		"parent_group_id":        "456",
		"health_threshold_count": 1,
		"monitors": []interface{}{
			"726000000002460",
			"726000000002464",
		},
	})
}
