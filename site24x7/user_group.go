package site24x7

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
)

// SAMPLE POST JSON
// {
// 	"display_name": "test_terraform"
// 	"attribute_group_id": "111111000021528003",
// 	"users": [
// 	  "111111000005937003"
// 	],
// 	"product_id": 0,
// }

var UserGroupSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name for the user group.",
	},
	"attribute_group_id": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Attribute Alert Group to be associated with the User Alert group.",
	},
	"users": {
		Type:     schema.TypeSet,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Description: "User IDs of the users to be associated to the group.",
	},
	"product_id": {
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      0,
		ValidateFunc: validation.IntInSlice([]int{0, 1, 2}),
		Description:  "Product for which the user group is being created. Default value is 0.",
	},
}

func ResourceSite24x7UserGroup() *schema.Resource {
	return &schema.Resource{
		Create: userGroupCreate,
		Read:   userGroupRead,
		Update: userGroupUpdate,
		Delete: userGroupDelete,
		Exists: userGroupExists,

		Schema: UserGroupSchema,
	}
}

func userGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	userGroup := resourceDataToUserGroup(d)

	userGroup, err := client.UserGroups().Create(userGroup)
	if err != nil {
		return err
	}

	d.SetId(userGroup.UserGroupID)

	return nil
}

func userGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	userGroup, err := client.UserGroups().Get(d.Id())
	if err != nil {
		return err
	}

	updateUserGroupResourceData(d, userGroup)

	return nil
}

func userGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	userGroup := resourceDataToUserGroup(d)

	userGroup, err := client.UserGroups().Update(userGroup)
	if err != nil {
		return err
	}

	d.SetId(userGroup.UserGroupID)

	return nil
}

func userGroupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	err := client.UserGroups().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func userGroupExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(Client)

	_, err := client.UserGroups().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToUserGroup(d *schema.ResourceData) *api.UserGroup {

	users := d.Get("users").(*schema.Set).List()
	userIDs := make([]string, 0, len(users))
	for _, v := range users {
		userIDs = append(userIDs, v.(string))
	}

	return &api.UserGroup{
		UserGroupID:      d.Id(),
		DisplayName:      d.Get("display_name").(string),
		Users:            userIDs,
		AttributeGroupID: d.Get("attribute_group_id").(string),
		ProductID:        d.Get("product_id").(int),
	}
}

// Called during read - populates the ResourceData with the userGroup in API response
func updateUserGroupResourceData(d *schema.ResourceData, userGroup *api.UserGroup) {
	d.Set("display_name", userGroup.DisplayName)
	d.Set("users", userGroup.Users)
	d.Set("attribute_group_id", userGroup.AttributeGroupID)
	d.Set("product_id", userGroup.ProductID)
}
