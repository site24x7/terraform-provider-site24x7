package endpoints

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type ThresholdProfiles interface {
	Get(profileID string) (*api.ThresholdProfile, error)
	Create(profile *api.ThresholdProfile) (*api.ThresholdProfile, error)
	Update(profile *api.ThresholdProfile) (*api.ThresholdProfile, error)
	Delete(profileID string) error
	List() ([]*api.ThresholdProfile, error)
}

type thresholdProfiles struct {
	client rest.Client
}

func NewThresholdProfiles(client rest.Client) ThresholdProfiles {
	return &thresholdProfiles{
		client: client,
	}
}

func (c *thresholdProfiles) Get(profileID string) (*api.ThresholdProfile, error) {
	profile := &api.ThresholdProfile{}
	err := c.client.
		Get().
		Resource("threshold_profiles").
		SetHeader("Accept", "application/json; version=2.1").
		ResourceID(profileID).
		Do().
		Parse(profile)

	return profile, err
}

func (c *thresholdProfiles) Create(profile *api.ThresholdProfile) (*api.ThresholdProfile, error) {
	newThresholdProfile := &api.ThresholdProfile{}
	err := c.client.
		Post().
		Resource("threshold_profiles").
		SetHeader("Accept", "application/json; version=2.1"). // We are overwriting the default API version 2.1 for HeartBeat monitor threshold profile addition to work.
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(profile).
		Do().
		Parse(newThresholdProfile)

	return newThresholdProfile, err
}

func (c *thresholdProfiles) Update(profile *api.ThresholdProfile) (*api.ThresholdProfile, error) {
	updatedThresholdProfile := &api.ThresholdProfile{}
	err := c.client.
		Put().
		Resource("threshold_profiles").
		ResourceID(profile.ProfileID).
		SetHeader("Accept", "application/json; version=2.1"). // We are overwriting the default API version 2.1 for HeartBeat monitor threshold profile addition to work.
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(profile).
		Do().
		Parse(updatedThresholdProfile)

	return updatedThresholdProfile, err
}

func (c *thresholdProfiles) Delete(profileID string) error {
	return c.client.
		Delete().
		Resource("threshold_profiles").
		ResourceID(profileID).
		Do().
		Err()
}

func (c *thresholdProfiles) List() ([]*api.ThresholdProfile, error) {
	api.ThresholdProfilesLock.Lock()
	defer api.ThresholdProfilesLock.Unlock()
	var err error
	if len(api.ThresholdProfiles) == 0 {
		thresholdProfiles := []*api.ThresholdProfile{}
		err = c.client.
			Get().
			Resource("threshold_profiles").
			Do().
			Parse(&thresholdProfiles)
		api.ThresholdProfiles = thresholdProfiles
	}
	return api.ThresholdProfiles, err
}
