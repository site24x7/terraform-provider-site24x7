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
	// Content Check
	"response_content_type": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "T",
		Description: "Response content type. Default value is 'T'. 'J' denotes JSON, 'T' denotes TEXT, 'X' denotes XML",
	},
	"match_json_path": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "Provide multiple JSON Path expressions to enable evaluation of JSON Path expression assertions. The assertions must successfully parse the JSON Path in the JSON. JSON expression assertions fails if the expressions does not match.",
	},
	"match_json_path_severity": {
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      2,
		ValidateFunc: validation.IntInSlice([]int{0, 2}), // 0 - Down, 2 - Trouble
		Description:  "Trigger an alert when the JSON path assertion fails during a test. Alert type constant. Can be either 0 or 2. '0' denotes Down and '2' denotes Trouble. Default value is 2.",
	},
	"json_schema": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "JSON schema to be validated against the JSON response.",
	},
	"json_schema_severity": {
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      2,
		ValidateFunc: validation.IntInSlice([]int{0, 2}), // 0 - Down, 2 - Trouble
		Description:  "Trigger an alert when the JSON schema assertion fails during a test. Alert type constant. Can be either 0 or 2. '0' denotes Down and '2' denotes Trouble. Default value is 2.",
	},
	"json_schema_check": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "JSON Schema check allows you to annotate and validate all JSON endpoints for your web service.",
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
		Description: "Check for the keyword in the website response.",
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
		Description: "Check for non existence of keyword in the website response.",
	},
	"match_case": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Perform case sensitive keyword search or not.",
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
		Description: "Match the regular expression in the website response.",
	},
	"response_headers": {
		Type:        schema.TypeMap,
		Optional:    true,
		Description: "A Map of Header name and value.",
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
	"graphql_query": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Provide the GraphQL query to get specific response from GraphQL based API service.",
	},
	"graphql_variables": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Provide the GraphQL variables to get specific response from GraphQL based API service.",
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
	"oauth2_provider": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Provider ID of the OAuth Provider to be associated with the monitor.",
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
	"jwt_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Token ID of the Web Token to be associated with the monitor.",
	},
	"use_name_server": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Resolve the IP address using Domain Name Server.",
	},
	"up_status_codes": {
		Type:        schema.TypeString,
		Optional:    true,
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
		Type: schema.TypeSet,
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

func ResourceSite24x7RestApiMonitor() *schema.Resource {
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
	client := meta.(site24x7.Client)

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
	client := meta.(site24x7.Client)

	restApiMonitor, err := client.RestApiMonitors().Get(d.Id())

	if err != nil {
		return err
	}

	updateRestApiMonitorResourceData(d, restApiMonitor)

	return nil
}

func restApiMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

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
	client := meta.(site24x7.Client)

	err := client.RestApiMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func restApiMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.RestApiMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToRestApiMonitor(d *schema.ResourceData, client site24x7.Client) (*api.RestApiMonitor, error) {

	var monitorGroups []string
	for _, group := range d.Get("monitor_groups").([]interface{}) {
		if group != nil {
			monitorGroups = append(monitorGroups, group.(string))
		}
	}

	dependencyIDs := d.Get("dependency_resource_ids").(*schema.Set).List()
	dependencyResourceIDs := make([]string, 0, len(dependencyIDs))
	for _, dependencyResourceID := range dependencyIDs {
		if dependencyResourceID != nil {
			dependencyResourceIDs = append(dependencyResourceIDs, dependencyResourceID.(string))
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

	var thirdPartyServiceIDs []string
	for _, id := range d.Get("third_party_service_ids").([]interface{}) {
		if id != nil {
			thirdPartyServiceIDs = append(thirdPartyServiceIDs, id.(string))
		}
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

	// Request Headers
	requestHeaderMap := d.Get("request_headers").(map[string]interface{})
	requestHeaderKeys := make([]string, 0, len(requestHeaderMap))
	for k := range requestHeaderMap {
		requestHeaderKeys = append(requestHeaderKeys, k)
	}
	sort.Strings(requestHeaderKeys)
	requestHeaders := make([]api.Header, len(requestHeaderKeys))
	for i, k := range requestHeaderKeys {
		requestHeaders[i] = api.Header{Name: k, Value: requestHeaderMap[k].(string)}
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
		CheckFrequency: d.Get("check_frequency").(string),
		Timeout:        d.Get("timeout").(int),
		HTTPMethod:     d.Get("http_method").(string),
		HTTPProtocol:   d.Get("http_protocol").(string),
		SSLProtocol:    d.Get("ssl_protocol").(string),
		UseAlpn:        d.Get("use_alpn").(bool),
		UseIPV6:        d.Get("use_ipv6").(bool),

		RequestContentType:        d.Get("request_content_type").(string),
		ResponseContentType:       d.Get("response_content_type").(string),
		RequestBody:               d.Get("request_body").(string),
		OAuth2Provider:            d.Get("oauth2_provider").(string),
		ClientCertificatePassword: d.Get("client_certificate_password").(string),
		JwtID:                     d.Get("jwt_id").(string),
		AuthMethod:                d.Get("auth_method").(string),
		AuthUser:                  d.Get("auth_user").(string),
		AuthPass:                  d.Get("auth_pass").(string),
		MatchCase:                 d.Get("match_case").(bool),
		JSONSchemaCheck:           d.Get("json_schema_check").(bool),
		UseNameServer:             d.Get("use_name_server").(bool),
		UserAgent:                 d.Get("user_agent").(string),
		RequestHeaders:            requestHeaders,
		ResponseHeaders:           httpResponseHeader,
		LocationProfileID:         d.Get("location_profile_id").(string),
		NotificationProfileID:     d.Get("notification_profile_id").(string),
		ThresholdProfileID:        d.Get("threshold_profile_id").(string),
		MonitorGroups:             monitorGroups,
		DependencyResourceIDs:     dependencyResourceIDs,
		UserGroupIDs:              userGroupIDs,
		TagIDs:                    tagIDs,
		ThirdPartyServiceIDs:      thirdPartyServiceIDs,
		ActionIDs:                 actionRefs,
		// HTTP Configuration
		UpStatusCodes: d.Get("up_status_codes").(string),
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

	if matchJSONPath, ok := d.GetOk("match_json_path"); ok {
		var jsonPathList []map[string]interface{}
		for _, jsonPath := range matchJSONPath.([]interface{}) {
			matchPathMap := make(map[string]interface{})
			matchPathMap["name"] = jsonPath.(string)
			jsonPathList = append(jsonPathList, matchPathMap)
		}
		matchJSONData := make(map[string]interface{})
		matchJSONData["jsonpath"] = jsonPathList
		matchJSONData["severity"] = d.Get("match_json_path_severity").(int)
		restApiMonitor.MatchJSON = matchJSONData
	}

	if jsonSchema, ok := d.GetOk("json_schema"); ok {
		jsonSchemaData := make(map[string]interface{})
		jsonSchemaData["severity"] = d.Get("json_schema_severity").(int)
		jsonSchemaData["schema_value"] = jsonSchema.(string)
		restApiMonitor.JSONSchema = jsonSchemaData
	}

	if graphqlQuery, ok := d.GetOk("graphql_query"); ok {
		graphqlMap := make(map[string]interface{})
		graphqlMap["query"] = graphqlQuery.(string)
		graphqlMap["variables"] = d.Get("graphql_variables").(string)
		restApiMonitor.GraphQL = graphqlMap
	}

	// Location Profile
	_, locationProfileErr := site24x7.SetLocationProfile(client, d, restApiMonitor)
	if locationProfileErr != nil {
		return nil, locationProfileErr
	}

	// Notification Profile
	_, notificationProfileErr := site24x7.SetNotificationProfile(client, d, restApiMonitor)
	if notificationProfileErr != nil {
		return nil, notificationProfileErr
	}

	// User Alert Groups
	_, userAlertGroupErr := site24x7.SetUserGroup(client, d, restApiMonitor)
	if userAlertGroupErr != nil {
		return nil, userAlertGroupErr
	}

	// Tags
	_, tagsErr := site24x7.SetTags(client, d, restApiMonitor)
	if tagsErr != nil {
		return nil, tagsErr
	}

	if restApiMonitor.ThresholdProfileID == "" {
		profile, err := site24x7.DefaultThresholdProfile(client, api.RESTAPI)
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
	d.Set("http_method", monitor.HTTPMethod)
	d.Set("http_protocol", monitor.HTTPProtocol)
	d.Set("ssl_protocol", monitor.SSLProtocol)
	d.Set("use_alpn", monitor.UseAlpn)
	d.Set("use_ipv6", monitor.UseIPV6)
	d.Set("request_content_type", monitor.RequestContentType)
	d.Set("response_content_type", monitor.ResponseContentType)
	d.Set("request_body", monitor.RequestBody)
	d.Set("auth_method", monitor.AuthMethod)
	d.Set("auth_user", monitor.AuthUser)
	d.Set("auth_pass", monitor.AuthPass)
	d.Set("oauth2_provider", monitor.OAuth2Provider)
	d.Set("client_certificate_password", monitor.ClientCertificatePassword)
	d.Set("jwt_id", monitor.JwtID)

	d.Set("location_profile_id", monitor.LocationProfileID)
	d.Set("notification_profile_id", monitor.NotificationProfileID)
	d.Set("threshold_profile_id", monitor.ThresholdProfileID)
	d.Set("monitor_groups", monitor.MonitorGroups)
	d.Set("dependency_resource_ids", monitor.DependencyResourceIDs)
	d.Set("user_group_ids", monitor.UserGroupIDs)
	d.Set("tag_ids", monitor.TagIDs)
	d.Set("third_party_service_ids", monitor.ThirdPartyServiceIDs)

	if monitor.GraphQL != nil {
		d.Set("graphql_query", monitor.GraphQL["query"].(string))
		d.Set("graphql_variables", monitor.GraphQL["variables"].(string))
	}

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
	if monitor.MatchJSON != nil {
		d.Set("match_json_path_severity", int(monitor.MatchJSON["severity"].(float64)))
		jsonPathMapArr := monitor.MatchJSON["jsonpath"].([]interface{})
		var jsonPathArr []string
		for _, jsonPathData := range jsonPathMapArr {
			jsonPathMap := jsonPathData.(map[string]interface{})
			jsonPath := jsonPathMap["name"].(string)
			jsonPathArr = append(jsonPathArr, jsonPath)
		}
		d.Set("match_json_path", jsonPathArr)
	}
	if monitor.JSONSchema != nil {
		d.Set("json_schema", monitor.JSONSchema["schema_value"].(string))
		d.Set("json_schema_severity", int(monitor.JSONSchema["severity"].(float64)))
	}

	d.Set("match_case", monitor.MatchCase)
	d.Set("json_schema_check", monitor.JSONSchemaCheck)
	d.Set("use_name_server", monitor.UseNameServer)
	d.Set("user_agent", monitor.UserAgent)

	// Request Headers
	requestHeaders := make(map[string]interface{})
	for _, h := range monitor.RequestHeaders {
		if h.Name == "" {
			continue
		}
		requestHeaders[h.Name] = h.Value
	}
	d.Set("request_headers", requestHeaders)

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

	// HTTP Configuration
	d.Set("up_status_codes", monitor.UpStatusCodes)
}
