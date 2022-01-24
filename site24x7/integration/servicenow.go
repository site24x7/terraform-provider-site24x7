package integration

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var serviceNowIntegrationSchema = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name for the ServiceNow Integration.",
	},
	"instance_url": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "ServiceNow instance URL.",
	},
	"title": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Title of the incident.",
	},
	"sender_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Name of the service who posted the incident.",
	},
	"user_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "User name for authentication.",
	},
	"password": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Password for authentication.",
		DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			// As password in API response is encrypted we are suppressing diff.
			return true
		},
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
	"send_custom_parameters": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Configuration to send custom parameters while executing the action.",
	},
	"custom_parameters": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Mandatory, if send_custom_parameters is set as true. Custom parameters to be passed while accessing the URL.",
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
}

func ResourceSite24x7ServiceNowIntegration() *schema.Resource {
	return &schema.Resource{
		Create: serviceNowIntegrationCreate,
		Read:   serviceNowIntegrationRead,
		Update: serviceNowIntegrationUpdate,
		Delete: serviceNowIntegrationDelete,
		Exists: serviceNowIntegrationExists,

		Schema: serviceNowIntegrationSchema,
	}
}

func serviceNowIntegrationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	serviceNowIntegration, err := resourceDataToServiceNowIntegration(d)
	if err != nil {
		return err
	}

	serviceNowIntegration, err = client.ServiceNowIntegration().Create(serviceNowIntegration)
	if err != nil {
		return err
	}

	d.SetId(serviceNowIntegration.ServiceID)

	return nil
}

func serviceNowIntegrationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	serviceNowIntegration, err := client.ServiceNowIntegration().Get(d.Id())
	if err != nil {
		return err
	}

	updateServiceNowIntegrationResourceData(d, serviceNowIntegration)

	return nil
}

func serviceNowIntegrationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	serviceNowIntegration, err := resourceDataToServiceNowIntegration(d)
	if err != nil {
		return err
	}

	serviceNowIntegration, err = client.ServiceNowIntegration().Update(serviceNowIntegration)
	if err != nil {
		return err
	}

	d.SetId(serviceNowIntegration.ServiceID)

	return nil
}

func serviceNowIntegrationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.ThirdPartyIntegrations().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func serviceNowIntegrationExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.ServiceNowIntegration().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToServiceNowIntegration(d *schema.ResourceData) (*api.ServiceNowIntegration, error) {

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

	serviceNowIntegration := &api.ServiceNowIntegration{
		ServiceID:            d.Id(),
		Name:                 d.Get("name").(string),
		InstanceURL:          d.Get("instance_url").(string),
		SenderName:           d.Get("sender_name").(string),
		Title:                d.Get("title").(string),
		UserName:             d.Get("user_name").(string),
		Password:             d.Get("password").(string),
		SelectionType:        api.ResourceType(d.Get("selection_type").(int)),
		TroubleAlert:         d.Get("trouble_alert").(bool),
		CriticalAlert:        d.Get("critical_alert").(bool),
		DownAlert:            d.Get("down_alert").(bool),
		SendCustomParameters: d.Get("send_custom_parameters").(bool),
		CustomParameters:     d.Get("custom_parameters").(string),
		Monitors:             monitorsIDs,
		Tags:                 tagIDs,
		AlertTagIDs:          alertTagIDs,
	}

	return serviceNowIntegration, nil
}

func updateServiceNowIntegrationResourceData(d *schema.ResourceData, serviceNowIntegration *api.ServiceNowIntegration) {
	d.Set("name", serviceNowIntegration.Name)
	d.Set("instance_url", serviceNowIntegration.InstanceURL)
	d.Set("sender_name", serviceNowIntegration.SenderName)
	d.Set("title", serviceNowIntegration.Title)
	d.Set("user_name", serviceNowIntegration.UserName)
	d.Set("password", serviceNowIntegration.Password)
	d.Set("selection_type", serviceNowIntegration.SelectionType)
	d.Set("trouble_alert", serviceNowIntegration.TroubleAlert)
	d.Set("critical_alert", serviceNowIntegration.CriticalAlert)
	d.Set("down_alert", serviceNowIntegration.DownAlert)
	d.Set("send_custom_parameters", serviceNowIntegration.SendCustomParameters)
	d.Set("custom_parameters", serviceNowIntegration.CustomParameters)
	d.Set("tags", serviceNowIntegration.Tags)
	d.Set("monitors", serviceNowIntegration.Monitors)
	d.Set("alert_tags_id", serviceNowIntegration.AlertTagIDs)
}
