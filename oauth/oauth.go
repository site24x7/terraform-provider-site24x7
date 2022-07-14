package oauth

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/oauth2"
)

const (
	// TokenURL is the URL from where access tokens for the Site24x7 API are
	// obtained.
	TokenURL = "https://accounts.zoho.com/oauth/v2/token"

	// TokenType is the type used in the Authorization header next to the
	// access token.
	TokenType = "Zoho-oauthtoken"
)

// Config is an OAuth config that is also aware of the refresh token.
type Config struct {
	*oauth2.Config

	// AccessToken is a token that's used by the application for identifying the user
	// and retrieve the data related to the user.
	AccessToken string

	// AccessToken expiry in seconds
	Expiry string

	// RefreshToken is a token that's used by the application
	// (as opposed to the user) to refresh the access token
	// if it expires.
	RefreshToken string
}

// NewConfig creates a new *Config for the provided client credentials.
func NewConfig(clientID, clientSecret, refreshToken, accessToken, expiry string) *Config {
	return &Config{
		Config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Endpoint: oauth2.Endpoint{
				AuthStyle: oauth2.AuthStyleInParams,
				TokenURL:  TokenURL,
			},
		},
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
		Expiry:       expiry,
	}
}

// Client returns a *http.Client which automatically retrieves OAuth access
// tokens and attaches them to any request made with it.
func (c *Config) Client(ctx context.Context) *http.Client {
	return &http.Client{
		Transport: &oauth2.Transport{
			Source: c.TokenSource(ctx),
		},
	}
}

// TokenSource creates an oauth2.TokenSource which obtains access tokens using
// the refresh token.
func (c *Config) TokenSource(ctx context.Context) oauth2.TokenSource {
	t := &oauth2.Token{
		RefreshToken: c.RefreshToken,
		TokenType:    TokenType,
	}
	if c.AccessToken != "" {
		t.AccessToken = c.AccessToken
		if c.Expiry != "" {
			if expiry, err := strconv.ParseInt(c.Expiry, 10, 64); err == nil {
				t.Expiry = time.Now().Local().Add(time.Duration(expiry) * time.Second)
			}
		}
	}
	tokenSrc := &tokenSource{
		delegate: c.Config.TokenSource(ctx, t),
	}
	return tokenSrc
}
