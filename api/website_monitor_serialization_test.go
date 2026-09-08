package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// WebsiteMonitor's bool fields must not carry `omitempty`.
//
// With `omitempty` a false bool is dropped from the payload entirely, so the
// API never sees it and keeps whatever it already had. That makes the value
// impossible to turn off from Terraform: the config says false, the API keeps
// reporting true, and every plan proposes the same change forever.
//
// Every other monitor struct in this package already serialises these fields
// unconditionally; WebsiteMonitor was the sole exception.
func TestWebsiteMonitorSerialisesFalseBools(t *testing.T) {
	monitor := &WebsiteMonitor{
		DisplayName:   "checkout",
		Type:          string(URL),
		Website:       "https://example.com",
		UseNameServer: false,
		UseIPV6:       false,
	}

	raw, err := json.Marshal(monitor)
	require.NoError(t, err)

	payload := map[string]interface{}{}
	require.NoError(t, json.Unmarshal(raw, &payload))

	for _, field := range []string{"use_name_server", "use_ipv6"} {
		value, present := payload[field]
		assert.Truef(t, present, "%s must be present in the payload even when false, otherwise it can never be turned off", field)
		assert.Equalf(t, false, value, "%s should serialise as false", field)
	}
}

// The true case was never broken, but pin it so a future change cannot regress
// both directions at once.
func TestWebsiteMonitorSerialisesTrueBools(t *testing.T) {
	monitor := &WebsiteMonitor{
		UseNameServer: true,
		UseIPV6:       true,
	}

	raw, err := json.Marshal(monitor)
	require.NoError(t, err)

	payload := map[string]interface{}{}
	require.NoError(t, json.Unmarshal(raw, &payload))

	assert.Equal(t, true, payload["use_name_server"])
	assert.Equal(t, true, payload["use_ipv6"])
}
