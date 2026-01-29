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

func TestScheduleReportCreate(t *testing.T) {
	d := scheduleReportTestResourceData(t)

	c := fake.NewClient()

	a := &api.ScheduleReport{
		DisplayName:     "Daily Summary",
		ReportType:      17,
		SelectionType:   2,
		ReportFormat:    1,
		ReportFrequency: 1,
		ScheduledTime:   10,
		ScheduledDay:    2,
		UserGroups:      []string{"100", "200"},
	}

	created := *a
	created.ReportID = "sr-123"

	c.FakeScheduleReport.On("Create", a).Return(&created, nil).Once()

	require.NoError(t, scheduleReportCreate(d, c))
	assert.Equal(t, "sr-123", d.Id())

	c.FakeScheduleReport.On("Create", a).
		Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := scheduleReportCreate(d, c)
	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestScheduleReportUpdate(t *testing.T) {
	d := scheduleReportTestResourceData(t)
	d.SetId("sr-123")

	c := fake.NewClient()

	payload := scheduleReportUpdatePayload{
		DisplayName:     "Daily Summary",
		ReportType:      17,
		SelectionType:   2,
		ReportFormat:    1,
		ReportFrequency: 1,
		ScheduledTime:   10,
		ScheduledDay:    2,
		UserGroups:      []string{"100", "200"},
	}

	updated := &api.ScheduleReport{
		ReportID:        "sr-123",
		DisplayName:     "Daily Summary",
		ReportType:      17,
		SelectionType:   2,
		ReportFormat:    1,
		ReportFrequency: 1,
		ScheduledTime:   10,
		ScheduledDay:    2,
		UserGroups:      []string{"100", "200"},
	}

	c.FakeScheduleReport.
		On("UpdateRaw", "sr-123", payload).
		Return(updated, nil).Once()

	require.NoError(t, scheduleReportUpdate(d, c))

	c.FakeScheduleReport.
		On("UpdateRaw", "sr-123", payload).
		Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := scheduleReportUpdate(d, c)
	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestScheduleReportRead(t *testing.T) {
	d := scheduleReportTestResourceData(t)
	d.SetId("sr-123")

	c := fake.NewClient()

	sr := &api.ScheduleReport{
		ReportID:        "sr-123",
		DisplayName:     "Daily Summary",
		ReportType:      17,
		SelectionType:   2,
		ReportFormat:    1,
		ReportFrequency: 1,
		ScheduledTime:   10,
		ScheduledDay:    2,
		UserGroups:      []string{"100", "200"},
	}

	c.FakeScheduleReport.On("Get", "sr-123").Return(sr, nil).Once()
	require.NoError(t, scheduleReportRead(d, c))

	assert.Equal(t, "Daily Summary", d.Get("display_name"))

	c.FakeScheduleReport.
		On("Get", "sr-123").
		Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := scheduleReportRead(d, c)
	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestScheduleReportDelete(t *testing.T) {
	d := scheduleReportTestResourceData(t)
	d.SetId("sr-123")

	c := fake.NewClient()

	c.FakeScheduleReport.On("Delete", "sr-123").Return(nil).Once()
	require.NoError(t, scheduleReportDelete(d, c))

	c.FakeScheduleReport.
		On("Delete", "sr-123").
		Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, scheduleReportDelete(d, c))
}

func TestScheduleReportExists(t *testing.T) {
	d := scheduleReportTestResourceData(t)
	d.SetId("sr-123")

	c := fake.NewClient()

	c.FakeScheduleReport.
		On("Get", "sr-123").
		Return(&api.ScheduleReport{}, nil).Once()

	exists, err := scheduleReportExists(d, c)
	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeScheduleReport.
		On("Get", "sr-123").
		Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = scheduleReportExists(d, c)
	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeScheduleReport.
		On("Get", "sr-123").
		Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = scheduleReportExists(d, c)
	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func scheduleReportTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, ScheduleReportSchema, map[string]interface{}{
		"display_name":     "Daily Summary",
		"report_type":      17,
		"selection_type":   2,
		"report_format":    1,
		"report_frequency": 1,
		"scheduled_time":   10,
		"scheduled_day":    2,
		"user_groups": []interface{}{
			"100",
			"200",
		},
	})
}
