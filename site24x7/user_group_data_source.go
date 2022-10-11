package site24x7

import (
	"errors"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
)

var userGroupDataSourceSchema = map[string]*schema.Schema{
	"name_regex": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Regular expression denoting the name of the user group.",
	},
	"matching_ids": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "List of user group IDs matching the name_regex.",
	},
	"matching_ids_and_names": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "List of user group IDs and names matching the name_regex.",
	},
	"display_name": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Display name for the user group.",
	},
	"attribute_group_id": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Attribute alert group associated with the user alert group.",
	},
	"users": {
		Type:     schema.TypeSet,
		Computed: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Description: "User IDs of the users associated to the group.",
	},
	"product_id": {
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "Product for which the user group was created.",
	},
}

func DataSourceSite24x7UserGroup() *schema.Resource {
	return &schema.Resource{
		Read:   userGroupDataSourceRead,
		Schema: userGroupDataSourceSchema,
	}
}

// userGroupDataSourceRead fetches all userGroup from Site24x7
func userGroupDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	userGroupList, err := client.UserGroups().List()
	if err != nil {
		return err
	}

	nameRegex := d.Get("name_regex")

	var userGroup *api.UserGroup
	var matchingUserGroupIDs []string
	var matchingUserGroupIDsAndNames []string

	if nameRegex != "" {
		// (?i) - Case insensitive match
		nameRegexPattern := regexp.MustCompile("(?i)" + nameRegex.(string))
		for _, groupInfo := range userGroupList {
			if len(groupInfo.DisplayName) > 0 {
				if userGroup == nil && nameRegexPattern.MatchString(groupInfo.DisplayName) {
					userGroup = new(api.UserGroup)
					userGroup.UserGroupID = groupInfo.UserGroupID
					userGroup.DisplayName = groupInfo.DisplayName
					userGroup.Users = groupInfo.Users
					userGroup.AttributeGroupID = groupInfo.AttributeGroupID
					userGroup.ProductID = groupInfo.ProductID
					matchingUserGroupIDs = append(matchingUserGroupIDs, groupInfo.UserGroupID)
					matchingUserGroupIDsAndNames = append(matchingUserGroupIDsAndNames, groupInfo.UserGroupID+"__"+groupInfo.DisplayName)
				} else if userGroup != nil && nameRegexPattern.MatchString(groupInfo.DisplayName) {
					matchingUserGroupIDs = append(matchingUserGroupIDs, groupInfo.UserGroupID)
					matchingUserGroupIDsAndNames = append(matchingUserGroupIDsAndNames, groupInfo.UserGroupID+"__"+groupInfo.DisplayName)
				}
			}
		}
	} else {
		return errors.New("Please enter a value for the attribute name_regex!")
	}

	if userGroup == nil {
		return errors.New("Unable to find user group matching the name : \"" + d.Get("name_regex").(string))
	}

	updateUserGroupDataSourceResourceData(d, userGroup, matchingUserGroupIDs, matchingUserGroupIDsAndNames)

	return nil
}

func updateUserGroupDataSourceResourceData(d *schema.ResourceData, userGroup *api.UserGroup, matchingUserGroupIDs []string, matchingUserGroupIDsAndNames []string) {
	d.SetId(userGroup.UserGroupID)
	d.Set("matching_ids", matchingUserGroupIDs)
	d.Set("matching_ids_and_names", matchingUserGroupIDsAndNames)
	d.Set("display_name", userGroup.DisplayName)
	d.Set("attribute_group_id", userGroup.AttributeGroupID)
	d.Set("users", userGroup.Users)
	d.Set("product_id", userGroup.ProductID)
}
