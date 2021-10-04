package site24x7

import (
	"log"

	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var MonitorGroupSchema = map[string]*schema.Schema{
	"display_name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Required: true,
	},
	"monitors": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional: true,
	},
	"health_threshold_count": {
		Type:     schema.TypeInt,
		Optional: true,
	},
	"dependency_resource_id": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional: true,
	},
	"suppress_alert": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	},
}

func resourceSite24x7MonitorGroup() *schema.Resource {
	return &schema.Resource{
		Create: monitorGroupCreate,
		Read:   monitorGroupRead,
		Update: monitorGroupUpdate,
		Delete: monitorGroupDelete,
		Exists: monitorGroupExists,

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
	return monitorGroupRead(d, meta)
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
	return monitorGroupRead(d, meta)
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
	var healthThresholdCount int
	if _, thresholdExistsInConf := d.GetOk("health_threshold_count"); !thresholdExistsInConf {
		healthThresholdCount = 1
	}

	var monitorIDs []string
	// If monitors are set in the configuration file iterate them and append to monitorIDs
	if _, monitorsExistsInConf := d.GetOk("monitors"); monitorsExistsInConf {
		for _, id := range d.Get("monitors").([]interface{}) {
			monitorIDs = append(monitorIDs, id.(string))
		}
	}

	var dependencyIDs []string
	// If dependency_resource_id's are set in the configuration file iterate them and append to dependencyIDs
	if _, monitorsExistsInConf := d.GetOk("dependency_resource_id"); monitorsExistsInConf {
		for _, id := range d.Get("dependency_resource_id").([]interface{}) {
			dependencyIDs = append(dependencyIDs, id.(string))
		}
	}

	// log.Println("CREATE ====================== monitors  : ", monitorIDs)
	// log.Println("CREATE ====================== healthThresholdCount  : ", healthThresholdCount)
	return &api.MonitorGroup{
		GroupID:              d.Id(),
		DisplayName:          d.Get("display_name").(string),
		Description:          d.Get("description").(string),
		Monitors:             monitorIDs,
		HealthThresholdCount: healthThresholdCount,
		//DependencyResourceID: dependencyIDs,
		//SuppressAlert:        d.Get("suppress_alert").(bool),
	}
}

func resourceDataToMonitorGroupUpdate(d *schema.ResourceData, monitorGroup *api.MonitorGroup) *api.MonitorGroup {
	// log.Println("UPDATE ++++++++++++++++++++++ monitors  : ", d.Get("monitors"))
	// log.Println("UPDATE ++++++++++++++++++++++ dependency_resource_id  : ", d.Get("dependency_resource_id"))
	// log.Println("UPDATE ++++++++++++++++++++++ suppress_alert  : ", d.Get("suppress_alert"))
	// log.Println("UPDATE ++++++++++++++++++++++ health_threshold_count  : ", d.Get("health_threshold_count"))

	// For all the three cases mentioned below we find the diff and append monitors to the monitorIDs list
	// 	- Monitors set in configuration file but not present in Site24x7
	// 	- Monitors not set in configuration file but present in Site24x7
	// 	- Monitors not set in configuration file and not present in Site24x7
	var monitorIDs []string
	if d.HasChange("monitors") {
		oldMonitors, newMonitors := d.GetChange("monitors")
		for _, id := range oldMonitors.([]interface{}) {
			monitorIDs = append(monitorIDs, id.(string))
		}
		for _, id := range newMonitors.([]interface{}) {
			monitorID := id.(string)
			_, found := api.Find(monitorIDs, monitorID)
			if !found {
				monitorIDs = append(monitorIDs, monitorID)
			}
		}
		log.Println("resourceDataToMonitorGroupUpdate ++++++++++++++++++++++ ResourceData oldMonitors : ", oldMonitors)
		log.Println("resourceDataToMonitorGroupUpdate ++++++++++++++++++++++ ResourceData newMonitors : ", newMonitors)
	}

	var healthThresholdCount int
	if d.HasChange("health_threshold_count") {
		oldThresholdCount, newThresholdCount := d.GetChange("health_threshold_count")
		if newThresholdCount != 0 {
			healthThresholdCount = newThresholdCount.(int)
		} else {
			healthThresholdCount = oldThresholdCount.(int)
		}
	}

	var dependencyResourceID string
	if d.HasChange("dependency_resource_id") {
		oldDependencyID, newDependencyID := d.GetChange("dependency_resource_id")
		if newDependencyID != "" {
			dependencyResourceID = newDependencyID.(string)
		} else {
			dependencyResourceID = oldDependencyID.(string)
		}
	}

	var dependencyResourceIDs []string
	if d.HasChange("dependency_resource_id") {
		oldDependencyIDs, newDependencyIDs := d.GetChange("dependency_resource_id")
		for _, id := range oldDependencyIDs.([]interface{}) {
			dependencyResourceIDs = append(dependencyResourceIDs, id.(string))
		}
		for _, id := range newDependencyIDs.([]interface{}) {
			monitorID := id.(string)
			_, found := api.Find(dependencyResourceIDs, monitorID)
			if !found {
				dependencyResourceIDs = append(dependencyResourceIDs, monitorID)
			}
		}
		log.Println("resourceDataToMonitorGroupUpdate ++++++++++++++++++++++ ResourceData oldDependencyIDs : ", oldDependencyIDs)
		log.Println("resourceDataToMonitorGroupUpdate ++++++++++++++++++++++ ResourceData newDependencyIDs : ", newDependencyIDs)
	}

	var suppressAlert bool
	if d.HasChange("suppress_alert") {
		oldSuppressAlert, newSuppressAlert := d.GetChange("suppress_alert")
		if _, suppressAlertExistsInConf := d.GetOk("suppress_alert"); suppressAlertExistsInConf {
			suppressAlert = newSuppressAlert.(bool)
		} else {
			log.Println("resourceDataToMonitorGroupUpdate ++++++++++++++++++++++ suppressAlert NOT PRESENT IN CONF")
			suppressAlert = oldSuppressAlert.(bool)
		}
	}

	// log.Println("TO UPDATE ++++++++++++++++++++++ monitorIDs  : ", monitorIDs)
	// log.Println("TO UPDATE ++++++++++++++++++++++ healthThresholdCount  : ", healthThresholdCount)
	log.Println("TO UPDATE ++++++++++++++++++++++ dependencyResourceID  : ", dependencyResourceID)
	log.Println("TO UPDATE ++++++++++++++++++++++ suppressAlert  : ", suppressAlert)

	return &api.MonitorGroup{
		GroupID:              d.Id(),
		DisplayName:          d.Get("display_name").(string),
		Description:          d.Get("description").(string),
		Monitors:             monitorIDs,
		HealthThresholdCount: healthThresholdCount,
		//DependencyResourceID: dependencyResourceIDs,
		//SuppressAlert:        suppressAlert,
	}
}

func updateMonitorGroupResourceData(d *schema.ResourceData, monitorGroup *api.MonitorGroup) {
	d.Set("display_name", monitorGroup.DisplayName)
	d.Set("description", monitorGroup.Description)
	d.Set("monitors", monitorGroup.Monitors)
	d.Set("health_threshold_count", monitorGroup.HealthThresholdCount)
	//d.Set("dependency_resource_id", monitorGroup.DependencyResourceID)
	//d.Set("suppress_alert", monitorGroup.SuppressAlert)
}
