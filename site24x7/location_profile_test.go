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

func TestLocationProfileCreate(t *testing.T) {
	d := locationProfileTestResourceData(t)

	c := fake.NewClient()

	a := &api.LocationProfile{
		ProfileName:        "prof",
		SecondaryLocations: []string{"123"},
		PrimaryLocation:    "1",
	}

	c.FakeLocationProfiles.On("Create", a).Return(a, nil).Once()

	require.NoError(t, locationProfileCreate(d, c))

	c.FakeLocationProfiles.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := locationProfileCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestLocationProfileUpdate(t *testing.T) {
	d := locationProfileTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.LocationProfile{
		ProfileID:          "123",
		ProfileName:        "prof",
		SecondaryLocations: []string{"123"},
		PrimaryLocation:    "1",
	}

	c.FakeLocationProfiles.On("Update", a).Return(a, nil).Once()

	require.NoError(t, locationProfileUpdate(d, c))

	c.FakeLocationProfiles.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := locationProfileUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestLocationProfileRead(t *testing.T) {
	d := locationProfileTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeLocationProfiles.On("Get", "123").Return(&api.LocationProfile{}, nil).Once()

	require.NoError(t, locationProfileRead(d, c))

	c.FakeLocationProfiles.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := locationProfileRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestLocationProfileDelete(t *testing.T) {
	d := locationProfileTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeLocationProfiles.On("Delete", "123").Return(nil).Once()

	require.NoError(t, locationProfileDelete(d, c))

	c.FakeLocationProfiles.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, locationProfileDelete(d, c))
}

func TestLocationProfileExists(t *testing.T) {
	d := locationProfileTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeLocationProfiles.On("Get", "123").Return(&api.LocationProfile{}, nil).Once()

	exists, err := locationProfileExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeLocationProfiles.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = locationProfileExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeLocationProfiles.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = locationProfileExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func locationProfileTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, LocationProfileSchema, map[string]interface{}{
		"profile_name": "prof",
		"secondary_locations": []interface{}{
			"123",
		},
		"primary_location": "1",
	})
}
