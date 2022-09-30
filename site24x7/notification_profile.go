package site24x7

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
)

var NotificationProfileSchema = map[string]*schema.Schema{
	"profile_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display Name for the notification profile.",
	},
	"rca_needed": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Configuration to send root cause analysis when the monitor is down. Default is true.",
	},
	"notify_after_executing_actions": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Configuration to raise alerts for downtime only after executing the pre-configured monitor actions. Default is false.",
	},
	"template_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     0,
		Description: "Email template ID for notification.",
	},
	"suppress_automation": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Configuration to stop automation from being executed on the dependent monitors. Default is true.",
	},
	"alert_configuration": {
		Type:        schema.TypeSet,
		Optional:    true,
		Description: "Configuration to alert the user. All alerts will be sent through the notification mode of your preference. You can also configure the business hours and the status for which you would like to receive an alert. If you do not set any specific business hours or status preferences, you'll receive alerts for all the status changes throughout the day.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"status": {
					Type:         schema.TypeInt,
					Optional:     true,
					Default:      -1,
					ValidateFunc: validation.IntInSlice([]int{-1, 0, 1, 2, 3}),
					Description:  "Status for which alerts should be raised. '-1' denotes 'Any', '0' denotes 'Down', '1' denotes 'Up', '2' denotes 'Trouble' and '3' denotes 'Critical'.",
				},
				"business_hours_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Default:     "-1",
					Description: "Alerting Period - Predefined business hours during which alerts should be sent. Default value is '-1' and it denotes 'All Hours'.",
				},
				"notification_medium": {
					Type:     schema.TypeList,
					Required: true,
					Elem: &schema.Schema{
						Type: schema.TypeInt,
					},
					Description: "Medium through which you’d wish to receive the notifications. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call' and '6' denotes 'Mobile push notification'.",
				},
				"outside_business_hours": {
					Type:        schema.TypeString,
					Optional:    true,
					Default:     0,
					Description: "To specify whether the user would receive alerts within or beyond business hours. Default value is '0' and it denotes 'Time within the business_hours_id configured', '1' denotes 'Time outside the business_hours_id configured'.",
				},
			},
		},
	},
	"notification_delay_configuration": {
		Type:        schema.TypeSet,
		Optional:    true,
		Description: "You can choose to delay and receive Down, Trouble, or Critical notifications if the monitor remains in the same state for a specific number of polls. If you haven't configured any Notification Delay for a specific period, you'll receive alerts immediately.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"status": {
					Type:         schema.TypeInt,
					Optional:     true,
					Default:      0,
					ValidateFunc: validation.IntInSlice([]int{0, 2, 3}),
					Description:  "Status for which alerts should be raised. '0' denotes 'Down', '2' denotes 'Trouble' and '3' denotes 'Critical'.",
				},
				"business_hours_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Default:     "-1",
					Description: "Alerting Period - Predefined business hours during which alerts should be sent. Default value is '-1' and it denotes 'All Hours'.",
				},
				"notification_delay": {
					Type:         schema.TypeInt,
					Optional:     true,
					Default:      1,
					ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5}),
					Description:  "Notify based on the downtime delay constants define here - https://www.site24x7.com/help/api/#notification-profile-constants. Default value is '1' and it denotes 'Notify immediately after failure'.",
				},
				"outside_business_hours": {
					Type:        schema.TypeString,
					Optional:    true,
					Default:     0,
					Description: "To specify whether the user would receive alerts within or beyond business hours. Default value is '0' and it denotes 'Time within the business_hours_id configured', '1' denotes 'Time outside the business_hours_id configured'.",
				},
			},
		},
	},
	"persistent_alert_configuration": {
		Type:        schema.TypeSet,
		Optional:    true,
		Description: "Persistent alerts provide continuous notifications until you acknowledge the Down/Critical/Trouble alarm. You will be receiving alerts until you acknowledge the alarms, at the frequency you've configured in the Notify Every Field.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"notify_every": {
					Type:         schema.TypeInt,
					Required:     true,
					ValidateFunc: validation.IntAtMost(60),
					Description:  "Denotes the number of times the error has to be ignored before sending a notification. Value ranges from 0-60.",
				},
				"notification_medium": {
					Type:     schema.TypeList,
					Required: true,
					Elem: &schema.Schema{
						Type: schema.TypeInt,
					},
					Description: "Medium through which you’d wish to receive the notifications. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call' and '6' denotes 'Mobile push notification'.",
				},
				"third_party_services": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Description: "Third-party services through which you’d wish to receive the notification.",
				},
			},
		},
	},
	"escalation_levels": {
		Type:        schema.TypeSet,
		Optional:    true,
		Description: "Configuration to receive persistent notifications after a specific number of errors.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"user_group_id": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "User group ID for downtime escalation.",
				},
				"escalation_wait_time": {
					Type:        schema.TypeInt,
					Required:    true,
					Description: "Mandatory, if any User Alert Group is added for escalation Downtime duration for escalation in mins.",
				},
				"notification_medium": {
					Type:     schema.TypeList,
					Required: true,
					Elem: &schema.Schema{
						Type: schema.TypeInt,
					},
					Description: "Medium through which you’d wish to receive the notifications. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call' and '6' denotes 'Mobile push notification'.",
				},
				"third_party_services": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Description: "Third-party services through which you’d wish to receive the notification.",
				},
			},
		},
	},
	"escalation_automations": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Description: "Execute configured IT automations during an escalation.",
	},
}

func ResourceSite24x7NotificationProfile() *schema.Resource {
	return &schema.Resource{
		Create: notificationProfileCreate,
		Read:   notificationProfileRead,
		Update: notificationProfileUpdate,
		Delete: notificationProfileDelete,
		Exists: notificationProfileExists,

		Schema: NotificationProfileSchema,
	}
}

func notificationProfileCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	notificationProfile := resourceDataToNotificationProfile(d)

	notificationProfile, err := client.NotificationProfiles().Create(notificationProfile)
	if err != nil {
		return err
	}

	d.SetId(notificationProfile.ProfileID)

	return nil
}

func notificationProfileRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	notificationProfile, err := client.NotificationProfiles().Get(d.Id())
	if err != nil {
		return err
	}

	updateNotificationProfileResourceData(d, notificationProfile)

	return nil
}

func notificationProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	notificationProfile := resourceDataToNotificationProfile(d)

	notificationProfile, err := client.NotificationProfiles().Update(notificationProfile)
	if err != nil {
		return err
	}

	d.SetId(notificationProfile.ProfileID)

	return nil
}

func notificationProfileDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	err := client.NotificationProfiles().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func notificationProfileExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(Client)

	_, err := client.NotificationProfiles().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToNotificationProfile(d *schema.ResourceData) *api.NotificationProfile {

	var alertConfiguration []map[string]interface{}
	for _, alertConf := range d.Get("alert_configuration").(*schema.Set).List() {
		if alertConf != nil {
			alertConfiguration = append(alertConfiguration, alertConf.(map[string]interface{}))
		}
	}

	var notificationDelayConfiguration []map[string]interface{}
	for _, notifationDelayConf := range d.Get("notification_delay_configuration").(*schema.Set).List() {
		if notifationDelayConf != nil {
			notificationDelayConfiguration = append(notificationDelayConfiguration, notifationDelayConf.(map[string]interface{}))
		}
	}

	var persistentAlertConfiguration []map[string]interface{}
	for _, persistentAlertConf := range d.Get("persistent_alert_configuration").(*schema.Set).List() {
		if persistentAlertConf != nil {
			persistentAlertConfiguration = append(persistentAlertConfiguration, persistentAlertConf.(map[string]interface{}))
		}
	}

	// Escalation Configuration
	var escalationLevels []map[string]interface{}
	for _, escalationlevelConf := range d.Get("escalation_levels").(*schema.Set).List() {
		if escalationlevelConf != nil {
			escalationLevels = append(escalationLevels, escalationlevelConf.(map[string]interface{}))
		}
	}

	var escalationAutomations []string
	for _, automationID := range d.Get("escalation_automations").(*schema.Set).List() {
		escalationAutomations = append(escalationAutomations, automationID.(string))
	}

	escalationConfiguration := make(map[string]interface{})
	if escalationLevels != nil {
		escalationConfiguration["escalation_levels"] = escalationLevels
	}
	if escalationAutomations != nil {
		escalationConfiguration["escalation_automations"] = escalationAutomations
	}

	notifProfile := &api.NotificationProfile{
		ProfileID:                      d.Id(),
		ProfileName:                    d.Get("profile_name").(string),
		RcaNeeded:                      d.Get("rca_needed").(bool),
		NotifyAfterExecutingActions:    d.Get("notify_after_executing_actions").(bool),
		TemplateID:                     d.Get("template_id").(string),
		SuppressAutomation:             d.Get("suppress_automation").(bool),
		AlertConfiguration:             alertConfiguration,
		NotificationDelayConfiguration: notificationDelayConfiguration,
		PersistentAlertConfiguration:   persistentAlertConfiguration,
	}

	if len(escalationConfiguration) != 0 {
		notifProfile.EscalationConfiguration = escalationConfiguration
	}

	return notifProfile
}

// Called during read and sets notificationProfile in API response to ResourceData
func updateNotificationProfileResourceData(d *schema.ResourceData, notificationProfile *api.NotificationProfile) {
	d.Set("profile_name", notificationProfile.ProfileName)
	d.Set("rca_needed", notificationProfile.RcaNeeded)
	d.Set("notify_after_executing_actions", notificationProfile.NotifyAfterExecutingActions)
	d.Set("template_id", notificationProfile.TemplateID)
	d.Set("suppress_automation", notificationProfile.SuppressAutomation)
	d.Set("alert_configuration", notificationProfile.AlertConfiguration)
	d.Set("notification_delay_configuration", notificationProfile.NotificationDelayConfiguration)
	d.Set("persistent_alert_configuration", notificationProfile.PersistentAlertConfiguration)
	if notificationProfile.EscalationConfiguration != nil {
		d.Set("escalation_automations", notificationProfile.EscalationConfiguration["escalation_automations"])
		d.Set("escalation_levels", notificationProfile.EscalationConfiguration["escalation_levels"])
	}

}
