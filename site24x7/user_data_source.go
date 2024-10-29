package site24x7

import (
	"errors"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
)

var userDataSourceSchema = map[string]*schema.Schema{
	"name_regex": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Regular expression denoting the name of the user.",
	},
	"matching_ids": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "List of user IDs matching the name_regex.",
	},
	"matching_ids_and_names": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "List of user IDs and names matching the name_regex.",
	},
	"display_name": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Display name for the user.",
	},
	"email": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Email address of the user.",
	},
	"role": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Role assigned to the user.",
	},
	"status": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Status of the user (active or inactive).",
	},
}

func DataSourceSite24x7User() *schema.Resource {
	return &schema.Resource{
		Read:   userDataSourceRead,
		Schema: userDataSourceSchema,
	}
}

// userDataSourceRead fetches users from Site24x7 based on name_regex
func userDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	userList, err := client.Users().List()
	if err != nil {
		return err
	}

	nameRegex := d.Get("name_regex").(string)
	var matchingUserIDs []string
	var matchingUserIDsAndNames []string
	var user *api.User

	if nameRegex != "" {
		nameRegexPattern := regexp.MustCompile("(?i)" + nameRegex)
		for _, userInfo := range userList {
			if nameRegexPattern.MatchString(userInfo.DisplayName) {
				user = userInfo
				matchingUserIDs = append(matchingUserIDs, userInfo.ID)
				matchingUserIDsAndNames = append(matchingUserIDsAndNames, userInfo.ID+"__"+userInfo.DisplayName)
			}
		}
	} else {
		return errors.New("Please enter a value for the attribute name_regex!")
	}

	if user == nil {
		return errors.New("Unable to find user matching the name: \"" + nameRegex + "\"")
	}

	updateUserDataSourceResourceData(d, user, matchingUserIDs, matchingUserIDsAndNames)

	return nil
}

func updateUserDataSourceResourceData(d *schema.ResourceData, user *api.User, matchingUserIDs []string, matchingUserIDsAndNames []string) {
	d.SetId(user.ID) // Set the ID to the matched user's ID
	d.Set("matching_ids", matchingUserIDs)
	d.Set("matching_ids_and_names", matchingUserIDsAndNames)
	d.Set("display_name", user.DisplayName)
	d.Set("email", user.Email)
	d.Set("user_role", user.UserRole)
}
