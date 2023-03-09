package endpoints

import (
	"github.com/jinzhu/copier"
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type Users interface {
	Get(groupID string) (*api.User, error)
	Create(group *api.User) (*api.User, error)
	Update(group *api.User) (*api.User, error)
	Delete(groupID string) error
	List() ([]*api.User, error)
}

type users struct {
	client rest.Client
}

func NewUsers(client rest.Client) Users {
	return &users{
		client: client,
	}
}

func (c *users) Get(groupID string) (*api.User, error) {
	user := &api.User{}
	err := c.client.
		Get().
		Resource("users").
		ResourceID(groupID).
		Do().
		Parse(user)

	return user, err
}

func (c *users) Create(group *api.User) (*api.User, error) {
	newUser := &api.User{}
	err := c.client.
		Post().
		Resource("users").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(group).
		Do().
		Parse(newUser)

	return newUser, err
}

func (c *users) Update(user *api.User) (*api.User, error) {
	updatedUser := &api.User{}
	userData := &api.User{}
	copier.Copy(userData, user)
	userData.ID = ""
	err := c.client.
		Put().
		Resource("users").
		ResourceID(user.ID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(userData).
		Do().
		Parse(updatedUser)

	return updatedUser, err
}

func (c *users) Delete(groupID string) error {
	return c.client.
		Delete().
		Resource("users").
		ResourceID(groupID).
		Do().
		Err()
}

func (c *users) List() ([]*api.User, error) {
	users := []*api.User{}
	err := c.client.
		Get().
		Resource("users").
		Do().
		Parse(&users)

	return users, err
}
