package common

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScheduleReport(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create scheduleReport",
			ExpectedVerb: "POST",
			ExpectedPath: "/scheduled_reports",
			ExpectedBody: validation.Fixture(t, "requests/create_schedule_report.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				scheduleReportCreate := &api.ScheduleReport{
					DisplayName:     "Daily Report",
					ReportType:      17, // Summary report
					SelectionType:   2,  // MONITOR
					ReportFrequency: 1,  // DAILY
					ReportFormat:    1,  // PDF
					UserGroups:      []string{"100"},
					ScheduledTime:   10,
				}

				_, err := NewScheduleReport(c).Create(scheduleReportCreate)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get scheduleReport",
			ExpectedVerb: "GET",
			ExpectedPath: "/scheduled_reports/123",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_schedule_report.json"),
			Fn: func(t *testing.T, c rest.Client) {
				report, err := NewScheduleReport(c).Get("123")
				require.NoError(t, err)

				expected := &api.ScheduleReport{
					ReportID:        "123",
					DisplayName:     "Daily Report",
					ReportType:      17,
					SelectionType:   2,
					ReportFrequency: 1,
					ReportFormat:    1,
					UserGroups:      []string{"100"},
					ScheduledTime:   10,
				}

				assert.Equal(t, expected, report)
			},
		},
		{
			Name:         "list scheduleReports",
			ExpectedVerb: "GET",
			ExpectedPath: "/scheduled_reports",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_schedule_report.json"),
			Fn: func(t *testing.T, c rest.Client) {
				reports, err := NewScheduleReport(c).List()
				require.NoError(t, err)

				expected := []*api.ScheduleReport{
					{
						ReportID:        "123",
						DisplayName:     "Daily Report",
						ReportType:      17,
						SelectionType:   2,
						ReportFrequency: 1,
						ReportFormat:    1,
						UserGroups:      []string{"100"},
						ScheduledTime:   10,
					},
					{
						ReportID:        "456",
						DisplayName:     "Weekly Report",
						ReportType:      17,
						SelectionType:   2,
						ReportFrequency: 2, // WEEKLY
						ReportFormat:    2, // CSV
						UserGroups:      []string{"200"},
						ScheduledTime:   11,
					},
				}

				assert.Equal(t, expected, reports)
			},
		},
		{
			Name:         "update scheduleReport",
			ExpectedVerb: "PUT",
			ExpectedPath: "/scheduled_reports/123",
			ExpectedBody: validation.Fixture(t, "requests/update_schedule_report.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				scheduleReportUpdate := &api.ScheduleReport{
					ReportID:        "123",
					DisplayName:     "Updated Report",
					ReportType:      17,
					SelectionType:   2,
					ReportFrequency: 2, // WEEKLY
					ReportFormat:    1, // PDF
					UserGroups:      []string{"300"},
					ScheduledTime:   12,
				}

				_, err := NewScheduleReport(c).Update(scheduleReportUpdate)
				require.NoError(t, err)
			},
		},
		{
			Name:         "updateRaw scheduleReport",
			ExpectedVerb: "PUT",
			ExpectedPath: "/scheduled_reports/123",
			ExpectedBody: validation.Fixture(t, "requests/update_raw_schedule_report.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				payload := map[string]interface{}{
					"display_name": "Raw Updated Report",
				}

				_, err := NewScheduleReport(c).UpdateRaw("123", payload)
				require.NoError(t, err)
			},
		},
		{
			Name:         "delete scheduleReport",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/scheduled_reports/123",
			StatusCode:   200,
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewScheduleReport(c).Delete("123"))
			},
		},
	})
}
