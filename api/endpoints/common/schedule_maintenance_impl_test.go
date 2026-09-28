package common

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScheduleMaintenance(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create scheduleMaintenances",
			ExpectedVerb: "POST",
			ExpectedPath: "/maintenance",
			// ⚠️ Updated: create now expects timezone
			ExpectedBody: validation.Fixture(t, "requests/create_schedule_maintenance.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				scheduleMaintenanceCreate := &api.ScheduleMaintenance{
					DisplayName:       "Schedule Maintenance",
					Description:       "Maintenance Window",
					MaintenanceType:   3,
					TimeZone:          "PST", // explicitly set → must be in fixture
					StartDate:         "2022-06-02",
					EndDate:           "2022-06-02",
					StartTime:         "19:41",
					EndTime:           "20:44",
					SelectionType:     0,
					Monitors:          []string{"123", "456"},
					PerformMonitoring: true,
				}

				_, err := NewScheduleMaintenance(c).Create(scheduleMaintenanceCreate)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get scheduleMaintenances",
			ExpectedVerb: "GET",
			ExpectedPath: "/maintenance/113770000041271035",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_schedule_maintenance.json"),
			Fn: func(t *testing.T, c rest.Client) {
				group, err := NewScheduleMaintenance(c).Get("113770000041271035")
				require.NoError(t, err)

				expected := &api.ScheduleMaintenance{
					DisplayName:       "Schedule Maintenance",
					Description:       "Maintenance Window",
					MaintenanceType:   3,
					StartDate:         "2022-06-02",
					EndDate:           "2022-06-02",
					StartTime:         "19:41",
					EndTime:           "20:44",
					SelectionType:     0,
					Monitors:          []string{"123", "456"},
					PerformMonitoring: true,
				}

				assert.Equal(t, expected, group)
			},
		},
		{
			Name:         "list scheduleMaintenances",
			ExpectedVerb: "GET",
			ExpectedPath: "/maintenance",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_schedule_maintenance.json"),
			Fn: func(t *testing.T, c rest.Client) {
				groups, err := NewScheduleMaintenance(c).List()
				require.NoError(t, err)

				expected := []*api.ScheduleMaintenance{
					{
						MaintenanceID:     "123",
						DisplayName:       "Schedule Maintenance",
						Description:       "Maintenance Window",
						MaintenanceType:   3,
						StartDate:         "2022-06-02",
						EndDate:           "2022-06-02",
						StartTime:         "19:41",
						EndTime:           "20:44",
						SelectionType:     0,
						Monitors:          []string{"123", "456"},
						PerformMonitoring: true,
					},
					{
						MaintenanceID:     "456",
						DisplayName:       "Schedule Maintenance",
						Description:       "Maintenance Window",
						MaintenanceType:   3,
						StartDate:         "2022-06-02",
						EndDate:           "2022-06-02",
						StartTime:         "19:41",
						EndTime:           "20:44",
						SelectionType:     0,
						Monitors:          []string{"123", "456"},
						PerformMonitoring: true,
					},
				}

				assert.Equal(t, expected, groups)
			},
		},
		{
			// Mirrors the List Maintenance example in the API reference, with
			// one maintenance per selection type, so every response field the
			// site24x7_schedule_maintenances data source reads is parsed.
			Name:         "list scheduleMaintenances of every selection type",
			ExpectedVerb: "GET",
			ExpectedPath: "/maintenance",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_schedule_maintenance_all_selection_types.json"),
			Fn: func(t *testing.T, c rest.Client) {
				maintenances, err := NewScheduleMaintenance(c).List()
				require.NoError(t, err)

				subgroupMonitors := false
				expected := []*api.ScheduleMaintenance{
					{
						MaintenanceID:        "113770000041409009",
						DisplayName:          "Weekly by days",
						Description:          "BI-WEEKLY",
						MaintenanceType:      8,
						MaintenanceStatus:    "V",
						TimeZone:             "IST",
						WeekDays:             []int{3, 5},
						StartTime:            "14:26",
						Duration:             float64(30),
						ExecuteEvery:         2,
						MaintenanceStartOn:   "2022-08-03",
						SelectionType:        2,
						Monitors:             []string{"113770000039133011"},
						MaintenanceStartTime: "2022-08-03T15:22:10+0530",
						ZUID:                 "65478659",
						PerformMonitoring:    true,
					},
					{
						MaintenanceID:        "113770000041409010",
						DisplayName:          "Monthly By day",
						Description:          "Patch Tuesday",
						MaintenanceType:      6,
						MaintenanceStatus:    "V",
						TimeZone:             "IST",
						StartWeek:            2,
						StartDay:             2,
						StartAfter:           2,
						StartTime:            "12:30",
						Duration:             float64(30),
						MaintenanceStartOn:   "2022-08-03",
						MaintenanceEndType:   2,
						MaintenanceEndOn:     "2023-08-03",
						SelectionType:        1,
						MonitorGroups:        []string{"113770000039135019"},
						SubgroupMonitors:     &subgroupMonitors,
						MaintenanceStartTime: "2022-08-03T15:22:10+0530",
						ZUID:                 "65478659",
						PerformMonitoring:    false,
					},
					{
						MaintenanceID:            "113770000041409012",
						DisplayName:              "Yearly",
						Description:              "Christmas",
						MaintenanceType:          9,
						MaintenanceStatus:        "V",
						TimeZone:                 "IST",
						MonthlyStartDate:         25,
						StartTime:                "00:00",
						EndTime:                  "23:59",
						Duration:                 "",
						MaintenanceStartOn:       "2022-08-03",
						MaintenanceEndType:       1,
						MaintenanceEndAfterTimes: 5,
						SelectionType:            3,
						Tags:                     []string{"113770000039135099"},
						MaintenanceStartTime:     "2022-08-03T15:22:10+0530",
						ZUID:                     "65478659",
						PerformMonitoring:        true,
					},
				}

				assert.Equal(t, expected, maintenances)
			},
		},
		{
			Name:         "update scheduleMaintenances",
			ExpectedVerb: "PUT",
			ExpectedPath: "/maintenance/123",
			ExpectedBody: validation.Fixture(t, "requests/update_schedule_maintenance.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				scheduleMaintenanceUpdate := &api.ScheduleMaintenance{
					MaintenanceID:     "123",
					DisplayName:       "Maintenance Update",
					Description:       "Maintenance Window",
					MaintenanceType:   3,
					StartDate:         "2022-06-02",
					EndDate:           "2022-06-02",
					StartTime:         "19:41",
					EndTime:           "20:44",
					SelectionType:     0,
					Monitors:          []string{"123", "456"},
					PerformMonitoring: true,
				}

				_, err := NewScheduleMaintenance(c).Update(scheduleMaintenanceUpdate)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete scheduleMaintenances",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/maintenance/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewScheduleMaintenance(c).Delete("123"))
			},
		},
	})
}
