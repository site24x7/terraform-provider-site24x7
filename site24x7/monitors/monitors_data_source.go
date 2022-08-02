package monitors

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var monitorsDataSourceSchema = map[string]*schema.Schema{
	"name_regex": {
		Type:     schema.TypeString,
		Optional: true,
		// ValidateFunc: validation.StringIsValidRegExp,
	},
	"monitor_type": {
		Type:     schema.TypeString,
		Optional: true,
	},
	// Computed values
	"ids": {
		Type:     schema.TypeList,
		Computed: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	},
	"ids_and_names": {
		Type:     schema.TypeList,
		Computed: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	},
}

func DataSourceSite24x7Monitors() *schema.Resource {
	return &schema.Resource{
		Read:   monitorsDataSourceRead,
		Schema: monitorsDataSourceSchema,
	}
}

// monitorDataSourceRead fetches all server monitors from Site24x7
func monitorsDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	allMonitorList, err := client.WebsiteMonitors().List()
	if err != nil {
		return err
	}
	var monitorIDs []string
	var monitorIDsAndNames []string

	monitorType := d.Get("monitor_type").(string)
	nameRegex := d.Get("name_regex").(string)

	if monitorType != "" && nameRegex == "" {
		for _, monitor := range allMonitorList {
			if monitorType == monitor.Type {
				monitorIDs = append(monitorIDs, monitor.MonitorID)
				monitorIDsAndNames = append(monitorIDsAndNames, monitor.MonitorID+"__"+monitor.DisplayName)
			}
		}
	} else if monitorType == "" && nameRegex != "" {
		r := regexp.MustCompile(nameRegex)
		for _, monitor := range allMonitorList {
			if len(monitor.DisplayName) > 0 && r.MatchString(monitor.DisplayName) {
				monitorIDs = append(monitorIDs, monitor.MonitorID)
				monitorIDsAndNames = append(monitorIDsAndNames, monitor.MonitorID+"__"+monitor.DisplayName)
			}
		}
	} else if monitorType != "" && nameRegex != "" {
		r := regexp.MustCompile(nameRegex)
		for _, monitor := range allMonitorList {
			if len(monitor.DisplayName) > 0 && r.MatchString(monitor.DisplayName) && monitorType == monitor.Type {
				monitorIDs = append(monitorIDs, monitor.MonitorID)
				monitorIDsAndNames = append(monitorIDsAndNames, monitor.MonitorID+"__"+monitor.DisplayName)
			}
		}
	}

	if len(monitorIDs) == 0 {
		return errors.New("Unable to find monitors matching the name : \"" + d.Get("name_regex").(string) + "\" and monitor type : \"" + monitorType + "\"")
	}

	updateMonitorIDsResourceData(d, monitorIDs, monitorIDsAndNames)

	return nil
}

func updateMonitorIDsResourceData(d *schema.ResourceData, monitorIDs []string, monitorIDsAndNames []string) {
	d.SetId(fmt.Sprintf("%d", hashcode.String(d.Get("name_regex").(string)+d.Get("monitor_type").(string))))
	d.Set("ids", monitorIDs)
	d.Set("ids_and_names", monitorIDsAndNames)
}
