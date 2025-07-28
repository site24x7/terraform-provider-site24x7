package monitors

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var AzureMonitorSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name for the Azure monitor.",
	},
	"tenant_id": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "The Tenant ID associated with the Microsoft Entra ID.",
	},
	"client_id": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "The Client ID for the Azure Service Principal.",
	},
	"client_secret": {
		Type:        schema.TypeString,
		Required:    true,
		Sensitive:   true,
		Description: "The Client Secret associated with the Azure Service Principal.",
	},
	"type": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Type of the monitor (should be AZURE).",
	},
	"services": {
		Type:        schema.TypeList,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Required:    true,
		Description: "List of Azure service types to be discovered.",
	},
	"management_group_reg": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "Use 0 for Azure Account, 1 for Management Group.",
	},
	"notification_profile_id": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Notification profile associated with the monitor.",
	},
	"user_group_ids": {
		Type:        schema.TypeList,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Required:    true,
		Description: "User group IDs to be notified during outages.",
	},
	"threshold_profile_id": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Threshold profile ID to be associated with the monitor.",
	},
	"discovery_interval": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Rediscovery interval (e.g., 30, 60, 360, etc.).",
	},
	"auto_add_subscription": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Automatically add newly-added subscriptions (1 for Yes, 0 for No).",
	},
	"azure_exclude_tags": {
		Type:        schema.TypeMap,
		Optional:    true,
		Description: "Tags to exclude Azure resources from discovery.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:     schema.TypeInt,
					Required: true,
				},
				"tags": {
					Type:     schema.TypeMap,
					Required: true,
					Elem: &schema.Schema{
						Type: schema.TypeList,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
				},
			},
		},
	},
	"azure_include_tags": {
		Type:        schema.TypeMap,
		Optional:    true,
		Description: "Tags to include Azure resources in discovery.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:     schema.TypeInt,
					Required: true,
				},
				"tags": {
					Type:     schema.TypeMap,
					Required: true,
					Elem: &schema.Schema{
						Type: schema.TypeList,
						Elem: &schema.Schema{Type: schema.TypeString},
					},
				},
			},
		},
	},
}

func ResourceSite24x7AzureMonitor() *schema.Resource {
	return &schema.Resource{
		Create: azureMonitorCreate,
		Read:   azureMonitorRead,
		Update: azureMonitorUpdate,
		Delete: azureMonitorDelete,
		Exists: azureMonitorExists,
		Schema: AzureMonitorSchema,
	}
}

func azureMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)
	monitor, err := resourceDataToAzureMonitor(d, client)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Azure Monitor payload: %+v", monitor)

	azureMonitor, err := client.AzureMonitors().Create(monitor)
	if err != nil {
		return err
	}
	d.SetId(azureMonitor.MonitorID)
	return nil
}

func azureMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)
	azureMonitor, err := client.AzureMonitors().Get(d.Id())
	if err != nil {
		return err
	}
	updateAzureMonitorResourceData(d, azureMonitor)
	return nil
}

func azureMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)
	monitor, err := resourceDataToAzureMonitor(d, client)
	if err != nil {
		return err
	}
	azureMonitor, err := client.AzureMonitors().Update(monitor)
	if err != nil {
		return err
	}
	d.SetId(azureMonitor.MonitorID)
	return nil
}

func azureMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)
	err := client.AzureMonitors().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}
	return err
}

func azureMonitorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)
	_, err := client.AzureMonitors().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}
	return err == nil, err
}

func expandAzureTagCondition(input map[string]interface{}) *api.AzureTagCondition {
	if input == nil {
		return nil
	}

	// Guard against missing or nil values
	typeVal, typeOk := input["type"]
	tagsVal, tagsOk := input["tags"]

	if !typeOk || !tagsOk {
		return nil // or handle the error appropriately
	}

	// Ensure type is an int
	typeInt, ok := typeVal.(int)
	if !ok {
		// Log error or handle the unexpected type
		return nil // or return an error if necessary
	}

	tagCondition := &api.AzureTagCondition{
		Type: typeInt,
		Tags: make(map[string][]string),
	}

	// Check if 'tags' is of the correct type
	rawTags, ok := tagsVal.(map[string]interface{})
	if !ok {
		// Handle the unexpected type
		return tagCondition // or handle accordingly
	}

	for k, v := range rawTags {
		interfaceList, ok := v.([]interface{})
		if !ok {
			continue // or handle the invalid type
		}
		strList := make([]string, len(interfaceList))
		for i, val := range interfaceList {
			strList[i] = val.(string)
		}
		tagCondition.Tags[k] = strList
	}

	return tagCondition
}

func flattenAzureTagCondition(condition *api.AzureTagCondition) map[string]interface{} {
	if condition == nil {
		return nil
	}

	tags := make(map[string]interface{})
	for k, v := range condition.Tags {
		strList := make([]interface{}, len(v))
		for i, s := range v {
			strList[i] = s
		}
		tags[k] = strList
	}

	return map[string]interface{}{
		"type": condition.Type,
		"tags": tags,
	}
}

func resourceDataToAzureMonitor(d *schema.ResourceData, client site24x7.Client) (*api.AzureMonitor, error) {
	var userGroupIDs []string
	for _, id := range d.Get("user_group_ids").([]interface{}) {
		if id != nil {
			userGroupIDs = append(userGroupIDs, id.(string))
		}
	}

	var services []string
	for _, v := range d.Get("services").([]interface{}) {
		if serviceType, ok := v.(string); ok {
			services = append(services, serviceType)
		}
	}

	monitor := &api.AzureMonitor{
		MonitorID:             d.Id(),
		DisplayName:           d.Get("display_name").(string),
		TenantID:              d.Get("tenant_id").(string),
		ClientID:              d.Get("client_id").(string),
		ClientSecret:          d.Get("client_secret").(string),
		Type:                  d.Get("type").(string),
		Services:              services,
		ManagementGroupReg:    d.Get("management_group_reg").(int),
		NotificationProfileID: d.Get("notification_profile_id").(string),
		UserGroupIDs:          userGroupIDs,
		ThresholdProfileID:    d.Get("threshold_profile_id").(string),
		DiscoveryInterval:     d.Get("discovery_interval").(string),
		AutoAddSubscription:   d.Get("auto_add_subscription").(int),
		AzureExcludeTags:      expandAzureTagCondition(d.Get("azure_exclude_tags").(map[string]interface{})),
		AzureIncludeTags:      expandAzureTagCondition(d.Get("azure_include_tags").(map[string]interface{})),
	}
	return monitor, nil
}

func updateAzureMonitorResourceData(d *schema.ResourceData, m *api.AzureMonitor) {
	d.Set("display_name", m.DisplayName)
	d.Set("type", m.Type)
	d.Set("services", m.Services)
	d.Set("management_group_reg", m.ManagementGroupReg)
	d.Set("notification_profile_id", m.NotificationProfileID)
	d.Set("user_group_ids", m.UserGroupIDs)
	d.Set("threshold_profile_id", m.ThresholdProfileID)
	d.Set("discovery_interval", m.DiscoveryInterval)
	d.Set("auto_add_subscription", m.AutoAddSubscription)
	d.Set("azure_exclude_tags", flattenAzureTagCondition(m.AzureExcludeTags))
	d.Set("azure_include_tags", flattenAzureTagCondition(m.AzureIncludeTags))
}
