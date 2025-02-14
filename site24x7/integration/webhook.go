package integration

import (
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
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
		Optional:    true,
		Default:     30,
		Description: "The amount of time a connection waits to time out. Default value is 30. Range 1 - 45.",
	},
	"method": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "G",
		Description: "HTTP Method to be used for accessing the website. PUT, PATCH and DELETE are not supported. Default value is 'G'.",
	},
	"selection_type": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     0,
		Description: "Resource Type associated with this integration. Default value is '0'. Can take values 0|2|3. '0' denotes 'All Monitors', '2' denotes 'Monitors', '3' denotes 'Tags'",
	},
	"trouble_alert": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Trouble'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications. Default value is 'true'",
	},
	"critical_alert": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Critical'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.",
	},
	"down_alert": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Setting this to 'true' will send alert notifications to this third-party integration when the monitor status changes to 'Down'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.",
	},
	"is_poller_webhook": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Boolean indicating whether it is an On-Premise Poller based Webhook.",
	},
	"poller": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Mandatory, if is_poller_webhook is set as true. Denotes On-Premise Poller ID.",
	},
	"send_incident_parameters": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Configuration to send incident parameters while executing the action.",
	},
	"send_custom_parameters": {
		Type:        schema.TypeBool,
		Optional:    true,
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
		Description: "Configuration to enable JSON format for post parameters.",
	},
	"auth_method": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Authentication method to access the action url.",
	},
	"user_name": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "User name for authentication.",
	},
	"password": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Password for authentication.",
		DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			// Suppress diff - Password in API response is encrypted.
			return true
		},
	},
	"oauth2_provider": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Provider ID of the OAuth Provider to be associated with the action.",
	},
	"custom_headers": {
		Type:        schema.TypeMap,
		Optional:    true,
		Description: "A Map of Header name and value.",
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
		Description: "Monitors to be associated with the integration when the selection_type = 2.",
	},
	"tags": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "Tags to be associated with the integration when the selection_type = 3.",
	},
	"alert_tags_id": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "Tag id’s to be associated with the integration.",
	},
	"manage_tickets": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Configuration to handle ticketing based integration.",
	},
	"response_format": {
		Type:        schema.TypeString,
		Default:     "J",
		Optional:    true,
		Description: "Expected response type for ticketing based integration: J(ason) or X(ml)",
	},
	"path_expression": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "JSON Path/XPath expression to obtain the ticket id of the request created.",
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
		Optional:    true,
		Description: "Configuration to send incident parameters while updating the ticket.",
	},
	"update_send_custom_parameters": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Configuration to send custom parameters while updating the ticket.",
	},
	"update_custom_parameters": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Configuration to send custom parameters while updating the ticket.",
	},
	"update_send_in_json_format": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Configuration to post in JSON format while updating the ticket.",
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
		Optional:    true,
		Description: "Configuration to send incident parameters while closing the ticket.",
	},
	"close_send_custom_parameters": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Configuration to send custom parameters while closing the ticket.",
	},
	"close_custom_parameters": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Mandatory, When close_send_custom_parameters is set as true.Custom parameters to be passed while accessing the URL.",
	},
	"close_send_in_json_format": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Configuration to post in JSON format while closing the ticket.",
	},
}

func ResourceSite24x7WebhookIntegration() *schema.Resource {
	return &schema.Resource{
		Create: webhookIntegrationCreate,
		Read:   webhookIntegrationRead,
		Update: webhookIntegrationUpdate,
		Delete: webhookIntegrationDelete,
		Exists: webhookIntegrationExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: WebhookIntegrationSchema,
	}
}

func webhookIntegrationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

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
	client := meta.(site24x7.Client)

	webhookIntegration, err := client.WebhookIntegration().Get(d.Id())
	if err != nil {
		return err
	}

	updateWebhookIntegrationResourceData(d, webhookIntegration)

	return nil
}

func webhookIntegrationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

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
	client := meta.(site24x7.Client)

	err := client.ThirdPartyIntegrations().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func webhookIntegrationExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

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

	var tagIDs []string
	for _, id := range d.Get("tags").([]interface{}) {
		tagIDs = append(tagIDs, id.(string))
	}

	var alertTagIDs []string
	for _, id := range d.Get("alert_tags_id").([]interface{}) {
		alertTagIDs = append(alertTagIDs, id.(string))
	}

	// Custom Headers
	customHeaderMap := d.Get("custom_headers").(map[string]interface{})
	keys := make([]string, 0, len(customHeaderMap))
	for k := range customHeaderMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	customHeaders := make([]api.Header, len(keys))
	for i, k := range keys {
		customHeaders[i] = api.Header{Name: k, Value: customHeaderMap[k].(string)}
	}

	webhookIntegration := &api.WebhookIntegration{
		ServiceID:                    d.Id(),
		Name:                         d.Get("name").(string),
		URL:                          d.Get("url").(string),
		Timeout:                      d.Get("timeout").(int),
		Method:                       d.Get("method").(string),
		SelectionType:                api.ResourceType(d.Get("selection_type").(int)),
		TroubleAlert:                 d.Get("trouble_alert").(bool),
		CriticalAlert:                d.Get("critical_alert").(bool),
		DownAlert:                    d.Get("down_alert").(bool),
		IsPollerWebhook:              d.Get("is_poller_webhook").(bool),
		Poller:                       d.Get("poller").(string),
		SendIncidentParameters:       d.Get("send_incident_parameters").(bool),
		SendCustomParameters:         d.Get("send_custom_parameters").(bool),
		CustomParameters:             d.Get("custom_parameters"),
		SendInJsonFormat:             d.Get("send_in_json_format").(bool),
		AuthMethod:                   d.Get("auth_method").(string),
		UserName:                     d.Get("user_name").(string),
		Password:                     d.Get("password").(string),
		OauthProvider:                d.Get("oauth2_provider").(string),
		UserAgent:                    d.Get("user_agent").(string),
		CustomHeaders:                customHeaders,
		Monitors:                     monitorsIDs,
		Tags:                         tagIDs,
		AlertTagIDs:                  alertTagIDs,
		ManageTickets:                d.Get("manage_tickets").(bool),
		ResponseFormat:               d.Get("response_format").(string),
		PathExpression:               d.Get("path_expression").(string),
		UpdateURL:                    d.Get("update_url").(string),
		UpdateMethod:                 d.Get("update_method").(string),
		UpdateSendIncidentParameters: d.Get("update_send_incident_parameters").(bool),
		UpdateSendCustomParameters:   d.Get("update_send_custom_parameters").(bool),
		UpdateCustomParameters:       d.Get("update_custom_parameters"),
		UpdateSendInJsonFormat:       d.Get("update_send_in_json_format").(bool),
		CloseURL:                     d.Get("close_url").(string),
		CloseMethod:                  d.Get("close_method").(string),
		CloseSendIncidentParameters:  d.Get("close_send_incident_parameters").(bool),
		CloseSendCustomParameters:    d.Get("close_send_custom_parameters").(bool),
		CloseCustomParameters:        d.Get("close_custom_parameters"),
		CloseSendInJsonFormat:        d.Get("close_send_in_json_format").(bool),
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
	customHeaders := make(map[string]interface{})
	for _, h := range webhookIntegration.CustomHeaders {
		if h.Name == "" {
			continue
		}
		customHeaders[h.Name] = h.Value
	}
	d.Set("name", webhookIntegration.Name)
	d.Set("url", webhookIntegration.URL)
	d.Set("timeout", webhookIntegration.Timeout)
	d.Set("method", webhookIntegration.Method)
	d.Set("selection_type", webhookIntegration.SelectionType)
	d.Set("trouble_alert", webhookIntegration.TroubleAlert)
	d.Set("critical_alert", webhookIntegration.CriticalAlert)
	d.Set("down_alert", webhookIntegration.DownAlert)
	d.Set("is_poller_webhook", webhookIntegration.IsPollerWebhook)
	d.Set("poller", webhookIntegration.Poller)
	d.Set("send_incident_parameters", webhookIntegration.SendIncidentParameters)
	d.Set("send_custom_parameters", webhookIntegration.SendCustomParameters)
	d.Set("custom_parameters", webhookIntegration.CustomParameters)
	d.Set("send_in_json_format", webhookIntegration.SendInJsonFormat)
	d.Set("auth_method", webhookIntegration.AuthMethod)
	d.Set("user_name", webhookIntegration.UserName)
	d.Set("password", webhookIntegration.Password)
	d.Set("oauth2_provider", webhookIntegration.OauthProvider)
	d.Set("user_agent", webhookIntegration.UserAgent)
	d.Set("custom_headers", customHeaders)
	d.Set("tags", webhookIntegration.Tags)
	d.Set("monitors", webhookIntegration.Monitors)
	d.Set("alert_tags_id", webhookIntegration.AlertTagIDs)
	// Manage tickets configuration
	d.Set("manage_tickets", webhookIntegration.ManageTickets)
	d.Set("response_format", webhookIntegration.ResponseFormat)
	d.Set("path_expression", webhookIntegration.PathExpression)
	// Update Ticket
	d.Set("update_url", webhookIntegration.UpdateURL)
	d.Set("update_method", webhookIntegration.UpdateMethod)
	d.Set("update_send_incident_parameters", webhookIntegration.UpdateSendIncidentParameters)
	d.Set("update_send_custom_parameters", webhookIntegration.UpdateSendCustomParameters)
	d.Set("update_custom_parameters", webhookIntegration.UpdateCustomParameters)
	d.Set("update_send_in_json_format", webhookIntegration.UpdateSendInJsonFormat)
	// Close ticket
	d.Set("close_url", webhookIntegration.CloseURL)
	d.Set("close_method", webhookIntegration.CloseMethod)
	d.Set("close_send_incident_parameters", webhookIntegration.CloseSendIncidentParameters)
	d.Set("close_send_custom_parameters", webhookIntegration.CloseSendCustomParameters)
	d.Set("close_custom_parameters", webhookIntegration.CloseCustomParameters)
	d.Set("close_send_in_json_format", webhookIntegration.CloseSendInJsonFormat)

}
