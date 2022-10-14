package site24x7

import (
	"context"
	"net/http"

	"github.com/site24x7/terraform-provider-site24x7/api/endpoints"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/common"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/integration"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/monitors"
	"github.com/site24x7/terraform-provider-site24x7/backoff"
	"github.com/site24x7/terraform-provider-site24x7/oauth"
	"github.com/site24x7/terraform-provider-site24x7/rest"
)

const (
	// APIBaseURL is the base url of the Site24x7 API.
	APIBaseURL = "https://www.site24x7.com/api"
)

// Config is the configuration for the Site24x7 API Client.
type Config struct {
	// ClientID is the OAuth client ID needed to obtain an access token for API
	// usage.
	ClientID string

	// ClientSecret is the OAuth client secret needed to obtain an access token
	// for API usage.
	ClientSecret string

	// RefreshToken is a token that's used by the application
	// (as opposed to the user) to refresh the access token
	// if it expires.
	RefreshToken string

	// AccessToken is a token that's used by the application for identifying the user
	// and retrieve the data related to him.
	AccessToken string

	// AccessToken expiry in seconds
	Expiry string

	// Application Account ID of the customer.
	ZAAID string

	// APIBaseURL allows overriding the default API base URL (https://www.site24x7.com/api).
	// See https://www.site24x7.com/help/api/index.html#introduction for options of data centers for top level domain.
	APIBaseURL string

	// TokenURL allows overriding the default token URL (https://accounts.zoho.com/oauth/v2/token).
	// See https://www.site24x7.com/help/api/index.html#authentication for options of data centers for top level domain.
	TokenURL string

	// RetryConfig contains the configuration of the backoff-retry behavior. If
	// nil, backoff.DefaultRetryConfig will be used.
	RetryConfig *backoff.RetryConfig
}

// OAuthClient creates a new *http.Client from c that transparently obtains and
// attaches OAuth access tokens to every request.
func (c *Config) OAuthClient(ctx context.Context) *http.Client {
	oauthConfig := oauth.NewConfig(c.ClientID, c.ClientSecret, c.RefreshToken, c.AccessToken, c.Expiry)
	if c.TokenURL != "" {
		oauthConfig.Endpoint.TokenURL = c.TokenURL
	}

	return oauthConfig.Client(ctx)
}

// HTTPClient is the interface of an http client that is compatible with
// *http.Client.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is the Site24x7 API Client interface. It provides methods to get
// clients for resource endpoints.
type Client interface {
	CurrentStatus() endpoints.CurrentStatus
	LocationProfiles() endpoints.LocationProfiles
	LocationTemplate() endpoints.LocationTemplate
	MonitorGroups() endpoints.MonitorGroups
	Subgroups() endpoints.Subgroups
	Tags() endpoints.Tags
	ScheduleMaintenance() common.ScheduleMaintenance
	WebsiteMonitors() monitors.WebsiteMonitors
	WebPageSpeedMonitors() monitors.WebPageSpeedMonitors
	SSLMonitors() monitors.SSLMonitors
	HeartbeatMonitors() monitors.HeartbeatMonitors
	ServerMonitors() monitors.ServerMonitors
	RestApiMonitors() monitors.RestApiMonitors
	AmazonMonitors() monitors.AmazonMonitors
	NotificationProfiles() endpoints.NotificationProfiles
	ThresholdProfiles() endpoints.ThresholdProfiles
	UserGroups() endpoints.UserGroups
	URLAutomations() endpoints.URLAutomations
	ThirdPartyIntegrations() integration.ThirdpartyIntegrations
	OpsgenieIntegration() integration.OpsgenieIntegration
	SlackIntegration() integration.SlackIntegration
	WebhookIntegration() integration.WebhookIntegration
	PagerDutyIntegration() integration.PagerDutyIntegration
	ServiceNowIntegration() integration.ServiceNowIntegration
}

type client struct {
	restClient rest.Client
}

// New creates a new Site24x7 API Client with Config c.
func New(c Config) Client {
	httpClient := backoff.WithRetries(
		c.OAuthClient(context.Background()),
		c.RetryConfig,
	)

	return NewClient(httpClient, c)
}

// NewClient creates a new Site24x7 API Client from httpClient with default API base URL.
// This can be used to provide a custom http client for use with the API. The custom http
// client has to transparently handle the Site24x7 OAuth flow.
func NewClient(httpClient HTTPClient, c Config) Client {

	clientConfig := rest.ClientConfig{
		APIBaseURL: c.APIBaseURL,
		TokenURL:   c.TokenURL,
		ZAAID:      c.ZAAID,
	}
	if c.ZAAID != "" {
		clientConfig.MSP = true
	}
	return &client{
		restClient: rest.NewClient(httpClient, clientConfig),
	}

}

// // NewClientWithBaseURL creates a new Site24x7 API Client from httpClient and given API base URL.
// // This can be used to provide a custom http client for use with the API. The custom http
// // client has to transparently handle the Site24x7 OAuth flow.
// func NewClientWithBaseURL(httpClient HTTPClient, baseURL string) Client {
// 	return &client{
// 		restClient: rest.NewClient(httpClient, baseURL),
// 	}
// }

// CurrentStatus implements Client.
func (c *client) CurrentStatus() endpoints.CurrentStatus {
	return endpoints.NewCurrentStatus(c.restClient)
}

// LocationProfiles implements Client.
func (c *client) LocationProfiles() endpoints.LocationProfiles {
	return endpoints.NewLocationProfiles(c.restClient)
}

// ScheduleMaintenance implements Client.
func (c *client) ScheduleMaintenance() common.ScheduleMaintenance {
	return common.NewScheduleMaintenance(c.restClient)
}

// LocationTemplate implements Client.
func (c *client) LocationTemplate() endpoints.LocationTemplate {
	return endpoints.NewLocationTemplate(c.restClient)
}

// AmazonMonitors implements Client.
func (c *client) AmazonMonitors() monitors.AmazonMonitors {
	return monitors.NewAmazonMonitors(c.restClient)
}

// WebsiteMonitors implements Client.
func (c *client) WebsiteMonitors() monitors.WebsiteMonitors {
	return monitors.NewMonitors(c.restClient)
}

// WebPageSpeedMonitors implements Client.
func (c *client) WebPageSpeedMonitors() monitors.WebPageSpeedMonitors {
	return monitors.NewWebPageSpeedMonitors(c.restClient)
}

// SSLMonitors implements Client.
func (c *client) SSLMonitors() monitors.SSLMonitors {
	return monitors.NewSSLMonitors(c.restClient)
}

// HeartbeatMonitors implements Client.
func (c *client) HeartbeatMonitors() monitors.HeartbeatMonitors {
	return monitors.NewHeartbeatMonitors(c.restClient)
}

// ServerMonitors implements Client.
func (c *client) ServerMonitors() monitors.ServerMonitors {
	return monitors.NewServerMonitors(c.restClient)
}

// RestApiMonitors implements Client.
func (c *client) RestApiMonitors() monitors.RestApiMonitors {
	return monitors.NewRestApiMonitors(c.restClient)
}

// MonitorGroups implements Client.
func (c *client) MonitorGroups() endpoints.MonitorGroups {
	return endpoints.NewMonitorGroups(c.restClient)
}

// Subgroups implements Client.
func (c *client) Subgroups() endpoints.Subgroups {
	return endpoints.NewSubgroups(c.restClient)
}

// Tags implements Client.
func (c *client) Tags() endpoints.Tags {
	return endpoints.NewTags(c.restClient)
}

// NotificationProfiles implements Client.
func (c *client) NotificationProfiles() endpoints.NotificationProfiles {
	return endpoints.NewNotificationProfiles(c.restClient)
}

// ThresholdProfiles implements Client.
func (c *client) ThresholdProfiles() endpoints.ThresholdProfiles {
	return endpoints.NewThresholdProfiles(c.restClient)
}

// UserGroups implements Client.
func (c *client) UserGroups() endpoints.UserGroups {
	return endpoints.NewUserGroups(c.restClient)
}

// ItAutomations implements Client.
func (c *client) URLAutomations() endpoints.URLAutomations {
	return endpoints.NewURLAutomations(c.restClient)
}

// OpsgenieIntegraion implements Client.
func (c *client) OpsgenieIntegration() integration.OpsgenieIntegration {
	return integration.NewOpsgenie(c.restClient)
}

// SlackIntegraion implements Client.
func (c *client) SlackIntegration() integration.SlackIntegration {
	return integration.NewSlack(c.restClient)
}

// WebhookIntegration implements Client.
func (c *client) WebhookIntegration() integration.WebhookIntegration {
	return integration.NewWebhook(c.restClient)
}

// PagerDutyIntegration implements Client.
func (c *client) PagerDutyIntegration() integration.PagerDutyIntegration {
	return integration.NewPagerDuty(c.restClient)
}

// ServiceNowIntegration implements Client.
func (c *client) ServiceNowIntegration() integration.ServiceNowIntegration {
	return integration.NewServiceNow(c.restClient)
}

// ThirdPartyIntegrations implements Client.
func (c *client) ThirdPartyIntegrations() integration.ThirdpartyIntegrations {
	return integration.NewThirdpartyIntegrations(c.restClient)
}
