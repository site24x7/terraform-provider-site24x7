package common

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type ScheduleMaintenance interface {
	Get(schedulemaintenanceID string) (*api.ScheduleMaintenance, error)
	Create(schedulemaintenance *api.ScheduleMaintenance) (*api.ScheduleMaintenance, error)
	Update(schedulemaintenance *api.ScheduleMaintenance) (*api.ScheduleMaintenance, error)
	Delete(schedulemaintenanceID string) error
	List() ([]*api.ScheduleMaintenance, error)
}

type schedulemaintenance struct {
	client rest.Client
}

func NewScheduleMaintenance(client rest.Client) ScheduleMaintenance {
	return &schedulemaintenance{
		client: client,
	}
}

func (c *schedulemaintenance) Get(schedulemaintenanceID string) (*api.ScheduleMaintenance, error) {
	schedulemaintenance := &api.ScheduleMaintenance{}
	err := c.client.
		Get().
		Resource("maintenance").
		ResourceID(schedulemaintenanceID).
		Do().
		Parse(schedulemaintenance)

	return schedulemaintenance, err
}

func (c *schedulemaintenance) Create(schedulemaintenance *api.ScheduleMaintenance) (*api.ScheduleMaintenance, error) {
	newScheduleMaintenance := &api.ScheduleMaintenance{}
	err := c.client.
		Post().
		Resource("maintenance").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(schedulemaintenance).
		Do().
		Parse(newScheduleMaintenance)

	return newScheduleMaintenance, err
}

func (c *schedulemaintenance) Update(schedulemaintenance *api.ScheduleMaintenance) (*api.ScheduleMaintenance, error) {
	updatedScheduleMaintenance := &api.ScheduleMaintenance{}
	err := c.client.
		Put().
		Resource("maintenance").
		ResourceID(schedulemaintenance.MaintenanceID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(schedulemaintenance).
		Do().
		Parse(updatedScheduleMaintenance)

	return updatedScheduleMaintenance, err
}

func (c *schedulemaintenance) Delete(schedulemaintenanceID string) error {
	return c.client.
		Delete().
		Resource("maintenance").
		ResourceID(schedulemaintenanceID).
		Do().
		Err()
}

func (c *schedulemaintenance) List() ([]*api.ScheduleMaintenance, error) {
	scheduleMaintenanceList := []*api.ScheduleMaintenance{}
	err := c.client.
		Get().
		Resource("maintenance").
		Do().
		Parse(&scheduleMaintenanceList)

	return scheduleMaintenanceList, err
}
