package apm

import (
	"fmt"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

// APMInstances is the interface for the APM Insight instance endpoints.
type APMInstances interface {
	List(timeWindow string) ([]*api.APMInstance, error)
	Get(instanceID, timeWindow string) (*api.APMInstance, error)
	Manage(instanceID string) error
	Unmanage(instanceID string) error
	Delete(instanceID string) error
}

type apmInstances struct {
	client rest.Client
}

// NewAPMInstances creates an APMInstances endpoint client.
func NewAPMInstances(client rest.Client) APMInstances {
	return &apmInstances{
		client: client,
	}
}

func (c *apmInstances) List(timeWindow string) ([]*api.APMInstance, error) {
	instances := []*api.APMInstance{}
	err := c.client.
		Get().
		Resource("apminsight/ins").
		ResourceID(timeWindowOrDefault(timeWindow)).
		Do().
		Parse(&instances)

	return instances, err
}

func (c *apmInstances) Get(instanceID, timeWindow string) (*api.APMInstance, error) {
	instance := &api.APMInstance{}
	err := c.client.
		Get().
		Resource("apminsight/ins").
		ResourceID(fmt.Sprintf("%s/%s", instanceID, timeWindowOrDefault(timeWindow))).
		Do().
		Parse(instance)

	return instance, err
}

func (c *apmInstances) Manage(instanceID string) error {
	return c.client.
		Post().
		Resource("apminsight/ins").
		ResourceID(fmt.Sprintf("%s/manage", instanceID)).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Do().
		Err()
}

func (c *apmInstances) Unmanage(instanceID string) error {
	return c.client.
		Post().
		Resource("apminsight/ins").
		ResourceID(fmt.Sprintf("%s/unmanage", instanceID)).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Do().
		Err()
}

func (c *apmInstances) Delete(instanceID string) error {
	return c.client.
		Delete().
		Resource("apminsight/ins").
		ResourceID(instanceID).
		Do().
		Err()
}
