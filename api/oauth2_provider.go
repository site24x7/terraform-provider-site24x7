package api

// SendTokenAs specifies how the access token should be sent (via query parameters or request headers).
type SendTokenAs struct {
	Method string `json:"method"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}

// RequestBodyParam specifies parameters to be sent in the request body.
type RequestBodyParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// QueryParam specifies additional query parameters for the authorization URL.
type QueryParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// OAuth2Scope specifies the level of access that the application is requesting.
type OAuth2Scope struct {
	Value string `json:"value"`
}

// OAuth2Provider represents an OAuth 2 Provider configuration in Site24x7.
type OAuth2Provider struct {
	_                struct{}           `type:"structure"`
	ProviderID       string             `json:"provider_id,omitempty"`
	ProviderName     string             `json:"provider_name"`
	OAuth2Flow       int                `json:"oauth2_flow"`
	ClientID         string             `json:"client_id"`
	ClientSecret     string             `json:"client_secret"`
	AccessTokenURI   string             `json:"access_token_uri"`
	AuthorizationURI string             `json:"authorization_uri,omitempty"`
	AuthMethod       string             `json:"auth_method,omitempty"`
	AuthUser         string             `json:"auth_user,omitempty"`
	AuthPass         string             `json:"auth_pass,omitempty"`
	AutoReauthorize  bool               `json:"auto_reauthorize,omitempty"`
	AccessToken      string             `json:"access_token,omitempty"`
	RefreshToken     string             `json:"refresh_token,omitempty"`
	ExpiryTime       string             `json:"expiry_time,omitempty"`
	SendTokenAs      *SendTokenAs       `json:"send_token_as"`
	RequestBody      []RequestBodyParam `json:"request_body,omitempty"`
	QueryParams      []QueryParam       `json:"query_params,omitempty"`
	OAuth2Scopes     []OAuth2Scope      `json:"oauth2_scopes,omitempty"`
	UserGroupIDs     []string           `json:"user_group_ids,omitempty"`
}
