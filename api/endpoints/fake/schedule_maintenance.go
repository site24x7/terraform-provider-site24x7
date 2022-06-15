package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/common"
	"github.com/stretchr/testify/mock"
)

var _ common.ScheduleMaintenance = &ScheduleMaintenance{}

type ScheduleMaintenance struct {
	mock.Mock
}

func (w *ScheduleMaintenance) Get(webhookIntegrationID string) (*api.ScheduleMaintenance, error) {
	args := w.Called(webhookIntegrationID)
	if obj, ok := args.Get(0).(*api.ScheduleMaintenance); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (w *ScheduleMaintenance) Create(webhookIntegration *api.ScheduleMaintenance) (*api.ScheduleMaintenance, error) {
	args := w.Called(webhookIntegration)
	if obj, ok := args.Get(0).(*api.ScheduleMaintenance); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *ScheduleMaintenance) Delete(profileID string) error {
	args := e.Called(profileID)
	return args.Error(0)
}

func (w *ScheduleMaintenance) Update(webhookIntegration *api.ScheduleMaintenance) (*api.ScheduleMaintenance, error) {
	args := w.Called(webhookIntegration)
	if obj, ok := args.Get(0).(*api.ScheduleMaintenance); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (e *ScheduleMaintenance) List() ([]*api.ScheduleMaintenance, error) {
	args := e.Called()
	if obj, ok := args.Get(0).([]*api.ScheduleMaintenance); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
