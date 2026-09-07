package api

import (
	"encoding/json"
	"strconv"
)

// parseNumericString decodes a value the APM Insight API sends inconsistently
// as either a quoted string or a bare number, always yielding a string. A JSON
// null yields the empty string.
//
// Several APM fields are documented as one wire type and sent as the other -
// agent_version is declared a string and sent as a number, last.modified.time
// is declared a long and sent as a string - so both forms have to be accepted.
func parseNumericString(rawValue []byte) (string, error) {
	if len(rawValue) == 0 || string(rawValue) == "null" {
		return "", nil
	}

	if rawValue[0] == '"' {
		var valueAsString string
		if err := json.Unmarshal(rawValue, &valueAsString); err != nil {
			return "", err
		}

		return valueAsString, nil
	}

	var valueAsFloat float64
	if err := json.Unmarshal(rawValue, &valueAsFloat); err != nil {
		return "", err
	}

	return strconv.FormatFloat(valueAsFloat, 'f', -1, 64), nil
}

// AgentVersion is the APM Insight agent version reported by an instance.
//
// The API is inconsistent about the wire type: the documentation declares it a
// string, but responses carry it as a bare number (e.g. 1.7, 4.5). Both forms
// unmarshal into a string here, following the same approach as Status.
type AgentVersion string

func (version *AgentVersion) UnmarshalJSON(rawValue []byte) error {
	parsed, err := parseNumericString(rawValue)
	if err != nil {
		return err
	}

	*version = AgentVersion(parsed)
	return nil
}

// APMLastModifiedTime is when an agent configuration profile was last updated,
// in milliseconds since the epoch.
//
// The API documents it as a long but sends it quoted, so both forms are
// accepted, the same way AgentVersion handles the mirror-image inconsistency.
type APMLastModifiedTime string

func (lastModified *APMLastModifiedTime) UnmarshalJSON(rawValue []byte) error {
	parsed, err := parseNumericString(rawValue)
	if err != nil {
		return err
	}

	*lastModified = APMLastModifiedTime(parsed)
	return nil
}

// APMAgentConfig is the agent_config object of an agent configuration profile.
//
// Every key is dotted on the wire, and every one of them is mandatory on
// create and update except cloud.instance.cleanup.threshold, so none of the
// fields carry omitempty: dropping a false or a zero would leave the request
// short of a required key.
//
// LastModifiedTime is the exception. It is server-maintained, absent from the
// documented request examples, and advances on every write, so it is only ever
// read.
type APMAgentConfig struct {
	_                             struct{}            `type:"structure"`
	TransactionTraceEnabled       bool                `json:"transaction.trace.enabled"`
	TransactionTraceThreshold     int                 `json:"transaction.trace.threshold"`
	ParametrizeSQLQuery           bool                `json:"transaction.trace.sql.parametrize"`
	SQLStackTraceThreshold        int                 `json:"transaction.trace.sql.stacktrace.threshold"`
	RequestTrackingInterval       int                 `json:"transaction.tracking.request.interval"`
	SQLCaptureEnabled             bool                `json:"sql.capture.enabled"`
	AutoUpgradeEnabled            bool                `json:"autoupgrade.enabled"`
	ShowInstancePortNumber        bool                `json:"show.instance.port.number"`
	ApdexThreshold                float64             `json:"apdex.threshold"`
	CloudInstanceCleanupThreshold int                 `json:"cloud.instance.cleanup.threshold"`
	LastModifiedTime              APMLastModifiedTime `json:"last.modified.time,omitempty"`
}

// APMAgentConfigProfile is a named set of agent settings applied to every
// instance of one APM Insight agent type.
//
// ProfileID is documented as an int but sent as a quoted string, like every
// other Site24x7 resource ID, so it is modelled as a string.
type APMAgentConfigProfile struct {
	_           struct{}       `type:"structure"`
	ProfileID   string         `json:"profile_id,omitempty"`
	ProfileName string         `json:"profile_name"`
	AgentType   string         `json:"agent_type"`
	IsDefault   bool           `json:"is_default"`
	AgentConfig APMAgentConfig `json:"agent_config"`
}

// APMRUMInfo links an APM Insight application to its Real User Monitoring
// application. It is returned as an empty object when no RUM app is linked.
type APMRUMInfo struct {
	_          struct{} `type:"structure"`
	RUMAppID   string   `json:"rumAppId,omitempty"`
	RUMAppKey  string   `json:"rumAppKey,omitempty"`
	RUMAppName string   `json:"rumAppName,omitempty"`
}

// APMInstanceInfo identifies a single reporting agent instance. In Java each
// JVM is an instance, in .NET each IIS app, in Ruby each Rails server, and in
// PHP each PHP server.
type APMInstanceInfo struct {
	_               struct{}     `type:"structure"`
	InstanceID      string       `json:"instance_id"`
	InstanceName    string       `json:"instance_name"`
	ApplicationID   string       `json:"application_id"`
	ApplicationName string       `json:"application_name"`
	Host            string       `json:"host"`
	Port            int          `json:"port"`
	InstanceType    string       `json:"ins_type"`
	AgentVersion    AgentVersion `json:"agent_version"`
}

// APMApplicationInfo is the identity and topology of an APM Insight
// application.
type APMApplicationInfo struct {
	_               struct{}                   `type:"structure"`
	ApplicationID   string                     `json:"application_id"`
	ApplicationName string                     `json:"application_name"`
	InstanceIDs     []string                   `json:"instance_ids"`
	InstanceCount   int                        `json:"instance_count"`
	HostCount       int                        `json:"host_count"`
	Hosts           map[string][]string        `json:"hosts"`
	Instances       map[string]APMInstanceInfo `json:"instances"`
	RUMInfo         APMRUMInfo                 `json:"rum_info"`
}

// APMAvailabilityHealthInfo carries the lifecycle state of an application or
// instance.
//
// Only the discrete status fields are modelled. The API also returns
// last_communication_time, a millisecond timestamp that advances on every
// agent heartbeat; it is deliberately omitted so that refreshing never
// produces a diff.
type APMAvailabilityHealthInfo struct {
	_                struct{} `type:"structure"`
	Availability     string   `json:"availability"`
	ManagedState     bool     `json:"managed_state"`
	UnderMaintenance bool     `json:"under_maintenance"`
}

// APMAppConfig describes how an instance is hosted.
type APMAppConfig struct {
	_          struct{} `type:"structure"`
	IsCloudApp bool     `json:"is_cloud_app"`
	AutoScale  bool     `json:"autoScale"`
}

// APMApplication is the identity and configuration view of an APM Insight
// application.
//
// The reporting API also returns aggregate performance metrics for the
// requested time window - response_time_data, apdex_data, exception_info and
// cpu_time. Those are deliberately not modelled: they are recomputed on every
// request, so carrying them in Terraform state would report drift that no
// apply could ever resolve.
type APMApplication struct {
	_                      struct{}                  `type:"structure"`
	ApplicationInfo        APMApplicationInfo        `json:"application_info"`
	AvailabilityHealthInfo APMAvailabilityHealthInfo `json:"availability_health_info"`
}

// APMInstance is the identity and configuration view of a single APM Insight
// agent instance. As with APMApplication, performance metrics are not
// modelled.
type APMInstance struct {
	_                      struct{}                  `type:"structure"`
	InstanceInfo           APMInstanceInfo           `json:"instance_info"`
	ApplicationInfo        APMApplicationInfo        `json:"application_info"`
	AvailabilityHealthInfo APMAvailabilityHealthInfo `json:"availability_health_info"`
	AppConfig              APMAppConfig              `json:"app_config"`
}
