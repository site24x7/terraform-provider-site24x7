package site24x7

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
)

var WebhookIntegrationSchema = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name for the OpsGenie Integration.",
	},
	"url": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "URL to be invoked for action execution.",
	},
	"timeout": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "The amount of time a connection waits to time out.Range 1 - 45.",
	},
	"method": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "HTTP Method to access the URL.",
	},
	"selection_type": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "Resource Type associated with this integration.Monitor Group not supported.",
	},
	"is_poller_webhook": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "URL to be invoked from an On-Premise Poller agent.",
	},
	"poller": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Mandatory, if is_poller_webhook is set as true.Denotes On-Premise Poller ID.",
	},
	"send_incident_parameters": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "Configuration to send incident parameters while executing the action.",
	},
	"send_custom_parameters": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "Configuration to send custom parameters while executing the action.",
	},
	"custom_parameters": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Mandatory, if send_custom_parameters is set as true.Custom parameters to be passed while accessing the URL.",
	},
	"send_in_json_format": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Configuration to enable json format for post parameters.",
	},
	"auth_method": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Authentication method to access the action url.",
	},
	"username": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Username for Authentication.",
	},
	"password": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Password for Authentication.",
	},
	"oauth2_provider": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Provider ID of the OAuth Provider to be associated with the action.",
	},
	"user_agent": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "User Agent to be used while monitoring the website.",
	},
	"monitors": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "Monitors associated with the integration.",
	},
	"manage_tickets": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "Configuration to handle ticketing based integration.",
	},
	"update_url": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "URL to be invoked to update the request.",
	},
	"update_method": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "HTTP Method to access the URL.",
	},
	"update_send_incident_parameters": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "Configuration to send incident parameters while executing the action.",
	},
	"update_send_custom_parameters": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "Configuration to send custom parameters while executing the action.",
	},
	"update_custom_parameters": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Configuration to send custom parameters while executing the action.",
	},
	"close_url": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "URL to be invoked to close the request.",
	},
	"close_method": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "HTTP Method to access the URL.",
	},
	"close_send_incident_parameters": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "Configuration to send incident parameters while executing the action.",
	},
	"close_send_custom_parameters": {
		Type:        schema.TypeBool,
		Required:    true,
		Description: "Configuration to send custom parameters while executing the action.",
	},
	"close_custom_parameters": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Mandatory, When close_send_custom_parameters is set as true.Custom parameters to be passed while accessing the URL.",
	},
	"alert_tags_id": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "Tag id’s to be associated with the integration.",
	},
}

func resourceSite24x7WebhookIntegration() *schema.Resource {
	return &schema.Resource{
		Create: webhookIntegrationCreate,
		Read:   webhookIntegrationRead,
		Update: webhookIntegrationUpdate,
		Delete: webhookIntegrationDelete,
		Exists: webhookIntegrationExists,

		Schema: WebhookIntegrationSchema,
	}
}

func webhookIntegrationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	webhookIntegration, err := resourceDataToWebhookIntegration(d)
	if err != nil {
		return err
	}

	webhookIntegration, err = client.WebhookIntegration().Create(webhookIntegration)
	if err != nil {
		return err
	}

	d.SetId(webhookIntegration.ServiceID)

	return nil
}

func webhookIntegrationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	webhookIntegration, err := client.WebhookIntegration().Get(d.Id())
	if err != nil {
		return err
	}

	updateWebhookIntegrationResourceData(d, webhookIntegration)

	return nil
}

func webhookIntegrationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	webhookIntegration, err := resourceDataToWebhookIntegration(d)
	if err != nil {
		return err
	}

	webhookIntegration, err = client.WebhookIntegration().Update(webhookIntegration)
	if err != nil {
		return err
	}

	d.SetId(webhookIntegration.ServiceID)

	return nil
}

func webhookIntegrationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	err := client.ThirdPartyIntegrations().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func webhookIntegrationExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(Client)

	_, err := client.WebhookIntegration().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToWebhookIntegration(d *schema.ResourceData) (*api.WebhookIntegration, error) {

	var monitorsIDs []string
	for _, id := range d.Get("monitors").([]interface{}) {
		monitorsIDs = append(monitorsIDs, id.(string))
	}

	var alertTagIDs []string
	for _, id := range d.Get("alert_tags_id").([]interface{}) {
		alertTagIDs = append(alertTagIDs, id.(string))
	}

	webhookIntegration := &api.WebhookIntegration{
		ServiceID:                    d.Id(),
		Name:                         d.Get("name").(string),
		URL:                          d.Get("url").(string),
		Timeout:                      d.Get("timeout").(int),
		Method:                       d.Get("method").(string),
		SelectionType:                api.ResourceType(d.Get("selection_type").(int)),
		IsPollerWebhook:              d.Get("is_poller_webhook").(bool),
		Poller:                       d.Get("poller").(string),
		SendIncidentParameters:       d.Get("send_incident_parameters").(bool),
		SendCustomParameters:         d.Get("send_custom_parameters").(bool),
		CustomParameters:             d.Get("custom_parameters"),
		SendInJsonFormat:             d.Get("send_in_json_format").(bool),
		AuthMethod:                   d.Get("auth_method").(string),
		Username:                     d.Get("username").(string),
		Password:                     d.Get("password").(string),
		OauthProvider:                d.Get("oauth2_provider").(string),
		UserAgent:                    d.Get("user_agent").(string),
		Monitors:                     monitorsIDs,
		ManageTickets:                d.Get("manage_tickets").(bool),
		UpdateURL:                    d.Get("update_url").(string),
		UpdateMethod:                 d.Get("update_method").(string),
		UpdateSendIncidentParameters: d.Get("update_send_incident_parameters").(bool),
		UpdateSendCustomParameters:   d.Get("update_send_custom_parameters").(bool),
		UpdateCustomParameters:       d.Get("update_custom_parameters"),
		CloseURL:                     d.Get("close_url").(string),
		CloseMethod:                  d.Get("close_method").(string),
		CloseSendIncidentParameters:  d.Get("close_send_incident_parameters").(bool),
		CloseSendCustomParameters:    d.Get("close_send_custom_parameters").(bool),
		CloseCustomParameters:        d.Get("close_custom_parameters"),
		AlertTagIDs:                  alertTagIDs,
	}

	if _, ok := d.GetOk("custom_parameters"); !webhookIntegration.SendCustomParameters && !ok {
		webhookIntegration.CustomParameters = nil
	}

	if _, ok := d.GetOk("update_custom_parameters"); !webhookIntegration.UpdateSendCustomParameters && !ok {
		webhookIntegration.UpdateCustomParameters = nil
	}

	if _, ok := d.GetOk("close_custom_parameters"); !webhookIntegration.CloseSendCustomParameters && !ok {
		webhookIntegration.CloseCustomParameters = nil
	}

	return webhookIntegration, nil
}

func updateWebhookIntegrationResourceData(d *schema.ResourceData, webhookIntegration *api.WebhookIntegration) {
	d.Set("name", webhookIntegration.Name)
	d.Set("url", webhookIntegration.URL)
	d.Set("timeout", webhookIntegration.Timeout)
	d.Set("method", webhookIntegration.Method)
	d.Set("selection_type", webhookIntegration.SelectionType)
	d.Set("is_poller_webhook", webhookIntegration.IsPollerWebhook)
	d.Set("poller", webhookIntegration.Poller)
	d.Set("send_incident_parameters", webhookIntegration.SendIncidentParameters)
	d.Set("send_custom_parameters", webhookIntegration.SendCustomParameters)
	d.Set("custom_parameters", webhookIntegration.CustomParameters)
	d.Set("send_in_json_format", webhookIntegration.SendInJsonFormat)
	d.Set("auth_method", webhookIntegration.AuthMethod)
	d.Set("username", webhookIntegration.Username)
	d.Set("password", webhookIntegration.Password)
	d.Set("oauth2_provider", webhookIntegration.OauthProvider)
	d.Set("user_agent", webhookIntegration.UserAgent)
	d.Set("monitors", webhookIntegration.Monitors)
	d.Set("alert_tags_id", webhookIntegration.AlertTagIDs)
}
