package monitors

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	// "github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var dnsServerMonitorSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display Name for the monitor.",
	},
	"type": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "DNS",
	},
	"dns_host": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "DNS Name Server to be monitored",
	},
	"dns_port": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "DNS Name Server to be monitored",
	},
	"use_ipv6": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Monitoring is performed over IPv6 from supported locations. IPv6 locations do not fall back to IPv4 on failure.",
	},
	"domain_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Domain name to be resolved.",
	},
	"check_frequency": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Check interval for monitoring. See https://www.site24x7.com/help/api/#check_interval.",
	},
	"timeout": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "Timeout for connecting to website. Default value is 10. Range 1 - 45.",
	},
	"deep_discovery": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Enable this attribute to auto discover and set up monitoring for all the related resources for the domain_name.",
	},

	// Configuration Profiles
	"location_profile_id": {
		Type:     schema.TypeString,
		Optional: true,
		// Computed:    true,
		Description: "Location profile to be associated with the monitor.",
	},
	"location_profile_name": {
		Type:        schema.TypeString,
		Required:    true,
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
		Required:    true,
		Description: "Name of the notification profile to be associated with the monitor.",
	},
	"threshold_profile_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Threshold profile to be associated with the monitor.",
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
		Required:    true,
		Description: "Name of the user groups to be associated with the monitor.",
	},
	"monitor_groups": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of monitor groups to which the monitor has to be associated.",
	},
	"lookup_type": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "DNS Server Lookup Type Constants. See https://www.site24x7.com/help/api/#dns_lookup_type",
	},
	"dnssec": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Pass dnssec parameter to enable Site24x7 to validate DNS responses.",
		Default:     false,
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
		Description: "Action to be performed on monitor status changes.",
	},
	"search_config": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"lookup_type": {
					Type:     schema.TypeString,
					Required: true,
				},
				"addr": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"ttlo": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"ttl": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"target": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"priority": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"port": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"wt": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"rcvd": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"pns": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"admin": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"serial": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"rff": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"rtf": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"expt": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"mttl": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"flg": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"prtcl": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"kalg": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"kid": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"key": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"tag": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"cert_auth": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"halg": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"hash": {
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
		Description: "Value to be checked against resolved values. Choose a JSON Format based on your configured lookup type. See https://www.site24x7.com/help/api/#constants for details",
	},
	"third_party_service_ids": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of Third Party Service IDs to be associated to the monitor.",
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
	"on_call_schedule_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "if user_group_ids is not choosen,	On-Call Schedule of your choice.",
	},
}

func ResourceSite24x7DNSServerMonitor() *schema.Resource {
	return &schema.Resource{
		Create: dnsServerMonitorCreate,
		Read:   dnsServerMonitorRead,
		Update: dnsServerMonitorUpdate,
		Delete: dnsServerMonitorDelete,
		Exists: dnsServerMonitorExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: dnsServerMonitorSchema,
	}
}

func dnsServerMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	dnsServerMonitor, err := resourceDataToDNSServerMonitor(d, client)
	if err != nil {
		return err
	}

	dnsServerMonitor, err = client.DNSServerMonitors().Create(dnsServerMonitor)
	if err != nil {
		return err
	}

	d.SetId(dnsServerMonitor.MonitorID)

	// return dnsServerMonitorRead(d, meta)
	return nil
}

func dnsServerMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	dnsServerMonitor, err := client.DNSServerMonitors().Get(d.Id())
	if err != nil {
		return err
	}

	updateDNSServerMonitorResourceData(d, dnsServerMonitor)

	return nil
}

func dnsServerMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	dnsServerMonitor, err := resourceDataToDNSServerMonitor(d, client)
	if err != nil {
		return err
	}

	dnsServerMonitor, err = client.DNSServerMonitors().Update(dnsServerMonitor)
	if err != nil {
		return err
	}

	d.SetId(dnsServerMonitor.MonitorID)

	// return websiteMonitorRead(d, meta)
	return nil
}

func dnsServerMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.DNSServerMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func dnsServerMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.DNSServerMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToDNSServerMonitor(d *schema.ResourceData, client site24x7.Client) (*api.DNSServerMonitor, error) {

	dnsServerMonitor := &api.DNSServerMonitor{
		MonitorID:      d.Id(),
		DisplayName:    d.Get("display_name").(string),
		Type:           string(api.DNS),
		CheckFrequency: d.Get("check_frequency").(string),
		Timeout:        d.Get("timeout").(int),
		DNSHost:        d.Get("dns_host").(string),
		DNSPort:        d.Get("dns_port").(string),
		UseIPV6:        d.Get("use_ipv6").(bool),
		DomainName:     d.Get("domain_name").(string),
		DNSSEC:         d.Get("dnssec").(bool),
		DeepDiscovery:  d.Get("deep_discovery").(bool),
		LookupType:     d.Get("lookup_type").(int),
	}

	// ================================ Configuration Profiles ================================
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

	actionMap := d.Get("actions").(map[string]interface{})
	keys := make([]string, 0, len(actionMap))
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

	searchConfigMap := d.Get("search_config").(*schema.Set)
	searchConfigItems := make([]api.SearchConfig, searchConfigMap.Len())
	for k, v := range searchConfigMap.List() {
		log.Println("Lookup Type is " + fmt.Sprint(v.(map[string]interface{})["lookup_type"]) + " in Site24x7")
		lookup_type := v.(map[string]interface{})["lookup_type"]
		switch lookup_type {
		case "A":
			searchConfigItems[k] = api.SearchConfig{
				Addr: v.(map[string]interface{})["addr"].(string),
				TTLO: v.(map[string]interface{})["ttlo"].(int),
				TTL:  v.(map[string]interface{})["ttl"].(int),
			}
		case "AAAA":
			searchConfigItems[k] = api.SearchConfig{
				Addr: v.(map[string]interface{})["addr"].(string),
				TTLO: v.(map[string]interface{})["ttlo"].(int),
				TTL:  v.(map[string]interface{})["ttl"].(int),
			}
		case "NS":
			searchConfigItems[k] = api.SearchConfig{
				Target: v.(map[string]interface{})["target"].(string),
				TTLO:   v.(map[string]interface{})["ttlo"].(int),
				TTL:    v.(map[string]interface{})["ttl"].(int),
			}
		case "CNAME":
			searchConfigItems[k] = api.SearchConfig{
				Target: v.(map[string]interface{})["target"].(string),
				TTLO:   v.(map[string]interface{})["ttlo"].(int),
				TTL:    v.(map[string]interface{})["ttl"].(int),
			}
		case "PTR":
			searchConfigItems[k] = api.SearchConfig{
				Target: v.(map[string]interface{})["target"].(string),
				TTLO:   v.(map[string]interface{})["ttlo"].(int),
				TTL:    v.(map[string]interface{})["ttl"].(int),
			}
		case "MX":
			searchConfigItems[k] = api.SearchConfig{
				Target:   v.(map[string]interface{})["target"].(string),
				Priority: v.(map[string]interface{})["target"].(int),
				TTLO:     v.(map[string]interface{})["ttlo"].(int),
				TTL:      v.(map[string]interface{})["ttl"].(int),
			}
		case "SRV":
			searchConfigItems[k] = api.SearchConfig{
				Port:     v.(map[string]interface{})["port"].(int),
				Target:   v.(map[string]interface{})["target"].(string),
				Wt:       v.(map[string]interface{})["wt"].(int),
				Priority: v.(map[string]interface{})["target"].(int),
				TTLO:     v.(map[string]interface{})["ttlo"].(int),
				TTL:      v.(map[string]interface{})["ttl"].(int),
			}
		case "TXT":
			searchConfigItems[k] = api.SearchConfig{
				Rcvd: v.(map[string]interface{})["rcvd"].(int),
				TTLO: v.(map[string]interface{})["ttlo"].(int),
				TTL:  v.(map[string]interface{})["ttl"].(int),
			}
		case "SOA":
			searchConfigItems[k] = api.SearchConfig{
				PNS:    v.(map[string]interface{})["pns"].(string),
				Admin:  v.(map[string]interface{})["admin"].(string),
				Serial: v.(map[string]interface{})["serial"].(int),
				RFF:    v.(map[string]interface{})["rff"].(int),
				RTF:    v.(map[string]interface{})["rtf"].(int),
				EXPT:   v.(map[string]interface{})["expt"].(int),
				MTTL:   v.(map[string]interface{})["mttl"].(int),
				TTLO:   v.(map[string]interface{})["ttlo"].(int),
				TTL:    v.(map[string]interface{})["ttl"].(int),
			}
		case "DNSKEY":
			searchConfigItems[k] = api.SearchConfig{
				Flg:   v.(map[string]interface{})["flg"].(int),
				Prtcl: v.(map[string]interface{})["prtcl"].(int),
				Kalg:  v.(map[string]interface{})["kalg"].(int),
				Kid:   v.(map[string]interface{})["kid"].(int),
				Key:   v.(map[string]interface{})["key"].(string),
				TTLO:  v.(map[string]interface{})["ttlo"].(int),
				TTL:   v.(map[string]interface{})["ttl"].(int),
			}
		case "CAA":
			searchConfigItems[k] = api.SearchConfig{
				Tag:  v.(map[string]interface{})["tag"].(string),
				CertAuth:  v.(map[string]interface{})["cert_auth"].(string),
				Flg:   v.(map[string]interface{})["flg"].(int),
				TTLO: v.(map[string]interface{})["ttlo"].(int),
				TTL:  v.(map[string]interface{})["ttl"].(int),
			}
		case "DS":
			searchConfigItems[k] = api.SearchConfig{
				Kid:   v.(map[string]interface{})["kid"].(int),
				Kalg:  v.(map[string]interface{})["kalg"].(int),
				Halg:  v.(map[string]interface{})["halg"].(int),
				Hash:  v.(map[string]interface{})["hash"].(string),
				TTLO: v.(map[string]interface{})["ttlo"].(int),
				TTL:  v.(map[string]interface{})["ttl"].(int),
			}
		}
	}

	dnsServerMonitor.LocationProfileID = d.Get("location_profile_id").(string)
	dnsServerMonitor.NotificationProfileID = d.Get("notification_profile_id").(string)
	dnsServerMonitor.ThresholdProfileID = d.Get("threshold_profile_id").(string)
	dnsServerMonitor.OnCallScheduleID = d.Get("on_call_schedule_id").(string)
	dnsServerMonitor.MonitorGroups = monitorGroups
	dnsServerMonitor.DependencyResourceIDs = dependencyResourceIDs
	dnsServerMonitor.UserGroupIDs = userGroupIDs
	dnsServerMonitor.TagIDs = tagIDs
	dnsServerMonitor.ThirdPartyServiceIDs = thirdPartyServiceIDs
	dnsServerMonitor.SearchConfig = searchConfigItems
	dnsServerMonitor.ActionIDs = actionRefs

	// Location Profile
	_, locationProfileErr := site24x7.SetLocationProfile(client, d, dnsServerMonitor)
	if locationProfileErr != nil {
		return nil, locationProfileErr
	}

	// Threshold Profile
	if dnsServerMonitor.ThresholdProfileID == "" {
		profile, err := site24x7.DefaultThresholdProfile(client, api.URL)
		if err != nil {
			return nil, err
		}
		dnsServerMonitor.ThresholdProfileID = profile.ProfileID
		d.Set("threshold_profile_id", profile)
	}
	// Notification Profile
	_, notificationProfileErr := site24x7.SetNotificationProfile(client, d, dnsServerMonitor)
	if notificationProfileErr != nil {
		return nil, notificationProfileErr
	}
	// User Alert Groups
	_, userAlertGroupErr := site24x7.SetUserGroup(client, d, dnsServerMonitor)
	if userAlertGroupErr != nil {
		return nil, userAlertGroupErr
	}
	// Tags
	_, tagsErr := site24x7.SetTags(client, d, dnsServerMonitor)
	if tagsErr != nil {
		return nil, tagsErr
	}

	return dnsServerMonitor, nil
}

func updateDNSServerMonitorResourceData(d *schema.ResourceData, monitor *api.DNSServerMonitor) {
	d.Set("display_name", monitor.DisplayName)
	d.Set("type", monitor.Type)
	d.Set("dns_host", monitor.DNSHost)
	d.Set("dns_port", monitor.DNSPort)
	d.Set("use_ipv6", monitor.UseIPV6)
	d.Set("domain_name", monitor.DomainName)
	d.Set("timeout", monitor.Timeout)
	d.Set("notification_profile_id", monitor.NotificationProfileID)
	d.Set("tag_ids", monitor.TagIDs)
	d.Set("user_group_ids", monitor.UserGroupIDs)
	d.Set("location_profile_id", monitor.LocationProfileID)
	d.Set("notification_profile_id", monitor.NotificationProfileID)
	d.Set("threshold_profile_id", monitor.ThresholdProfileID)
	d.Set("user_group_ids", monitor.UserGroupIDs)
	d.Set("monitor_groups", monitor.MonitorGroups)
	d.Set("dependency_resource_ids", monitor.DependencyResourceIDs)
	d.Set("lookup_type", monitor.LookupType)
	d.Set("dnssec", monitor.DNSSEC)
	actions := make(map[string]interface{})
	for _, action := range monitor.ActionIDs {
		actions[fmt.Sprintf("%d", action.AlertType)] = action.ActionID
	}
	d.Set("actions", actions)
	d.Set("third_party_service_ids", monitor.ThirdPartyServiceIDs)
	d.Set("tag_ids", monitor.TagIDs)
	d.Set("deep_discovery", monitor.DeepDiscovery)
	d.Set("check_frequency", monitor.CheckFrequency)
}
