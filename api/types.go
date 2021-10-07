package api

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Status int

type MonitorType string

// Custom unmarshaller that allows the value to be both a
// string and an integer, always unmarshals into integer
//
// The Site24x7 API has a bug where it accepts integers, but
// returns them as strings.
func (status *Status) UnmarshalJSON(rawValue []byte) error {
	if rawValue[0] != '"' {
		return json.Unmarshal(rawValue, (*int)(status))
	}

	var valueAsString string
	if err := json.Unmarshal(rawValue, &valueAsString); err != nil {
		return err
	}

	valueAsInt, err := strconv.Atoi(valueAsString)
	if err != nil {
		return err
	}

	*status = Status(valueAsInt)
	return nil
}

const (
	Down               Status = 0
	Up                 Status = 1
	Trouble            Status = 2
	Critical           Status = 3
	Suspended          Status = 5
	Maintenance        Status = 7
	Discovery          Status = 9
	ConfigurationError Status = 10
)

const (
	URL      MonitorType = "URL"
	SSL_CERT MonitorType = "SSL_CERT"
	RESTAPI  MonitorType = "RESTAPI"
	AMAZON   MonitorType = "AMAZON"
)

type ValueAndSeverity struct {
	Value    string `json:"value"`
	Severity Status `json:"severity"`
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ActionRef struct {
	ActionID  string `json:"action_id"`
	AlertType Status `json:"alert_type"`
}

// Monitor performance of websites and internet services
type Monitor struct {
	MonitorID             string            `json:"monitor_id,omitempty"`
	DisplayName           string            `json:"display_name"`
	Type                  string            `json:"type"`
	Website               string            `json:"website"`
	CheckFrequency        string            `json:"check_frequency"`
	HTTPMethod            string            `json:"http_method"`
	AuthUser              string            `json:"auth_user"`
	AuthPass              string            `json:"auth_pass"`
	MatchingKeyword       *ValueAndSeverity `json:"matching_keyword,omitempty"`
	UnmatchingKeyword     *ValueAndSeverity `json:"unmatching_keyword,omitempty"`
	MatchRegex            *ValueAndSeverity `json:"match_regex,omitempty"`
	MatchCase             bool              `json:"match_case"`
	UserAgent             string            `json:"user_agent"`
	CustomHeaders         []Header          `json:"custom_headers,omitempty"`
	Timeout               int               `json:"timeout"`
	LocationProfileID     string            `json:"location_profile_id"`
	NotificationProfileID string            `json:"notification_profile_id"`
	ThresholdProfileID    string            `json:"threshold_profile_id"`
	MonitorGroups         []string          `json:"monitor_groups,omitempty"`
	UserGroupIDs          []string          `json:"user_group_ids,omitempty"`
	ActionIDs             []ActionRef       `json:"action_ids,omitempty"`
	UseNameServer         bool              `json:"use_name_server"`
	UpStatusCodes         string            `json:"up_status_codes"`
}

// Monitor performance of websites and internet services
type WebsiteMonitor struct {
	MonitorID             string            `json:"monitor_id,omitempty"`
	DisplayName           string            `json:"display_name"`
	Type                  string            `json:"type"`
	Website               string            `json:"website"`
	CheckFrequency        string            `json:"check_frequency"`
	HTTPMethod            string            `json:"http_method"`
	AuthUser              string            `json:"auth_user"`
	AuthPass              string            `json:"auth_pass"`
	MatchingKeyword       *ValueAndSeverity `json:"matching_keyword,omitempty"`
	UnmatchingKeyword     *ValueAndSeverity `json:"unmatching_keyword,omitempty"`
	MatchRegex            *ValueAndSeverity `json:"match_regex,omitempty"`
	MatchCase             bool              `json:"match_case"`
	UserAgent             string            `json:"user_agent"`
	CustomHeaders         []Header          `json:"custom_headers,omitempty"`
	Timeout               int               `json:"timeout"`
	LocationProfileID     string            `json:"location_profile_id"`
	NotificationProfileID string            `json:"notification_profile_id"`
	ThresholdProfileID    string            `json:"threshold_profile_id"`
	MonitorGroups         []string          `json:"monitor_groups,omitempty"`
	UserGroupIDs          []string          `json:"user_group_ids,omitempty"`
	ActionIDs             []ActionRef       `json:"action_ids,omitempty"`
	UseNameServer         bool              `json:"use_name_server"`
	UpStatusCodes         string            `json:"up_status_codes"`
}

func (websiteMonitor *WebsiteMonitor) String() string {
	return ToString(websiteMonitor)
}

type SSLMonitor struct {
	MonitorID             string      `json:"monitor_id,omitempty"`
	DisplayName           string      `json:"display_name"`
	DomainName            string      `json:"domain_name"`
	Type                  string      `json:"type"`
	Timeout               int         `json:"timeout"`
	Protocol              string      `json:"protocol"`
	Port                  interface{} `json:"port"` // To Fix: API accepts int and returns string in response.
	ExpireDays            int         `json:"expire_days"`
	HTTPProtocolVersion   string      `json:"http_protocol"`
	IgnoreDomainMismatch  bool        `json:"ignore_domain_mismatch"`
	IgnoreTrust           bool        `json:"ignore_trust"`
	LocationProfileID     string      `json:"location_profile_id"`
	NotificationProfileID string      `json:"notification_profile_id"`
	ThresholdProfileID    string      `json:"threshold_profile_id"`
	MonitorGroups         []string    `json:"monitor_groups,omitempty"`
	UserGroupIDs          []string    `json:"user_group_ids,omitempty"`
	// ActionIDs             []ActionRef `json:"action_ids,omitempty"`
}

func (sslMonitor *SSLMonitor) String() string {
	return ToString(sslMonitor)
}

type RestApiMonitor struct {
	MonitorID                 string                 `json:"monitor_id,omitempty"`
	DisplayName               string                 `json:"display_name"`
	Type                      string                 `json:"type"`
	Website                   string                 `json:"website"`
	CheckFrequency            string                 `json:"check_frequency"`
	Timeout                   int                    `json:"timeout"`
	HttpMethod                string                 `json:"http_method,omitempty"`
	HttpProtocol              string                 `json:"http_protocol,omitempty"`
	SslProtocol               string                 `json:"ssl_protocol,omitempty"`
	LocationProfileID         string                 `json:"location_profile_id"`
	NotificationProfileID     string                 `json:"notification_profile_id"`
	ThresholdProfileID        string                 `json:"threshold_profile_id"`
	MonitorGroups             []string               `json:"monitor_groups,omitempty"`
	UserGroupIDs              []string               `json:"user_group_ids,omitempty"`
	ActionIDs                 []ActionRef            `json:"action_ids,omitempty"`
	RequestContentType        string                 `json:"request_content_type,omitempty"`
	ResponseContentType       string                 `json:"response_type,omitempty"`
	RequestParam              string                 `json:"request_param,omitempty"`
	AuthUser                  string                 `json:"auth_user,omitempty"`
	AuthPass                  string                 `json:"auth_pass,omitempty"`
	OAuth2Provider            string                 `json:"oauth2_provider,omitempty"`
	ClientCertificatePassword string                 `json:"client_certificate_password,omitempty"`
	JwtID                     string                 `json:"jwt_id,omitempty"`
	MatchingKeyword           map[string]interface{} `json:"matching_keyword,omitempty"`
	UnmatchingKeyword         map[string]interface{} `json:"unmatching_keyword,omitempty"`
	MatchRegex                map[string]interface{} `json:"match_regex,omitempty"`
	UseAlpn                   bool                   `json:"use_alpn"`
	UseIPV6                   bool                   `json:"use_ipv6"`
	MatchCase                 bool                   `json:"match_case"`
	JSONSchemaCheck           bool                   `json:"json_schema_check"`
	UseNameServer             bool                   `json:"use_name_server"`
	UserAgent                 string                 `json:"user_agent"`
	CustomHeaders             []Header               `json:"custom_headers,omitempty"`
}

func (restApiMonitor *RestApiMonitor) String() string {
	return ToString(restApiMonitor)
}

type AmazonMonitor struct {
	MonitorID             string   `json:"monitor_id,omitempty"`
	DisplayName           string   `json:"display_name"`
	Type                  string   `json:"type"`
	SecretKey             string   `json:"aws_secret_key"`
	AccessKey             string   `json:"aws_access_key"`
	DiscoverFrequency     int      `json:"aws_discovery_frequency"`
	DiscoverServices      []string `json:"aws_discover_services"`
	NotificationProfileID string   `json:"notification_profile_id"`
	UserGroupIDs          []string `json:"user_group_ids"`
}

func (amazonMonitor *AmazonMonitor) String() string {
	return ToString(amazonMonitor)
}

// MonitorGroup organizes Monitor resources into groups.
type MonitorGroup struct {
	GroupID                string   `json:"group_id,omitempty"`
	DisplayName            string   `json:"display_name"`
	Description            string   `json:"description,omitempty"`
	Monitors               []string `json:"monitors,omitempty"`
	HealthThresholdCount   int      `json:"health_threshold_count,omitempty"`
	DependencyResourceID   []string `json:"dependency_resource_ids,omitempty"`
	SuppressAlert          bool     `json:"suppress_alert"`
	DependencyResourceType int      `json:"selection_type,omitempty"`
}

func (monitorGroup *MonitorGroup) String() string {
	return ToString(monitorGroup)
}

// ThresholdProfile help the alarms engine to decide if a specific resource has to be declared critical or down
type ThresholdProfile struct {
	_                      struct{}                 `type:"structure"` // Enforces key based initialization.
	ProfileID              string                   `json:"profile_id,omitempty"`
	Type                   string                   `json:"type"` // Denotes monitor type
	ProfileName            string                   `json:"profile_name"`
	ProfileType            int                      `json:"profile_type"` // 1- Static Threshold or 2 - AI-based Threshold
	DownLocationThreshold  int                      `json:"down_location_threshold"`
	WebsiteContentModified bool                     `json:"website_content_modified"`
	WebsiteContentChanges  []map[string]interface{} `json:"website_content_changes,omitempty"`
	ResponseTimeThreshold  map[string]interface{}   `json:"response_time_threshold,omitempty"`
}

func (thresholdProfile *ThresholdProfile) String() string {
	return ToString(thresholdProfile)
}

func (thresholdProfile *ThresholdProfile) UnmarshalJSON(rawValue []byte) error {
	var f interface{}
	if err := json.Unmarshal(rawValue, &f); err != nil {
		return err
	}
	if f == nil {
		return nil
	}
	m := f.(map[string]interface{})
	for k, v := range m {
		if k == "profile_id" {
			thresholdProfile.ProfileID, _ = v.(string)
		} else if k == "type" {
			thresholdProfile.Type, _ = v.(string)
		} else if k == "profile_name" {
			thresholdProfile.ProfileName, _ = v.(string)
		} else if k == "profile_type" {
			thresholdProfile.ProfileType = int(v.(float64))
		} else if k == "down_location_threshold" {
			thresholdProfile.DownLocationThreshold = int(v.(float64))
		} else if k == "website_content_modified" {
			thresholdProfile.WebsiteContentModified, _ = v.(bool)
		} else if k == "website_content_changes" {
			thresholdProfile.WebsiteContentChanges = make([]map[string]interface{}, 1, 1)
			switch val := v.(type) {
			case []interface{}:
				for _, x := range val {
					fmt.Println("this is b", x.(map[string]interface{}))
					thresholdProfile.WebsiteContentChanges = append(thresholdProfile.WebsiteContentChanges, x.(map[string]interface{}))
				}
			}
		} else if k == "response_time_threshold" {
			thresholdProfile.ResponseTimeThreshold = v.(map[string]interface{})
		}
	}
	return nil
}

// NotificationProfile allows tweaking when alerts have to be sent out.
type NotificationProfile struct {
	ProfileID                   string   `json:"profile_id,omitempty"`
	ProfileName                 string   `json:"profile_name"`
	RcaNeeded                   bool     `json:"rca_needed"`
	NotifyAfterExecutingActions bool     `json:"notify_after_executing_actions"`
	DowntimeNotificationDelay   int      `json:"downtime_notification_delay,omitempty"`
	PersistentNotification      int      `json:"persistent_notification,omitempty"`
	EscalationUserGroupId       string   `json:"escalation_user_group_id,omitempty"`
	EscalationWaitTime          int      `json:"escalation_wait_time"`
	EscalationAutomations       []string `json:"escalation_automations,omitempty"`
	EscalationServices          []string `json:"escalation_services,omitempty"`
	TemplateID                  string   `json:"template_id,omitempty"`
}

func (notificationProfile *NotificationProfile) String() string {
	return ToString(notificationProfile)
}

// LocationProfiles make it convenient to set monitoring locations consistently across many websites or monitors
type LocationProfile struct {
	ProfileID          string   `json:"profile_id,omitempty"`
	ProfileName        string   `json:"profile_name"`
	PrimaryLocation    string   `json:"primary_location"`
	SecondaryLocations []string `json:"secondary_locations,omitempty"`
	RestrictAltLoc     bool     `json:"restrict_alt_loc"`
}

// LocationTemplate holds locations Site24x7 performs their monitor checks
// from.
type LocationTemplate struct {
	Locations []*Location `json:"locations"`
}

// Location is a physical location Site24x7 performs monitor checks from. The
// LocationID field maps to the IDs used in the PrimaryLocation and
// SecondaryLocations fields of LocationProfile values.
type Location struct {
	LocationID  string `json:"location_id"`
	CountryName string `json:"country_name"`
	DisplayName string `json:"display_name"`
	UseIPV6     bool   `json:"use_ipv6"`
	CityName    string `json:"city_name"`
	CityShort   string `json:"city_short"`
	Continent   string `json:"continent"`
}

// UserGroup help organize individuals so that they receive alerts and reports based on their responsibility.
type UserGroup struct {
	UserGroupID      string   `json:"user_group_id,omitempty"`
	DisplayName      string   `json:"display_name"`
	Users            []string `json:"users"`
	AttributeGroupID string   `json:"attribute_group_id"`
	ProductID        int      `json:"product_id"`
}

func (userGroup *UserGroup) String() string {
	return ToString(userGroup)
}

// URLAutomation prioritize and remediate routine actions automatically,
// increase IT efficiency and streamline your processes to reduce performance degrade
type URLAutomation struct {
	ActionID               string `json:"action_id,omitempty"`
	ActionType             int    `json:"action_type"`
	ActionName             string `json:"action_name"`
	ActionUrl              string `json:"action_url"`
	ActionTimeout          int    `json:"action_timeout"`
	ActionMethod           string `json:"action_method"`
	SuppressAlert          bool   `json:"suppress_alert,omitempty"`
	SendIncidentParameters bool   `json:"send_incident_parameters"`
	SendCustomParameters   bool   `json:"send_custom_parameters"`
	CustomParameters       string `json:"custom_parameters"`
	SendInJsonFormat       bool   `json:"send_in_json_format"`
	SendEmail              bool   `json:"send_mail"`
	AuthMethod             string `json:"auth_method,omitempty"`
	Username               string `json:"username,omitempty"`
	Password               string `json:"password,omitempty"`
	OAuth2Provider         string `json:"oauth2_provider,omitempty"`
	UserAgent              string `json:"user_agent,omitempty"`
}

func (urlAutomation *URLAutomation) String() string {
	return ToString(urlAutomation)
}

// MonitorsStatus describes the response for the current status endpoint as
// defined here: https://www.site24x7.com/help/api/#retrieve-current-status.
type MonitorsStatus struct {
	Monitors []*MonitorStatus `json:"monitors"`
}

// MonitorStatus describes a monitor status response as defined here:
// https://www.site24x7.com/help/api/#retrieve-current-status.
type MonitorStatus struct {
	Name           string   `json:"name"`
	MonitorID      string   `json:"monitor_id"`
	MonitorType    string   `json:"monitor_type"`
	Status         Status   `json:"status"`
	LastPolledTime string   `json:"last_polled_time"`
	Unit           string   `json:"unit"`
	OutageID       string   `json:"outage_id"`
	DowntimeMillis string   `json:"downtime_millis"`
	DownReason     string   `json:"down_reason"`
	Duration       string   `json:"duration"`
	ServerType     string   `json:"server_type"`
	Tags           []string `json:"tags"`
}

// CurrentStatusListOptions hold the options that can be specified to filter
// current monitor statuses.
type CurrentStatusListOptions struct {
	APMRequired       *bool   `url:"apm_required,omitempty"`
	GroupRequired     *bool   `url:"group_required,omitempty"`
	SuspendedRequired *bool   `url:"suspended_required,omitempty"`
	LocationsRequired *bool   `url:"locations_required,omitempty"`
	StatusRequired    *string `url:"status_required,omitempty"`
}
