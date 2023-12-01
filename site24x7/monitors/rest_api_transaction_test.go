package monitors

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRestApiTransactionMonitorCreate(t *testing.T) {
	d := restApiTransactionMonitorTestResourceData(t)

	c := fake.NewClient()

	a := &api.RestApiTransactionMonitor{
		DisplayName:           "foo",
		Type:                  string(api.RESTAPISEQ),
		CheckFrequency:        "5",
		LocationProfileID:     "456",
		NotificationProfileID: "789",
		ThresholdProfileID:    "012",
		MonitorGroups:         []string{"234", "567"},
		DependencyResourceIDs: []string{"234", "567"},
		UserGroupIDs:          []string{"123", "456"},
		TagIDs:                []string{"123"},
		Steps: []api.Steps{
			{
				DisplayName: "Step2",
				StepsDetails: []api.StepDetails{
					{
						StepUrl:                   "www.test.tld",
						Timeout:                   "0",
						HTTPMethod:                "G",
						DisplayName:               "Step2",
						HTTPProtocol:              "H1.1",
						SSLProtocol:               "Auto",
						UseAlpn:                   false,
						RequestBody:               "req_param",
						RequestContentType:        "JSON",
						ResponseContentType:       "T",
						OAuth2Provider:            "provider",
						ClientCertificatePassword: "",
						JwtID:                     "111",
						AuthMethod:                "B",
						AuthUser:                  "username",
						AuthPass:                  "",
						RequestHeaders: []api.Header{
							{
								Name:  "Accept",
								Value: "application/json",
							},
							{
								Name:  "Cache-Control",
								Value: "nocache",
							},
						},
						UseNameServer:     true,
						MatchCase:         true,
						JSONSchemaCheck:   false,
						UserAgent:         "firefox",
						MatchingKeyword:   map[string]interface{}{},
						UnmatchingKeyword: map[string]interface{}{},
						MatchRegex:        map[string]interface{}{},
					},
				},
			},
			{
				DisplayName: "Step1",
				StepsDetails: []api.StepDetails{
					{
						StepUrl:                   "www.test.tld",
						Timeout:                   "0",
						DisplayName:               "Step1",
						HTTPMethod:                "G",
						HTTPProtocol:              "H1.1",
						SSLProtocol:               "Auto",
						UseAlpn:                   false,
						RequestBody:               "req_param",
						RequestContentType:        "JSON",
						ResponseContentType:       "T",
						OAuth2Provider:            "provider",
						ClientCertificatePassword: "",
						JwtID:                     "111",
						AuthMethod:                "B",
						AuthUser:                  "username",
						AuthPass:                  "",
						RequestHeaders: []api.Header{
							{
								Name:  "Accept",
								Value: "application/json",
							},
							{
								Name:  "Cache-Control",
								Value: "nocache",
							},
						},
						UseNameServer:     true,
						MatchCase:         true,
						JSONSchemaCheck:   false,
						UserAgent:         "firefox",
						MatchingKeyword:   map[string]interface{}{},
						UnmatchingKeyword: map[string]interface{}{},
						MatchRegex:        map[string]interface{}{},
					},
				},
			},
		},
	}

	locationProfiles := []*api.LocationProfile{
		{
			ProfileID:   "123",
			ProfileName: "Location Profile",
		},
		{
			ProfileID:   "456",
			ProfileName: "TEST",
		},
	}
	c.FakeLocationProfiles.On("List").Return(locationProfiles, nil)

	notificationProfiles := []*api.NotificationProfile{
		{
			ProfileID:   "123",
			ProfileName: "Notifi Profile",
			RcaNeeded:   true,
		},
		{
			ProfileID:   "456",
			ProfileName: "TEST",
			RcaNeeded:   false,
		},
	}
	c.FakeNotificationProfiles.On("List").Return(notificationProfiles, nil)

	userGroups := []*api.UserGroup{
		{
			DisplayName:      "Admin Group",
			Users:            []string{"123", "456"},
			AttributeGroupID: "789",
			ProductID:        0,
		},
		{
			DisplayName:      "Network Group",
			Users:            []string{"123", "456"},
			AttributeGroupID: "345",
			ProductID:        0,
		},
	}
	c.FakeUserGroups.On("List").Return(userGroups, nil)

	tags := []*api.Tag{
		{
			TagID:    "123",
			TagName:  "aws tag",
			TagValue: "baz",
			TagColor: "#B7DA9E",
		},
		{
			TagID:    "456",
			TagName:  "website tag",
			TagValue: "baz 1",
			TagColor: "#B7DA9E",
		},
	}

	c.FakeTags.On("List").Return(tags, nil)

	c.FakeRestApiTransactionMonitors.On("Create", a).Return(a, nil).Once()

	require.NoError(t, restApiTransactionMonitorCreate(d, c))

	c.FakeRestApiTransactionMonitors.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := restApiTransactionMonitorCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestRestApiTransactionMonitorUpdate(t *testing.T) {
	d := restApiTransactionMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.RestApiTransactionMonitor{
		MonitorID:      "123",
		DisplayName:    "foo",
		Type:           string(api.RESTAPISEQ),
		CheckFrequency: "5",
		Steps: []api.Steps{
			{
				DisplayName: "Step2",
				MonitorID:   "123",
				StepsDetails: []api.StepDetails{
					{
						StepUrl:                   "www.test.tld",
						DisplayName:               "Step2",
						Timeout:                   "0",
						HTTPMethod:                "G",
						HTTPProtocol:              "H1.1",
						SSLProtocol:               "Auto",
						UseAlpn:                   false,
						RequestContentType:        "JSON",
						ResponseContentType:       "T",
						RequestBody:               "req_param",
						OAuth2Provider:            "provider",
						ClientCertificatePassword: "",
						JwtID:                     "111",
						UseNameServer:             true,
						MatchCase:                 true,
						JSONSchemaCheck:           false,
						UserAgent:                 "firefox",
						AuthMethod:                "B",
						AuthUser:                  "username",
						AuthPass:                  "",
						MatchingKeyword:           map[string]interface{}{},
						UnmatchingKeyword:         map[string]interface{}{},
						MatchRegex:                map[string]interface{}{},
						RequestHeaders: []api.Header{
							{
								Name:  "Accept",
								Value: "application/json",
							},
							{
								Name:  "Cache-Control",
								Value: "nocache",
							},
						},
					},
				},
			},
			{
				DisplayName: "Step1",
				MonitorID:   "123",
				StepsDetails: []api.StepDetails{
					{
						StepUrl:                   "www.test.tld",
						DisplayName:               "Step1",
						Timeout:                   "0",
						HTTPMethod:                "G",
						HTTPProtocol:              "H1.1",
						SSLProtocol:               "Auto",
						UseAlpn:                   false,
						RequestContentType:        "JSON",
						ResponseContentType:       "T",
						RequestBody:               "req_param",
						OAuth2Provider:            "provider",
						ClientCertificatePassword: "",
						JwtID:                     "111",
						UseNameServer:             true,
						MatchCase:                 true,
						JSONSchemaCheck:           false,
						UserAgent:                 "firefox",
						AuthMethod:                "B",
						AuthUser:                  "username",
						AuthPass:                  "",
						MatchingKeyword:           map[string]interface{}{},
						UnmatchingKeyword:         map[string]interface{}{},
						MatchRegex:                map[string]interface{}{},
						RequestHeaders: []api.Header{
							{
								Name:  "Accept",
								Value: "application/json",
							},
							{
								Name:  "Cache-Control",
								Value: "nocache",
							},
						},
					},
				},
			},
		},
		LocationProfileID:     "456",
		NotificationProfileID: "789",
		ThresholdProfileID:    "012",
		MonitorGroups:         []string{"234", "567"},
		DependencyResourceIDs: []string{"234", "567"},
		UserGroupIDs:          []string{"123", "456"},
		TagIDs:                []string{"123"},
		// ActionIDs: []api.ActionRef{
		// 	{
		// 		ActionID:  "123action",
		// 		AlertType: 1,
		// 	},
		// 	{
		// 		ActionID:  "234action",
		// 		AlertType: 5,
		// 	},
		// },
		// MatchingKeyword: map[string]interface{}{
		// 	"severity": "2",
		// 	"value":    "aaa",
		// },
		// UnmatchingKeyword: map[string]interface{}{
		// 	"severity": "2",
		// 	"value":    "bbb",
		// },
		// MatchRegex: map[string]interface{}{
		// 	"severity": "0",
		// 	"value":    "*.a.*",
		// },
	}

	locationProfiles := []*api.LocationProfile{
		{
			ProfileID:   "123",
			ProfileName: "Location Profile",
		},
		{
			ProfileID:   "456",
			ProfileName: "TEST",
		},
	}
	c.FakeLocationProfiles.On("List").Return(locationProfiles, nil)

	notificationProfiles := []*api.NotificationProfile{
		{
			ProfileID:   "123",
			ProfileName: "Notifi Profile",
			RcaNeeded:   true,
		},
		{
			ProfileID:   "456",
			ProfileName: "TEST",
			RcaNeeded:   false,
		},
	}
	c.FakeNotificationProfiles.On("List").Return(notificationProfiles, nil)

	userGroups := []*api.UserGroup{
		{
			DisplayName:      "Admin Group",
			Users:            []string{"123", "456"},
			AttributeGroupID: "789",
			ProductID:        0,
		},
		{
			DisplayName:      "Network Group",
			Users:            []string{"123", "456"},
			AttributeGroupID: "345",
			ProductID:        0,
		},
	}
	c.FakeUserGroups.On("List").Return(userGroups, nil)

	tags := []*api.Tag{
		{
			TagID:    "123",
			TagName:  "aws tag",
			TagValue: "baz",
			TagColor: "#B7DA9E",
		},
		{
			TagID:    "456",
			TagName:  "website tag",
			TagValue: "baz 1",
			TagColor: "#B7DA9E",
		},
	}
	c.FakeTags.On("List").Return(tags, nil)

	c.FakeRestApiTransactionMonitors.On("Update", a).Return(a, nil).Once()

	require.NoError(t, restApiTransactionMonitorUpdate(d, c))

	c.FakeRestApiTransactionMonitors.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := restApiTransactionMonitorUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestRestApiTransactionMonitorRead(t *testing.T) {
	d := restApiTransactionMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeRestApiTransactionMonitors.On("Get", "123").Return(&api.RestApiMonitor{}, nil).Once()

	require.NoError(t, restApiTransactionMonitorRead(d, c))

	c.FakeRestApiTransactionMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := restApiTransactionMonitorRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestRestApiTransactionMonitorDelete(t *testing.T) {
	d := restApiTransactionMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeRestApiTransactionMonitors.On("Delete", "123").Return(nil).Once()

	require.NoError(t, restApiTransactionMonitorDelete(d, c))

	c.FakeRestApiTransactionMonitors.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, restApiTransactionMonitorDelete(d, c))
}

func TestRestApiTransactionMonitorExists(t *testing.T) {
	d := restApiTransactionMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeRestApiTransactionMonitors.On("Get", "123").Return(&api.RestApiMonitor{}, nil).Once()

	exists, err := restApiTransactionMonitorExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeRestApiTransactionMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = restApiTransactionMonitorExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeRestApiTransactionMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = restApiTransactionMonitorExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func restApiTransactionMonitorTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, RestApiTransactionMonitorSchema, map[string]interface{}{
		"display_name":            "foo",
		"type":                    string(api.RESTAPISEQ),
		"check_frequency":         "5",
		"location_profile_id":     "456",
		"notification_profile_id": "789",
		"threshold_profile_id":    "012",
		"monitor_groups": []interface{}{
			"234",
			"567",
		},
		"dependency_resource_ids": []interface{}{
			"234",
			"567",
		},
		"user_group_ids": []interface{}{
			"123",
			"456",
		},
		"steps": []interface{}{
			map[string]interface{}{
				"display_name": "Step1",
				"step_details": []interface{}{
					map[string]interface{}{
						"step_url":                    "www.test.tld",
						"timeout":                     10,
						"http_method":                 "G",
						"http_protocol":               "H1.1",
						"ssl_protocol":                "Auto",
						"use_alpn":                    false,
						"use_ipv6":                    false,
						"request_content_type":        "JSON",
						"response_content_type":       "T",
						"request_body":                "req_param",
						"auth_user":                   "username",
						"auth_pass":                   "",
						"oauth2_provider":             "provider",
						"client_certificate_password": "",
						"jwt_id":                      "111",
						"match_case":                  true,
						"user_agent":                  "firefox",
						"use_name_server":             true,
						"json_schema_check":           false,
						"request_headers": map[string]interface{}{
							"Accept":        "application/json",
							"Cache-Control": "nocache",
						},
					},
				},
			},
			map[string]interface{}{
				"display_name": "Step2",
				"step_details": []interface{}{
					map[string]interface{}{
						"step_url":                    "www.test.tld",
						"timeout":                     10,
						"http_method":                 "G",
						"http_protocol":               "H1.1",
						"ssl_protocol":                "Auto",
						"use_alpn":                    false,
						"use_ipv6":                    false,
						"request_content_type":        "JSON",
						"response_content_type":       "T",
						"request_body":                "req_param",
						"auth_user":                   "username",
						"auth_pass":                   "",
						"oauth2_provider":             "provider",
						"client_certificate_password": "",
						"jwt_id":                      "111",
						"match_case":                  true,
						"user_agent":                  "firefox",
						"use_name_server":             true,
						"json_schema_check":           false,
						"request_headers": map[string]interface{}{
							"Accept":        "application/json",
							"Cache-Control": "nocache",
						},
					},
				},
			},
		},
		// "actions": map[string]interface{}{
		// 	"1": "123action",
		// 	"5": "234action",
		// },
		// "matching_keyword": map[string]interface{}{
		// 	"severity": "2",
		// 	"value":    "aaa",
		// },
		// "unmatching_keyword": map[string]interface{}{
		// 	"severity": "2",
		// 	"value":    "bbb",
		// },
		// "match_regex": map[string]interface{}{
		// 	"severity": "0",
		// 	"value":    ".*a.*",
		// },
	})
}
