package monitors

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var SSLMonitorSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display Name for the monitor.",
	},
	"domain_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Domain name to be verified for SSL Certificate.",
	},
	"timeout": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     30,
		Description: "Timeout for connecting to the host. Range 1 - 45.",
	},
	"protocol": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "HTTPS",
		Description: "Supported protocols are HTTPS, SMTPS, POPS, IMAPS, FTPS or CUSTOM",
	},
	"port": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     443,
		Description: "Server Port.",
	},
	"expire_days": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     30,
		Description: "Day threshold for certificate expiry notification. Range 1 - 999.",
	},
	"http_protocol_version": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "H1.1",
		Description: "Version of the HTTP protocol.",
	},
	"ignore_domain_mismatch": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Description: "Boolean to ignore domain name mismatch errors.",
	},
	"ignore_trust": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Description: "To ignore the validation of SSL/TLS certificate chain.",
	},
	"location_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Location profile to be associated with the monitor.",
	},
	"location_profile_name": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Name of the location profile to be associated with the monitor.",
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

func ResourceSite24x7SSLMonitor() *schema.Resource {
	return &schema.Resource{
		Create: sslMonitorCreate,
		Read:   sslMonitorRead,
		Update: sslMonitorUpdate,
		Delete: sslMonitorDelete,
		Exists: sslMonitorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: SSLMonitorSchema,
	}
}

func sslMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	sslMonitor, err := resourceDataToSSLMonitor(d, client)
	if err != nil {
		return err
	}

	sslMonitor, err = client.SSLMonitors().Create(sslMonitor)
	if err != nil {
		return err
	}

	d.SetId(sslMonitor.MonitorID)

	return nil
}

func sslMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	sslMonitor, err := client.SSLMonitors().Get(d.Id())
	if err != nil {
		return err
	}

	updateSSLMonitorResourceData(d, sslMonitor)

	return nil
}

func sslMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	sslMonitor, err := resourceDataToSSLMonitor(d, client)
	if err != nil {
		return err
	}

	sslMonitor, err = client.SSLMonitors().Update(sslMonitor)
	if err != nil {
		return err
	}

	d.SetId(sslMonitor.MonitorID)

	// return sslMonitorRead(d, meta)
	return nil
}

func sslMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.SSLMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func sslMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.SSLMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToSSLMonitor(d *schema.ResourceData, client site24x7.Client) (*api.SSLMonitor, error) {

	var monitorGroups []string
	for _, group := range d.Get("monitor_groups").([]interface{}) {
		monitorGroups = append(monitorGroups, group.(string))
	}

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

	sslMonitor := &api.SSLMonitor{
		MonitorID:             d.Id(),
		DisplayName:           d.Get("display_name").(string),
		Type:                  string(api.SSL_CERT),
		DomainName:            d.Get("domain_name").(string),
		Protocol:              d.Get("protocol").(string),
		Timeout:               d.Get("timeout").(int),
		Port:                  d.Get("port"),
		ExpireDays:            d.Get("expire_days").(int),
		HTTPProtocolVersion:   d.Get("http_protocol_version").(string),
		IgnoreDomainMismatch:  d.Get("ignore_domain_mismatch").(bool),
		IgnoreTrust:           d.Get("ignore_trust").(bool),
		LocationProfileID:     d.Get("location_profile_id").(string),
		NotificationProfileID: d.Get("notification_profile_id").(string),
		ThresholdProfileID:    d.Get("threshold_profile_id").(string),
		MonitorGroups:         monitorGroups,
		UserGroupIDs:          userGroupIDs,
		TagIDs:                tagIDs,
		ThirdPartyServiceIDs:  thirdPartyServiceIDs,
	}

	if sslMonitor.LocationProfileID == "" {
		locationProfileNameToMatch := d.Get("location_profile_name").(string)
		profile, err := site24x7.DefaultLocationProfile(client, locationProfileNameToMatch)
		if err != nil {
			return nil, err
		}
		sslMonitor.LocationProfileID = profile.ProfileID
		d.Set("location_profile_id", profile.ProfileID)
	}

	// Notification Profile
	_, notificationProfileErr := site24x7.SetNotificationProfile(client, d, sslMonitor)
	if notificationProfileErr != nil {
		return nil, notificationProfileErr
	}

	// User Alert Groups
	_, userAlertGroupErr := site24x7.SetUserGroup(client, d, sslMonitor)
	if userAlertGroupErr != nil {
		return nil, userAlertGroupErr
	}

	// Tags
	_, tagsErr := site24x7.SetTags(client, d, sslMonitor)
	if tagsErr != nil {
		return nil, tagsErr
	}

	if sslMonitor.ThresholdProfileID == "" {
		profile, err := site24x7.DefaultThresholdProfile(client, api.SSL_CERT)
		if err != nil {
			return nil, err
		}
		sslMonitor.ThresholdProfileID = profile.ProfileID
		d.Set("threshold_profile_id", profile)
	}

	return sslMonitor, nil
}

func updateSSLMonitorResourceData(d *schema.ResourceData, monitor *api.SSLMonitor) {
	d.Set("display_name", monitor.DisplayName)
	d.Set("type", monitor.Type)
	d.Set("domain_name", monitor.DomainName)
	d.Set("timeout", monitor.Timeout)
	d.Set("protocol", monitor.Protocol)
	d.Set("port", monitor.Port)
	d.Set("expire_days", monitor.ExpireDays)
	d.Set("http_protocol_version", monitor.HTTPProtocolVersion)
	d.Set("ignore_domain_mismatch", monitor.IgnoreDomainMismatch)
	d.Set("ignore_trust", monitor.IgnoreTrust)
	d.Set("location_profile_id", monitor.LocationProfileID)
	d.Set("notification_profile_id", monitor.NotificationProfileID)
	d.Set("threshold_profile_id", monitor.ThresholdProfileID)
	d.Set("monitor_groups", monitor.MonitorGroups)
	d.Set("user_group_ids", monitor.UserGroupIDs)
	d.Set("tag_ids", monitor.TagIDs)
	d.Set("third_party_service_ids", monitor.ThirdPartyServiceIDs)
}
