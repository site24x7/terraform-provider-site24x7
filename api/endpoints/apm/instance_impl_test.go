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

const instanceResponse = `{
  "response_time_data": {
    "response_time": 30.7,
    "request_count": 106915
  },
  "apdex_data": {
    "apdex": 0.996
  },
  "cpu_time": 5.5,
  "instance_info": {
    "port": 8080,
    "host": "192.168.1.1",
    "instance_name": "192.168.1.1:8080",
    "application_name": "plus.site24x7",
    "application_id": "101071000000059001",
    "instance_id": "101071000000059011",
    "ins_type": "JAVA",
    "agent_version": 1.8
  },
  "availability_health_info": {
    "availability": "AVAILABLE",
    "managed_state": true,
    "under_maintenance": false,
    "last_communication_time": 1440480146375
  },
  "app_config": {
    "is_cloud_app": true,
    "autoScale": false
  }
}`

func expectedInstance() *api.APMInstance {
	return &api.APMInstance{
		InstanceInfo: api.APMInstanceInfo{
			InstanceID:      "101071000000059011",
			InstanceName:    "192.168.1.1:8080",
			ApplicationID:   "101071000000059001",
			ApplicationName: "plus.site24x7",
			Host:            "192.168.1.1",
			Port:            8080,
			InstanceType:    "JAVA",
			AgentVersion:    "1.8",
		},
		AvailabilityHealthInfo: api.APMAvailabilityHealthInfo{
			Availability:     "AVAILABLE",
			ManagedState:     true,
			UnderMaintenance: false,
		},
		AppConfig: api.APMAppConfig{
			IsCloudApp: true,
			AutoScale:  false,
		},
	}
}

func TestAPMInstances(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "List APM instances",
			ExpectedVerb: "GET",
			ExpectedPath: "/apminsight/ins/H",
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, json.RawMessage("["+instanceResponse+"]")),
			Fn: func(t *testing.T, c rest.Client) {
				instances, err := NewAPMInstances(c).List("H")
				require.NoError(t, err)

				assert.Equal(t, []*api.APMInstance{expectedInstance()}, instances)
			},
		},
		{
			Name:         "Get APM instance",
			ExpectedVerb: "GET",
			ExpectedPath: "/apminsight/ins/101071000000059011/D",
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, json.RawMessage(instanceResponse)),
			Fn: func(t *testing.T, c rest.Client) {
				instance, err := NewAPMInstances(c).Get("101071000000059011", "D")
				require.NoError(t, err)

				assert.Equal(t, expectedInstance(), instance)
			},
		},
		{
			Name:         "Manage APM instance",
			ExpectedVerb: "POST",
			ExpectedPath: "/apminsight/ins/15698000009961011/manage",
			StatusCode:   200,
			ResponseBody: validation.JsonBody(t, &api.Response{Code: 0, Message: "success"}),
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewAPMInstances(c).Manage("15698000009961011"))
			},
		},
		{
			Name:         "Unmanage APM instance",
			ExpectedVerb: "POST",
			ExpectedPath: "/apminsight/ins/15698000009961011/unmanage",
			StatusCode:   200,
			ResponseBody: validation.JsonBody(t, &api.Response{Code: 0, Message: "success"}),
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewAPMInstances(c).Unmanage("15698000009961011"))
			},
		},
		{
			Name:         "Delete APM instance",
			ExpectedVerb: "DELETE",
			ExpectedPath: "/apminsight/ins/15698000009961011",
			StatusCode:   200,
			ResponseBody: validation.JsonBody(t, &api.Response{Code: 0, Message: "success"}),
			Fn: func(t *testing.T, c rest.Client) {
				require.NoError(t, NewAPMInstances(c).Delete("15698000009961011"))
			},
		},
	})
}
