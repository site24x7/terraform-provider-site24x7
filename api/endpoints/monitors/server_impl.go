package monitors

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type ServerMonitors interface {
	Get(monitorID string) (*api.ServerMonitor, error)
	Create(monitor *api.ServerMonitor) (*api.ServerMonitor, error)
	Update(monitor *api.ServerMonitor) (*api.ServerMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.ServerMonitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type servermonitors struct {
	client rest.Client
}

func NewServerMonitors(client rest.Client) ServerMonitors {
	return &servermonitors{
		client: client,
	}
}

func (c *servermonitors) Get(monitorID string) (*api.ServerMonitor, error) {
	monitor := &api.ServerMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *servermonitors) Create(monitor *api.ServerMonitor) (*api.ServerMonitor, error) {
	newMonitor := &api.ServerMonitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)

	return newMonitor, err
}

func (c *servermonitors) Update(monitor *api.ServerMonitor) (*api.ServerMonitor, error) {
	updatedMonitor := &api.ServerMonitor{}
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

func (c *servermonitors) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *servermonitors) List() ([]*api.ServerMonitor, error) {
	servermonitors := []*api.ServerMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&servermonitors)

	return servermonitors, err
}

func (c *servermonitors) Activate(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/activate").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *servermonitors) Suspend(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/suspend").
		ResourceID(monitorID).
		Do().
		Err()
}
