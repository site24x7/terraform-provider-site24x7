package endpoints

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type URLActions interface {
	Get(actionID string) (*api.URLAction, error)
	Create(automation *api.URLAction) (*api.URLAction, error)
	Update(automation *api.URLAction) (*api.URLAction, error)
	Delete(actionID string) error
	List() ([]*api.URLAction, error)
}

type urlActions struct {
	client rest.Client
}

func NewURLActions(client rest.Client) URLActions {
	return &urlActions{
		client: client,
	}
}

func (c *urlActions) Get(actionID string) (*api.URLAction, error) {
	automation := &api.URLAction{}
	err := c.client.
		Get().
		Resource("it_automation").
		ResourceID(actionID).
		Do().
		Parse(automation)

	return automation, err
}

func (c *urlActions) Create(automation *api.URLAction) (*api.URLAction, error) {
	newURLAction := &api.URLAction{}
	err := c.client.
		Post().
		Resource("it_automation").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(automation).
		Do().
		Parse(newURLAction)

	return newURLAction, err
}

func (c *urlActions) Update(automation *api.URLAction) (*api.URLAction, error) {
	urlAction := &api.URLAction{}
	err := c.client.
		Put().
		Resource("it_automation").
		ResourceID(automation.ActionID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(automation).
		Do().
		Parse(urlAction)

	return urlAction, err
}

func (c *urlActions) Delete(actionID string) error {
	return c.client.
		Delete().
		Resource("it_automation").
		ResourceID(actionID).
		Do().
		Err()
}

func (c *urlActions) List() ([]*api.URLAction, error) {
	urlAction := []*api.URLAction{}
	err := c.client.
		Get().
		Resource("it_automation").
		Do().
		Parse(&urlAction)

	return urlAction, err
}
