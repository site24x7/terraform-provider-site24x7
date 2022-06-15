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

func TestScheduleMaintenanceCreate(t *testing.T) {
	d := scheduleMaintenanceTestResourceData(t)

	c := fake.NewClient()

	a := &api.ScheduleMaintenance{
		DisplayName:       "Schedule Maintenance",
		Description:       "Maintenance Window",
		MaintenanceType:   3,
		StartDate:         "2022-06-02",
		EndDate:           "2022-06-02",
		StartTime:         "19:41",
		EndTime:           "20:44",
		SelectionType:     2,
		Monitors:          []string{"123", "456"},
		PerformMonitoring: true,
	}

	c.FakeScheduleMaintenance.On("Create", a).Return(a, nil).Once()

	require.NoError(t, scheduleMaintenanceCreate(d, c))

	c.FakeScheduleMaintenance.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := scheduleMaintenanceCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestScheduleMaintenanceUpdate(t *testing.T) {
	d := scheduleMaintenanceTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.ScheduleMaintenance{
		MaintenanceID:     "123",
		DisplayName:       "Schedule Maintenance",
		Description:       "Maintenance Window",
		MaintenanceType:   3,
		StartDate:         "2022-06-02",
		EndDate:           "2022-06-02",
		StartTime:         "19:41",
		EndTime:           "20:44",
		SelectionType:     2,
		Monitors:          []string{"123", "456"},
		PerformMonitoring: true,
	}

	c.FakeScheduleMaintenance.On("Update", a).Return(a, nil).Once()

	require.NoError(t, scheduleMaintenanceUpdate(d, c))

	c.FakeScheduleMaintenance.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := scheduleMaintenanceUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestScheduleMaintenanceRead(t *testing.T) {
	d := scheduleMaintenanceTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeScheduleMaintenance.On("Get", "123").Return(&api.ScheduleMaintenance{}, nil).Once()

	require.NoError(t, scheduleMaintenanceRead(d, c))

	c.FakeScheduleMaintenance.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := scheduleMaintenanceRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestScheduleMaintenanceDelete(t *testing.T) {
	d := scheduleMaintenanceTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeScheduleMaintenance.On("Delete", "123").Return(nil).Once()

	require.NoError(t, scheduleMaintenanceDelete(d, c))

	c.FakeScheduleMaintenance.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, scheduleMaintenanceDelete(d, c))
}

func TestScheduleMaintenanceExists(t *testing.T) {
	d := scheduleMaintenanceTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeScheduleMaintenance.On("Get", "123").Return(&api.ScheduleMaintenance{}, nil).Once()

	exists, err := scheduleMaintenanceExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeScheduleMaintenance.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = scheduleMaintenanceExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeScheduleMaintenance.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = scheduleMaintenanceExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func scheduleMaintenanceTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, ScheduleMaintenanceSchema, map[string]interface{}{
		"display_name":     "Schedule Maintenance",
		"description":      "Maintenance Window",
		"start_date":       "2022-06-02",
		"end_date":         "2022-06-02",
		"start_time":       "19:41",
		"end_time":         "20:44",
		"selection_type":   2,
		"maintenance_type": 3,
		"monitors": []interface{}{
			"123",
			"456",
		},
		"perform_monitoring": true,
	})
}
