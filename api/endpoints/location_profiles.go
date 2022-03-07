package endpoints

import (
	"github.com/jinzhu/copier"
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type LocationProfiles interface {
	Get(profileID string) (*api.LocationProfile, error)
	Create(profile *api.LocationProfile) (*api.LocationProfile, error)
	Update(profile *api.LocationProfile) (*api.LocationProfile, error)
	Delete(profileID string) error
	List() ([]*api.LocationProfile, error)
}

type locationProfiles struct {
	client rest.Client
}

func NewLocationProfiles(client rest.Client) LocationProfiles {
	return &locationProfiles{
		client: client,
	}
}

func (c *locationProfiles) Get(profileID string) (*api.LocationProfile, error) {
	profile := &api.LocationProfile{}
	err := c.client.
		Get().
		Resource("location_profiles").
		ResourceID(profileID).
		Do().
		Parse(profile)

	return profile, err
}

func (c *locationProfiles) Create(profile *api.LocationProfile) (*api.LocationProfile, error) {
	newProfile := &api.LocationProfile{}
	err := c.client.
		Post().
		Resource("location_profiles").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(profile).
		Do().
		Parse(newProfile)

	return newProfile, err
}

func (c *locationProfiles) Update(profile *api.LocationProfile) (*api.LocationProfile, error) {
	updatedProfile := &api.LocationProfile{}
	locationProfileData := &api.LocationProfile{}
	copier.Copy(locationProfileData, profile)
	locationProfileData.ProfileID = ""
	err := c.client.
		Put().
		Resource("location_profiles").
		ResourceID(profile.ProfileID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(locationProfileData).
		Do().
		Parse(updatedProfile)

	return updatedProfile, err
}

func (c *locationProfiles) Delete(profileID string) error {
	return c.client.
		Delete().
		Resource("location_profiles").
		ResourceID(profileID).
		Do().
		Err()
}

func (c *locationProfiles) List() ([]*api.LocationProfile, error) {
	api.LocationProfilesLock.Lock()
	defer api.LocationProfilesLock.Unlock()
	var err error
	if len(api.LocationProfiles) == 0 {
		profiles := []*api.LocationProfile{}
		err = c.client.
			Get().
			Resource("location_profiles").
			Do().
			Parse(&profiles)
		api.LocationProfiles = profiles
	}
	return api.LocationProfiles, err
}
