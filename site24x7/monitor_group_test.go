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

func TestMonitorGroupCreate(t *testing.T) {
	d := monitorGroupTestResourceData(t)

	c := fake.NewClient()

	a := &api.MonitorGroup{
		DisplayName:            "foobar",
		Description:            "baz",
		DependencyResourceIDs:  []string{"234", "567"},
		DependencyResourceType: 2,
	}

	c.FakeMonitorGroups.On("Create", a).Return(a, nil).Once()

	require.NoError(t, monitorGroupCreate(d, c))

	c.FakeMonitorGroups.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := monitorGroupCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

// func TestMonitorGroupUpdate(t *testing.T) {
// 	d := monitorGroupTestResourceData(t)
// 	d.SetId("123")

// 	c := fake.NewClient()

// 	a := &api.MonitorGroup{
// 		GroupID:     "123",
// 		DisplayName: "foobar",
// 		Description: "baz",
// 	}

// 	c.FakeMonitorGroups.On("Update", a).Return(a, nil).Once()

// 	require.NoError(t, monitorGroupUpdate(d, c))

// 	c.FakeMonitorGroups.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

// 	err := monitorGroupUpdate(d, c)

// 	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
// }

func TestMonitorGroupRead(t *testing.T) {
	d := monitorGroupTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeMonitorGroups.On("Get", "123").Return(&api.MonitorGroup{}, nil).Once()

	require.NoError(t, monitorGroupRead(d, c))

	c.FakeMonitorGroups.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := monitorGroupRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestMonitorGroupDelete(t *testing.T) {
	d := monitorGroupTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeMonitorGroups.On("Delete", "123").Return(nil).Once()

	require.NoError(t, monitorGroupDelete(d, c))

	c.FakeMonitorGroups.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, monitorGroupDelete(d, c))
}

func TestMonitorGroupExists(t *testing.T) {
	d := monitorGroupTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeMonitorGroups.On("Get", "123").Return(&api.MonitorGroup{}, nil).Once()

	exists, err := monitorGroupExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeMonitorGroups.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = monitorGroupExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeMonitorGroups.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = monitorGroupExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func monitorGroupTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, MonitorGroupSchema, map[string]interface{}{
		"display_name":   "foobar",
		"description":    "baz",
		"selection_type": 2,
		"dependency_resource_ids": []interface{}{
			"234",
			"567",
		},
	})
}
