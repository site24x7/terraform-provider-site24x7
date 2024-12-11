package common

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var BusinessHourSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name for the business hour configuration.",
	},
	"description": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Description for the business hour configuration.",
	},
	"time_config": {
		Type:        schema.TypeList,
		Required:    true,
		Description: "Configuration for each day's business hours.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"day": {
					Type:        schema.TypeInt,
					Required:    true,
					Description: "Day of the week (1 for Sunday, 7 for Saturday).",
				},
				"start_time": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Start time in HH:mm format.",
				},
				"end_time": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "End time in HH:mm format.",
				},
			},
		},
	},
}

func ResourceSite24x7BusinessHour() *schema.Resource {
	return &schema.Resource{
		Create: businessHourCreate,
		Read:   businessHourRead,
		Update: businessHourUpdate,
		Delete: businessHourDelete,
		Exists: businessHourExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: BusinessHourSchema,
	}
}

func businessHourCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	log.Println("[INFO] Creating new BusinessHour resource")

	businessHour := resourceDataToBusinessHour(d)

	businessHour, err := client.BusinessHour().Create(businessHour)
	if err != nil {
		log.Printf("[ERROR] Failed to create BusinessHour: %v", err)
		return fmt.Errorf("failed to create business hour: %w", err)
	}
	log.Printf("[DEBUG] Created BusinessHour with ID: %s, DisplayName: %s", businessHour.ID, businessHour.DisplayName)
	d.SetId(businessHour.ID)

	// Check for errors when setting fields in the schema
	if err := d.Set("display_name", businessHour.DisplayName); err != nil {
		log.Printf("[ERROR] Failed to set display_name in state: %v", err)
		return fmt.Errorf("error setting display_name in state after creation: %w", err)
	}

	return nil
}

func businessHourRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	businessHour, err := client.BusinessHour().Get(d.Id())
	if apierrors.IsNotFound(err) {
		log.Printf("[WARN] BusinessHour with ID %s not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}
	if err != nil {
		log.Printf("[ERROR] Failed to read BusinessHour with ID %s: %v", d.Id(), err)
		return fmt.Errorf("failed to read business hour with ID %s: %w", d.Id(), err)
	}
	log.Printf("[DEBUG] Retrieved BusinessHour with ID: %s, DisplayName: %s", businessHour.ID, businessHour.DisplayName)
	updateBusinessHourResourceData(d, businessHour)
	log.Printf("[INFO] Successfully updated state for BusinessHour ID: %s", d.Id())
	return nil
}

func businessHourUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)
	log.Printf("[INFO] Updating BusinessHour with ID: %s", d.Id())
	businessHour := resourceDataToBusinessHour(d)

	businessHour, err := client.BusinessHour().Update(businessHour)
	if err != nil {
		log.Printf("[ERROR] Failed to update BusinessHour with ID %s: %v", d.Id(), err)
		return fmt.Errorf("failed to update business hour with ID %s: %w", businessHour.ID, err)
	}
	log.Printf("[DEBUG] Updated BusinessHour with ID: %s, DisplayName: %s", businessHour.ID, businessHour.DisplayName)
	d.SetId(businessHour.ID)

	// Handle potential errors in state updates
	if err := d.Set("display_name", businessHour.DisplayName); err != nil {
		log.Printf("[ERROR] Failed to set display_name in state after update: %v", err)
		return fmt.Errorf("error setting display_name in state after update: %w", err)
	}

	return nil
}

func businessHourDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.BusinessHour().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to delete business hour with ID %s: %w", d.Id(), err)
	}

	return nil
}

func businessHourExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.BusinessHour().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("failed to check existence of business hour with ID %s: %w", d.Id(), err)
	}

	return true, nil
}

func resourceDataToBusinessHour(d *schema.ResourceData) *api.BusinessHour {
	var timeConfig []api.TimeSlot
	for _, tc := range d.Get("time_config").([]interface{}) {
		data := tc.(map[string]interface{})
		timeConfig = append(timeConfig, api.TimeSlot{
			Day:       data["day"].(int),
			StartTime: data["start_time"].(string),
			EndTime:   data["end_time"].(string),
		})
	}

	return &api.BusinessHour{
		ID:          d.Id(),
		DisplayName: d.Get("display_name").(string),
		Description: d.Get("description").(string),
		TimeConfig:  timeConfig,
	}
}

func updateBusinessHourResourceData(d *schema.ResourceData, businessHour *api.BusinessHour) {
	d.Set("display_name", businessHour.DisplayName)
	d.Set("description", businessHour.Description)

	var timeConfig []map[string]interface{}
	for _, tc := range businessHour.TimeConfig {
		timeConfig = append(timeConfig, map[string]interface{}{
			"day":        tc.Day,
			"start_time": tc.StartTime,
			"end_time":   tc.EndTime,
		})
	}
	d.Set("time_config", timeConfig)
}
