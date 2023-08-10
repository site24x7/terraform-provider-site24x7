package aws

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type AWSExternalID interface {
	Get() (*api.AWSExternalID, error)
}

type awsexternalid struct {
	client rest.Client
}

func NewAWSExternalID(client rest.Client) AWSExternalID {
	return &awsexternalid{
		client: client,
	}
}

func (c *awsexternalid) Get() (*api.AWSExternalID, error) {
	externalID := &api.AWSExternalID{}
	err := c.client.
		Get().
		Resource("aws/external_id").
		Do().
		Parse(externalID)
	return externalID, err
}
