package endpoints

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type SSLMonitors interface {
	Get(monitorID string) (*api.SSLMonitor, error)
	Create(monitor *api.SSLMonitor) (*api.SSLMonitor, error)
	Update(monitor *api.SSLMonitor) (*api.SSLMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.SSLMonitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type sslmonitors struct {
	client rest.Client
}

func NewSSLMonitors(client rest.Client) SSLMonitors {
	return &sslmonitors{
		client: client,
	}
}

func (c *sslmonitors) Get(monitorID string) (*api.SSLMonitor, error) {
	monitor := &api.SSLMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *sslmonitors) Create(monitor *api.SSLMonitor) (*api.SSLMonitor, error) {
	newMonitor := &api.SSLMonitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)

	return newMonitor, err
}

func (c *sslmonitors) Update(monitor *api.SSLMonitor) (*api.SSLMonitor, error) {
	updatedMonitor := &api.SSLMonitor{}
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

func (c *sslmonitors) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *sslmonitors) List() ([]*api.SSLMonitor, error) {
	sslmonitors := []*api.SSLMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&sslmonitors)

	return sslmonitors, err
}

func (c *sslmonitors) Activate(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/activate").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *sslmonitors) Suspend(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/suspend").
		ResourceID(monitorID).
		Do().
		Err()
}
