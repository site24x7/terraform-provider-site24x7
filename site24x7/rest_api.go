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
		Type:     schema.TypeString,
		Required: true,
	},
	"website": {
		Type:     schema.TypeString,
		Required: true,
	},
	"check_frequency": {
		Type:     schema.TypeInt,
		Optional: true,
		Default:  1,
	},
	"timeout": {
		Type:     schema.TypeInt,
		Optional: true,
		Default:  10,
	},
	"location_profile_id": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
	"location_profile_name": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"notification_profile_id": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
	"threshold_profile_id": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
	"monitor_groups": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional: true,
	},
	"user_group_ids": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional: true,
		Computed: true,
	},
	"http_method": {
		Type:     schema.TypeString,
		Optional: true,
		Default:  "G",
	},
	"http_protocol": {
		Type:     schema.TypeString,
		Optional: true,
		Default:  "H1.1",
	},
	"ssl_protocol": {
		Type:     schema.TypeString,
		Optional: true,
		Default:  "Auto",
		// DiffSuppressFunc: func(k, sslProtocolInState, sslProtocolInConf string, d *schema.ResourceData) bool {
		// 	// GET API response doesn't contain 'ssl_protocol' attribute. Temporary fix to suppress the
		// 	// misleading plan.
		// 	if sslProtocolInState == "" {
		// 		return true
		// 	} else {
		// 		return false
		// 	}
		// },
	},
	"use_ipv6": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"request_content_type": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"response_content_type": {
		Type:     schema.TypeString,
		Optional: true,
		Default:  "T",
	},
	"request_param": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"auth_user": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"auth_pass": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"oauth2_provider": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"client_certificate_password": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"jwt_id": {
		Type:     schema.TypeString,
		Optional: true,
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
	},
	"custom_headers": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	},
	"match_case": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"json_schema_check": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"use_name_server": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"use_alpn": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	},
	"user_agent": {
		Type:     schema.TypeString,
		Optional: true,
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

	var actionRefs []api.ActionRef
	if actionData, ok := d.GetOk("actions"); ok {
		actionMap := actionData.(map[string]interface{})

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

		AuthUser:        d.Get("auth_user").(string),
		AuthPass:        d.Get("auth_pass").(string),
		MatchCase:       d.Get("match_case").(bool),
		JSONSchemaCheck: d.Get("json_schema_check").(bool),
		UseNameServer:   d.Get("use_name_server").(bool),
		UserAgent:       d.Get("user_agent").(string),
		CustomHeaders:   customHeaders,

		LocationProfileID:     d.Get("location_profile_id").(string),
		NotificationProfileID: d.Get("notification_profile_id").(string),
		ThresholdProfileID:    d.Get("threshold_profile_id").(string),
		MonitorGroups:         monitorGroups,
		UserGroupIDs:          userGroupIDs,
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

	if restApiMonitor.NotificationProfileID == "" {
		profile, err := DefaultNotificationProfile(client)
		if err != nil {
			return nil, err
		}
		restApiMonitor.NotificationProfileID = profile.ProfileID
		d.Set("notification_profile_id", profile.ProfileID)
	}

	if restApiMonitor.ThresholdProfileID == "" {
		profile, err := DefaultThresholdProfile(client, api.RESTAPI)
		if err != nil {
			return nil, err
		}
		restApiMonitor.ThresholdProfileID = profile.ProfileID
		d.Set("threshold_profile_id", profile)
	}

	if len(restApiMonitor.UserGroupIDs) == 0 {
		userGroup, err := DefaultUserGroup(client)
		if err != nil {
			return nil, err
		}
		restApiMonitor.UserGroupIDs = []string{userGroup.UserGroupID}
		d.Set("user_group_ids", []string{userGroup.UserGroupID})
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

	customHeaders := make(map[string]interface{})
	for _, h := range monitor.CustomHeaders {
		if h.Name == "" {
			continue
		}
		customHeaders[h.Name] = h.Value
	}

	d.Set("custom_headers", customHeaders)

	actions := make(map[string]interface{})
	for _, action := range monitor.ActionIDs {
		actions[fmt.Sprintf("%d", action.AlertType)] = action.ActionID
	}

	d.Set("actions", actions)
}
