package site24x7

import (
	"errors"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
)

var itAutomationDataSourceSchema = map[string]*schema.Schema{
	"name_regex": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Regular expression denoting the name of the IT action.",
	},
	"matching_ids": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "List of IT action IDs matching the name_regex.",
	},
	"matching_ids_and_names": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "List of IT action IDs and names matching the name_regex.",
	},
	"action_name": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Display name for the Action.",
	},
	"action_type": {
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "Type of the Action.",
	},
	"url": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "URL invoked for action execution.",
	},
	"method": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "HTTP Method to access the action url.",
	},
	"timeout": {
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "Timeout for connecting the website.",
	},
	"send_custom_parameters": {
		Type:        schema.TypeBool,
		Computed:    true,
		Description: "Custom parameters sent while executing the action.",
	},
	"custom_parameters": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Custom parameters passed while accessing the action url.",
	},
	"send_in_json_format": {
		Type:        schema.TypeBool,
		Computed:    true,
		Description: "Configuration to enable json format for post parameters.",
	},
	"send_email": {
		Type:        schema.TypeBool,
		Computed:    true,
		Description: "Boolean indicating whether to send email or not.",
	},
	"send_incident_parameters": {
		Type:        schema.TypeBool,
		Computed:    true,
		Description: "Configuration to send incident parameters while executing the action.",
	},
	"auth_method": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Authentication method to access the action url.",
	},
	"user_agent": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "User Agent used while monitoring the website.",
	},
}

func DataSourceSite24x7ITAutomation() *schema.Resource {
	return &schema.Resource{
		Read:   itAutomationDataSourceRead,
		Schema: itAutomationDataSourceSchema,
	}
}

// itAutomationDataSourceRead fetches all itAutomation from Site24x7
func itAutomationDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	itAutomationList, err := client.URLActions().List()
	if err != nil {
		return err
	}

	nameRegex := d.Get("name_regex")

	var itAutomation *api.URLAction
	var matchingITAutomationIDs []string
	var matchingITAutomationIDsAndNames []string
	if nameRegex != "" {
		// (?i) - Case insensitive match
		nameRegexPattern := regexp.MustCompile("(?i)" + nameRegex.(string))
		for _, actionInfo := range itAutomationList {
			if len(actionInfo.ActionName) > 0 {
				if itAutomation == nil && nameRegexPattern.MatchString(actionInfo.ActionName) {
					itAutomation = new(api.URLAction)
					itAutomation.ActionID = actionInfo.ActionID
					itAutomation.ActionName = actionInfo.ActionName
					itAutomation.ActionType = actionInfo.ActionType
					itAutomation.ActionUrl = actionInfo.ActionUrl
					itAutomation.ActionTimeout = actionInfo.ActionTimeout
					itAutomation.ActionMethod = actionInfo.ActionMethod
					itAutomation.SuppressAlert = actionInfo.SuppressAlert
					itAutomation.SendIncidentParameters = actionInfo.SendIncidentParameters
					itAutomation.SendCustomParameters = actionInfo.SendCustomParameters
					itAutomation.CustomParameters = actionInfo.CustomParameters
					itAutomation.SendInJsonFormat = actionInfo.SendInJsonFormat
					itAutomation.SendEmail = actionInfo.SendEmail
					itAutomation.AuthMethod = actionInfo.AuthMethod
					itAutomation.OAuth2Provider = actionInfo.OAuth2Provider
					itAutomation.UserAgent = actionInfo.UserAgent

					matchingITAutomationIDs = append(matchingITAutomationIDs, actionInfo.ActionID)
					matchingITAutomationIDsAndNames = append(matchingITAutomationIDsAndNames, actionInfo.ActionID+"__"+actionInfo.ActionName)
				} else if itAutomation != nil && nameRegexPattern.MatchString(actionInfo.ActionName) {
					matchingITAutomationIDs = append(matchingITAutomationIDs, actionInfo.ActionID)
					matchingITAutomationIDsAndNames = append(matchingITAutomationIDsAndNames, actionInfo.ActionID+"__"+actionInfo.ActionName)
				}
			}
		}
	} else {
		return errors.New("Please enter a value for the attribute name_regex!")
	}

	if itAutomation == nil {
		return errors.New("Unable to find IT action matching the name : \"" + d.Get("name_regex").(string))
	}

	updateITAutomationDataSourceResourceData(d, itAutomation, matchingITAutomationIDs, matchingITAutomationIDsAndNames)

	return nil
}

func updateITAutomationDataSourceResourceData(d *schema.ResourceData, itAutomation *api.URLAction, matchingITAutomationIDs []string, matchingITAutomationIDsAndNames []string) {
	d.SetId(itAutomation.ActionID)
	d.Set("matching_ids", matchingITAutomationIDs)
	d.Set("matching_ids_and_names", matchingITAutomationIDsAndNames)
	d.Set("action_name", itAutomation.ActionName)
	d.Set("action_type", itAutomation.ActionType)
	d.Set("url", itAutomation.ActionUrl)
	d.Set("method", itAutomation.ActionMethod)
	d.Set("timeout", itAutomation.ActionTimeout)
	d.Set("send_custom_parameters", itAutomation.SendCustomParameters)
	d.Set("custom_parameters", itAutomation.CustomParameters)
	d.Set("send_in_json_format", itAutomation.SendInJsonFormat)
	d.Set("send_email", itAutomation.SendEmail)
	d.Set("send_incident_parameters", itAutomation.SendIncidentParameters)
	d.Set("auth_method", itAutomation.AuthMethod)
	d.Set("user_agent", itAutomation.UserAgent)
}
