package monitors

import (
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var ISPMonitorSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Name for the monitor",
	},
	"hostname": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Registered domain name.",
	},
	"use_ipv6": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "whether to use ipv6 or not",
	},
	"port": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     443,
		Description: "Who is Server Port",
	},
	"timeout": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     10,
		Description: "Timeout for connecting to website. Default value is 10. Range 1 - 45",
	},
	"protocol": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "1",
		Description: "ICMP,TCP,UDP",
	},
	"threshold_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Threshold profile to be associated with the monitor.",
	},
	"on_call_schedule_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "if user_group_ids is not choosen,	On-Call Schedule of your choice.",
	},
	"check_frequency": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "5",
		Description: "Interval at which your RESRAPI has to be monitored. Default value is 5 minute.",
	},
	"perform_automation": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "To perform automation or not",
	},
	"location_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Location profile to be associated with the monitor",
	},
	"location_profile_name": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Name of the location profile to be associated with the monitor",
	},
	"notification_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Notification profile to be associated with the monitor",
	},
	"notification_profile_name": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Name of the notification profile to be associated with the monitor",
	},
	"user_group_ids": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Computed:    true,
		Description: "List of user groups to be notified when the monitor is down",
	},
	"user_group_names": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "Name of the user groups to be associated with the monitor",
	},
	"monitor_groups": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of monitor groups to which the monitor has to be associated.",
	},
	"actions": {
		Type:        schema.TypeMap,
		Optional:    true,
		Elem:        schema.TypeString,
		Description: "Action to be performed on monitor status changes",
	},
	"third_party_service_ids": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of Third Party Service IDs to be associated to the monitor",
	},
	"tag_ids": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Computed:    true,
		Description: "List of tag IDs to be associated to the monitor",
	},
	"tag_names": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of tag names to be associated to the monitor",
	},
	"dependency_resource_ids": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.",
	},
}

func ResourceSite24x7ISPMonitor() *schema.Resource {
	return &schema.Resource{
		Create: ispMonitorCreate,
		Read:   ispMonitorRead,
		Update: ispMonitorUpdate,
		Delete: ispMonitorDelete,
		Exists: ispMonitorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: ISPMonitorSchema,
	}
}

func ispMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	ispMonitor, err := resourceDataToISPMonitor(d, client)
	if err != nil {
		return err
	}
	ispMonitor, err = client.ISPMonitors().Create(ispMonitor)
	if err != nil {
		return err
	}

	d.SetId(ispMonitor.MonitorID)

	// return domainExpiryMonitorRead(d, meta)
	return nil
}

func ispMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	ispMonitor, err := client.ISPMonitors().Get(d.Id())

	if err != nil {
		return err
	}

	updateISPMonitorResourceData(d, ispMonitor)

	return nil
}

func ispMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	ispMonitor, err := resourceDataToISPMonitor(d, client)

	if err != nil {
		return err
	}

	ispMonitor, err = client.ISPMonitors().Update(ispMonitor)
	if err != nil {
		return err
	}

	d.SetId(ispMonitor.MonitorID)

	// return domainExpiryMonitorRead(d, meta)
	return nil
}

func ispMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.ISPMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func ispMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.ISPMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToISPMonitor(d *schema.ResourceData, client site24x7.Client) (*api.ISPMonitor, error) {

	var monitorGroups []string
	for _, group := range d.Get("monitor_groups").([]interface{}) {
		if group != nil {
			monitorGroups = append(monitorGroups, group.(string))
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
	//dependencyid
	dependencyIDs := d.Get("dependency_resource_ids").(*schema.Set).List()
	dependencyResourceIDs := make([]string, 0, len(dependencyIDs))
	for _, dependencyResourceID := range dependencyIDs {
		if dependencyResourceID != nil {
			dependencyResourceIDs = append(dependencyResourceIDs, dependencyResourceID.(string))
		}
	}
	var thirdPartyServiceIDs []string
	for _, id := range d.Get("third_party_service_ids").([]interface{}) {
		if id != nil {
			thirdPartyServiceIDs = append(thirdPartyServiceIDs, id.(string))
		}
	}

	actionMap := d.Get("actions").(map[string]interface{})
	var keys = make([]string, 0, len(actionMap))
	for k := range actionMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	actionRefs := make([]api.ActionRef, len(keys))
	for i, k := range keys {
		status, err := strconv.Atoi(k)
		if err != nil {
			return nil, err
		}
		actionRefs[i] = api.ActionRef{
			ActionID:  actionMap[k].(string),
			AlertType: api.Status(status),
		}
	}

	ispMonitor := &api.ISPMonitor{
		MonitorID:             d.Id(),
		DisplayName:           d.Get("display_name").(string),
		Type:                  string(api.ISP),
		Hostname:              d.Get("hostname").(string),
		UseIPV6:               d.Get("use_ipv6").(bool),
		Port:                  d.Get("port").(int),
		Timeout:               d.Get("timeout").(int),
		Protocol:              d.Get("protocol").(string),
		OnCallScheduleID:      d.Get("on_call_schedule_id").(string),
		PerformAutomation:     d.Get("perform_automation").(bool),
		LocationProfileID:     d.Get("location_profile_id").(string),
		NotificationProfileID: d.Get("notification_profile_id").(string),
		ThresholdProfileID:    d.Get("threshold_profile_id").(string),
		CheckFrequency:        d.Get("check_frequency").(string),
		DependencyResourceIDs: dependencyResourceIDs,
		MonitorGroups:         monitorGroups,
		UserGroupIDs:          userGroupIDs,
		TagIDs:                tagIDs,
		ThirdPartyServiceIDs:  thirdPartyServiceIDs,
		ActionIDs:             actionRefs,
	}
	_, locationProfileErr := site24x7.SetLocationProfile(client, d, ispMonitor)
	if locationProfileErr != nil {
		return nil, locationProfileErr
	}

	// Notification Profile
	_, notificationProfileErr := site24x7.SetNotificationProfile(client, d, ispMonitor)
	if notificationProfileErr != nil {
		return nil, notificationProfileErr
	}

	// User Alert Groups
	_, userAlertGroupErr := site24x7.SetUserGroup(client, d, ispMonitor)
	if userAlertGroupErr != nil {
		return nil, userAlertGroupErr
	}

	// Tags
	_, tagsErr := site24x7.SetTags(client, d, ispMonitor)
	if tagsErr != nil {
		return nil, tagsErr
	}

	//Threshold
	if ispMonitor.ThresholdProfileID == "" {
		profile, err := site24x7.DefaultThresholdProfile(client, api.ISP)
		if err != nil {
			return nil, err
		}
		ispMonitor.ThresholdProfileID = profile.ProfileID
	}
	return ispMonitor, nil
}

func updateISPMonitorResourceData(d *schema.ResourceData, monitor *api.ISPMonitor) {
	d.Set("display_name", monitor.DisplayName)
	d.Set("type", monitor.Type)
	d.Set("hostname", monitor.Hostname)
	d.Set("use_ipv6", monitor.UseIPV6)
	d.Set("port", monitor.Port)
	d.Set("timeout", monitor.Timeout)
	d.Set("check_frequency", monitor.CheckFrequency)
	d.Set("protocol", monitor.Protocol)
	d.Set("on_call_schedule_id", monitor.OnCallScheduleID)
	d.Set("location_profile_id", monitor.LocationProfileID)
	d.Set("notification_profile_id", monitor.NotificationProfileID)
	d.Set("perform_automation", monitor.PerformAutomation)
	d.Set("dependency_resource_ids", monitor.DependencyResourceIDs)
	d.Set("threshold_profile_id", monitor.ThresholdProfileID)
	d.Set("monitor_groups", monitor.MonitorGroups)
	d.Set("user_group_ids", monitor.UserGroupIDs)
	d.Set("tag_ids", monitor.TagIDs)
	d.Set("third_party_service_ids", monitor.ThirdPartyServiceIDs)
}
