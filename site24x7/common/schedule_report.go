package common

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var ScheduleReportSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name for the Summary Report.",
	},
	"report_type": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     17,
		Description: "Report type constant. Summary Report = 17.",
	},
	"selection_type": {
		Type:         schema.TypeInt,
		Required:     true,
		Description:  "Resource type for the report. Allowed values: 0 (All Monitors), 2 (Monitors), 3 (Tags), 4 (Monitor Type).",
	},
	"report_format": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "Report format constant (e.g., 1 = PDF, 2 = CSV).",
	},
	"report_frequency": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     1,
		Description: "Frequency for the report. 1 = Daily.",
	},
	"scheduled_time": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "Hour of day at which the report is generated (0–23).",
	},
	"scheduled_day": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "day at which the report is generated (0–23).",
	},
	"user_groups": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Required:    true,
		Description: "List of user group IDs that will receive the report.",
	},
}

func ResourceSite24x7ScheduleReport() *schema.Resource {
	return &schema.Resource{
		Create: scheduleReportCreate,
		Read:   scheduleReportRead,
		Update: scheduleReportUpdate,
		Delete: scheduleReportDelete,
		Exists: scheduleReportExists,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: ScheduleReportSchema,
	}
}

func scheduleReportCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	sr := resourceDataToScheduleReport(d)

	created, err := client.ScheduleReport().Create(sr)
	if err != nil {
		return err
	}

	d.SetId(created.ReportID)
	return nil
}

func scheduleReportRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	sr, err := client.ScheduleReport().Get(d.Id())
	if apierrors.IsNotFound(err) {
		d.SetId("")
		return nil
	}
	if err != nil {
		return err
	}

	updateScheduleReportResourceData(d, sr)
	return nil
}

type scheduleReportUpdatePayload struct {
	DisplayName     string   `json:"display_name"`
	ReportType      int      `json:"report_type"`
	SelectionType   int      `json:"selection_type"`
	ReportFormat    int      `json:"report_format"`
	ReportFrequency int      `json:"report_frequency"`
	ScheduledTime   int      `json:"scheduled_time"`
	UserGroups      []string `json:"user_groups"`
	ScheduledDay    int      `json:"scheduled_day"`
}

func scheduleReportUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

		var userGroups []string
	for _, e := range d.Get("user_groups").([]interface{}) {
		userGroups = append(userGroups, e.(string))
	}
	// ✅ Shadow payload (NO report_id)
	payload := scheduleReportUpdatePayload{
		DisplayName:     d.Get("display_name").(string),
		ReportType:      d.Get("report_type").(int),
		SelectionType:   d.Get("selection_type").(int),
		ReportFormat:    d.Get("report_format").(int),
		ReportFrequency: d.Get("report_frequency").(int),
		ScheduledTime:   d.Get("scheduled_time").(int),
		UserGroups:      userGroups,
		ScheduledDay:    d.Get("scheduled_day").(int),
	}

	updated, err := client.ScheduleReport().UpdateRaw(d.Id(), payload)
	if err != nil {
		return err
	}

	d.SetId(updated.ReportID)
	return nil
}

func scheduleReportDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.ScheduleReport().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}
	return err
}

func scheduleReportExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.ScheduleReport().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func resourceDataToScheduleReport(d *schema.ResourceData) *api.ScheduleReport {
	var userGroups []string

	for _, e := range d.Get("user_groups").([]interface{}) {
		userGroups = append(userGroups, e.(string))
	}

	return &api.ScheduleReport{
		ReportID:         d.Id(), // ✅ CRITICAL FIX
		DisplayName:     d.Get("display_name").(string),
		ReportType:      d.Get("report_type").(int),
		SelectionType:   d.Get("selection_type").(int),
		ReportFormat:    d.Get("report_format").(int),
		ReportFrequency: d.Get("report_frequency").(int),
		ScheduledTime:   d.Get("scheduled_time").(int),
		UserGroups:      userGroups,
		ScheduledDay:    d.Get("scheduled_day").(int),
	}
}

func updateScheduleReportResourceData(d *schema.ResourceData, sr *api.ScheduleReport) {
	d.Set("display_name", sr.DisplayName)
	d.Set("report_type", sr.ReportType)
	d.Set("selection_type", sr.SelectionType)
	d.Set("report_format", sr.ReportFormat)
	d.Set("report_frequency", sr.ReportFrequency)
	d.Set("scheduled_time", sr.ScheduledTime)
	d.Set("scheduled_day", sr.ScheduledDay)

	if sr.UserGroups != nil {
		var iface []interface{}
		for _, u := range sr.UserGroups {
			iface = append(iface, u)
		}
		d.Set("user_groups", iface)
	} else {
		d.Set("user_groups", nil)
	}
}
