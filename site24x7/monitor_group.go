package site24x7

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
)

var MonitorGroupSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display Name for the Monitor Group.",
	},
	"description": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Description for the Monitor Group.",
	},
	// As of now we don't support associating monitors via configuration file
	// "monitors": {
	// 	Type: schema.TypeList,
	// 	Elem: &schema.Schema{
	// 		Type: schema.TypeString,
	// 	},
	// 	Optional: true,
	// },
	"health_threshold_count": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     1,
		Description: "Number of monitors' health that decide the group status. ‘0’ implies that all the monitors are considered for determining the group status. Default value is 1.",
	},
	"dependency_resource_ids": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.",
	},
	"suppress_alert": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Description: "Boolean value indicating whether to suppress alert when the dependent monitor is down. Setting suppress_alert = true with an empty dependency_resource_id is meaningless.",
	},
	"tag_ids": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Computed:    true,
		Description: "List of tag IDs to be associated to the monitor group.",
	},
}

func ResourceSite24x7MonitorGroup() *schema.Resource {
	return &schema.Resource{
		Create: monitorGroupCreate,
		Read:   monitorGroupRead,
		Update: monitorGroupUpdate,
		Delete: monitorGroupDelete,
		Exists: monitorGroupExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: MonitorGroupSchema,
	}
}

func monitorGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	monitorGroup := resourceDataToMonitorGroupCreate(d)

	monitorGroup, err := client.MonitorGroups().Create(monitorGroup)
	if err != nil {
		return err
	}

	d.SetId(monitorGroup.GroupID)

	// Read is called for updating state after modification
	// https://www.terraform.io/docs/extend/best-practices/detecting-drift.html#update-state-after-modification
	// return monitorGroupRead(d, meta)
	return nil
}

func monitorGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	monitorGroup, err := client.MonitorGroups().Get(d.Id())
	if err != nil {
		return err
	}

	updateMonitorGroupResourceData(d, monitorGroup)

	return nil
}

func monitorGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	monitorGroup, err := client.MonitorGroups().Get(d.Id())
	if err != nil {
		return err
	}

	monitorGroup = resourceDataToMonitorGroupUpdate(d, monitorGroup)

	monitorGroup, err = client.MonitorGroups().Update(monitorGroup)
	if err != nil {
		return err
	}

	d.SetId(monitorGroup.GroupID)

	// Read is called for updating state after modification
	// https://www.terraform.io/docs/extend/best-practices/detecting-drift.html#update-state-after-modification
	// return monitorGroupRead(d, meta)
	return nil
}

func monitorGroupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	err := client.MonitorGroups().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func monitorGroupExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(Client)

	_, err := client.MonitorGroups().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToMonitorGroupCreate(d *schema.ResourceData) *api.MonitorGroup {

	var monitorIDs []string
	// If monitors are set in the configuration file iterate them and append to monitorIDs
	if _, monitorsExistsInConf := d.GetOk("monitors"); monitorsExistsInConf {
		for _, id := range d.Get("monitors").([]interface{}) {
			monitorIDs = append(monitorIDs, id.(string))
		}
	}

	dependencyIDs := d.Get("dependency_resource_ids").(*schema.Set).List()
	dependencyResourceIDs := make([]string, 0, len(dependencyIDs))
	for _, dependencyResourceID := range dependencyIDs {
		if dependencyResourceID != nil {
			dependencyResourceIDs = append(dependencyResourceIDs, dependencyResourceID.(string))
		}
	}

	var tagIDs []string
	for _, id := range d.Get("tag_ids").(*schema.Set).List() {
		if id != nil {
			tagIDs = append(tagIDs, id.(string))
		}
	}

	return &api.MonitorGroup{
		GroupID:                d.Id(),
		DisplayName:            d.Get("display_name").(string),
		Description:            d.Get("description").(string),
		Monitors:               monitorIDs,
		HealthThresholdCount:   d.Get("health_threshold_count").(int),
		DependencyResourceIDs:  dependencyResourceIDs,
		SuppressAlert:          d.Get("suppress_alert").(bool),
		DependencyResourceType: 2,
		TagIDs:                 tagIDs,
	}
}

func resourceDataToMonitorGroupUpdate(d *schema.ResourceData, monitorGroup *api.MonitorGroup) *api.MonitorGroup {
	// For all the three cases mentioned below we find the diff and append monitors to the monitorIDs list
	// 	- Monitors set in configuration file but not present in Site24x7
	// 	- Monitors not set in configuration file but present in Site24x7
	// 	- Monitors not set in configuration file and not present in Site24x7

	// var monitorIDs []string
	// if d.HasChange("monitors") {
	// 	oldMonitors, newMonitors := d.GetChange("monitors")
	// 	for _, id := range oldMonitors.([]interface{}) {
	// 		monitorIDs = append(monitorIDs, id.(string))
	// 	}
	// 	for _, id := range newMonitors.([]interface{}) {
	// 		monitorID := id.(string)
	// 		_, found := api.Find(monitorIDs, monitorID)
	// 		if !found {
	// 			monitorIDs = append(monitorIDs, monitorID)
	// 		}
	// 	}
	// }

	dependencyIDs := d.Get("dependency_resource_ids").(*schema.Set).List()
	dependencyResourceIDs := make([]string, 0, len(dependencyIDs))
	for _, dependencyResourceID := range dependencyIDs {
		if dependencyResourceID != nil {
			dependencyResourceIDs = append(dependencyResourceIDs, dependencyResourceID.(string))
		}
	}

	var suppressAlert bool
	if d.HasChange("suppress_alert") {
		oldSuppressAlert, newSuppressAlert := d.GetChange("suppress_alert")
		if _, suppressAlertExistsInConf := d.GetOk("suppress_alert"); suppressAlertExistsInConf {
			suppressAlert = newSuppressAlert.(bool)
		} else {
			suppressAlert = oldSuppressAlert.(bool)
		}
	}

	var tagIDs []string
	for _, id := range d.Get("tag_ids").(*schema.Set).List() {
		if id != nil {
			tagIDs = append(tagIDs, id.(string))
		}
	}

	return &api.MonitorGroup{
		GroupID:     d.Id(),
		DisplayName: d.Get("display_name").(string),
		Description: d.Get("description").(string),
		// Setting monitors from GET response. Empty "monitors" in PUT request dissociates all monitors from the monitor group.
		Monitors:               monitorGroup.Monitors,
		HealthThresholdCount:   d.Get("health_threshold_count").(int),
		DependencyResourceIDs:  dependencyResourceIDs,
		SuppressAlert:          suppressAlert,
		DependencyResourceType: 2,
		TagIDs:                 tagIDs,
	}
}

func updateMonitorGroupResourceData(d *schema.ResourceData, monitorGroup *api.MonitorGroup) {
	d.Set("display_name", monitorGroup.DisplayName)
	d.Set("description", monitorGroup.Description)
	// d.Set("monitors", monitorGroup.Monitors)
	d.Set("health_threshold_count", monitorGroup.HealthThresholdCount)
	d.Set("dependency_resource_ids", monitorGroup.DependencyResourceIDs)
	d.Set("suppress_alert", monitorGroup.SuppressAlert)
	d.Set("tag_ids", monitorGroup.TagIDs)
}
