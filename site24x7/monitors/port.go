package monitors

import (
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var PortMonitorSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Name for the monitor",
	},
	"host_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Registered domain name.",
	},
	"port": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     80,
		Description: "Specify the port the host is listening to",
	},
	"timeout": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     10,
		Description: "Timeout for connecting to website. Default value is 10. Range 1 - 45",
	},
	"invert_port_check": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Invert the default behaviour of PORT check.",
	},
	"use_ipv6": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Monitoring is performed over IPv6 from supported locations. IPv6 locations do not fall back to IPv4 on failure.",
	},
	"use_ssl": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Monitoring is performed over IPv6 from supported locations. IPv6 locations do not fall back to IPv4 on failure.",
	},
	"application_type": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "",
	},
	"command": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "",
	},
	// Content Checks
	"matching_keyword_value": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Check for the keyword in the website response.",
	},
	"matching_keyword_severity": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     2,
		Description: "Severity with which alert has to raised when the matching keyword is found in the website response.",
	},
	"threshold_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Threshold profile to be associated with the monitor.",
	},
	"unmatching_keyword_value": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Check for the absence of the keyword in the website response.",
	},
	"unmatching_keyword_severity": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     2,
		Description: "Severity with which alert has to raised when the keyword is not present in the website response.",
	},
	"perform_automation": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "To perform automation or not",
	},
	"check_frequency": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "5",
		Description: "Interval at which your website has to be monitored. Default value is 5 minute.",
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
	"dependency_resource_ids": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.",
	},
	"on_call_schedule_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "A new On Call schedule to be associated with monitors when user group id  is not chosen",
	},
	"monitor_groups": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of monitor groups to which the monitor has to be associated",
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
}

func ResourceSite24x7PortMonitor() *schema.Resource {
	return &schema.Resource{
		Create: portMonitorCreate,
		Read:   portMonitorRead,
		Update: portMonitorUpdate,
		Delete: portMonitorDelete,
		Exists: portMonitorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: PortMonitorSchema,
	}
}

func portMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)
	portMonitor, err := resourceDataToPortMonitor(d, client)
	if err != nil {
		return err
	}
	portMonitor, err = client.PortMonitors().Create(portMonitor)
	if err != nil {
		return err
	}

	d.SetId(portMonitor.MonitorID)

	return nil
}

func portMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	portMonitor, err := client.PortMonitors().Get(d.Id())

	if err != nil {
		return err
	}

	updatePortMonitorResourceData(d, portMonitor)

	return nil
}

func portMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	portMonitor, err := resourceDataToPortMonitor(d, client)

	if err != nil {
		return err
	}

	portMonitor, err = client.PortMonitors().Update(portMonitor)
	if err != nil {
		return err
	}

	d.SetId(portMonitor.MonitorID)

	return nil
}

func portMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.PortMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func portMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.PortMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToPortMonitor(d *schema.ResourceData, client site24x7.Client) (*api.PortMonitor, error) {

	var monitorGroups []string
	for _, group := range d.Get("monitor_groups").([]interface{}) {
		if group != nil {
			monitorGroups = append(monitorGroups, group.(string))
		}
	}
	sort.Strings(monitorGroups)

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
	dependencyIDs := d.Get("dependency_resource_ids").(*schema.Set).List()
	dependencyResourceIDs := make([]string, 0, len(dependencyIDs))
	for _, dependencyResourceID := range dependencyIDs {
		if dependencyResourceID != nil {
			dependencyResourceIDs = append(dependencyResourceIDs, dependencyResourceID.(string))
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

	portMonitor := &api.PortMonitor{
		MonitorID:             d.Id(),
		DisplayName:           d.Get("display_name").(string),
		Type:                  string(api.PORT),
		HostName:              d.Get("host_name").(string),
		Port:                  d.Get("port").(int),
		Timeout:               d.Get("timeout").(int),
		UseIPV6:               d.Get("use_ipv6").(bool),
		UseSSL:                d.Get("use_ssl").(bool),
		CheckFrequency:        d.Get("check_frequency").(string),
		InvertPortCheck:       d.Get("invert_port_check").(bool),
		ApplicationType:       d.Get("application_type").(string),
		Command:               d.Get("command").(string),
		PerformAutomation:     d.Get("perform_automation").(bool),
		LocationProfileID:     d.Get("location_profile_id").(string),
		NotificationProfileID: d.Get("notification_profile_id").(string),
		OnCallScheduleID:      d.Get("on_call_schedule_id").(string),
		ThresholdProfileID:    d.Get("threshold_profile_id").(string),
		MonitorGroups:         monitorGroups,
		UserGroupIDs:          userGroupIDs,
		TagIDs:                tagIDs,
		DependencyResourceIDs: dependencyResourceIDs,
		ThirdPartyServiceIDs:  thirdPartyServiceIDs,
		ActionIDs:             actionRefs,
	}
	if _, ok := d.GetOk("matching_keyword_value"); ok {
		portMonitor.MatchingKeyword = &api.ValueAndSeverity{
			Value:    d.Get("matching_keyword_value").(string),
			Severity: api.Status(d.Get("matching_keyword_severity").(int)),
		}
	}

	if _, ok := d.GetOk("unmatching_keyword_value"); ok {
		portMonitor.UnmatchingKeyword = &api.ValueAndSeverity{
			Value:    d.Get("unmatching_keyword_value").(string),
			Severity: api.Status(d.Get("unmatching_keyword_severity").(int)),
		}
	}
	_, locationProfileErr := site24x7.SetLocationProfile(client, d, portMonitor)
	if locationProfileErr != nil {
		return nil, locationProfileErr
	}

	// Notification Profile
	_, notificationProfileErr := site24x7.SetNotificationProfile(client, d, portMonitor)
	if notificationProfileErr != nil {
		return nil, notificationProfileErr
	}

	// User Alert Groups
	_, userAlertGroupErr := site24x7.SetUserGroup(client, d, portMonitor)
	if userAlertGroupErr != nil {
		return nil, userAlertGroupErr
	}

	// Tags
	_, tagsErr := site24x7.SetTags(client, d, portMonitor)
	if tagsErr != nil {
		return nil, tagsErr
	}

	//Threshold
	if portMonitor.ThresholdProfileID == "" {
		profile, err := site24x7.DefaultThresholdProfile(client, api.PORT)
		if err != nil {
			return nil, err
		}
		portMonitor.ThresholdProfileID = profile.ProfileID
	}
	return portMonitor, nil
}

func updatePortMonitorResourceData(d *schema.ResourceData, monitor *api.PortMonitor) {
	d.Set("display_name", monitor.DisplayName)
	d.Set("type", monitor.Type)
	d.Set("host_name", monitor.HostName)
	d.Set("port", monitor.Port)
	d.Set("timeout", monitor.Timeout)
	d.Set("use_ipv6", monitor.UseIPV6)
	d.Set("use_ssl", monitor.UseSSL)
	d.Set("invert_port_check", monitor.InvertPortCheck)
	d.Set("application_type", monitor.ApplicationType)
	d.Set("command", monitor.Command)
	d.Set("check_frequency", monitor.CheckFrequency)
	d.Set("location_profile_id", monitor.LocationProfileID)
	d.Set("dependency_resource_ids", monitor.DependencyResourceIDs)
	d.Set("threshold_profile_id", monitor.ThresholdProfileID)
	d.Set("notification_profile_id", monitor.NotificationProfileID)
	d.Set("on_call_schedule_id", monitor.OnCallScheduleID)
	d.Set("monitor_groups", monitor.MonitorGroups)
	d.Set("user_group_ids", monitor.UserGroupIDs)
	d.Set("perform_automation", monitor.PerformAutomation)
	d.Set("tag_ids", monitor.TagIDs)
	d.Set("third_party_service_ids", monitor.ThirdPartyServiceIDs)
	// Content Check
	if monitor.MatchingKeyword != nil {
		d.Set("matching_keyword_value", monitor.MatchingKeyword.Value)
		d.Set("matching_keyword_severity", monitor.MatchingKeyword.Severity)
	}
	if monitor.UnmatchingKeyword != nil {
		d.Set("unmatching_keyword_value", monitor.UnmatchingKeyword.Value)
		d.Set("unmatching_keyword_severity", monitor.UnmatchingKeyword.Severity)
	}
}
