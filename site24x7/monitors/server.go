package monitors

import (
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var ServerMonitorSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name for the monitor.",
	},
	"poll_interval": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     1,
		Description: "Frequency at which data has to be collected for the server monitor.",
	},
	"log_needed": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Boolean to enable/disable Event Log/Syslog monitoring.",
	},
	"perform_automation": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Execute IT Automation during scheduled maintenance.",
	},
	"notification_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Notification profile to be associated with the monitor.",
	},
	"notification_profile_name": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Name of the notification profile to be associated with the monitor.",
	},
	"threshold_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Threshold profile to be associated with the monitor.",
	},
	"monitor_groups": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Computed:    true,
		Description: "List of monitor groups to which the monitor has to be associated.",
	},
	"user_group_ids": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Computed:    true,
		Description: "List of user groups to be notified when the monitor is down.",
	},
	"user_group_names": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "Name of the user groups to be associated with the monitor.",
	},
	"tag_ids": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Computed:    true,
		Description: "List of tag IDs to be associated to the monitor.",
	},
	"tag_names": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of tag names to be associated to the monitor.",
	},
	"third_party_service_ids": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of Third Party Service IDs to be associated to the monitor.",
	},
}

func ResourceSite24x7ServerMonitor() *schema.Resource {
	return &schema.Resource{
		Read:   serverMonitorRead,
		Update: serverMonitorUpdate,
		Delete: serverMonitorDelete,
		Exists: serverMonitorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: ServerMonitorSchema,
	}
}

func serverMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	serverMonitor, err := client.ServerMonitors().Get(d.Id())
	if err != nil {
		return err
	}

	updateServerMonitorResourceData(d, serverMonitor)

	return nil
}

func serverMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)
	serverMon, err := client.ServerMonitors().Get(d.Id())

	serverMonitor, err := resourceDataToServerMonitor(d, client, serverMon)
	if err != nil {
		return err
	}

	serverMonitor, err = client.ServerMonitors().Update(serverMonitor)
	if err != nil {
		return err
	}

	d.SetId(serverMonitor.MonitorID)

	// return serverMonitorRead(d, meta)
	return nil
}

func serverMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.ServerMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func serverMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.ServerMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToServerMonitor(d *schema.ResourceData, client site24x7.Client, serverMon *api.ServerMonitor) (*api.ServerMonitor, error) {

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

	serverMonitor := &api.ServerMonitor{
		MonitorID:          serverMon.MonitorID,
		DisplayName:        d.Get("display_name").(string),
		Type:               string(api.SERVER),
		HostName:           serverMon.HostName,
		IPAddress:          serverMon.IPAddress,
		TemplateID:         serverMon.TemplateID,
		PollInterval:       d.Get("poll_interval").(int),
		ITAutomationModule: serverMon.ITAutomationModule,
		PluginModule:       serverMon.PluginModule,
		LogNeeded:          d.Get("log_needed").(bool),
		PerformAutomation:  d.Get("perform_automation").(bool),

		NotificationProfileID: d.Get("notification_profile_id").(string),
		ThresholdProfileID:    d.Get("threshold_profile_id").(string),
		MonitorGroups:         monitorGroups,
		UserGroupIDs:          userGroupIDs,
		TagIDs:                tagIDs,
		ThirdPartyServiceIDs:  thirdPartyServiceIDs,
	}

	// Notification Profile
	_, notificationProfileErr := site24x7.SetNotificationProfile(client, d, serverMonitor)
	if notificationProfileErr != nil {
		return nil, notificationProfileErr
	}

	// User Alert Groups
	_, userAlertGroupErr := site24x7.SetUserGroup(client, d, serverMonitor)
	if userAlertGroupErr != nil {
		return nil, userAlertGroupErr
	}

	// Tags
	_, tagsErr := site24x7.SetTags(client, d, serverMonitor)
	if tagsErr != nil {
		return nil, tagsErr
	}

	if serverMonitor.ThresholdProfileID == "" {
		profile, err := site24x7.DefaultThresholdProfile(client, api.SSL_CERT)
		if err != nil {
			return nil, err
		}
		serverMonitor.ThresholdProfileID = profile.ProfileID
		d.Set("threshold_profile_id", profile)
	}

	return serverMonitor, nil
}

func updateServerMonitorResourceData(d *schema.ResourceData, monitor *api.ServerMonitor) {
	d.Set("monitor_id", monitor.MonitorID)
	d.Set("display_name", monitor.DisplayName)
	d.Set("type", monitor.Type)
	d.Set("host_name", monitor.HostName)
	d.Set("ip_address", monitor.IPAddress)
	d.Set("template_id", monitor.TemplateID)
	d.Set("poll_interval", monitor.PollInterval)
	d.Set("it_automation_module", monitor.ITAutomationModule)
	d.Set("plugin_module", monitor.PluginModule)
	d.Set("log_needed", monitor.LogNeeded)
	d.Set("perform_automation", monitor.PerformAutomation)

	d.Set("notification_profile_id", monitor.NotificationProfileID)
	d.Set("threshold_profile_id", monitor.ThresholdProfileID)
	d.Set("monitor_groups", monitor.MonitorGroups)
	d.Set("user_group_ids", monitor.UserGroupIDs)
	d.Set("tag_ids", monitor.TagIDs)
	d.Set("third_party_service_ids", monitor.ThirdPartyServiceIDs)
}
