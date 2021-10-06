package site24x7

import (
	"errors"
	"strings"

	"github.com/site24x7/terraform-provider-site24x7/api"
)

// DefaultLocationProfile fetches all location profiles from the server
// and tries to find a match for the input profile name. If no match is found
// the first location profile from the list is returned. If no location profiles are configured,
// DefaultLocationProfile will return an error.
func DefaultLocationProfile(client Client, profileNameToMatch string) (*api.LocationProfile, error) {
	locationProfiles, err := client.LocationProfiles().List()
	if err != nil {
		return nil, err
	}

	if len(locationProfiles) == 0 {
		return nil, errors.New("No Location Profiles Configured")
	}

	if profileNameToMatch != "" {
		for _, p := range locationProfiles {
			if strings.Contains(p.ProfileName, profileNameToMatch) {
				return p, nil
			}
		}
	}

	return locationProfiles[0], nil
}

// DefaultNotificationProfile fetches the first notification profile returned by the
// client. If no notification profiles are configured, DefaultNotificationProfile will
// return an error.
func DefaultNotificationProfile(client Client) (*api.NotificationProfile, error) {
	profiles, err := client.NotificationProfiles().List()

	if err != nil {
		return nil, err
	}

	if len(profiles) == 0 {
		return nil, errors.New("No Notification Profiles Configured")
	}

	return profiles[0], nil
}

// DefaultThresholdProfile fetches all threshold profiles from the server
// and tries to match threshold profile type and the given monitor type.
// If no match is found the first threshold profile from the list is returned.
// If no threshold profiles are configured, DefaultThresholdProfile will return an error.
func DefaultThresholdProfile(client Client, monitorType api.MonitorType) (*api.ThresholdProfile, error) {
	profiles, err := client.ThresholdProfiles().List()
	if err != nil {
		return nil, err
	}

	if len(profiles) == 0 {
		return nil, errors.New("No Threshold Profiles Configured")
	}

	for _, p := range profiles {
		if p.Type == string(monitorType) {
			return p, nil
		}
	}

	return profiles[0], nil
}

// DefaultUserGroup fetches the first user group returned by the
// client. If no user groups are configured, DefaultUserGroup will
// return an error.
func DefaultUserGroup(client Client) (*api.UserGroup, error) {
	userGroups, err := client.UserGroups().List()
	if err != nil {
		return nil, err
	}

	if len(userGroups) == 0 {
		return nil, errors.New("No User Groups Configured")
	}

	return userGroups[0], nil
}
