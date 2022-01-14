package site24x7

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
)

var PagerDutyIntegrationSchema = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name for the PagerDuty Integration.",
	},
	"service_key": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Unique integration key provided by PagerDuty to facilitate incident creation in PagerDuty.",
	},
	"selection_type": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     0,
		Description: "Resource Type associated with this integration. Can take values 0|2|3. '0' denotes 'All Monitors', '2' denotes 'Monitors', '3' denotes 'Tags'",
	},
	"sender_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Name of the service who posted the incident.",
	},
	"title": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Title of the incident.",
	},
	"trouble_alert": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Setting this to 'true' will send alert notifications through this third-party integration when the monitor status changes to 'Trouble'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.",
	},
	"critical_alert": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Setting this to 'true' will send alert notifications through this third-party integration when the monitor status changes to 'Critical'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.",
	},
	"down_alert": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Setting this to 'true' will send alert notifications through this third-party integration when the monitor status changes to 'Down'. One among trouble_alert|critical_alert|down_alert should be set to true for receiving notifications.",
	},
	"manual_resolve": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Configuration to resolve the incidents manually when the monitor changes to UP status.",
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

func resourceSite24x7PagerDutyIntegration() *schema.Resource {
	return &schema.Resource{
		Create: pagerDutyIntegrationCreate,
		Read:   pagerDutyIntegrationRead,
		Update: pagerDutyIntegrationUpdate,
		Delete: pagerDutyIntegrationDelete,
		Exists: pagerDutyIntegrationExists,

		Schema: PagerDutyIntegrationSchema,
	}
}

func pagerDutyIntegrationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	pagerDutyIntegration, err := resourceDataToPagerDutyIntegration(d)
	if err != nil {
		return err
	}

	pagerDutyIntegration, err = client.PagerDutyIntegration().Create(pagerDutyIntegration)
	if err != nil {
		return err
	}

	d.SetId(pagerDutyIntegration.ServiceID)

	return nil
}

func pagerDutyIntegrationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	pagerDutyIntegration, err := client.PagerDutyIntegration().Get(d.Id())
	if err != nil {
		return err
	}

	updatePagerDutyIntegrationResourceData(d, pagerDutyIntegration)

	return nil
}

func pagerDutyIntegrationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	pagerDutyIntegration, err := resourceDataToPagerDutyIntegration(d)
	if err != nil {
		return err
	}

	pagerDutyIntegration, err = client.PagerDutyIntegration().Update(pagerDutyIntegration)
	if err != nil {
		return err
	}

	d.SetId(pagerDutyIntegration.ServiceID)

	return nil
}

func pagerDutyIntegrationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	err := client.ThirdPartyIntegrations().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func pagerDutyIntegrationExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(Client)

	_, err := client.PagerDutyIntegration().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToPagerDutyIntegration(d *schema.ResourceData) (*api.PagerDutyIntegration, error) {

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

	pagerDutyIntegration := &api.PagerDutyIntegration{
		ServiceID:            d.Id(),
		Name:                 d.Get("name").(string),
		ServiceKey:           d.Get("service_key").(string),
		SelectionType:        api.ResourceType(d.Get("selection_type").(int)),
		SenderName:           d.Get("sender_name").(string),
		Title:                d.Get("title").(string),
		TroubleAlert:         d.Get("trouble_alert").(bool),
		CriticalAlert:        d.Get("critical_alert").(bool),
		DownAlert:            d.Get("down_alert").(bool),
		ManualResolve:        d.Get("manual_resolve").(bool),
		SendCustomParameters: d.Get("send_custom_parameters").(bool),
		CustomParameters:     d.Get("custom_parameters").(string),
		Monitors:             monitorsIDs,
		Tags:                 tagIDs,
		AlertTagIDs:          alertTagIDs,
	}

	return pagerDutyIntegration, nil
}

func updatePagerDutyIntegrationResourceData(d *schema.ResourceData, pagerDutyIntegration *api.PagerDutyIntegration) {
	d.Set("name", pagerDutyIntegration.Name)
	d.Set("service_key", pagerDutyIntegration.ServiceKey)
	d.Set("selection_type", pagerDutyIntegration.SelectionType)
	d.Set("sender_name", pagerDutyIntegration.SenderName)
	d.Set("title", pagerDutyIntegration.Title)
	d.Set("trouble_alert", pagerDutyIntegration.TroubleAlert)
	d.Set("critical_alert", pagerDutyIntegration.CriticalAlert)
	d.Set("down_alert", pagerDutyIntegration.DownAlert)
	d.Set("manual_resolve", pagerDutyIntegration.ManualResolve)
	d.Set("send_custom_parameters", pagerDutyIntegration.SendCustomParameters)
	d.Set("custom_parameters", pagerDutyIntegration.CustomParameters)
	d.Set("tags", pagerDutyIntegration.Tags)
	d.Set("monitors", pagerDutyIntegration.Monitors)
	d.Set("alert_tags_id", pagerDutyIntegration.AlertTagIDs)
}
