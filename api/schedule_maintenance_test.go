package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The API sends "" for the attributes a maintenance type does not use.
func TestScheduleMaintenanceUnmarshalJSONAcceptsEmptyStrings(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		expected *ScheduleMaintenance
	}{
		{
			name: "once",
			body: `{"maintenance_id":"1","display_name":"Once","maintenance_type":3,"start_date":"2026-10-20",
				"end_date":"2026-10-20","start_time":"19:30","end_time":"21:00","start_day":"","end_day":"",
				"duration":"","week_days":"","execute_every":"","start_week":"","start_after":"",
				"monthly_start_date":"","maintenance_end_type":"","maintenance_end_after_times":"",
				"selection_type":2,"monitors":["123"],"subgroup_monitors":"","perform_monitoring":true}`,
			expected: &ScheduleMaintenance{
				MaintenanceID: "1", DisplayName: "Once", MaintenanceType: MaintenanceTypeOnce,
				StartDate: "2026-10-20", EndDate: "2026-10-20", StartTime: "19:30", EndTime: "21:00",
				Duration: "", SelectionType: 2, Monitors: []string{"123"}, PerformMonitoring: true,
			},
		},
		{
			name: "weekly by time on Sunday with numbers as strings",
			body: `{"maintenance_id":"2","maintenance_type":"2","start_day":"0","end_day":0,"start_time":"22:00",
				"end_time":"04:00","execute_every":"2","start_week":"","maintenance_end_type":2,
				"maintenance_end_on":"2027-03-31","selection_type":"2","perform_monitoring":"true"}`,
			expected: &ScheduleMaintenance{
				MaintenanceID: "2", MaintenanceType: MaintenanceTypeWeeklyByTime, StartDay: 0, EndDay: 0,
				StartTime: "22:00", EndTime: "04:00", ExecuteEvery: 2, MaintenanceEndType: 2,
				MaintenanceEndOn: "2027-03-31", SelectionType: 2, PerformMonitoring: true,
			},
		},
		{
			name: "weekly by day",
			body: `{"maintenance_id":"3","maintenance_type":8,"week_days":[3,"5"],"start_time":"14:30",
				"duration":30,"execute_every":2,"end_day":"","start_day":"","selection_type":1,
				"monitor_groups":["9"],"subgroup_monitors":false,"perform_monitoring":false}`,
			expected: &ScheduleMaintenance{
				MaintenanceID: "3", MaintenanceType: MaintenanceTypeWeeklyByDay, WeekDays: []int{3, 5},
				StartTime: "14:30", Duration: float64(30), ExecuteEvery: 2, SelectionType: 1,
				MonitorGroups: []string{"9"}, SubgroupMonitors: boolPointer(false),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := &ScheduleMaintenance{}
			require.NoError(t, json.Unmarshal([]byte(test.body), got))
			assert.Equal(t, test.expected, got)
		})
	}
}

func TestScheduleMaintenanceUnmarshalJSONRejectsNonNumbers(t *testing.T) {
	err := json.Unmarshal([]byte(`{"start_day":"Sunday"}`), &ScheduleMaintenance{})
	assert.Error(t, err)
}

// A list response goes through the same decoder.
func TestScheduleMaintenanceListUnmarshalJSONAcceptsEmptyStrings(t *testing.T) {
	var list []*ScheduleMaintenance
	require.NoError(t, json.Unmarshal([]byte(`[{"maintenance_id":"1","maintenance_type":1,"end_day":"","start_week":""}]`), &list))
	require.Len(t, list, 1)
	assert.Equal(t, MaintenanceTypeDaily, list[0].MaintenanceType)
}

func boolPointer(value bool) *bool {
	return &value
}

func TestScheduleMaintenanceMarshalJSONSendsDaysOnlyForTheTypesThatUseThem(t *testing.T) {
	tests := []struct {
		name        string
		maintenance *ScheduleMaintenance
		expected    map[string]interface{}
		absent      []string
	}{
		{
			name:        "weekly by time on Sunday",
			maintenance: &ScheduleMaintenance{MaintenanceType: MaintenanceTypeWeeklyByTime, StartDay: 0, EndDay: 0, EndTime: "02:00"},
			expected:    map[string]interface{}{"start_day": float64(0), "end_day": float64(0), "end_time": "02:00"},
		},
		{
			name:        "monthly by day on Sunday",
			maintenance: &ScheduleMaintenance{MaintenanceType: MaintenanceTypeMonthlyByDay, StartDay: 0, StartWeek: 1, Duration: 30},
			expected:    map[string]interface{}{"start_day": float64(0), "start_week": float64(1), "duration": float64(30)},
			absent:      []string{"end_day", "end_time", "start_date", "end_date"},
		},
		{
			name:        "once",
			maintenance: &ScheduleMaintenance{MaintenanceType: MaintenanceTypeOnce, StartDay: 3, EndDay: 4, StartDate: "2026-10-01"},
			expected:    map[string]interface{}{"start_date": "2026-10-01"},
			absent:      []string{"start_day", "end_day", "duration", "end_date"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body, err := json.Marshal(test.maintenance)
			require.NoError(t, err)

			var sent map[string]interface{}
			require.NoError(t, json.Unmarshal(body, &sent))

			for key, value := range test.expected {
				assert.Equal(t, value, sent[key], key)
			}
			for _, key := range test.absent {
				assert.NotContains(t, sent, key)
			}
		})
	}
}
