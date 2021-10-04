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

func TestThresholdProfileCreate(t *testing.T) {
	d := thresholdProfileTestResourceData(t)

	c := fake.NewClient()

	a := &api.ThresholdProfile{
		ProfileName:            "threshold_profile_name",
		Type:                   "URL",
		ProfileType:            1,
		DownLocationThreshold:  1,
		WebsiteContentModified: false,
		//WebsiteContentChanges:  websiteContentChanges,
	}

	c.FakeThresholdProfiles.On("Create", a).Return(a, nil).Once()

	require.NoError(t, thresholdProfileCreate(d, c))

	c.FakeThresholdProfiles.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := thresholdProfileCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestThresholdProfileUpdate(t *testing.T) {
	d := thresholdProfileTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.ThresholdProfile{
		ProfileID:              "123",
		ProfileName:            "threshold_profile_name",
		Type:                   "URL",
		ProfileType:            1,
		DownLocationThreshold:  1,
		WebsiteContentModified: false,
		//WebsiteContentChanges:  websiteContentChanges,

	}

	c.FakeThresholdProfiles.On("Update", a).Return(a, nil).Once()

	require.NoError(t, thresholdProfileUpdate(d, c))

	c.FakeThresholdProfiles.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := thresholdProfileUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestThresholdProfileRead(t *testing.T) {
	d := thresholdProfileTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeThresholdProfiles.On("Get", "123").Return(&api.ThresholdProfile{}, nil).Once()

	require.NoError(t, thresholdProfileRead(d, c))

	c.FakeThresholdProfiles.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := thresholdProfileRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestThresholdProfileDelete(t *testing.T) {
	d := thresholdProfileTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeThresholdProfiles.On("Delete", "123").Return(nil).Once()

	require.NoError(t, thresholdProfileDelete(d, c))

	c.FakeThresholdProfiles.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, thresholdProfileDelete(d, c))
}

func TestThresholdProfileExists(t *testing.T) {
	d := thresholdProfileTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeThresholdProfiles.On("Get", "123").Return(&api.ThresholdProfile{}, nil).Once()

	exists, err := thresholdProfileExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeThresholdProfiles.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = thresholdProfileExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeThresholdProfiles.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = thresholdProfileExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func thresholdProfileTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, ThresholdProfileSchema, map[string]interface{}{
		"profile_name":             "threshold_profile_name",
		"type":                     "URL",
		"profile_type":             1,
		"down_location_threshold":  1,
		"website_content_modified": false,
	})
}
