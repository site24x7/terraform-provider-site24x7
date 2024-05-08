package monitors

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type WebsiteMonitors interface {
	Get(monitorID string) (*api.WebsiteMonitor, error)
	Create(monitor *api.WebsiteMonitor) (*api.WebsiteMonitor, error)
	Update(monitor *api.WebsiteMonitor) (*api.WebsiteMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.WebsiteMonitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type websitemonitors struct {
	client rest.Client
}

func NewMonitors(client rest.Client) WebsiteMonitors {
	return &websitemonitors{
		client: client,
	}
}

func (c *websitemonitors) Get(monitorID string) (*api.WebsiteMonitor, error) {
	monitor := &api.WebsiteMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *websitemonitors) Create(monitor *api.WebsiteMonitor) (*api.WebsiteMonitor, error) {
	newMonitor := &api.WebsiteMonitor{}
	newMonitor.UseIPV6 = monitor.UseIPV6
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)
	return newMonitor, err
}

func (c *websitemonitors) Update(monitor *api.WebsiteMonitor) (*api.WebsiteMonitor, error) {
	updatedMonitor := &api.WebsiteMonitor{}
	updatedMonitor.UseIPV6 = monitor.UseIPV6
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

func (c *websitemonitors) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *websitemonitors) List() ([]*api.WebsiteMonitor, error) {
	monitors := []*api.WebsiteMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&monitors)

	return monitors, err
}

func (c *websitemonitors) Activate(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/activate").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *websitemonitors) Suspend(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/suspend").
		ResourceID(monitorID).
		Do().
		Err()
}
