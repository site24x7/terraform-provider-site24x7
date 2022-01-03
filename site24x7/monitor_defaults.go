package site24x7

import (
	"errors"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	log "github.com/sirupsen/logrus"
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

func SetNotificationProfile(client Client, d *schema.ResourceData, monitor api.Site24x7Monitor) (*api.NotificationProfile, error) {
	var notificationProfile *api.NotificationProfile
	notificationProfiles, err := client.NotificationProfiles().List()
	if err != nil {
		return nil, err
	}
	if len(notificationProfiles) == 0 {
		return nil, errors.New("Unable to find notification profiles in Site24x7. Please configure them by visiting Admin -> Configuration Profiles -> Notification Profiles")
	}
	// notification_profile_id will be set for existing resources.
	// If notification_profile_name is defined we try to find a match in Site24x7 and override notification_profile_id else raise an error.
	if _, notificationProfileNameExistsInConf := d.GetOk("notification_profile_name"); notificationProfileNameExistsInConf {
		notificationProfileNameToMatch := d.Get("notification_profile_name").(string)
		log.Println("Finding match for the notification profile name : \"" + notificationProfileNameToMatch + "\" in Site24x7")
		if notificationProfileNameToMatch != "" {
			for _, p := range notificationProfiles {
				if strings.Contains(p.ProfileName, notificationProfileNameToMatch) {
					notificationProfile = p
				}
			}
		}
		if notificationProfile == nil {
			return nil, errors.New("Unable to find notification profile matching the string : \"" + notificationProfileNameToMatch + "\" in Site24x7. Please configure a valid value for the argument \"notification_profile_name\"")
		}
		monitor.SetNotificationProfileID(notificationProfile.ProfileID)
		d.Set("notification_profile_id", notificationProfile.ProfileID)
	} else if monitor.GetNotificationProfileID() == "" { // This will be true when notification_profile_id in the configuration file is empty during resource addition.
		notificationProfile = notificationProfiles[0]
		monitor.SetNotificationProfileID(notificationProfile.ProfileID)
		d.Set("notification_profile_id", notificationProfile.ProfileID)
	}
	return notificationProfile, nil
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

func SetUserGroup(client Client, d *schema.ResourceData, monitor api.Site24x7Monitor) ([]string, error) {
	var userGroupNamesInConf []string
	var userGroupIDs []string
	userGroups, err := client.UserGroups().List()
	if err != nil {
		return nil, err
	}
	if len(userGroups) == 0 {
		return nil, errors.New("Unable to find user groups in Site24x7. Please configure them by visiting Admin -> User & Alert Management -> User Alert Group")
	}

	// If user_group_names are defined we try to find a match in Site24x7 and override user_group_ids else raise an error.
	if _, userGroupNamesExistsInConf := d.GetOk("user_group_names"); userGroupNamesExistsInConf {
		for _, userGrpName := range d.Get("user_group_names").([]interface{}) {
			userGroupNamesInConf = append(userGroupNamesInConf, userGrpName.(string))
		}
		log.Println("Finding match for the user group names : [" + strings.Join(userGroupNamesInConf, ", ") + "] in Site24x7")
		for _, userGroupName := range userGroupNamesInConf {
			if userGroupName != "" {
				for _, userGroup := range userGroups {
					if strings.Contains(userGroup.DisplayName, userGroupName) {
						userGroupIDs = append(userGroupIDs, userGroup.UserGroupID)
						log.Println("Match found for user group name : " + userGroupName + ", user group id : " + userGroup.UserGroupID)
					}
				}
			}
		}

		if len(userGroupIDs) == 0 {
			return nil, errors.New("Unable to find user group matching the List : \"" + strings.Join(userGroupNamesInConf, ", ") + "\" in Site24x7. Please configure a valid value for the argument \"user_group_names\"")
		}
		monitor.SetUserGroupIDs(userGroupIDs)
		d.Set("user_group_ids", userGroupIDs)
	} else if len(monitor.GetUserGroupIDs()) == 0 { // This will be true when user_group_ids in the configuration file is empty during resource addition.
		userGroup := userGroups[0]
		monitor.SetUserGroupIDs([]string{userGroup.UserGroupID})
		d.Set("user_group_id", []string{userGroup.UserGroupID})
	}
	return userGroupIDs, nil
}
