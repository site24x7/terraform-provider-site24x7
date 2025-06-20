package monitors

import (
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var CronMonitorSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display Name for the monitor.",
	},
	"cron_expression": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Cron expression to denote the job schedule.",
	},
	"cron_tz": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Timezone of the server where job runs.",
	},
	"wait_time": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "Provide an extended period of time (seconds) to define when your alerts should be triggered. This is basically to avoid false alerts.",
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

func ResourceSite24x7CronMonitor() *schema.Resource {
	return &schema.Resource{
		Create: cronMonitorCreate,
		Read:   cronMonitorRead,
		Update: cronMonitorUpdate,
		Delete: cronMonitorDelete,
		Exists: cronMonitorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: CronMonitorSchema,
	}
}

func cronMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	cronMonitor, err := resourceDataToCronMonitor(d, client)
	if err != nil {
		return err
	}

	cronMonitor, err = client.CronMonitors().Create(cronMonitor)
	if err != nil {
		return err
	}

	d.SetId(cronMonitor.MonitorID)

	return nil
}

func cronMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	cronMonitor, err := client.CronMonitors().Get(d.Id())
	if err != nil {
		return err
	}

	updateCronMonitorResourceData(d, cronMonitor)

	return nil
}

func cronMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	cronMonitor, err := resourceDataToCronMonitor(d, client)
	if err != nil {
		return err
	}

	cronMonitor, err = client.CronMonitors().Update(cronMonitor)
	if err != nil {
		return err
	}

	d.SetId(cronMonitor.MonitorID)

	// return cronMonitorRead(d, meta)
	return nil
}

func cronMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.CronMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func cronMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.CronMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToCronMonitor(d *schema.ResourceData, client site24x7.Client) (*api.CronMonitor, error) {

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

	cronMonitor := &api.CronMonitor{
		MonitorID:             d.Id(),
		DisplayName:           d.Get("display_name").(string),
		CronExpression:        d.Get("cron_expression").(string),
		CronTz:                d.Get("cron_tz").(string),
		WaitTime:              d.Get("wait_time").(int),
		Type:                  string(api.CRON),
		ThresholdProfileID:    d.Get("threshold_profile_id").(string),
		NotificationProfileID: d.Get("notification_profile_id").(string),
		MonitorGroups:         monitorGroups,
		UserGroupIDs:          userGroupIDs,
		TagIDs:                tagIDs,
		ThirdPartyServiceIDs:  thirdPartyServiceIDs,
		OnCallScheduleID:      d.Get("on_call_schedule_id").(string),
	}

	// Notification Profile
	_, notificationProfileErr := site24x7.SetNotificationProfile(client, d, cronMonitor)
	if notificationProfileErr != nil {
		return nil, notificationProfileErr
	}

	// User Alert Groups
	_, userAlertGroupErr := site24x7.SetUserGroup(client, d, cronMonitor)
	if userAlertGroupErr != nil {
		return nil, userAlertGroupErr
	}

	// Tags
	_, tagsErr := site24x7.SetTags(client, d, cronMonitor)
	if tagsErr != nil {
		return nil, tagsErr
	}
	
	// Threshold profile
	if cronMonitor.ThresholdProfileID == "" {
		profile, err := site24x7.DefaultThresholdProfile(client, api.CRON)
		if err != nil {
			return nil, err
		}
		cronMonitor.ThresholdProfileID = profile.ProfileID
		d.Set("threshold_profile_id", profile)
	}

	return cronMonitor, nil
}

func updateCronMonitorResourceData(d *schema.ResourceData, monitor *api.CronMonitor) {
	d.Set("display_name", monitor.DisplayName)
	d.Set("type", monitor.Type)
	d.Set("cron_expression", monitor.CronExpression)
	d.Set("cron_tz", monitor.CronTz)
	d.Set("wait_time", monitor.WaitTime)
	d.Set("threshold_profile_id", monitor.ThresholdProfileID)
	d.Set("notification_profile_id", monitor.NotificationProfileID)
	d.Set("monitor_groups", monitor.MonitorGroups)
	d.Set("user_group_ids", monitor.UserGroupIDs)
	d.Set("tag_ids", monitor.TagIDs)
	d.Set("third_party_service_ids", monitor.ThirdPartyServiceIDs)
	d.Set("on_call_schedule_id", monitor.OnCallScheduleID)
}
