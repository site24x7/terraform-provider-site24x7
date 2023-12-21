package monitors

import (
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var FTPTransferMonitorSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Name for the monitor",
	},
	"host_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Registered domain name or ip addresss",
	},
	"protocol": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "FTP",
		Description: "HTTPS,SMTPS,POPS,IMAPS,FTPS or CUSTOM",
	},
	"type": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "FTP",
	},
	"port": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     21,
		Description: "Who is Server Port",
	},
	"check_frequency": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "5",
		Description: "Interval at which your RESRAPI has to be monitored. Default value is 5 minute.",
	},
	"timeout": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     10,
		Description: "Timeout for connecting to website. Default value is 10. Range 1 - 45",
	},
	"check_upload": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "To check upload or not",
	},
	"check_download": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "To check download or not",
	},
	"user_name": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "username to access the file",
	},
	"password": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "password to access the file",
	},
	"destination": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Destination of the file path",
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
	"monitor_groups": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of monitor groups to which the monitor has to be associated.",
	},
	"credential_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Credential Profile to associate.",
	},
	"dependency_resource_ids": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.",
	},
	"threshold_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Threshold profile to be associated with the monitor.",
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
	"perform_automation": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "To perform automation or not",
	},
	"actions": {
		Type:        schema.TypeMap,
		Optional:    true,
		Elem:        schema.TypeString,
		Description: "Action to be performed on monitor status changes",
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
	"on_call_schedule_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "if user_group_ids is not choosen,	On-Call Schedule of your choice.",
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
	"third_party_service_ids": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of Third Party Service IDs to be associated to the monitor",
	},
}

func ResourceSite24x7FTPTransferMonitor() *schema.Resource {
	return &schema.Resource{
		Create: ftpTransferMonitorCreate,
		Read:   ftpTransferMonitorRead,
		Update: ftpTransferMonitorUpdate,
		Delete: ftpTransferMonitorDelete,
		Exists: ftpTransferMonitorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: FTPTransferMonitorSchema,
	}
}

func ftpTransferMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	ftpTransferMonitor, err := resourceDataToFTPTransferMonitor(d, client)
	if err != nil {
		return err
	}

	ftpTransferMonitor, err = client.FTPTransferMonitors().Create(ftpTransferMonitor)
	if err != nil {
		return err
	}

	d.SetId(ftpTransferMonitor.MonitorID)
	return nil
}

func ftpTransferMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	ftpTransferMonitor, err := client.FTPTransferMonitors().Get(d.Id())

	if err != nil {
		return err
	}

	updateFTPTransferMonitorResourceData(d, ftpTransferMonitor)

	return nil
}

func ftpTransferMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	ftpTransferMonitor, err := resourceDataToFTPTransferMonitor(d, client)

	if err != nil {
		return err
	}

	ftpTransferMonitor, err = client.FTPTransferMonitors().Update(ftpTransferMonitor)
	if err != nil {
		return err
	}

	d.SetId(ftpTransferMonitor.MonitorID)

	return nil
}

func ftpTransferMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.FTPTransferMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func ftpTransferMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.FTPTransferMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToFTPTransferMonitor(d *schema.ResourceData, client site24x7.Client) (*api.FTPTransferMonitor, error) {

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

	ftpTransferMonitor := &api.FTPTransferMonitor{
		MonitorID:             d.Id(),
		DisplayName:           d.Get("display_name").(string),
		HostName:              d.Get("host_name").(string),
		Protocol:              d.Get("protocol").(string),
		Type:                  string(api.FTP),
		Port:                  d.Get("port").(int),
		CheckFrequency:        d.Get("check_frequency").(string),
		Timeout:               d.Get("timeout").(int),
		CheckUpload:           d.Get("check_upload").(bool),
		CheckDownload:         d.Get("check_download").(bool),
		Username:              d.Get("user_name").(string),
		Password:              d.Get("password").(string),
		Destination:           d.Get("destination").(string),
		LocationProfileID:     d.Get("location_profile_id").(string),
		MonitorGroups:         monitorGroups,
		CredentialProfileID:   d.Get("credential_profile_id").(string),
		DependencyResourceIDs: dependencyResourceIDs,
		ThresholdProfileID:    d.Get("threshold_profile_id").(string),
		TagIDs:                tagIDs,
		PerformAutomation:     d.Get("perform_automation").(bool),
		ActionIDs:             actionRefs,
		UserGroupIDs:          userGroupIDs,
		OnCallScheduleID:      d.Get("on_call_schedule_id").(string),
		NotificationProfileID: d.Get("notification_profile_id").(string),
		ThirdPartyServiceIDs:  thirdPartyServiceIDs,
	}
	_, locationProfileErr := site24x7.SetLocationProfile(client, d, ftpTransferMonitor)
	if locationProfileErr != nil {
		return nil, locationProfileErr
	}

	// Notification Profile
	_, notificationProfileErr := site24x7.SetNotificationProfile(client, d, ftpTransferMonitor)
	if notificationProfileErr != nil {
		return nil, notificationProfileErr
	}

	// User Alert Groups
	_, userAlertGroupErr := site24x7.SetUserGroup(client, d, ftpTransferMonitor)
	if userAlertGroupErr != nil {
		return nil, userAlertGroupErr
	}

	// Tags
	_, tagsErr := site24x7.SetTags(client, d, ftpTransferMonitor)
	if tagsErr != nil {
		return nil, tagsErr
	}
	//Threshold
	if ftpTransferMonitor.ThresholdProfileID == "" {
		profile, err := site24x7.DefaultThresholdProfile(client, api.FTP)
		if err != nil {
			return nil, err
		}
		ftpTransferMonitor.ThresholdProfileID = profile.ProfileID
	}

	return ftpTransferMonitor, nil
}

func updateFTPTransferMonitorResourceData(d *schema.ResourceData, monitor *api.FTPTransferMonitor) {
	d.Set("display_name", monitor.DisplayName)
	d.Set("host_name", monitor.HostName)
	d.Set("protocol", monitor.Protocol)
	d.Set("type", monitor.Type)
	d.Set("port", monitor.Port)
	d.Set("check_frequency", monitor.CheckFrequency)
	d.Set("timeout", monitor.Timeout)
	d.Set("check_upload", monitor.CheckUpload)
	d.Set("check_download", monitor.CheckDownload)
	d.Set("user_name", monitor.Username)
	d.Set("password", monitor.Password)
	d.Set("destination", monitor.Destination)
	d.Set("location_profile_id", monitor.LocationProfileID)
	d.Set("monitor_groups", monitor.MonitorGroups)
	d.Set("credential_profile_id", monitor.CredentialProfileID)
	d.Set("dependency_resource_ids", monitor.DependencyResourceIDs)
	d.Set("threshold_profile_id", monitor.ThresholdProfileID)
	d.Set("tag_ids", monitor.TagIDs)
	d.Set("perform_automation", monitor.PerformAutomation)
	d.Set("action_ids", monitor.ActionIDs)
	d.Set("user_group_ids", monitor.UserGroupIDs)
	d.Set("on_call_schedule_id", monitor.OnCallScheduleID)
	d.Set("notification_profile_id", monitor.NotificationProfileID)
	d.Set("third_party_service_ids", monitor.ThirdPartyServiceIDs)

}
