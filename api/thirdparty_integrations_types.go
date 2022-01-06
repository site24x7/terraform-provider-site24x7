package api

import (
	"encoding/json"
	"strconv"
)

type ResourceType int

func (resourceType *ResourceType) UnmarshalJSON(rawValue []byte) error {
	if rawValue[0] != '"' {
		return json.Unmarshal(rawValue, (*int)(resourceType))
	}

	var valueAsString string
	if err := json.Unmarshal(rawValue, &valueAsString); err != nil {
		return err
	}

	valueAsInt, err := strconv.Atoi(valueAsString)
	if err != nil {
		return err
	}

	*resourceType = ResourceType(valueAsInt)
	return nil
}

type ThirdPartyIntegrations struct {
	_             struct{}     `type:"structure"` // Enforces key based initialization.
	ServiceID     string       `json:"service_id"`
	ServiceStatus int          `json:"service_status"`
	ServiceKey    string       `json:"service_key,omitempty"`
	Name          string       `json:"name"`
	SenderName    string       `json:"sender_name,omitempty"`
	Title         string       `json:"title,omitempty"`
	SelectionType ResourceType `json:"selection_type"`
	TroubleAlert  bool         `json:"trouble_alert"`
	Type          int          `json:"type"`
}

// Denotes opsgenie integration resource in Site24x7.
type OpsgenieIntegration struct {
	_             struct{}     `type:"structure"` // Enforces key based initialization.
	ServiceID     string       `json:"service_id,omitempty"`
	ServiceStatus int          `json:"service_status,omitempty"`
	Name          string       `json:"name"`
	URL           string       `json:"url"`
	SelectionType ResourceType `json:"selection_type"`
	Monitors      []string     `json:"monitors,omitempty"`
	TroubleAlert  bool         `json:"trouble_alert"`
	ManualResolve bool         `json:"manual_resolve"`
	AlertTagIDs   []string     `json:"alert_tags_id,omitempty"`
}

// Denotes slack integration resource in Site24x7.
type SlackIntegration struct {
	_             struct{}     `type:"structure"` // Enforces key based initialization.
	ServiceID     string       `json:"service_id,omitempty"`
	ServiceStatus int          `json:"service_status,omitempty"`
	Name          string       `json:"name"`
	URL           string       `json:"url"`
	SelectionType ResourceType `json:"selection_type"`
	SenderName    string       `json:"sender_name"`
	Title         string       `json:"title"`
	Monitors      []string     `json:"monitors,omitempty"`
	AlertTagIDs   []string     `json:"alert_tags_id,omitempty"`
}

// Denotes webhook integration resource in Site24x7.
type WebhookIntegration struct {
	_                            struct{}     `type:"structure"` // Enforces key based initialization.
	ServiceID                    string       `json:"service_id,omitempty"`
	ServiceStatus                int          `json:"service_status,omitempty"`
	Name                         string       `json:"name"`
	URL                          string       `json:"url"`
	Timeout                      int          `json:"timeout"`
	Method                       string       `json:"method"`
	SelectionType                ResourceType `json:"selection_type"`
	IsPollerWebhook              bool         `json:"is_poller_webhook"`
	Poller                       string       `json:"poller,omitempty"`
	SendIncidentParameters       bool         `json:"send_incident_parameters"`
	SendCustomParameters         bool         `json:"send_custom_parameters,omitempty"`
	CustomParameters             string       `json:"custom_parameters,omitempty"`
	SendInJsonFormat             bool         `json:"send_in_json_format,omitempty"`
	AuthMethod                   string       `json:"auth_method,omitempty"`
	Username                     string       `json:"username,omitempty"`
	Password                     string       `json:"password,omitempty"`
	OauthProvider                string       `json:"oauth2_provider,omitempty"`
	UserAgent                    string       `json:"user_agent,omitempty"`
	Monitors                     []string     `json:"monitors,omitempty"`
	ManageTickets                bool         `json:"manage_tickets"`
	UpdateURL                    string       `json:"update_url,omitempty"`
	UpdateMethod                 string       `json:"update_method,omitempty"`
	UpdateSendIncidentParameters bool         `json:"update_send_incident_parameters,omitempty"`
	UpdateSendCustomParameters   bool         `json:"update_send_custom_parameters,omitempty"`
	UpdateCustomParameters       string       `json:"update_custom_parameters,omitempty"`
	CloseURL                     string       `json:"close_url,omitempty"`
	CloseMethod                  string       `json:"close_method,omitempty"`
	CloseSendIncidentParameters  bool         `json:"close_send_incident_parameters,omitempty"`
	CloseSendCustomParameters    bool         `json:"close_send_custom_parameters,omitempty"`
	CloseCustomParameters        string       `json:"close_custom_parameters,omitempty"`
	AlertTagIDs                  []string     `json:"alert_tags_id,omitempty"`
}
