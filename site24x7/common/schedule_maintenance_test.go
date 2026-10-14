package common

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
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
		TimeZone:          "PST",
		EndDate:           "2022-06-02",
		StartTime:         "19:41",
		EndTime:           "20:44",
		SelectionType:     2,
		Monitors:          []string{"123", "456"},
		PerformMonitoring: true,
	}

	created := &api.ScheduleMaintenance{MaintenanceID: "123"}

	c.FakeScheduleMaintenance.On("Create", a).Return(created, nil).Once()
	c.FakeScheduleMaintenance.On("Get", "123").Return(a, nil).Once()

	require.NoError(t, scheduleMaintenanceCreate(d, c))
	assert.Equal(t, "123", d.Id())

	d = scheduleMaintenanceTestResourceData(t)

	c.FakeScheduleMaintenance.On("Create", a).Return(nil, apierrors.NewStatusError(500, "error")).Once()

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
		TimeZone:          "PST",
		StartTime:         "19:41",
		EndTime:           "20:44",
		SelectionType:     2,
		Monitors:          []string{"123", "456"},
		PerformMonitoring: true,
	}

	c.FakeScheduleMaintenance.On("Update", a).Return(a, nil).Once()
	c.FakeScheduleMaintenance.On("Get", "123").Return(a, nil).Once()

	require.NoError(t, scheduleMaintenanceUpdate(d, c))

	c.FakeScheduleMaintenance.On("Update", a).Return(nil, apierrors.NewStatusError(500, "error")).Once()

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

// Read must replace every attribute, so a change of type or an import leaves
// nothing stale, and normalise what the API leaves out.
func TestScheduleMaintenanceReadSetsEveryAttribute(t *testing.T) {
	d := scheduleMaintenanceTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeScheduleMaintenance.On("Get", "123").Return(&api.ScheduleMaintenance{
		MaintenanceID:      "123",
		DisplayName:        "Monthly By day",
		MaintenanceType:    api.MaintenanceTypeMonthlyByDay,
		TimeZone:           "IST",
		StartWeek:          2,
		StartDay:           2,
		StartAfter:         2,
		StartTime:          "12:30",
		Duration:           float64(30), // as decoded from JSON
		MaintenanceStartOn: "2022-08-03",
		SelectionType:      maintenanceSelectionMonitorGroups,
		MonitorGroups:      []string{"789"},
		PerformMonitoring:  true,
	}, nil).Once()

	require.NoError(t, scheduleMaintenanceRead(d, c))

	assert.Equal(t, api.MaintenanceTypeMonthlyByDay, d.Get("maintenance_type"))
	assert.Equal(t, 2, d.Get("start_week"))
	assert.Equal(t, 2, d.Get("start_day"))
	assert.Equal(t, 2, d.Get("start_after"))
	assert.Equal(t, 30, d.Get("duration"))
	assert.Equal(t, "2022-08-03", d.Get("maintenance_start_on"))
	assert.Equal(t, "IST", d.Get("time_zone"))
	assert.Equal(t, 1, d.Get("selection_type"))
	assert.Equal(t, []interface{}{"789"}, d.Get("monitor_groups"))

	// Values of the previous Once configuration are cleared.
	assert.Equal(t, "", d.Get("start_date"))
	assert.Equal(t, "", d.Get("end_date"))
	assert.Equal(t, "", d.Get("end_time"))
	assert.Empty(t, d.Get("monitors"))

	// Absent values fall back to the API defaults.
	assert.Equal(t, 1, d.Get("execute_every"))
	assert.Equal(t, true, d.Get("subgroup_monitors"))
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

// Each maintenance type sends only the attributes the API reads for it.
func TestResourceDataToScheduleMaintenancePerType(t *testing.T) {
	boolPtr := func(b bool) *bool { return &b }

	tests := []struct {
		name     string
		config   map[string]interface{}
		expected *api.ScheduleMaintenance
	}{
		{
			name: "daily ending after 5 runs",
			config: map[string]interface{}{
				"maintenance_type":            api.MaintenanceTypeDaily,
				"end_time":                    "02:00",
				"maintenance_start_on":        "2026-10-01",
				"maintenance_end_type":        1,
				"maintenance_end_after_times": 5,
			},
			expected: &api.ScheduleMaintenance{
				MaintenanceType:          api.MaintenanceTypeDaily,
				EndTime:                  "02:00",
				MaintenanceStartOn:       "2026-10-01",
				MaintenanceEndType:       1,
				MaintenanceEndAfterTimes: 5,
			},
		},
		{
			name: "weekly by time from Sunday, bi-weekly",
			config: map[string]interface{}{
				"maintenance_type": api.MaintenanceTypeWeeklyByTime,
				"start_day":        0,
				"end_day":          1,
				"end_time":         "02:00",
				"execute_every":    2,
			},
			expected: &api.ScheduleMaintenance{
				MaintenanceType: api.MaintenanceTypeWeeklyByTime,
				StartDay:        0,
				EndDay:          1,
				EndTime:         "02:00",
				ExecuteEvery:    2,
			},
		},
		{
			name: "weekly by day ending on a date",
			config: map[string]interface{}{
				"maintenance_type":     api.MaintenanceTypeWeeklyByDay,
				"week_days":            []interface{}{3, 5},
				"duration":             30,
				"maintenance_end_type": 2,
				"maintenance_end_on":   "2027-01-01",
			},
			expected: &api.ScheduleMaintenance{
				MaintenanceType:    api.MaintenanceTypeWeeklyByDay,
				WeekDays:           []int{3, 5},
				Duration:           30,
				ExecuteEvery:       1,
				MaintenanceEndType: 2,
				MaintenanceEndOn:   "2027-01-01",
			},
		},
		{
			name: "monthly by date",
			config: map[string]interface{}{
				"maintenance_type":   api.MaintenanceTypeMonthlyByDate,
				"monthly_start_date": 15,
				"duration":           120,
			},
			expected: &api.ScheduleMaintenance{
				MaintenanceType:  api.MaintenanceTypeMonthlyByDate,
				MonthlyStartDate: 15,
				Duration:         120,
			},
		},
		{
			name: "monthly by day for monitor groups without subgroups",
			config: map[string]interface{}{
				"maintenance_type":  api.MaintenanceTypeMonthlyByDay,
				"start_week":        2,
				"start_day":         2,
				"start_after":       2,
				"duration":          30,
				"selection_type":    1,
				"monitor_groups":    []interface{}{"789"},
				"subgroup_monitors": false,
			},
			expected: &api.ScheduleMaintenance{
				MaintenanceType:  api.MaintenanceTypeMonthlyByDay,
				StartWeek:        2,
				StartDay:         2,
				StartAfter:       2,
				Duration:         30,
				SelectionType:    maintenanceSelectionMonitorGroups,
				MonitorGroups:    []string{"789"},
				SubgroupMonitors: boolPtr(false),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config := map[string]interface{}{
				"display_name": "Maintenance",
				"start_time":   "01:00",
				"monitors":     []interface{}{"123"},
			}
			for k, v := range test.config {
				config[k] = v
			}
			if test.config["selection_type"] != nil {
				delete(config, "monitors")
			}

			d := schema.TestResourceDataRaw(t, ScheduleMaintenanceSchema, config)

			expected := test.expected
			expected.DisplayName = "Maintenance"
			expected.StartTime = "01:00"
			expected.PerformMonitoring = true
			if expected.SelectionType == 0 {
				expected.SelectionType = maintenanceSelectionMonitors
				expected.Monitors = []string{"123"}
			}

			assert.Equal(t, expected, resourceDataToScheduleMaintenance(d))
		})
	}
}

// unknownValue is how the SDK marks a value not known until apply.
const unknownValue = "74D93920-ED26-11E3-AC10-0800200C9A66"

func TestScheduleMaintenanceCustomizeDiff(t *testing.T) {
	tests := []struct {
		name   string
		config map[string]interface{}
		errors []string
	}{
		{
			name: "valid once",
			config: map[string]interface{}{
				"maintenance_type": 3, "start_date": "2026-10-01", "end_date": "2026-10-01", "end_time": "02:00",
			},
		},
		{
			name:   "once without dates",
			config: map[string]interface{}{"maintenance_type": 3, "end_time": "02:00"},
			errors: []string{
				"start_date is required when maintenance_type is 3 (Once)",
				"end_date is required when maintenance_type is 3 (Once)",
			},
		},
		{
			name: "once with a recurrence end",
			config: map[string]interface{}{
				"maintenance_type": 3, "start_date": "2026-10-01", "end_date": "2026-10-01", "end_time": "02:00",
				"maintenance_end_type": 2, "maintenance_end_on": "2027-01-01",
			},
			errors: []string{"maintenance_end_type is not used when maintenance_type is 3 (Once)"},
		},
		{
			name: "weekly by time starting Sunday",
			config: map[string]interface{}{
				"maintenance_type": 2, "start_day": 0, "end_day": 0, "end_time": "02:00",
			},
		},
		{
			name:   "weekly by time without days",
			config: map[string]interface{}{"maintenance_type": 2, "end_time": "02:00"},
			errors: []string{
				"start_day is required when maintenance_type is 2 (Weekly (By Time))",
				"end_day is required when maintenance_type is 2 (Weekly (By Time))",
			},
		},
		{
			name:   "weekly by day",
			config: map[string]interface{}{"maintenance_type": 8, "week_days": []interface{}{3, 5}, "duration": 30, "execute_every": 2},
		},
		{
			name:   "weekly by day with an end time",
			config: map[string]interface{}{"maintenance_type": 8, "duration": 30, "end_time": "02:00"},
			errors: []string{"end_time is not used when maintenance_type is 8 (Weekly (By Day))"},
		},
		{
			name:   "monthly by date without duration",
			config: map[string]interface{}{"maintenance_type": 5, "monthly_start_date": 15},
			errors: []string{"duration is required when maintenance_type is 5 (Monthly (By Date))"},
		},
		{
			name:   "monthly by date with a weekly interval",
			config: map[string]interface{}{"maintenance_type": 5, "monthly_start_date": 15, "duration": 30, "execute_every": 2},
			errors: []string{"execute_every is not used when maintenance_type is 5 (Monthly (By Date))"},
		},
		{
			name:   "monthly by day",
			config: map[string]interface{}{"maintenance_type": 6, "start_week": 5, "start_day": 2, "duration": 30},
		},
		{
			name:   "monthly by day without week or day",
			config: map[string]interface{}{"maintenance_type": 6, "duration": 30},
			errors: []string{
				"start_day is required when maintenance_type is 6 (Monthly (By Day))",
				"start_week is required when maintenance_type is 6 (Monthly (By Day))",
			},
		},
		{
			name:   "end after times missing",
			config: map[string]interface{}{"maintenance_type": 1, "end_time": "02:00", "maintenance_end_type": 1},
			errors: []string{"maintenance_end_after_times is required when maintenance_end_type is 1"},
		},
		{
			name:   "end date without the matching end type",
			config: map[string]interface{}{"maintenance_type": 1, "end_time": "02:00", "maintenance_end_on": "2027-01-01"},
			errors: []string{"maintenance_end_on is only used when maintenance_end_type is 2"},
		},
		{
			name: "monitor groups selection with monitors",
			config: map[string]interface{}{
				"maintenance_type": 1, "end_time": "02:00", "selection_type": 1, "monitors": []interface{}{"123"},
			},
			errors: []string{"monitors is not used when selection_type is 1"},
		},
		{
			name: "monitors list unknown until apply",
			config: map[string]interface{}{
				"maintenance_type": 1, "end_time": "02:00", "monitors": unknownValue,
			},
		},
		{
			name: "monitor ID unknown until apply",
			config: map[string]interface{}{
				"maintenance_type": 1, "end_time": "02:00", "monitors": []interface{}{unknownValue},
			},
		},
		{
			name: "week_days unknown until apply",
			config: map[string]interface{}{
				"maintenance_type": 8, "duration": 30, "week_days": unknownValue,
			},
		},
		{
			name: "subgroups off without monitor groups",
			config: map[string]interface{}{
				"maintenance_type": 1, "end_time": "02:00", "subgroup_monitors": false,
			},
			errors: []string{"subgroup_monitors is only used when selection_type is 1"},
		},
	}

	resource := ResourceSite24x7ScheduleMaintenance()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config := map[string]interface{}{
				"display_name": "Maintenance",
				"start_time":   "01:00",
				"monitors":     []interface{}{"123"},
			}
			for k, v := range test.config {
				config[k] = v
			}

			_, err := resource.Diff(nil, terraform.NewResourceConfigRaw(config), nil)

			if len(test.errors) == 0 {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			for _, expected := range test.errors {
				assert.Contains(t, err.Error(), expected)
			}
		})
	}
}

// After Read stores what the API returns, planning the same configuration
// again must show no changes and pass validation.
func TestScheduleMaintenancePlanIsCleanAfterRead(t *testing.T) {
	tests := []struct {
		name     string
		config   map[string]interface{}
		response *api.ScheduleMaintenance
	}{
		{
			name: "once",
			config: map[string]interface{}{
				"maintenance_type": 3, "start_date": "2026-10-01", "end_date": "2026-10-01", "end_time": "02:00",
			},
			// The API derives duration from start_time and end_time.
			response: &api.ScheduleMaintenance{
				MaintenanceType: 3, StartDate: "2026-10-01", EndDate: "2026-10-01", EndTime: "02:00", Duration: float64(60),
			},
		},
		{
			name:   "daily",
			config: map[string]interface{}{"maintenance_type": 1, "end_time": "02:00"},
			response: &api.ScheduleMaintenance{
				MaintenanceType: 1, EndTime: "02:00", Duration: float64(60), MaintenanceStartOn: "2026-09-30",
			},
		},
		{
			name:   "weekly by time on Sunday",
			config: map[string]interface{}{"maintenance_type": 2, "start_day": 0, "end_day": 0, "end_time": "02:00"},
			response: &api.ScheduleMaintenance{
				MaintenanceType: 2, StartDay: 0, EndDay: 0, EndTime: "02:00", ExecuteEvery: 1, MaintenanceStartOn: "2026-09-30",
			},
		},
		{
			name:   "weekly by day",
			config: map[string]interface{}{"maintenance_type": 8, "week_days": []interface{}{3, 5}, "duration": 30, "execute_every": 2},
			response: &api.ScheduleMaintenance{
				MaintenanceType: 8, WeekDays: []int{3, 5}, Duration: float64(30), ExecuteEvery: 2, MaintenanceStartOn: "2026-09-30",
			},
		},
		{
			name:   "monthly by date",
			config: map[string]interface{}{"maintenance_type": 5, "monthly_start_date": 15, "duration": 60},
			response: &api.ScheduleMaintenance{
				MaintenanceType: 5, MonthlyStartDate: 15, Duration: float64(60), MaintenanceStartOn: "2026-09-30",
			},
		},
		{
			name: "monthly by day ending after 3 runs",
			config: map[string]interface{}{
				"maintenance_type": 6, "start_week": 2, "start_day": 2, "duration": 30,
				"maintenance_end_type": 1, "maintenance_end_after_times": 3,
			},
			response: &api.ScheduleMaintenance{
				MaintenanceType: 6, StartWeek: 2, StartDay: 2, Duration: float64(30), MaintenanceStartOn: "2026-09-30",
				MaintenanceEndType: 1, MaintenanceEndAfterTimes: 3,
			},
		},
	}

	resource := ResourceSite24x7ScheduleMaintenance()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config := map[string]interface{}{
				"display_name": "Maintenance",
				"start_time":   "01:00",
				"monitors":     []interface{}{"123"},
			}
			for k, v := range test.config {
				config[k] = v
			}

			response := test.response
			response.MaintenanceID = "123"
			response.DisplayName = "Maintenance"
			response.StartTime = "01:00"
			response.TimeZone = "IST"
			response.SelectionType = maintenanceSelectionMonitors
			response.Monitors = []string{"123"}
			response.PerformMonitoring = true

			d := resource.TestResourceData()
			d.SetId("123")
			updateScheduleMaintenanceResourceData(d, response)

			diff, err := resource.Diff(d.State(), terraform.NewResourceConfigRaw(config), nil)

			require.NoError(t, err)
			if diff != nil {
				assert.Empty(t, diff.Attributes)
			}
		})
	}
}

func TestScheduleMaintenanceSchemaValidation(t *testing.T) {
	resource := ResourceSite24x7ScheduleMaintenance()

	tests := []struct {
		name   string
		config map[string]interface{}
	}{
		{"yearly is not supported", map[string]interface{}{"maintenance_type": 9}},
		{"day of week out of range", map[string]interface{}{"maintenance_type": 2, "start_day": 7}},
		{"week day out of range", map[string]interface{}{"maintenance_type": 8, "week_days": []interface{}{7}}},
		{"execute_every out of range", map[string]interface{}{"maintenance_type": 8, "execute_every": 5}},
		{"start_after out of range", map[string]interface{}{"maintenance_type": 6, "start_after": 26}},
		{"start_week out of range", map[string]interface{}{"maintenance_type": 6, "start_week": 6}},
		{"date format", map[string]interface{}{"maintenance_type": 3, "start_date": "01-10-2026"}},
		{"all monitors selection", map[string]interface{}{"maintenance_type": 1, "selection_type": 0}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config := map[string]interface{}{
				"display_name": "Maintenance",
				"start_time":   "01:00",
			}
			for k, v := range test.config {
				config[k] = v
			}

			_, errs := resource.Validate(terraform.NewResourceConfigRaw(config))

			assert.NotEmpty(t, errs)
		})
	}
}

func scheduleMaintenanceTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, ScheduleMaintenanceSchema, map[string]interface{}{
		"display_name":     "Schedule Maintenance",
		"description":      "Maintenance Window",
		"start_date":       "2022-06-02",
		"end_date":         "2022-06-02",
		"time_zone":        "PST",
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
