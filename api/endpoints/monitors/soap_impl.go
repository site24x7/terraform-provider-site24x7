package monitors

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type SOAPMonitors interface {
	Get(monitorID string) (*api.SOAPMonitor, error)
	Create(monitor *api.SOAPMonitor) (*api.SOAPMonitor, error)
	Update(monitor *api.SOAPMonitor) (*api.SOAPMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.SOAPMonitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type soapMonitors struct {
	client rest.Client
}

func NewSOAPMonitors(client rest.Client) SOAPMonitors {
	return &soapMonitors{
		client: client,
	}
}

func (c *soapMonitors) Get(monitorID string) (*api.SOAPMonitor, error) {
	monitor := &api.SOAPMonitor{}

	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *soapMonitors) Create(monitor *api.SOAPMonitor) (*api.SOAPMonitor, error) {
	newMonitor := &api.SOAPMonitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)

	return newMonitor, err
}

func (c *soapMonitors) Update(monitor *api.SOAPMonitor) (*api.SOAPMonitor, error) {
	updatedMonitor := &api.SOAPMonitor{}
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

func (c *soapMonitors) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *soapMonitors) List() ([]*api.SOAPMonitor, error) {
	soapMonitors := []*api.SOAPMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&soapMonitors)

	return soapMonitors, err
}

func (c *soapMonitors) Activate(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/activate").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *soapMonitors) Suspend(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/suspend").
		ResourceID(monitorID).
		Do().
		Err()
}
