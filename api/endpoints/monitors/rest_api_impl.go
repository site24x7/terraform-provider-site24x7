package monitors

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type RestApiMonitors interface {
	Get(monitorID string) (*api.RestApiMonitor, error)
	Create(monitor *api.RestApiMonitor) (*api.RestApiMonitor, error)
	Update(monitor *api.RestApiMonitor) (*api.RestApiMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.RestApiMonitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type restapimonitors struct {
	client rest.Client
}

func NewRestApiMonitors(client rest.Client) RestApiMonitors {
	return &restapimonitors{
		client: client,
	}
}

func (c *restapimonitors) Get(monitorID string) (*api.RestApiMonitor, error) {
	monitor := &api.RestApiMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *restapimonitors) Create(monitor *api.RestApiMonitor) (*api.RestApiMonitor, error) {
	newMonitor := &api.RestApiMonitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)

	return newMonitor, err
}

func (c *restapimonitors) Update(monitor *api.RestApiMonitor) (*api.RestApiMonitor, error) {
	updatedMonitor := &api.RestApiMonitor{}
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

func (c *restapimonitors) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *restapimonitors) List() ([]*api.RestApiMonitor, error) {
	restapimonitors := []*api.RestApiMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&restapimonitors)

	return restapimonitors, err
}

func (c *restapimonitors) Activate(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/activate").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *restapimonitors) Suspend(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/suspend").
		ResourceID(monitorID).
		Do().
		Err()
}
