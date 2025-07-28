package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
)

// AWSExternalID is a fake implementation of the aws.AWSExternalID interface.
type AWSExternalID struct{}

// Get returns a mocked AWSExternalID response.
func (f *AWSExternalID) Get() (*api.AWSExternalID, error) {
	return &api.AWSExternalID{
		ID: "mocked-external-id",
	}, nil
}
