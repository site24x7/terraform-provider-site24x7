package apm

import (
	"encoding/json"
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const agentConfigProfileResponse = `{
  "agent_type": "JAVA",
  "agent_config": {
    "transaction.trace.enabled": true,
    "transaction.trace.sql.parametrize": true,
    "autoupgrade.enabled": false,
    "transaction.trace.sql.stacktrace.threshold": 3,
    "transaction.tracking.request.interval": 1,
    "last.modified.time": "1563790347139",
    "transaction.trace.threshold": 2,
    "sql.capture.enabled": true,
    "show.instance.port.number": true,
    "apdex.threshold": 0.5,
    "cloud.instance.cleanup.threshold": 7
  },
  "profile_name": "Configuration Profile - Java",
  "profile_id": "1000000009001",
  "is_default": false
}`

func expectedAgentConfigProfile() *api.APMAgentConfigProfile {
	return &api.APMAgentConfigProfile{
		ProfileID:   "1000000009001",
		ProfileName: "Configuration Profile - Java",
		AgentType:   "JAVA",
		IsDefault:   false,
		AgentConfig: api.APMAgentConfig{
			TransactionTraceEnabled:       true,
			TransactionTraceThreshold:     2,
			ParametrizeSQLQuery:           true,
			SQLStackTraceThreshold:        3,
			RequestTrackingInterval:       1,
			SQLCaptureEnabled:             true,
			AutoUpgradeEnabled:            false,
			ShowInstancePortNumber:        true,
			ApdexThreshold:                0.5,
			CloudInstanceCleanupThreshold: 7,
			LastModifiedTime:              "1563790347139",
		},
	}
}

// agentConfigProfileToWrite is what a caller hands to Create or Update: the
// same profile without the server-maintained last.modified.time.
func agentConfigProfileToWrite() *api.APMAgentConfigProfile {
	profile := expectedAgentConfigProfile()
	profile.AgentConfig.LastModifiedTime = ""
	return profile
}

// expectedWriteBody pins the wire format rather than round-tripping the struct:
// the keys are dotted, every mandatory setting is present even when it is false
// or zero, and the server-maintained last.modified.time is not sent.
const expectedWriteBody = `{
  "profile_name": "Configuration Profile - Java",
  "agent_type": "JAVA",
  "is_default": false,
  "agent_config": {
    "transaction.trace.enabled": true,
    "transaction.trace.threshold": 2,
    "transaction.trace.sql.parametrize": true,
    "transaction.trace.sql.stacktrace.threshold": 3,
    "transaction.tracking.request.interval": 1,
    "sql.capture.enabled": true,
    "autoupgrade.enabled": false,
    "show.instance.port.number": true,
    "apdex.threshold": 0.5,
    "cloud.instance.cleanup.threshold": 7
  }
}`

func TestAPMAgentConfigProfiles(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "Create APM agent configuration profile",
			ExpectedVerb: "POST",
			ExpectedPath: "/apminsight/agent_config_profile",
			ExpectedBody: []byte(expectedWriteBody),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, json.RawMessage(agentConfigProfileResponse)),
			Fn: func(t *testing.T, c rest.Client) {
				profile := agentConfigProfileToWrite()
				profile.ProfileID = ""

				created, err := NewAPMAgentConfigProfiles(c).Create(profile)
				require.NoError(t, err)

				assert.Equal(t, expectedAgentConfigProfile(), created)
			},
		},
		{
			Name:         "Update APM agent configuration profile",
			ExpectedVerb: "PUT",
			ExpectedPath: "/apminsight/agent_config_profile/1000000009001",
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, json.RawMessage(agentConfigProfileResponse)),
			Fn: func(t *testing.T, c rest.Client) {
				updated, err := NewAPMAgentConfigProfiles(c).Update(agentConfigProfileToWrite())
				require.NoError(t, err)

				assert.Equal(t, expectedAgentConfigProfile(), updated)
			},
		},
		{
			// There is no get-by-profile-id endpoint, so Get lists and filters.
			Name:         "Get APM agent configuration profile",
			ExpectedVerb: "GET",
			ExpectedPath: "/apminsight/agent_config_profiles",
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, json.RawMessage("["+agentConfigProfileResponse+"]")),
			Fn: func(t *testing.T, c rest.Client) {
				profile, err := NewAPMAgentConfigProfiles(c).Get("1000000009001")
				require.NoError(t, err)

				assert.Equal(t, expectedAgentConfigProfile(), profile)
			},
		},
		{
			// A profile deleted outside Terraform has to be reported as a 404 so
			// the resource can drop it from state instead of failing the refresh.
			Name:         "Get APM agent configuration profile that no longer exists",
			ExpectedVerb: "GET",
			ExpectedPath: "/apminsight/agent_config_profiles",
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, json.RawMessage("["+agentConfigProfileResponse+"]")),
			Fn: func(t *testing.T, c rest.Client) {
				profile, err := NewAPMAgentConfigProfiles(c).Get("1000000009002")

				assert.Nil(t, profile)
				require.Error(t, err)
				assert.True(t, apierrors.IsNotFound(err), "expected a not found error, got %v", err)
			},
		},
		{
			Name:         "Get default APM agent configuration profile",
			ExpectedVerb: "GET",
			ExpectedPath: "/apminsight/default_agent_config_profile/JAVA",
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, json.RawMessage(agentConfigProfileResponse)),
			Fn: func(t *testing.T, c rest.Client) {
				profile, err := NewAPMAgentConfigProfiles(c).GetDefault("JAVA")
				require.NoError(t, err)

				assert.Equal(t, expectedAgentConfigProfile(), profile)
			},
		},
		{
			Name:         "List APM agent configuration profiles",
			ExpectedVerb: "GET",
			ExpectedPath: "/apminsight/agent_config_profiles",
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, json.RawMessage("["+agentConfigProfileResponse+"]")),
			Fn: func(t *testing.T, c rest.Client) {
				profiles, err := NewAPMAgentConfigProfiles(c).List()
				require.NoError(t, err)

				assert.Equal(t, []*api.APMAgentConfigProfile{expectedAgentConfigProfile()}, profiles)
			},
		},
		{
			Name:         "List APM agent configuration profiles by agent type",
			ExpectedVerb: "GET",
			ExpectedPath: "/apminsight/agent_config_profiles/JAVA",
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, json.RawMessage("["+agentConfigProfileResponse+"]")),
			Fn: func(t *testing.T, c rest.Client) {
				profiles, err := NewAPMAgentConfigProfiles(c).ListByAgentType("JAVA")
				require.NoError(t, err)

				assert.Equal(t, []*api.APMAgentConfigProfile{expectedAgentConfigProfile()}, profiles)
			},
		},
		{
			// The delete response is a {name, id} pair rather than a profile, and
			// is discarded.
			Name:         "Delete APM agent configuration profile",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/apminsight/agent_config_profile/1000000009001",
			StatusCode:   200,
			ResponseBody: validation.JsonBody(t, &api.Response{Code: 0, Message: "success"}),
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewAPMAgentConfigProfiles(c).Delete("1000000009001"))
			},
		},
	})
}
