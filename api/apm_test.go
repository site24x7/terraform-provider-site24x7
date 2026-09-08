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

// last.modified.time is documented as a long but sent quoted, the mirror image
// of agent_version, so both forms have to round-trip into the same string.
func TestAPMLastModifiedTimeUnmarshalJSON(t *testing.T) {
	for _, test := range []struct {
		name     string
		raw      string
		expected APMLastModifiedTime
	}{
		{name: "quoted string", raw: `"1563790347139"`, expected: "1563790347139"},
		{name: "bare number", raw: `1563790347139`, expected: "1563790347139"},
		{name: "null", raw: `null`, expected: ""},
	} {
		t.Run(test.name, func(t *testing.T) {
			var lastModified APMLastModifiedTime
			require.NoError(t, json.Unmarshal([]byte(test.raw), &lastModified))
			assert.Equal(t, test.expected, lastModified)
		})
	}
}

// Every mandatory agent_config key must be serialized even when its value is
// the zero value, and last.modified.time must never be sent: it is
// server-maintained and absent from the documented request bodies.
func TestAPMAgentConfigProfileSerializesMandatoryZeroValues(t *testing.T) {
	profile := &APMAgentConfigProfile{
		ProfileName: "Configuration Profile - Java",
		AgentType:   "JAVA",
		AgentConfig: APMAgentConfig{
			ApdexThreshold:                0.5,
			CloudInstanceCleanupThreshold: 7,
		},
	}

	serialized, err := json.Marshal(profile)
	require.NoError(t, err)

	var body map[string]interface{}
	require.NoError(t, json.Unmarshal(serialized, &body))

	// profile_id is omitted on create, when there is no ID yet.
	assert.NotContains(t, body, "profile_id")
	assert.Equal(t, false, body["is_default"])

	agentConfig := body["agent_config"].(map[string]interface{})
	assert.NotContains(t, agentConfig, "last.modified.time")

	for _, key := range []string{
		"transaction.trace.enabled",
		"transaction.trace.threshold",
		"transaction.trace.sql.parametrize",
		"transaction.trace.sql.stacktrace.threshold",
		"transaction.tracking.request.interval",
		"sql.capture.enabled",
		"autoupgrade.enabled",
		"show.instance.port.number",
		"apdex.threshold",
		"cloud.instance.cleanup.threshold",
	} {
		assert.Contains(t, agentConfig, key)
	}
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
