package endpoints

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type AmazonMonitors interface {
	Get(monitorID string) (*api.AmazonMonitor, error)
	Create(monitor *api.AmazonMonitor) (*api.AmazonMonitor, error)
	Update(monitor *api.AmazonMonitor) (*api.AmazonMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.AmazonMonitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type amazonMonitors struct {
	client rest.Client
}

func NewAmazonMonitors(client rest.Client) AmazonMonitors {
	return &amazonMonitors{
		client: client,
	}
}

func (c *amazonMonitors) Get(monitorID string) (*api.AmazonMonitor, error) {
	monitor := &api.AmazonMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *amazonMonitors) Create(monitor *api.AmazonMonitor) (*api.AmazonMonitor, error) {
	newMonitor := &api.AmazonMonitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)

	return newMonitor, err
}

func (c *amazonMonitors) Update(monitor *api.AmazonMonitor) (*api.AmazonMonitor, error) {
	updatedMonitor := &api.AmazonMonitor{}
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

func (c *amazonMonitors) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *amazonMonitors) List() ([]*api.AmazonMonitor, error) {
	monitors := []*api.AmazonMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&monitors)

	return monitors, err
}

func (c *amazonMonitors) Activate(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/activate").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *amazonMonitors) Suspend(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/suspend").
		ResourceID(monitorID).
		Do().
		Err()
}
