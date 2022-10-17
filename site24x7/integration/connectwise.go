package integration

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var ConnectwiseIntegrationSchema = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name for the ConnectWise Integration.",
	},
	"url": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "URL to be invoked for action execution.",
	},
	"company": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Company for Authentication.",
	},
	"public_key": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Public Key for Authentication",
	},
	"private_key": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Private Key for Authentication.",
	},
	"company_id": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Tickets to your ConnectWise account will be assigned to this Company ID.",
	},
	"close_status": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "The configuration settings to resolve or close incidents automatically in Connectwise, when the monitor status changes to UP.",
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
	"user_groups": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "User Alert Group to be notified when there is an error in ConnectWise Manage Integration.",
	},
}


func ResourceSite24x7ConnectwiseIntegration() *schema.Resource {
	return &schema.Resource{
		Create: connectwiseIntegrationCreate,
		Read:   connectwiseIntegrationRead,
		Update: connectwiseIntegrationUpdate,
		Delete: connectwiseIntegrationDelete,
		Exists: connectwiseIntegrationExists,

		Schema: ConnectwiseIntegrationSchema,
	}
}

func connectwiseIntegrationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	connectwiseIntegration, err := resourceDataToConnectwiseIntegration(d)
	if err != nil {
		return err
	}

	connectwiseIntegration, err = client.ConnectwiseIntegration().Create(connectwiseIntegration)
	if err != nil {
		return err
	}

	d.SetId(connectwiseIntegration.ServiceID)

	return nil
}

func connectwiseIntegrationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	connectwiseIntegration, err := client.ConnectwiseIntegration().Get(d.Id())
	if err != nil {
		return err
	}

	updateConnectwiseIntegrationResourceData(d, connectwiseIntegration)

	return nil
}

func connectwiseIntegrationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	connectwiseIntegration, err := resourceDataToConnectwiseIntegration(d)
	if err != nil {
		return err
	}

	connectwiseIntegration, err = client.ConnectwiseIntegration().Update(connectwiseIntegration)
	if err != nil {
		return err
	}

	d.SetId(connectwiseIntegration.ServiceID)

	return nil
}

func connectwiseIntegrationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.ThirdPartyIntegrations().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func connectwiseIntegrationExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.ConnectwiseIntegration().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToConnectwiseIntegration(d *schema.ResourceData) (*api.ConnectwiseIntegration, error) {

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
	var userGroups []string
	for _, id := range d.Get("user_groups").([]interface{}) {
		userGroups = append(userGroups, id.(string))
	}

	connectwiseIntegration := &api.ConnectwiseIntegration{
		ServiceID:            d.Id(),
		Name:                 d.Get("name").(string),
		URL:           		  d.Get("url").(string),
		Company:              d.Get("company").(string),
		PublicKey:            d.Get("public_key").(string),
		PrivateKey:           d.Get("private_key").(string),
		CompanyId:            d.Get("company_id").(string),
		CloseStatus:          d.Get("close_status").(string),
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
		UserGroups:			  userGroups,
	}

	return connectwiseIntegration, nil
}

func updateConnectwiseIntegrationResourceData(d *schema.ResourceData, connectwiseIntegration *api.ConnectwiseIntegration) {
	d.Set("name", connectwiseIntegration.Name)
	d.Set("url", connectwiseIntegration.URL)
	d.Set("company", connectwiseIntegration.Company)
	d.Set("public_key", connectwiseIntegration.PublicKey)
	d.Set("private_key", connectwiseIntegration.PrivateKey)
	d.Set("company_id", connectwiseIntegration.CompanyId)
	d.Set("close_status", connectwiseIntegration.CloseStatus)
	d.Set("selection_type", connectwiseIntegration.SelectionType)
	d.Set("trouble_alert", connectwiseIntegration.TroubleAlert)
	d.Set("critical_alert", connectwiseIntegration.CriticalAlert)
	d.Set("down_alert", connectwiseIntegration.DownAlert)
	d.Set("manual_resolve", connectwiseIntegration.ManualResolve)
	d.Set("send_custom_parameters", connectwiseIntegration.SendCustomParameters)
	d.Set("custom_parameters", connectwiseIntegration.CustomParameters)
	d.Set("tags", connectwiseIntegration.Tags)
	d.Set("monitors", connectwiseIntegration.Monitors)
	d.Set("alert_tags_id", connectwiseIntegration.AlertTagIDs)
	d.Set("user_groups", connectwiseIntegration.UserGroups)
}
