package api

type Site24x7Monitor interface {
	SetLocationProfileID(notificationProfileID string)
	GetLocationProfileID() string
	SetNotificationProfileID(notificationProfileID string)
	GetNotificationProfileID() string
	SetUserGroupIDs(userGroupIDs []string)
	GetUserGroupIDs() []string
	SetTagIDs(tagIDs []string)
	GetTagIDs() []string
	String() string
}

// Generic type for denoting a resource in Site24x7.
type GenericMonitor struct {
	_                     struct{} `type:"structure"` // Enforces key based initialization.
	MonitorID             string   `json:"monitor_id,omitempty"`
	DisplayName           string   `json:"display_name"`
	Type                  string   `json:"type"`
	LocationProfileID     string   `json:"location_profile_id"`
	NotificationProfileID string   `json:"notification_profile_id"`
	ThresholdProfileID    string   `json:"threshold_profile_id"`
	MonitorGroups         []string `json:"monitor_groups,omitempty"`
	UserGroupIDs          []string `json:"user_group_ids,omitempty"`
	TagIDs                []string `json:"tag_ids,omitempty"`
	ThirdPartyServiceIDs  []string `json:"third_party_services,omitempty"`
}

func (monitor *GenericMonitor) SetNotificationProfileID(notificationProfileID string) {
	monitor.NotificationProfileID = notificationProfileID
}

func (monitor *GenericMonitor) GetNotificationProfileID() string {
	return monitor.NotificationProfileID
}

func (monitor *GenericMonitor) SetUserGroupIDs(userGroupIDs []string) {
	monitor.UserGroupIDs = userGroupIDs
}

func (monitor *GenericMonitor) GetUserGroupIDs() []string {
	return monitor.UserGroupIDs
}

func (monitor *GenericMonitor) SetTagIDs(tagIDs []string) {
	monitor.TagIDs = tagIDs
}

func (monitor *GenericMonitor) GetTagIDs() []string {
	return monitor.TagIDs
}

func (monitor *GenericMonitor) String() string {
	return ToString(monitor)
}

// Denotes the website monitor resource in Site24x7.
type WebsiteMonitor struct {
	_              struct{} `type:"structure"` // Enforces key based initialization.
	MonitorID      string   `json:"monitor_id,omitempty"`
	DisplayName    string   `json:"display_name"`
	Type           string   `json:"type"`
	Website        string   `json:"website"`
	CheckFrequency string   `json:"check_frequency"`
	Timeout        int      `json:"timeout"`
	UseIPV6        bool     `json:"use_ipv6"`
	// HTTP Configuration
	HTTPMethod                string   `json:"http_method"`
	RequestContentType        string   `json:"request_content_type,omitempty"`
	RequestBody               string   `json:"request_param,omitempty"`
	RequestHeaders            []Header `json:"custom_headers,omitempty"`
	UserAgent                 string   `json:"user_agent,omitempty"`
	AuthMethod                string   `json:"auth_method,omitempty"`
	AuthUser                  string   `json:"auth_user,omitempty"`
	AuthPass                  string   `json:"auth_pass,omitempty"`
	CredentialProfileID       string   `json:"credential_profile_id,omitempty"`
	ClientCertificatePassword string   `json:"client_certificate_password,omitempty"`
	UseNameServer             bool     `json:"use_name_server,omitempty"`
	ForcedIPs                 string   `json:"forced_ips,omitempty"`
	UpStatusCodes             string   `json:"up_status_codes,omitempty"`
	FollowHTTPRedirection     bool     `json:"follow_redirect"`
	SSLProtocol               string   `json:"ssl_protocol,omitempty"`
	HTTPProtocol              string   `json:"http_protocol,omitempty"`
	UseAlpn                   bool     `json:"use_alpn"`
	// Content Check
	MatchingKeyword   *ValueAndSeverity  `json:"matching_keyword,omitempty"`
	UnmatchingKeyword *ValueAndSeverity  `json:"unmatching_keyword,omitempty"`
	MatchCase         bool               `json:"match_case"`
	MatchRegex        *ValueAndSeverity  `json:"match_regex,omitempty"`
	ResponseHeaders   HTTPResponseHeader `json:"response_headers_check,omitempty"`
	// Configuration Profiles
	LocationProfileID     string      `json:"location_profile_id"`
	NotificationProfileID string      `json:"notification_profile_id"`
	ThresholdProfileID    string      `json:"threshold_profile_id"`
	MonitorGroups         []string    `json:"monitor_groups,omitempty"`
	DependencyResourceIDs []string    `json:"dependency_resource_ids,omitempty"`
	UserGroupIDs          []string    `json:"user_group_ids,omitempty"`
	TagIDs                []string    `json:"tag_ids,omitempty"`
	ThirdPartyServiceIDs  []string    `json:"third_party_services,omitempty"`
	ActionIDs             []ActionRef `json:"action_ids,omitempty"`
}

func (websiteMonitor *WebsiteMonitor) SetLocationProfileID(locationProfileID string) {
	websiteMonitor.LocationProfileID = locationProfileID
}

func (websiteMonitor *WebsiteMonitor) GetLocationProfileID() string {
	return websiteMonitor.LocationProfileID
}

func (websiteMonitor *WebsiteMonitor) SetNotificationProfileID(notificationProfileID string) {
	websiteMonitor.NotificationProfileID = notificationProfileID
}

func (websiteMonitor *WebsiteMonitor) GetNotificationProfileID() string {
	return websiteMonitor.NotificationProfileID
}

func (websiteMonitor *WebsiteMonitor) SetUserGroupIDs(userGroupIDs []string) {
	websiteMonitor.UserGroupIDs = userGroupIDs
}

func (websiteMonitor *WebsiteMonitor) GetUserGroupIDs() []string {
	return websiteMonitor.UserGroupIDs
}

func (websiteMonitor *WebsiteMonitor) SetTagIDs(tagIDs []string) {
	websiteMonitor.TagIDs = tagIDs
}

func (websiteMonitor *WebsiteMonitor) GetTagIDs() []string {
	return websiteMonitor.TagIDs
}

func (websiteMonitor *WebsiteMonitor) String() string {
	return ToString(websiteMonitor)
}

// Denotes the web page speed monitor resource in Site24x7.
type WebPageSpeedMonitor struct {
	_              struct{} `type:"structure"` // Enforces key based initialization.
	MonitorID      string   `json:"monitor_id,omitempty"`
	DisplayName    string   `json:"display_name"`
	Type           string   `json:"type"`
	Website        string   `json:"website"`
	CheckFrequency string   `json:"check_frequency"`
	Timeout        int      `json:"timeout"`
	UseIPV6        bool     `json:"use_ipv6"`
	WebsiteType    int      `json:"website_type"`
	BrowserType    int      `json:"browser_type"`
	BrowserVersion int      `json:"browser_version"`
	DeviceType     string   `json:"device_type"`
	WPAResolution  string   `json:"wpa_resolution"`
	// HTTP Configuration
	HTTPMethod    string   `json:"http_method"`
	CustomHeaders []Header `json:"custom_headers,omitempty"`
	AuthUser      string   `json:"auth_user,omitempty"`
	AuthPass      string   `json:"auth_pass,omitempty"`
	UserAgent     string   `json:"user_agent,omitempty"`
	UpStatusCodes string   `json:"up_status_codes,omitempty"`
	// Content Check
	MatchingKeyword   *ValueAndSeverity `json:"matching_keyword,omitempty"`
	UnmatchingKeyword *ValueAndSeverity `json:"unmatching_keyword,omitempty"`
	MatchRegex        *ValueAndSeverity `json:"match_regex,omitempty"`
	MatchCase         bool              `json:"match_case"`

	LocationProfileID     string      `json:"location_profile_id"`
	NotificationProfileID string      `json:"notification_profile_id"`
	ThresholdProfileID    string      `json:"threshold_profile_id"`
	MonitorGroups         []string    `json:"monitor_groups,omitempty"`
	DependencyResourceIDs []string    `json:"dependency_resource_ids,omitempty"`
	UserGroupIDs          []string    `json:"user_group_ids,omitempty"`
	TagIDs                []string    `json:"tag_ids,omitempty"`
	ThirdPartyServiceIDs  []string    `json:"third_party_services,omitempty"`
	ActionIDs             []ActionRef `json:"action_ids,omitempty"`
}

func (webPageSpeedMonitor *WebPageSpeedMonitor) SetLocationProfileID(locationProfileID string) {
	webPageSpeedMonitor.LocationProfileID = locationProfileID
}

func (webPageSpeedMonitor *WebPageSpeedMonitor) GetLocationProfileID() string {
	return webPageSpeedMonitor.LocationProfileID
}

func (webPageSpeedMonitor *WebPageSpeedMonitor) SetNotificationProfileID(notificationProfileID string) {
	webPageSpeedMonitor.NotificationProfileID = notificationProfileID
}

func (webPageSpeedMonitor *WebPageSpeedMonitor) GetNotificationProfileID() string {
	return webPageSpeedMonitor.NotificationProfileID
}

func (webPageSpeedMonitor *WebPageSpeedMonitor) SetUserGroupIDs(userGroupIDs []string) {
	webPageSpeedMonitor.UserGroupIDs = userGroupIDs
}

func (webPageSpeedMonitor *WebPageSpeedMonitor) GetUserGroupIDs() []string {
	return webPageSpeedMonitor.UserGroupIDs
}

func (webPageSpeedMonitor *WebPageSpeedMonitor) SetTagIDs(tagIDs []string) {
	webPageSpeedMonitor.TagIDs = tagIDs
}

func (webPageSpeedMonitor *WebPageSpeedMonitor) GetTagIDs() []string {
	return webPageSpeedMonitor.TagIDs
}

func (webPageSpeedMonitor *WebPageSpeedMonitor) String() string {
	return ToString(webPageSpeedMonitor)
}

// Denotes the SSL monitor resource in Site24x7.
type SSLMonitor struct {
	_                     struct{}    `type:"structure"` // Enforces key based initialization.
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
	DependencyResourceIDs []string    `json:"dependency_resource_ids,omitempty"`
	UserGroupIDs          []string    `json:"user_group_ids,omitempty"`
	TagIDs                []string    `json:"tag_ids,omitempty"`
	ThirdPartyServiceIDs  []string    `json:"third_party_services,omitempty"`
	// ActionIDs             []ActionRef `json:"action_ids,omitempty"`
}

func (sslMonitor *SSLMonitor) SetLocationProfileID(locationProfileID string) {
	sslMonitor.LocationProfileID = locationProfileID
}

func (sslMonitor *SSLMonitor) GetLocationProfileID() string {
	return sslMonitor.LocationProfileID
}

func (sslMonitor *SSLMonitor) SetNotificationProfileID(notificationProfileID string) {
	sslMonitor.NotificationProfileID = notificationProfileID
}

func (sslMonitor *SSLMonitor) GetNotificationProfileID() string {
	return sslMonitor.NotificationProfileID
}

func (sslMonitor *SSLMonitor) SetUserGroupIDs(userGroupIDs []string) {
	sslMonitor.UserGroupIDs = userGroupIDs
}

func (sslMonitor *SSLMonitor) GetUserGroupIDs() []string {
	return sslMonitor.UserGroupIDs
}

func (sslMonitor *SSLMonitor) SetTagIDs(tagIDs []string) {
	sslMonitor.TagIDs = tagIDs
}

func (sslMonitor *SSLMonitor) GetTagIDs() []string {
	return sslMonitor.TagIDs
}

func (sslMonitor *SSLMonitor) String() string {
	return ToString(sslMonitor)
}

// Denotes the REST API monitor resource in Site24x7.
type RestApiMonitor struct {
	_              struct{} `type:"structure"` // Enforces key based initialization.
	MonitorID      string   `json:"monitor_id,omitempty"`
	DisplayName    string   `json:"display_name"`
	Type           string   `json:"type"`
	Website        string   `json:"website"`
	CheckFrequency string   `json:"check_frequency"`
	Timeout        int      `json:"timeout"`
	UseIPV6        bool     `json:"use_ipv6"`
	// HTTP Configuration
	HTTPMethod                string                 `json:"http_method"`
	RequestContentType        string                 `json:"request_content_type,omitempty"`
	RequestBody               string                 `json:"request_param,omitempty"`
	RequestHeaders            []Header               `json:"custom_headers,omitempty"`
	GraphQL                   map[string]interface{} `json:"graphql,omitempty"`
	UserAgent                 string                 `json:"user_agent,omitempty"`
	AuthMethod                string                 `json:"auth_method,omitempty"`
	AuthUser                  string                 `json:"auth_user,omitempty"`
	AuthPass                  string                 `json:"auth_pass,omitempty"`
	CredentialProfileID       string                 `json:"credential_profile_id,omitempty"`
	OAuth2Provider            string                 `json:"oauth2_provider,omitempty"`
	ClientCertificatePassword string                 `json:"client_certificate_password,omitempty"`
	JwtID                     string                 `json:"jwt_id,omitempty"`
	UseNameServer             bool                   `json:"use_name_server"`
	HTTPProtocol              string                 `json:"http_protocol,omitempty"`
	SSLProtocol               string                 `json:"ssl_protocol,omitempty"`
	UpStatusCodes             string                 `json:"up_status_codes,omitempty"`
	UseAlpn                   bool                   `json:"use_alpn"`
	// Content Check
	ResponseContentType string                 `json:"response_type"`
	MatchJSON           map[string]interface{} `json:"match_json,omitempty"`
	JSONSchema          map[string]interface{} `json:"json_schema,omitempty"`
	JSONSchemaCheck     bool                   `json:"json_schema_check,omitempty"`
	MatchingKeyword     map[string]interface{} `json:"matching_keyword,omitempty"`
	UnmatchingKeyword   map[string]interface{} `json:"unmatching_keyword,omitempty"`
	MatchCase           bool                   `json:"match_case"`
	MatchRegex          map[string]interface{} `json:"match_regex,omitempty"`
	ResponseHeaders     HTTPResponseHeader     `json:"response_headers_check,omitempty"`
	// Configuration Profiles
	LocationProfileID     string      `json:"location_profile_id"`
	NotificationProfileID string      `json:"notification_profile_id"`
	ThresholdProfileID    string      `json:"threshold_profile_id"`
	MonitorGroups         []string    `json:"monitor_groups,omitempty"`
	DependencyResourceIDs []string    `json:"dependency_resource_ids,omitempty"`
	UserGroupIDs          []string    `json:"user_group_ids,omitempty"`
	TagIDs                []string    `json:"tag_ids,omitempty"`
	ThirdPartyServiceIDs  []string    `json:"third_party_services,omitempty"`
	ActionIDs             []ActionRef `json:"action_ids,omitempty"`
}

func (restApiMonitor *RestApiMonitor) SetLocationProfileID(locationProfileID string) {
	restApiMonitor.LocationProfileID = locationProfileID
}

func (restApiMonitor *RestApiMonitor) GetLocationProfileID() string {
	return restApiMonitor.LocationProfileID
}

func (restApiMonitor *RestApiMonitor) SetNotificationProfileID(notificationProfileID string) {
	restApiMonitor.NotificationProfileID = notificationProfileID
}

func (restApiMonitor *RestApiMonitor) GetNotificationProfileID() string {
	return restApiMonitor.NotificationProfileID
}

func (restApiMonitor *RestApiMonitor) SetUserGroupIDs(userGroupIDs []string) {
	restApiMonitor.UserGroupIDs = userGroupIDs
}

func (restApiMonitor *RestApiMonitor) GetUserGroupIDs() []string {
	return restApiMonitor.UserGroupIDs
}

func (restApiMonitor *RestApiMonitor) SetTagIDs(tagIDs []string) {
	restApiMonitor.TagIDs = tagIDs
}

func (restApiMonitor *RestApiMonitor) GetTagIDs() []string {
	return restApiMonitor.TagIDs
}

func (restApiMonitor *RestApiMonitor) String() string {
	return ToString(restApiMonitor)
}

// Denotes the REST API Transaction monitor resource in Site24x7.
type RestApiTransactionMonitor struct {
	_              struct{} `type:"structure"` // Enforces key based initialization.
	MonitorID      string   `json:"monitor_id,omitempty"`
	DisplayName    string   `json:"display_name"`
	Type           string   `json:"type"`
	CheckFrequency string   `json:"check_frequency"`
	Steps          []Steps  `json:"steps"`
	// Configuration Profiles
	LocationProfileID     string      `json:"location_profile_id"`
	NotificationProfileID string      `json:"notification_profile_id"`
	ThresholdProfileID    string      `json:"threshold_profile_id"`
	MonitorGroups         []string    `json:"monitor_groups,omitempty"`
	DependencyResourceIDs []string    `json:"dependency_resource_ids,omitempty"`
	UserGroupIDs          []string    `json:"user_group_ids,omitempty"`
	TagIDs                []string    `json:"tag_ids,omitempty"`
	ThirdPartyServiceIDs  []string    `json:"third_party_services,omitempty"`
	ActionIDs             []ActionRef `json:"action_ids,omitempty"`
}

func (restApiTransactionMonitor *RestApiTransactionMonitor) SetLocationProfileID(locationProfileID string) {
	restApiTransactionMonitor.LocationProfileID = locationProfileID
}

func (restApiTransactionMonitor *RestApiTransactionMonitor) GetLocationProfileID() string {
	return restApiTransactionMonitor.LocationProfileID
}

func (restApiTransactionMonitor *RestApiTransactionMonitor) SetNotificationProfileID(notificationProfileID string) {
	restApiTransactionMonitor.NotificationProfileID = notificationProfileID
}

func (restApiTransactionMonitor *RestApiTransactionMonitor) GetNotificationProfileID() string {
	return restApiTransactionMonitor.NotificationProfileID
}

func (restApiTransactionMonitor *RestApiTransactionMonitor) SetUserGroupIDs(userGroupIDs []string) {
	restApiTransactionMonitor.UserGroupIDs = userGroupIDs
}

func (restApiTransactionMonitor *RestApiTransactionMonitor) GetUserGroupIDs() []string {
	return restApiTransactionMonitor.UserGroupIDs
}

func (restApiTransactionMonitor *RestApiTransactionMonitor) SetTagIDs(tagIDs []string) {
	restApiTransactionMonitor.TagIDs = tagIDs
}

func (restApiTransactionMonitor *RestApiTransactionMonitor) GetTagIDs() []string {
	return restApiTransactionMonitor.TagIDs
}

func (restApiTransactionMonitor *RestApiTransactionMonitor) String() string {
	return ToString(restApiTransactionMonitor)
}

// Denotes the Amazon monitor resource in Site24x7.
type AmazonMonitor struct {
	_                     struct{} `type:"structure"` // Enforces key based initialization.
	MonitorID             string   `json:"monitor_id,omitempty"`
	DisplayName           string   `json:"display_name"`
	Type                  string   `json:"type"`
	AWSExternalID         string   `json:"aws_external_id"`
	RoleARN               string   `json:"role_arn"`
	DiscoverFrequency     int      `json:"aws_discovery_frequency"`
	DiscoverServices      []string `json:"aws_discover_services"`
	NotificationProfileID string   `json:"notification_profile_id"`
	UserGroupIDs          []string `json:"user_group_ids"`
	TagIDs                []string `json:"tag_ids,omitempty"`
	ThirdPartyServiceIDs  []string `json:"third_party_services,omitempty"`
}

func (amazonMonitor *AmazonMonitor) SetLocationProfileID(locationProfileID string) {
}

func (amazonMonitor *AmazonMonitor) GetLocationProfileID() string {
	return ""
}

func (amazonMonitor *AmazonMonitor) SetNotificationProfileID(notificationProfileID string) {
	amazonMonitor.NotificationProfileID = notificationProfileID
}

func (amazonMonitor *AmazonMonitor) GetNotificationProfileID() string {
	return amazonMonitor.NotificationProfileID
}

func (amazonMonitor *AmazonMonitor) SetUserGroupIDs(userGroupIDs []string) {
	amazonMonitor.UserGroupIDs = userGroupIDs
}

func (amazonMonitor *AmazonMonitor) GetUserGroupIDs() []string {
	return amazonMonitor.UserGroupIDs
}

func (amazonMonitor *AmazonMonitor) SetTagIDs(tagIDs []string) {
	amazonMonitor.TagIDs = tagIDs
}

func (amazonMonitor *AmazonMonitor) GetTagIDs() []string {
	return amazonMonitor.TagIDs
}

func (amazonMonitor *AmazonMonitor) String() string {
	return ToString(amazonMonitor)
}

// Denotes the server monitor resource in Site24x7.
type ServerMonitor struct {
	_                     struct{} `type:"structure"` // Enforces key based initialization.
	MonitorID             string   `json:"monitor_id,omitempty"`
	DisplayName           string   `json:"display_name"`
	HostName              string   `json:"hostname"`
	IPAddress             string   `json:"ipaddress"`
	Type                  string   `json:"type"`
	TemplateID            string   `json:"templateid"`
	PollInterval          int      `json:"sm_poll_interval"`
	ITAutomationModule    bool     `json:"server_setting_it_aut"`
	PluginModule          bool     `json:"server_setting_plugins"`
	LogNeeded             bool     `json:"log_needed"`
	PerformAutomation     bool     `json:"perform_automation"`
	NotificationProfileID string   `json:"notification_profile_id"`
	ThresholdProfileID    string   `json:"threshold_profile_id"`
	ResourceProfileID     string   `json:"resource_profile_id"`
	MonitorGroups         []string `json:"monitor_groups,omitempty"`
	UserGroupIDs          []string `json:"user_group_ids,omitempty"`
	TagIDs                []string `json:"tag_ids,omitempty"`
	ThirdPartyServiceIDs  []string `json:"third_party_services,omitempty"`
	// ActionIDs             []ActionRef `json:"action_ids,omitempty"`
}

func (serverMonitor *ServerMonitor) SetLocationProfileID(locationProfileID string) {
}

func (serverMonitor *ServerMonitor) GetLocationProfileID() string {
	return ""
}

func (serverMonitor *ServerMonitor) SetNotificationProfileID(notificationProfileID string) {
	serverMonitor.NotificationProfileID = notificationProfileID
}

func (serverMonitor *ServerMonitor) GetNotificationProfileID() string {
	return serverMonitor.NotificationProfileID
}

func (serverMonitor *ServerMonitor) SetUserGroupIDs(userGroupIDs []string) {
	serverMonitor.UserGroupIDs = userGroupIDs
}

func (serverMonitor *ServerMonitor) GetUserGroupIDs() []string {
	return serverMonitor.UserGroupIDs
}

func (serverMonitor *ServerMonitor) SetTagIDs(tagIDs []string) {
	serverMonitor.TagIDs = tagIDs
}

func (serverMonitor *ServerMonitor) GetTagIDs() []string {
	return serverMonitor.TagIDs
}

func (serverMonitor *ServerMonitor) String() string {
	return ToString(serverMonitor)
}

// Denotes the Heartbeat monitor resource in Site24x7.
type HeartbeatMonitor struct {
	_                     struct{} `type:"structure"` // Enforces key based initialization.
	MonitorID             string   `json:"monitor_id,omitempty"`
	DisplayName           string   `json:"display_name"`
	NameInPingURL         string   `json:"unique_name"`
	Type                  string   `json:"type"`
	ThresholdProfileID    string   `json:"threshold_profile_id"`
	NotificationProfileID string   `json:"notification_profile_id"`
	MonitorGroups         []string `json:"monitor_groups,omitempty"`
	TagIDs                []string `json:"tag_ids,omitempty"`
	ThirdPartyServiceIDs  []string `json:"third_party_services,omitempty"`
	UserGroupIDs          []string `json:"user_group_ids,omitempty"`
	OnCallScheduleID      string   `json:"on_call_schedule_id,omitempty"`
}

func (heartbeatMonitor *HeartbeatMonitor) SetLocationProfileID(locationProfileID string) {
}

func (heartbeatMonitor *HeartbeatMonitor) GetLocationProfileID() string {
	return ""
}

func (heartbeatMonitor *HeartbeatMonitor) SetNotificationProfileID(notificationProfileID string) {
	heartbeatMonitor.NotificationProfileID = notificationProfileID
}

func (heartbeatMonitor *HeartbeatMonitor) GetNotificationProfileID() string {
	return heartbeatMonitor.NotificationProfileID
}

func (heartbeatMonitor *HeartbeatMonitor) SetUserGroupIDs(userGroupIDs []string) {
	heartbeatMonitor.UserGroupIDs = userGroupIDs
}

func (heartbeatMonitor *HeartbeatMonitor) GetUserGroupIDs() []string {
	return heartbeatMonitor.UserGroupIDs
}

func (heartbeatMonitor *HeartbeatMonitor) SetTagIDs(tagIDs []string) {
	heartbeatMonitor.TagIDs = tagIDs
}

func (heartbeatMonitor *HeartbeatMonitor) GetTagIDs() []string {
	return heartbeatMonitor.TagIDs
}

func (heartbeatMonitor *HeartbeatMonitor) String() string {
	return ToString(heartbeatMonitor)
}

// Denotes the DNS Server monitor resource in Site24x7.
type DNSServerMonitor struct {
	_                struct{}       `type:"structure"` // Enforces key based initialization.
	MonitorID        string         `json:"monitor_id,omitempty"`
	DisplayName      string         `json:"display_name"`
	DomainName       string         `json:"domain_name"`
	Type             string         `json:"type"`
	DNSHost          string         `json:"dns_host"`
	DNSPort          string         `json:"dns_port"`
	UseIPV6          bool           `json:"use_ipv6"`
	CheckFrequency   string         `json:"check_frequency"`
	Timeout          int            `json:"timeout"`
	LookupType       int            `json:"lookup_type,omitempty"`
	DNSSEC           bool           `json:"dnssec,omitempty"`
	DeepDiscovery    bool           `json:"deep_discovery,omitempty"`
	SearchConfig     []SearchConfig `json:"search_config,omitempty"`
	OnCallScheduleID string         `json:"on_call_schedule_id,omitempty"`

	// Configuration Profiles
	LocationProfileID     string      `json:"location_profile_id"`
	NotificationProfileID string      `json:"notification_profile_id"`
	ThresholdProfileID    string      `json:"threshold_profile_id"`
	MonitorGroups         []string    `json:"monitor_groups,omitempty"`
	DependencyResourceIDs []string    `json:"dependency_resource_ids,omitempty"`
	UserGroupIDs          []string    `json:"user_group_ids,omitempty"`
	TagIDs                []string    `json:"tag_ids,omitempty"`
	ThirdPartyServiceIDs  []string    `json:"third_party_services,omitempty"`
	ActionIDs             []ActionRef `json:"action_ids,omitempty"`
}

func (DNSServerMonitor *DNSServerMonitor) SetLocationProfileID(locationProfileID string) {
	DNSServerMonitor.LocationProfileID = locationProfileID
}

func (DNSServerMonitor *DNSServerMonitor) GetLocationProfileID() string {
	return DNSServerMonitor.LocationProfileID
}

func (DNSServerMonitor *DNSServerMonitor) SetNotificationProfileID(notificationProfileID string) {
	DNSServerMonitor.NotificationProfileID = notificationProfileID
}

func (DNSServerMonitor *DNSServerMonitor) GetNotificationProfileID() string {
	return DNSServerMonitor.NotificationProfileID
}

func (DNSServerMonitor *DNSServerMonitor) SetUserGroupIDs(userGroupIDs []string) {
	DNSServerMonitor.UserGroupIDs = userGroupIDs
}

func (DNSServerMonitor *DNSServerMonitor) GetUserGroupIDs() []string {
	return DNSServerMonitor.UserGroupIDs
}

func (DNSServerMonitor *DNSServerMonitor) SetTagIDs(tagIDs []string) {
	DNSServerMonitor.TagIDs = tagIDs
}

func (DNSServerMonitor *DNSServerMonitor) GetTagIDs() []string {
	return DNSServerMonitor.TagIDs
}

func (DNSServerMonitor *DNSServerMonitor) String() string {
	return ToString(DNSServerMonitor)
}
