package api

// MilestoneMarker represents a milestone marker in Site24x7.
type MilestoneMarker struct {
	_             struct{} `type:"structure"`
	MonitorID     string   `json:"monitor_id"`
	MarkerTime    string   `json:"marker_time"`
	NewMarkerTime string   `json:"new_marker_time,omitempty"`
	Label         string   `json:"label"`
	Message       string   `json:"message,omitempty"`
	DisplayName   string   `json:"display_name,omitempty"`
	MilestoneType int      `json:"milestone_type,omitempty"`
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
