package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/common"
	"github.com/stretchr/testify/mock"
)

var _ common.MilestoneMarker = &MilestoneMarker{}

type MilestoneMarker struct {
	mock.Mock
}

func (e *MilestoneMarker) Create(marker *api.MilestoneMarker) (*api.MilestoneMarker, error) {
	args := e.Called(marker)
	if obj, ok := args.Get(0).(*api.MilestoneMarker); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *MilestoneMarker) Update(marker *api.MilestoneMarker) (*api.MilestoneMarker, error) {
	args := e.Called(marker)
	if obj, ok := args.Get(0).(*api.MilestoneMarker); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *MilestoneMarker) Delete(monitorID string, markerTime string) error {
	args := e.Called(monitorID, markerTime)
	return args.Error(0)
}

func (e *MilestoneMarker) Get(monitorID string, markerTime string) (*api.MilestoneMarker, error) {
	args := e.Called(monitorID, markerTime)
	if obj, ok := args.Get(0).(*api.MilestoneMarker); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
