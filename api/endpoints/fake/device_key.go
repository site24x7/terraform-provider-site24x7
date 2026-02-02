package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/common"
	"github.com/stretchr/testify/mock"
)

// Ensuring the mock struct implements the interface
var _ common.DeviceKey = &DeviceKey{}

// DeviceKey is the mock implementation of the DeviceKey interface
type DeviceKey struct {
	mock.Mock
}

// Get retrieves a DeviceKey
func (c *DeviceKey) Get() (*api.DeviceKey, error) {
	args := c.Called()
	if obj, ok := args.Get(0).(*api.DeviceKey); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
