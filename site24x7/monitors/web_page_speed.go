package monitors

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

// SAMPLE POST JSON
// {
// 	"display_name": "IANA - Home Page",
// 	"website": "https://www.iana.org/",
// 	"check_frequency": "15",
// 	"http_method": "G",
//  "auth_user": "user name",
//  "auth_pass": "user pass",
// 	"use_ipv6": false,

// 	"browser_type": 3,
// 	"device_type": "1",
// 	"type": "HOMEPAGE",
// 	"user_group_ids": [
// 	  "123456000000025005"
// 	],
// 	"timeout": 30,
// 	"match_case": false,
// 	"wpa_resolution": "1024,768",
// 	"matching_keyword": {
// 	  "value": "Management",
// 	  "severity": 2
// 	},
// 	"notification_profile_id": "123456000000029001",
// 	"user_agent": "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:94.0) Gecko/20100101 Firefox/83.0",
// 	"threshold_profile_id": "123456000030676013",
// 	"perform_automation": false,
// 	"website_type": 2,
// 	"location_profile_id": "123456000000025013",
// 	"ignore_filetype": {
// 	  "value": [
// 		{
// 		  "rc": "",
// 		  "ft": ""
// 		},
// 		{
// 		  "rc": "401,500",
// 		  "ft": "png,js"
// 		}
// 	  ]
// 	},
// 	"custom_headers": [
// 	  {
// 		"name": "requestheader",
// 		"value": "headerval"
// 	  }
// 	],
// 	"link_validation": {
// 	  "severity": 2,
// 	  "value": [
// 		{
// 		  "link": "https://www.iana.org/protocols",
// 		  "header_name": "Cache-Control",
// 		  "header_value": "dddd",
// 		  "type": "3"
// 		}
// 	  ]
// 	},
// 	"unmatching_keyword": {
// 	  "severity": 2,
// 	  "value": "aaaaa"
// 	},
// 	"match_regex": {
// 	  "severity": 2,
// 	  "value": "regex"
// 	},
// 	"up_status_codes": "200,300,400"
// }

// https://www.site24x7.com/help/api/#web-page-speed-(browser)
var webPageSpeedMonitorSchema = map[string]*schema.Schema{
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
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "5",
		Description: "Interval at which your website has to be monitored. Default value is 5 minute.",
	},
	"timeout": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     30,
		Description: "Timeout for connecting to website. Default value is 30. Range 1 - 45.",
	},
	"use_ipv6": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Monitoring is performed over IPv6 from supported locations. IPv6 locations do not fall back to IPv4 on failure.",
	},
	"website_type": {
		Type:     schema.TypeInt,
		Optional: true,
		Default:  1,
		Description: "Type of content the website page has. 1 - Static Website,	2 - Dynamic Website, 3 - Flash-Based Website.",
	},
	"browser_type": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     1,
		Description: "Type of the browser. 1 - Firefox, 2 - Chrome.",
	},
	"device_type": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "1",
		Description: "Type of the device used for running the speed test. 1 - Desktop, 2 - Mobile, 3 - Tab.",
	},
	"wpa_resolution": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "1024,768",
		Description: "Set a resolution based on your preferred device type.",
	},
	// HTTP Configuration
	"http_method": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "G",
		Description: "HTTP Method to be used for accessing the website. PUT, PATCH and DELETE are not supported. Default value is 'G'.",
	},
	"custom_headers": {
		Type:        schema.TypeMap,
		Optional:    true,
		Description: "A Map of Header name and value.",
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
	"user_agent": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "User Agent to be used while monitoring the website.",
	},
	"up_status_codes": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "",
		Description: "Provide a comma-separated list of HTTP status codes that indicate a successful response. You can specify individual status codes, as well as ranges separated with a colon.",
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
	"dependency_resource_ids": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.",
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
	"actions": {
		Type:        schema.TypeMap,
		Optional:    true,
		Elem:        schema.TypeString,
		Description: "Action to be performed on monitor status changes.",
	},
}

func ResourceSite24x7WebPageSpeedMonitor() *schema.Resource {
	return &schema.Resource{
		Create: webPageSpeedMonitorCreate,
		Read:   webPageSpeedMonitorRead,
		Update: webPageSpeedMonitorUpdate,
		Delete: webPageSpeedMonitorDelete,
		Exists: webPageSpeedMonitorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: webPageSpeedMonitorSchema,
	}
}

func webPageSpeedMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	webPageSpeedMonitor, err := resourceDataToWebPageSpeedMonitor(d, client)
	if err != nil {
		return err
	}

	webPageSpeedMonitor, err = client.WebPageSpeedMonitors().Create(webPageSpeedMonitor)
	if err != nil {
		return err
	}

	d.SetId(webPageSpeedMonitor.MonitorID)

	// return webPageSpeedMonitorRead(d, meta)
	return nil
}

func webPageSpeedMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	webPageSpeedMonitor, err := client.WebPageSpeedMonitors().Get(d.Id())
	if err != nil {
		return err
	}

	updateWebPageSpeedMonitorResourceData(d, webPageSpeedMonitor)

	return nil
}

func webPageSpeedMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	webPageSpeedMonitor, err := resourceDataToWebPageSpeedMonitor(d, client)
	if err != nil {
		return err
	}

	webPageSpeedMonitor, err = client.WebPageSpeedMonitors().Update(webPageSpeedMonitor)
	if err != nil {
		return err
	}

	d.SetId(webPageSpeedMonitor.MonitorID)

	// return webPageSpeedMonitorRead(d, meta)
	return nil
}

func webPageSpeedMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.WebPageSpeedMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func webPageSpeedMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.WebPageSpeedMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToWebPageSpeedMonitor(d *schema.ResourceData, client site24x7.Client) (*api.WebPageSpeedMonitor, error) {

	// Custom Headers
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

	var tagIDs []string
	for _, id := range d.Get("tag_ids").([]interface{}) {
		tagIDs = append(tagIDs, id.(string))
	}

	var thirdPartyServiceIDs []string
	for _, id := range d.Get("third_party_service_ids").([]interface{}) {
		thirdPartyServiceIDs = append(thirdPartyServiceIDs, id.(string))
	}

	var monitorGroups []string
	for _, group := range d.Get("monitor_groups").([]interface{}) {
		monitorGroups = append(monitorGroups, group.(string))
	}

	dependencyIDs := d.Get("dependency_resource_ids").(*schema.Set).List()
	dependencyResourceIDs := make([]string, 0, len(dependencyIDs))
	for _, dependencyResourceID := range dependencyIDs {
		dependencyResourceIDs = append(dependencyResourceIDs, dependencyResourceID.(string))
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

	webPageSpeedMonitor := &api.WebPageSpeedMonitor{
		MonitorID:      d.Id(),
		DisplayName:    d.Get("display_name").(string),
		Type:           string(api.HOMEPAGE),
		Website:        d.Get("website").(string),
		CheckFrequency: d.Get("check_frequency").(string),
		Timeout:        d.Get("timeout").(int),
		UseIPV6:        d.Get("use_ipv6").(bool),
		WebsiteType:    d.Get("website_type").(int),
		BrowserType:    d.Get("browser_type").(int),
		DeviceType:     d.Get("device_type").(string),
		WPAResolution:  d.Get("wpa_resolution").(string),
		// HTTP Configuration
		HTTPMethod:    d.Get("http_method").(string),
		CustomHeaders: customHeaders,
		AuthUser:      d.Get("auth_user").(string),
		AuthPass:      d.Get("auth_pass").(string),
		UserAgent:     d.Get("user_agent").(string),
		UpStatusCodes: d.Get("up_status_codes").(string),
		// Content Check
		MatchCase:             d.Get("match_case").(bool),
		LocationProfileID:     d.Get("location_profile_id").(string),
		NotificationProfileID: d.Get("notification_profile_id").(string),
		ThresholdProfileID:    d.Get("threshold_profile_id").(string),
		MonitorGroups:         monitorGroups,
		DependencyResourceIDs: dependencyResourceIDs,
		UserGroupIDs:          userGroupIDs,
		TagIDs:                tagIDs,
		ThirdPartyServiceIDs:  thirdPartyServiceIDs,
		ActionIDs:             actionRefs,
	}

	if _, ok := d.GetOk("match_regex_value"); ok {
		webPageSpeedMonitor.MatchRegex = &api.ValueAndSeverity{
			Value:    d.Get("match_regex_value").(string),
			Severity: api.Status(d.Get("match_regex_severity").(int)),
		}
	}

	if _, ok := d.GetOk("unmatching_keyword_value"); ok {
		webPageSpeedMonitor.UnmatchingKeyword = &api.ValueAndSeverity{
			Value:    d.Get("unmatching_keyword_value").(string),
			Severity: api.Status(d.Get("unmatching_keyword_severity").(int)),
		}
	}

	if _, ok := d.GetOk("matching_keyword_value"); ok {
		webPageSpeedMonitor.MatchingKeyword = &api.ValueAndSeverity{
			Value:    d.Get("matching_keyword_value").(string),
			Severity: api.Status(d.Get("matching_keyword_severity").(int)),
		}
	}

	if webPageSpeedMonitor.LocationProfileID == "" {
		locationProfileNameToMatch := d.Get("location_profile_name").(string)
		profile, err := site24x7.DefaultLocationProfile(client, locationProfileNameToMatch)
		if err != nil {
			return nil, err
		}
		webPageSpeedMonitor.LocationProfileID = profile.ProfileID
		d.Set("location_profile_id", profile.ProfileID)
	}

	// Notification Profile
	_, notificationProfileErr := site24x7.SetNotificationProfile(client, d, webPageSpeedMonitor)
	if notificationProfileErr != nil {
		return nil, notificationProfileErr
	}

	// User Alert Groups
	_, userAlertGroupErr := site24x7.SetUserGroup(client, d, webPageSpeedMonitor)
	if userAlertGroupErr != nil {
		return nil, userAlertGroupErr
	}

	// Tags
	_, tagsErr := site24x7.SetTags(client, d, webPageSpeedMonitor)
	if tagsErr != nil {
		return nil, tagsErr
	}

	if webPageSpeedMonitor.ThresholdProfileID == "" {
		profile, err := site24x7.DefaultThresholdProfile(client, api.HOMEPAGE)
		if err != nil {
			return nil, err
		}
		webPageSpeedMonitor.ThresholdProfileID = profile.ProfileID
		d.Set("threshold_profile_id", profile)
	}

	return webPageSpeedMonitor, nil
}

func updateWebPageSpeedMonitorResourceData(d *schema.ResourceData, monitor *api.WebPageSpeedMonitor) {
	d.Set("display_name", monitor.DisplayName)
	d.Set("type", monitor.Type)
	d.Set("website", monitor.Website)
	d.Set("check_frequency", monitor.CheckFrequency)
	d.Set("timeout", monitor.Timeout)
	d.Set("use_ipv6", monitor.UseIPV6)
	d.Set("website_type", monitor.WebsiteType)
	d.Set("browser_type", monitor.BrowserType)
	d.Set("device_type", monitor.DeviceType)
	d.Set("wpa_resolution", monitor.WPAResolution)
	// HTTP Configuration
	customHeaders := make(map[string]interface{})
	for _, h := range monitor.CustomHeaders {
		if h.Name == "" {
			continue
		}
		customHeaders[h.Name] = h.Value
	}
	d.Set("http_method", monitor.HTTPMethod)
	d.Set("custom_headers", customHeaders)
	d.Set("auth_user", monitor.AuthUser)
	d.Set("auth_pass", monitor.AuthPass)
	d.Set("user_agent", monitor.UserAgent)
	d.Set("up_status_codes", monitor.UpStatusCodes)

	// Content Check
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

	d.Set("location_profile_id", monitor.LocationProfileID)
	d.Set("notification_profile_id", monitor.NotificationProfileID)
	d.Set("threshold_profile_id", monitor.ThresholdProfileID)
	d.Set("monitor_groups", monitor.MonitorGroups)
	d.Set("dependency_resource_ids", monitor.DependencyResourceIDs)
	d.Set("user_group_ids", monitor.UserGroupIDs)
	d.Set("tag_ids", monitor.TagIDs)
	d.Set("third_party_service_ids", monitor.ThirdPartyServiceIDs)

	actions := make(map[string]interface{})
	for _, action := range monitor.ActionIDs {
		actions[fmt.Sprintf("%d", action.AlertType)] = action.ActionID
	}

	d.Set("actions", actions)
}
