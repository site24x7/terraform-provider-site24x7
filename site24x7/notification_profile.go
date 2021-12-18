package site24x7

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
)

// SAMPLE POST JSON

// {
// 	"template_id": "123456000024578001",
// 	"notify_after_executing_actions": false,
// 	"downtime_notification_delay": 1,
// 	"persistent_notification": 1,
// 	"escalation_wait_time": "30",
// 	"rca_needed": true,
// 	"suppress_automation": true,
// 	"profile_name": "Test Notification Profile",
// 	"escalation_user_group_id": "123456000000025005",
// 	"escalation_automations": [
// 	  "123456000000047001"
// 	],
// 	"escalation_services": [
// 	  "123456000008777001"
// 	]
// }

var NotificationProfileSchema = map[string]*schema.Schema{
	"profile_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display Name for the notification profile",
	},
	"rca_needed": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Settings to send root cause analysis when monitor goes down.",
	},
	"notify_after_executing_actions": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Description: "Settings to downtime only after executing configured monitor actions.",
	},
	"escalation_wait_time": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Duration of Downtime before Escalation. Mandatory if any user group is added for escalation.",
	},
	"downtime_notification_delay": {
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      1,
		ValidateFunc: validation.IntInSlice([]int{1, 2, 3, 4, 5}),
		Description:  "Configuration for delayed notification. Default value is 1.",
	},
	"persistent_notification": {
		Type:         schema.TypeInt,
		Optional:     true,
		ValidateFunc: validation.IntInSlice([]int{-1, 1, 2, 3, 4, 5}),
		Description:  "Settings to receive persistent notification after number of errors.",
	},
	"escalation_user_group_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "User group ID for downtime escalation.",
	},
	"template_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     0,
		Description: "Email template ID for notification",
	},
	"suppress_automation": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Settings to stop an automation being executed on the dependent monitors.",
	},
	"escalation_automations": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Description: "Execute configured IT automations during an escalation.",
	},
	"escalation_services": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Description: "Invoke and manage escalations in your preferred third party services.",
	},
}

func resourceSite24x7NotificationProfile() *schema.Resource {
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

	var escalationAutomations []string
	for _, automationID := range d.Get("escalation_automations").([]interface{}) {
		escalationAutomations = append(escalationAutomations, automationID.(string))
	}

	var escalationServices []string
	for _, thirdPartyServiceID := range d.Get("escalation_services").([]interface{}) {
		escalationServices = append(escalationServices, thirdPartyServiceID.(string))
	}

	return &api.NotificationProfile{
		ProfileID:                   d.Id(),
		ProfileName:                 d.Get("profile_name").(string),
		RcaNeeded:                   d.Get("rca_needed").(bool),
		NotifyAfterExecutingActions: d.Get("notify_after_executing_actions").(bool),
		DowntimeNotificationDelay:   d.Get("downtime_notification_delay").(int),
		PersistentNotification:      d.Get("persistent_notification").(int),
		EscalationUserGroupId:       d.Get("escalation_user_group_id").(string),
		EscalationWaitTime:          d.Get("escalation_wait_time").(int),
		SuppressAutomation:          d.Get("suppress_automation").(bool),
		EscalationAutomations:       escalationAutomations,
		EscalationServices:          escalationServices,
		TemplateID:                  d.Get("template_id").(string),
	}
}

// Called during read and sets notificationProfile in API response to ResourceData
func updateNotificationProfileResourceData(d *schema.ResourceData, notificationProfile *api.NotificationProfile) {
	d.Set("profile_name", notificationProfile.ProfileName)
	d.Set("rca_needed", notificationProfile.RcaNeeded)
	d.Set("notify_after_executing_actions", notificationProfile.NotifyAfterExecutingActions)
	d.Set("downtime_notification_delay", notificationProfile.DowntimeNotificationDelay)
	d.Set("persistent_notification", notificationProfile.PersistentNotification)
	d.Set("escalation_user_group_id", notificationProfile.EscalationUserGroupId)
	d.Set("escalation_wait_time", notificationProfile.EscalationWaitTime)
	d.Set("escalation_user_group_id", notificationProfile.EscalationUserGroupId)
	d.Set("suppress_automation", notificationProfile.SuppressAutomation)
	d.Set("escalation_automations", notificationProfile.EscalationAutomations)
	d.Set("escalation_services", notificationProfile.EscalationServices)
	d.Set("template_id", notificationProfile.TemplateID)
}
