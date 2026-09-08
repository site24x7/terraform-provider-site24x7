package apm

import (
	"encoding/json"
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// applicationResponse mirrors a real payload, metrics included, so that the
// tests prove the metric blocks are parsed past and dropped rather than merely
// being absent from the fixture.
const applicationResponse = `{
  "response_time_data": {
    "response_time": 5112.5,
    "request_count": 66533,
    "throughput": 1108.88
  },
  "exception_info": {
    "warning_count": 0,
    "fatal_count": 0
  },
  "apdex_data": {
    "apdex": 0.875,
    "satisfied": 83.6
  },
  "cpu_time": 98.8,
  "application_info": {
    "instance_ids": ["101071000000034035"],
    "host_count": 1,
    "application_name": "Site24x7",
    "hosts": {
      "192.168.1.1": ["101071000000034035"]
    },
    "application_id": "101071000000034001",
    "rum_info": {
      "rumAppId": "15698000010568013",
      "rumAppKey": "61cc526aa43df99a460d1e7bce9e635d",
      "rumAppName": "BookingApp"
    },
    "instances": {
      "101071000000034035": {
        "port": 8080,
        "host": "192.168.1.1",
        "instance_name": "192.168.1.1:8080",
        "application_name": "Site24x7",
        "application_id": "101071000000034001",
        "instance_id": "101071000000034035",
        "ins_type": "JAVA",
        "agent_version": 1.7
      }
    },
    "instance_count": 1
  },
  "availability_health_info": {
    "availability": "AVAILABLE",
    "managed_state": true,
    "under_maintenance": false,
    "last_communication_time": 1440480146375
  }
}`

func expectedApplication() *api.APMApplication {
	return &api.APMApplication{
		ApplicationInfo: api.APMApplicationInfo{
			ApplicationID:   "101071000000034001",
			ApplicationName: "Site24x7",
			InstanceIDs:     []string{"101071000000034035"},
			InstanceCount:   1,
			HostCount:       1,
			Hosts: map[string][]string{
				"192.168.1.1": {"101071000000034035"},
			},
			Instances: map[string]api.APMInstanceInfo{
				"101071000000034035": {
					InstanceID:      "101071000000034035",
					InstanceName:    "192.168.1.1:8080",
					ApplicationID:   "101071000000034001",
					ApplicationName: "Site24x7",
					Host:            "192.168.1.1",
					Port:            8080,
					InstanceType:    "JAVA",
					AgentVersion:    "1.7",
				},
			},
			RUMInfo: api.APMRUMInfo{
				RUMAppID:   "15698000010568013",
				RUMAppKey:  "61cc526aa43df99a460d1e7bce9e635d",
				RUMAppName: "BookingApp",
			},
		},
		AvailabilityHealthInfo: api.APMAvailabilityHealthInfo{
			Availability:     "AVAILABLE",
			ManagedState:     true,
			UnderMaintenance: false,
		},
	}
}

func TestAPMApplications(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "List APM applications",
			ExpectedVerb: "GET",
			ExpectedPath: "/apminsight/app/D",
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, json.RawMessage("["+applicationResponse+"]")),
			Fn: func(t *testing.T, c rest.Client) {
				applications, err := NewAPMApplications(c).List("D")
				require.NoError(t, err)

				assert.Equal(t, []*api.APMApplication{expectedApplication()}, applications)
			},
		},
		{
			Name:         "List APM applications falls back to the default time window",
			ExpectedVerb: "GET",
			ExpectedPath: "/apminsight/app/H",
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, json.RawMessage("[]")),
			Fn: func(t *testing.T, c rest.Client) {
				_, err := NewAPMApplications(c).List("")
				require.NoError(t, err)
			},
		},
		{
			Name:         "Get APM application",
			ExpectedVerb: "GET",
			ExpectedPath: "/apminsight/app/101071000000034001/H",
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, json.RawMessage(applicationResponse)),
			Fn: func(t *testing.T, c rest.Client) {
				application, err := NewAPMApplications(c).Get("101071000000034001", "H")
				require.NoError(t, err)

				assert.Equal(t, expectedApplication(), application)
			},
		},
		{
			Name:         "Manage APM application",
			ExpectedVerb: "POST",
			ExpectedPath: "/apminsight/app/15698000009961011/manage",
			StatusCode:   200,
			ResponseBody: validation.JsonBody(t, &api.Response{Code: 0, Message: "success"}),
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewAPMApplications(c).Manage("15698000009961011"))
			},
		},
		{
			Name:         "Unmanage APM application",
			ExpectedVerb: "POST",
			ExpectedPath: "/apminsight/app/15698000009961011/unmanage",
			StatusCode:   200,
			ResponseBody: validation.JsonBody(t, &api.Response{Code: 0, Message: "success"}),
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewAPMApplications(c).Unmanage("15698000009961011"))
			},
		},
		{
			Name:         "Delete APM application",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/apminsight/app/15698000009961011",
			StatusCode:   200,
			ResponseBody: validation.JsonBody(t, &api.Response{Code: 0, Message: "success"}),
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewAPMApplications(c).Delete("15698000009961011"))
			},
		},
	})
}
