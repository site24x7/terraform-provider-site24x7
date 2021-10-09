package site24x7

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
)

var WebsiteMonitorSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display Name for the monitor.",
	},
	"website": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Website address to monitor.",
	},
	"check_frequency": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     1,
		Description: "Interval at which your website has to be monitored. Default value is 1 minute.",
	},
	"http_method": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "G",
		Description: "HTTP Method to be used for accessing the website. PUT, PATCH and DELETE are not supported",
	},
	"auth_user": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Authentication user name to access the website.",
	},
	"auth_pass": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Authentication password to access the website.",
	},
	"matching_keyword_value": {
		Type:     schema.TypeString,
		Optional: true,
		Default:  "Check for the keyword in the website response.",
	},
	"matching_keyword_severity": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     2,
		Description: "Severity with which alert has to raised when the matching keyword is found in the website response.",
	},
	"unmatching_keyword_value": {
		Type:     schema.TypeString,
		Optional: true,
		Default:  "Check for the absence of the keyword in the website response.",
	},
	"unmatching_keyword_severity": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     2,
		Description: "Severity with which alert has to raised when the keyword is not present in the website response.",
	},
	"match_regex_value": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Match the regular expression in the website response.",
	},
	"match_regex_severity": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     2,
		Description: "Severity with which alert has to raised when the matching regex is found in the website response.",
	},
	"match_case": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Perform case sensitive keyword search or not.",
	},
	"user_agent": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "User Agent to be used while monitoring the website.",
	},
	"custom_headers": {
		Type:        schema.TypeMap,
		Optional:    true,
		Description: "Header name and value in a string array.",
	},
	"timeout": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     10,
		Description: "Timeout for connecting to website. Default value is 10. Range 1 - 45.",
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
	"actions": {
		Type:        schema.TypeMap,
		Optional:    true,
		Elem:        schema.TypeString,
		Description: "Action to be performed on monitor status changes.",
	},
	"use_name_server": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Resolve the IP address using Domain Name Server.",
	},
	"up_status_codes": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "",
		Description: "Provide a comma-separated list of HTTP status codes that indicate a successful response. You can specify individual status codes, as well as ranges separated with a colon.",
	},
}

func resourceSite24x7WebsiteMonitor() *schema.Resource {
	return &schema.Resource{
		Create: websiteMonitorCreate,
		Read:   websiteMonitorRead,
		Update: websiteMonitorUpdate,
		Delete: websiteMonitorDelete,
		Exists: websiteMonitorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: WebsiteMonitorSchema,
	}
}

func websiteMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	websiteMonitor, err := resourceDataToWebsiteMonitor(d, client)
	if err != nil {
		return err
	}

	websiteMonitor, err = client.WebsiteMonitors().Create(websiteMonitor)
	if err != nil {
		return err
	}

	d.SetId(websiteMonitor.MonitorID)

	// return websiteMonitorRead(d, meta)
	return nil
}

func websiteMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	websiteMonitor, err := client.WebsiteMonitors().Get(d.Id())
	if err != nil {
		return err
	}

	updateWebsiteMonitorResourceData(d, websiteMonitor)

	return nil
}

func websiteMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	websiteMonitor, err := resourceDataToWebsiteMonitor(d, client)
	if err != nil {
		return err
	}

	websiteMonitor, err = client.WebsiteMonitors().Update(websiteMonitor)
	if err != nil {
		return err
	}

	d.SetId(websiteMonitor.MonitorID)

	// return websiteMonitorRead(d, meta)
	return nil
}

func websiteMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	err := client.WebsiteMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func websiteMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(Client)

	_, err := client.WebsiteMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToWebsiteMonitor(d *schema.ResourceData, client Client) (*api.WebsiteMonitor, error) {
	customHeaderMap := d.Get("custom_headers").(map[string]interface{})

	keys := make([]string, 0, len(customHeaderMap))
	for k := range customHeaderMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	customHeaders := make([]api.Header, len(keys))
	for i, k := range keys {
		customHeaders[i] = api.Header{Name: k, Value: customHeaderMap[k].(string)}
	}

	var userGroupIDs []string
	for _, id := range d.Get("user_group_ids").([]interface{}) {
		userGroupIDs = append(userGroupIDs, id.(string))
	}

	var monitorGroups []string
	for _, group := range d.Get("monitor_groups").([]interface{}) {
		monitorGroups = append(monitorGroups, group.(string))
	}

	actionMap := d.Get("actions").(map[string]interface{})

	keys = make([]string, 0, len(actionMap))
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

	websiteMonitor := &api.WebsiteMonitor{
		MonitorID:             d.Id(),
		DisplayName:           d.Get("display_name").(string),
		Type:                  string(api.URL),
		Website:               d.Get("website").(string),
		CheckFrequency:        strconv.Itoa(d.Get("check_frequency").(int)),
		HTTPMethod:            d.Get("http_method").(string),
		AuthUser:              d.Get("auth_user").(string),
		AuthPass:              d.Get("auth_pass").(string),
		MatchCase:             d.Get("match_case").(bool),
		UserAgent:             d.Get("user_agent").(string),
		CustomHeaders:         customHeaders,
		Timeout:               d.Get("timeout").(int),
		LocationProfileID:     d.Get("location_profile_id").(string),
		NotificationProfileID: d.Get("notification_profile_id").(string),
		ThresholdProfileID:    d.Get("threshold_profile_id").(string),
		MonitorGroups:         monitorGroups,
		UserGroupIDs:          userGroupIDs,
		ActionIDs:             actionRefs,
		UseNameServer:         d.Get("use_name_server").(bool),
		UpStatusCodes:         d.Get("up_status_codes").(string),
	}

	if _, ok := d.GetOk("match_regex_value"); ok {
		websiteMonitor.MatchRegex = &api.ValueAndSeverity{
			Value:    d.Get("match_regex_value").(string),
			Severity: api.Status(d.Get("match_regex_severity").(int)),
		}
	}

	if _, ok := d.GetOk("unmatching_keyword_value"); ok {
		websiteMonitor.UnmatchingKeyword = &api.ValueAndSeverity{
			Value:    d.Get("unmatching_keyword_value").(string),
			Severity: api.Status(d.Get("unmatching_keyword_severity").(int)),
		}
	}

	if _, ok := d.GetOk("matching_keyword_value"); ok {
		websiteMonitor.MatchingKeyword = &api.ValueAndSeverity{
			Value:    d.Get("matching_keyword_value").(string),
			Severity: api.Status(d.Get("matching_keyword_severity").(int)),
		}
	}

	if websiteMonitor.LocationProfileID == "" {
		locationProfileNameToMatch := d.Get("location_profile_name").(string)
		profile, err := DefaultLocationProfile(client, locationProfileNameToMatch)
		if err != nil {
			return nil, err
		}
		websiteMonitor.LocationProfileID = profile.ProfileID
		d.Set("location_profile_id", profile.ProfileID)
	}

	if websiteMonitor.NotificationProfileID == "" {
		profile, err := DefaultNotificationProfile(client)
		if err != nil {
			return nil, err
		}
		websiteMonitor.NotificationProfileID = profile.ProfileID
		d.Set("notification_profile_id", profile.ProfileID)
	}

	if websiteMonitor.ThresholdProfileID == "" {
		profile, err := DefaultThresholdProfile(client, api.URL)
		if err != nil {
			return nil, err
		}
		websiteMonitor.ThresholdProfileID = profile.ProfileID
		d.Set("threshold_profile_id", profile)
	}

	if len(websiteMonitor.UserGroupIDs) == 0 {
		userGroup, err := DefaultUserGroup(client)
		if err != nil {
			return nil, err
		}
		websiteMonitor.UserGroupIDs = []string{userGroup.UserGroupID}
		d.Set("user_group_ids", []string{userGroup.UserGroupID})
	}

	return websiteMonitor, nil
}

func updateWebsiteMonitorResourceData(d *schema.ResourceData, monitor *api.WebsiteMonitor) {
	d.Set("display_name", monitor.DisplayName)
	d.Set("type", monitor.Type)
	d.Set("website", monitor.Website)
	d.Set("check_frequency", monitor.CheckFrequency)
	d.Set("http_method", monitor.HTTPMethod)
	d.Set("auth_user", monitor.AuthUser)
	d.Set("auth_pass", monitor.AuthPass)
	if monitor.MatchingKeyword != nil {
		d.Set("matching_keyword_value", monitor.MatchingKeyword.Value)
		d.Set("matching_keyword_severity", monitor.MatchingKeyword.Severity)
	}
	if monitor.UnmatchingKeyword != nil {
		d.Set("unmatching_keyword_value", monitor.UnmatchingKeyword.Value)
		d.Set("unmatching_keyword_severity", monitor.UnmatchingKeyword.Severity)
	}
	if monitor.MatchRegex != nil {
		d.Set("match_regex_value", monitor.MatchRegex.Value)
		d.Set("match_regex_severity", monitor.MatchRegex.Severity)
	}
	d.Set("match_case", monitor.MatchCase)
	d.Set("user_agent", monitor.UserAgent)

	customHeaders := make(map[string]interface{})
	for _, h := range monitor.CustomHeaders {
		if h.Name == "" {
			continue
		}
		customHeaders[h.Name] = h.Value
	}

	d.Set("custom_headers", customHeaders)
	d.Set("timeout", monitor.Timeout)
	d.Set("location_profile_id", monitor.LocationProfileID)
	d.Set("notification_profile_id", monitor.NotificationProfileID)
	d.Set("threshold_profile_id", monitor.ThresholdProfileID)
	d.Set("monitor_groups", monitor.MonitorGroups)
	d.Set("user_group_ids", monitor.UserGroupIDs)

	actions := make(map[string]interface{})
	for _, action := range monitor.ActionIDs {
		actions[fmt.Sprintf("%d", action.AlertType)] = action.ActionID
	}

	d.Set("actions", actions)
	d.Set("use_name_server", monitor.UseNameServer)
	d.Set("up_status_codes", monitor.UpStatusCodes)
}
