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

// Sample RESTAPI TRANSACTION POST JSON

//{
//"check_frequency": "5",
//"display_name": "foo",
//"location_profile_id": "111111000000025013",
//"notification_profile_id": "111111000000029001",
//"steps": [{
//"display_name": "Step 1",
//"monitor_id": "",
//"step_details": [{
//"auth_pass": "password",
//"auth_user": "username",
//"client_certificate_password": "pass",
//"display_name": "Step 1",
//"http_method": "G",
//"http_protocol": "H1.1",
//"match_case": true,
//"oauth2_provider": "provider",
//"request_content_type": "JSON",
//"request_param": "req_param",
//"response_headers_check": {
//"severity": 2,
//"value": [{
//"name": "Accept-Encoding",
//"value": "gzip"
//}, {
//"name": "Cache-Control",
//"value": "nocache"
//}]
//},
//"unmatching_keyword": {
//"value": "aaa",
//"severity": 2
//},
//"matching_keyword": {
//"value": "bbb",
//"severity": 2
//},
//"match_regex": {
//"severity": 0,
//"value": "*.a.*"
//},
//"response_type": "T",
//"ssl_protocol": "Auto",
//"step_url": "www.test.tld",
//"timeout": 10,
//"use_alpn": false,
//"use_name_server": true,
//"user_agent": "firefox"
//}]
//},{
//"display_name": "Step 2",
//"monitor_id": "",
//"step_details": [{
//"auth_pass": "password",
//"auth_user": "username",
//"client_certificate_password": "pass",
//"display_name": "Step 2",
//"http_method": "G",
//"http_protocol": "H1.1",
//"match_case": true,
//"oauth2_provider": "provider",
//"request_content_type": "JSON",
//"request_param": "req_param",
//"unmatching_keyword": {
//"value": "aaa",
//"severity": 2
//},
//"matching_keyword": {
//"value": "bbb",
//"severity": 2
//},
//"match_regex": {
//"severity": 0,
//"value": "*.a.*"
//},
//"response_headers_check": {
//"severity": 2,
//"value": [{
//"name": "Accept-Encoding",
//"value": "gzip"
//}, {
//"name": "Cache-Control",
//"value": "nocache"
//}]
//},
//"response_type": "T",
//"ssl_protocol": "Auto",
//"step_url": "www.test.tld",
//"timeout": 10,
//"use_alpn": false,
//"use_name_server": true,
//"user_agent": "firefox"
//}]
//}
//],
//"tag_ids": [],
//"threshold_profile_id": "111111000021519001",
//"type": "RESTAPISEQ",
//"user_group_ids": ["111111000000025005"]
//}

var RestApiTransactionMonitorSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display Name for the monitor.",
	},
	"check_frequency": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "5",
		Description: "Interval at which your RESRAPI has to be monitored. Default value is 5 minute.",
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
	"steps": {
		Type:     schema.TypeSet,
		Optional: true,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"display_name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Display Name for the monitor.",
				},
				"step_details": {
					Type:     schema.TypeSet,
					Computed: true,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"step_url": {
								Type:        schema.TypeString,
								Required:    true,
								Description: "Rest API Url to monitors",
							},
							"timeout": {
								Type:        schema.TypeString,
								Optional:    true,
								Default:     "10",
								Description: "Timeout for connecting to website. Default value is 10. Range 1 - 45.",
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
								Description: "HTTP Method to be used for accessing the website. Default value is 'G'. 'G' denotes GET, 'P' denotes POST, 'U' denotes PUT and 'D' denotes DELETE. HEAD is not supported.",
							},
							"use_ipv6": {
								Type:        schema.TypeBool,
								Optional:    true,
								Description: "Monitoring is performed over IPv6 from supported locations. IPv6 locations do not fall back to IPv4 on failure.",
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
							"response_variables": {
								Type:     schema.TypeMap,
								Optional: true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"response_type": {
											Type:         schema.TypeString,
											Required:     true,
											ValidateFunc: validation.IntInSlice([]int{0, 2}), // Trouble or Down
										},
										"variables": {
											Type:     schema.TypeSet,
											Optional: true,
											Computed: true,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"name": {
														Type:        schema.TypeString,
														Required:    true,
														Description: "Store the name of parameter forwarding variable path",
													},
													"value": {
														Type:        schema.TypeString,
														Required:    true,
														Description: "Storing the parameter forwarding variable path",
													},
												},
											},
											Description: "List of Parameter forwarding variables",
										},
									},
								},
								Description: "Response Format to send response type and parameter forwarding variable",
							},
							"dynamic_header_params": {
								Type:     schema.TypeSet,
								Optional: true,
								Computed: true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"name": {
											Type:        schema.TypeString,
											Required:    true,
											Description: "Store the name of parameter forwarding variable path",
										},
										"value": {
											Type:        schema.TypeString,
											Required:    true,
											Description: "Storing the parameter forwarding variable path",
										},
									},
								},
								Description: "List of Response Header/Cookies Format to send the parameter forwarding variable",
							},
						},
					},
					Description: "API request details related to this step",
				},
			},
		},
		Description: "List of Monitors steps",
	},
}

func ResourceSite24x7RestApiTransactionMonitor() *schema.Resource {
	return &schema.Resource{
		Create: restApiTransactionMonitorCreate,
		Read:   restApiTransactionMonitorRead,
		Update: restApiTransactionMonitorUpdate,
		Delete: restApiTransactionMonitorDelete,
		Exists: restApiTransactionMonitorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: RestApiTransactionMonitorSchema,
	}
}

func restApiTransactionMonitorCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(site24x7.Client)

	restApiTransactionMonitor, err := resourceDataToRestApiTransactionMonitor(d, client)
	if err != nil {
		return err
	}

	restApiTransactionMonitor, err = client.RestApiTransactionMonitors().Create(restApiTransactionMonitor)
	if err != nil {
		return err
	}

	d.SetId(restApiTransactionMonitor.MonitorID)

	// return restApiMonitorRead(d, meta)
	return nil
}

func restApiTransactionMonitorRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(site24x7.Client)

	restApiTransactionMonitors, err := client.RestApiTransactionMonitors().Get(d.Id())

	restApiTransactionMonitorsSteps, steperr := client.RestApiTransactionMonitors().GetSteps(d.Id())

	if err != nil {
		return err
	}
	if steperr != nil {
		return steperr
	}

	updateRestApiTransactionMonitorResourceData(d, restApiTransactionMonitors, restApiTransactionMonitorsSteps)

	return nil
}

func restApiTransactionMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	restApiTransactionMonitors, err := resourceDataToRestApiTransactionMonitor(d, client)

	if err != nil {
		return err
	}

	restApiTransactionMonitors, err = client.RestApiTransactionMonitors().Update(restApiTransactionMonitors)
	if err != nil {
		return err
	}

	d.SetId(restApiTransactionMonitors.MonitorID)

	// return restApiMonitorRead(d, meta)
	return nil
}

func restApiTransactionMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.RestApiTransactionMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func restApiTransactionMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.RestApiTransactionMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToRestApiTransactionMonitor(d *schema.ResourceData, client site24x7.Client) (*api.RestApiTransactionMonitor, error) {

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

	// Steps Configuration

	Steps := d.Get("steps").(*schema.Set)
	StepsItems := make([]api.Steps, Steps.Len())
	for k, v := range Steps.List() {
		StepsDetails := v.(map[string]interface{})["step_details"].(*schema.Set)
		StepsDetailsItem := make([]api.StepDetails, StepsDetails.Len())

		for i, j := range StepsDetails.List() {
			// Request Headers
			requestHeaderMap := j.(map[string]interface{})["request_headers"].(map[string]interface{})
			requestHeaderKeys := make([]string, 0, len(requestHeaderMap))
			for k := range requestHeaderMap {
				requestHeaderKeys = append(requestHeaderKeys, k)
			}
			sort.Strings(requestHeaderKeys)
			requestHeaders := make([]api.Header, len(requestHeaderKeys))
			for i, k := range requestHeaderKeys {
				requestHeaders[i] = api.Header{Name: k, Value: requestHeaderMap[k].(string)}
			}

			// Dynamic Header Params changes

			var httpResponseVariable api.HTTPResponseVariable
			responseVariableMap := j.(map[string]interface{})["response_variable"].(map[string]interface{})
			if len(responseVariableMap) > 0 {
				responseVariableKeys := make([]string, 0, len(responseVariableMap))
				for k := range responseVariableMap {
					responseVariableKeys = append(responseVariableKeys, k)
				}
				sort.Strings(responseVariableKeys)
				responseVariables := make([]api.Header, len(responseVariableKeys))

				for i, k := range responseVariableKeys {
					responseVariables[i] = api.Header{Name: k, Value: responseVariableMap[k].(string)}
				}
				httpResponseVariable.ResponseType = api.ResponseType(j.(map[string]interface{})["response_type"].(string))
				httpResponseVariable.Variables = responseVariables
			}

			var dynamicHeaderParams api.HTTPDynamicHeaderParams
			dynamicHeaderParamsMap := j.(map[string]interface{})["dynamic_header_params"].(map[string]interface{})
			if len(dynamicHeaderParamsMap) > 0 {
				dynamicHeaderParamsKeys := make([]string, 0, len(dynamicHeaderParamsMap))
				for k := range dynamicHeaderParamsMap {
					dynamicHeaderParamsKeys = append(dynamicHeaderParamsKeys, k)
				}
				sort.Strings(dynamicHeaderParamsKeys)
				dynamicReponseVariables := make([]api.Header, len(dynamicHeaderParamsKeys))

				for i, k := range dynamicHeaderParamsKeys {
					dynamicReponseVariables[i] = api.Header{Name: k, Value: dynamicHeaderParamsMap[k].(string)}
				}
				dynamicHeaderParams.Variables = dynamicReponseVariables
			}

			// HTTP Response Headers
			var httpResponseHeader api.HTTPResponseHeader
			responseHeaderMap := j.(map[string]interface{})["response_headers"].(map[string]interface{})
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
				httpResponseHeader.Severity = api.Status(j.(map[string]interface{})["response_headers_severity"].(int))
				httpResponseHeader.Value = responseHeaders
			}

			var MatchRegex map[string]interface{}
			if matchingRegex, ok := j.(map[string]interface{})["match_regex"]; ok {
				MatchRegex = matchingRegex.(map[string]interface{})
			}

			var MatchingKeyword map[string]interface{}
			if matchingKeyword, ok := j.(map[string]interface{})["matching_keyword"]; ok {
				MatchingKeyword = matchingKeyword.(map[string]interface{})
			}

			var UnmatchingKeyword map[string]interface{}
			if unmatchingKeyword, ok := j.(map[string]interface{})["unmatching_keyword"]; ok {
				UnmatchingKeyword = unmatchingKeyword.(map[string]interface{})
			}

			var MatchJSON map[string]interface{}
			if matchJSONPath, ok := d.GetOk("match_json_path"); ok {
				var jsonPathList []map[string]interface{}
				for _, jsonPath := range matchJSONPath.([]interface{}) {
					matchPathMap := make(map[string]interface{})
					matchPathMap["name"] = jsonPath.(string)
					jsonPathList = append(jsonPathList, matchPathMap)
				}
				matchJSONData := make(map[string]interface{})
				matchJSONData["jsonpath"] = jsonPathList
				matchJSONData["severity"] = j.(map[string]interface{})["match_json_path_severity"].(int)
				MatchJSON = matchJSONData
			}

			var JSONSchema map[string]interface{}
			if jsonSchema, ok := d.GetOk("json_schema"); ok {
				jsonSchemaData := make(map[string]interface{})
				jsonSchemaData["severity"] = d.Get("json_schema_severity").(int)
				jsonSchemaData["schema_value"] = jsonSchema.(string)
				JSONSchema = jsonSchemaData
			}

			var GraphQL map[string]interface{}
			if graphqlQuery, ok := d.GetOk("graphql_query"); ok {
				graphqlMap := make(map[string]interface{})
				graphqlMap["query"] = graphqlQuery.(string)
				graphqlMap["variables"] = d.Get("graphql_variables").(string)
				GraphQL = graphqlMap
			}

			StepsDetailsItem[i] = api.StepDetails{
				StepUrl:                   j.(map[string]interface{})["step_url"].(string),
				DisplayName:               v.(map[string]interface{})["display_name"].(string),
				HTTPMethod:                j.(map[string]interface{})["http_method"].(string),
				RequestContentType:        j.(map[string]interface{})["request_content_type"].(string),
				RequestBody:               j.(map[string]interface{})["request_body"].(string),
				RequestHeaders:            requestHeaders,
				GraphQL:                   GraphQL,
				UserAgent:                 j.(map[string]interface{})["user_agent"].(string),
				AuthMethod:                j.(map[string]interface{})["auth_method"].(string),
				AuthUser:                  j.(map[string]interface{})["auth_user"].(string),
				AuthPass:                  j.(map[string]interface{})["auth_pass"].(string),
				OAuth2Provider:            j.(map[string]interface{})["oauth2_provider"].(string),
				ClientCertificatePassword: j.(map[string]interface{})["client_certificate_password"].(string),
				JwtID:                     j.(map[string]interface{})["jwt_id"].(string),
				UseNameServer:             j.(map[string]interface{})["use_name_server"].(bool),
				HTTPProtocol:              j.(map[string]interface{})["http_protocol"].(string),
				SSLProtocol:               j.(map[string]interface{})["ssl_protocol"].(string),
				UpStatusCodes:             j.(map[string]interface{})["up_status_codes"].(string),
				UseAlpn:                   j.(map[string]interface{})["use_alpn"].(bool),
				ResponseContentType:       j.(map[string]interface{})["response_content_type"].(string),
				MatchJSON:                 MatchJSON,
				JSONSchema:                JSONSchema,
				JSONSchemaCheck:           j.(map[string]interface{})["json_schema_check"].(bool),
				MatchingKeyword:           MatchingKeyword,
				UnmatchingKeyword:         UnmatchingKeyword,
				MatchCase:                 j.(map[string]interface{})["match_case"].(bool),
				MatchRegex:                MatchRegex,
				ResponseHeaders:           httpResponseHeader,
				ResponseVariable:          httpResponseVariable,
				DynamicHeaderParams:       dynamicHeaderParams,
			}
		}

		StepsItems[k] = api.Steps{
			DisplayName:  v.(map[string]interface{})["display_name"].(string),
			StepsDetails: StepsDetailsItem,
			MonitorID:    d.Id(),
		}
	}

	restApiTransactionMonitor := &api.RestApiTransactionMonitor{
		MonitorID:             d.Id(),
		DisplayName:           d.Get("display_name").(string),
		Type:                  string(api.RESTAPISEQ),
		CheckFrequency:        d.Get("check_frequency").(string),
		LocationProfileID:     d.Get("location_profile_id").(string),
		NotificationProfileID: d.Get("notification_profile_id").(string),
		ThresholdProfileID:    d.Get("threshold_profile_id").(string),
		MonitorGroups:         monitorGroups,
		DependencyResourceIDs: dependencyResourceIDs,
		UserGroupIDs:          userGroupIDs,
		TagIDs:                tagIDs,
		ThirdPartyServiceIDs:  thirdPartyServiceIDs,
		ActionIDs:             actionRefs,
		Steps:                 StepsItems,
	}

	// Location Profile
	_, locationProfileErr := site24x7.SetLocationProfile(client, d, restApiTransactionMonitor)
	if locationProfileErr != nil {
		return nil, locationProfileErr
	}

	// Notification Profile
	_, notificationProfileErr := site24x7.SetNotificationProfile(client, d, restApiTransactionMonitor)
	if notificationProfileErr != nil {
		return nil, notificationProfileErr
	}

	// User Alert Groups
	_, userAlertGroupErr := site24x7.SetUserGroup(client, d, restApiTransactionMonitor)
	if userAlertGroupErr != nil {
		return nil, userAlertGroupErr
	}

	// Tags
	_, tagsErr := site24x7.SetTags(client, d, restApiTransactionMonitor)
	if tagsErr != nil {
		return nil, tagsErr
	}

	if restApiTransactionMonitor.ThresholdProfileID == "" {
		profile, err := site24x7.DefaultThresholdProfile(client, api.RESTAPISEQ)
		if err != nil {
			return nil, err
		}
		restApiTransactionMonitor.ThresholdProfileID = profile.ProfileID
		d.Set("threshold_profile_id", profile)
	}

	return restApiTransactionMonitor, nil
}

func updateRestApiTransactionMonitorResourceData(d *schema.ResourceData, monitor *api.RestApiTransactionMonitor, steps *[]api.Steps) {

	d.Set("display_name", monitor.DisplayName)
	d.Set("type", monitor.Type)
	d.Set("check_frequency", monitor.CheckFrequency)
	d.Set("steps", *steps)
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
