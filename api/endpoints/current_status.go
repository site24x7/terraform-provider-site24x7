package endpoints

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type CurrentStatus interface {
	Get(monitorID string) (*api.MonitorStatus, error)
	ListGroup(groupID string) (*api.MonitorsStatus, error)
	ListType(monitorType string) (*api.MonitorsStatus, error)
	List(options *api.CurrentStatusListOptions) (*api.MonitorsStatus, error)
}

type currentStatus struct {
	client rest.Client
}

func NewCurrentStatus(client rest.Client) CurrentStatus {
	return &currentStatus{
		client: client,
	}
}

func (c *currentStatus) Get(monitorID string) (*api.MonitorStatus, error) {
	status := &api.MonitorStatus{}
	err := c.client.
		Get().
		Resource("current_status").
		ResourceID(monitorID).
		Do().
		Parse(status)

	return status, err
}

func (c *currentStatus) ListGroup(groupID string) (*api.MonitorsStatus, error) {
	status := &api.MonitorsStatus{}
	err := c.client.
		Get().
		Resource("current_status/group").
		ResourceID(groupID).
		Do().
		Parse(status)

	return status, err
}

func (c *currentStatus) ListType(monitorType string) (*api.MonitorsStatus, error) {
	status := &api.MonitorsStatus{}
	err := c.client.
		Get().
		Resource("current_status/type").
		ResourceID(monitorType).
		Do().
		Parse(status)

	return status, err
}

func (c *currentStatus) List(options *api.CurrentStatusListOptions) (*api.MonitorsStatus, error) {
	status := &api.MonitorsStatus{}
	err := c.client.
		Get().
		Resource("current_status").
		QueryParams(options).
		Do().
		Parse(status)

	return status, err
}
