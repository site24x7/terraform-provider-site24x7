package monitors

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type WebTransactionBrowserMonitors interface {
	Get(monitorID string) (*api.WebTransactionBrowserMonitor, error)
	Create(monitor *api.WebTransactionBrowserMonitor) (*api.WebTransactionBrowserMonitor, error)
	Update(monitor *api.WebTransactionBrowserMonitor) (*api.WebTransactionBrowserMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.WebTransactionBrowserMonitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type webTransactionBrowserMonitors struct {
	client rest.Client
}

func NewWebTransactionBrowserMonitors(client rest.Client) WebTransactionBrowserMonitors {
	return &webTransactionBrowserMonitors{
		client: client,
	}
}

func (c *webTransactionBrowserMonitors) Get(monitorID string) (*api.WebTransactionBrowserMonitor, error) {
	monitor := &api.WebTransactionBrowserMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *webTransactionBrowserMonitors) Create(monitor *api.WebTransactionBrowserMonitor) (*api.WebTransactionBrowserMonitor, error) {
	newMonitor := &api.WebTransactionBrowserMonitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)

	return newMonitor, err
}

func (c *webTransactionBrowserMonitors) Update(monitor *api.WebTransactionBrowserMonitor) (*api.WebTransactionBrowserMonitor, error) {
	updatedMonitor := &api.WebTransactionBrowserMonitor{}
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

func (c *webTransactionBrowserMonitors) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *webTransactionBrowserMonitors) List() ([]*api.WebTransactionBrowserMonitor, error) {
	webTransactionBrowserMonitors := []*api.WebTransactionBrowserMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&webTransactionBrowserMonitors)

	return webTransactionBrowserMonitors, err
}

func (c *webTransactionBrowserMonitors) Activate(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/activate").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *webTransactionBrowserMonitors) Suspend(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/suspend").
		ResourceID(monitorID).
		Do().
		Err()
}
