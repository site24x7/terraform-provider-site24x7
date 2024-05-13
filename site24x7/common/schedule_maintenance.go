package common

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

// SAMPLE POST JSON

// {
// 	"maintenance_type": 3,
// 	"selection_type": 2,
// 	"start_time": "19:41",
// 	"end_time": "20:44",
//  "timezone": "PST",
// 	"perform_monitoring": true,
// 	"display_name": "Test Schedule Maintenance",
// 	"description": "Test Schedule Maintenance",
// 	"monitors": [
// 	"123456000000631008"
// 	],
// 	"start_date": "2022-06-02",
// 	"end_date": "2022-06-02"
// }

var ScheduleMaintenanceSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name for the maintenance.",
	},
	"description": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Description for the maintenance.",
	},
	// As of now maintenance_type "3" is supported
	// "maintenance_type": {
	// 	Type:         schema.TypeInt,
	// 	Optional:     true,
	// 	Default:      3,
	// 	ValidateFunc: validation.IntInSlice([]int{3}),
	// 	Description:  "Configuration for Once/Daily/Weekly/Monthly only maintenance. Default is 3 - Once. Refer https://www.site24x7.com/help/api/#schedule_maintenance_constants",
	// },
	"start_date": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Mandatory, if the maintenance_type chosen is Once. Maintenance start date. Format - yyyy-mm-dd.",
	},
	"start_time": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Maintenance start time. Format - hh:mm",
	},
	"time_zone": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Time zone for your scheduled maintenance. Default value is your account timezone.",
	},
	"end_date": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Mandatory, if the maintenance_type chosen is Once. Maintenance end date. Format - yyyy-mm-dd.",
	},
	"end_time": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Maintenance end time. Format - hh:mm",
	},
	"selection_type": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     2,
		Description: "Resource Type associated with this integration. Default value is '0'. Can take values 1|2|3. '1' denotes 'Monitor Group', '2' denotes 'Monitors', '3' denotes 'Tags'",
	},
	"monitors": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "Monitors that need to be associated with the maintenance window when the selection_type = 2.",
	},
	"monitor_groups": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "Monitor Groups that need to be associated with the maintenance window when the selection_type = 1.",
	},
	"tags": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "Tags that need to be associated with the maintenance window when the selection_type = 3.",
	},
	"perform_monitoring": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Enable this to perform uptime monitoring of the resource during the maintenance window.",
	},
}

func ResourceSite24x7ScheduleMaintenance() *schema.Resource {
	return &schema.Resource{
		Create: scheduleMaintenanceCreate,
		Read:   scheduleMaintenanceRead,
		Update: scheduleMaintenanceUpdate,
		Delete: scheduleMaintenanceDelete,
		Exists: scheduleMaintenanceExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: ScheduleMaintenanceSchema,
	}
}

func scheduleMaintenanceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	scheduleMaintenance := resourceDataToScheduleMaintenance(d)

	scheduleMaintenance, err := client.ScheduleMaintenance().Create(scheduleMaintenance)
	if err != nil {
		return err
	}

	d.SetId(scheduleMaintenance.MaintenanceID)

	return nil
}

func scheduleMaintenanceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	scheduleMaintenance, err := client.ScheduleMaintenance().Get(d.Id())
	if err != nil {
		return err
	}

	updateScheduleMaintenanceResourceData(d, scheduleMaintenance)

	return nil
}

func scheduleMaintenanceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	scheduleMaintenance := resourceDataToScheduleMaintenance(d)

	scheduleMaintenance, err := client.ScheduleMaintenance().Update(scheduleMaintenance)
	if err != nil {
		return err
	}

	d.SetId(scheduleMaintenance.MaintenanceID)

	return nil
}

func scheduleMaintenanceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.ScheduleMaintenance().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func scheduleMaintenanceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.ScheduleMaintenance().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToScheduleMaintenance(d *schema.ResourceData) *api.ScheduleMaintenance {

	var monitorsIDs []string
	for _, id := range d.Get("monitors").([]interface{}) {
		monitorsIDs = append(monitorsIDs, id.(string))
	}

	var monitorGroupIDs []string
	for _, id := range d.Get("monitor_groups").([]interface{}) {
		monitorGroupIDs = append(monitorGroupIDs, id.(string))
	}

	var tagIDs []string
	for _, id := range d.Get("tags").([]interface{}) {
		tagIDs = append(tagIDs, id.(string))
	}

	return &api.ScheduleMaintenance{
		MaintenanceID:     d.Id(),
		DisplayName:       d.Get("display_name").(string),
		Description:       d.Get("description").(string),
		MaintenanceType:   3,
		TimeZone:          d.Get("time_zone").(string),
		StartDate:         d.Get("start_date").(string),
		EndDate:           d.Get("end_date").(string),
		StartTime:         d.Get("start_time").(string),
		EndTime:           d.Get("end_time").(string),
		SelectionType:     api.ResourceType(d.Get("selection_type").(int)),
		Monitors:          monitorsIDs,
		MonitorGroups:     monitorGroupIDs,
		Tags:              tagIDs,
		PerformMonitoring: d.Get("perform_monitoring").(bool),
	}
}

// Called during read and sets scheduleMaintenance in API response to ResourceData
func updateScheduleMaintenanceResourceData(d *schema.ResourceData, scheduleMaintenance *api.ScheduleMaintenance) {
	d.Set("display_name", scheduleMaintenance.DisplayName)
	d.Set("description", scheduleMaintenance.Description)
	d.Set("start_date", scheduleMaintenance.StartDate)
	d.Set("time_zone", scheduleMaintenance.TimeZone)
	d.Set("end_date", scheduleMaintenance.EndDate)
	d.Set("start_time", scheduleMaintenance.StartTime)
	d.Set("end_time", scheduleMaintenance.EndTime)
	d.Set("perform_monitoring", scheduleMaintenance.PerformMonitoring)
	d.Set("selection_type", scheduleMaintenance.SelectionType)
	d.Set("monitors", scheduleMaintenance.Monitors)
	d.Set("monitor_groups", scheduleMaintenance.MonitorGroups)
	d.Set("tags", scheduleMaintenance.Tags)
}
