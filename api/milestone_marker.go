package api

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// MilestoneMarker represents a milestone marker in Site24x7.
// MonitorID can be a monitor ID string, a group ID string, or "-1" for global milestones.
type MilestoneMarker struct {
	_             struct{} `type:"structure"`
	MonitorID     string   `json:"-"` // Custom marshal/unmarshal - excluded from default JSON handling
	MarkerTime    string   `json:"marker_time"`
	NewMarkerTime string   `json:"new_marker_time,omitempty"`
	Label         string   `json:"label"`
	Message       string   `json:"message,omitempty"`
	DisplayName   string   `json:"display_name,omitempty"`
	MilestoneType int      `json:"milestone_type,omitempty"`
}

// milestoneMarkerJSON is an internal type used for custom JSON marshaling/unmarshaling.
type milestoneMarkerJSON struct {
	MonitorID     interface{} `json:"monitor_id"`
	MarkerTime    string      `json:"marker_time"`
	NewMarkerTime string      `json:"new_marker_time,omitempty"`
	Label         string      `json:"label"`
	Message       string      `json:"message,omitempty"`
	DisplayName   string      `json:"display_name,omitempty"`
	MilestoneType int         `json:"milestone_type,omitempty"`
}

// MarshalJSON implements custom JSON marshaling for MilestoneMarker.
// For global milestones (MonitorID == "-1"), monitor_id is serialized as integer -1.
// For monitor/group milestones, monitor_id is serialized as a string.
func (m MilestoneMarker) MarshalJSON() ([]byte, error) {
	aux := milestoneMarkerJSON{
		MarkerTime:    m.MarkerTime,
		NewMarkerTime: m.NewMarkerTime,
		Label:         m.Label,
		Message:       m.Message,
		DisplayName:   m.DisplayName,
		MilestoneType: m.MilestoneType,
	}

	if m.MonitorID == "-1" {
		aux.MonitorID = -1
	} else {
		aux.MonitorID = m.MonitorID
	}

	return json.Marshal(aux)
}

// UnmarshalJSON implements custom JSON unmarshaling for MilestoneMarker.
// Handles monitor_id as either an integer (e.g. -1 for global) or a string (e.g. "15698000017614001" for monitor-specific).
func (m *MilestoneMarker) UnmarshalJSON(data []byte) error {
	aux := &milestoneMarkerJSON{}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	m.MarkerTime = aux.MarkerTime
	m.NewMarkerTime = aux.NewMarkerTime
	m.Label = aux.Label
	m.Message = aux.Message
	m.DisplayName = aux.DisplayName
	m.MilestoneType = aux.MilestoneType

	// Handle monitor_id which can be a JSON number or string.
	switch v := aux.MonitorID.(type) {
	case string:
		m.MonitorID = v
	case float64:
		// JSON numbers are unmarshaled as float64 by default.
		m.MonitorID = strconv.FormatInt(int64(v), 10)
	case nil:
		// Global milestones don't have monitor_id in the API response.
		m.MonitorID = "-1"
	default:
		return fmt.Errorf("unexpected type for monitor_id: %T", aux.MonitorID)
	}

	return nil
}

// MilestoneMarkerListResponse represents the response from the list milestone markers API.
type MilestoneMarkerListResponse struct {
	MilestoneList []MilestoneMarker `json:"milestone_list"`
	TotalCount    int               `json:"total_count"`
}

// MilestoneMarkerDeleteParams holds query parameters for deleting a milestone marker.
type MilestoneMarkerDeleteParams struct {
	MarkerTime string `url:"marker_time"`
	MonitorID  string `url:"monitor_id"`
}

// MilestoneMarkerListParams holds query parameters for listing milestone markers.
type MilestoneMarkerListParams struct {
	Page int `url:"page"`
}
