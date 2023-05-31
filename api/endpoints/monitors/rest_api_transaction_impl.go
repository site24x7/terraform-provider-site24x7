package monitors

import (
	log "github.com/sirupsen/logrus"
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type RestApiTransactionMonitors interface {
	Get(monitorID string) (*api.RestApiTransactionMonitor, error)
	Create(monitor *api.RestApiTransactionMonitor) (*api.RestApiTransactionMonitor, error)
	Update(monitor *api.RestApiTransactionMonitor) (*api.RestApiTransactionMonitor, error)
	Delete(monitorID string) error
	List() ([]*api.RestApiTransactionMonitor, error)
	Activate(monitorID string) error
	Suspend(monitorID string) error
}

type restapitransactions struct {
	client rest.Client
}

func NewRestApiTransactionMonitors(client rest.Client) RestApiTransactionMonitors {
	return &restapitransactions{
		client: client,
	}
}

func (c *restapitransactions) Get(monitorID string) (*api.RestApiTransactionMonitor, error) {
	monitor := &api.RestApiTransactionMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Parse(monitor)

	return monitor, err
}

func (c *restapitransactions) Create(monitor *api.RestApiTransactionMonitor) (*api.RestApiTransactionMonitor, error) {
	log.Println("GetAPIBaseURL : ",monitor)
	newMonitor := &api.RestApiTransactionMonitor{}
	err := c.client.
		Post().
		Resource("monitors").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitor).
		Do().
		Parse(newMonitor)

	return newMonitor, err
}

func (c *restapitransactions) Update(monitor *api.RestApiTransactionMonitor) (*api.RestApiTransactionMonitor, error) {
	updatedMonitor := &api.RestApiTransactionMonitor{}
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

func (c *restapitransactions) Delete(monitorID string) error {
	return c.client.
		Delete().
		Resource("monitors").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *restapitransactions) List() ([]*api.RestApiTransactionMonitor, error) {
	restapimonitors := []*api.RestApiTransactionMonitor{}
	err := c.client.
		Get().
		Resource("monitors").
		Do().
		Parse(&restapimonitors)

	return restapimonitors, err
}

func (c *restapitransactions) Activate(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/activate").
		ResourceID(monitorID).
		Do().
		Err()
}

func (c *restapitransactions) Suspend(monitorID string) error {
	return c.client.
		Put().
		Resource("monitors/suspend").
		ResourceID(monitorID).
		Do().
		Err()
}

