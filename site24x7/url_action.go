package site24x7

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
)

// SAMPLE POST JSON
// {
// 	"action_name": "URL IT action",
// 	"action_type": 1,
// 	"action_timeout": 15,
// 	"action_url": "http://www.example.com",
// 	"action_method": "G",
// 	"custom_headers": [
// 	  {
// 		"name": "testheader",
// 		"value": "header value"
// 	  }
// 	],
// 	"send_mail": false,
// 	"auth_method": "B",
// 	"user_agent": "user agent",
// 	"send_custom_parameters": true,
// 	"custom_parameters": "customparam=value"
//   }

var URLActionSchema = map[string]*schema.Schema{
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"url": {
		Type:     schema.TypeString,
		Required: true,
	},
	"type": {
		Type:     schema.TypeInt,
		Optional: true,
		Default:  1,
	},
	"method": {
		Type:     schema.TypeString,
		Optional: true,
		Default:  "G",
	},
	"timeout": {
		Type:     schema.TypeInt,
		Optional: true,
		Default:  15,
	},
	// @TODO(mohmann): requires_authentication is no valid field in the
	// URLAutomations API anymore and is thus ignored. We should remove
	// it completely from the resource in the future. This is just here
	// for backwards compatibility.
	"requires_authentication": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	},
	"custom_parameters": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"send_custom_parameters": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	},
	"send_in_json_format": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	},
	"send_email": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	},
	"send_incident_parameters": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	},
	"auth_method": {
		Type:     schema.TypeString,
		Optional: true,
		Default:  "B",
	},
	"user_agent": {
		Type:     schema.TypeString,
		Optional: true,
	},
}

func resourceSite24x7URLAction() *schema.Resource {
	return &schema.Resource{
		Create: urlActionCreate,
		Read:   urlActionRead,
		Update: urlActionUpdate,
		Delete: urlActionDelete,
		Exists: urlActionExists,

		Schema: URLActionSchema,
	}
}

func urlActionCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	automation := resourceDataToUrlAction(d)

	automation, err := client.URLAutomations().Create(automation)
	if err != nil {
		return err
	}

	d.SetId(automation.ActionID)

	return nil
}

func urlActionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	automation, err := client.URLAutomations().Get(d.Id())
	if err != nil {
		return err
	}

	updateUrlActionResourceData(d, automation)

	return nil
}

func urlActionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	automation := resourceDataToUrlAction(d)

	automation, err := client.URLAutomations().Update(automation)
	if err != nil {
		return err
	}

	d.SetId(automation.ActionID)

	return nil
}

func urlActionDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	err := client.URLAutomations().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func urlActionExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(Client)

	_, err := client.URLAutomations().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToUrlAction(d *schema.ResourceData) *api.URLAutomation {
	return &api.URLAutomation{
		ActionID:               d.Id(),
		ActionName:             d.Get("name").(string),
		ActionType:             1,
		ActionUrl:              d.Get("url").(string),
		ActionMethod:           d.Get("method").(string),
		ActionTimeout:          d.Get("timeout").(int),
		SendEmail:              d.Get("send_email").(bool),
		SendCustomParameters:   d.Get("send_custom_parameters").(bool),
		CustomParameters:       d.Get("custom_parameters").(string),
		AuthMethod:             d.Get("auth_method").(string),
		SendInJsonFormat:       d.Get("send_in_json_format").(bool),
		SendIncidentParameters: d.Get("send_incident_parameters").(bool),
		UserAgent:              d.Get("user_agent").(string),
	}
}

func updateUrlActionResourceData(d *schema.ResourceData, automation *api.URLAutomation) {
	d.Set("name", automation.ActionName)
	d.Set("type", 1)
	d.Set("url", automation.ActionUrl)
	d.Set("method", automation.ActionMethod)
	d.Set("timeout", automation.ActionTimeout)
	d.Set("send_email", automation.SendEmail)
	d.Set("send_custom_parameters", automation.SendCustomParameters)
	d.Set("custom_parameters", automation.CustomParameters)
	d.Set("auth_method", automation.AuthMethod)
	d.Set("send_in_json_format", automation.SendInJsonFormat)
	d.Set("send_incident_parameters", automation.SendIncidentParameters)
	d.Set("user_agent", automation.UserAgent)
}
