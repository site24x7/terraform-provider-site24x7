package site24x7

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
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
	"selection_type": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "Resource Type associated with this integration.Monitor Group not supported.",
	},
	"monitors": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "Monitors associated with the integration.",
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
	"alert_tags_id": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "Tag id’s to be associated with the integration.",
	},
}

func resourceSite24x7SlackIntegration() *schema.Resource {
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
	client := meta.(Client)

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
	client := meta.(Client)

	slackIntegration, err := client.SlackIntegration().Get(d.Id())
	if err != nil {
		return err
	}

	updateSlackIntegrationResourceData(d, slackIntegration)

	return nil
}

func slackIntegrationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

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
	client := meta.(Client)

	err := client.ThirdPartyIntegrations().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func slackIntegrationExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(Client)

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

	var alertTagIDs []string
	for _, id := range d.Get("alert_tags_id").([]interface{}) {
		alertTagIDs = append(alertTagIDs, id.(string))
	}

	slackIntegration := &api.SlackIntegration{
		ServiceID:     d.Id(),
		Name:          d.Get("name").(string),
		URL:           d.Get("url").(string),
		SelectionType: api.ResourceType(d.Get("selection_type").(int)),
		Monitors:      monitorsIDs,
		SenderName:    d.Get("sender_name").(string),
		Title:         d.Get("title").(string),
		AlertTagIDs:   alertTagIDs,
	}

	return slackIntegration, nil
}

func updateSlackIntegrationResourceData(d *schema.ResourceData, slackIntegration *api.SlackIntegration) {
	d.Set("name", slackIntegration.Name)
	d.Set("url", slackIntegration.URL)
	d.Set("selection_type", slackIntegration.SelectionType)
	d.Set("monitors", slackIntegration.Monitors)
	d.Set("sender_name", slackIntegration.SenderName)
	d.Set("title", slackIntegration.Title)
	d.Set("alert_tags_id", slackIntegration.AlertTagIDs)
}
