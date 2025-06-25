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
	_                         struct{} `type:"structure"` // Enforces key based initialization.
	MonitorID                 string   `json:"monitor_id,omitempty"`
	DisplayName               string   `json:"display_name"`
	Type                      string   `json:"type"`
	Website                   string   `json:"website"`
	CheckFrequency            string   `json:"check_frequency"`
	Timeout                   int      `json:"timeout"`
	IPType                    int      `json:"ip_type"`
	PrimaryProtocol           int      `json:"primary_protocol,omitempty"`
	SecondaryProtocolSeverity int      `json:"secondary_protocol_severity,omitempty"`
	HiddenMonAdded            int      `json:"hidden_mon_added,omitempty"`
	UseIPV6                   bool     `json:"use_ipv6,omitempty"`
	State                     int      `json:"state"`
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
	IgnoreCertError           bool     `json:"ignore_cert_err"`
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
	IPType         int      `json:"ip_type"`
	WebsiteType    int      `json:"website_type"`
	BrowserType    int      `json:"browser_type"`
	BrowserVersion int      `json:"browser_version"`
	DeviceType     string   `json:"device_type"`
	WPAResolution  string   `json:"wpa_resolution"`
	ElementCheck   string   `json:"element_check,omitempty"`
	JwtID          string   `json:"jwt_id,omitempty"`
	// HTTP Configuration
	HTTPMethod          string   `json:"http_method"`
	CustomHeaders       []Header `json:"custom_headers,omitempty"`
	AuthUser            string   `json:"auth_user,omitempty"`
	AuthPass            string   `json:"auth_pass,omitempty"`
	CredentialProfileID string   `json:"credential_profile_id,omitempty"`
	UserAgent           string   `json:"user_agent,omitempty"`
	UpStatusCodes       string   `json:"up_status_codes,omitempty"`
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

// Denotes the web transaction browser monitor resource in Site24x7.
type WebTransactionBrowserMonitor struct {
	_                     struct{}               `type:"structure"` // Enforces key based initialization.
	MonitorID             string                 `json:"monitor_id,omitempty"`
	DisplayName           string                 `json:"display_name"`
	Type                  string                 `json:"type"`
	BaseURL               string                 `json:"base_url"`
	SeleniumScript        string                 `json:"selenium_script,omitempty"`
	ScriptType            string                 `json:"script_type,omitempty"`
	PerformAutomation     bool                   `json:"perform_automation,omitempty"`
	CheckFrequency        string                 `json:"check_frequency,omitempty"`
	AsyncDCEnabled        bool                   `json:"async_dc_enabled,omitempty"`
	BrowserType           int                    `json:"browser_type,omitempty"`
	ThinkTime             int                    `json:"think_time,omitempty"`
	IgnoreCertError       bool                   `json:"ignore_cert_err,omitempty"`
	IPType                int                    `json:"ip_type,omitempty"`
	UserAgent             string                 `json:"user_agent,omitempty"`
	BrowserVersion        int                    `json:"browser_version,omitempty"`
	PageLoadTime          int                    `json:"page_load_time,omitempty"`
	Resolution            string                 `json:"resolution,omitempty"`
	ProxyDetails          map[string]interface{} `json:"proxy_details,omitempty"`
	AuthDetails           map[string]interface{} `json:"auth_details,omitempty"`
	CustomHeaders         []Header               `json:"custom_headers,omitempty"`
	Cookies               []Header               `json:"cookies,omitempty"`
	ThresholdProfileID    string                 `json:"threshold_profile_id,omitempty"`
	LocationProfileID     string                 `json:"location_profile_id"`
	NotificationProfileID string                 `json:"notification_profile_id"`
	UserGroupIDs          []string               `json:"user_group_ids,omitempty"`
	OnCallScheduleID      string                 `json:"on_call_schedule_id,omitempty"`
	DependencyResourceIDs []string               `json:"dependency_resource_ids,omitempty"`
	MonitorGroups         []string               `json:"monitor_groups,omitempty"`
	ActionIDs             []ActionRef            `json:"action_ids,omitempty"`
	ThirdPartyServiceIDs  []string               `json:"third_party_services,omitempty"`
	TagIDs                []string               `json:"tag_ids,omitempty"`
}

func (webTransactionBrowserMonitor *WebTransactionBrowserMonitor) SetLocationProfileID(locationProfileID string) {
	webTransactionBrowserMonitor.LocationProfileID = locationProfileID
}

func (webTransactionBrowserMonitor *WebTransactionBrowserMonitor) GetLocationProfileID() string {
	return webTransactionBrowserMonitor.LocationProfileID
}

func (webTransactionBrowserMonitor *WebTransactionBrowserMonitor) SetNotificationProfileID(notificationProfileID string) {
	webTransactionBrowserMonitor.NotificationProfileID = notificationProfileID
}

func (webTransactionBrowserMonitor *WebTransactionBrowserMonitor) GetNotificationProfileID() string {
	return webTransactionBrowserMonitor.NotificationProfileID
}

func (webTransactionBrowserMonitor *WebTransactionBrowserMonitor) SetUserGroupIDs(userGroupIDs []string) {
	webTransactionBrowserMonitor.UserGroupIDs = userGroupIDs
}

func (webTransactionBrowserMonitor *WebTransactionBrowserMonitor) GetUserGroupIDs() []string {
	return webTransactionBrowserMonitor.UserGroupIDs
}

func (webTransactionBrowserMonitor *WebTransactionBrowserMonitor) SetTagIDs(tagIDs []string) {
	webTransactionBrowserMonitor.TagIDs = tagIDs
}

func (webTransactionBrowserMonitor *WebTransactionBrowserMonitor) GetTagIDs() []string {
	return webTransactionBrowserMonitor.TagIDs
}

func (webTransactionBrowserMonitor *WebTransactionBrowserMonitor) String() string {
	return ToString(webTransactionBrowserMonitor)
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

// Denotes the FTP Transfer monitor resource in Site24x7.
type FTPTransferMonitor struct {
	_                     struct{}    `type:"structure"` // Enforces key based initialization.
	MonitorID             string      `json:"monitor_id,omitempty"`
	DisplayName           string      `json:"display_name"`
	HostName              string      `json:"host_name"`
	Protocol              string      `json:"protocol"`
	Type                  string      `json:"type"`
	Port                  int         `json:"port"` // To Fix: API accepts int and returns string in response.
	CheckFrequency        string      `json:"check_frequency"`
	Timeout               int         `json:"timeout"`
	CheckUpload           bool        `json:"check_upload"`
	CheckDownload         bool        `json:"check_download"`
	Username              string      `json:"user_name"`
	Password              string      `json:"password,omitempty"`
	Destination           string      `json:"destination,omitempty"`
	LocationProfileID     string      `json:"location_profile_id"`
	MonitorGroups         []string    `json:"monitor_groups,omitempty"`
	CredentialProfileID   string      `json:"credential_profile_id,omitempty"`
	DependencyResourceIDs []string    `json:"dependency_resource_ids,omitempty"`
	ThresholdProfileID    string      `json:"threshold_profile_id"`
	TagIDs                []string    `json:"tag_ids,omitempty"`
	PerformAutomation     bool        `json:"perform_automation"`
	ActionIDs             []ActionRef `json:"action_ids,omitempty"`
	UserGroupIDs          []string    `json:"user_group_ids,omitempty"`
	OnCallScheduleID      string      `json:"on_call_schedule_id,omitempty"`
	NotificationProfileID string      `json:"notification_profile_id"`
	ThirdPartyServiceIDs  []string    `json:"third_party_services,omitempty"`
}

func (ftpTransferMonitor *FTPTransferMonitor) SetLocationProfileID(locationProfileID string) {
	ftpTransferMonitor.LocationProfileID = locationProfileID
}

func (ftpTransferMonitor *FTPTransferMonitor) GetLocationProfileID() string {
	return ftpTransferMonitor.LocationProfileID
}

func (ftpTransferMonitor *FTPTransferMonitor) SetNotificationProfileID(notificationProfileID string) {
	ftpTransferMonitor.NotificationProfileID = notificationProfileID
}

func (ftpTransferMonitor *FTPTransferMonitor) GetNotificationProfileID() string {
	return ftpTransferMonitor.NotificationProfileID
}

func (ftpTransferMonitor *FTPTransferMonitor) SetUserGroupIDs(userGroupIDs []string) {
	ftpTransferMonitor.UserGroupIDs = userGroupIDs
}

func (ftpTransferMonitor *FTPTransferMonitor) GetUserGroupIDs() []string {
	return ftpTransferMonitor.UserGroupIDs
}

func (ftpTransferMonitor *FTPTransferMonitor) String() string {
	return ToString(ftpTransferMonitor)
}

func (ftpTransferMonitor *FTPTransferMonitor) SetTagIDs(tagIDs []string) {
	ftpTransferMonitor.TagIDs = tagIDs
}

func (ftpTransferMonitor *FTPTransferMonitor) GetTagIDs() []string {
	return ftpTransferMonitor.TagIDs
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

// Denotes the ISP monitor resource in Site24x7.
type ISPMonitor struct {
	_                     struct{}    `type:"structure"` // Enforces key based initialization.
	MonitorID             string      `json:"monitor_id,omitempty"`
	DisplayName           string      `json:"display_name"`
	Hostname              string      `json:"hostname"`
	UseIPV6               bool        `json:"use_ipv6"`
	Type                  string      `json:"type"`
	Timeout               int         `json:"timeout,omitempty"`
	Protocol              string      `json:"protocol,omitempty"`
	Port                  int         `json:"port,omitempty"` // To Fix: API accepts int and returns string in response.
	CheckFrequency        string      `json:"check_frequency"`
	OnCallScheduleID      string      `json:"on_call_schedule_id,omitempty"`
	LocationProfileID     string      `json:"location_profile_id"`
	PerformAutomation     bool        `json:"perform_automation"`
	NotificationProfileID string      `json:"notification_profile_id"`
	ThresholdProfileID    string      `json:"threshold_profile_id"`
	MonitorGroups         []string    `json:"monitor_groups,omitempty"`
	DependencyResourceIDs []string    `json:"dependency_resource_ids,omitempty"`
	UserGroupIDs          []string    `json:"user_group_ids,omitempty"`
	TagIDs                []string    `json:"tag_ids,omitempty"`
	ThirdPartyServiceIDs  []string    `json:"third_party_services,omitempty"`
	ActionIDs             []ActionRef `json:"action_ids,omitempty"`
}

func (ispMonitor *ISPMonitor) SetLocationProfileID(locationProfileID string) {
	ispMonitor.LocationProfileID = locationProfileID
}

func (ispMonitor *ISPMonitor) GetLocationProfileID() string {
	return ispMonitor.LocationProfileID
}

func (ispMonitor *ISPMonitor) SetNotificationProfileID(notificationProfileID string) {
	ispMonitor.NotificationProfileID = notificationProfileID
}

func (ispMonitor *ISPMonitor) GetNotificationProfileID() string {
	return ispMonitor.NotificationProfileID
}

func (ispMonitor *ISPMonitor) SetUserGroupIDs(userGroupIDs []string) {
	ispMonitor.UserGroupIDs = userGroupIDs
}

func (ispMonitor *ISPMonitor) GetUserGroupIDs() []string {
	return ispMonitor.UserGroupIDs
}

func (ispMonitor *ISPMonitor) String() string {
	return ToString(ispMonitor)
}

func (ispMonitor *ISPMonitor) SetTagIDs(tagIDs []string) {
	ispMonitor.TagIDs = tagIDs
}

func (ispMonitor *ISPMonitor) GetTagIDs() []string {
	return ispMonitor.TagIDs
}

// Denotes the Domain Expiry monitor resource in Site24x7.
type DomainExpiryMonitor struct {
	_                     struct{}    `type:"structure"` // Enforces key based initialization.
	MonitorID             string      `json:"monitor_id,omitempty"`
	DisplayName           string      `json:"display_name"`
	Type                  string      `json:"type"`
	HostName              string      `json:"host_name"`
	DomainName            string      `json:"domain_name"`
	Port                  interface{} `json:"port"`
	Timeout               int         `json:"timeout"`
	UseIPV6               bool        `json:"use_ipv6"`
	ExpireDays            int         `json:"expire_days"`
	PerformAutomation     bool        `json:"perform_automation"`
	LocationProfileID     string      `json:"location_profile_id"`
	NotificationProfileID string      `json:"notification_profile_id"`
	UserGroupIDs          []string    `json:"user_group_ids,omitempty"`
	OnCallScheduleID      string      `json:"on_call_schedule_id,omitempty"`
	IgnoreRegistryDate    bool        `json:"ignore_registry_date"`
	MonitorGroups         []string    `json:"monitor_groups,omitempty"`
	ActionIDs             []ActionRef `json:"action_ids,omitempty"`
	ThirdPartyServiceIDs  []string    `json:"third_party_services,omitempty"`
	TagIDs                []string    `json:"tag_ids,omitempty"`
	// Content Check
	MatchingKeyword   map[string]interface{} `json:"matching_keyword,omitempty"`
	UnmatchingKeyword map[string]interface{} `json:"unmatching_keyword,omitempty"`
	MatchCase         bool                   `json:"match_case"`
	MatchRegex        map[string]interface{} `json:"match_regex,omitempty"`
}

func (domainExpiryMonitor *DomainExpiryMonitor) SetLocationProfileID(locationProfileID string) {
	domainExpiryMonitor.LocationProfileID = locationProfileID
}

func (domainExpiryMonitor *DomainExpiryMonitor) GetLocationProfileID() string {
	return domainExpiryMonitor.LocationProfileID
}

func (domainExpiryMonitor *DomainExpiryMonitor) SetNotificationProfileID(notificationProfileID string) {
	domainExpiryMonitor.NotificationProfileID = notificationProfileID
}

func (domainExpiryMonitor *DomainExpiryMonitor) GetNotificationProfileID() string {
	return domainExpiryMonitor.NotificationProfileID
}

func (domainExpiryMonitor *DomainExpiryMonitor) SetUserGroupIDs(userGroupIDs []string) {
	domainExpiryMonitor.UserGroupIDs = userGroupIDs
}

func (domainExpiryMonitor *DomainExpiryMonitor) GetUserGroupIDs() []string {
	return domainExpiryMonitor.UserGroupIDs
}

func (domainExpiryMonitor *DomainExpiryMonitor) SetTagIDs(tagIDs []string) {
	domainExpiryMonitor.TagIDs = tagIDs
}

func (domainExpiryMonitor *DomainExpiryMonitor) GetTagIDs() []string {
	return domainExpiryMonitor.TagIDs
}

func (domainExpiryMonitor *DomainExpiryMonitor) String() string {
	return ToString(domainExpiryMonitor)
}

// Denotes the Port monitor resource in Site24x7.
type PortMonitor struct {
	_                     struct{}          `type:"structure"` // Enforces key based initialization.
	MonitorID             string            `json:"monitor_id,omitempty"`
	DisplayName           string            `json:"display_name"`
	HostName              string            `json:"host_name"`
	UseIPV6               bool              `json:"use_ipv6"`
	InvertPortCheck       bool              `json:"invert_port_check,omitempty"`
	UseSSL                bool              `json:"use_ssl,omitempty"`
	Type                  string            `json:"type"`
	Timeout               int               `json:"timeout,omitempty"`
	ApplicationType       string            `json:"application_type,omitempty"`
	Command               string            `json:"command,omitempty"`
	MatchingKeyword       *ValueAndSeverity `json:"matching_keyword,omitempty"`
	UnmatchingKeyword     *ValueAndSeverity `json:"unmatching_keyword,omitempty"`
	Port                  int               `json:"port,omitempty"` // To Fix: API accepts int and returns string in response.
	PerformAutomation     bool              `json:"perform_automation"`
	CheckFrequency        string            `json:"check_frequency"`
	OnCallScheduleID      string            `json:"on_call_schedule_id,omitempty"`
	LocationProfileID     string            `json:"location_profile_id"`
	NotificationProfileID string            `json:"notification_profile_id"`
	ThresholdProfileID    string            `json:"threshold_profile_id"`
	MonitorGroups         []string          `json:"monitor_groups,omitempty"`
	DependencyResourceIDs []string          `json:"dependency_resource_ids,omitempty"`
	UserGroupIDs          []string          `json:"user_group_ids,omitempty"`
	TagIDs                []string          `json:"tag_ids,omitempty"`
	ThirdPartyServiceIDs  []string          `json:"third_party_services,omitempty"`
	ActionIDs             []ActionRef       `json:"action_ids,omitempty"`
}

func (portMonitor *PortMonitor) SetLocationProfileID(locationProfileID string) {
	portMonitor.LocationProfileID = locationProfileID
}

func (portMonitor *PortMonitor) GetLocationProfileID() string {
	return portMonitor.LocationProfileID
}

func (portMonitor *PortMonitor) SetNotificationProfileID(notificationProfileID string) {
	portMonitor.NotificationProfileID = notificationProfileID
}

func (portMonitor *PortMonitor) GetNotificationProfileID() string {
	return portMonitor.NotificationProfileID
}

func (portMonitor *PortMonitor) SetUserGroupIDs(userGroupIDs []string) {
	portMonitor.UserGroupIDs = userGroupIDs
}

func (portMonitor *PortMonitor) GetUserGroupIDs() []string {
	return portMonitor.UserGroupIDs
}

func (portMonitor *PortMonitor) String() string {
	return ToString(portMonitor)
}

func (portMonitor *PortMonitor) SetTagIDs(tagIDs []string) {
	portMonitor.TagIDs = tagIDs
}

func (portMonitor *PortMonitor) GetTagIDs() []string {
	return portMonitor.TagIDs
}

// Denotes the PING monitor resource in Site24x7.
type PINGMonitor struct {
	_                     struct{}    `type:"structure"` // Enforces key based initialization.
	MonitorID             string      `json:"monitor_id,omitempty"`
	DisplayName           string      `json:"display_name"`
	HostName              string      `json:"host_name"`
	UseIPV6               bool        `json:"use_ipv6"`
	Type                  string      `json:"type"`
	Timeout               int         `json:"timeout,omitempty"`
	CheckFrequency        string      `json:"check_frequency"`
	OnCallScheduleID      string      `json:"on_call_schedule_id,omitempty"`
	LocationProfileID     string      `json:"location_profile_id"`
	NotificationProfileID string      `json:"notification_profile_id"`
	PerformAutomation     bool        `json:"perform_automation"`
	ThresholdProfileID    string      `json:"threshold_profile_id"`
	MonitorGroups         []string    `json:"monitor_groups,omitempty"`
	DependencyResourceIDs []string    `json:"dependency_resource_ids,omitempty"`
	UserGroupIDs          []string    `json:"user_group_ids,omitempty"`
	TagIDs                []string    `json:"tag_ids,omitempty"`
	ThirdPartyServiceIDs  []string    `json:"third_party_services,omitempty"`
	ActionIDs             []ActionRef `json:"action_ids,omitempty"`
}

func (pingMonitor *PINGMonitor) SetLocationProfileID(locationProfileID string) {
	pingMonitor.LocationProfileID = locationProfileID
}

func (pingMonitor *PINGMonitor) GetLocationProfileID() string {
	return pingMonitor.LocationProfileID
}

func (pingMonitor *PINGMonitor) SetNotificationProfileID(notificationProfileID string) {
	pingMonitor.NotificationProfileID = notificationProfileID
}

func (pingMonitor *PINGMonitor) GetNotificationProfileID() string {
	return pingMonitor.NotificationProfileID
}

func (pingMonitor *PINGMonitor) SetUserGroupIDs(userGroupIDs []string) {
	pingMonitor.UserGroupIDs = userGroupIDs
}

func (pingMonitor *PINGMonitor) GetUserGroupIDs() []string {
	return pingMonitor.UserGroupIDs
}

func (pingMonitor *PINGMonitor) String() string {
	return ToString(pingMonitor)
}

func (pingMonitor *PINGMonitor) SetTagIDs(tagIDs []string) {
	pingMonitor.TagIDs = tagIDs
}

func (pingMonitor *PINGMonitor) GetTagIDs() []string {
	return pingMonitor.TagIDs
}

// Denotes the SOAP monitor resource in Site24x7.
type SOAPMonitor struct {
	_                         struct{}           `type:"structure"` // Enforces key based initialization.
	Type                      string             `json:"type"`
	MonitorID                 string             `json:"monitor_id,omitempty"`
	DisplayName               string             `json:"display_name"`
	Website                   string             `json:"website"`
	RequestParam              string             `json:"request_param"`
	SOAPAttributesSeverity    int                `json:"soap_attributes_severity"`
	SOAPAttributes            []Header           `json:"soap_attributes,omitempty"`
	ResponseHeaders           HTTPResponseHeader `json:"response_headers_check,omitempty"`
	Timeout                   int                `json:"timeout,omitempty"`
	RequestContentType        string             `json:"request_content_type"`
	HTTPMethod                string             `json:"http_method"`
	UseNameServer             bool               `json:"use_name_server"`
	HTTPProtocol              string             `json:"http_protocol"`
	UseIPV6                   bool               `json:"use_ipv6"`
	ResponseType              string             `json:"response_type"`
	CheckFrequency            string             `json:"check_frequency"`
	CredentialProfileID       string             `json:"credential_profile_id,omitempty"`
	ClientCertificatePassword string             `json:"client_certificate_password,omitempty"`
	UpStatusCodes             string             `json:"up_status_codes,omitempty"`
	OnCallScheduleID          string             `json:"on_call_schedule_id,omitempty"`
	LocationProfileID         string             `json:"location_profile_id"`
	NotificationProfileID     string             `json:"notification_profile_id"`
	PerformAutomation         bool               `json:"perform_automation"`
	ThresholdProfileID        string             `json:"threshold_profile_id"`
	SSLProtocol               string             `json:"ssl_protocol"`
	UseAlpn                   bool               `json:"use_alpn"`
	MonitorGroups             []string           `json:"monitor_groups,omitempty"`
	DependencyResourceIDs     []string           `json:"dependency_resource_ids,omitempty"`
	UserGroupIDs              []string           `json:"user_group_ids,omitempty"`
	TagIDs                    []string           `json:"tag_ids,omitempty"`
	ThirdPartyServiceIDs      []string           `json:"third_party_services,omitempty"`
	ActionIDs                 []ActionRef        `json:"action_ids,omitempty"`
}

func (soapMonitor *SOAPMonitor) SetLocationProfileID(locationProfileID string) {
	soapMonitor.LocationProfileID = locationProfileID
}

func (soapMonitor *SOAPMonitor) GetLocationProfileID() string {
	return soapMonitor.LocationProfileID
}

func (soapMonitor *SOAPMonitor) SetNotificationProfileID(notificationProfileID string) {
	soapMonitor.NotificationProfileID = notificationProfileID
}

func (soapMonitor *SOAPMonitor) GetNotificationProfileID() string {
	return soapMonitor.NotificationProfileID
}

func (soapMonitor *SOAPMonitor) SetUserGroupIDs(userGroupIDs []string) {
	soapMonitor.UserGroupIDs = userGroupIDs
}

func (soapMonitor *SOAPMonitor) GetUserGroupIDs() []string {
	return soapMonitor.UserGroupIDs
}

func (soapMonitor *SOAPMonitor) String() string {
	return ToString(soapMonitor)
}

func (soapMonitor *SOAPMonitor) SetTagIDs(tagIDs []string) {
	soapMonitor.TagIDs = tagIDs
}

func (soapMonitor *SOAPMonitor) GetTagIDs() []string {
	return soapMonitor.TagIDs
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

// Cron Monitor resource in Site24x7
type CronMonitor struct {
	_                     struct{} `type:"structure"` // Enforces key based initialization.
	MonitorID             string   `json:"monitor_id,omitempty"`
	DisplayName           string   `json:"display_name"`
	CronExpression        string   `json:"cron_expression"`
	CronTz                string   `json:"cron_tz"`
	WaitTime              int      `json:"wait_time"`
	Type                  string   `json:"type"`
	ThresholdProfileID    string   `json:"threshold_profile_id"`
	NotificationProfileID string   `json:"notification_profile_id"`
	MonitorGroups         []string `json:"monitor_groups,omitempty"`
	TagIDs                []string `json:"tag_ids,omitempty"`
	ThirdPartyServiceIDs  []string `json:"third_party_services,omitempty"`
	UserGroupIDs          []string `json:"user_group_ids,omitempty"`
	OnCallScheduleID      string   `json:"on_call_schedule_id,omitempty"`
}

func (cronMonitor *CronMonitor) SetLocationProfileID(locationProfileID string) {
}

func (cronMonitor *CronMonitor) GetLocationProfileID() string {
	return ""
}

func (cronMonitor *CronMonitor) SetNotificationProfileID(notificationProfileID string) {
	cronMonitor.NotificationProfileID = notificationProfileID
}

func (cronMonitor *CronMonitor) GetNotificationProfileID() string {
	return cronMonitor.NotificationProfileID
}

func (cronMonitor *CronMonitor) SetUserGroupIDs(userGroupIDs []string) {
	cronMonitor.UserGroupIDs = userGroupIDs
}

func (cronMonitor *CronMonitor) GetUserGroupIDs() []string {
	return cronMonitor.UserGroupIDs
}

func (cronMonitor *CronMonitor) SetTagIDs(tagIDs []string) {
	cronMonitor.TagIDs = tagIDs
}

func (cronMonitor *CronMonitor) GetTagIDs() []string {
	return cronMonitor.TagIDs
}

func (cronMonitor *CronMonitor) String() string {
	return ToString(cronMonitor)
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

type GCPMonitor struct {
	_                     struct{} `type:"structure"` // Enforces key-based initialization.
	MonitorID             string   `json:"monitor_id,omitempty"`
	DisplayName           string   `json:"display_name"`
	Type                  string   `json:"type"`
	ProjectID             string   `json:"project_id"`
	DiscoverServices      []int    `json:"gcp_discover_services,omitempty"`
	CheckFrequency        string   `json:"check_frequency"`
	GcpRegistrationMethod string   `json:"gcp_registration_method,omitempty"`
	StopRediscoverOption  int      `json:"stop_rediscover_option"`
	GCPSAContent          struct {
		PrivateKey  string `json:"private_key"`
		ClientEmail string `json:"client_email"`
	} `json:"gcp_sa_content"`
	UserGroupIDs          []string `json:"user_group_ids"`
	TagIDs                []string `json:"tag_ids,omitempty"`
	NotificationProfileID string   `json:"notification_profile_id"`
	GCPTagsType           int      `json:"gcp_tags_type,omitempty"`
	GCPTags               []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"gcp_tags,omitempty"`
}

func (gcpMonitor *GCPMonitor) SetTagIDs(tagIDs []string) {
	gcpMonitor.TagIDs = tagIDs
}

func (gcpMonitor *GCPMonitor) GetTagIDs() []string {
	return gcpMonitor.TagIDs
}

func (gcpMonitor *GCPMonitor) SetLocationProfileID(locationProfileID string) {
}

func (gcpMonitor *GCPMonitor) GetLocationProfileID() string {
	return ""
}
func (gcpMonitor *GCPMonitor) SetNotificationProfileID(notificationProfileID string) {
	gcpMonitor.NotificationProfileID = notificationProfileID
}

func (gcpMonitor *GCPMonitor) GetNotificationProfileID() string {
	return gcpMonitor.NotificationProfileID
}

func (gcpMonitor *GCPMonitor) SetUserGroupIDs(userGroupIDs []string) {
	gcpMonitor.UserGroupIDs = userGroupIDs
}

func (gcpMonitor *GCPMonitor) GetUserGroupIDs() []string {
	return gcpMonitor.UserGroupIDs
}

func (gcpMonitor *GCPMonitor) SetGCPTags(gcpTags []struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}) {
	gcpMonitor.GCPTags = gcpTags
}

func (gcpMonitor *GCPMonitor) GetGCPTags() []struct {
	Name  string `json:"name"`
	Value string `json:"value"`
} {
	return gcpMonitor.GCPTags
}

func (gcpMonitor *GCPMonitor) String() string {
	return ToString(gcpMonitor)
}
