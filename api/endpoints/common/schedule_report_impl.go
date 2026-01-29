package common

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

// ScheduleReport defines operations for Site24x7 scheduled reports.
type ScheduleReport interface {
	Get(reportID string) (*api.ScheduleReport, error)
	Create(scheduleReport *api.ScheduleReport) (*api.ScheduleReport, error)
	Update(scheduleReport *api.ScheduleReport) (*api.ScheduleReport, error)
	Delete(reportID string) error
	UpdateRaw(reportID string, payload interface{}) (*api.ScheduleReport, error)
	List() ([]*api.ScheduleReport, error)
}

type schedulereport struct {
	client rest.Client
}

func NewScheduleReport(client rest.Client) ScheduleReport {
	return &schedulereport{
		client: client,
	}
}

func (c *schedulereport) Get(reportID string) (*api.ScheduleReport, error) {
	scheduleReport := &api.ScheduleReport{}
	err := c.client.
		Get().
		Resource("scheduled_reports").
		ResourceID(reportID).
		Do().
		Parse(scheduleReport)

	return scheduleReport, err
}

func (c *schedulereport) Create(scheduleReport *api.ScheduleReport) (*api.ScheduleReport, error) {
	newScheduleReport := &api.ScheduleReport{}
	err := c.client.
		Post().
		Resource("scheduled_reports").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(scheduleReport).
		Do().
		Parse(newScheduleReport)

	return newScheduleReport, err
}

func (c *schedulereport) Update(scheduleReport *api.ScheduleReport) (*api.ScheduleReport, error) {
	updatedScheduleReport := &api.ScheduleReport{}
	err := c.client.
		Put().
		Resource("scheduled_reports").
		ResourceID(scheduleReport.ReportID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(scheduleReport).
		Do().
		Parse(updatedScheduleReport)

	return updatedScheduleReport, err
}

func (c *schedulereport) UpdateRaw(reportID string, payload interface{}) (*api.ScheduleReport, error) {
	updated := &api.ScheduleReport{}
	err := c.client.
		Put().
		Resource("scheduled_reports").
		ResourceID(reportID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(payload).
		Do().
		Parse(updated)

	return updated, err
}

func (c *schedulereport) Delete(reportID string) error {
	return c.client.
		Delete().
		Resource("scheduled_reports").
		ResourceID(reportID).
		Do().
		Err()
}

func (c *schedulereport) List() ([]*api.ScheduleReport, error) {
	scheduleReportList := []*api.ScheduleReport{}
	err := c.client.
		Get().
		Resource("scheduled_reports").
		Do().
		Parse(&scheduleReportList)

	return scheduleReportList, err
}
