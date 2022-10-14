package monitors

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type HeartbeatMonitors interface {
	Get(monitorID string) (*api.HeartbeatMonitor, error)
	Create(monitor *api.HeartbeatMonitor) (*api.HeartbeatMonitor, error)
	Update(monitor *api.HeartbeatMonitor) (*api.HeartbeatMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.HeartbeatMonitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type heartbeatmonitors struct {
	client rest.Client
}

func NewHeartbeatMonitors(client rest.Client) HeartbeatMonitors {
	return &heartbeatmonitors{
		client: client,
	}
}

func (c *heartbeatmonitors) Get(monitorID string) (*api.HeartbeatMonitor, error) {
	monitor := &api.HeartbeatMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *heartbeatmonitors) Create(monitor *api.HeartbeatMonitor) (*api.HeartbeatMonitor, error) {
	newMonitor := &api.HeartbeatMonitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)

	return newMonitor, err
}

func (c *heartbeatmonitors) Update(monitor *api.HeartbeatMonitor) (*api.HeartbeatMonitor, error) {
	updatedMonitor := &api.HeartbeatMonitor{}
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

func (c *heartbeatmonitors) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *heartbeatmonitors) List() ([]*api.HeartbeatMonitor, error) {
	heartbeatmonitors := []*api.HeartbeatMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&heartbeatmonitors)

	return heartbeatmonitors, err
}

func (c *heartbeatmonitors) Activate(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/activate").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *heartbeatmonitors) Suspend(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/suspend").
		ResourceID(monitorID).
		Do().
		Err()
}
