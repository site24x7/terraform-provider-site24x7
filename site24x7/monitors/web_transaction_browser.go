package monitors

import (
	"fmt"
	"log"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var WebTransactionBrowserMonitorSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Name for the monitor",
	},
	"type": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "REALBROWSER",
	},
	"base_url": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "BaseURL of the transaction",
	},
	"selenium_script": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Recorded Trasanction script to create a monitor",
	},
	"script_type": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Recorded transaction script type.(txt , side)",
	},
	"check_frequency": {
		Type:        schema.TypeString,
		Default:     15,
		Optional:    true,
		Description: "Check interval for monitoring.",
	},
	"async_dc_enabled": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "When asynchronous data collection is enabled, polling will be carried out from all the locations at the same time. If it is disabled, polling will be done consecutively from the selected locations.",
	},
	"browser_type": {
		Type:        schema.TypeInt,
		Default:     1,
		Optional:    true,
		Description: "Choose the browser type.",
	},
	"browser_version": {
		Type:        schema.TypeInt,
		Default:     10101,
		Optional:    true,
		Description: "Choose the browser version",
	},
	"think_time": {
		Type:        schema.TypeInt,
		Default:     1,
		Optional:    true,
		Description: "Think time between each steps",
	},
	"page_load_time": {
		Type:        schema.TypeInt,
		Default:     60,
		Optional:    true,
		Description: "Timeout for page load.",
	},
	"resolution": {
		Type:        schema.TypeString,
		Default:     "1600,900",
		Optional:    true,
		Description: "Screen resolution for running the script.",
	},
	"ip_type": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "IP Type for monitor.",
	},
	"ignore_cert_err": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Enter true or false to Trust the Server SSL Certificate. Default value is true.",
	},
	"perform_automation": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "To perform automation or not",
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
	"custom_headers": {
		Type:        schema.TypeMap,
		Optional:    true,
		Description: "A Map of Header name and value.",
	},
	"threshold_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Threshold profile to be associated with the monitor.",
	},
	"user_agent": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "User Agent to be used while monitoring the website.",
	},
	// "parallel_polling": {
	// 	Type:        schema.TypeBool,
	// 	Optional:    true,
	// 	Description: "Parallel polling option",
	// },
	"proxy_details": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"webProxyUrl": {
					Type:     schema.TypeString,
					Required: true,
				},
				"webProxyUname": {
					Type:     schema.TypeString,
					Required: true,
				},
				"webProxyPass": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
		Description: "Check for the proxy in the website response.",
	},
	"auth_details": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"username": {
					Type:     schema.TypeString,
					Required: true,
				},
				"password": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
		Description: "Check for the auth details in the website response.",
	},
	"cookies": {
		Type:        schema.TypeMap,
		Optional:    true,
		Description: "A Map of cookies name and value.",
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
		Description: "A new On Call schedule to be associated with monitors when user group id  is not chosen",
	},
	"monitor_groups": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of monitor groups to which the monitor has to be associated",
	},
	"dependency_resource_ids": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.",
	},
	"actions": {
		Type:        schema.TypeMap,
		Optional:    true,
		Elem:        schema.TypeString,
		Description: "Action to be performed on monitor IT Automation templates.",
	},
	"third_party_service_ids": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of Third Party Service IDs to be associated to the monitor",
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
}

func ResourceSite24x7WebTransactionBrowserMonitor() *schema.Resource {
	return &schema.Resource{
		Create: webTransactionBrowserMonitorCreate,
		Read:   webTransactionBrowserMonitorRead,
		Update: webTransactionBrowserMonitorUpdate,
		Delete: webTransactionBrowserMonitorDelete,
		Exists: webTransactionBrowserMonitorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: WebTransactionBrowserMonitorSchema,
	}
}

func webTransactionBrowserMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	webTransactionBrowserMonitor, err := resourceDataToWebTransactionBrowserMonitorCreate(d, client)
	if err != nil {
		return err
	}

	webTransactionBrowserMonitor, err = client.WebTransactionBrowserMonitors().Create(webTransactionBrowserMonitor)
	log.Println("GetTokenURL : ", webTransactionBrowserMonitor.AsyncDCEnabled)
	if err != nil {
		return err
	}

	d.SetId(webTransactionBrowserMonitor.MonitorID)

	// return webTransactionBrowserMonitorRead(d, meta)
	return nil
}

func webTransactionBrowserMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	webTransactionBrowserMonitor, err := client.WebTransactionBrowserMonitors().Get(d.Id())

	if err != nil {
		return err
	}

	updateWebTransactionBrowserMonitorResourceData(d, webTransactionBrowserMonitor)

	return nil
}

func webTransactionBrowserMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	webTransactionBrowserMonitor, err := resourceDataToWebTransactionBrowserMonitorCreate(d, client)

	if err != nil {
		return err
	}

	webTransactionBrowserMonitor.SeleniumScript = ""
	webTransactionBrowserMonitor.ScriptType = ""

	webTransactionBrowserMonitor, err = client.WebTransactionBrowserMonitors().Update(webTransactionBrowserMonitor)
	if err != nil {
		return err
	}

	d.SetId(webTransactionBrowserMonitor.MonitorID)

	// return webTransactionBrowserMonitorRead(d, meta)
	return nil
}

func webTransactionBrowserMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.WebTransactionBrowserMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func webTransactionBrowserMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.WebTransactionBrowserMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToWebTransactionBrowserMonitorCreate(d *schema.ResourceData, client site24x7.Client) (*api.WebTransactionBrowserMonitor, error) {

	// Cookies
	cookiesMap := d.Get("cookies").(map[string]interface{})
	keys := make([]string, 0, len(cookiesMap))
	for k := range cookiesMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	cookies := make([]api.Header, len(keys))
	for i, k := range keys {
		cookies[i] = api.Header{Name: k, Value: cookiesMap[k].(string)}
	}

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
	var akeys = make([]string, 0, len(actionMap))
	for k := range actionMap {
		akeys = append(akeys, k)
	}
	sort.Strings(akeys)
	actionRefs := make([]api.ActionRef, len(akeys))
	for i, k := range akeys {
		status, err := strconv.Atoi(k)
		if err != nil {
			return nil, err
		}
		actionRefs[i] = api.ActionRef{
			ActionID:  actionMap[k].(string),
			AlertType: api.Status(status),
		}
	}

	// Custom Headers
	customHeaderMap := d.Get("custom_headers").(map[string]interface{})
	var key = make([]string, 0, len(customHeaderMap))
	for k := range customHeaderMap {
		key = append(key, k)
	}
	sort.Strings(key)
	customHeaders := make([]api.Header, len(key))
	for i, k := range key {
		customHeaders[i] = api.Header{Name: k, Value: customHeaderMap[k].(string)}
	}

	webTransactionBrowserMonitor := &api.WebTransactionBrowserMonitor{
		MonitorID:         d.Id(),
		DisplayName:       d.Get("display_name").(string),
		Type:              string(api.REALBROWSER),
		BaseURL:           d.Get("base_url").(string),
		SeleniumScript:    d.Get("selenium_script").(string),
		ScriptType:        d.Get("script_type").(string),
		CheckFrequency:    d.Get("check_frequency").(string),
		AsyncDCEnabled:    d.Get("async_dc_enabled").(bool),
		BrowserType:       d.Get("browser_type").(int),
		BrowserVersion:    d.Get("browser_version").(int),
		ThinkTime:         d.Get("think_time").(int),
		PageLoadTime:      d.Get("page_load_time").(int),
		Resolution:        d.Get("resolution").(string),
		PerformAutomation: d.Get("perform_automation").(bool),
		IgnoreCertError:   d.Get("ignore_cert_err").(bool),
		IPType:            d.Get("ip_type").(int),
		UserAgent:         d.Get("user_agent").(string),
		//MatchCase:             d.Get("match_case").(bool),
		//ParallelPolling:       d.Get("parallel_polling").(bool),
		LocationProfileID:     d.Get("location_profile_id").(string),
		NotificationProfileID: d.Get("notification_profile_id").(string),
		ThresholdProfileID:    d.Get("threshold_profile_id").(string),
		OnCallScheduleID:      d.Get("on_call_schedule_id").(string),
		Cookies:               cookies,
		CustomHeaders:         customHeaders,
		DependencyResourceIDs: dependencyResourceIDs,
		MonitorGroups:         monitorGroups,
		UserGroupIDs:          userGroupIDs,
		TagIDs:                tagIDs,
		ThirdPartyServiceIDs:  thirdPartyServiceIDs,
		ActionIDs:             actionRefs,
	}
	if proxyDetails, ok := d.GetOk("proxy_details"); ok {
		webTransactionBrowserMonitor.ProxyDetails = proxyDetails.(map[string]interface{})
	}

	if authDetails, ok := d.GetOk("auth_details"); ok {
		webTransactionBrowserMonitor.AuthDetails = authDetails.(map[string]interface{})
	}
	_, locationProfileErr := site24x7.SetLocationProfile(client, d, webTransactionBrowserMonitor)
	if locationProfileErr != nil {
		return nil, locationProfileErr
	}

	// Notification Profile
	_, notificationProfileErr := site24x7.SetNotificationProfile(client, d, webTransactionBrowserMonitor)
	if notificationProfileErr != nil {
		return nil, notificationProfileErr
	}

	// User Alert Groups
	_, userAlertGroupErr := site24x7.SetUserGroup(client, d, webTransactionBrowserMonitor)
	if userAlertGroupErr != nil {
		return nil, userAlertGroupErr
	}

	// Tags
	_, tagsErr := site24x7.SetTags(client, d, webTransactionBrowserMonitor)
	if tagsErr != nil {
		return nil, tagsErr
	}
	if webTransactionBrowserMonitor.ThresholdProfileID == "" {
		profile, err := site24x7.DefaultThresholdProfile(client, api.REALBROWSER)
		if err != nil {
			return nil, err
		}
		webTransactionBrowserMonitor.ThresholdProfileID = profile.ProfileID
		d.Set("threshold_profile_id", profile)
	}
	if webTransactionBrowserMonitor.DisplayName == "" {

		urlString := d.Get("base_url").(string)

		parsedURL, err := url.Parse(urlString)
		if err != nil {
			return nil, err
		}

		// Extract domain name (without subdomain)
		hostnameParts := strings.Split(parsedURL.Hostname(), ".")
		var domain = ""
		if len(hostnameParts) >= 2 {
			domain = "RBM-" + hostnameParts[len(hostnameParts)-2] // Get the second-to-last part
			fmt.Println("Domain:", domain)
		}
		webTransactionBrowserMonitor.DisplayName = domain
		d.Set("display_name", "RBM-"+domain)
	}
	return webTransactionBrowserMonitor, nil
}
func updateWebTransactionBrowserMonitorResourceData(d *schema.ResourceData, monitor *api.WebTransactionBrowserMonitor) {
	customHeaders := make(map[string]interface{})
	for _, h := range monitor.CustomHeaders {
		if h.Name == "" {
			continue
		}
		customHeaders[h.Name] = h.Value
	}

	d.Set("display_name", monitor.DisplayName)
	d.Set("type", monitor.Type)
	d.Set("base_url", monitor.BaseURL)
	//d.Set("selenium_script", monitor.SeleniumScript)
	//d.Set("script_type", monitor.ScriptType)
	d.Set("check_frequency", monitor.CheckFrequency)
	d.Set("async_dc_enabled", monitor.AsyncDCEnabled)
	d.Set("browser_type", monitor.BrowserType)
	d.Set("browser_version", monitor.BrowserVersion)
	d.Set("think_time", monitor.ThinkTime)
	d.Set("ignore_cert_err", monitor.IgnoreCertError)
	d.Set("page_load_time", monitor.PageLoadTime)
	d.Set("resolution", monitor.Resolution)
	d.Set("location_profile_id", monitor.LocationProfileID)
	d.Set("notification_profile_id", monitor.NotificationProfileID)
	d.Set("threshold_profile_id", monitor.ThresholdProfileID)
	d.Set("custom_headers", customHeaders)
	d.Set("user_group_ids", monitor.UserGroupIDs)
	d.Set("on_call_schedule_id", monitor.OnCallScheduleID)
	d.Set("monitor_groups", monitor.MonitorGroups)
	d.Set("ip_type", monitor.IPType)
	d.Set("user_agent", monitor.UserAgent)
	d.Set("dependency_resource_ids", monitor.DependencyResourceIDs)
	d.Set("action_ids", monitor.ActionIDs)
	d.Set("third_party_services", monitor.ThirdPartyServiceIDs)
	d.Set("tag_ids", monitor.TagIDs)
	//d.Set("parallel_polling", monitor.ParallelPolling)
	d.Set("perform_automation", monitor.PerformAutomation)

	if monitor.ProxyDetails != nil {
		proxyDetailsMap := make(map[string]interface{})
		proxyDetailsMap["webProxyUrl"] = monitor.ProxyDetails["webProxyUrl"].(string)
		proxyDetailsMap["webProxyUname"] = monitor.ProxyDetails["webProxyUname"].(string)
		proxyDetailsMap["webProxyPass"] = monitor.ProxyDetails["webProxyPass"].(string)
		d.Set("proxy_details", proxyDetailsMap)
	}
	if monitor.AuthDetails != nil {
		authDetailsMap := make(map[string]interface{})
		authDetailsMap["userName"] = monitor.AuthDetails["userName"].(string)
		authDetailsMap["password"] = monitor.AuthDetails["password"].(string)
		d.Set("auth_details", authDetailsMap)
	}
	// cookies
	cookies := make(map[string]interface{})
	for _, h := range monitor.Cookies {
		if h.Name == "" {
			continue
		}
		cookies[h.Name] = h.Value
	}
	d.Set("cookies", cookies)
}
