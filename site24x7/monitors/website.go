package monitors

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

// SAMPLE POST JSON
// {
// 	"website": "http://www.zoho.com",
// 	"use_ipv6": false,
// 	"http_protocol": "H1.1",
// 	"check_frequency": "5",
// 	"unmatching_keyword": {
// 	"value": "unmatching_bbbb",
// 	"severity": 2
// 	},
// 	"threshold_profile_id": "123456000000815001",
// 	"type": "URL",
// 	"display_name": "google",
// 	"user_group_ids": [
// 	"123456000000025005"
// 	],
// 	"timeout": 15,
// 	"match_case": false,
// 	"response_headers_check": {
// 	"value": [
// 		{
// 		"name": "Accept-Patch",
// 		"value": "b"
// 		},
// 		{
// 		"name": "Allow",
// 		"value": "a"
// 		}
// 	],
// 	"severity": 2
// 	},
// 	"auth_method": "B",
// 	"http_method": "G",
// 	"perform_automation": false,
// 	"use_name_server": false,
// 	"use_alpn": false,
// 	"matching_keyword": {
// 	"value": "matching_aaaa",
// 	"severity": 2
// 	},
// 	"notification_profile_id": "123456000000029001",
// 	"location_profile_id": "123456000000025021",
// 	"ssl_protocol": "Auto",
// 	"custom_headers": [
// 	{
// 		"name": "Accept-Language",
// 		"value": "english"
// 	},
// 	{
// 		"name": "Accept-Datetime",
// 		"value": "dateheader"
// 	}
// 	],
// 	"tag_ids": [],
//  "third_party_services": [
//     "123456000024411001"
//  ]
// }

var websiteMonitorSchema = map[string]*schema.Schema{
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
		Default:     "1",
		Description: "Interval at which your website has to be monitored. Default value is 1 minute.",
	},
	"timeout": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     10,
		Description: "Timeout for connecting to website. Default value is 10. Range 1 - 45.",
	},
	"use_ipv6": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Monitoring is performed over IPv6 from supported locations. IPv6 locations do not fall back to IPv4 on failure.",
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
	"match_case": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Perform case sensitive keyword search or not.",
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
	"response_headers": {
		Type:        schema.TypeMap,
		Optional:    true,
		Description: "A Map of respone header name and value.",
	},
	"response_headers_severity": {
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      2,
		ValidateFunc: validation.IntInSlice([]int{0, 2}), // 0 - Down, 2 - Trouble
		Description:  "Alert type constant. Can be either 0 or 2. '0' denotes Down and '2' denotes Trouble",
	},
	// HTTP Configuration
	"http_method": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "G",
		Description: "HTTP Method to be used for accessing the website. Default value is 'G'. 'G' denotes GET, 'P' denotes POST and 'H' denotes HEAD. PUT, PATCH and DELETE are not supported.",
	},
	"request_content_type": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Provide content type for request params when http_method is 'P'. 'J' denotes JSON, 'T' denotes TEXT, 'X' denotes XML and 'F' denotes FORM",
	},
	"request_body": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Provide the content to be passed in the request body while accessing the website.",
	},
	"request_headers": {
		Type:        schema.TypeMap,
		Optional:    true,
		Description: "A Map of request header name and value.",
	},
	"user_agent": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "User Agent to be used while monitoring the website.",
	},
	"auth_method": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "B",
		Description: "Authentication method to access the website. Default value is 'B'. 'B' denotes Basic/NTLM. 'O' denotes OAuth 2 and 'W' denotes Web Token.",
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
		DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			// Suppress diff - Password in API response is encrypted.
			return true
		},
	},
	"credential_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Credential Profile to associate.",
	},
	"client_certificate_password": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Password of the client certificate.",
		DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			// Suppress diff - Password in API response is encrypted.
			return true
		},
	},
	"use_name_server": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Resolve the IP address using Domain Name Server.",
	},
	"forced_ips": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Provide the domain name or IP addresses to be used for monitoring instead of using the IPs resolved from the given URL.",
	},
	"up_status_codes": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "",
		Description: "Provide a comma-separated list of HTTP status codes that indicate a successful response. You can specify individual status codes, as well as ranges separated with a colon.",
	},
	"ssl_protocol": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "Auto",
		Description: "Specify the version of the SSL protocol. If you are not sure about the version, use Auto.",
	},
	"http_protocol": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "H1.1",
		Description: "Specify the version of the HTTP protocol. Default value is H1.1.",
	},
	"use_alpn": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Description: "Enable ALPN to send supported protocols as part of the TLS handshake.",
	},
	// Configuration Profiles
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

func ResourceSite24x7WebsiteMonitor() *schema.Resource {
	return &schema.Resource{
		Create: websiteMonitorCreate,
		Read:   websiteMonitorRead,
		Update: websiteMonitorUpdate,
		Delete: websiteMonitorDelete,
		Exists: websiteMonitorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: websiteMonitorSchema,
	}
}

func websiteMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

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
	client := meta.(site24x7.Client)

	websiteMonitor, err := client.WebsiteMonitors().Get(d.Id())
	if err != nil {
		return err
	}

	updateWebsiteMonitorResourceData(d, websiteMonitor)

	return nil
}

func websiteMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

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
	client := meta.(site24x7.Client)

	err := client.WebsiteMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func websiteMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.WebsiteMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToWebsiteMonitor(d *schema.ResourceData, client site24x7.Client) (*api.WebsiteMonitor, error) {

	websiteMonitor := &api.WebsiteMonitor{
		MonitorID:      d.Id(),
		DisplayName:    d.Get("display_name").(string),
		Type:           string(api.URL),
		Website:        d.Get("website").(string),
		CheckFrequency: d.Get("check_frequency").(string),
		Timeout:        d.Get("timeout").(int),
	}

	/*  ================================ Content Checks ================================
	MatchingKeyword   *ValueAndSeverity  `json:"matching_keyword,omitempty"`
	UnmatchingKeyword *ValueAndSeverity  `json:"unmatching_keyword,omitempty"`
	MatchCase         bool               `json:"match_case"`
	MatchRegex        *ValueAndSeverity  `json:"match_regex,omitempty"`
	ResponseHeaders   HTTPResponseHeader `json:"response_headers_check,omitempty"`
	*/
	if _, ok := d.GetOk("matching_keyword_value"); ok {
		websiteMonitor.MatchingKeyword = &api.ValueAndSeverity{
			Value:    d.Get("matching_keyword_value").(string),
			Severity: api.Status(d.Get("matching_keyword_severity").(int)),
		}
	}
	if _, ok := d.GetOk("unmatching_keyword_value"); ok {
		websiteMonitor.UnmatchingKeyword = &api.ValueAndSeverity{
			Value:    d.Get("unmatching_keyword_value").(string),
			Severity: api.Status(d.Get("unmatching_keyword_severity").(int)),
		}
	}
	websiteMonitor.MatchCase = d.Get("match_case").(bool)
	if _, ok := d.GetOk("match_regex_value"); ok {
		websiteMonitor.MatchRegex = &api.ValueAndSeverity{
			Value:    d.Get("match_regex_value").(string),
			Severity: api.Status(d.Get("match_regex_severity").(int)),
		}
	}
	var httpResponseHeader api.HTTPResponseHeader
	responseHeaderMap := d.Get("response_headers").(map[string]interface{})
	if len(responseHeaderMap) > 0 {
		reponseHeaderKeys := make([]string, 0, len(responseHeaderMap))
		for k := range responseHeaderMap {
			reponseHeaderKeys = append(reponseHeaderKeys, k)
		}
		sort.Strings(reponseHeaderKeys)
		responseHeaders := make([]api.Header, len(reponseHeaderKeys))
		for i, k := range reponseHeaderKeys {
			responseHeaders[i] = api.Header{Name: k, Value: responseHeaderMap[k].(string)}
		}
		httpResponseHeader.Severity = api.Status(d.Get("response_headers_severity").(int))
		httpResponseHeader.Value = responseHeaders
	}
	websiteMonitor.ResponseHeaders = httpResponseHeader

	/* ================================ HTTP Configuration ================================
	HTTPMethod                string   `json:"http_method"`
	RequestContentType        string   `json:"request_content_type,omitempty"`
	RequestBody               string   `json:"request_param,omitempty"`
	RequestHeaders            []Header `json:"custom_headers,omitempty"`
	UserAgent                 string   `json:"user_agent,omitempty"`
	AuthMethod                string   `json:"auth_method,omitempty"`
	AuthUser                  string   `json:"auth_user,omitempty"`
	AuthPass                  string   `json:"auth_pass,omitempty"`
	CredentialProfileID       string   `json:"credential_profile_id,omitempty"`
	ClientCertificatePassword string   `json:"client_certificate_password,omitempty"`
	UseNameServer             bool     `json:"use_name_server,omitempty"`
	ForcedIPs                 string   `json:"forced_ips,omitempty"`
	UpStatusCodes             string   `json:"up_status_codes,omitempty"`
	SSLProtocol               string   `json:"ssl_protocol,omitempty"`
	HTTPProtocol              string   `json:"http_protocol,omitempty"`
	UseAlpn                   bool     `json:"use_alpn"`
	*/
	websiteMonitor.HTTPMethod = d.Get("http_method").(string)
	websiteMonitor.RequestContentType = d.Get("request_content_type").(string)
	websiteMonitor.RequestBody = d.Get("request_body").(string)
	requestHeaderMap := d.Get("request_headers").(map[string]interface{})
	keys := make([]string, 0, len(requestHeaderMap))
	for k := range requestHeaderMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	requestHeaders := make([]api.Header, len(keys))
	for i, k := range keys {
		requestHeaders[i] = api.Header{Name: k, Value: requestHeaderMap[k].(string)}
	}
	websiteMonitor.RequestHeaders = requestHeaders
	websiteMonitor.UserAgent = d.Get("user_agent").(string)
	websiteMonitor.AuthMethod = d.Get("auth_method").(string)
	websiteMonitor.AuthUser = d.Get("auth_user").(string)
	websiteMonitor.AuthPass = d.Get("auth_pass").(string)
	websiteMonitor.CredentialProfileID = d.Get("credential_profile_id").(string)
	websiteMonitor.ClientCertificatePassword = d.Get("client_certificate_password").(string)
	websiteMonitor.UseNameServer = d.Get("use_name_server").(bool)
	websiteMonitor.ForcedIPs = d.Get("forced_ips").(string)
	websiteMonitor.UpStatusCodes = d.Get("up_status_codes").(string)
	websiteMonitor.SSLProtocol = d.Get("ssl_protocol").(string)
	websiteMonitor.HTTPProtocol = d.Get("http_protocol").(string)
	websiteMonitor.UseAlpn = d.Get("use_alpn").(bool)

	// ================================ Configuration Profiles ================================
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

	websiteMonitor.LocationProfileID = d.Get("location_profile_id").(string)
	websiteMonitor.NotificationProfileID = d.Get("notification_profile_id").(string)
	websiteMonitor.ThresholdProfileID = d.Get("threshold_profile_id").(string)
	websiteMonitor.MonitorGroups = monitorGroups
	websiteMonitor.DependencyResourceIDs = dependencyResourceIDs
	websiteMonitor.UserGroupIDs = userGroupIDs
	websiteMonitor.TagIDs = tagIDs
	websiteMonitor.ThirdPartyServiceIDs = thirdPartyServiceIDs
	websiteMonitor.ActionIDs = actionRefs
	// Location Profile
	if websiteMonitor.LocationProfileID == "" {
		locationProfileNameToMatch := d.Get("location_profile_name").(string)
		profile, err := site24x7.DefaultLocationProfile(client, locationProfileNameToMatch)
		if err != nil {
			return nil, err
		}
		websiteMonitor.LocationProfileID = profile.ProfileID
		d.Set("location_profile_id", profile.ProfileID)
	}
	// Threshold Profile
	if websiteMonitor.ThresholdProfileID == "" {
		profile, err := site24x7.DefaultThresholdProfile(client, api.URL)
		if err != nil {
			return nil, err
		}
		websiteMonitor.ThresholdProfileID = profile.ProfileID
		d.Set("threshold_profile_id", profile)
	}
	// Notification Profile
	_, notificationProfileErr := site24x7.SetNotificationProfile(client, d, websiteMonitor)
	if notificationProfileErr != nil {
		return nil, notificationProfileErr
	}
	// User Alert Groups
	_, userAlertGroupErr := site24x7.SetUserGroup(client, d, websiteMonitor)
	if userAlertGroupErr != nil {
		return nil, userAlertGroupErr
	}
	// Tags
	_, tagsErr := site24x7.SetTags(client, d, websiteMonitor)
	if tagsErr != nil {
		return nil, tagsErr
	}

	return websiteMonitor, nil
}

func updateWebsiteMonitorResourceData(d *schema.ResourceData, monitor *api.WebsiteMonitor) {
	d.Set("display_name", monitor.DisplayName)
	d.Set("type", monitor.Type)
	d.Set("website", monitor.Website)
	d.Set("check_frequency", monitor.CheckFrequency)
	d.Set("timeout", monitor.Timeout)
	d.Set("use_ipv6", monitor.UseIPV6)
	// ================================ HTTP Configuration ================================
	d.Set("http_method", monitor.HTTPMethod)
	d.Set("request_content_type", monitor.RequestContentType)
	d.Set("request_body", monitor.RequestBody)
	requestHeaders := make(map[string]interface{})
	for _, h := range monitor.RequestHeaders {
		if h.Name == "" {
			continue
		}
		requestHeaders[h.Name] = h.Value
	}
	d.Set("request_headers", requestHeaders)
	d.Set("user_agent", monitor.UserAgent)
	d.Set("auth_method", monitor.AuthMethod)
	d.Set("auth_user", monitor.AuthUser)
	d.Set("auth_pass", monitor.AuthPass)
	d.Set("credential_profile_id", monitor.CredentialProfileID)
	d.Set("client_certificate_password", monitor.ClientCertificatePassword)
	d.Set("use_name_server", monitor.UseNameServer)
	d.Set("forced_ips", monitor.ForcedIPs)
	d.Set("up_status_codes", monitor.UpStatusCodes)
	d.Set("ssl_protocol", monitor.SSLProtocol)
	d.Set("http_protocol", monitor.HTTPProtocol)
	d.Set("use_alpn", monitor.UseAlpn)
	// ================================ Content Checks ================================
	if monitor.MatchingKeyword != nil {
		d.Set("matching_keyword_value", monitor.MatchingKeyword.Value)
		d.Set("matching_keyword_severity", monitor.MatchingKeyword.Severity)
	}
	if monitor.UnmatchingKeyword != nil {
		d.Set("unmatching_keyword_value", monitor.UnmatchingKeyword.Value)
		d.Set("unmatching_keyword_severity", monitor.UnmatchingKeyword.Severity)
	}
	d.Set("match_case", monitor.MatchCase)
	if monitor.MatchRegex != nil {
		d.Set("match_regex_value", monitor.MatchRegex.Value)
		d.Set("match_regex_severity", monitor.MatchRegex.Severity)
	}
	// Response Headers
	responseHeaders := make(map[string]interface{})
	for _, h := range monitor.ResponseHeaders.Value {
		if h.Name == "" {
			continue
		}
		responseHeaders[h.Name] = h.Value
	}
	d.Set("response_headers", responseHeaders)
	d.Set("response_headers_severity", monitor.ResponseHeaders.Severity)
	// ================================ Configuration Profiles ================================
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
