package common

import (
	"fmt"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type AttributeAlertGroup interface {
	Get(groupID string) (*api.AttributeAlertGroup, error)
	Create(group *api.AttributeAlertGroup) (*api.AttributeAlertGroup, error)
	Update(group *api.AttributeAlertGroup) (*api.AttributeAlertGroup, error)
	Delete(groupID string) error
	List() ([]*api.AttributeAlertGroup, error)
}

type attributealertgroup struct {
	client rest.Client
}

func NewAttributeAlertGroup(client rest.Client) AttributeAlertGroup {
	return &attributealertgroup{
		client: client,
	}
}

func (c *attributealertgroup) Get(groupID string) (*api.AttributeAlertGroup, error) {
	// The API does not provide a direct GET-by-ID endpoint.
	// We list all and filter by group_id.
	groups, err := c.List()
	if err != nil {
		return nil, err
	}

	for _, g := range groups {
		if g.GroupID == groupID {
			return g, nil
		}
	}

	return nil, fmt.Errorf("attribute alert group not found for group_id: %s", groupID)
}

func (c *attributealertgroup) Create(group *api.AttributeAlertGroup) (*api.AttributeAlertGroup, error) {
	newGroup := &api.AttributeAlertGroup{}
	err := c.client.
		Post().
		Resource("attribute_groups").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(group).
		Do().
		Parse(newGroup)

	return newGroup, err
}

func (c *attributealertgroup) Update(group *api.AttributeAlertGroup) (*api.AttributeAlertGroup, error) {
	updatedGroup := &api.AttributeAlertGroup{}
	err := c.client.
		Put().
		Resource("attribute_groups").
		ResourceID(group.GroupID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(group).
		Do().
		Parse(updatedGroup)

	return updatedGroup, err
}

func (c *attributealertgroup) Delete(groupID string) error {
	return c.client.
		Delete().
		Resource("attribute_groups").
		ResourceID(groupID).
		Do().
		Err()
}

func (c *attributealertgroup) List() ([]*api.AttributeAlertGroup, error) {
	groups := []*api.AttributeAlertGroup{}
	err := c.client.
		Get().
		Resource("attribute_groups").
		Do().
		Parse(&groups)

	return groups, err
}
