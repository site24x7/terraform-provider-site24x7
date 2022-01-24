package integration

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var SlackIntegrationSchema = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name for the Slack Integration.",
	},
	"url": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Hook URL to which the message will be posted.",
	},
	"sender_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Name of the service who posted the message.",
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

func ResourceSite24x7SlackIntegration() *schema.Resource {
	return &schema.Resource{
		Create: slackIntegrationCreate,
		Read:   slackIntegrationRead,
		Update: slackIntegrationUpdate,
		Delete: slackIntegrationDelete,
		Exists: slackIntegrationExists,

		Schema: SlackIntegrationSchema,
	}
}

func slackIntegrationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	slackIntegration, err := resourceDataToSlackIntegration(d)
	if err != nil {
		return err
	}

	slackIntegration, err = client.SlackIntegration().Create(slackIntegration)
	if err != nil {
		return err
	}

	d.SetId(slackIntegration.ServiceID)

	return nil
}

func slackIntegrationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	slackIntegration, err := client.SlackIntegration().Get(d.Id())
	if err != nil {
		return err
	}

	updateSlackIntegrationResourceData(d, slackIntegration)

	return nil
}

func slackIntegrationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	slackIntegration, err := resourceDataToSlackIntegration(d)
	if err != nil {
		return err
	}

	slackIntegration, err = client.SlackIntegration().Update(slackIntegration)
	if err != nil {
		return err
	}

	d.SetId(slackIntegration.ServiceID)

	return nil
}

func slackIntegrationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.ThirdPartyIntegrations().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func slackIntegrationExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.SlackIntegration().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToSlackIntegration(d *schema.ResourceData) (*api.SlackIntegration, error) {

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

	slackIntegration := &api.SlackIntegration{
		ServiceID:     d.Id(),
		Name:          d.Get("name").(string),
		URL:           d.Get("url").(string),
		SenderName:    d.Get("sender_name").(string),
		Title:         d.Get("title").(string),
		SelectionType: api.ResourceType(d.Get("selection_type").(int)),
		TroubleAlert:  d.Get("trouble_alert").(bool),
		CriticalAlert: d.Get("critical_alert").(bool),
		DownAlert:     d.Get("down_alert").(bool),
		Monitors:      monitorsIDs,
		Tags:          tagIDs,
		AlertTagIDs:   alertTagIDs,
	}

	return slackIntegration, nil
}

func updateSlackIntegrationResourceData(d *schema.ResourceData, slackIntegration *api.SlackIntegration) {
	d.Set("name", slackIntegration.Name)
	d.Set("url", slackIntegration.URL)
	d.Set("sender_name", slackIntegration.SenderName)
	d.Set("title", slackIntegration.Title)
	d.Set("selection_type", slackIntegration.SelectionType)
	d.Set("trouble_alert", slackIntegration.TroubleAlert)
	d.Set("critical_alert", slackIntegration.CriticalAlert)
	d.Set("down_alert", slackIntegration.DownAlert)
	d.Set("tags", slackIntegration.Tags)
	d.Set("monitors", slackIntegration.Monitors)
	d.Set("alert_tags_id", slackIntegration.AlertTagIDs)
}
