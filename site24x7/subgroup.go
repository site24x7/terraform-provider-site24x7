package site24x7

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
)

// POST data :
// {
// 	"display_name": "a - subgroup",
// 	"description": "description",
// 	"health_threshold_count": 1,
// 	"monitors": [
// 	  "123456000000880011",
// 	  "123456000024411005",
// 	  "123456000007534005"
// 	],
// 	"group_type": 2,
// 	"parent_group_id": "123456000033743001",
// 	"top_group_id": "123456000033743001"
//   }

var SubgroupSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display Name for the Subgroup.",
	},
	"parent_group_id": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Unique ID of the parent group under which subgroup has to be configured. It can be a subgroup or Monitor group. (In case of level 1 subgroup, top_group_id is monitor group id. In other cases it will be subgroup id. You can get the subgroup Ids configured for top_group_id by using business view API).",
	},
	"top_group_id": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Unique ID of the top monitor group for which business view has been configured.",
	},
	"description": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Description for the Subgroup.",
	},
	"health_threshold_count": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     1,
		Description: "Number of monitors' health that decide the group status. ‘0’ implies that all the monitors are considered for determining the group status. Default value is 1.",
	},
	"monitors": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "Monitors to be associated with the Subgroup.",
	},
	"group_type": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     1,
		Description: "Denotes the type of monitors that can be associated. ‘1’ implies that all type of monitors can be associated with this subgroup. Default value is 1. '2' - Web, '3' - Port/Ping, '4' - Server, '5' - Database, '6' - Synthetic Transaction, '7' - Web API, '8' - APM Insight,'9' - Network Devices, '10' - RUM, '11' - AppLogs Monitor",
	},
	"healthcheck_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Threshold profile to be associated with the Subgroup.",
	},
	"notification_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Notification profile to be associated with the Subgroup.",
	},
	"user_group_ids": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Computed:    true,
		Description: "The user groups to be notified during an outage. Mandatory, if on_call_schedule_id is not chosen.",
	},
	"on_call_schedule_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "On-Call Schedule of your choice. Mandatory, if user_group_ids is not chosen.",
	},
	"third_party_service_ids": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of Third Party Service IDs to be associated to the Subgroup.",
	},
	"tag_ids": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Computed:    true,
		Description: "List of tag IDs to be associated to the Subgroup.",
	},
	"check_frequency": {
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "Health Check polling interval.",
	},
}

func ResourceSite24x7Subgroup() *schema.Resource {
	return &schema.Resource{
		Create: subgroupCreate,
		Read:   subgroupRead,
		Update: subgroupUpdate,
		Delete: subgroupDelete,
		Exists: subgroupExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: SubgroupSchema,
	}
}

func subgroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	subgroup := resourceDataToSubgroup(d, client)

	subgroup, err := client.Subgroups().Create(subgroup)
	if err != nil {
		return err
	}

	d.SetId(subgroup.ID)

	// Read is called for updating state after modification
	// https://www.terraform.io/docs/extend/best-practices/detecting-drift.html#update-state-after-modification
	// return subgroupRead(d, meta)
	return nil
}

func subgroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	subgroup, err := client.Subgroups().Get(d.Id())
	if err != nil {
		return err
	}

	updateSubgroupResourceData(d, subgroup)

	return nil
}

func subgroupUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	subgrp := resourceDataToSubgroup(d, client)

	subgroup, err := client.Subgroups().Update(subgrp)
	if err != nil {
		return err
	}

	d.SetId(subgroup.ID)

	// Read is called for updating state after modification
	// https://www.terraform.io/docs/extend/best-practices/detecting-drift.html#update-state-after-modification
	// return subgroupRead(d, meta)
	return nil
}

func subgroupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	err := client.Subgroups().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func subgroupExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(Client)

	_, err := client.Subgroups().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToSubgroup(d *schema.ResourceData, client Client) *api.Subgroup {

	monitors := d.Get("monitors").(*schema.Set).List()
	monitorIDs := make([]string, 0, len(monitors))
	for _, monitorID := range monitors {
		if monitorID != nil {
			monitorIDs = append(monitorIDs, monitorID.(string))
		}
	}

	var userGroupIDs []string
	for _, id := range d.Get("user_group_ids").([]interface{}) {
		if id != nil {
			userGroupIDs = append(userGroupIDs, id.(string))
		}
	}

	var tagIDs []string
	for _, id := range d.Get("tag_ids").(*schema.Set).List() {
		if id != nil {
			tagIDs = append(tagIDs, id.(string))
		}
	}

	var thirdPartyServiceIDs []string
	for _, id := range d.Get("third_party_service_ids").([]interface{}) {
		if id != nil {
			thirdPartyServiceIDs = append(thirdPartyServiceIDs, id.(string))
		}
	}

	// Setting default notification profile
	notificationProfileID := d.Get("notification_profile_id").(string)
	if notificationProfileID == "" {
		notificationProfiles, err := client.NotificationProfiles().List()
		if err == nil && len(notificationProfiles) > 0 {
			notificationProfileID = notificationProfiles[0].ProfileID
		}
	}

	// Setting default health check profile
	healthCheckProfileID := d.Get("healthcheck_profile_id").(string)
	if healthCheckProfileID == "" {
		monitorType := "HEALTHCHECK"
		profiles, err := client.ThresholdProfiles().List()
		if err == nil {
			var thresholdProf *api.ThresholdProfile
			for _, p := range profiles {
				if p.Type == monitorType {
					thresholdProf = p
				}
			}
			if thresholdProf != nil {
				healthCheckProfileID = thresholdProf.ProfileID
			}
		}
	}

	return &api.Subgroup{
		ID:                    d.Id(),
		DisplayName:           d.Get("display_name").(string),
		Description:           d.Get("description").(string),
		TopGroupID:            d.Get("top_group_id").(string),
		ParentGroupID:         d.Get("parent_group_id").(string),
		Type:                  d.Get("group_type").(int),
		Monitors:              monitorIDs,
		HealthThresholdCount:  d.Get("health_threshold_count").(int),
		HealthCheckProfileID:  healthCheckProfileID,
		NotificationProfileID: notificationProfileID,
		UserGroupIDs:          userGroupIDs,
		OnCallScheduleID:      d.Get("on_call_schedule_id").(string),
		ThirdPartyServiceIDs:  thirdPartyServiceIDs,
		TagIDs:                tagIDs,
	}
}

func updateSubgroupResourceData(d *schema.ResourceData, subgroup *api.Subgroup) {
	d.Set("display_name", subgroup.DisplayName)
	d.Set("description", subgroup.Description)
	d.Set("monitors", subgroup.Monitors)
	d.Set("health_threshold_count", subgroup.HealthThresholdCount)
	d.Set("top_group_id", subgroup.TopGroupID)
	d.Set("parent_group_id", subgroup.ParentGroupID)
	d.Set("group_type", subgroup.Type)
	d.Set("healthcheck_profile_id", subgroup.HealthCheckProfileID)
	d.Set("notification_profile_id", subgroup.NotificationProfileID)
	d.Set("user_group_ids", subgroup.UserGroupIDs)
	d.Set("on_call_schedule_id", subgroup.OnCallScheduleID)
	d.Set("third_party_service_ids", subgroup.ThirdPartyServiceIDs)
	d.Set("tag_ids", subgroup.TagIDs)
	d.Set("check_frequency", subgroup.CheckFrequency)
}
