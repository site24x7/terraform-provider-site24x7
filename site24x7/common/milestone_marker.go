package common

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var milestoneMarkerSchema = map[string]*schema.Schema{
	"monitor_id": {
		Type:        schema.TypeString,
		Optional:    true,
		ForceNew:    true,
		Default:     "-1",
		Description: "Monitor ID, Group ID, or -1 for a global milestone marker. When omitted, a global milestone marker is created that applies to all monitors.",
	},
	"marker_time": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Timestamp of milestone creation (e.g. 2026-06-03T11:00:00+0530).",
	},
	"label": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Milestone label.",
	},
	"message": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "",
		Description: "Milestone description.",
	},
	"display_name": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Display name of the milestone.",
	},
	"milestone_type": {
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "Milestone marker level.",
	},
}

func ResourceSite24x7MilestoneMarker() *schema.Resource {
	return &schema.Resource{
		Create: resourceSite24x7MilestoneMarkerCreate,
		Read:   resourceSite24x7MilestoneMarkerRead,
		Update: resourceSite24x7MilestoneMarkerUpdate,
		Delete: resourceSite24x7MilestoneMarkerDelete,
		Schema: milestoneMarkerSchema,
	}
}

// compositeID builds a composite ID from monitor_id and marker_time.
func compositeID(monitorID, markerTime string) string {
	return monitorID + ":" + markerTime
}

// parseCompositeID splits a composite ID into monitor_id and marker_time.
func parseCompositeID(id string) (string, string, error) {
	parts := strings.SplitN(id, ":", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid milestone marker ID format: %s (expected monitor_id:marker_time)", id)
	}
	return parts[0], parts[1], nil
}

func resourceSite24x7MilestoneMarkerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	marker := resourceDataToMilestoneMarker(d)

	createdMarker, err := client.MilestoneMarker().Create(marker)
	if err != nil {
		return err
	}

	d.SetId(compositeID(createdMarker.MonitorID, createdMarker.MarkerTime))
	return nil
}

func resourceSite24x7MilestoneMarkerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	monitorID, markerTime, err := parseCompositeID(d.Id())
	if err != nil {
		return err
	}

	marker, err := client.MilestoneMarker().Get(monitorID, markerTime)
	if err != nil {
		// If not found, remove from state.
		d.SetId("")
		return nil
	}

	updateMilestoneMarkerResourceData(d, marker)
	return nil
}

func resourceSite24x7MilestoneMarkerUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	_, oldMarkerTime, err := parseCompositeID(d.Id())
	if err != nil {
		return err
	}

	marker := resourceDataToMilestoneMarker(d)
	newMarkerTime := marker.MarkerTime

	// The API uses marker_time as the old time and new_marker_time as the updated time.
	marker.MarkerTime = oldMarkerTime
	if newMarkerTime != oldMarkerTime {
		marker.NewMarkerTime = newMarkerTime
	}

	updatedMarker, err := client.MilestoneMarker().Update(marker)
	if err != nil {
		return err
	}

	d.SetId(compositeID(updatedMarker.MonitorID, updatedMarker.MarkerTime))
	return nil
}

func resourceSite24x7MilestoneMarkerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	monitorID, markerTime, err := parseCompositeID(d.Id())
	if err != nil {
		return err
	}

	err = client.MilestoneMarker().Delete(monitorID, markerTime)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourceDataToMilestoneMarker(d *schema.ResourceData) *api.MilestoneMarker {
	return &api.MilestoneMarker{
		MonitorID:  d.Get("monitor_id").(string),
		MarkerTime: d.Get("marker_time").(string),
		Label:      d.Get("label").(string),
		Message:    d.Get("message").(string),
	}
}

func updateMilestoneMarkerResourceData(d *schema.ResourceData, marker *api.MilestoneMarker) {
	d.Set("monitor_id", marker.MonitorID)
	d.Set("marker_time", marker.MarkerTime)
	d.Set("label", marker.Label)
	d.Set("message", marker.Message)
	d.Set("display_name", marker.DisplayName)
	d.Set("milestone_type", marker.MilestoneType)
}
