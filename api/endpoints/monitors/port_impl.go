package monitors

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type PortMonitors interface {
	Get(monitorID string) (*api.PortMonitor, error)
	Create(monitor *api.PortMonitor) (*api.PortMonitor, error)
	Update(monitor *api.PortMonitor) (*api.PortMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.PortMonitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type portMonitors struct {
	client rest.Client
}

func NewPortMonitors(client rest.Client) PortMonitors {
	return &portMonitors{
		client: client,
	}
}

func (c *portMonitors) Get(monitorID string) (*api.PortMonitor, error) {
	monitor := &api.PortMonitor{}

	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *portMonitors) Create(monitor *api.PortMonitor) (*api.PortMonitor, error) {
	newMonitor := &api.PortMonitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)

	return newMonitor, err
}

func (c *portMonitors) Update(monitor *api.PortMonitor) (*api.PortMonitor, error) {
	updatedMonitor := &api.PortMonitor{}
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

func (c *portMonitors) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *portMonitors) List() ([]*api.PortMonitor, error) {
	portMonitors := []*api.PortMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&portMonitors)

	return portMonitors, err
}

func (c *portMonitors) Activate(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/activate").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *portMonitors) Suspend(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/suspend").
		ResourceID(monitorID).
		Do().
		Err()
}
