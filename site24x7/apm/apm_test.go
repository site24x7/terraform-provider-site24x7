package apm

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apmendpoint "github.com/site24x7/terraform-provider-site24x7/api/endpoints/apm"
	"github.com/site24x7/terraform-provider-site24x7/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The API returns hosts and instances as JSON objects. Go randomises map
// iteration order, so without sorting these would land in state in a different
// order on every read and show a diff that no apply could settle.
func TestFlattenHostsIsSorted(t *testing.T) {
	hosts := map[string][]string{
		"10.0.0.3": {"c"},
		"10.0.0.1": {"a"},
		"10.0.0.2": {"b"},
	}

	expected := []string{"10.0.0.1", "10.0.0.2", "10.0.0.3"}

	// Repeat so a lucky map ordering cannot pass by accident.
	for i := 0; i < 25; i++ {
		assert.Equal(t, expected, flattenHosts(hosts))
	}
}

func TestFlattenHostsHandlesEmpty(t *testing.T) {
	assert.Equal(t, []string{}, flattenHosts(nil))
}

func TestFlattenInstancesIsSortedByInstanceID(t *testing.T) {
	instances := map[string]api.APMInstanceInfo{
		"300": {InstanceID: "300", InstanceName: "c", InstanceType: "JAVA", AgentVersion: "1.9"},
		"100": {InstanceID: "100", InstanceName: "a", InstanceType: "JAVA", AgentVersion: "1.7"},
		"200": {InstanceID: "200", InstanceName: "b", InstanceType: "PHP", AgentVersion: "1.8"},
	}

	for i := 0; i < 25; i++ {
		flattened := flattenInstances(instances)

		ids := make([]string, 0, len(flattened))
		for _, instance := range flattened {
			ids = append(ids, instance.(map[string]interface{})["instance_id"].(string))
		}

		assert.Equal(t, []string{"100", "200", "300"}, ids)
	}
}

func TestFlattenInstanceInfo(t *testing.T) {
	info := api.APMInstanceInfo{
		InstanceID:   "101071000000034035",
		InstanceName: "192.168.1.1:8080",
		Host:         "192.168.1.1",
		Port:         8080,
		InstanceType: "JAVA",
		AgentVersion: "1.7",
	}

	assert.Equal(t, map[string]interface{}{
		"instance_id":   "101071000000034035",
		"instance_name": "192.168.1.1:8080",
		"host":          "192.168.1.1",
		"port":          8080,
		"ins_type":      "JAVA",
		"agent_version": "1.7",
	}, flattenInstanceInfo(info))
}

// sortedInstanceIDs must not reorder the caller's slice, which belongs to the
// parsed API response.
func TestSortedInstanceIDsDoesNotMutateInput(t *testing.T) {
	original := []string{"300", "100", "200"}

	assert.Equal(t, []string{"100", "200", "300"}, sortedInstanceIDs(original))
	assert.Equal(t, []string{"300", "100", "200"}, original)
}

func TestTimeWindowOrDefault(t *testing.T) {
	assert.Equal(t, apmendpoint.DefaultTimeWindow, timeWindowOrDefault(""))
	assert.Equal(t, "D", timeWindowOrDefault("D"))
}

// An imported application starts with no time_window in state. Read has to
// record the window it queried, or the first plan after the import proposes
// "" -> "H" and the adoption never reaches a clean plan.
func TestAPMApplicationReadRecordsTimeWindowAfterImport(t *testing.T) {
	c := fake.NewClient()

	// TestResourceDataRaw builds from a config map, so it applies the schema
	// default. Clearing the attribute is what models an imported resource,
	// whose state carries nothing but the ID.
	d := schema.TestResourceDataRaw(t, apmApplicationResourceSchema(), map[string]interface{}{})
	d.SetId("101071000000034001")
	d.Set("time_window", "")
	require.Equal(t, "", d.Get("time_window"), "an imported resource starts with no time_window")

	application := &api.APMApplication{
		ApplicationInfo:        api.APMApplicationInfo{ApplicationID: "101071000000034001", ApplicationName: "checkout-api"},
		AvailabilityHealthInfo: api.APMAvailabilityHealthInfo{ManagedState: true},
	}

	// The endpoint client falls back to the default window for an empty one.
	c.FakeAPMApplications.On("Get", "101071000000034001", "").Return(application, nil).Once()

	require.NoError(t, resourceSite24x7APMApplicationRead(d, c))

	assert.Equal(t, apmendpoint.DefaultTimeWindow, d.Get("time_window"))
	assert.Equal(t, "checkout-api", d.Get("application_name"))
	assert.Equal(t, true, d.Get("managed"))
}

// An explicitly configured window is preserved rather than overwritten with the
// default.
func TestAPMApplicationReadPreservesExplicitTimeWindow(t *testing.T) {
	c := fake.NewClient()

	d := schema.TestResourceDataRaw(t, apmApplicationResourceSchema(), map[string]interface{}{
		"time_window": "D",
	})
	d.SetId("101071000000034001")

	application := &api.APMApplication{
		ApplicationInfo: api.APMApplicationInfo{ApplicationID: "101071000000034001"},
	}

	c.FakeAPMApplications.On("Get", "101071000000034001", "D").Return(application, nil).Once()

	require.NoError(t, resourceSite24x7APMApplicationRead(d, c))

	assert.Equal(t, "D", d.Get("time_window"))
}

func TestAPMApplicationMatches(t *testing.T) {
	managed := &api.APMApplication{
		ApplicationInfo:        api.APMApplicationInfo{ApplicationName: "checkout-api"},
		AvailabilityHealthInfo: api.APMAvailabilityHealthInfo{ManagedState: true},
	}
	unmanaged := &api.APMApplication{
		ApplicationInfo:        api.APMApplicationInfo{ApplicationName: "legacy-batch"},
		AvailabilityHealthInfo: api.APMAvailabilityHealthInfo{ManagedState: false},
	}

	assert.True(t, apmApplicationMatches(managed, nil, "any"))
	assert.True(t, apmApplicationMatches(managed, nil, "managed"))
	assert.False(t, apmApplicationMatches(managed, nil, "unmanaged"))

	assert.True(t, apmApplicationMatches(unmanaged, nil, "unmanaged"))
	assert.False(t, apmApplicationMatches(unmanaged, nil, "managed"))
}

func TestAPMInstanceMatches(t *testing.T) {
	instance := &api.APMInstance{
		InstanceInfo: api.APMInstanceInfo{
			InstanceName:  "192.168.1.1:8080",
			ApplicationID: "app-1",
			InstanceType:  "JAVA",
		},
		AvailabilityHealthInfo: api.APMAvailabilityHealthInfo{ManagedState: true},
	}

	assert.True(t, apmInstanceMatches(instance, nil, "", "", "any"))
	assert.True(t, apmInstanceMatches(instance, nil, "app-1", "", "any"))
	assert.False(t, apmInstanceMatches(instance, nil, "app-2", "", "any"))

	// ins_type is matched case-insensitively, since the API reports it
	// uppercase but practitioners write it however they like.
	assert.True(t, apmInstanceMatches(instance, nil, "", "java", "any"))
	assert.True(t, apmInstanceMatches(instance, nil, "", "JAVA", "any"))
	assert.False(t, apmInstanceMatches(instance, nil, "", "PHP", "any"))
}
