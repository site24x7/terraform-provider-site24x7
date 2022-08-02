package monitors

import (
	"errors"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var monitorDataSourceSchema = map[string]*schema.Schema{
	"name_regex": {
		Type:     schema.TypeString,
		Optional: true,
		// ValidateFunc: validation.StringIsValidRegExp,
	},
	"monitor_type": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"display_name": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
	"location_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Location profile associated with the monitor.",
	},
	"notification_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Notification profile associated with the monitor.",
	},
	"threshold_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Threshold profile associated with the monitor.",
	},
	"monitor_groups": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of monitor groups to which the monitor is associated.",
	},
	"user_group_ids": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Computed:    true,
		Description: "List of user groups notified when the monitor is down.",
	},
	"tag_ids": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Computed:    true,
		Description: "List of Tag IDs associated to the monitor.",
	},
	"third_party_service_ids": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of Third Party Service IDs associated to the monitor.",
	},
}

func DataSourceSite24x7Monitor() *schema.Resource {
	return &schema.Resource{
		Read:   monitorDataSourceRead,
		Schema: monitorDataSourceSchema,
	}
}

// monitorDataSourceRead fetches all server monitors from Site24x7
func monitorDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	allMonitorList, err := client.WebsiteMonitors().List()
	if err != nil {
		return err
	}
	monitorType := d.Get("monitor_type").(string)
	var genericMonitor *api.GenericMonitor
	if nameRegex, ok := d.GetOk("name_regex"); ok {
		r := regexp.MustCompile(nameRegex.(string))
		for _, monitor := range allMonitorList {
			if len(monitor.DisplayName) > 0 {
				if r.MatchString(monitor.DisplayName) && monitorType != "" && monitorType == monitor.Type {
					genericMonitor = new(api.GenericMonitor)
					genericMonitor.DisplayName = monitor.DisplayName
					genericMonitor.MonitorID = monitor.MonitorID
					genericMonitor.Type = monitor.Type
					genericMonitor.LocationProfileID = monitor.LocationProfileID
					genericMonitor.NotificationProfileID = monitor.NotificationProfileID
					genericMonitor.ThresholdProfileID = monitor.ThresholdProfileID
					genericMonitor.MonitorGroups = monitor.MonitorGroups
					genericMonitor.ThirdPartyServiceIDs = monitor.ThirdPartyServiceIDs
					genericMonitor.UserGroupIDs = monitor.UserGroupIDs
					genericMonitor.TagIDs = monitor.TagIDs
				}
			}
		}
	}

	if genericMonitor == nil {
		return errors.New("Unable to find monitor matching the name : \"" + d.Get("name_regex").(string) + "\" and monitor type : \"" + monitorType + "\"")
	}

	updateResourceData(d, genericMonitor)

	return nil
}

func updateResourceData(d *schema.ResourceData, monitor *api.GenericMonitor) {
	d.SetId(monitor.MonitorID)
	d.Set("monitor_type", monitor.Type)
	d.Set("display_name", monitor.DisplayName)
	d.Set("location_profile_id", monitor.LocationProfileID)
	d.Set("notification_profile_id", monitor.NotificationProfileID)
	d.Set("threshold_profile_id", monitor.ThresholdProfileID)
	d.Set("monitor_groups", monitor.MonitorGroups)
	d.Set("user_group_ids", monitor.UserGroupIDs)
	d.Set("tag_ids", monitor.TagIDs)
	d.Set("third_party_service_ids", monitor.ThirdPartyServiceIDs)
}
