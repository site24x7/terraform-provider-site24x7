package monitors

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type PINGMonitors interface {
	Get(monitorID string) (*api.PINGMonitor, error)
	Create(monitor *api.PINGMonitor) (*api.PINGMonitor, error)
	Update(monitor *api.PINGMonitor) (*api.PINGMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.PINGMonitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type pingMonitors struct {
	client rest.Client
}

func NewPINGMonitors(client rest.Client) PINGMonitors {
	return &pingMonitors{
		client: client,
	}
}

func (c *pingMonitors) Get(monitorID string) (*api.PINGMonitor, error) {
	monitor := &api.PINGMonitor{}

	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *pingMonitors) Create(monitor *api.PINGMonitor) (*api.PINGMonitor, error) {
	newMonitor := &api.PINGMonitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)

	return newMonitor, err
}

func (c *pingMonitors) Update(monitor *api.PINGMonitor) (*api.PINGMonitor, error) {
	updatedMonitor := &api.PINGMonitor{}
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

func (c *pingMonitors) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *pingMonitors) List() ([]*api.PINGMonitor, error) {
	pingMonitors := []*api.PINGMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&pingMonitors)

	return pingMonitors, err
}

func (c *pingMonitors) Activate(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/activate").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *pingMonitors) Suspend(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/suspend").
		ResourceID(monitorID).
		Do().
		Err()
}
