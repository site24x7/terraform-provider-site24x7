package site24x7

import (
	"errors"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
)

var notificationProfileDataSourceSchema = map[string]*schema.Schema{
	"name_regex": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Regular expression denoting the name of the notification profile.",
	},
	"matching_ids": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "List of notification profile IDs matching the name_regex.",
	},
	"matching_ids_and_names": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "List of notification profile IDs and names matching the name_regex.",
	},
	"profile_name": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Display name for the notification profile.",
	},
	"rca_needed": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Configuration denoting whether send root cause analysis when the monitor is down is enabled for this profile.",
	},
	"notify_after_executing_actions": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Configuration denoting whether to raise alerts for downtime only after executing the pre-configured monitor actions.",
	},
	"template_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Email template ID for notification.",
	},
	"suppress_automation": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Configuration denoting whether to stop automation from being executed on the dependent monitors.",
	},
}

func DataSourceSite24x7NotificationProfile() *schema.Resource {
	return &schema.Resource{
		Read:   notificationProfileDataSourceRead,
		Schema: notificationProfileDataSourceSchema,
	}
}

// notificationProfileDataSourceRead fetches all notificationProfile from Site24x7
func notificationProfileDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	notificationProfileList, err := client.NotificationProfiles().List()
	if err != nil {
		return err
	}

	nameRegex := d.Get("name_regex")

	var notificationProfile *api.NotificationProfile
	var matchingNotificationProfileIDs []string
	var matchingNotificationProfileIDsAndNames []string

	if nameRegex != "" {
		// (?i) - Case insensitive match
		nameRegexPattern := regexp.MustCompile("(?i)" + nameRegex.(string))
		for _, profileInfo := range notificationProfileList {
			if len(profileInfo.ProfileName) > 0 {
				if notificationProfile == nil && nameRegexPattern.MatchString(profileInfo.ProfileName) {
					notificationProfile = new(api.NotificationProfile)
					notificationProfile.ProfileID = profileInfo.ProfileID
					notificationProfile.ProfileName = profileInfo.ProfileName
					notificationProfile.RcaNeeded = profileInfo.RcaNeeded
					notificationProfile.NotifyAfterExecutingActions = profileInfo.NotifyAfterExecutingActions
					notificationProfile.SuppressAutomation = profileInfo.SuppressAutomation
					notificationProfile.TemplateID = profileInfo.TemplateID
					matchingNotificationProfileIDs = append(matchingNotificationProfileIDs, profileInfo.ProfileID)
					matchingNotificationProfileIDsAndNames = append(matchingNotificationProfileIDsAndNames, profileInfo.ProfileID+"__"+profileInfo.ProfileName)
				} else if notificationProfile != nil && nameRegexPattern.MatchString(profileInfo.ProfileName) {
					matchingNotificationProfileIDs = append(matchingNotificationProfileIDs, profileInfo.ProfileID)
					matchingNotificationProfileIDsAndNames = append(matchingNotificationProfileIDsAndNames, profileInfo.ProfileID+"__"+profileInfo.ProfileName)
				}
			}
		}
	} else {
		return errors.New("Please enter a value for the attribute name_regex!")
	}

	if notificationProfile == nil {
		return errors.New("Unable to find notification profile matching the name : \"" + d.Get("name_regex").(string))
	}

	updateNotificationProfileDataSourceResourceData(d, notificationProfile, matchingNotificationProfileIDs, matchingNotificationProfileIDsAndNames)

	return nil
}

func updateNotificationProfileDataSourceResourceData(d *schema.ResourceData, notificationProfile *api.NotificationProfile, matchingNotificationProfileIDs []string, matchingNotificationProfileIDsAndNames []string) {
	d.SetId(notificationProfile.ProfileID)
	d.Set("matching_ids", matchingNotificationProfileIDs)
	d.Set("matching_ids_and_names", matchingNotificationProfileIDsAndNames)
	d.Set("profile_name", notificationProfile.ProfileName)
	d.Set("rca_needed", notificationProfile.RcaNeeded)
	d.Set("notify_after_executing_actions", notificationProfile.NotifyAfterExecutingActions)
	d.Set("suppress_automation", notificationProfile.SuppressAutomation)
	d.Set("template_id", notificationProfile.TemplateID)
}
