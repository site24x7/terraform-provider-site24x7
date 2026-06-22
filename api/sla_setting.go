package api

// SLATarget represents an SLA target to be achieved.
type SLATarget struct {
	_               struct{} `type:"structure"`
	TargetID        string   `json:"target_id,omitempty"`
	TargetName      string   `json:"target_name"`
	TargetColor     string   `json:"target_color"`
	TargetCondition int      `json:"target_condition"`
	TargetValue     float64  `json:"target_value"`
}

// SLOAvailability represents the SLA target for monitor availability.
type SLOAvailability struct {
	Availability float64 `json:"availability"`
	Condition    int     `json:"condition"`
	Weightage    float64 `json:"weightage"`
}

// SLOResponseTime represents the SLA target for monitor response time.
type SLOResponseTime struct {
	ResponseTime  float64 `json:"responsetime"`
	TimeAvailable float64 `json:"time_available"`
	Condition     int     `json:"condition"`
	Weightage     float64 `json:"weightage"`
}

// SLASetting represents an SLA report configuration in Site24x7.
type SLASetting struct {
	_               struct{}         `type:"structure"`
	SLAID           string           `json:"sla_id,omitempty"`
	DisplayName     string           `json:"display_name"`
	Type            int              `json:"type"`
	Description     string           `json:"description,omitempty"`
	BusinessHoursID string           `json:"business_hours_id,omitempty"`
	SelectionType   int              `json:"selection_type"`
	Monitors        []string         `json:"monitors,omitempty"`
	MonitorGroups   []string         `json:"monitor_groups,omitempty"`
	SLATargets      []SLATarget      `json:"sla_targets"`
	SLOAvailability *SLOAvailability `json:"slo_availability,omitempty"`
	SLOResponseTime *SLOResponseTime `json:"slo_responsetime,omitempty"`
}
