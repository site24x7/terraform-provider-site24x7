package common

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

/*
Supported Maintenance Types:
3 - Once
2 - Weekly
*/

var ScheduleMaintenanceSchema = map[string]*schema.Schema{
	"display_name": {
		Type:     schema.TypeString,
		Required: true,
	},

	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},

	"maintenance_type": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "Maintenance type. 3 = Once, 2 = Weekly",
	},

	// ---------- Common ----------
	"start_time": {
		Type:     schema.TypeString,
		Required: true,
	},

	"end_time": {
		Type:     schema.TypeString,
		Required: true,
	},

	"time_zone": {
		Type:     schema.TypeString,
		Optional: true,
		Description: "Time zone for your scheduled maintenance. Default is account timezone.",
	},

	"perform_monitoring": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  true,
	},

	// ---------- Once ----------
	"start_date": {
		Type:     schema.TypeString,
		Optional: true,
		Description: "Required for once maintenance. Format: yyyy-mm-dd",
	},

	"end_date": {
		Type:     schema.TypeString,
		Optional: true,
		Description: "Required for once maintenance. Format: yyyy-mm-dd",
	},

	// ---------- Weekly ----------
	"start_day": {
		Type:     schema.TypeInt,
		Optional: true,
		Description: "Start day for weekly maintenance (1=Sun ... 7=Sat)",
	},
	"end_day": {
		Type:     schema.TypeInt,
		Optional: true,
		Description: "End day for weekly maintenance (1=Sun ... 7=Sat)",
	},

	"duration": {
		Type:     schema.TypeInt,
		Optional: true,
		Description: "Duration in minutes (required for weekly maintenance)",
	},

	"week_days": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeInt,
		},
		Description: "Days of week on which maintenance should recur",
	},

	"execute_every": {
		Type:     schema.TypeInt,
		Optional: true,
		Default:  1,
		Description: "Interval at which weekly maintenance recurs (1–4)",
	},

	"maintenance_start_on": {
		Type:     schema.TypeString,
		Optional: true,
		Description: "Date on which weekly maintenance should start. Format: yyyy-mm-dd",
	},

	// ---------- Resource Selection ----------
	"selection_type": {
		Type:     schema.TypeInt,
		Optional: true,
		Default:  2,
		Description: "1=Monitor Groups, 2=Monitors, 3=Tags",
	},

	"monitors": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{Type: schema.TypeString},
	},

	"monitor_groups": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{Type: schema.TypeString},
	},

	"tags": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{Type: schema.TypeString},
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

	var monitorsIDs, monitorGroupIDs, tagIDs []string

	for _, id := range d.Get("monitors").([]interface{}) {
		monitorsIDs = append(monitorsIDs, id.(string))
	}

	for _, id := range d.Get("monitor_groups").([]interface{}) {
		monitorGroupIDs = append(monitorGroupIDs, id.(string))
	}

	for _, id := range d.Get("tags").([]interface{}) {
		tagIDs = append(tagIDs, id.(string))
	}

	sm := &api.ScheduleMaintenance{
		MaintenanceID:     d.Id(),
		DisplayName:       d.Get("display_name").(string),
		Description:       d.Get("description").(string),
		MaintenanceType:   d.Get("maintenance_type").(int),
		StartTime:         d.Get("start_time").(string),
		EndTime:           d.Get("end_time").(string),
		TimeZone:          d.Get("time_zone").(string),
		SelectionType:     api.ResourceType(d.Get("selection_type").(int)),
		Monitors:          monitorsIDs,
		MonitorGroups:     monitorGroupIDs,
		Tags:              tagIDs,
		PerformMonitoring: d.Get("perform_monitoring").(bool),
	}

	// Once
	if sm.MaintenanceType == 3 {
		sm.StartDate = d.Get("start_date").(string)
		sm.EndDate = d.Get("end_date").(string)
	}

	// Weekly
	if sm.MaintenanceType == 2 {
		sm.StartDay = d.Get("start_day").(int)
		sm.EndDay = d.Get("end_day").(int)

		sm.Duration = d.Get("duration").(int)
		sm.ExecuteEvery = d.Get("execute_every").(int)
		sm.MaintenanceStartOn = d.Get("maintenance_start_on").(string)

		if v, ok := d.GetOk("week_days"); ok {
			for _, day := range v.([]interface{}) {
				sm.WeekDays = append(sm.WeekDays, day.(int))
			}
		}
	}

	return sm
}

// Called during read
func updateScheduleMaintenanceResourceData(d *schema.ResourceData, sm *api.ScheduleMaintenance) {
	d.Set("display_name", sm.DisplayName)
	d.Set("description", sm.Description)
	d.Set("maintenance_type", sm.MaintenanceType)
	d.Set("start_time", sm.StartTime)
	d.Set("end_time", sm.EndTime)
	d.Set("time_zone", sm.TimeZone)
	d.Set("perform_monitoring", sm.PerformMonitoring)
	d.Set("selection_type", sm.SelectionType)
	d.Set("monitors", sm.Monitors)
	d.Set("monitor_groups", sm.MonitorGroups)
	d.Set("tags", sm.Tags)

	if sm.MaintenanceType == 3 {
		if sm.StartDate != "" {
			d.Set("start_date", sm.StartDate)
		}
		if sm.EndDate != "" {
			d.Set("end_date", sm.EndDate)
		}
	}

	if sm.MaintenanceType == 2 {
		if sm.StartDay > 0 {
			d.Set("start_day", sm.StartDay)
		}
		if sm.EndDay > 0 {
			d.Set("end_day", sm.EndDay)
		}

		// duration can be "", number, or null
		if v, ok := sm.Duration.(float64); ok {
			d.Set("duration", int(v))
		}

		if sm.ExecuteEvery > 0 {
			d.Set("execute_every", sm.ExecuteEvery)
		}

		// ✅ Critical fix: only set if non-empty
		if len(sm.WeekDays) > 0 {
			d.Set("week_days", sm.WeekDays)
		}

		if sm.MaintenanceStartOn != "" {
			d.Set("maintenance_start_on", sm.MaintenanceStartOn)
		}
	}
}