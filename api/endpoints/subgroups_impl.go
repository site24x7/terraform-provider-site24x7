package endpoints

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type Subgroups interface {
	Get(groupID string) (*api.Subgroup, error)
	Create(group *api.Subgroup) (*api.Subgroup, error)
	Update(group *api.Subgroup) (*api.Subgroup, error)
	Delete(groupID string) error
	List() ([]*api.Subgroup, error)
}

type subgroups struct {
	client rest.Client
}

func NewSubgroups(client rest.Client) Subgroups {
	return &subgroups{
		client: client,
	}
}

func (c *subgroups) Get(groupID string) (*api.Subgroup, error) {
	subgroup := &api.Subgroup{}
	err := c.client.
		Get().
		Resource("subgroups").
		ResourceID(groupID).
		Do().
		Parse(subgroup)

	return subgroup, err
}

func (c *subgroups) Create(subgroup *api.Subgroup) (*api.Subgroup, error) {
	newSubgroup := &api.Subgroup{}
	err := c.client.
		Post().
		Resource("subgroups").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(subgroup).
		Do().
		Parse(newSubgroup)

	return newSubgroup, err
}

func (c *subgroups) Update(subgroup *api.Subgroup) (*api.Subgroup, error) {
	updatedGroup := &api.Subgroup{}
	err := c.client.
		Put().
		Resource("subgroups").
		ResourceID(subgroup.ID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(subgroup).
		Do().
		Parse(updatedGroup)

	return updatedGroup, err
}

func (c *subgroups) Delete(groupID string) error {
	return c.client.
		Delete().
		Resource("subgroups").
		ResourceID(groupID).
		Do().
		Err()
}

func (c *subgroups) List() ([]*api.Subgroup, error) {
	subgroups := []*api.Subgroup{}
	err := c.client.
		Get().
		Resource("subgroups").
		Do().
		Parse(&subgroups)

	return subgroups, err
}
