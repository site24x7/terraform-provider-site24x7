package monitors

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type GCPMonitors interface {
	Get(monitorID string) (*api.GCPMonitor, error)
	Create(monitor *api.GCPMonitor) (*api.GCPMonitor, error)
	Update(monitor *api.GCPMonitor) (*api.GCPMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.GCPMonitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type gcpMonitors struct {
	client rest.Client
}

func NewGCPMonitors(client rest.Client) GCPMonitors {
	return &gcpMonitors{
		client: client,
	}
}

func (c *gcpMonitors) Get(monitorID string) (*api.GCPMonitor, error) {
	monitor := &api.GCPMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *gcpMonitors) Create(monitor *api.GCPMonitor) (*api.GCPMonitor, error) {
	newMonitor := &api.GCPMonitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)

	return newMonitor, err
}

func (c *gcpMonitors) Update(monitor *api.GCPMonitor) (*api.GCPMonitor, error) {
	updatedMonitor := &api.GCPMonitor{}
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

func (c *gcpMonitors) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *gcpMonitors) List() ([]*api.GCPMonitor, error) {
	monitors := []*api.GCPMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&monitors)

	return monitors, err
}

func (c *gcpMonitors) Activate(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/activate").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *gcpMonitors) Suspend(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/suspend").
		ResourceID(monitorID).
		Do().
		Err()
}
