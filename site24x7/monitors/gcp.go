package monitors

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var GCPMonitorSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name for the GCP monitor.",
	},
	"project_id": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Project ID of the GCP account.",
	},
	"private_key": {
		Type:        schema.TypeString,
		Required:    true,
		Sensitive:   true,
		Description: "Private key for GCP Service Account authentication.",
	},
	"client_email": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Client email for GCP Service Account authentication.",
	},
	"gcp_discovery_frequency": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "1",
		Description: "Rediscovery polling interval for the GCP account.",
	},
	"gcp_registration_method": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "1",
		Description: "Onboarding Method.",
	},
	"stop_rediscover_option": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "Preferred rediscovery frequency.",
	},
	"gcp_discover_services": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeInt,
		},
		Optional:    true,
		Description: "List of GCP services that need to be discovered.",
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
		Description: "List of tag IDs to be associated with the monitor.",
	},
	"tag_names": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of tag names to be associated with the monitor.",
	},
	"third_party_service_ids": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of Third Party Service IDs to be associated with the monitor.",
	},
}

func ResourceSite24x7GCPMonitor() *schema.Resource {
	return &schema.Resource{
		Create: gcpMonitorCreate,
		Read:   gcpMonitorRead,
		Update: gcpMonitorUpdate,
		Delete: gcpMonitorDelete,
		Exists: gcpMonitorExists,

		Schema: GCPMonitorSchema,
	}
}

func gcpMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	monitor, err := resourceDataToGCPMonitor(d, client)
	if err != nil {
		return err
	}

	gcpMonitor, err := client.GCPMonitors().Create(monitor)
	if err != nil {
		return err
	}

	d.SetId(gcpMonitor.MonitorID)

	return nil
}

func gcpMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	gcpMonitor, err := client.GCPMonitors().Get(d.Id())
	if err != nil {
		return err
	}

	updateGCPMonitorResourceData(d, gcpMonitor)

	return nil
}

func gcpMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	gcpMonitor, err := resourceDataToGCPMonitor(d, client)
	if err != nil {
		return err
	}
	if d.HasChange("gcp_discover_services") {
		var gcpServicesToDiscover []int
		if v, ok := d.GetOk("gcp_discover_services"); ok {
			for _, id := range v.([]interface{}) {
				if num, ok := id.(int); ok {
					gcpServicesToDiscover = append(gcpServicesToDiscover, num)
				}
			}
		}
		gcpMonitor.DiscoverServices = gcpServicesToDiscover
	}
	gcpMonitor, err = client.GCPMonitors().Update(gcpMonitor)
	if err != nil {
		return err
	}

	d.SetId(gcpMonitor.MonitorID)

	return nil
}

func gcpMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.GCPMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func gcpMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.GCPMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToGCPMonitor(d *schema.ResourceData, client site24x7.Client) (*api.GCPMonitor, error) {
	var userGroupIDs []string
	for _, id := range d.Get("user_group_ids").([]interface{}) {
		if id != nil {
			userGroupIDs = append(userGroupIDs, id.(string))
		}
	}

	var gcpServicesToDiscover []int
	for _, id := range d.Get("gcp_discover_services").([]interface{}) {
		if num, ok := id.(int); ok {
			gcpServicesToDiscover = append(gcpServicesToDiscover, num)
		}
	}

	if len(userGroupIDs) == 0 {
		userGroup, err := site24x7.DefaultUserGroup(client)
		if err != nil {
			return nil, err
		}
		userGroupIDs = append(userGroupIDs, userGroup.UserGroupID)
	}
	var gcpTags []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	var tagIDs []string
	for _, id := range d.Get("tag_ids").(*schema.Set).List() {
		if id != nil {
			tagIDs = append(tagIDs, id.(string))
		}
	}

	if v, ok := d.Get("gcp_tags").([]interface{}); ok {
		for _, tag := range v {
			if tagMap, ok := tag.(map[string]interface{}); ok {
				gcpTags = append(gcpTags, struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				}{
					Name:  tagMap["name"].(string),
					Value: tagMap["value"].(string),
				})
			}
		}
	}

	gcpMonitor := &api.GCPMonitor{
		MonitorID:             d.Id(),
		DisplayName:           d.Get("display_name").(string),
		Type:                  string(api.GCP),
		CheckFrequency:        d.Get("gcp_discovery_frequency").(string),
		StopRediscoverOption:  d.Get("stop_rediscover_option").(int),
		GcpRegistrationMethod: d.Get("gcp_registration_method").(string),
		DiscoverServices:      gcpServicesToDiscover,
		ProjectID:             d.Get("project_id").(string),
		NotificationProfileID: d.Get("notification_profile_id").(string),
		UserGroupIDs:          userGroupIDs,
		TagIDs:                tagIDs,
		GCPTags:               gcpTags,
		GCPSAContent: struct {
			PrivateKey  string `json:"private_key"`
			ClientEmail string `json:"client_email"`
		}{
			PrivateKey:  d.Get("private_key").(string),
			ClientEmail: d.Get("client_email").(string),
		},
	}

	// Notification Profile
	_, notificationProfileErr := site24x7.SetNotificationProfile(client, d, gcpMonitor)
	if notificationProfileErr != nil {
		return nil, notificationProfileErr
	}
	// Tags
	_, tagsErr := site24x7.SetTags(client, d, gcpMonitor)
	if tagsErr != nil {
		return nil, tagsErr
	}
	// User Alert Groups
	_, userAlertGroupErr := site24x7.SetUserGroup(client, d, gcpMonitor)
	if userAlertGroupErr != nil {
		return nil, userAlertGroupErr
	}

	return gcpMonitor, nil
}

func updateGCPMonitorResourceData(d *schema.ResourceData, gcpmonitor *api.GCPMonitor) {
	d.Set("display_name", gcpmonitor.DisplayName)
	d.Set("project_id", gcpmonitor.ProjectID)
	if v, ok := d.GetOk("private_key"); ok && v.(string) != "" {
		d.Set("private_key", v)
	}
	if v, ok := d.GetOk("client_email"); ok && v.(string) != "" {
		d.Set("client_email", v)
	}
	d.Set("gcp_discovery_frequency", gcpmonitor.CheckFrequency)
	d.Set("stop_rediscover_option", gcpmonitor.StopRediscoverOption)
	d.Set("gcp_discover_services", gcpmonitor.DiscoverServices)
	d.Set("gcp_registration_method", gcpmonitor.GcpRegistrationMethod)
	d.Set("notification_profile_id", gcpmonitor.NotificationProfileID)
	d.Set("user_group_ids", gcpmonitor.UserGroupIDs)
	d.Set("tag_ids", gcpmonitor.TagIDs)
	var gcpTags []map[string]string
	for _, tag := range gcpmonitor.GCPTags {
		gcpTags = append(gcpTags, map[string]string{
			"name":  tag.Name,
			"value": tag.Value,
		})
	}
	d.Set("gcp_tags", gcpTags)
}
