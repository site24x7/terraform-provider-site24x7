package site24x7

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
)

// SAMPLE POST JSON
//   {
// 	"profile_name": "location profile",
// 	"primary_location": "8",
// 	"secondary_locations": [
// 	  "106",
// 	  "20",
// 	  "113",
// 	  "94"
// 	],
// 	"restrict_alt_loc": false
//   }

var LocationProfileSchema = map[string]*schema.Schema{
	"profile_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name for the location profile.",
	},
	"primary_location": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Primary location for monitoring.",
	},
	"secondary_locations": {
		Type:     schema.TypeSet,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Description: "List of secondary locations for monitoring",
	},
	"restrict_alternate_location_polling": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Restricts polling of the resource from the selected locations alone in the Location Profile, overrides the alternate location poll logic.",
	},
	"outer_regions_location_consent": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Consent is mandatory for monitoring from countries outside the European Economic Area (EEA) and the Adequate countries. To provide your consent, set outer_regions_location_consent as true.",
	},
}

func ResourceSite24x7LocationProfile() *schema.Resource {
	return &schema.Resource{
		Create: locationProfileCreate,
		Read:   locationProfileRead,
		Update: locationProfileUpdate,
		Delete: locationProfileDelete,
		Exists: locationProfileExists,

		Schema: LocationProfileSchema,
	}
}

func locationProfileCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	locationProfile := resourceDataToLocationProfile(d)

	locationProfile, err := client.LocationProfiles().Create(locationProfile)
	if err != nil {
		return err
	}

	d.SetId(locationProfile.ProfileID)

	return nil
}

func locationProfileRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	locationProfile, err := client.LocationProfiles().Get(d.Id())
	if err != nil {
		return err
	}

	updateLocationProfileResourceData(d, locationProfile)

	return nil
}

func locationProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	locationProfile := resourceDataToLocationProfile(d)

	locationProfile, err := client.LocationProfiles().Update(locationProfile)
	if err != nil {
		return err
	}

	d.SetId(locationProfile.ProfileID)

	return nil
}

func locationProfileDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	err := client.LocationProfiles().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func locationProfileExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(Client)

	_, err := client.LocationProfiles().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToLocationProfile(d *schema.ResourceData) *api.LocationProfile {

	secondaryLocations := d.Get("secondary_locations").(*schema.Set).List()
	secondaryLocationIDs := make([]string, 0, len(secondaryLocations))
	for _, v := range secondaryLocations {
		secondaryLocationIDs = append(secondaryLocationIDs, v.(string))
	}
	return &api.LocationProfile{
		ProfileID:                        d.Id(),
		ProfileName:                      d.Get("profile_name").(string),
		PrimaryLocation:                  d.Get("primary_location").(string),
		SecondaryLocations:               secondaryLocationIDs,
		RestrictAlternateLocationPolling: d.Get("restrict_alternate_location_polling").(bool),
		LocationConsentForOuterRegions:   d.Get("outer_regions_location_consent").(bool),
	}
}

// Called during read - populates the ResourceData with the locationProfile in API response
func updateLocationProfileResourceData(d *schema.ResourceData, locationProfile *api.LocationProfile) {
	d.Set("display_name", locationProfile.ProfileName)
	d.Set("primary_location", locationProfile.PrimaryLocation)
	d.Set("secondary_locations", locationProfile.SecondaryLocations)
	d.Set("restrict_alternate_location_polling", locationProfile.RestrictAlternateLocationPolling)
	d.Set("outer_regions_location_consent", locationProfile.LocationConsentForOuterRegions)
}
