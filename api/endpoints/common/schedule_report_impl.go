package common

import (
	"fmt"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

// ScheduleReport defines operations for Site24x7 scheduled reports.
type ScheduleReport interface {
	Get(reportID string) (*api.ScheduleReport, error)
	Create(scheduleReport *api.ScheduleReport) (*api.ScheduleReport, error)
	Update(scheduleReport *api.ScheduleReport) (*api.ScheduleReport, error)
	Delete(reportID string) error
	List() ([]*api.ScheduleReport, error)
}

type schedulereport struct {
	client rest.Client
}

// NewScheduleReport creates a new ScheduleReport client instance.
func NewScheduleReport(client rest.Client) ScheduleReport {
	return &schedulereport{
		client: client,
	}
}

// Get fetches a single scheduled report by its report_id.
func (c *schedulereport) Get(reportID string) (*api.ScheduleReport, error) {
	sr := &api.ScheduleReport{}
	err := c.client.
		Get().
		Resource("scheduled_reports").
		ResourceID(reportID).
		Do().
		Parse(sr)

	return sr, err
}

// Create schedules a new report.
func (c *schedulereport) Create(scheduleReport *api.ScheduleReport) (*api.ScheduleReport, error) {
	newSR := &api.ScheduleReport{}
	err := c.client.
		Post().
		Resource("scheduled_reports").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(scheduleReport).
		Do().
		Parse(newSR)

	if err != nil {
		return nil, err
	}

	// Defensive check — ensure report_id is returned
	if newSR.ReportID == "" {
		return nil, fmt.Errorf("no report_id returned in create response")
	}

	return newSR, nil
}

// Update modifies an existing scheduled report.
func (c *schedulereport) Update(scheduleReport *api.ScheduleReport) (*api.ScheduleReport, error) {
	if scheduleReport.ReportID == "" {
		return nil, fmt.Errorf("cannot update scheduled report: missing ReportID")
	}

	updatedSR := &api.ScheduleReport{}
	err := c.client.
		Put().
		Resource("scheduled_reports").
		ResourceID(scheduleReport.ReportID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(scheduleReport).
		Do().
		Parse(updatedSR)

	if err != nil {
		return nil, err
	}

	return updatedSR, nil
}

// Delete removes a scheduled report by report_id.
func (c *schedulereport) Delete(reportID string) error {
	if reportID == "" {
		return fmt.Errorf("cannot delete scheduled report: missing reportID")
	}

	return c.client.
		Delete().
		Resource("scheduled_reports").
		ResourceID(reportID).
		Do().
		Err()
}

// List retrieves all scheduled reports.
func (c *schedulereport) List() ([]*api.ScheduleReport, error) {
	list := []*api.ScheduleReport{}
	err := c.client.
		Get().
		Resource("scheduled_reports").
		Do().
		Parse(&list)

	return list, err
}
