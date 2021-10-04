package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.LocationTemplate = &LocationTemplate{}

type LocationTemplate struct {
	mock.Mock
}

func (e *LocationTemplate) Get() (*api.LocationTemplate, error) {
	args := e.Called()
	if obj, ok := args.Get(0).(*api.LocationTemplate); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
