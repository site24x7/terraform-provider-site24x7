package integration

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var OpsgenieIntegrationSchema = map[string]*schema.Schema{
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
	"tags": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "Tags to be associated with the integration when the selection_type = 3.",
	},
	"monitors": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "Monitors to be associated with the integration when the selection_type = 2.",
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

func ResourceSite24x7OpsgenieIntegration() *schema.Resource {
	return &schema.Resource{
		Create: opsgenieIntegrationCreate,
		Read:   opsgenieIntegrationRead,
		Update: opsgenieIntegrationUpdate,
		Delete: opsgenieIntegrationDelete,
		Exists: opsgenieIntegrationExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: OpsgenieIntegrationSchema,
	}
}

func opsgenieIntegrationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	opsgenieIntegration, err := resourceDataToOpsgenieIntegration(d)
	if err != nil {
		return err
	}

	opsgenieIntegration, err = client.OpsgenieIntegration().Create(opsgenieIntegration)
	if err != nil {
		return err
	}

	d.SetId(opsgenieIntegration.ServiceID)

	return nil
}

func opsgenieIntegrationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	opsgenieIntegration, err := client.OpsgenieIntegration().Get(d.Id())
	if err != nil {
		return err
	}

	updateOpsgenieIntegrationResourceData(d, opsgenieIntegration)

	return nil
}

func opsgenieIntegrationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	opsgenieIntegration, err := resourceDataToOpsgenieIntegration(d)
	if err != nil {
		return err
	}

	opsgenieIntegration, err = client.OpsgenieIntegration().Update(opsgenieIntegration)
	if err != nil {
		return err
	}

	d.SetId(opsgenieIntegration.ServiceID)

	return nil
}

func opsgenieIntegrationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.ThirdPartyIntegrations().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func opsgenieIntegrationExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.OpsgenieIntegration().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToOpsgenieIntegration(d *schema.ResourceData) (*api.OpsgenieIntegration, error) {

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

	opsgenieIntegration := &api.OpsgenieIntegration{
		ServiceID:            d.Id(),
		Name:                 d.Get("name").(string),
		URL:                  d.Get("url").(string),
		SelectionType:        api.ResourceType(d.Get("selection_type").(int)),
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

	return opsgenieIntegration, nil
}

func updateOpsgenieIntegrationResourceData(d *schema.ResourceData, opsgenieIntegration *api.OpsgenieIntegration) {
	d.Set("name", opsgenieIntegration.Name)
	d.Set("url", opsgenieIntegration.URL)
	d.Set("selection_type", opsgenieIntegration.SelectionType)
	d.Set("trouble_alert", opsgenieIntegration.TroubleAlert)
	d.Set("critical_alert", opsgenieIntegration.CriticalAlert)
	d.Set("down_alert", opsgenieIntegration.DownAlert)
	d.Set("manual_resolve", opsgenieIntegration.ManualResolve)
	d.Set("send_custom_parameters", opsgenieIntegration.SendCustomParameters)
	d.Set("custom_parameters", opsgenieIntegration.CustomParameters)
	d.Set("tags", opsgenieIntegration.Tags)
	d.Set("monitors", opsgenieIntegration.Monitors)
	d.Set("alert_tags_id", opsgenieIntegration.AlertTagIDs)
}
