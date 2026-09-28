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

// businessHourRequest is what the resource builds from businessHourTestResourceData
// before it has an ID, i.e. what a create sends.
func businessHourRequest() *api.BusinessHour {
	return &api.BusinessHour{
		DisplayName: "Business Hour",
		Description: "Test description",
		TimeConfig: []api.TimeSlot{
			{Day: 1, StartTime: "09:00", EndTime: "18:00"},
		},
	}
}

func TestBusinessHourCreate(t *testing.T) {
	d := businessHourTestResourceData(t)
	c := fake.NewClient()

	created := businessHourRequest()
	created.ID = "123"

	c.FakeBusinesshour.On("Create", businessHourRequest()).Return(created, nil).Once()

	require.NoError(t, businessHourCreate(d, c))
	assert.Equal(t, "123", d.Id())

	// A fresh ResourceData, because the successful create above set an ID that
	// would otherwise be carried into the next request.
	failing := businessHourTestResourceData(t)

	c.FakeBusinesshour.On("Create", businessHourRequest()).
		Return(nil, apierrors.NewStatusError(500, "error")).Once()

	// businessHourCreate wraps the API error with %w, so the raw StatusError is
	// no longer comparable by value.
	assert.EqualError(t, businessHourCreate(failing, c), "failed to create business hour: error")
}

func TestBusinessHourUpdate(t *testing.T) {
	d := businessHourTestResourceData(t)
	d.SetId("123")
	c := fake.NewClient()

	a := businessHourRequest()
	a.ID = "123"

	c.FakeBusinesshour.On("Update", a).Return(a, nil).Once()
	require.NoError(t, businessHourUpdate(d, c))

	// The real client returns a non-nil (empty) struct alongside the error, and
	// this test mirrors that. It matters: businessHourUpdate reassigns its local
	// variable to the returned value and then reads businessHour.ID when
	// building the error, so a client returning (nil, err) would panic, and the
	// ID in the message comes from the empty response rather than the request.
	c.FakeBusinesshour.On("Update", a).
		Return(&api.BusinessHour{}, apierrors.NewStatusError(500, "error")).Once()

	updateErr := businessHourUpdate(d, c)
	require.Error(t, updateErr)
	assert.Contains(t, updateErr.Error(), "failed to update business hour")
}

func TestBusinessHourRead(t *testing.T) {
	d := businessHourTestResourceData(t)
	d.SetId("123")
	c := fake.NewClient()

	c.FakeBusinesshour.On("Get", "123").Return(&api.BusinessHour{ID: "123"}, nil).Once()
	require.NoError(t, businessHourRead(d, c))

	c.FakeBusinesshour.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	assert.EqualError(t, businessHourRead(d, c), "failed to read business hour with ID 123: error")
}

// A business hour deleted outside Terraform is dropped from state rather than
// failing the refresh.
func TestBusinessHourReadDropsDeletedBusinessHour(t *testing.T) {
	d := businessHourTestResourceData(t)
	d.SetId("123")
	c := fake.NewClient()

	c.FakeBusinesshour.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, businessHourRead(d, c))
	assert.Equal(t, "", d.Id())
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

	c.FakeBusinesshour.On("Get", "123").Return(&api.BusinessHour{ID: "123"}, nil).Once()

	exists, err := businessHourExists(d, c)
	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeBusinesshour.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = businessHourExists(d, c)
	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeBusinesshour.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = businessHourExists(d, c)
	assert.EqualError(t, err, "failed to check existence of business hour with ID 123: error")
	assert.False(t, exists)
}

// The keys here are the ones BusinessHourSchema actually declares. The previous
// version used timezone/work_hours/weekdays, which the schema has never had, so
// every field except display_name arrived empty.
func businessHourTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, BusinessHourSchema, map[string]interface{}{
		"display_name": "Business Hour",
		"description":  "Test description",
		"time_config": []interface{}{
			map[string]interface{}{
				"day":        1,
				"start_time": "09:00",
				"end_time":   "18:00",
			},
		},
	})
}
