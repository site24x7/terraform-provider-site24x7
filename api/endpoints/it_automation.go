package endpoints

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type URLAutomations interface {
	Get(actionID string) (*api.URLAutomation, error)
	Create(automation *api.URLAutomation) (*api.URLAutomation, error)
	Update(automation *api.URLAutomation) (*api.URLAutomation, error)
	Delete(actionID string) error
	List() ([]*api.URLAutomation, error)
}

type urlAutomations struct {
	client rest.Client
}

func NewURLAutomations(client rest.Client) URLAutomations {
	return &urlAutomations{
		client: client,
	}
}

func (c *urlAutomations) Get(actionID string) (*api.URLAutomation, error) {
	automation := &api.URLAutomation{}
	err := c.client.
		Get().
		Resource("it_automation").
		ResourceID(actionID).
		Do().
		Parse(automation)

	return automation, err
}

func (c *urlAutomations) Create(automation *api.URLAutomation) (*api.URLAutomation, error) {
	newURLAutomation := &api.URLAutomation{}
	err := c.client.
		Post().
		Resource("it_automation").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(automation).
		Do().
		Parse(newURLAutomation)

	return newURLAutomation, err
}

func (c *urlAutomations) Update(automation *api.URLAutomation) (*api.URLAutomation, error) {
	urlAutomation := &api.URLAutomation{}
	err := c.client.
		Put().
		Resource("it_automation").
		ResourceID(automation.ActionID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(automation).
		Do().
		Parse(urlAutomation)

	return urlAutomation, err
}

func (c *urlAutomations) Delete(actionID string) error {
	return c.client.
		Delete().
		Resource("it_automation").
		ResourceID(actionID).
		Do().
		Err()
}

func (c *urlAutomations) List() ([]*api.URLAutomation, error) {
	urlAutomation := []*api.URLAutomation{}
	err := c.client.
		Get().
		Resource("it_automation").
		Do().
		Parse(&urlAutomation)

	return urlAutomation, err
}
