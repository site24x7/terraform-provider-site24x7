package endpoints

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type OpsgenieIntegration interface {
	Get(opsgenieIntegrationID string) (*api.OpsgenieIntegration, error)
	Create(opsgenieIntegration *api.OpsgenieIntegration) (*api.OpsgenieIntegration, error)
	Update(opsgenieIntegration *api.OpsgenieIntegration) (*api.OpsgenieIntegration, error)
}

type opsgenie struct {
	client rest.Client
}

func NewOpsgenie(client rest.Client) OpsgenieIntegration {
	return &opsgenie{
		client: client,
	}
}

func (o *opsgenie) Get(opsgenieIntegrationID string) (*api.OpsgenieIntegration, error) {
	opsgenieIntegration := &api.OpsgenieIntegration{}
	err := o.client.
		Get().
		Resource("integration/opsgenie").
		ResourceID(opsgenieIntegrationID).
		Do().
		Parse(opsgenieIntegration)

	return opsgenieIntegration, err
}

func (o *opsgenie) Create(opsgenieIntegration *api.OpsgenieIntegration) (*api.OpsgenieIntegration, error) {
	newOpsgenieIntegration := &api.OpsgenieIntegration{}
	err := o.client.
		Post().
		Resource("integration/opsgenie").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(opsgenieIntegration).
		Do().
		Parse(newOpsgenieIntegration)

	return newOpsgenieIntegration, err
}

func (o *opsgenie) Update(opsgenieIntegration *api.OpsgenieIntegration) (*api.OpsgenieIntegration, error) {
	updatedOpsgenieIntegration := &api.OpsgenieIntegration{}
	err := o.client.
		Put().
		Resource("integration/opsgenie").
		ResourceID(opsgenieIntegration.ServiceID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(opsgenieIntegration).
		Do().
		Parse(updatedOpsgenieIntegration)

	return updatedOpsgenieIntegration, err
}
