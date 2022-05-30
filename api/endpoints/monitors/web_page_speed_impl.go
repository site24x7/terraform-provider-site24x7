package monitors

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type WebPageSpeedMonitors interface {
	Get(monitorID string) (*api.WebPageSpeedMonitor, error)
	Create(monitor *api.WebPageSpeedMonitor) (*api.WebPageSpeedMonitor, error)
	Update(monitor *api.WebPageSpeedMonitor) (*api.WebPageSpeedMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.WebPageSpeedMonitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type webpagespeedmonitors struct {
	client rest.Client
}

func NewWebPageSpeedMonitors(client rest.Client) WebPageSpeedMonitors {
	return &webpagespeedmonitors{
		client: client,
	}
}

func (c *webpagespeedmonitors) Get(monitorID string) (*api.WebPageSpeedMonitor, error) {
	monitor := &api.WebPageSpeedMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *webpagespeedmonitors) Create(monitor *api.WebPageSpeedMonitor) (*api.WebPageSpeedMonitor, error) {
	newMonitor := &api.WebPageSpeedMonitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)

	return newMonitor, err
}

func (c *webpagespeedmonitors) Update(monitor *api.WebPageSpeedMonitor) (*api.WebPageSpeedMonitor, error) {
	updatedMonitor := &api.WebPageSpeedMonitor{}
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

func (c *webpagespeedmonitors) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *webpagespeedmonitors) List() ([]*api.WebPageSpeedMonitor, error) {
	monitors := []*api.WebPageSpeedMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&monitors)

	return monitors, err
}

func (c *webpagespeedmonitors) Activate(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/activate").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *webpagespeedmonitors) Suspend(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/suspend").
		ResourceID(monitorID).
		Do().
		Err()
}
