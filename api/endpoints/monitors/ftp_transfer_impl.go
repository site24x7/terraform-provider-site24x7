package monitors

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type FTPTransferMonitors interface {
	Get(monitorID string) (*api.FTPTransferMonitor, error)
	Create(monitor *api.FTPTransferMonitor) (*api.FTPTransferMonitor, error)
	Update(monitor *api.FTPTransferMonitor) (*api.FTPTransferMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.FTPTransferMonitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type ftptransfermonitors struct {
	client rest.Client
}

func NewFTPTransferMonitors(client rest.Client) FTPTransferMonitors {
	return &ftptransfermonitors{
		client: client,
	}
}

func (c *ftptransfermonitors) Get(monitorID string) (*api.FTPTransferMonitor, error) {
	monitor := &api.FTPTransferMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *ftptransfermonitors) Create(monitor *api.FTPTransferMonitor) (*api.FTPTransferMonitor, error) {
	newMonitor := &api.FTPTransferMonitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)

	return newMonitor, err
}

func (c *ftptransfermonitors) Update(monitor *api.FTPTransferMonitor) (*api.FTPTransferMonitor, error) {
	updatedMonitor := &api.FTPTransferMonitor{}
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

func (c *ftptransfermonitors) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *ftptransfermonitors) List() ([]*api.FTPTransferMonitor, error) {
	ftptransfermonitors := []*api.FTPTransferMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&ftptransfermonitors)

	return ftptransfermonitors, err
}

func (c *ftptransfermonitors) Activate(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/activate").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *ftptransfermonitors) Suspend(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/suspend").
		ResourceID(monitorID).
		Do().
		Err()
}
