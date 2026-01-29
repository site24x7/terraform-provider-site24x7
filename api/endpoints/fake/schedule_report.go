package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/stretchr/testify/mock"
)

type ScheduleReport struct {
	mock.Mock
}

func (f *ScheduleReport) Create(sr *api.ScheduleReport) (*api.ScheduleReport, error) {
	args := f.Called(sr)
	if obj, ok := args.Get(0).(*api.ScheduleReport); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (f *ScheduleReport) Get(id string) (*api.ScheduleReport, error) {
	args := f.Called(id)
	if obj, ok := args.Get(0).(*api.ScheduleReport); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

// High-level update (used by Terraform layer)
func (f *ScheduleReport) Update(sr *api.ScheduleReport) (*api.ScheduleReport, error) {
	args := f.Called(sr)
	if obj, ok := args.Get(0).(*api.ScheduleReport); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

// Low-level raw update (used by REST layer)
func (f *ScheduleReport) UpdateRaw(id string, payload interface{}) (*api.ScheduleReport, error) {
	args := f.Called(id, payload)
	if obj, ok := args.Get(0).(*api.ScheduleReport); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (f *ScheduleReport) Delete(id string) error {
	args := f.Called(id)
	return args.Error(0)
}

func (f *ScheduleReport) List() ([]*api.ScheduleReport, error) {
	args := f.Called()
	if obj, ok := args.Get(0).([]*api.ScheduleReport); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
