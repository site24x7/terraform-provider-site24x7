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

func TestThresholdProfileCreate(t *testing.T) {
	d := thresholdProfileTestResourceData(t)

	c := fake.NewClient()

	a := &api.ThresholdProfile{
		ProfileName:            "threshold_profile_name",
		Type:                   "URL",
		ProfileType:            1,
		DownLocationThreshold:  1,
		WebsiteContentModified: true,
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
		WebsiteContentModified: true,
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
		"website_content_modified": true,
	})
}

func TestServerThresholdProfileCreate(t *testing.T) {
	d := serverThresholdProfileTestResourceData(t)

	c := fake.NewClient()

	a := &api.ThresholdProfile{
		ProfileName: "server_threshold_profile",
		Type:        "SERVER",
		ProfileType: 1,
		CpuThreshold: []map[string]interface{}{
			{
				"severity":            "2",
				"comparison_operator": "1",
				"value":               "80",
				"strategy":            "1",
				"polls_check":         "5",
			},
		},
		MemoryThreshold: []map[string]interface{}{
			{
				"severity":            "2",
				"comparison_operator": "1",
				"value":               "90",
				"strategy":            "1",
				"polls_check":         "5",
			},
		},
		ServerResourceDownAlert: map[string]interface{}{
			"severity": "2",
			"value":    "true",
		},
	}

	c.FakeThresholdProfiles.On("Create", a).Return(a, nil).Once()

	require.NoError(t, thresholdProfileCreate(d, c))

	c.FakeThresholdProfiles.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := thresholdProfileCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestServerThresholdProfileUpdate(t *testing.T) {
	d := serverThresholdProfileTestResourceData(t)
	d.SetId("456")

	c := fake.NewClient()

	a := &api.ThresholdProfile{
		ProfileID:   "456",
		ProfileName: "server_threshold_profile",
		Type:        "SERVER",
		ProfileType: 1,
		CpuThreshold: []map[string]interface{}{
			{
				"severity":            "2",
				"comparison_operator": "1",
				"value":               "80",
				"strategy":            "1",
				"polls_check":         "5",
			},
		},
		MemoryThreshold: []map[string]interface{}{
			{
				"severity":            "2",
				"comparison_operator": "1",
				"value":               "90",
				"strategy":            "1",
				"polls_check":         "5",
			},
		},
		ServerResourceDownAlert: map[string]interface{}{
			"severity": "2",
			"value":    "true",
		},
	}

	c.FakeThresholdProfiles.On("Update", a).Return(a, nil).Once()

	require.NoError(t, thresholdProfileUpdate(d, c))

	c.FakeThresholdProfiles.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := thresholdProfileUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestServerThresholdProfileRead(t *testing.T) {
	d := serverThresholdProfileTestResourceData(t)
	d.SetId("456")

	c := fake.NewClient()

	serverProfile := &api.ThresholdProfile{
		ProfileID:   "456",
		ProfileName: "server_threshold_profile",
		Type:        "SERVER",
		ProfileType: 1,
		CpuThreshold: []map[string]interface{}{
			{
				"severity":            float64(2),
				"comparison_operator": float64(1),
				"value":               float64(80),
				"strategy":            float64(1),
				"polls_check":         float64(5),
			},
		},
		ServerResourceDownAlert: map[string]interface{}{
			"severity": float64(2),
			"value":    true,
		},
	}

	c.FakeThresholdProfiles.On("Get", "456").Return(serverProfile, nil).Once()

	require.NoError(t, thresholdProfileRead(d, c))

	assert.Equal(t, "server_threshold_profile", d.Get("profile_name"))
	assert.Equal(t, "SERVER", d.Get("type"))
}

func serverThresholdProfileTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, ThresholdProfileSchema, map[string]interface{}{
		"profile_name": "server_threshold_profile",
		"type":         "SERVER",
		"profile_type": 1,
		"server_resource_down_alert": map[string]interface{}{
			"severity": "2",
			"value":    "true",
		},
		"cpu_trouble_threshold": map[string]interface{}{
			"severity":            "2",
			"comparison_operator": "1",
			"value":               "80",
			"strategy":            "1",
			"polls_check":         "5",
		},
		"memory_trouble_threshold": map[string]interface{}{
			"severity":            "2",
			"comparison_operator": "1",
			"value":               "90",
			"strategy":            "1",
			"polls_check":         "5",
		},
	})
}
