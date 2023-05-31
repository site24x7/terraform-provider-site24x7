package site24x7

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
)

// SAMPLE POST JSON
// {
// 	"is_edit_allowed": true,
// 	"is_client_portal_user": false,
// 	"is_account_contact": false,
// 	"is_invited": true,
// 	"subscribe_newsletter": false,
// 	"alert_settings": {
// 	  "email_format": 1,
// 	  "anomaly": [],
// 	  "critical": [
// 		1
// 	  ],
// 	  "applogs": [
// 		1
// 	  ],
// 	  "trouble": [
// 		1
// 	  ],
// 	  "up": [
// 		1
// 	  ],
// 	  "alerting_period": {
// 		"start_time": "00:00",
// 		"end_time": "00:00"
// 	  },
// 	  "down": [
// 		1
// 	  ]
// 	},
// 	"display_name": "User - Terraform",
// 	"is_contact": false,
// 	"selection_type": 1,
// 	"user_role": 10,
// 	"email_address": "jim@example1.com",
// 	"mobile_settings": {
// 	  "country_code": 93,
// 	  "is_confirmed": false,
// 	  "call_provider_id": 0,
// 	  "sms_provider_id": 2,
// 	  "mobile_number": 434388234
// 	},
// 	"image_present": false,
// 	"notify_medium": [
// 	  1
// 	],
// 	"user_groups": [
// 	  "306947000000025005",
// 	  "306947000000025009",
// 	  "306947000000025007"
// 	],
// 	"monitor_groups": [
// 	  "306947000021059031",
// 	  "306947000033224882"
// 	]
// }

var UserSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Name of the User.",
	},
	"email_address": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Email address of the user. Email verification has to be done manually.",
	},
	"user_role": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "Role assigned to the user for accessing Site24x7. Role will be updated only after the user accepts the invitation. Refer https://www.site24x7.com/help/api/#site24x7_user_constants",
	},
	"job_title": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Provide your job title to be added in Site24x7. Refer https://www.site24x7.com/help/api/#job_title",
	},
	"selection_type": {
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      0,
		ValidateFunc: validation.IntInSlice([]int{0, 1}),
		Description:  "Resource type associated to this user. Default value is '0'. Can take values 0|1. '0' denotes 'All Monitors', '1' denotes 'Monitor Group'. 'monitor_groups' attribute is mandatory when the 'selection_type' is '1'.",
	},
	"monitor_groups": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Computed:    true,
		Description: "List of monitor groups to which the user has access to. 'monitor_groups' attribute is mandatory when the 'selection_type' is '1'",
	},
	"notification_medium": {
		Type:     schema.TypeSet,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeInt,
		},
		Description: "Medium through which you’d wish to receive the notifications. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.",
	},
	"user_group_ids": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Computed:    true,
		Description: "List of groups to be associated for the user for receiving alerts.",
	},
	"mobile_settings": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"country_code": {
					Type:     schema.TypeInt,
					Required: true,
				},
				"mobile_number": {
					Type:     schema.TypeInt,
					Required: true,
				},
				"is_confirmed": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
				"call_provider_id": {
					Type:     schema.TypeInt,
					Optional: true,
					Default:  0,
				},
				"sms_provider_id": {
					Type:     schema.TypeInt,
					Optional: true,
					Default:  2,
				},
			},
		},
		Description: "Phone number configurations to receive alerts.",
	},
	// Alert Settings
	"down_notification_medium": {
		Type:     schema.TypeSet,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeInt,
		},
		Description: "Medium through which you’d wish to receive the Down alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.",
	},
	"critical_notification_medium": {
		Type:     schema.TypeSet,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeInt,
		},
		Description: "Medium through which you’d wish to receive the Critical alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.",
	},
	"trouble_notification_medium": {
		Type:     schema.TypeSet,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeInt,
		},
		Description: "Medium through which you’d wish to receive the Trouble alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.",
	},
	"up_notification_medium": {
		Type:     schema.TypeSet,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeInt,
		},
		Description: "Medium through which you’d wish to receive the Up alerts. Default value is 1. '1' denotes 'Email', '2' denotes 'SMS', '3' denotes 'Voice Call'.",
	},
	"alerting_period_start_time": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "00:00",
		Description: "Define a time window so you can receive Voice/SMS status alerts during this period alone. You can't define this window for email or IM based notifications.",
	},
	"alerting_period_end_time": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "00:00",
		Description: "Define a time window so you can receive Voice/SMS status alerts during this period alone. You can't define this window for email or IM based notifications.",
	},
	"email_format": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     1,
		Description: "Denotes the email format. '0' - Text, '1' - HTML",
	},
	"consent_for_non_eu_alerts": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Description: "The third-party providers we use to send SMS and voice alerts will process the data outside the EU region.",
	},
}

func ResourceSite24x7User() *schema.Resource {
	return &schema.Resource{
		Create: userCreate,
		Read:   userRead,
		Update: userUpdate,
		Delete: userDelete,
		Exists: userExists,

		Schema: UserSchema,
	}
}

func userCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	user := resourceDataToUser(d)

	user, err := client.Users().Create(user)
	if err != nil {
		return err
	}

	d.SetId(user.ID)

	return nil
}

func userRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	user, err := client.Users().Get(d.Id())
	if err != nil {
		return err
	}

	updateUserResourceData(d, user)

	return nil
}

func userUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	user := resourceDataToUser(d)

	user, err := client.Users().Update(user)
	if err != nil {
		return err
	}

	d.SetId(user.ID)

	return nil
}

func userDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	err := client.Users().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func userExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(Client)

	_, err := client.Users().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToUser(d *schema.ResourceData) *api.User {
	var userGroupIDs []string
	for _, id := range d.Get("user_group_ids").(*schema.Set).List() {
		if id != nil {
			userGroupIDs = append(userGroupIDs, id.(string))
		}
	}

	var monitorGroupIDs []string
	for _, id := range d.Get("monitor_groups").(*schema.Set).List() {
		if id != nil {
			monitorGroupIDs = append(monitorGroupIDs, id.(string))
		}
	}

	var notificationMedium []int
	for _, id := range d.Get("notification_medium").(*schema.Set).List() {
		if id != nil {
			notificationMedium = append(notificationMedium, id.(int))
		}
	}

	user := &api.User{
		ID:                        d.Id(),
		DisplayName:               d.Get("display_name").(string),
		Email:                     d.Get("email_address").(string),
		UserRole:                  d.Get("user_role").(int),
		JobTitle:                  d.Get("job_title").(int),
		MobileSettings:            d.Get("mobile_settings").(map[string]interface{}),
		SelectionType:             api.ResourceType(d.Get("selection_type").(int)),
		NotificationMedium:        notificationMedium,
		UserGroupIDs:              userGroupIDs,
		MonitorGroups:             monitorGroupIDs,
		Consent_for_non_eu_alerts: d.Get("consent_for_non_eu_alerts").(bool),
	}

	// Alert Settings
	var downNotificationMedium []int
	for _, id := range d.Get("down_notification_medium").(*schema.Set).List() {
		if id != nil {
			downNotificationMedium = append(downNotificationMedium, id.(int))
		}
	}
	if len(downNotificationMedium) == 0 {
		downNotificationMedium = append(downNotificationMedium, 1)
	}

	var criticalNotificationMedium []int
	for _, id := range d.Get("critical_notification_medium").(*schema.Set).List() {
		if id != nil {
			criticalNotificationMedium = append(criticalNotificationMedium, id.(int))
		}
	}
	if len(criticalNotificationMedium) == 0 {
		criticalNotificationMedium = append(criticalNotificationMedium, 1)
	}

	var troubleNotificationMedium []int
	for _, id := range d.Get("trouble_notification_medium").(*schema.Set).List() {
		if id != nil {
			troubleNotificationMedium = append(troubleNotificationMedium, id.(int))
		}
	}
	if len(troubleNotificationMedium) == 0 {
		troubleNotificationMedium = append(troubleNotificationMedium, 1)
	}

	var upNotificationMedium []int
	for _, id := range d.Get("up_notification_medium").(*schema.Set).List() {
		if id != nil {
			upNotificationMedium = append(upNotificationMedium, id.(int))
		}
	}
	if len(upNotificationMedium) == 0 {
		upNotificationMedium = append(upNotificationMedium, 1)
	}

	alertingPeriod := make(map[string]string)
	alertingPeriod["start_time"] = d.Get("alerting_period_start_time").(string)
	alertingPeriod["end_time"] = d.Get("alerting_period_end_time").(string)
	alertSettings := make(map[string]interface{})
	alertSettings["email_format"] = d.Get("email_format").(int)
	alertSettings["alerting_period"] = alertingPeriod
	alertSettings["down"] = downNotificationMedium
	alertSettings["critical"] = criticalNotificationMedium
	alertSettings["trouble"] = troubleNotificationMedium
	alertSettings["up"] = upNotificationMedium
	alertSettings["applogs"] = []int{1}
	alertSettings["anomaly"] = []int{}
	user.AlertSettings = alertSettings

	return user
}

// Called during read - populates the ResourceData with the user in API response
func updateUserResourceData(d *schema.ResourceData, user *api.User) {
	d.Set("display_name", user.DisplayName)
	d.Set("email_address", user.Email)
	d.Set("user_role", user.UserRole)
	d.Set("job_title", user.JobTitle)
	d.Set("selection_type", user.SelectionType)
	d.Set("notification_medium", user.NotificationMedium)
	d.Set("user_group_ids", user.UserGroupIDs)
	d.Set("monitor_groups", user.MonitorGroups)
	d.Set("mobile_settings", user.MobileSettings)

	// Alert Settings

	if user.AlertSettings != nil {
		d.Set("email_format", int(user.AlertSettings["email_format"].(float64)))
		// downNotificationMedium :=
		// var jsonPathArr []string
		// for _, jsonPathData := range jsonPathMapArr {
		// 	jsonPathMap := jsonPathData.(map[string]interface{})
		// 	jsonPath := jsonPathMap["name"].(string)
		// 	jsonPathArr = append(jsonPathArr, jsonPath)
		// }
		alertingPeriodMap := user.AlertSettings["alerting_period"].(map[string]interface{})

		d.Set("alerting_period_start_time", alertingPeriodMap["start_time"].(string))
		d.Set("alerting_period_end_time", alertingPeriodMap["end_time"].(string))

		downNotificationMedium := user.AlertSettings["down"].([]interface{})
		var downNotifArr []int
		for _, notifMedium := range downNotificationMedium {
			downNotifArr = append(downNotifArr, int(notifMedium.(float64)))
		}
		d.Set("down_notification_medium", downNotifArr)

		criticalNotificationMedium := user.AlertSettings["critical"].([]interface{})
		var criticalNotifArr []int
		for _, notifMedium := range criticalNotificationMedium {
			criticalNotifArr = append(criticalNotifArr, int(notifMedium.(float64)))
		}
		d.Set("critical_notification_medium", criticalNotifArr)
		troubleNotificationMedium := user.AlertSettings["trouble"].([]interface{})
		var troubleNotifArr []int
		for _, notifMedium := range troubleNotificationMedium {
			troubleNotifArr = append(troubleNotifArr, int(notifMedium.(float64)))
		}
		d.Set("trouble_notification_medium", troubleNotifArr)
		upNotificationMedium := user.AlertSettings["up"].([]interface{})
		var upNotifArr []int
		for _, notifMedium := range upNotificationMedium {
			upNotifArr = append(upNotifArr, int(notifMedium.(float64)))
		}
		log.Println("upNotifArr =================== ", upNotifArr)
		d.Set("up_notification_medium", upNotifArr)

	}
}
