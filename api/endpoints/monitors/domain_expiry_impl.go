package monitors

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type DomainExpiryMonitors interface {
	Get(monitorID string) (*api.DomainExpiryMonitor, error)
	Create(monitor *api.DomainExpiryMonitor) (*api.DomainExpiryMonitor, error)
	Update(monitor *api.DomainExpiryMonitor) (*api.DomainExpiryMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.DomainExpiryMonitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type domainExpiryMonitors struct {
	client rest.Client
}

func NewDomainExpiryMonitors(client rest.Client) DomainExpiryMonitors {
	return &domainExpiryMonitors{
		client: client,
	}
}

func (c *domainExpiryMonitors) Get(monitorID string) (*api.DomainExpiryMonitor, error) {
	monitor := &api.DomainExpiryMonitor{}

	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *domainExpiryMonitors) Create(monitor *api.DomainExpiryMonitor) (*api.DomainExpiryMonitor, error) {
	newMonitor := &api.DomainExpiryMonitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)

	return newMonitor, err
}

func (c *domainExpiryMonitors) Update(monitor *api.DomainExpiryMonitor) (*api.DomainExpiryMonitor, error) {
	updatedMonitor := &api.DomainExpiryMonitor{}
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

func (c *domainExpiryMonitors) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *domainExpiryMonitors) List() ([]*api.DomainExpiryMonitor, error) {
	domainExpiryMonitors := []*api.DomainExpiryMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&domainExpiryMonitors)

	return domainExpiryMonitors, err
}

func (c *domainExpiryMonitors) Activate(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/activate").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *domainExpiryMonitors) Suspend(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/suspend").
		ResourceID(monitorID).
		Do().
		Err()
}
