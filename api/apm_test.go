package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The APM Insight API documents agent_version as a string but sends it as a
// bare number, so both forms have to round-trip into the same string.
func TestAgentVersionUnmarshalJSON(t *testing.T) {
	for _, test := range []struct {
		name     string
		raw      string
		expected AgentVersion
	}{
		{name: "bare number", raw: `1.7`, expected: "1.7"},
		{name: "number without a fraction", raw: `2`, expected: "2"},
		{name: "number with trailing precision", raw: `4.55`, expected: "4.55"},
		{name: "quoted string", raw: `"1.8"`, expected: "1.8"},
		{name: "non-numeric string", raw: `"1.8.2-beta"`, expected: "1.8.2-beta"},
		{name: "null", raw: `null`, expected: ""},
	} {
		t.Run(test.name, func(t *testing.T) {
			var version AgentVersion
			require.NoError(t, json.Unmarshal([]byte(test.raw), &version))
			assert.Equal(t, test.expected, version)
		})
	}
}

func TestAgentVersionUnmarshalJSONRejectsUnexpectedTypes(t *testing.T) {
	var version AgentVersion
	assert.Error(t, json.Unmarshal([]byte(`{"version": 1.7}`), &version))
}

// An application with no linked RUM app comes back as an empty object rather
// than a null, and must not fail the whole parse.
func TestAPMApplicationParsesEmptyRUMInfo(t *testing.T) {
	raw := `{
	  "application_info": {
	    "application_id": "123",
	    "application_name": "checkout",
	    "rum_info": {},
	    "instance_ids": [],
	    "instance_count": 0
	  },
	  "availability_health_info": {
	    "availability": "AVAILABLE",
	    "managed_state": true,
	    "under_maintenance": false
	  }
	}`

	application := &APMApplication{}
	require.NoError(t, json.Unmarshal([]byte(raw), application))

	assert.Equal(t, "123", application.ApplicationInfo.ApplicationID)
	assert.Equal(t, APMRUMInfo{}, application.ApplicationInfo.RUMInfo)
	assert.True(t, application.AvailabilityHealthInfo.ManagedState)
}
