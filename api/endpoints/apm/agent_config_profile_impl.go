package apm

import (
	"fmt"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

// The write and read endpoints differ in more than the verb: writes go to the
// singular "agent_config_profile" and reads to the plural
// "agent_config_profiles". Keeping them as named constants avoids silently
// pluralising the wrong one.
const (
	agentConfigProfileResource        = "apminsight/agent_config_profile"
	agentConfigProfilesResource       = "apminsight/agent_config_profiles"
	defaultAgentConfigProfileResource = "apminsight/default_agent_config_profile"
)

// APMAgentConfigProfiles is the interface for the APM Insight agent
// configuration profile endpoints.
type APMAgentConfigProfiles interface {
	Get(profileID string) (*api.APMAgentConfigProfile, error)
	GetDefault(agentType string) (*api.APMAgentConfigProfile, error)
	Create(profile *api.APMAgentConfigProfile) (*api.APMAgentConfigProfile, error)
	Update(profile *api.APMAgentConfigProfile) (*api.APMAgentConfigProfile, error)
	Delete(profileID string) error
	List() ([]*api.APMAgentConfigProfile, error)
	ListByAgentType(agentType string) ([]*api.APMAgentConfigProfile, error)
}

type apmAgentConfigProfiles struct {
	client rest.Client
}

// NewAPMAgentConfigProfiles creates an APMAgentConfigProfiles endpoint client.
func NewAPMAgentConfigProfiles(client rest.Client) APMAgentConfigProfiles {
	return &apmAgentConfigProfiles{
		client: client,
	}
}

// Get resolves a profile by ID.
//
// The API has no get-by-profile-id endpoint - only list-all, list-by-agent-type
// and get-default - so this lists and filters, the same approach
// AttributeAlertGroup takes for the same reason.
//
// A missing profile is reported as a 404 StatusError rather than a plain one so
// that callers can tell it apart with apierrors.IsNotFound and remove the
// profile from Terraform state instead of failing the refresh.
func (c *apmAgentConfigProfiles) Get(profileID string) (*api.APMAgentConfigProfile, error) {
	profiles, err := c.List()
	if err != nil {
		return nil, err
	}

	for _, profile := range profiles {
		if profile.ProfileID == profileID {
			return profile, nil
		}
	}

	return nil, apierrors.NewStatusError(
		http.StatusNotFound,
		fmt.Sprintf("APM agent configuration profile not found for profile_id: %s", profileID),
	)
}

func (c *apmAgentConfigProfiles) GetDefault(agentType string) (*api.APMAgentConfigProfile, error) {
	profile := &api.APMAgentConfigProfile{}
	err := c.client.
		Get().
		Resource(defaultAgentConfigProfileResource).
		ResourceID(agentType).
		Do().
		Parse(profile)

	return profile, err
}

func (c *apmAgentConfigProfiles) Create(profile *api.APMAgentConfigProfile) (*api.APMAgentConfigProfile, error) {
	created := &api.APMAgentConfigProfile{}
	err := c.client.
		Post().
		Resource(agentConfigProfileResource).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(profile).
		Do().
		Parse(created)

	return created, err
}

func (c *apmAgentConfigProfiles) Update(profile *api.APMAgentConfigProfile) (*api.APMAgentConfigProfile, error) {
	updated := &api.APMAgentConfigProfile{}
	err := c.client.
		Put().
		Resource(agentConfigProfileResource).
		ResourceID(profile.ProfileID).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(profile).
		Do().
		Parse(updated)

	return updated, err
}

// Delete discards the response body, which is a {name, id} pair rather than a
// profile.
func (c *apmAgentConfigProfiles) Delete(profileID string) error {
	return c.client.
		Delete().
		Resource(agentConfigProfileResource).
		ResourceID(profileID).
		Do().
		Err()
}

func (c *apmAgentConfigProfiles) List() ([]*api.APMAgentConfigProfile, error) {
	profiles := []*api.APMAgentConfigProfile{}
	err := c.client.
		Get().
		Resource(agentConfigProfilesResource).
		Do().
		Parse(&profiles)

	return profiles, err
}

func (c *apmAgentConfigProfiles) ListByAgentType(agentType string) ([]*api.APMAgentConfigProfile, error) {
	profiles := []*api.APMAgentConfigProfile{}
	err := c.client.
		Get().
		Resource(agentConfigProfilesResource).
		ResourceID(agentType).
		Do().
		Parse(&profiles)

	return profiles, err
}
