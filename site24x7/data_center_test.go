package site24x7

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The exact endpoints per data center, from https://www.site24x7.com/help/api/
//
// Japan is the odd one out: its API is served from app.site24x7.jp, not from a
// www host like every other region. www.site24x7.jp is the Japanese marketing
// site, so pointing the provider at it returns HTML, never API responses.
func TestDataCenterEndpoints(t *testing.T) {
	for _, test := range []struct {
		code     string
		baseURL  string
		tokenURL string
	}{
		{"US", "https://www.site24x7.com/api", "https://accounts.zoho.com/oauth/v2/token"},
		{"EU", "https://www.site24x7.eu/api", "https://accounts.zoho.eu/oauth/v2/token"},
		{"IN", "https://www.site24x7.in/api", "https://accounts.zoho.in/oauth/v2/token"},
		{"AU", "https://www.site24x7.net.au/api", "https://accounts.zoho.com.au/oauth/v2/token"},
		{"CN", "https://www.site24x7.cn/api", "https://accounts.zoho.com.cn/oauth/v2/token"},
		{"JP", "https://app.site24x7.jp/api", "https://accounts.zoho.jp/oauth/v2/token"},
		{"CA", "https://www.site24x7.ca/api", "https://accounts.zohocloud.ca/oauth/v2/token"},
	} {
		t.Run(test.code, func(t *testing.T) {
			dc := GetDataCenter(test.code)

			assert.Equal(t, test.baseURL, dc.GetAPIBaseURL())
			assert.Equal(t, test.tokenURL, dc.GetTokenURL())
		})
	}
}

// rest.Request builds every request URL by concatenating the base URL with
// "/" + resource, so a base URL carrying a trailing or doubled slash silently
// produces a path the API does not route - "//api/monitors" rather than
// "/api/monitors". These invariants catch that class of typo for every region,
// including ones added later.
func TestDataCenterURLsAreWellFormed(t *testing.T) {
	for code, dc := range dataCenter {
		t.Run(code, func(t *testing.T) {
			assert.Equal(t, code, dc.code, "map key and code field must agree")
			assert.NotEmpty(t, dc.displayName)

			baseURL, err := url.Parse(dc.GetAPIBaseURL())
			require.NoError(t, err)
			assert.Equal(t, "https", baseURL.Scheme)
			assert.NotEmpty(t, baseURL.Host)
			assert.Equal(t, "/api", baseURL.Path, "base URL path must be exactly /api")

			tokenURL, err := url.Parse(dc.GetTokenURL())
			require.NoError(t, err)
			assert.Equal(t, "https", tokenURL.Scheme)
			assert.NotEmpty(t, tokenURL.Host)
			assert.Equal(t, "/oauth/v2/token", tokenURL.Path)
		})
	}
}
