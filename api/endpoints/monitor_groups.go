package endpoints

import (
	"github.com/jinzhu/copier"
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type MonitorGroups interface {
	Get(groupID string) (*api.MonitorGroup, error)
	Create(group *api.MonitorGroup) (*api.MonitorGroup, error)
	Update(group *api.MonitorGroup) (*api.MonitorGroup, error)
	Delete(groupID string) error
	List() ([]*api.MonitorGroup, error)
}

type monitorGroups struct {
	client rest.Client
}

func NewMonitorGroups(client rest.Client) MonitorGroups {
	return &monitorGroups{
		client: client,
	}
}

func (c *monitorGroups) Get(groupID string) (*api.MonitorGroup, error) {
	monitorGroup := &api.MonitorGroup{}
	err := c.client.
		Get().
		Resource("monitor_groups").
		ResourceID(groupID).
		Do().
		Parse(monitorGroup)

	return monitorGroup, err
}

func (c *monitorGroups) Create(group *api.MonitorGroup) (*api.MonitorGroup, error) {
	newMonitorGroup := &api.MonitorGroup{}
	err := c.client.
		Post().
		Resource("monitor_groups").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(group).
		Do().
		Parse(newMonitorGroup)

	return newMonitorGroup, err
}

func (c *monitorGroups) Update(group *api.MonitorGroup) (*api.MonitorGroup, error) {
	updatedGroup := &api.MonitorGroup{}
	monitorGroupData := &api.MonitorGroup{}
	copier.Copy(monitorGroupData, group)
	monitorGroupData.GroupID = ""
	err := c.client.
		Put().
		Resource("monitor_groups").
		ResourceID(group.GroupID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(monitorGroupData).
		Do().
		Parse(updatedGroup)

	return updatedGroup, err
}

func (c *monitorGroups) Delete(groupID string) error {
	return c.client.
		Delete().
		Resource("monitor_groups").
		ResourceID(groupID).
		Do().
		Err()
}

func (c *monitorGroups) List() ([]*api.MonitorGroup, error) {
	monitorGroups := []*api.MonitorGroup{}
	err := c.client.
		Get().
		Resource("monitor_groups").
		Do().
		Parse(&monitorGroups)

	return monitorGroups, err
}
