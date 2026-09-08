package apm

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func agentConfigProfileTestResourceData(t *testing.T) *schema.ResourceData {
	// Every attribute is set explicitly because TestResourceDataRaw does not
	// apply schema defaults.
	return schema.TestResourceDataRaw(t, apmAgentConfigProfileSchema, map[string]interface{}{
		"profile_name":                               "Configuration Profile - Java",
		"agent_type":                                 "JAVA",
		"is_default":                                 false,
		"transaction_trace_enabled":                  true,
		"transaction_trace_threshold":                2,
		"transaction_trace_sql_parametrize":          true,
		"transaction_trace_sql_stacktrace_threshold": 3,
		"transaction_tracking_request_interval":      1,
		"sql_capture_enabled":                        true,
		"autoupgrade_enabled":                        false,
		"show_instance_port_number":                  true,
		"apdex_threshold":                            0.5,
		"cloud_instance_cleanup_threshold":           7,
	})
}

func testAgentConfigProfile(profileID string) *api.APMAgentConfigProfile {
	return &api.APMAgentConfigProfile{
		ProfileID:   profileID,
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
		},
	}
}

func TestAPMAgentConfigProfileCreate(t *testing.T) {
	d := agentConfigProfileTestResourceData(t)
	c := fake.NewClient()

	// The request carries no profile ID, and the ID from the response becomes
	// the Terraform ID.
	request := testAgentConfigProfile("")
	response := testAgentConfigProfile("1000000009001")
	response.AgentConfig.LastModifiedTime = "1563790347139"

	c.FakeAPMAgentConfigProfiles.On("Create", request).Return(response, nil).Once()

	require.NoError(t, resourceSite24x7APMAgentConfigProfileCreate(d, c))

	assert.Equal(t, "1000000009001", d.Id())
	assert.Equal(t, "1563790347139", d.Get("last_modified_time"))

	// A fresh ResourceData, because the successful create above set an ID that
	// would otherwise end up in the next request.
	failing := agentConfigProfileTestResourceData(t)

	c.FakeAPMAgentConfigProfiles.On("Create", request).Return(nil, apierrors.NewStatusError(500, "error")).Once()

	assert.Equal(t, apierrors.NewStatusError(500, "error"), resourceSite24x7APMAgentConfigProfileCreate(failing, c))
}

func TestAPMAgentConfigProfileRead(t *testing.T) {
	d := agentConfigProfileTestResourceData(t)
	d.SetId("1000000009001")

	c := fake.NewClient()

	c.FakeAPMAgentConfigProfiles.On("Get", "1000000009001").Return(testAgentConfigProfile("1000000009001"), nil).Once()

	require.NoError(t, resourceSite24x7APMAgentConfigProfileRead(d, c))
	assert.Equal(t, "1000000009001", d.Id())

	c.FakeAPMAgentConfigProfiles.On("Get", "1000000009001").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	assert.Equal(t, apierrors.NewStatusError(500, "error"), resourceSite24x7APMAgentConfigProfileRead(d, c))
}

// A profile deleted outside Terraform is dropped from state rather than failing
// the refresh, so the next plan proposes recreating it.
func TestAPMAgentConfigProfileReadDropsDeletedProfile(t *testing.T) {
	d := agentConfigProfileTestResourceData(t)
	d.SetId("1000000009001")

	c := fake.NewClient()

	c.FakeAPMAgentConfigProfiles.On("Get", "1000000009001").
		Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, resourceSite24x7APMAgentConfigProfileRead(d, c))
	assert.Equal(t, "", d.Id())
}

func TestAPMAgentConfigProfileUpdate(t *testing.T) {
	d := agentConfigProfileTestResourceData(t)
	d.SetId("1000000009001")

	c := fake.NewClient()

	profile := testAgentConfigProfile("1000000009001")
	c.FakeAPMAgentConfigProfiles.On("Update", profile).Return(profile, nil).Once()

	require.NoError(t, resourceSite24x7APMAgentConfigProfileUpdate(d, c))
	assert.Equal(t, "1000000009001", d.Id())

	c.FakeAPMAgentConfigProfiles.On("Update", profile).Return(nil, apierrors.NewStatusError(500, "error")).Once()

	assert.Equal(t, apierrors.NewStatusError(500, "error"), resourceSite24x7APMAgentConfigProfileUpdate(d, c))
}

func TestAPMAgentConfigProfileDelete(t *testing.T) {
	d := agentConfigProfileTestResourceData(t)
	d.SetId("1000000009001")

	c := fake.NewClient()

	c.FakeAPMAgentConfigProfiles.On("Delete", "1000000009001").Return(nil).Once()
	require.NoError(t, resourceSite24x7APMAgentConfigProfileDelete(d, c))

	// A profile that is already gone is not an error.
	c.FakeAPMAgentConfigProfiles.On("Delete", "1000000009001").
		Return(apierrors.NewStatusError(404, "not found")).Once()
	require.NoError(t, resourceSite24x7APMAgentConfigProfileDelete(d, c))

	c.FakeAPMAgentConfigProfiles.On("Delete", "1000000009001").
		Return(apierrors.NewStatusError(500, "error")).Once()
	assert.Equal(t, apierrors.NewStatusError(500, "error"), resourceSite24x7APMAgentConfigProfileDelete(d, c))
}

// The flattened attribute names are the wire keys with the dots replaced by
// underscores, which is the contract both the resource and the data sources
// rely on.
func TestFlattenAgentConfigProfile(t *testing.T) {
	profile := testAgentConfigProfile("1000000009001")
	profile.AgentConfig.LastModifiedTime = "1563790347139"

	assert.Equal(t, map[string]interface{}{
		"profile_id":                                 "1000000009001",
		"profile_name":                               "Configuration Profile - Java",
		"agent_type":                                 "JAVA",
		"is_default":                                 false,
		"transaction_trace_enabled":                  true,
		"transaction_trace_threshold":                2,
		"transaction_trace_sql_parametrize":          true,
		"transaction_trace_sql_stacktrace_threshold": 3,
		"transaction_tracking_request_interval":      1,
		"sql_capture_enabled":                        true,
		"autoupgrade_enabled":                        false,
		"show_instance_port_number":                  true,
		"apdex_threshold":                            0.5,
		"cloud_instance_cleanup_threshold":           7,
		"last_modified_time":                         "1563790347139",
	}, flattenAgentConfigProfile(profile))
}

// Every key flattenAgentConfigProfile produces has to exist in the data source
// schema, or the list data source would silently drop it.
func TestAgentConfigProfileComputedSchemaCoversFlattenedKeys(t *testing.T) {
	computedSchema := agentConfigProfileComputedSchema()

	for name := range flattenAgentConfigProfile(testAgentConfigProfile("1")) {
		assert.Contains(t, computedSchema, name)
	}

	// ... and the writable schema covers all of them but profile_id, which the
	// resource keeps in the Terraform ID.
	for name := range flattenAgentConfigProfile(testAgentConfigProfile("1")) {
		if name == "profile_id" {
			continue
		}
		assert.Contains(t, apmAgentConfigProfileSchema, name)
	}
}

func TestAPMAgentConfigProfileDataSourceLookupByDefault(t *testing.T) {
	c := fake.NewClient()

	d := schema.TestResourceDataRaw(t, apmAgentConfigProfileDataSourceSchema(), map[string]interface{}{
		"agent_type":      "JAVA",
		"default_profile": true,
	})

	c.FakeAPMAgentConfigProfiles.On("GetDefault", "JAVA").
		Return(testAgentConfigProfile("1000000009003"), nil).Once()

	require.NoError(t, apmAgentConfigProfileDataSourceRead(d, c))

	assert.Equal(t, "1000000009003", d.Id())
	assert.Equal(t, "1000000009003", d.Get("profile_id"))
	assert.Equal(t, "Configuration Profile - Java", d.Get("profile_name"))
}

func TestAPMAgentConfigProfileDataSourceRequiresALookupKey(t *testing.T) {
	c := fake.NewClient()

	d := schema.TestResourceDataRaw(t, apmAgentConfigProfileDataSourceSchema(), map[string]interface{}{})

	assert.EqualError(
		t,
		apmAgentConfigProfileDataSourceRead(d, c),
		"one of profile_id, name_regex or default_profile must be set",
	)
}

func TestAPMAgentConfigProfileDataSourceRequiresAgentTypeForDefault(t *testing.T) {
	c := fake.NewClient()

	d := schema.TestResourceDataRaw(t, apmAgentConfigProfileDataSourceSchema(), map[string]interface{}{
		"default_profile": true,
	})

	assert.EqualError(
		t,
		apmAgentConfigProfileDataSourceRead(d, c),
		"agent_type must be set when default_profile is true",
	)
}

// An ambiguous name is refused rather than resolved by API ordering.
func TestFindAPMAgentConfigProfileByName(t *testing.T) {
	c := fake.NewClient()

	java := testAgentConfigProfile("1000000009001")
	javaDefault := testAgentConfigProfile("1000000009003")
	javaDefault.ProfileName = "Default profile-JAVA"
	javaDefault.IsDefault = true

	// agent_type set narrows the search to the dedicated endpoint.
	c.FakeAPMAgentConfigProfiles.On("ListByAgentType", "JAVA").
		Return([]*api.APMAgentConfigProfile{java, javaDefault}, nil).Once()

	profile, err := findAPMAgentConfigProfileByName(c, "^Default profile", "JAVA")
	require.NoError(t, err)
	assert.Equal(t, "1000000009003", profile.ProfileID)

	c.FakeAPMAgentConfigProfiles.On("List").
		Return([]*api.APMAgentConfigProfile{java, javaDefault}, nil).Once()

	_, err = findAPMAgentConfigProfileByName(c, "JAVA|Java", "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "matched 2 APM agent configuration profiles")

	c.FakeAPMAgentConfigProfiles.On("List").
		Return([]*api.APMAgentConfigProfile{java}, nil).Once()

	_, err = findAPMAgentConfigProfileByName(c, "nothing-matches-this", "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to find an APM agent configuration profile")

	_, err = findAPMAgentConfigProfileByName(c, "([", "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "is not a valid regular expression")
}

func TestAPMAgentConfigProfilesDataSourceIsSortedByProfileID(t *testing.T) {
	c := fake.NewClient()

	d := schema.TestResourceDataRaw(t, apmAgentConfigProfilesDataSourceSchema, map[string]interface{}{})

	first := testAgentConfigProfile("1000000009001")
	second := testAgentConfigProfile("1000000009002")
	second.ProfileName = "Configuration Profile - PHP"
	second.AgentType = "PHP"

	// Returned out of order to prove the data source sorts.
	c.FakeAPMAgentConfigProfiles.On("List").
		Return([]*api.APMAgentConfigProfile{second, first}, nil).Once()

	require.NoError(t, apmAgentConfigProfilesDataSourceRead(d, c))

	assert.Equal(t, []interface{}{"1000000009001", "1000000009002"}, d.Get("ids"))
	assert.Equal(t, []interface{}{
		"1000000009001__Configuration Profile - Java",
		"1000000009002__Configuration Profile - PHP",
	}, d.Get("ids_and_names"))

	profiles := d.Get("profiles").([]interface{})
	require.Len(t, profiles, 2)
	assert.Equal(t, "JAVA", profiles[0].(map[string]interface{})["agent_type"])
	assert.Equal(t, "PHP", profiles[1].(map[string]interface{})["agent_type"])
}

func TestAPMAgentConfigProfilesDataSourceFiltersByAgentTypeAndName(t *testing.T) {
	c := fake.NewClient()

	d := schema.TestResourceDataRaw(t, apmAgentConfigProfilesDataSourceSchema, map[string]interface{}{
		"agent_type": "JAVA",
		"name_regex": "^Default profile",
	})

	java := testAgentConfigProfile("1000000009001")
	javaDefault := testAgentConfigProfile("1000000009003")
	javaDefault.ProfileName = "Default profile-JAVA"

	c.FakeAPMAgentConfigProfiles.On("ListByAgentType", "JAVA").
		Return([]*api.APMAgentConfigProfile{java, javaDefault}, nil).Once()

	require.NoError(t, apmAgentConfigProfilesDataSourceRead(d, c))

	assert.Equal(t, []interface{}{"1000000009003"}, d.Get("ids"))
}

// An empty result is a legitimate answer for the list data source.
func TestAPMAgentConfigProfilesDataSourceAllowsNoMatches(t *testing.T) {
	c := fake.NewClient()

	d := schema.TestResourceDataRaw(t, apmAgentConfigProfilesDataSourceSchema, map[string]interface{}{
		"name_regex": "nothing-matches-this",
	})

	c.FakeAPMAgentConfigProfiles.On("List").
		Return([]*api.APMAgentConfigProfile{testAgentConfigProfile("1000000009001")}, nil).Once()

	require.NoError(t, apmAgentConfigProfilesDataSourceRead(d, c))

	assert.Empty(t, d.Get("ids"))
	assert.NotEmpty(t, d.Id())
}
