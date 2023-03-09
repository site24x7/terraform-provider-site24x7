package site24x7

import (
	"errors"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
)

var subgroupDataSourceSchema = map[string]*schema.Schema{
	"name_regex": {
		Type:     schema.TypeString,
		Required: true,
		// ValidateFunc: validation.StringIsValidRegExp,
	},
	"display_name": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Display Name for the Subgroup.",
	},
	"description": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Description for the Subgroup.",
	},
	"parent_group_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Unique ID of the parent group under which subgroup has to be configured. It can be a subgroup or Monitor group. (In case of level 1 subgroup, top_group_id is monitor group id. In other cases it will be subgroup id. You can get the subgroup Ids configured for top_group_id by using business view API).",
	},
	"top_group_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Unique ID of the top monitor group for which business view has been configured.",
	},
	"group_type": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     1,
		Description: "Denotes the type of monitors that can be associated. ‘1’ implies that all type of monitors can be associated with this subgroup. Default value is 1. '2' - Web, '3' - Port/Ping, '4' - Server, '5' - Database, '6' - Synthetic Transaction, '7' - Web API, '8' - APM Insight,'9' - Network Devices, '10' - RUM, '11' - AppLogs Monitor",
	},
	"health_threshold_count": {
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "Number of monitors' health that decide the group status. ‘0’ implies that all the monitors are considered for determining the group status.",
	},
	"monitors": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Computed:    true,
		Description: "List of monitors associated to the group.",
	},
}

func DataSourceSite24x7Subgroup() *schema.Resource {
	return &schema.Resource{
		Read:   subgroupDataSourceRead,
		Schema: subgroupDataSourceSchema,
	}
}

// subgroupDataSourceRead fetches all subgroup from Site24x7
func subgroupDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	subgroupList, err := client.Subgroups().List()
	if err != nil {
		return err
	}

	nameRegex := d.Get("name_regex")

	var subgroup *api.Subgroup
	if nameRegex != "" {
		// (?i) - Case insensitive match
		nameRegexPattern := regexp.MustCompile("(?i)" + nameRegex.(string))
		for _, groupInfo := range subgroupList {
			if len(groupInfo.DisplayName) > 0 {
				if nameRegexPattern.MatchString(groupInfo.DisplayName) {
					subgroup = new(api.Subgroup)
					subgroup.ID = groupInfo.ID
					subgroup.DisplayName = groupInfo.DisplayName
					subgroup.Description = groupInfo.Description
					subgroup.HealthThresholdCount = groupInfo.HealthThresholdCount
					subgroup.Monitors = groupInfo.Monitors
					subgroup.ParentGroupID = groupInfo.ParentGroupID
					subgroup.TopGroupID = groupInfo.TopGroupID
					subgroup.Type = groupInfo.Type
					break
				}
			}
		}
	} else {
		return errors.New("Please enter a value for the attribute name_regex!")
	}

	if subgroup == nil {
		return errors.New("Unable to find subgroup matching the name : \"" + d.Get("name_regex").(string))
	}

	updateSubgroupDataSourceResourceData(d, subgroup)

	return nil
}

func updateSubgroupDataSourceResourceData(d *schema.ResourceData, subgroup *api.Subgroup) {
	d.SetId(subgroup.ID)
	d.Set("display_name", subgroup.DisplayName)
	d.Set("description", subgroup.Description)
	d.Set("monitors", subgroup.Monitors)
	d.Set("parent_group_id", subgroup.ParentGroupID)
	d.Set("top_group_id", subgroup.TopGroupID)
	d.Set("group_type", subgroup.Type)
	d.Set("health_threshold_count", subgroup.HealthThresholdCount)
	d.Set("monitors", subgroup.Monitors)
}
