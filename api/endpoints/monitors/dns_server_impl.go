package monitors

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type DNSServerMonitors interface {
	Get(monitorID string) (*api.DNSServerMonitor, error)
	Create(monitor *api.DNSServerMonitor) (*api.DNSServerMonitor, error)
	Update(monitor *api.DNSServerMonitor) (*api.DNSServerMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.DNSServerMonitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type dnsservermonitors struct {
	client rest.Client
}

func NewDNSServerMonitors(client rest.Client) DNSServerMonitors {
	return &dnsservermonitors{
		client: client,
	}
}

func (c *dnsservermonitors) Get(monitorID string) (*api.DNSServerMonitor, error) {
	monitor := &api.DNSServerMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *dnsservermonitors) Create(monitor *api.DNSServerMonitor) (*api.DNSServerMonitor, error) {
	newMonitor := &api.DNSServerMonitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)
	return newMonitor, err
}

func (c *dnsservermonitors) Update(monitor *api.DNSServerMonitor) (*api.DNSServerMonitor, error) {
	updatedMonitor := &api.DNSServerMonitor{}
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

func (c *dnsservermonitors) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *dnsservermonitors) List() ([]*api.DNSServerMonitor, error) {
	monitors := []*api.DNSServerMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&monitors)

	return monitors, err
}

func (c *dnsservermonitors) Activate(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/activate").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *dnsservermonitors) Suspend(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/suspend").
		ResourceID(monitorID).
		Do().
		Err()
}
