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
