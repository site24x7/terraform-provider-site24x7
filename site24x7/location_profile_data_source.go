package site24x7

import (
	"errors"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
)

var locationProfileDataSourceSchema = map[string]*schema.Schema{
	"name_regex": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Regular expression denoting the name of the location profile.",
	},
	"matching_ids_and_names": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "List of location profile IDs and names matching the name_regex.",
	},
	"profile_name": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Display name for the location profile.",
	},
	"primary_location": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Primary location for monitoring.",
	},
	"secondary_locations": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Computed:    true,
		Description: "List of secondary locations for monitoring.",
	},
	"outer_regions_location_consent": {
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "Attribute denoting whether consent is mandatory for monitoring from countries outside the European Economic Area (EEA) and the Adequate countries.",
	},
	"restrict_alt_loc": {
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "Restricts polling of the resource from the selected locations alone in the Location Profile, overrides the alternate location poll logic.",
	},
}

func DataSourceSite24x7LocationProfile() *schema.Resource {
	return &schema.Resource{
		Read:   locationProfileDataSourceRead,
		Schema: locationProfileDataSourceSchema,
	}
}

// locationProfileDataSourceRead fetches all locationProfile from Site24x7
func locationProfileDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	locationProfileList, err := client.LocationProfiles().List()
	if err != nil {
		return err
	}

	nameRegex := d.Get("name_regex")

	var locationProfile *api.LocationProfile
	var matchingLocationProfileIDsAndNames []string

	if nameRegex != "" {
		// (?i) - Case insensitive match
		nameRegexPattern := regexp.MustCompile("(?i)" + nameRegex.(string))
		for _, profileInfo := range locationProfileList {
			if len(profileInfo.ProfileName) > 0 {
				if locationProfile == nil && nameRegexPattern.MatchString(profileInfo.ProfileName) {
					locationProfile = new(api.LocationProfile)
					locationProfile.ProfileID = profileInfo.ProfileID
					locationProfile.ProfileName = profileInfo.ProfileName
					locationProfile.PrimaryLocation = profileInfo.PrimaryLocation
					locationProfile.SecondaryLocations = profileInfo.SecondaryLocations
					locationProfile.LocationConsentForOuterRegions = profileInfo.LocationConsentForOuterRegions
					locationProfile.RestrictAlternateLocationPolling = profileInfo.RestrictAlternateLocationPolling
					matchingLocationProfileIDsAndNames = append(matchingLocationProfileIDsAndNames, profileInfo.ProfileID+"__"+profileInfo.ProfileName)
				} else if locationProfile != nil && nameRegexPattern.MatchString(profileInfo.ProfileName) {
					matchingLocationProfileIDsAndNames = append(matchingLocationProfileIDsAndNames, profileInfo.ProfileID+"__"+profileInfo.ProfileName)
				}
			}
		}
	} else {
		return errors.New("Please enter a value for the attribute name_regex!")
	}

	if locationProfile == nil {
		return errors.New("Unable to find location profile matching the name : \"" + d.Get("name_regex").(string))
	}

	updateLocationProfileDataSourceResourceData(d, locationProfile, matchingLocationProfileIDsAndNames)

	return nil
}

func updateLocationProfileDataSourceResourceData(d *schema.ResourceData, locationProfile *api.LocationProfile, matchingLocationProfileIDsAndNames []string) {
	d.SetId(locationProfile.ProfileID)
	d.Set("matching_ids_and_names", matchingLocationProfileIDsAndNames)
	d.Set("profile_name", locationProfile.ProfileName)
	d.Set("primary_location", locationProfile.PrimaryLocation)
	d.Set("secondary_locations", locationProfile.SecondaryLocations)
	d.Set("outer_regions_location_consent", locationProfile.LocationConsentForOuterRegions)
	d.Set("restrict_alt_loc", locationProfile.RestrictAlternateLocationPolling)
}
