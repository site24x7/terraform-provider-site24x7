package site24x7

import (
	"errors"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
)

var thresholdProfileDataSourceSchema = map[string]*schema.Schema{
	"name_regex": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Regular expression denoting the name of the threshold profile.",
	},
	"matching_ids": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "List of threshold profile IDs matching the name_regex.",
	},
	"matching_ids_and_names": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "List of threshold profile IDs and names matching the name_regex.",
	},
	"profile_name": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Display name for the threshold profile.",
	},
	"type": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Type of the monitor for which the threshold profile is being created.",
	},
	"profile_type": {
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "Static Threshold(1) or AI-based Threshold(2)",
	},
	"down_location_threshold": {
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "Triggers alert when the monitor is down from configured number of locations. Default value is '3'",
	},
	"website_content_modified": {
		Type:        schema.TypeBool,
		Computed:    true,
		Description: "Triggers alert when the website content is modified.",
	},
}

func DataSourceSite24x7ThresholdProfile() *schema.Resource {
	return &schema.Resource{
		Read:   thresholdProfileDataSourceRead,
		Schema: thresholdProfileDataSourceSchema,
	}
}

// thresholdProfileDataSourceRead fetches all thresholdProfile from Site24x7
func thresholdProfileDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	thresholdProfileList, err := client.ThresholdProfiles().List()
	if err != nil {
		return err
	}

	nameRegex := d.Get("name_regex")

	var thresholdProfile *api.ThresholdProfile
	var matchingThresholdProfileIDs []string
	var matchingThresholdProfileIDsAndNames []string
	if nameRegex != "" {
		// (?i) - Case insensitive match
		nameRegexPattern := regexp.MustCompile("(?i)" + nameRegex.(string))
		for _, profileInfo := range thresholdProfileList {
			if len(profileInfo.ProfileName) > 0 {
				if thresholdProfile == nil && nameRegexPattern.MatchString(profileInfo.ProfileName) {
					thresholdProfile = new(api.ThresholdProfile)
					thresholdProfile.ProfileID = profileInfo.ProfileID
					thresholdProfile.ProfileName = profileInfo.ProfileName
					thresholdProfile.Type = profileInfo.Type
					thresholdProfile.ProfileType = profileInfo.ProfileType
					thresholdProfile.DownLocationThreshold = profileInfo.DownLocationThreshold
					thresholdProfile.WebsiteContentModified = profileInfo.WebsiteContentModified
					matchingThresholdProfileIDs = append(matchingThresholdProfileIDs, profileInfo.ProfileID)
					matchingThresholdProfileIDsAndNames = append(matchingThresholdProfileIDsAndNames, profileInfo.ProfileID+"__"+profileInfo.ProfileName)
				} else if thresholdProfile != nil && nameRegexPattern.MatchString(profileInfo.ProfileName) {
					matchingThresholdProfileIDs = append(matchingThresholdProfileIDs, profileInfo.ProfileID)
					matchingThresholdProfileIDsAndNames = append(matchingThresholdProfileIDsAndNames, profileInfo.ProfileID+"__"+profileInfo.ProfileName)
				}
			}
		}
	} else {
		return errors.New("Please enter a value for the attribute name_regex!")
	}

	if thresholdProfile == nil {
		return errors.New("Unable to find threshold profile matching the name : \"" + d.Get("name_regex").(string))
	}

	updateThresholdProfileDataSourceResourceData(d, thresholdProfile, matchingThresholdProfileIDs, matchingThresholdProfileIDsAndNames)

	return nil
}

func updateThresholdProfileDataSourceResourceData(d *schema.ResourceData, thresholdProfile *api.ThresholdProfile, matchingThresholdProfileIDs []string, matchingThresholdProfileIDsAndNames []string) {
	d.SetId(thresholdProfile.ProfileID)
	d.Set("matching_ids", matchingThresholdProfileIDs)
	d.Set("matching_ids_and_names", matchingThresholdProfileIDsAndNames)
	d.Set("profile_name", thresholdProfile.ProfileName)
	d.Set("type", thresholdProfile.Type)
	d.Set("profile_type", thresholdProfile.ProfileType)
	d.Set("down_location_threshold", thresholdProfile.DownLocationThreshold)
	d.Set("website_content_modified", thresholdProfile.WebsiteContentModified)
}
