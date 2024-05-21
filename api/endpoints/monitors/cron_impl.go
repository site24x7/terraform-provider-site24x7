package monitors

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type CronMonitors interface {
	Get(monitorID string) (*api.CronMonitor, error)
	Create(monitor *api.CronMonitor) (*api.CronMonitor, error)
	Update(monitor *api.CronMonitor) (*api.CronMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.CronMonitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type cronMonitors struct {
	client rest.Client
}

func NewCronMonitors(client rest.Client) CronMonitors {
	return &cronMonitors{
		client: client,
	}
}

func (c *cronMonitors) Get(monitorID string) (*api.CronMonitor, error) {
	monitor := &api.CronMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *cronMonitors) Create(monitor *api.CronMonitor) (*api.CronMonitor, error) {
	newMonitor := &api.CronMonitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)

	return newMonitor, err
}

func (c *cronMonitors) Update(monitor *api.CronMonitor) (*api.CronMonitor, error) {
	updatedMonitor := &api.CronMonitor{}
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

func (c *cronMonitors) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *cronMonitors) List() ([]*api.CronMonitor, error) {
	cronMonitors := []*api.CronMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&cronMonitors)

	return cronMonitors, err
}

func (c *cronMonitors) Activate(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/activate").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *cronMonitors) Suspend(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/suspend").
		ResourceID(monitorID).
		Do().
		Err()
}
