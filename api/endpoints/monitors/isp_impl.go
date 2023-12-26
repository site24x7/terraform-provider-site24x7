package monitors

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type ISPMonitors interface {
	Get(monitorID string) (*api.ISPMonitor, error)
	Create(monitor *api.ISPMonitor) (*api.ISPMonitor, error)
	Update(monitor *api.ISPMonitor) (*api.ISPMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.ISPMonitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type ispmonitors struct {
	client rest.Client
}

func NewISPMonitors(client rest.Client) ISPMonitors {
	return &ispmonitors{
		client: client,
	}
}

func (c *ispmonitors) Get(monitorID string) (*api.ISPMonitor, error) {
	monitor := &api.ISPMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *ispmonitors) Create(monitor *api.ISPMonitor) (*api.ISPMonitor, error) {
	newMonitor := &api.ISPMonitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)

	return newMonitor, err
}

func (c *ispmonitors) Update(monitor *api.ISPMonitor) (*api.ISPMonitor, error) {
	updatedMonitor := &api.ISPMonitor{}
	err := c.client.
		Put().
		Resource("monitors").
		ResourceID(monitor.MonitorID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(updatedMonitor)

	return updatedMonitor, err
}

func (c *ispmonitors) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *ispmonitors) List() ([]*api.ISPMonitor, error) {
	ispmonitors := []*api.ISPMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&ispmonitors)

	return ispmonitors, err
}

func (c *ispmonitors) Activate(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/activate").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *ispmonitors) Suspend(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/suspend").
		ResourceID(monitorID).
		Do().
		Err()
}
