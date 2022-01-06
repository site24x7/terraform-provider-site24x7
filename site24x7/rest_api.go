package site24x7

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
)

// Sample RESTAPI POST JSON

// {
// 	"type": "RESTAPI",
// 	"http_protocol": "H1.1",
// 	"ssl_protocol": "Auto",
// 	"response_headers_check": {
// 	  "severity": 2,
// 	  "value": [
// 		{
// 		  "name": "Allow",
// 		  "value": "a"
// 		}
// 	  ]
// 	},
// 	"match_xml": {
// 	  "xpath": [
// 		{
// 		  "name": ""
// 		}
// 	  ],
// 	  "severity": 2
// 	},
// 	"match_json": {
// 	  "jsonpath": [
// 		{
// 		  "name": ""
// 		}
// 	  ],
// 	  "severity": 2
// 	},
// 	"user_group_ids": [
// 	  "111111000000025005"
// 	],
// 	"website": "https://dummy.restapiexample.com/",
// 	"check_frequency": "5",
// 	"response_type": "T",
// 	"timeout": 30,
// 	"notification_profile_id": "111111000000029001",
// 	"http_method": "G",
// 	"auth_method": "B",
// 	"matching_keyword": {
// 	  "severity": 2,
// 	  "value": "a"
// 	},
// 	"match_case": true,
// 	"json_schema_check": false,
// 	"use_name_server": false,
// 	"use_alpn": false,
// 	"use_ipv6": false,
// 	"tag_ids": [],
// 	"location_profile_id": "111111000000025013",
// 	"threshold_profile_id": "111111000021519001",
// 	"display_name": "rest api - terraform"
//   }

var RestApiMonitorSchema = map[string]*schema.Schema{
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
	"http_method": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "G",
		Description: "HTTP Method used for accessing the website. Default value is G.",
	},
	"http_protocol": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "H1.1",
		Description: "Specify the version of the HTTP protocol. Default value is H1.1.",
	},
	"ssl_protocol": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "Auto",
		Description: "Specify the version of the SSL protocol. If you are not sure about the version, use Auto.",
	},
	"use_ipv6": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Select IPv6 for monitoring the websites hosted with IPv6 address. If you choose non IPv6 supported locations, monitoring will happen through IPv4.",
	},
	"request_content_type": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Provide content type for request params.",
	},
	"response_content_type": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "T",
		Description: "Response content type.",
	},
	"request_param": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Provide parameters to be passed while accessing the website.",
	},
	"auth_user": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Authentication password to access the website.",
	},
	"auth_pass": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Authentication user name to access the website.",
	},
	"oauth2_provider": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Provider ID of the OAuth Provider to be associated with the monitor.",
	},
	"client_certificate_password": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Password of the uploaded client certificate.",
	},
	"jwt_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Token ID of the Web Token to be associated with the monitor.",
	},
	"matching_keyword": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{0, 2}), // Trouble or Down
				},
				"value": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
		Description: "",
	},
	"unmatching_keyword": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{0, 2}), // Trouble or Down
				},
				"value": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
		Description: "",
	},
	"match_regex": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"severity": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntInSlice([]int{0, 2}), // Trouble or Down
				},
				"value": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
		Description: "",
	},
	"match_case": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Perform case sensitive keyword search or not.",
	},
	"json_schema_check": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Enable this option to perform the JSON schema check.",
	},
	"use_name_server": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Resolve the IP address using Domain Name Server.",
	},
	"use_alpn": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Description: "Enable ALPN to send supported protocols as part of the TLS handshake.",
	},
	"user_agent": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "User Agent to be used while monitoring the website.",
	},
	"custom_headers": {
		Type:        schema.TypeMap,
		Optional:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "A Map of Header name and value.",
	},
	"response_headers_severity": {
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      2,
		ValidateFunc: validation.IntInSlice([]int{0, 2}), // 0 - Down, 2 - Trouble
		Description:  "Alert type constant. Can be either 0 or 2. '0' denotes Down and '2' denotes Trouble",
	},
	"response_headers": {
		Type:        schema.TypeMap,
		Optional:    true,
		Description: "A Map of Header name and value.",
	},
}

func resourceSite24x7RestApiMonitor() *schema.Resource {
	return &schema.Resource{
		Create: restApiMonitorCreate,
		Read:   restApiMonitorRead,
		Update: restApiMonitorUpdate,
		Delete: restApiMonitorDelete,
		Exists: restApiMonitorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: RestApiMonitorSchema,
	}
}

func restApiMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	restApiMonitor, err := resourceDataToRestApiMonitor(d, client)
	if err != nil {
		return err
	}

	restApiMonitor, err = client.RestApiMonitors().Create(restApiMonitor)
	if err != nil {
		return err
	}

	d.SetId(restApiMonitor.MonitorID)

	// return restApiMonitorRead(d, meta)
	return nil
}

func restApiMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	restApiMonitor, err := client.RestApiMonitors().Get(d.Id())

	if err != nil {
		return err
	}

	updateRestApiMonitorResourceData(d, restApiMonitor)

	return nil
}

func restApiMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	restApiMonitor, err := resourceDataToRestApiMonitor(d, client)

	if err != nil {
		return err
	}

	restApiMonitor, err = client.RestApiMonitors().Update(restApiMonitor)
	if err != nil {
		return err
	}

	d.SetId(restApiMonitor.MonitorID)

	// return restApiMonitorRead(d, meta)
	return nil
}

func restApiMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	err := client.RestApiMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func restApiMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(Client)

	_, err := client.RestApiMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToRestApiMonitor(d *schema.ResourceData, client Client) (*api.RestApiMonitor, error) {

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

	var actionRefs []api.ActionRef
	if actionData, ok := d.GetOk("actions"); ok {
		actionMap := actionData.(map[string]interface{})
		actionKeys := make([]string, 0, len(actionMap))
		for k := range actionMap {
			actionKeys = append(actionKeys, k)
		}
		sort.Strings(actionKeys)
		actionRefs := make([]api.ActionRef, len(actionKeys))
		for i, k := range actionKeys {
			status, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}

			actionRefs[i] = api.ActionRef{
				ActionID:  actionMap[k].(string),
				AlertType: api.Status(status),
			}
		}
	}

	// Custom Headers
	customHeaderMap := d.Get("custom_headers").(map[string]interface{})
	customHeaderKeys := make([]string, 0, len(customHeaderMap))
	for k := range customHeaderMap {
		customHeaderKeys = append(customHeaderKeys, k)
	}
	sort.Strings(customHeaderKeys)
	customHeaders := make([]api.Header, len(customHeaderKeys))
	for i, k := range customHeaderKeys {
		customHeaders[i] = api.Header{Name: k, Value: customHeaderMap[k].(string)}
	}

	// HTTP Response Headers
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

	restApiMonitor := &api.RestApiMonitor{
		MonitorID:      d.Id(),
		DisplayName:    d.Get("display_name").(string),
		Type:           string(api.RESTAPI),
		Website:        d.Get("website").(string),
		CheckFrequency: strconv.Itoa(d.Get("check_frequency").(int)),
		Timeout:        d.Get("timeout").(int),
		HttpMethod:     d.Get("http_method").(string),
		HttpProtocol:   d.Get("http_protocol").(string),
		SslProtocol:    d.Get("ssl_protocol").(string),
		UseAlpn:        d.Get("use_alpn").(bool),
		UseIPV6:        d.Get("use_ipv6").(bool),

		RequestContentType:        d.Get("request_content_type").(string),
		ResponseContentType:       d.Get("response_content_type").(string),
		RequestParam:              d.Get("request_param").(string),
		OAuth2Provider:            d.Get("oauth2_provider").(string),
		ClientCertificatePassword: d.Get("client_certificate_password").(string),
		JwtID:                     d.Get("jwt_id").(string),

		AuthUser:              d.Get("auth_user").(string),
		AuthPass:              d.Get("auth_pass").(string),
		MatchCase:             d.Get("match_case").(bool),
		JSONSchemaCheck:       d.Get("json_schema_check").(bool),
		UseNameServer:         d.Get("use_name_server").(bool),
		UserAgent:             d.Get("user_agent").(string),
		CustomHeaders:         customHeaders,
		ResponseHeaders:       httpResponseHeader,
		LocationProfileID:     d.Get("location_profile_id").(string),
		NotificationProfileID: d.Get("notification_profile_id").(string),
		ThresholdProfileID:    d.Get("threshold_profile_id").(string),
		MonitorGroups:         monitorGroups,
		UserGroupIDs:          userGroupIDs,
		TagIDs:                tagIDs,
		ThirdPartyServiceIDs:  thirdPartyServiceIDs,
		ActionIDs:             actionRefs,
	}

	if matchingRegex, ok := d.GetOk("match_regex"); ok {
		restApiMonitor.MatchRegex = matchingRegex.(map[string]interface{})
	}

	if matchingKeyword, ok := d.GetOk("matching_keyword"); ok {
		restApiMonitor.MatchingKeyword = matchingKeyword.(map[string]interface{})
	}

	if unmatchingKeyword, ok := d.GetOk("unmatching_keyword"); ok {
		restApiMonitor.UnmatchingKeyword = unmatchingKeyword.(map[string]interface{})
	}

	if restApiMonitor.LocationProfileID == "" {
		locationProfileNameToMatch := d.Get("location_profile_name").(string)
		profile, err := DefaultLocationProfile(client, locationProfileNameToMatch)
		if err != nil {
			return nil, err
		}
		restApiMonitor.LocationProfileID = profile.ProfileID
		d.Set("location_profile_id", profile.ProfileID)
	}

	// Notification Profile
	_, notificationProfileErr := SetNotificationProfile(client, d, restApiMonitor)
	if notificationProfileErr != nil {
		return nil, notificationProfileErr
	}

	// User Alert Groups
	_, userAlertGroupErr := SetUserGroup(client, d, restApiMonitor)
	if userAlertGroupErr != nil {
		return nil, userAlertGroupErr
	}

	// Tags
	_, tagsErr := SetTags(client, d, restApiMonitor)
	if tagsErr != nil {
		return nil, tagsErr
	}

	if restApiMonitor.ThresholdProfileID == "" {
		profile, err := DefaultThresholdProfile(client, api.RESTAPI)
		if err != nil {
			return nil, err
		}
		restApiMonitor.ThresholdProfileID = profile.ProfileID
		d.Set("threshold_profile_id", profile)
	}

	return restApiMonitor, nil
}

func updateRestApiMonitorResourceData(d *schema.ResourceData, monitor *api.RestApiMonitor) {
	d.Set("display_name", monitor.DisplayName)
	d.Set("type", monitor.Type)
	d.Set("website", monitor.Website)
	d.Set("check_frequency", monitor.CheckFrequency)
	d.Set("timeout", monitor.Timeout)
	d.Set("http_method", monitor.HttpMethod)
	d.Set("http_protocol", monitor.HttpProtocol)
	d.Set("ssl_protocol", monitor.SslProtocol)
	d.Set("use_alpn", monitor.UseAlpn)
	d.Set("use_ipv6", monitor.UseIPV6)
	d.Set("request_content_type", monitor.RequestContentType)
	d.Set("response_content_type", monitor.ResponseContentType)
	d.Set("request_param", monitor.RequestParam)
	d.Set("auth_user", monitor.AuthUser)
	d.Set("auth_pass", monitor.AuthPass)
	d.Set("oauth2_provider", monitor.OAuth2Provider)
	d.Set("client_certificate_password", monitor.ClientCertificatePassword)
	d.Set("jwt_id", monitor.JwtID)

	d.Set("location_profile_id", monitor.LocationProfileID)
	d.Set("notification_profile_id", monitor.NotificationProfileID)
	d.Set("threshold_profile_id", monitor.ThresholdProfileID)
	d.Set("monitor_groups", monitor.MonitorGroups)
	d.Set("user_group_ids", monitor.UserGroupIDs)
	d.Set("tag_ids", monitor.TagIDs)
	d.Set("third_party_service_ids", monitor.ThirdPartyServiceIDs)
	// if monitor.MatchingKeyword != nil {
	// 	d.Set("matching_keyword", monitor.MatchingKeyword)
	// }
	// if monitor.UnmatchingKeyword != nil {
	// 	d.Set("unmatching_keyword", monitor.UnmatchingKeyword)
	// }
	// if monitor.MatchRegex != nil {
	// 	d.Set("match_regex", monitor.MatchRegex)
	// }

	if monitor.MatchingKeyword != nil {
		matchingKeywordMap := make(map[string]interface{})
		matchingKeywordMap["severity"] = int(monitor.MatchingKeyword["severity"].(float64))
		matchingKeywordMap["value"] = monitor.MatchingKeyword["value"].(string)
		d.Set("matching_keyword", matchingKeywordMap)
	}
	if monitor.UnmatchingKeyword != nil {
		unmatchingKeywordMap := make(map[string]interface{})
		unmatchingKeywordMap["severity"] = int(monitor.UnmatchingKeyword["severity"].(float64))
		unmatchingKeywordMap["value"] = monitor.UnmatchingKeyword["value"].(string)
		d.Set("unmatching_keyword", unmatchingKeywordMap)
	}
	if monitor.MatchRegex != nil {
		matchRegexMap := make(map[string]interface{})
		matchRegexMap["severity"] = int(monitor.MatchRegex["severity"].(float64))
		matchRegexMap["value"] = monitor.MatchRegex["value"].(string)
		d.Set("match_regex", matchRegexMap)
	}

	d.Set("match_case", monitor.MatchCase)
	d.Set("json_schema_check", monitor.JSONSchemaCheck)
	d.Set("use_name_server", monitor.UseNameServer)
	d.Set("user_agent", monitor.UserAgent)

	// Custom Headers
	customHeaders := make(map[string]interface{})
	for _, h := range monitor.CustomHeaders {
		if h.Name == "" {
			continue
		}
		customHeaders[h.Name] = h.Value
	}
	d.Set("custom_headers", customHeaders)

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

	actions := make(map[string]interface{})
	for _, action := range monitor.ActionIDs {
		actions[fmt.Sprintf("%d", action.AlertType)] = action.ActionID
	}

	d.Set("actions", actions)
}
