package endpoints

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type UserGroups interface {
	Get(groupID string) (*api.UserGroup, error)
	Create(group *api.UserGroup) (*api.UserGroup, error)
	Update(group *api.UserGroup) (*api.UserGroup, error)
	Delete(groupID string) error
	List() ([]*api.UserGroup, error)
}

type userGroups struct {
	client rest.Client
}

func NewUserGroups(client rest.Client) UserGroups {
	return &userGroups{
		client: client,
	}
}

func (c *userGroups) Get(groupID string) (*api.UserGroup, error) {
	userGroup := &api.UserGroup{}
	err := c.client.
		Get().
		Resource("user_groups").
		ResourceID(groupID).
		Do().
		Parse(userGroup)

	return userGroup, err
}

func (c *userGroups) Create(group *api.UserGroup) (*api.UserGroup, error) {
	newUserGroup := &api.UserGroup{}
	err := c.client.
		Post().
		Resource("user_groups").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(group).
		Do().
		Parse(newUserGroup)

	return newUserGroup, err
}

func (c *userGroups) Update(group *api.UserGroup) (*api.UserGroup, error) {
	updatedGroup := &api.UserGroup{}
	err := c.client.
		Put().
		Resource("user_groups").
		ResourceID(group.UserGroupID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(group).
		Do().
		Parse(updatedGroup)

	return updatedGroup, err
}

func (c *userGroups) Delete(groupID string) error {
	return c.client.
		Delete().
		Resource("user_groups").
		ResourceID(groupID).
		Do().
		Err()
}

func (c *userGroups) List() ([]*api.UserGroup, error) {
	userGroups := []*api.UserGroup{}
	err := c.client.
		Get().
		Resource("user_groups").
		Do().
		Parse(&userGroups)

	return userGroups, err
}
