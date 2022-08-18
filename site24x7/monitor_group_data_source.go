package site24x7

import (
	"errors"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	log "github.com/sirupsen/logrus"
	"github.com/site24x7/terraform-provider-site24x7/api"
)

var monitorGroupDataSourceSchema = map[string]*schema.Schema{
	"name_regex": {
		Type:     schema.TypeString,
		Required: true,
		// ValidateFunc: validation.StringIsValidRegExp,
	},
	"display_name": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Display Name for the Monitor Group.",
	},
	"description": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Description for the Monitor Group.",
	},
	"health_threshold_count": {
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "Number of monitors' health that decide the group status. ‘0’ implies that all the monitors are considered for determining the group status.",
	},
	"monitors": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Computed:    true,
		Description: "List of monitors associated to the group.",
	},
	"dependency_resource_ids": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Computed:    true,
		Description: "List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.",
	},
	"suppress_alert": {
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "Boolean value indicating whether to suppress alert when the dependent monitor is down.",
	},
}

func DataSourceSite24x7MonitorGroup() *schema.Resource {
	return &schema.Resource{
		Read:   monitorGroupDataSourceRead,
		Schema: monitorGroupDataSourceSchema,
	}
}

// monitorGroupDataSourceRead fetches all monitorGroup from Site24x7
func monitorGroupDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	monitorGroupList, err := client.MonitorGroups().List()
	if err != nil {
		return err
	}
	log.Println("MonitorGroup list ============================ ", monitorGroupList)

	nameRegex := d.Get("name_regex")

	var monitorGroup *api.MonitorGroup
	if nameRegex != "" {
		// (?i) - Case insensitive match
		nameRegexPattern := regexp.MustCompile("(?i)" + nameRegex.(string))
		for _, groupInfo := range monitorGroupList {
			if len(groupInfo.DisplayName) > 0 {
				if nameRegexPattern.MatchString(groupInfo.DisplayName) {
					monitorGroup = new(api.MonitorGroup)
					monitorGroup.GroupID = groupInfo.GroupID
					monitorGroup.DisplayName = groupInfo.DisplayName
					monitorGroup.Description = groupInfo.Description
					monitorGroup.HealthThresholdCount = groupInfo.HealthThresholdCount
					monitorGroup.Monitors = groupInfo.Monitors
					monitorGroup.DependencyResourceIDs = groupInfo.DependencyResourceIDs
					monitorGroup.SuppressAlert = groupInfo.SuppressAlert
					break
				}
			}
		}
	} else {
		return errors.New("Please enter a value for the attribute name_regex!")
	}

	if monitorGroup == nil {
		return errors.New("Unable to find monitor group matching the name : \"" + d.Get("name_regex").(string))
	}

	updateMonitorGroupDataSourceResourceData(d, monitorGroup)

	return nil
}

func updateMonitorGroupDataSourceResourceData(d *schema.ResourceData, monitorGroup *api.MonitorGroup) {
	d.SetId(monitorGroup.GroupID)
	d.Set("display_name", monitorGroup.DisplayName)
	d.Set("description", monitorGroup.Description)
	d.Set("monitors", monitorGroup.Monitors)
	d.Set("dependency_resource_ids", monitorGroup.DependencyResourceIDs)
	d.Set("health_threshold_count", monitorGroup.HealthThresholdCount)
	d.Set("suppress_alert", monitorGroup.SuppressAlert)
}
