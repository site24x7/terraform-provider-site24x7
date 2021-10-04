package site24x7

import (
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

// SAMPLE POST JSON
// {
// 	"type": "URL",
//  "profile_name": "Website Threshold Profile",
// 	"profile_type": 1,
// 	"down_location_threshold": 1,
//  "website_content_modified": {
//     "severity": 2,
//     "value": false
//   },
// 	"website_content_changes": [
// 		{
// 		"severity": 2,
// 		"comparison_operator": 1,
// 		"value": 90
// 		}
// 	],
// "response_time_threshold": {
//     "primary": [
//       {
//         "severity": 2,
//         "comparison_operator": 1,
//         "strategy": 1,
//         "value": 10000,
//         "polls_check": 1
//       }
//     ],
//     "secondary": [
//       {
//         "severity": 2,
//         "comparison_operator": 1,
//         "strategy": 1,
//         "polls_check": 5,
//         "value": 10000
//       }
//     ]
//   },
// }

var ThresholdProfileSchema = map[string]*schema.Schema{
	"profile_name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"type": {
		Type:     schema.TypeString,
		Required: true,
	},
	"profile_type": {
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      1,
		ValidateFunc: validation.IntInSlice([]int{1, 2}),
	},
	"down_location_threshold": {
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      1,
		ValidateFunc: validation.IntInSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}),
	},
	"website_content_modified": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"website_content_changes": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{2, 3}), // Trouble or Critical
				},
				"comparison_operator": {
					Type:         schema.TypeInt,
					Optional:     true,
					Default:      1,
					ValidateFunc: validation.IntInSlice([]int{1}),
				},
				"value": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntAtMost(100),
				},
			},
		},
	},
	"response_time_threshold": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"primary": {
					Type:     schema.TypeList,
					Optional: true,
					MaxItems: 2, // Thresholds can be configured only for Trouble/Critical
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"severity": {
								Type:         schema.TypeInt,
								Required:     true,
								ValidateFunc: validation.IntInSlice([]int{2, 3}), // Trouble or Critical
							},
							"comparison_operator": {
								Type:         schema.TypeInt,
								Required:     true,
								ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5}),
							},
							"value": {
								Type:     schema.TypeInt,
								Required: true,
							},
							"strategy": {
								Type:         schema.TypeInt,
								Required:     true,
								ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4}),
							},
							"polls_check": {
								Type:     schema.TypeInt,
								Required: true,
							},
						},
					},
				},
				"secondary": {
					Type:     schema.TypeList,
					Optional: true,
					MaxItems: 2, // Thresholds can be configured only for Trouble/Critical
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"severity": {
								Type:         schema.TypeInt,
								Required:     true,
								ValidateFunc: validation.IntInSlice([]int{2, 3}), // Trouble or Critical
							},
							"comparison_operator": {
								Type:         schema.TypeInt,
								Required:     true,
								ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5}),
							},
							"value": {
								Type:     schema.TypeInt,
								Required: true,
							},
							"strategy": {
								Type:         schema.TypeInt,
								Required:     true,
								ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4}),
							},
							"polls_check": {
								Type:     schema.TypeInt,
								Required: true,
							},
						},
					},
				},
			},
		},
	},
}

func resourceSite24x7ThresholdProfile() *schema.Resource {
	return &schema.Resource{
		Create: thresholdProfileCreate,
		Read:   thresholdProfileRead,
		Update: thresholdProfileUpdate,
		Delete: thresholdProfileDelete,
		Exists: thresholdProfileExists,

		Schema: ThresholdProfileSchema,
	}
}

func thresholdProfileCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	thresholdProfile := resourceDataToThresholdProfile(d)

	thresholdProfile, err := client.ThresholdProfiles().Create(thresholdProfile)
	if err != nil {
		return err
	}

	d.SetId(thresholdProfile.ProfileID)

	return nil
}

func thresholdProfileRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	thresholdProfile, err := client.ThresholdProfiles().Get(d.Id())
	if err != nil {
		return err
	}

	updateThresholdProfileResourceData(d, thresholdProfile)

	return nil
}

func thresholdProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	thresholdProfile := resourceDataToThresholdProfile(d)

	thresholdProfile, err := client.ThresholdProfiles().Update(thresholdProfile)
	if err != nil {
		return err
	}

	d.SetId(thresholdProfile.ProfileID)

	return nil
}

func thresholdProfileDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	err := client.ThresholdProfiles().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func thresholdProfileExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(Client)

	_, err := client.ThresholdProfiles().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToThresholdProfile(d *schema.ResourceData) *api.ThresholdProfile {

	var websiteContentChanges []map[string]interface{}
	if contentChangesList, ok := d.GetOk("website_content_changes"); ok {
		for _, urlContentChanges := range contentChangesList.([]interface{}) {
			urlContentChangesMap, ok := urlContentChanges.(map[string]interface{})
			if ok {
				websiteContentChanges = append(websiteContentChanges, urlContentChangesMap)
			}
		}
	}
	return &api.ThresholdProfile{
		ProfileID:              d.Id(),
		ProfileName:            d.Get("profile_name").(string),
		Type:                   d.Get("type").(string),
		ProfileType:            d.Get("profile_type").(int),
		DownLocationThreshold:  d.Get("down_location_threshold").(int),
		WebsiteContentModified: d.Get("website_content_modified").(bool),
		WebsiteContentChanges:  websiteContentChanges,
		// ResponseTimeThreshold:  d.Get("response_time_threshold"),
	}
}

// Called during read and sets thresholdProfile in API response to ResourceData
func updateThresholdProfileResourceData(d *schema.ResourceData, thresholdProfile *api.ThresholdProfile) {
	d.Set("profile_name", thresholdProfile.ProfileName)
	d.Set("type", thresholdProfile.Type)
	d.Set("profile_type", thresholdProfile.ProfileType)
	d.Set("down_location_threshold", thresholdProfile.DownLocationThreshold)
	d.Set("website_content_modified", thresholdProfile.WebsiteContentModified)
	d.Set("website_content_changes", thresholdProfile.WebsiteContentChanges)
	// d.Set("response_time_threshold", thresholdProfile.ResponseTimeThreshold)
}
