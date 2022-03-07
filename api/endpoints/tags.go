package endpoints

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type Tags interface {
	Get(tagID string) (*api.Tag, error)
	Create(tag *api.Tag) (*api.Tag, error)
	Update(tag *api.Tag) (*api.Tag, error)
	Delete(tagID string) error
	List() ([]*api.Tag, error)
}

type tags struct {
	client rest.Client
}

func NewTags(client rest.Client) Tags {
	return &tags{
		client: client,
	}
}

func (c *tags) Get(tagID string) (*api.Tag, error) {
	tag := &api.Tag{}
	err := c.client.
		Get().
		Resource("tags").
		ResourceID(tagID).
		Do().
		Parse(tag)

	return tag, err
}

func (c *tags) Create(tag *api.Tag) (*api.Tag, error) {
	newTag := &api.Tag{}
	err := c.client.
		Post().
		Resource("tags").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(tag).
		Do().
		Parse(newTag)

	return newTag, err
}

func (c *tags) Update(tag *api.Tag) (*api.Tag, error) {
	updatedTag := &api.Tag{}
	err := c.client.
		Put().
		Resource("tags").
		ResourceID(tag.TagID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(tag).
		Do().
		Parse(updatedTag)

	return updatedTag, err
}

func (c *tags) Delete(tagID string) error {
	return c.client.
		Delete().
		Resource("tags").
		ResourceID(tagID).
		Do().
		Err()
}

func (c *tags) List() ([]*api.Tag, error) {
	api.TagsListLock.Lock()
	defer api.TagsListLock.Unlock()
	var err error
	if len(api.TagsList) == 0 {
		tags := []*api.Tag{}
		err = c.client.
			Get().
			Resource("tags").
			Do().
			Parse(&tags)
		api.TagsList = tags
	}
	return api.TagsList, err
}
