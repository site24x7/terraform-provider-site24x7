package api

type Site24x7Monitor interface {
	SetNotificationProfileID(notificationProfileID string)
	GetNotificationProfileID() string
	SetUserGroupIDs(userGroupIDs []string)
	GetUserGroupIDs() []string
	SetTagIDs(tagIDs []string)
	GetTagIDs() []string
	String() string
}

// Denotes the website monitor resource in Site24x7.
type WebsiteMonitor struct {
	_                     struct{}           `type:"structure"` // Enforces key based initialization.
	MonitorID             string             `json:"monitor_id,omitempty"`
	DisplayName           string             `json:"display_name"`
	Type                  string             `json:"type"`
	Website               string             `json:"website"`
	CheckFrequency        string             `json:"check_frequency"`
	HTTPMethod            string             `json:"http_method"`
	AuthUser              string             `json:"auth_user"`
	AuthPass              string             `json:"auth_pass"`
	MatchingKeyword       *ValueAndSeverity  `json:"matching_keyword,omitempty"`
	UnmatchingKeyword     *ValueAndSeverity  `json:"unmatching_keyword,omitempty"`
	MatchRegex            *ValueAndSeverity  `json:"match_regex,omitempty"`
	MatchCase             bool               `json:"match_case"`
	UserAgent             string             `json:"user_agent"`
	Timeout               int                `json:"timeout"`
	UseNameServer         bool               `json:"use_name_server"`
	UpStatusCodes         string             `json:"up_status_codes"`
	CustomHeaders         []Header           `json:"custom_headers,omitempty"`
	ResponseHeaders       HTTPResponseHeader `json:"response_headers_check,omitempty"`
	LocationProfileID     string             `json:"location_profile_id"`
	NotificationProfileID string             `json:"notification_profile_id"`
	ThresholdProfileID    string             `json:"threshold_profile_id"`
	MonitorGroups         []string           `json:"monitor_groups,omitempty"`
	UserGroupIDs          []string           `json:"user_group_ids,omitempty"`
	TagIDs                []string           `json:"tag_ids,omitempty"`
	ThirdPartyServiceIDs  []string           `json:"third_party_services,omitempty"`
	ActionIDs             []ActionRef        `json:"action_ids,omitempty"`
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
	UserGroupIDs          []string    `json:"user_group_ids,omitempty"`
	TagIDs                []string    `json:"tag_ids,omitempty"`
	ThirdPartyServiceIDs  []string    `json:"third_party_services,omitempty"`
	// ActionIDs             []ActionRef `json:"action_ids,omitempty"`
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
	_                         struct{}               `type:"structure"` // Enforces key based initialization.
	MonitorID                 string                 `json:"monitor_id,omitempty"`
	DisplayName               string                 `json:"display_name"`
	Type                      string                 `json:"type"`
	Website                   string                 `json:"website"`
	CheckFrequency            string                 `json:"check_frequency"`
	Timeout                   int                    `json:"timeout"`
	HttpMethod                string                 `json:"http_method,omitempty"`
	HttpProtocol              string                 `json:"http_protocol,omitempty"`
	SslProtocol               string                 `json:"ssl_protocol,omitempty"`
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
	ResponseHeaders           HTTPResponseHeader     `json:"response_headers_check,omitempty"`
	LocationProfileID         string                 `json:"location_profile_id"`
	NotificationProfileID     string                 `json:"notification_profile_id"`
	ThresholdProfileID        string                 `json:"threshold_profile_id"`
	MonitorGroups             []string               `json:"monitor_groups,omitempty"`
	UserGroupIDs              []string               `json:"user_group_ids,omitempty"`
	TagIDs                    []string               `json:"tag_ids,omitempty"`
	ThirdPartyServiceIDs      []string               `json:"third_party_services,omitempty"`
	ActionIDs                 []ActionRef            `json:"action_ids,omitempty"`
	// HTTP Configuration
	UpStatusCodes string `json:"up_status_codes,omitempty"`
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

// Denotes the Amazon monitor resource in Site24x7.
type AmazonMonitor struct {
	_                     struct{} `type:"structure"` // Enforces key based initialization.
	MonitorID             string   `json:"monitor_id,omitempty"`
	DisplayName           string   `json:"display_name"`
	Type                  string   `json:"type"`
	SecretKey             string   `json:"aws_secret_key"`
	AccessKey             string   `json:"aws_access_key"`
	DiscoverFrequency     int      `json:"aws_discovery_frequency"`
	DiscoverServices      []string `json:"aws_discover_services"`
	NotificationProfileID string   `json:"notification_profile_id"`
	UserGroupIDs          []string `json:"user_group_ids"`
	TagIDs                []string `json:"tag_ids,omitempty"`
	ThirdPartyServiceIDs  []string `json:"third_party_services,omitempty"`
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
