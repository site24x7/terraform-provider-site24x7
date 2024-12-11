package common

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBusinessHourCreate(t *testing.T) {
	d := businessHourTestResourceData(t)
	c := fake.NewClient()

	a := &api.BusinessHour{
		ID:          "123",
		DisplayName: "Business Hour",
		Description: "Test description",
		TimeConfig: []api.TimeSlot{
			{Day: 1, StartTime: "09:00", EndTime: "18:00"},
		},
	}

	c.FakeBusinesshour.On("Create", a).Return(a, nil).Once()
	require.NoError(t, businessHourCreate(d, c))

	c.FakeBusinesshour.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()
	err := businessHourCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestBusinessHourUpdate(t *testing.T) {
	d := businessHourTestResourceData(t)
	d.SetId("123")
	c := fake.NewClient()

	a := &api.BusinessHour{
		ID:          "123",
		DisplayName: "Business Hour",
		Description: "Test description",
		TimeConfig: []api.TimeSlot{
			{Day: 1, StartTime: "09:00", EndTime: "18:00"},
		},
	}

	c.FakeBusinesshour.On("Update", a).Return(a, nil).Once()
	require.NoError(t, businessHourUpdate(d, c))

	c.FakeBusinesshour.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()
	err := businessHourUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestBusinessHourRead(t *testing.T) {
	d := businessHourTestResourceData(t)
	d.SetId("123")
	c := fake.NewClient()

	c.FakeBusinesshour.On("Get", "123").Return(&api.BusinessHour{}, nil).Once()
	require.NoError(t, businessHourRead(d, c))

	c.FakeBusinesshour.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()
	err := businessHourRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestBusinessHourDelete(t *testing.T) {
	d := businessHourTestResourceData(t)
	d.SetId("123")
	c := fake.NewClient()

	c.FakeBusinesshour.On("Delete", "123").Return(nil).Once()
	require.NoError(t, businessHourDelete(d, c))

	c.FakeBusinesshour.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()
	require.NoError(t, businessHourDelete(d, c))
}

func TestBusinessHourExists(t *testing.T) {
	d := businessHourTestResourceData(t)
	d.SetId("123")
	c := fake.NewClient()

	c.FakeBusinesshour.On("Get", "123").Return(&api.BusinessHour{}, nil).Once()

	exists, err := businessHourExists(d, c)
	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeBusinesshour.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = businessHourExists(d, c)
	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeBusinesshour.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = businessHourExists(d, c)
	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func businessHourTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, BusinessHourSchema, map[string]interface{}{
		"display_name": "Business Hour",
		"timezone":     "PST",
		"work_hours": []interface{}{
			"09:00-18:00",
		},
		"weekdays": []interface{}{1, 2, 3, 4, 5},
	})
}
