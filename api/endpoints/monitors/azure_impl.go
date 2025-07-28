package monitors

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type AzureMonitors interface {
	Get(monitorID string) (*api.AzureMonitor, error)
	Create(monitor *api.AzureMonitor) (*api.AzureMonitor, error)
	Update(monitor *api.AzureMonitor) (*api.AzureMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.AzureMonitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type azureMonitors struct {
	client rest.Client
}

func NewAzureMonitors(client rest.Client) AzureMonitors {
	return &azureMonitors{
		client: client,
	}
}

func (c *azureMonitors) Get(monitorID string) (*api.AzureMonitor, error) {
	monitor := &api.AzureMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *azureMonitors) Create(monitor *api.AzureMonitor) (*api.AzureMonitor, error) {
	newMonitor := &api.AzureMonitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)

	return newMonitor, err
}

func (c *azureMonitors) Update(monitor *api.AzureMonitor) (*api.AzureMonitor, error) {
	updatedMonitor := &api.AzureMonitor{}
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

func (c *azureMonitors) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *azureMonitors) List() ([]*api.AzureMonitor, error) {
	monitors := []*api.AzureMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&monitors)

	return monitors, err
}

func (c *azureMonitors) Activate(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/activate").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *azureMonitors) Suspend(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/suspend").
		ResourceID(monitorID).
		Do().
		Err()
}
