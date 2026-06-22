package common

import (
	"fmt"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

type MilestoneMarker interface {
	Create(marker *api.MilestoneMarker) (*api.MilestoneMarker, error)
	Update(marker *api.MilestoneMarker) (*api.MilestoneMarker, error)
	Delete(monitorID string, markerTime string) error
	Get(monitorID string, markerTime string) (*api.MilestoneMarker, error)
}

type milestonemarker struct {
	client rest.Client
}

func NewMilestoneMarker(client rest.Client) MilestoneMarker {
	return &milestonemarker{
		client: client,
	}
}

func (c *milestonemarker) Create(marker *api.MilestoneMarker) (*api.MilestoneMarker, error) {
	err := c.client.
		Post().
		Resource("milestone").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(marker).
		Do().
		Err()

	if err != nil {
		return nil, err
	}

	// API returns only {"code": 0, "message": "success"}, so return the input marker.
	return marker, nil
}

func (c *milestonemarker) Update(marker *api.MilestoneMarker) (*api.MilestoneMarker, error) {
	err := c.client.
		Put().
		Resource("milestone").
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		Body(marker).
		Do().
		Err()

	if err != nil {
		return nil, err
	}

	// After update, the effective marker_time becomes new_marker_time if provided.
	updatedMarker := &api.MilestoneMarker{
		MonitorID:  marker.MonitorID,
		MarkerTime: marker.MarkerTime,
		Label:      marker.Label,
		Message:    marker.Message,
	}
	if marker.NewMarkerTime != "" {
		updatedMarker.MarkerTime = marker.NewMarkerTime
	}

	return updatedMarker, nil
}

func (c *milestonemarker) Delete(monitorID string, markerTime string) error {
	return c.client.
		Delete().
		Resource("milestone").
		QueryParams(&api.MilestoneMarkerDeleteParams{
			MarkerTime: markerTime,
			MonitorID:  monitorID,
		}).
		Do().
		Err()
}

// Get retrieves a specific milestone marker by paginating through the list endpoint
// and matching on monitor_id + marker_time.
func (c *milestonemarker) Get(monitorID string, markerTime string) (*api.MilestoneMarker, error) {
	for page := 1; ; page++ {
		listResponse := &api.MilestoneMarkerListResponse{}
		err := c.client.
			Get().
			Resource("milestone").
			QueryParams(&api.MilestoneMarkerListParams{Page: page}).
			Do().
			Parse(listResponse)

		if err != nil {
			return nil, err
		}

		for _, m := range listResponse.MilestoneList {
			if m.MonitorID == monitorID && m.MarkerTime == markerTime {
				return &m, nil
			}
		}

		// If we've exhausted all pages, break out.
		if len(listResponse.MilestoneList) == 0 {
			break
		}
	}

	return nil, fmt.Errorf("milestone marker not found for monitor_id: %s, marker_time: %s", monitorID, markerTime)
}
