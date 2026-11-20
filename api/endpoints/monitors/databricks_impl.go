package monitors

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type DatabricksMonitors interface {
	Get(monitorID string) (*api.DatabricksMonitor, error)
	Create(monitor *api.DatabricksMonitor) (*api.DatabricksMonitor, error)
	Update(monitor *api.DatabricksMonitor) (*api.DatabricksMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.DatabricksMonitor, error)
}

type databricksMonitors struct {
	client rest.Client
}

func NewDatabricksMonitors(client rest.Client) DatabricksMonitors {
	return &databricksMonitors{
		client: client,
	}
}

func (c *databricksMonitors) Get(monitorID string) (*api.DatabricksMonitor, error) {
	monitor := &api.DatabricksMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *databricksMonitors) Create(monitor *api.DatabricksMonitor) (*api.DatabricksMonitor, error) {
	newMonitor := &api.DatabricksMonitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)

	return newMonitor, err
}

func (c *databricksMonitors) Update(monitor *api.DatabricksMonitor) (*api.DatabricksMonitor, error) {
	updatedMonitor := &api.DatabricksMonitor{}
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

func (c *databricksMonitors) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *databricksMonitors) List() ([]*api.DatabricksMonitor, error) {
	monitors := []*api.DatabricksMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&monitors)

	return monitors, err
}
