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
		Type:     schema.TypeList,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Description: "List of secondary locations for monitoring",
		// DiffSuppressFunc: func(k, secLocationsInState, secLocationsInConf string, d *schema.ResourceData) bool {
		// 	log.Println("secLocationsInConf ++++++++++++++++++ ", secLocationsInConf)
		// 	log.Println("secLocationsInState +++++++++++++++++ ", secLocationsInState)
		// 	return false
		// },
	},
	"restrict_alternate_location_polling": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Restricts polling of the resource from the selected locations alone in the Location Profile, overrides the alternate location poll logic.",
	},
}

func resourceSite24x7LocationProfile() *schema.Resource {
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
	var secondaryLocationIDs []string
	for _, secondaryLocationID := range d.Get("secondary_locations").([]interface{}) {
		secondaryLocationIDs = append(secondaryLocationIDs, secondaryLocationID.(string))
	}
	return &api.LocationProfile{
		ProfileID:                        d.Id(),
		ProfileName:                      d.Get("profile_name").(string),
		PrimaryLocation:                  d.Get("primary_location").(string),
		SecondaryLocations:               secondaryLocationIDs,
		RestrictAlternateLocationPolling: d.Get("restrict_alternate_location_polling").(bool),
	}
}

// Called during read - populates the ResourceData with the locationProfile in API response
func updateLocationProfileResourceData(d *schema.ResourceData, locationProfile *api.LocationProfile) {
	d.Set("display_name", locationProfile.ProfileName)
	d.Set("primary_location", locationProfile.PrimaryLocation)
	// sort.Strings(locationProfile.SecondaryLocations)
	d.Set("secondary_locations", locationProfile.SecondaryLocations)
	d.Set("restrict_alternate_location_polling", locationProfile.RestrictAlternateLocationPolling)
}
