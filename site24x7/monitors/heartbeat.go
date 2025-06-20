package monitors

import (
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var HeartbeatMonitorSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display Name for the monitor.",
	},
	"name_in_ping_url": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Unique name to be used in the ping URL.",
	},
	"threshold_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Threshold profile to be associated with the monitor.",
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
	"monitor_groups": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of monitor groups to which the monitor has to be associated.",
	},
	"user_group_ids": {
		Type: schema.TypeSet,
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
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of Third Party Service IDs to be associated to the monitor.",
	},
	"on_call_schedule_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Mandatory, if the user group ID is not given. On-Call Schedule ID of your choice.",
	},
}

func ResourceSite24x7HeartbeatMonitor() *schema.Resource {
	return &schema.Resource{
		Create: heartbeatMonitorCreate,
		Read:   heartbeatMonitorRead,
		Update: heartbeatMonitorUpdate,
		Delete: heartbeatMonitorDelete,
		Exists: heartbeatMonitorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: HeartbeatMonitorSchema,
	}
}

func heartbeatMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	heartbeatMonitor, err := resourceDataToHeartbeatMonitor(d, client)
	if err != nil {
		return err
	}

	heartbeatMonitor, err = client.HeartbeatMonitors().Create(heartbeatMonitor)
	if err != nil {
		return err
	}

	d.SetId(heartbeatMonitor.MonitorID)

	return nil
}

func heartbeatMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	heartbeatMonitor, err := client.HeartbeatMonitors().Get(d.Id())
	if err != nil {
		return err
	}

	updateHeartbeatMonitorResourceData(d, heartbeatMonitor)

	return nil
}

func heartbeatMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	heartbeatMonitor, err := resourceDataToHeartbeatMonitor(d, client)
	if err != nil {
		return err
	}

	heartbeatMonitor, err = client.HeartbeatMonitors().Update(heartbeatMonitor)
	if err != nil {
		return err
	}

	d.SetId(heartbeatMonitor.MonitorID)

	// return heartbeatMonitorRead(d, meta)
	return nil
}

func heartbeatMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.HeartbeatMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func heartbeatMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.HeartbeatMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToHeartbeatMonitor(d *schema.ResourceData, client site24x7.Client) (*api.HeartbeatMonitor, error) {

	var monitorGroups []string
	for _, group := range d.Get("monitor_groups").(*schema.Set).List() {
		if group != nil {
			monitorGroups = append(monitorGroups, group.(string))
		}
	}
	sort.Strings(monitorGroups)

	var userGroupIDs []string
	for _, id := range d.Get("user_group_ids").(*schema.Set).List() {
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
	for _, id := range d.Get("third_party_service_ids").(*schema.Set).List() {
		if id != nil {
			thirdPartyServiceIDs = append(thirdPartyServiceIDs, id.(string))
		}
	}

	heartbeatMonitor := &api.HeartbeatMonitor{
		MonitorID:             d.Id(),
		DisplayName:           d.Get("display_name").(string),
		NameInPingURL:         d.Get("name_in_ping_url").(string),
		Type:                  string(api.HEARTBEAT),
		ThresholdProfileID:    d.Get("threshold_profile_id").(string),
		NotificationProfileID: d.Get("notification_profile_id").(string),
		MonitorGroups:         monitorGroups,
		UserGroupIDs:          userGroupIDs,
		TagIDs:                tagIDs,
		ThirdPartyServiceIDs:  thirdPartyServiceIDs,
		OnCallScheduleID:      d.Get("on_call_schedule_id").(string),
	}

	// Notification Profile
	_, notificationProfileErr := site24x7.SetNotificationProfile(client, d, heartbeatMonitor)
	if notificationProfileErr != nil {
		return nil, notificationProfileErr
	}

	// User Alert Groups
	_, userAlertGroupErr := site24x7.SetUserGroup(client, d, heartbeatMonitor)
	if userAlertGroupErr != nil {
		return nil, userAlertGroupErr
	}

	// Tags
	_, tagsErr := site24x7.SetTags(client, d, heartbeatMonitor)
	if tagsErr != nil {
		return nil, tagsErr
	}

	if heartbeatMonitor.ThresholdProfileID == "" {
		profile, err := site24x7.DefaultThresholdProfile(client, api.HEARTBEAT)
		if err != nil {
			return nil, err
		}
		heartbeatMonitor.ThresholdProfileID = profile.ProfileID
		d.Set("threshold_profile_id", profile)
	}

	return heartbeatMonitor, nil
}

func updateHeartbeatMonitorResourceData(d *schema.ResourceData, monitor *api.HeartbeatMonitor) {
	d.Set("display_name", monitor.DisplayName)
	d.Set("type", monitor.Type)
	d.Set("name_in_ping_url", monitor.NameInPingURL)
	d.Set("threshold_profile_id", monitor.ThresholdProfileID)
	d.Set("notification_profile_id", monitor.NotificationProfileID)
	d.Set("monitor_groups", monitor.MonitorGroups)
	d.Set("user_group_ids", monitor.UserGroupIDs)
	d.Set("tag_ids", monitor.TagIDs)
	d.Set("third_party_service_ids", monitor.ThirdPartyServiceIDs)
	d.Set("on_call_schedule_id", monitor.OnCallScheduleID)
}
