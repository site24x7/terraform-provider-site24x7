package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints"
	"github.com/stretchr/testify/mock"
)

var _ endpoints.CurrentStatus = &CurrentStatus{}

type CurrentStatus struct {
	mock.Mock
}

func (e *CurrentStatus) Get(monitorID string) (*api.MonitorStatus, error) {
	args := e.Called(monitorID)
	if obj, ok := args.Get(0).(*api.MonitorStatus); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *CurrentStatus) ListGroup(groupID string) (*api.MonitorsStatus, error) {
	args := e.Called(groupID)
	if obj, ok := args.Get(0).(*api.MonitorsStatus); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *CurrentStatus) ListType(monitorType string) (*api.MonitorsStatus, error) {
	args := e.Called(monitorType)
	if obj, ok := args.Get(0).(*api.MonitorsStatus); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *CurrentStatus) List(options *api.CurrentStatusListOptions) (*api.MonitorsStatus, error) {
	args := e.Called(options)
	if obj, ok := args.Get(0).(*api.MonitorsStatus); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
