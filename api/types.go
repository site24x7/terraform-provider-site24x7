package api

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type Status int

type ResponseType string

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

// Type of the Site24x7 resource.
type MonitorType string

type ValueAndSeverity struct {
	Severity Status `json:"severity"`
	Value    string `json:"value"`
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SearchConfig struct {
	Addr     string `json:"addr,omitempty"`
	TTLO     int    `json:"ttlo,omitempty"`
	TTL      int    `json:"ttl,omitempty"`
	Target   string `json:"target,omitempty"`
	Priority int    `json:"priority,omitempty"`
	Port     int    `json:"port,omitempty"`
	Wt       int    `json:"wt,omitempty"`
	Rcvd     int    `json:"rcvd,omitempty"`
	PNS      string `json:"pns,omitempty"`
	Admin    string `json:"admin,omitempty"`
	Serial   int    `json:"serial,omitempty"`
	RFF      int    `json:"rff,omitempty"`
	RTF      int    `json:"rtf,omitempty"`
	EXPT     int    `json:"expt,omitempty"`
	MTTL     int    `json:"mttl,omitempty"`
	Flg      int    `json:"flg,omitempty"`
	Prtcl    int    `json:"prtcl,omitempty"`
	Kalg     int    `json:"kalg,omitempty"`
	Kid      int    `json:"kid,omitempty"`
	Key      string `json:"key,omitempty"`
	Tag      string `json:"tag,omitempty"`
	CertAuth string `json:"certauth,omitempty"`
	Halg     int    `json:"halg,omitempty"`
	Hash     string `json:"hash,omitempty"`
}

type Steps struct {
	DisplayName  string        `json:"display_name"`
	StepsDetails []StepDetails `json:"step_details"`
	MonitorID    string        `json:"monitor_id,omitempty"`
	StepId       string        `json:"step_id,omitempty"`
}

type StepDetails struct {
	StepId      string `json:"step_id,omitempty"`
	StepUrl     string `json:"step_url"`
	StopOnErr   int    `json:"severity"`
	DisplayName string `json:"display_name"`
	// HTTP Configuration
	Timeout                   string                 `json:"timeout"`
	HTTPMethod                string                 `json:"http_method"`
	RequestContentType        string                 `json:"request_content_type,omitempty"`
	RequestBody               string                 `json:"request_param,omitempty"`
	RequestHeaders            []Header               `json:"custom_headers,omitempty"`
	GraphQL                   map[string]interface{} `json:"graphql,omitempty"`
	UserAgent                 string                 `json:"user_agent,omitempty"`
	AuthMethod                string                 `json:"auth_method,omitempty"`
	AuthUser                  string                 `json:"auth_user,omitempty"`
	AuthPass                  string                 `json:"auth_pass,omitempty"`
	OAuth2Provider            string                 `json:"oauth2_provider,omitempty"`
	ClientCertificatePassword string                 `json:"client_certificate_password,omitempty"`
	JwtID                     string                 `json:"jwt_id,omitempty"`
	UseNameServer             bool                   `json:"use_name_server"`
	HTTPProtocol              string                 `json:"http_protocol,omitempty"`
	SSLProtocol               string                 `json:"ssl_protocol,omitempty"`
	UpStatusCodes             string                 `json:"up_status_codes,omitempty"`
	UseAlpn                   bool                   `json:"use_alpn"`
	// Content Check
	ResponseContentType string                  `json:"response_type"`
	MatchJSON           map[string]interface{}  `json:"match_json,omitempty"`
	JSONSchema          map[string]interface{}  `json:"json_schema,omitempty"`
	JSONSchemaCheck     bool                    `json:"json_schema_check,omitempty"`
	MatchingKeyword     map[string]interface{}  `json:"matching_keyword,omitempty"`
	UnmatchingKeyword   map[string]interface{}  `json:"unmatching_keyword,omitempty"`
	MatchCase           bool                    `json:"match_case"`
	MatchRegex          map[string]interface{}  `json:"match_regex,omitempty"`
	ResponseHeaders     HTTPResponseHeader      `json:"response_headers_check,omitempty"`
	ResponseVariable    HTTPResponseVariable    `json:"response_variables,omitempty"`
	DynamicHeaderParams HTTPDynamicHeaderParams `json:"dynamic_header_params,omitempty"`
}

type HTTPResponseHeader struct {
	Severity Status   `json:"severity"`
	Value    []Header `json:"value"`
}

type HTTPResponseVariable struct {
	ResponseType ResponseType `json:"response_type"`
	Variables    []Header     `json:"variables"`
}

type HTTPDynamicHeaderParams struct {
	Variables []Header `json:"variables"`
}

type ActionRef struct {
	ActionID  string `json:"action_id"`
	AlertType Status `json:"alert_type"`
}

// MonitorGroup organizes monitor resources into groups.
type MonitorGroup struct {
	_           struct{} `type:"structure"` // Enforces key based initialization.
	GroupID     string   `json:"group_id,omitempty"`
	DisplayName string   `json:"display_name"`
	Description string   `json:"description,omitempty"`
	// GroupType                int      `json:"group_type,omitempty"`
	Monitors                 []string `json:"monitors,omitempty"`
	HealthThresholdCount     int      `json:"health_threshold_count"`
	DependencyResourceIDs    []string `json:"dependency_resource_ids,omitempty"`
	SuppressAlert            bool     `json:"suppress_alert"`
	DependencyResourceType   int      `json:"selection_type,omitempty"`
	NotificationProfileID    string   `json:"notification_profile_id"`
	HealthCheckProfileID     string   `json:"healthcheck_profile_id"`
	TagIDs                   []string `json:"tags,omitempty"`
	UserGroupIDs             []string `json:"user_group_ids,omitempty"`
	ThirdPartyServiceIDs     []string `json:"third_party_services,omitempty"`
	EnableIncidentManagement bool     `json:"enable_incident_management"`
	HealingPeriod            int      `json:"healing_period"`
	AlertFrequency           int      `json:"alert_frequency"`
	AlertPeriodically        bool     `json:"alert_periodically"`
}

func (monitorGroup *MonitorGroup) String() string {
	return ToString(monitorGroup)
}

// Subgroups help you revisualize the high level architecture of your monitor group in a business view inside the web client. Create nested subgroups under your monitor group. Its a handy concept for easy administration.
type Subgroup struct {
	_                    struct{} `type:"structure"` // Enforces key based initialization.
	ID                   string   `json:"group_id,omitempty"`
	DisplayName          string   `json:"display_name"`
	TopGroupID           string   `json:"top_group_id"`
	ParentGroupID        string   `json:"parent_group_id"`
	Description          string   `json:"description,omitempty"`
	Type                 int      `json:"group_type"`
	Monitors             []string `json:"monitors,omitempty"`
	HealthThresholdCount int      `json:"health_threshold_count,omitempty"`
}

func (subgroup *Subgroup) String() string {
	return ToString(subgroup)
}

// Tags helps you to organize monitor resources and locate related items that have the same tag.
type Tag struct {
	_        struct{} `type:"structure"` // Enforces key based initialization.
	TagID    string   `json:"tag_id,omitempty"`
	TagName  string   `json:"tag_name"`
	TagValue string   `json:"tag_value,omitempty"`
	TagType  int      `json:"tag_type,omitempty"`  // 1 - User Defined Tag or 2 - AWS System Generated Tag
	TagColor string   `json:"tag_color,omitempty"` // Possible Colors - '#B7DA9E','#73C7A3','#B5DCDF','#D4ABBB','#4895A8','#DFE897','#FCEA8B','#FFC36D','#F79953','#F16B3C','#E55445','#F2E2B6','#DEC57B','#CBBD80','#AAB3D4','#7085BA','#F6BDAE','#EFAB6D','#CA765C','#999','#4A148C','#009688','#00ACC1','#0091EA','#8BC34A','#558B2F'
}

func (tag *Tag) String() string {
	return ToString(tag)
}

// ThresholdProfile help the alarms engine to decide if a specific resource has to be declared critical or down
type ThresholdProfile struct {
	_                      struct{}                 `type:"structure"` // Enforces key based initialization.
	ProfileID              string                   `json:"profile_id,omitempty"`
	Type                   string                   `json:"type"` // Denotes monitor type
	ProfileName            string                   `json:"profile_name"`
	ProfileType            int                      `json:"profile_type"` // 1 - Static Threshold or 2 - AI-based Threshold
	DownLocationThreshold  int                      `json:"down_location_threshold"`
	WebsiteContentModified bool                     `json:"website_content_modified,omitempty"`
	WebsiteContentChanges  []map[string]interface{} `json:"website_content_changes,omitempty"`
	ReadTimeOut            map[string]interface{}   `json:"read_time_out,omitempty"`
	ResponseTimeThreshold  map[string]interface{}   `json:"response_time_threshold,omitempty"`
	// SSL_CERT attributes
	SSLCertificateFingerprintModified map[string]interface{}   `json:"ssl_fingerprint_modified,omitempty"`
	SSLCertificateDaysUntilExpiry     []map[string]interface{} `json:"days_until_expiry,omitempty"`

	// CRON attributes
	CronNoRunAlert    map[string]interface{} `json:"cron_no_run_alert,omitempty"`
	CronDurationAlert map[string]interface{} `json:"cron_duration_alert,omitempty"`

	// HEARTBEAT attributes
	TroubleIfNotPingedMoreThan map[string]interface{} `json:"hb_availability1,omitempty"`
	DownIfNotPingedMoreThan    map[string]interface{} `json:"hb_availability2,omitempty"`
	TroubleIfPingedWithin      map[string]interface{} `json:"hb_availability3,omitempty"`
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
			typeOfContentModified := reflect.TypeOf(v).Kind()
			if typeOfContentModified == reflect.Map {
				contentModifiedMap := v.(map[string]interface{})
				thresholdProfile.WebsiteContentModified, _ = contentModifiedMap["value"].(bool)
			} else {
				thresholdProfile.WebsiteContentModified, _ = v.(bool)
			}
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
		} else if k == "read_time_out" {
			thresholdProfile.ReadTimeOut = v.(map[string]interface{})
		} else if k == "cron_no_run_alert" {
			thresholdProfile.CronNoRunAlert = v.(map[string]interface{})
		} else if k == "cron_duration_alert" {
			thresholdProfile.CronDurationAlert = v.(map[string]interface{})
		}
	}
	return nil
}

// NotificationProfile allows tweaking when alerts have to be sent out.
type NotificationProfile struct {
	_                              struct{}                 `type:"structure"` // Enforces key based initialization.
	ProfileID                      string                   `json:"profile_id,omitempty"`
	ProfileName                    string                   `json:"profile_name"`
	RcaNeeded                      bool                     `json:"rca_needed"`
	NotifyAfterExecutingActions    bool                     `json:"notify_after_executing_actions"`
	TemplateID                     string                   `json:"template_id,omitempty"`
	SuppressAutomation             bool                     `json:"suppress_automation"`
	AlertConfiguration             []map[string]interface{} `json:"alert_configuration,omitempty"`
	NotificationDelayConfiguration []map[string]interface{} `json:"notification_delay_configuration,omitempty"`
	PersistentAlertConfiguration   []map[string]interface{} `json:"persistent_alert_configuration,omitempty"`
	EscalationConfiguration        map[string]interface{}   `json:"escalation_configuration,omitempty"`
}

func (notificationProfile *NotificationProfile) String() string {
	return ToString(notificationProfile)
}

// LocationProfile make it convenient to set monitoring locations consistently across many websites or monitors
type LocationProfile struct {
	_                                struct{} `type:"structure"` // Enforces key based initialization.
	ProfileID                        string   `json:"profile_id,omitempty"`
	ProfileName                      string   `json:"profile_name"`
	PrimaryLocation                  string   `json:"primary_location"`
	SecondaryLocations               []string `json:"secondary_locations"`
	RestrictAlternateLocationPolling bool     `json:"restrict_alt_loc"`
	LocationConsentForOuterRegions   bool     `json:"outer_regions_location_consent"`
}

func (locationProfile *LocationProfile) String() string {
	return ToString(locationProfile)
}

// Schedule a maintenance window to collaborate effectively within the IT team. It prevents redundant alerts from being triggered.
type ScheduleMaintenance struct {
	_                 struct{}     `type:"structure"` // Enforces key based initialization.
	MaintenanceID     string       `json:"maintenance_id,omitempty"`
	DisplayName       string       `json:"display_name"`
	Description       string       `json:"description"`
	MaintenanceType   int          `json:"maintenance_type"`
	StartTime         string       `json:"start_time"`
	TimeZone          string       `json:"timezone"`
	EndTime           string       `json:"end_time"`
	StartDate         string       `json:"start_date"`
	EndDate           string       `json:"end_date"`
	PerformMonitoring bool         `json:"perform_monitoring"`
	SelectionType     ResourceType `json:"selection_type"`
	Monitors          []string     `json:"monitors,omitempty"`
	MonitorGroups     []string     `json:"monitor_groups,omitempty"`
	Tags              []string     `json:"tags,omitempty"`
}

func (scheduleMaintenance *ScheduleMaintenance) String() string {
	return ToString(scheduleMaintenance)
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
	_           struct{} `type:"structure"` // Enforces key based initialization.
	LocationID  string   `json:"location_id"`
	CountryName string   `json:"country_name"`
	DisplayName string   `json:"display_name"`
	UseIPV6     bool     `json:"use_ipv6"`
	CityName    string   `json:"city_name"`
	CityShort   string   `json:"city_short"`
	Continent   string   `json:"continent"`
}

// Setup other users who can login to Site24x7 and receive instant notifications about outages.
type User struct {
	_                     struct{}               `type:"structure"` // Enforces key based initialization.
	ID                    string                 `json:"user_id,omitempty"`
	DisplayName           string                 `json:"display_name"`
	Email                 string                 `json:"email_address"`
	SelectionType         ResourceType           `json:"selection_type"`
	UserRole              int                    `json:"user_role"`
	StatusIQRole          int                    `json:"statusiq_role,omitempty"`
	CloudSpendRole        int                    `json:"cloudspend_role,omitempty"`
	JobTitle              int                    `json:"job_title,omitempty"`
	MobileSettings        map[string]interface{} `json:"mobile_settings,omitempty"`
	AlertSettings         map[string]interface{} `json:"alert_settings"`
	TwitterSettings       map[string]interface{} `json:"twitter_settings,omitempty"`
	IsEditAllowed         bool                   `json:"is_edit_allowed,omitempty"`
	IsClientPortalUser    bool                   `json:"is_client_portal_user,omitempty"`
	IsAccountContact      bool                   `json:"is_account_contact,omitempty"`
	IsContact             bool                   `json:"is_contact,omitempty"`
	IsInvited             bool                   `json:"is_invited,omitempty"`
	SubscribeNewsletter   bool                   `json:"subscribe_newsletter,omitempty"`
	UserGroupIDs          []string               `json:"user_groups,omitempty"`
	NotificationMedium    []int                  `json:"notify_medium"`
	Monitors              []string               `json:"monitors,omitempty"`
	MonitorGroups         []string               `json:"monitor_groups,omitempty"`
	ConsentForNonEUAlerts bool                   `json:"consent_for_non_eu_alerts"`
}

func (user *User) String() string {
	return ToString(user)
}

// UserGroup help organize individuals so that they receive alerts and reports based on their responsibility.
type UserGroup struct {
	_                struct{} `type:"structure"` // Enforces key based initialization.
	UserGroupID      string   `json:"user_group_id,omitempty"`
	DisplayName      string   `json:"display_name"`
	Users            []string `json:"users"`
	AttributeGroupID string   `json:"attribute_group_id"`
	ProductID        int      `json:"product_id"`
}

func (userGroup *UserGroup) String() string {
	return ToString(userGroup)
}

// URLAction prioritize and remediate routine actions automatically,
// increase IT efficiency and streamline your processes to reduce performance degrade
type URLAction struct {
	_                      struct{} `type:"structure"` // Enforces key based initialization.
	ActionID               string   `json:"action_id,omitempty"`
	ActionType             int      `json:"action_type"`
	ActionName             string   `json:"action_name"`
	ActionUrl              string   `json:"action_url"`
	ActionTimeout          int      `json:"action_timeout"`
	ActionMethod           string   `json:"action_method"`
	SuppressAlert          bool     `json:"suppress_alert,omitempty"`
	SendIncidentParameters bool     `json:"send_incident_parameters"`
	SendCustomParameters   bool     `json:"send_custom_parameters"`
	CustomParameters       string   `json:"custom_parameters"`
	SendInJsonFormat       bool     `json:"send_in_json_format"`
	SendEmail              bool     `json:"send_mail"`
	AuthMethod             string   `json:"auth_method,omitempty"`
	Username               string   `json:"username,omitempty"`
	Password               string   `json:"password,omitempty"`
	OAuth2Provider         string   `json:"oauth2_provider,omitempty"`
	UserAgent              string   `json:"user_agent,omitempty"`
}

func (urlAction *URLAction) String() string {
	return ToString(urlAction)
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

// Denotes a MSP customer
type MSPCustomer struct {
	Name   string `json:"name"`
	ZAAID  string `json:"zaaid"`
	UserID string `json:"user_id"`
}

func (mspCustomer *MSPCustomer) String() string {
	return ToString(mspCustomer)
}

// Denotes a AWS external ID
type AWSExternalID struct {
	ID string `json:"external_id"`
}

func (awsExternalID *AWSExternalID) String() string {
	return ToString(awsExternalID)
}

// Denotes a Device Key
type DeviceKey struct {
	ID string `json:"device_key"`
}

func (deviceKey *DeviceKey) String() string {
	return ToString(deviceKey)
}

type CredentialProfile struct {
	_              struct{} `type:"structure"` // Enforces key based initialization.
	ID             string   `json:"credential_profile_id"`
	CredentialType int      `json:"credential_type"`
	CredentialName string   `json:"credential_name"`
	UserName       string   `json:"username"`
	Password       string   `json:"password"`
}

func (credentialProfile *CredentialProfile) String() string {
	return ToString(credentialProfile)
}

type BusinessHour struct {
	_           struct{}   `type:"structure"` // Enforces key based initialization.
	ID          string     `json:"business_hours_id,omitempty"`
	DisplayName string     `json:"display_name"`
	Description string     `json:"description"`
	TimeConfig  []TimeSlot `json:"time_config"`
}
type TimeSlot struct {
	Day       int    `json:"day"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// String representation for debugging
func (bh *BusinessHour) String() string {
	return ToString(bh)
}

type Customer struct {
	_               struct{} `type:"structure"` // Enforces key-based initialization.
	UserID          string   `json:"user_id,omitempty"`
	CountryCode     string   `json:"country_code,omitempty"`
	Timezone        string   `json:"timezone,omitempty"`
	LanguageCode    string   `json:"language_code,omitempty"`
	Industry        string   `json:"industry,omitempty"`
	RoleTitle       string   `json:"roletitle,omitempty"`
	Invite          bool     `json:"invite,omitempty"`
	CustomerGroups  []string `json:"customer_groups,omitempty"`
	Digest          string   `json:"digest,omitempty"`
	Zuids           []string `json:"zuids,omitempty"`
	CustomerCompany string   `json:"customer_company,omitempty"`
	DisplayName     string   `json:"display_name,omitempty"`
	CustomerWebsite string   `json:"customer_website,omitempty"`
	EmailAddress    string   `json:"email_address,omitempty"`
	PortalName      string   `json:"portal_name,omitempty"`
	Captcha         string   `json:"captcha,omitempty"`
	Zaaid           string   `json:"zaaid,omitempty"`
}

func (c *Customer) String() string {
	return ToString(c)
}
