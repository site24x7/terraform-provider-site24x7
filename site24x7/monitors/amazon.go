package monitors

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var AmazonMonitorSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name for the AWS monitor.",
	},
	"aws_access_key": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Access Key ID for the AWS account.",
		DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			// We are suppressing diff since aws_access_key in API response is encrypted.
			return true
		},
	},
	"aws_secret_key": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Secret Access key for the AWS account.",
		DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			// We are suppressing diff since aws_secret_key in API response is encrypted.
			return true
		},
	},
	"aws_discovery_frequency": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     1,
		Description: "Rediscovery polling interval for the AWS account.",
	},
	"aws_discover_services": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of AWS services that needs to be discovered. https://www.site24x7.com/help/api/#aws_discover_services",
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
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Computed:    true,
		Description: "List of Tag IDs to be associated to the monitor.",
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

func ResourceSite24x7AmazonMonitor() *schema.Resource {
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
	client := meta.(site24x7.Client)

	monitor, err := resourceDataToAmazonMonitor(d, client)
	if err != nil {
		return err
	}

	amazonMonitor, err := client.AmazonMonitors().Create(monitor)
	if err != nil {
		return err
	}

	d.SetId(amazonMonitor.MonitorID)

	return nil
}

func amazonMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	AmazonMonitor, err := client.AmazonMonitors().Get(d.Id())
	if err != nil {
		return err
	}

	updateAmazonMonitorResourceData(d, AmazonMonitor)

	return nil
}

func amazonMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

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
	client := meta.(site24x7.Client)

	err := client.AmazonMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func amazonMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.AmazonMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToAmazonMonitor(d *schema.ResourceData, client site24x7.Client) (*api.AmazonMonitor, error) {

	var userGroupIDs []string
	for _, id := range d.Get("user_group_ids").([]interface{}) {
		userGroupIDs = append(userGroupIDs, id.(string))
	}

	var tagIDs []string
	for _, id := range d.Get("tag_ids").([]interface{}) {
		tagIDs = append(tagIDs, id.(string))
	}

	var thirdPartyServiceIDs []string
	for _, id := range d.Get("third_party_service_ids").([]interface{}) {
		thirdPartyServiceIDs = append(thirdPartyServiceIDs, id.(string))
	}

	var awsServicesToDiscover []string
	for _, id := range d.Get("aws_discover_services").([]interface{}) {
		awsServicesToDiscover = append(awsServicesToDiscover, id.(string))
	}

	if len(userGroupIDs) == 0 {
		userGroup, err := site24x7.DefaultUserGroup(client)
		if err != nil {
			return nil, err
		}
		userGroupIDs = append(userGroupIDs, userGroup.UserGroupID)
	}

	amazonMonitor := &api.AmazonMonitor{
		MonitorID:             d.Id(),
		DisplayName:           d.Get("display_name").(string),
		Type:                  string(api.AMAZON),
		DiscoverFrequency:     d.Get("aws_discovery_frequency").(int),
		DiscoverServices:      awsServicesToDiscover,
		SecretKey:             d.Get("aws_secret_key").(string),
		AccessKey:             d.Get("aws_access_key").(string),
		NotificationProfileID: d.Get("notification_profile_id").(string),
		UserGroupIDs:          userGroupIDs,
		TagIDs:                tagIDs,
		ThirdPartyServiceIDs:  thirdPartyServiceIDs,
	}

	// Notification Profile
	_, notificationProfileErr := site24x7.SetNotificationProfile(client, d, amazonMonitor)
	if notificationProfileErr != nil {
		return nil, notificationProfileErr
	}

	// User Alert Groups
	_, userAlertGroupErr := site24x7.SetUserGroup(client, d, amazonMonitor)
	if userAlertGroupErr != nil {
		return nil, userAlertGroupErr
	}

	// Tags
	_, tagsErr := site24x7.SetTags(client, d, amazonMonitor)
	if tagsErr != nil {
		return nil, tagsErr
	}
	return amazonMonitor, nil
}

func updateAmazonMonitorResourceData(d *schema.ResourceData, amazonMonitor *api.AmazonMonitor) {
	d.Set("display_name", amazonMonitor.DisplayName)
	d.Set("aws_discovery_frequency", amazonMonitor.DiscoverFrequency)
	d.Set("notification_profile_id", amazonMonitor.NotificationProfileID)
	d.Set("user_group_ids", amazonMonitor.UserGroupIDs)
	d.Set("tag_ids", amazonMonitor.TagIDs)
	d.Set("third_party_service_ids", amazonMonitor.ThirdPartyServiceIDs)
	d.Set("aws_discover_services", amazonMonitor.DiscoverServices)
	d.Set("aws_secret_key", amazonMonitor.SecretKey)
	d.Set("aws_access_key", amazonMonitor.AccessKey)
}
