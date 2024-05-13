package integration

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var TelegramIntegrationSchema = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name for the Telegram Integration.",
	},
	"channel_url": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Hook URL to which the message will be posted.",
	},
	"token": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Hook URL to which the message will be posted.",
	},

	"title": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Title of the incident.",
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

func ResourceSite24x7TelegramIntegration() *schema.Resource {
	return &schema.Resource{
		Create: telegramIntegrationCreate,
		Read:   telegramIntegrationRead,
		Update: telegramIntegrationUpdate,
		Delete: telegramIntegrationDelete,
		Exists: telegramIntegrationExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: TelegramIntegrationSchema,
	}
}

func telegramIntegrationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	telegramIntegration, err := resourceDataToTelegramIntegration(d)
	if err != nil {
		return err
	}

	telegramIntegration, err = client.TelegramIntegration().Create(telegramIntegration)
	if err != nil {
		return err
	}

	d.SetId(telegramIntegration.ServiceID)

	return nil
}

func telegramIntegrationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	telegramIntegration, err := client.TelegramIntegration().Get(d.Id())
	if err != nil {
		return err
	}

	updateTelegramIntegrationResourceData(d, telegramIntegration)

	return nil
}

func telegramIntegrationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	telegramIntegration, err := resourceDataToTelegramIntegration(d)
	if err != nil {
		return err
	}

	telegramIntegration, err = client.TelegramIntegration().Update(telegramIntegration)
	if err != nil {
		return err
	}

	d.SetId(telegramIntegration.ServiceID)

	return nil
}

func telegramIntegrationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.ThirdPartyIntegrations().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func telegramIntegrationExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.TelegramIntegration().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToTelegramIntegration(d *schema.ResourceData) (*api.TelegramIntegration, error) {

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

	telegramIntegration := &api.TelegramIntegration{
		ServiceID:     d.Id(),
		Name:          d.Get("name").(string),
		URL:           d.Get("channel_url").(string),
		BotToken:      d.Get("token").(string),
		Title:         d.Get("title").(string),
		SelectionType: api.ResourceType(d.Get("selection_type").(int)),
		TroubleAlert:  d.Get("trouble_alert").(bool),
		CriticalAlert: d.Get("critical_alert").(bool),
		DownAlert:     d.Get("down_alert").(bool),
		Monitors:      monitorsIDs,
		Tags:          tagIDs,
		AlertTagIDs:   alertTagIDs,
	}

	return telegramIntegration, nil
}

func updateTelegramIntegrationResourceData(d *schema.ResourceData, telegramIntegration *api.TelegramIntegration) {
	d.Set("name", telegramIntegration.Name)
	d.Set("channel_url", telegramIntegration.URL)
	d.Set("token", telegramIntegration.BotToken)
	d.Set("title", telegramIntegration.Title)
	d.Set("selection_type", telegramIntegration.SelectionType)
	d.Set("trouble_alert", telegramIntegration.TroubleAlert)
	d.Set("critical_alert", telegramIntegration.CriticalAlert)
	d.Set("down_alert", telegramIntegration.DownAlert)
	d.Set("tags", telegramIntegration.Tags)
	d.Set("monitors", telegramIntegration.Monitors)
	d.Set("alert_tags_id", telegramIntegration.AlertTagIDs)
}
