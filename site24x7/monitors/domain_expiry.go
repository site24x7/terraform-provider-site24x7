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

var DomainExpiryMonitorSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Name for the monitor",
	},
	"type": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "DOMAINEXPIRY",
	},
	"host_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Registered domain name.",
	},
	"domain_name": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "whois.iana.org",
		Description: "Specify the name of the Whois server from where you wish to query the domain data.",
	},
	"port": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     443,
		Description: "Who is Server Port",
	},
	"timeout": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     10,
		Description: "Timeout for connecting to website. Default value is 10. Range 1 - 45",
	},
	"expire_days": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     30,
		Description: "Day threshold for domain expiry notification.Range 1 - 999",
	},
	"use_ipv6": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Monitoring is performed over IPv6 from supported locations. IPv6 locations do not fall back to IPv4 on failure.",
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
	"ignore_registry_date": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Ignores the registry expiry date and prefer registrar expiry date when notifying for domain expiry",
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

func ResourceSite24x7DomainExpiryMonitor() *schema.Resource {
	return &schema.Resource{
		Create: domainExpiryMonitorCreate,
		Read:   domainExpiryMonitorRead,
		Update: domainExpiryMonitorUpdate,
		Delete: domainExpiryMonitorDelete,
		Exists: domainExpiryMonitorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: DomainExpiryMonitorSchema,
	}
}

func domainExpiryMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	domainExpiryMonitor, err := resourceDataToDomainExpiryMonitor(d, client)
	if err != nil {
		return err
	}

	domainExpiryMonitor, err = client.DomainExpiryMonitors().Create(domainExpiryMonitor)
	if err != nil {
		return err
	}

	d.SetId(domainExpiryMonitor.MonitorID)

	// return domainExpiryMonitorRead(d, meta)
	return nil
}

func domainExpiryMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	domainExpiryMonitor, err := client.DomainExpiryMonitors().Get(d.Id())

	if err != nil {
		return err
	}

	updateDomainExpiryMonitorResourceData(d, domainExpiryMonitor)

	return nil
}

func domainExpiryMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)
	log.Println("ignore_registry_date inside create : ", d.Get("ignore_registry_date"))
	domainExpiryMonitor, err := resourceDataToDomainExpiryMonitor(d, client)

	if err != nil {
		return err
	}

	domainExpiryMonitor, err = client.DomainExpiryMonitors().Update(domainExpiryMonitor)
	if err != nil {
		return err
	}

	d.SetId(domainExpiryMonitor.MonitorID)
	// return domainExpiryMonitorRead(d, meta)
	return nil
}

func domainExpiryMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.DomainExpiryMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func domainExpiryMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.DomainExpiryMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToDomainExpiryMonitor(d *schema.ResourceData, client site24x7.Client) (*api.DomainExpiryMonitor, error) {

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

	var thirdPartyServiceIDs []string
	for _, id := range d.Get("third_party_service_ids").([]interface{}) {
		if id != nil {
			thirdPartyServiceIDs = append(thirdPartyServiceIDs, id.(string))
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

	domainExpiryMonitor := &api.DomainExpiryMonitor{
		MonitorID:             d.Id(),
		DisplayName:           d.Get("display_name").(string),
		Type:                  string(api.DOMAINEXPIRY),
		HostName:              d.Get("host_name").(string),
		WhoIsServer:           d.Get("domain_name").(string),
		Port:                  d.Get("port"),
		Timeout:               d.Get("timeout").(int),
		ExpireDays:            d.Get("expire_days").(int),
		UseIPV6:               d.Get("use_ipv6").(bool),
		LocationProfileID:     d.Get("location_profile_id").(string),
		NotificationProfileID: d.Get("notification_profile_id").(string),
		OnCallScheduleID:      d.Get("on_call_schedule_id").(string),
		MatchCase:             d.Get("match_case").(bool),
		IgnoreRegistryDate:    d.Get("ignore_registry_date").(bool),
		MonitorGroups:         monitorGroups,
		UserGroupIDs:          userGroupIDs,
		TagIDs:                tagIDs,
		ThirdPartyServiceIDs:  thirdPartyServiceIDs,
		ActionIDs:             actionRefs,
	}
	if matchingRegex, ok := d.GetOk("match_regex"); ok {
		domainExpiryMonitor.MatchRegex = matchingRegex.(map[string]interface{})
	}

	if matchingKeyword, ok := d.GetOk("matching_keyword"); ok {
		domainExpiryMonitor.MatchingKeyword = matchingKeyword.(map[string]interface{})
	}

	if unmatchingKeyword, ok := d.GetOk("unmatching_keyword"); ok {
		domainExpiryMonitor.UnmatchingKeyword = unmatchingKeyword.(map[string]interface{})
	}

	_, locationProfileErr := site24x7.SetLocationProfile(client, d, domainExpiryMonitor)
	if locationProfileErr != nil {
		return nil, locationProfileErr
	}

	// Notification Profile
	_, notificationProfileErr := site24x7.SetNotificationProfile(client, d, domainExpiryMonitor)
	if notificationProfileErr != nil {
		return nil, notificationProfileErr
	}

	// User Alert Groups
	_, userAlertGroupErr := site24x7.SetUserGroup(client, d, domainExpiryMonitor)
	if userAlertGroupErr != nil {
		return nil, userAlertGroupErr
	}

	// Tags
	_, tagsErr := site24x7.SetTags(client, d, domainExpiryMonitor)
	if tagsErr != nil {
		return nil, tagsErr
	}
	return domainExpiryMonitor, nil
}

func updateDomainExpiryMonitorResourceData(d *schema.ResourceData, monitor *api.DomainExpiryMonitor) {
	d.Set("display_name", monitor.DisplayName)
	d.Set("type", monitor.Type)
	d.Set("host_name", monitor.HostName)
	d.Set("domain_name", monitor.WhoIsServer)
	d.Set("port", monitor.Port)
	d.Set("timeout", monitor.Timeout)
	d.Set("expire_days", monitor.ExpireDays)
	d.Set("use_ipv6", monitor.UseIPV6)
	d.Set("location_profile_id", monitor.LocationProfileID)
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

	d.Set("notification_profile_id", monitor.NotificationProfileID)
	d.Set("on_call_schedule_id", monitor.OnCallScheduleID)
	d.Set("ignore_registry_date", monitor.IgnoreRegistryDate)
	d.Set("monitor_groups", monitor.MonitorGroups)
	d.Set("user_group_ids", monitor.UserGroupIDs)
	d.Set("tag_ids", monitor.TagIDs)
	d.Set("third_party_service_ids", monitor.ThirdPartyServiceIDs)
	actions := make(map[string]interface{})
	for _, action := range monitor.ActionIDs {
		actions[fmt.Sprintf("%d", action.AlertType)] = action.ActionID
	}

	d.Set("actions", actions)
}
