package site24x7

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
)

var AmazonMonitorSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "",
	},
	"aws_access_key": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "",
	},
	"aws_secret_key": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "",
	},
	"aws_discovery_frequency": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     1,
		Description: "",
	},
	"aws_discover_services": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "",
	},
	"notification_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "",
	},
	"user_group_ids": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of user groups to be notified when the monitor is down.",
	},
	"tag_ids": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of Tag IDs to be associated to the monitor.",
	},
}

func resourceSite24x7AmazonMonitor() *schema.Resource {
	return &schema.Resource{
		Create: amazonMonitorCreate,
		Read:   amazonMonitorRead,
		Update: amazonMonitorUpdate,
		Delete: amazonMonitorDelete,
		Exists: amazonMonitorExists,

		Schema: AmazonMonitorSchema,
	}
}

func amazonMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	amazonMonitor, err := resourceDataToAmazonMonitor(d, client)

	amazonMonitor, err = client.AmazonMonitors().Create(amazonMonitor)
	if err != nil {
		return err
	}

	d.SetId(amazonMonitor.MonitorID)

	return nil
}

func amazonMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	AmazonMonitor, err := client.AmazonMonitors().Get(d.Id())
	if err != nil {
		return err
	}

	updateAmazonMonitorResourceData(d, AmazonMonitor)

	return nil
}

func amazonMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	amazonMonitor, err := resourceDataToAmazonMonitor(d, client)
	if err != nil {
		return err
	}

	amazonMonitor, err = client.AmazonMonitors().Update(amazonMonitor)
	if err != nil {
		return err
	}

	d.SetId(amazonMonitor.MonitorID)

	return nil
}

func amazonMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	err := client.AmazonMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func amazonMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(Client)

	_, err := client.AmazonMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToAmazonMonitor(d *schema.ResourceData, client Client) (*api.AmazonMonitor, error) {

	var userGroupIDs []string
	for _, id := range d.Get("user_group_ids").([]interface{}) {
		userGroupIDs = append(userGroupIDs, id.(string))
	}

	var tagIDs []string
	for _, id := range d.Get("tag_ids").([]interface{}) {
		tagIDs = append(tagIDs, id.(string))
	}

	var awsServicesToDiscover []string
	for _, id := range d.Get("aws_discover_services").([]interface{}) {
		awsServicesToDiscover = append(awsServicesToDiscover, id.(string))
	}

	var notificationProfileID string
	notificationProfileID = d.Get("notification_profile_id").(string)
	if notificationProfileID == "" {
		profile, err := DefaultNotificationProfile(client)
		if err != nil {
			return nil, err
		}
		notificationProfileID = profile.ProfileID
	}

	if len(userGroupIDs) == 0 {
		userGroup, err := DefaultUserGroup(client)
		if err != nil {
			return nil, err
		}
		userGroupIDs = append(userGroupIDs, userGroup.UserGroupID)
	}

	return &api.AmazonMonitor{
		DisplayName:           d.Get("display_name").(string),
		Type:                  string(api.AMAZON),
		DiscoverFrequency:     d.Get("aws_discovery_frequency").(int),
		DiscoverServices:      awsServicesToDiscover,
		SecretKey:             d.Get("aws_secret_key").(string),
		AccessKey:             d.Get("aws_access_key").(string),
		NotificationProfileID: notificationProfileID,
		UserGroupIDs:          userGroupIDs,
		TagIDs:                tagIDs,
	}, nil
}

func updateAmazonMonitorResourceData(d *schema.ResourceData, amazonMonitor *api.AmazonMonitor) {
	d.Set("display_name", amazonMonitor.DisplayName)
	d.Set("aws_discovery_frequency", amazonMonitor.DiscoverFrequency)
	d.Set("notification_profile_id", amazonMonitor.NotificationProfileID)
	d.Set("user_group_ids", amazonMonitor.UserGroupIDs)
	d.Set("tag_ids", amazonMonitor.TagIDs)
	d.Set("aws_discover_services", amazonMonitor.DiscoverServices)
	d.Set("aws_secret_key", amazonMonitor.SecretKey)
	d.Set("aws_access_key", amazonMonitor.AccessKey)
}
