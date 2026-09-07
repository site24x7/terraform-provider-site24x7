package apm

import (
	"fmt"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

// DefaultTimeWindow is used when a caller does not specify one. The APM
// Insight reporting endpoints require a time window in the path, but the
// identity and configuration attributes this provider reads are the same
// whichever window is requested, so the shortest one is the cheapest default.
const DefaultTimeWindow = "H"

// APMApplications is the interface for the APM Insight application endpoints.
type APMApplications interface {
	List(timeWindow string) ([]*api.APMApplication, error)
	Get(applicationID, timeWindow string) (*api.APMApplication, error)
	Manage(applicationID string) error
	Unmanage(applicationID string) error
	Delete(applicationID string) error
}

type apmApplications struct {
	client rest.Client
}

// NewAPMApplications creates an APMApplications endpoint client.
func NewAPMApplications(client rest.Client) APMApplications {
	return &apmApplications{
		client: client,
	}
}

func (c *apmApplications) List(timeWindow string) ([]*api.APMApplication, error) {
	applications := []*api.APMApplication{}
	err := c.client.
		Get().
		Resource("apminsight/app").
		ResourceID(timeWindowOrDefault(timeWindow)).
		Do().
		Parse(&applications)

	return applications, err
}

func (c *apmApplications) Get(applicationID, timeWindow string) (*api.APMApplication, error) {
	application := &api.APMApplication{}
	err := c.client.
		Get().
		Resource("apminsight/app").
		ResourceID(fmt.Sprintf("%s/%s", applicationID, timeWindowOrDefault(timeWindow))).
		Do().
		Parse(application)

	return application, err
}

func (c *apmApplications) Manage(applicationID string) error {
	return c.client.
		Post().
		Resource("apminsight/app").
		ResourceID(fmt.Sprintf("%s/manage", applicationID)).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Do().
		Err()
}

func (c *apmApplications) Unmanage(applicationID string) error {
	return c.client.
		Post().
		Resource("apminsight/app").
		ResourceID(fmt.Sprintf("%s/unmanage", applicationID)).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Do().
		Err()
}

func (c *apmApplications) Delete(applicationID string) error {
	return c.client.
		Delete().
		Resource("apminsight/app").
		ResourceID(applicationID).
		Do().
		Err()
}

func timeWindowOrDefault(timeWindow string) string {
	if timeWindow == "" {
		return DefaultTimeWindow
	}
	return timeWindow
}
