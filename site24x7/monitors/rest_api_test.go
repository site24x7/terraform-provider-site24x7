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

func TestRestApiMonitorCreate(t *testing.T) {
	d := restApiMonitorTestResourceData(t)

	c := fake.NewClient()

	a := &api.RestApiMonitor{
		DisplayName:               "foo",
		Type:                      string(api.RESTAPI),
		Website:                   "www.test.tld",
		CheckFrequency:            "5",
		Timeout:                   10,
		HttpMethod:                "G",
		HttpProtocol:              "H1.1",
		SslProtocol:               "Auto",
		UseAlpn:                   false,
		UseIPV6:                   false,
		RequestParam:              "req_param",
		RequestContentType:        "JSON",
		ResponseContentType:       "T",
		OAuth2Provider:            "provider",
		ClientCertificatePassword: "pass",
		JwtID:                     "111",
		LocationProfileID:         "456",
		NotificationProfileID:     "789",
		ThresholdProfileID:        "012",
		UseNameServer:             true,
		MatchCase:                 true,
		JSONSchemaCheck:           false,
		UserAgent:                 "firefox",
		MonitorGroups:             []string{"234", "567"},
		DependencyResourceIDs:     []string{"234", "567"},
		UserGroupIDs:              []string{"123", "456"},
		TagIDs:                    []string{"123"},
		AuthUser:                  "username",
		AuthPass:                  "password",
		CustomHeaders: []api.Header{
			{
				Name:  "Accept",
				Value: "application/json",
			},
			{
				Name:  "Cache-Control",
				Value: "nocache",
			},
		},

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

	c.FakeRestApiMonitors.On("Create", a).Return(a, nil).Once()

	require.NoError(t, restApiMonitorCreate(d, c))

	c.FakeRestApiMonitors.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := restApiMonitorCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestRestApiMonitorUpdate(t *testing.T) {
	d := restApiMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.RestApiMonitor{
		MonitorID:                 "123",
		DisplayName:               "foo",
		Type:                      string(api.RESTAPI),
		Website:                   "www.test.tld",
		CheckFrequency:            "5",
		Timeout:                   10,
		HttpMethod:                "G",
		HttpProtocol:              "H1.1",
		SslProtocol:               "Auto",
		UseAlpn:                   false,
		UseIPV6:                   false,
		RequestContentType:        "JSON",
		ResponseContentType:       "T",
		RequestParam:              "req_param",
		OAuth2Provider:            "provider",
		ClientCertificatePassword: "pass",
		JwtID:                     "111",
		LocationProfileID:         "456",
		NotificationProfileID:     "789",
		ThresholdProfileID:        "012",
		UseNameServer:             true,
		MatchCase:                 true,
		JSONSchemaCheck:           false,
		UserAgent:                 "firefox",
		MonitorGroups:             []string{"234", "567"},
		DependencyResourceIDs:     []string{"234", "567"},
		UserGroupIDs:              []string{"123", "456"},
		TagIDs:                    []string{"123"},
		AuthUser:                  "username",
		AuthPass:                  "password",
		CustomHeaders: []api.Header{
			{
				Name:  "Accept",
				Value: "application/json",
			},
			{
				Name:  "Cache-Control",
				Value: "nocache",
			},
		},
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

	c.FakeRestApiMonitors.On("Update", a).Return(a, nil).Once()

	require.NoError(t, restApiMonitorUpdate(d, c))

	c.FakeRestApiMonitors.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := restApiMonitorUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestRestApiMonitorRead(t *testing.T) {
	d := restApiMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeRestApiMonitors.On("Get", "123").Return(&api.RestApiMonitor{}, nil).Once()

	require.NoError(t, restApiMonitorRead(d, c))

	c.FakeRestApiMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := restApiMonitorRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestRestApiMonitorDelete(t *testing.T) {
	d := restApiMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeRestApiMonitors.On("Delete", "123").Return(nil).Once()

	require.NoError(t, restApiMonitorDelete(d, c))

	c.FakeRestApiMonitors.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, restApiMonitorDelete(d, c))
}

func TestRestApiMonitorExists(t *testing.T) {
	d := restApiMonitorTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeRestApiMonitors.On("Get", "123").Return(&api.RestApiMonitor{}, nil).Once()

	exists, err := restApiMonitorExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeRestApiMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = restApiMonitorExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeRestApiMonitors.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = restApiMonitorExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func restApiMonitorTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, RestApiMonitorSchema, map[string]interface{}{
		"display_name":                "foo",
		"type":                        string(api.RESTAPI),
		"website":                     "www.test.tld",
		"check_frequency":             "5",
		"timeout":                     10,
		"http_method":                 "G",
		"http_protocol":               "H1.1",
		"ssl_protocol":                "Auto",
		"use_alpn":                    false,
		"use_ipv6":                    false,
		"request_content_type":        "JSON",
		"response_content_type":       "T",
		"request_param":               "req_param",
		"auth_user":                   "username",
		"auth_pass":                   "password",
		"oauth2_provider":             "provider",
		"client_certificate_password": "pass",
		"jwt_id":                      "111",
		"match_case":                  true,
		"user_agent":                  "firefox",
		"custom_headers": map[string]interface{}{
			"Accept":        "application/json",
			"Cache-Control": "nocache",
		},
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
		"use_name_server":   true,
		"json_schema_check": false,
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
