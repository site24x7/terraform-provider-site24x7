package monitors

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var SOAPMonitorSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Name for the monitor",
	},
	"website": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Registered domain name.",
	},
	"request_param": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Request params to given for the soap monitor",
	},
	"soap_attributes_severity": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Registered domain name.",
	},
	"soap_attributes": {
		Type:        schema.TypeMap,
		Optional:    true,
		Description: "A Map of respone header name and value.",
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
	"timeout": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     10,
		Description: "Timeout for connecting to website. Default value is 10. Range 1 - 45",
	},
	"request_content_type": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "",
	},
	"http_method": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "",
	},
	"use_name_server": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "",
	},
	"up_status_codes": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "",
		Description: "Provide a comma-separated list of HTTP status codes that indicate a successful response. You can specify individual status codes, as well as ranges separated with a colon.",
	},
	"http_protocol": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "",
	},
	"use_ipv6": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Monitoring is performed over IPv6 from supported locations. IPv6 locations do not fall back to IPv4 on failure.",
	},
	"response_type": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "",
	},
	"check_frequency": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "5",
		Description: "Interval at which your website has to be monitored. Default value is 5 minute.",
	},
	"perform_automation": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "To perform automation or not",
	},
	"use_alpn": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "",
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
	"threshold_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Threshold profile to be associated with the monitor.",
	},
	"ssl_protocol": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "SSL Protocol to be associated with the monitor.",
	},
	"credential_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Credential Profile to associate.",
	},
	// "client_certificate_password": {
	// 	Type:        schema.TypeString,
	// 	Optional:    true,
	// 	Description: "Password of the client certificate.",
	// 	DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
	// 		// Suppress diff - Password in API response is encrypted.
	// 		return true
	// 	},
	// },
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
	"dependency_resource_ids": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of dependent resource IDs. Suppress alert when dependent monitor(s) is down.",
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
	"actions": {
		Type:        schema.TypeMap,
		Optional:    true,
		Elem:        schema.TypeString,
		Description: "Action to be performed on monitor status changes",
	},
	"third_party_services": {
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

func ResourceSite24x7SOAPMonitor() *schema.Resource {
	return &schema.Resource{
		Create: soapMonitorCreate,
		Read:   soapMonitorRead,
		Update: soapMonitorUpdate,
		Delete: soapMonitorDelete,
		Exists: soapMonitorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: SOAPMonitorSchema,
	}
}

func soapMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)
	log.Println("***********************************Inside soap monitor creation*******************************************")
	soapMonitor, err := resourceDataToSOAPMonitor(d, client)
	if err != nil {
		return err
	}
	soapMonitor, err = client.SOAPMonitors().Create(soapMonitor)
	if err != nil {
		return err
	}

	d.SetId(soapMonitor.MonitorID)

	return nil
}

func soapMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	soapMonitor, err := client.SOAPMonitors().Get(d.Id())

	if err != nil {
		return err
	}

	updateSOAPMonitorResourceData(d, soapMonitor)

	return nil
}

func soapMonitorUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(site24x7.Client)

	soapMonitor, err := resourceDataToSOAPMonitor(d, client)

	if err != nil {
		return err
	}
	soapMonitor, err = client.SOAPMonitors().Update(soapMonitor)
	if err != nil {
		return err
	}

	d.SetId(soapMonitor.MonitorID)

	return nil
}

func soapMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.PINGMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func soapMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.SOAPMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToSOAPMonitor(d *schema.ResourceData, client site24x7.Client) (*api.SOAPMonitor, error) {
	log.Println("***********************RESOURCEDATATOSOAPMONITOR************************************")
	var monitorGroups []string
	for _, group := range d.Get("monitor_groups").([]interface{}) {
		if group != nil {
			monitorGroups = append(monitorGroups, group.(string))
		}
	}
	sort.Strings(monitorGroups)
	
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
	for _, id := range d.Get("third_party_services").([]interface{}) {
		if id != nil {
			thirdPartyServiceIDs = append(thirdPartyServiceIDs, id.(string))
		}
	}
	dependencyIDs := d.Get("dependency_resource_ids").(*schema.Set).List()
	dependencyResourceIDs := make([]string, 0, len(dependencyIDs))
	for _, dependencyResourceID := range dependencyIDs {
		if dependencyResourceID != nil {
			dependencyResourceIDs = append(dependencyResourceIDs, dependencyResourceID.(string))
		}
	}

	actionMap := d.Get("actions").(map[string]interface{})
	var keys = make([]string, 0, len(actionMap))
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

	soapMonitor := &api.SOAPMonitor{
		MonitorID:              d.Id(),
		DisplayName:            d.Get("display_name").(string),
		Type:                   string(api.SOAP),
		Website:                d.Get("website").(string),
		RequestParam:           d.Get("request_param").(string),
		SOAPAttributesSeverity: d.Get("soap_attributes_severity").(int),
		Timeout:                d.Get("timeout").(int),
		RequestContentType:     d.Get("request_content_type").(string),
		HTTPMethod:             d.Get("http_method").(string),
		UseNameServer:          d.Get("use_name_server").(bool),
		HTTPProtocol:           d.Get("http_protocol").(string),
		UseIPV6:                d.Get("use_ipv6").(bool),
		ResponseType:           d.Get("response_type").(string),
		CheckFrequency:         d.Get("check_frequency").(string),
		SSLProtocol:            d.Get("ssl_protocol").(string),
		UseAlpn:                d.Get("use_alpn").(bool),
		LocationProfileID:      d.Get("location_profile_id").(string),
		NotificationProfileID:  d.Get("notification_profile_id").(string),
		PerformAutomation:      d.Get("perform_automation").(bool),
		OnCallScheduleID:       d.Get("on_call_schedule_id").(string),
		ThresholdProfileID:     d.Get("threshold_profile_id").(string),
		MonitorGroups:          monitorGroups,
		UserGroupIDs:           userGroupIDs,
		TagIDs:                 tagIDs,
		DependencyResourceIDs:  dependencyResourceIDs,
		ThirdPartyServiceIDs:   thirdPartyServiceIDs,
		ActionIDs:              actionRefs,
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
	soapMonitor.CredentialProfileID = d.Get("credential_profile_id").(string)
	//soapMonitor.ClientCertificatePassword = d.Get("client_certificate_password").(string)
	soapMonitor.UpStatusCodes = d.Get("up_status_codes").(string)
	soapMonitor.ResponseHeaders = httpResponseHeader
	soapAttributeMap := d.Get("soap_attributes").(map[string]interface{})
	attributekeys := make([]string, 0, len(soapAttributeMap))
	for k := range soapAttributeMap {
		attributekeys = append(attributekeys, k)
	}
	sort.Strings(attributekeys)
	soapAttributes := make([]api.Header, len(attributekeys))
	for i, k := range attributekeys {
		soapAttributes[i] = api.Header{Name: k, Value: soapAttributeMap[k].(string)}
	}
	soapMonitor.SOAPAttributes = soapAttributes
	_, locationProfileErr := site24x7.SetLocationProfile(client, d, soapMonitor)
	if locationProfileErr != nil {
		return nil, locationProfileErr
	}

	_, notificationProfileErr := site24x7.SetNotificationProfile(client, d, soapMonitor)
	if notificationProfileErr != nil {
		return nil, notificationProfileErr
	}

	// User Alert Groups
	_, userAlertGroupErr := site24x7.SetUserGroup(client, d, soapMonitor)
	if userAlertGroupErr != nil {
		return nil, userAlertGroupErr
	}

	// Tags
	_, tagsErr := site24x7.SetTags(client, d, soapMonitor)
	if tagsErr != nil {
		return nil, tagsErr
	}

	//Threshold
	if soapMonitor.ThresholdProfileID == "" {
		profile, err := site24x7.DefaultThresholdProfile(client, api.SOAP)
		if err != nil {
			return nil, err
		}
		soapMonitor.ThresholdProfileID = profile.ProfileID
	}
	return soapMonitor, nil
}

func updateSOAPMonitorResourceData(d *schema.ResourceData, monitor *api.SOAPMonitor) {

	d.Set("display_name", monitor.DisplayName)
	d.Set("website", monitor.Website)
	d.Set("request_params", monitor.RequestParam)
	d.Set("type", monitor.Type)
	d.Set("host_name", monitor.DisplayName)
	d.Set("timeout", monitor.Timeout)
	d.Set("use_ipv6", monitor.UseIPV6)
	d.Set("request_content_type", monitor.RequestContentType)
	d.Set("http_method", monitor.HTTPMethod)
	d.Set("use_name_server", monitor.UseNameServer)
	d.Set("http_protocol", monitor.HTTPProtocol)
	d.Set("response_type", monitor.ResponseType)
	d.Set("check_frequency", monitor.CheckFrequency)
	d.Set("perform_automation", monitor.PerformAutomation)
	d.Set("use_alpn", monitor.UseAlpn)
	d.Set("location_profile_id", monitor.LocationProfileID)
	d.Set("up_status_codes", monitor.UpStatusCodes)
	d.Set("notification_profile_id", monitor.NotificationProfileID)
	d.Set("threshold_profile_id", monitor.ThresholdProfileID)
	d.Set("ssl_protocol", monitor.SSLProtocol)
	d.Set("user_group_ids", monitor.UserGroupIDs)
	d.Set("dependency_resource_ids", monitor.DependencyResourceIDs)
	d.Set("on_call_schedule_id", monitor.OnCallScheduleID)
	d.Set("monitor_groups", monitor.MonitorGroups)
	d.Set("third_party_services", monitor.ThirdPartyServiceIDs)
	d.Set("tag_ids", monitor.TagIDs)
	d.Set("soap_attributes_severity", monitor.SOAPAttributesSeverity)
	d.Set("credential_profile_id", monitor.CredentialProfileID)
	//d.Set("client_certificate_password", monitor.ClientCertificatePassword)
	soapAttributes := make(map[string]interface{})
	for _, h := range monitor.SOAPAttributes {
		if h.Name == "" {
			continue
		}
		soapAttributes[h.Name] = h.Value
	}
	d.Set("soap_attributes", soapAttributes)

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
