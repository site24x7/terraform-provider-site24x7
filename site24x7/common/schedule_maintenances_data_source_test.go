package common

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// testScheduleMaintenances returns one maintenance per selection type, listed
// out of ID order to prove the data source sorts its output.
func testScheduleMaintenances() []*api.ScheduleMaintenance {
	subgroupMonitors := false
	return []*api.ScheduleMaintenance{
		{
			MaintenanceID:     "300",
			DisplayName:       "Tag window",
			MaintenanceType:   9,
			MaintenanceStatus: "V",
			SelectionType:     3,
			Tags:              []string{"tag-1"},
			Duration:          "",
		},
		{
			MaintenanceID:     "100",
			DisplayName:       "Weekly patching",
			MaintenanceType:   8,
			MaintenanceStatus: "V",
			SelectionType:     2,
			Monitors:          []string{"mon-1", "mon-2"},
			WeekDays:          []int{3, 5},
			Duration:          float64(30),
			PerformMonitoring: true,
		},
		{
			MaintenanceID:     "200",
			DisplayName:       "Group window",
			MaintenanceType:   3,
			MaintenanceStatus: "C",
			SelectionType:     1,
			MonitorGroups:     []string{"group-1"},
			SubgroupMonitors:  &subgroupMonitors,
		},
		{
			MaintenanceID:     "400",
			DisplayName:       "Another monitor window",
			MaintenanceType:   3,
			MaintenanceStatus: "V",
			SelectionType:     2,
			Monitors:          []string{"mon-2"},
		},
	}
}

func scheduleMaintenancesTestData(t *testing.T, raw map[string]interface{}) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, scheduleMaintenancesDataSourceSchema, raw)
}

func readScheduleMaintenances(t *testing.T, raw map[string]interface{}) *schema.ResourceData {
	d := scheduleMaintenancesTestData(t, raw)
	c := fake.NewClient()
	c.FakeScheduleMaintenance.On("List").Return(testScheduleMaintenances(), nil).Once()

	require.NoError(t, scheduleMaintenancesDataSourceRead(d, c))
	c.FakeScheduleMaintenance.AssertExpectations(t)
	return d
}

func TestScheduleMaintenancesDataSourceNoFilterReturnsAllSorted(t *testing.T) {
	d := readScheduleMaintenances(t, map[string]interface{}{})

	assert.Equal(t, []interface{}{"100", "200", "300", "400"}, d.Get("ids"))
	assert.Equal(t, "100__Weekly patching", d.Get("ids_and_names.0"))
	assert.NotEmpty(t, d.Id())
}

func TestScheduleMaintenancesDataSourceFilters(t *testing.T) {
	cases := []struct {
		name     string
		raw      map[string]interface{}
		expected []interface{}
	}{
		{"by monitor", map[string]interface{}{"monitor_id": "mon-2"}, []interface{}{"100", "400"}},
		{"by monitor group", map[string]interface{}{"monitor_group_id": "group-1"}, []interface{}{"200"}},
		{"by tag", map[string]interface{}{"tag_id": "tag-1"}, []interface{}{"300"}},
		{"by maintenance type", map[string]interface{}{"maintenance_type": 3}, []interface{}{"200", "400"}},
		{"by status", map[string]interface{}{"maintenance_status": "C"}, []interface{}{"200"}},
		{"by name", map[string]interface{}{"name_regex": "^Weekly"}, []interface{}{"100"}},
		{"filters are combined", map[string]interface{}{"monitor_id": "mon-2", "maintenance_type": 3}, []interface{}{"400"}},
		// A monitor ID never matches a group or tag maintenance, even when the
		// same string appears in that maintenance's own list.
		{"monitor id does not match group list", map[string]interface{}{"monitor_id": "group-1"}, []interface{}{}},
		{"no match is not an error", map[string]interface{}{"monitor_id": "unknown"}, []interface{}{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d := readScheduleMaintenances(t, tc.raw)
			assert.Equal(t, tc.expected, d.Get("ids"))
		})
	}
}

func TestScheduleMaintenancesDataSourceFlattensDetails(t *testing.T) {
	d := readScheduleMaintenances(t, map[string]interface{}{})

	// Sorted order: 100, 200, 300, 400.
	assert.Equal(t, "100", d.Get("maintenances.0.maintenance_id"))
	assert.Equal(t, 2, d.Get("maintenances.0.selection_type"))
	assert.Equal(t, []interface{}{"mon-1", "mon-2"}, d.Get("maintenances.0.monitors"))
	assert.Equal(t, []interface{}{3, 5}, d.Get("maintenances.0.week_days"))
	assert.Equal(t, 30, d.Get("maintenances.0.duration"))
	assert.Equal(t, true, d.Get("maintenances.0.perform_monitoring"))
	// subgroup_monitors defaults to true when the API omits it.
	assert.Equal(t, true, d.Get("maintenances.0.subgroup_monitors"))

	assert.Equal(t, []interface{}{"group-1"}, d.Get("maintenances.1.monitor_groups"))
	assert.Equal(t, false, d.Get("maintenances.1.subgroup_monitors"))

	// An empty-string duration is reported as 0.
	assert.Equal(t, 0, d.Get("maintenances.2.duration"))
}

func TestScheduleMaintenancesDataSourceInvalidRegex(t *testing.T) {
	d := scheduleMaintenancesTestData(t, map[string]interface{}{"name_regex": "("})
	c := fake.NewClient()

	err := scheduleMaintenancesDataSourceRead(d, c)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "name_regex")
	// The regex is validated before any API call is made.
	c.FakeScheduleMaintenance.AssertNotCalled(t, "List")
}

func TestScheduleMaintenancesDataSourceListError(t *testing.T) {
	d := scheduleMaintenancesTestData(t, map[string]interface{}{})
	c := fake.NewClient()
	c.FakeScheduleMaintenance.On("List").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := scheduleMaintenancesDataSourceRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestMaintenanceDurationMinutes(t *testing.T) {
	assert.Equal(t, 30, maintenanceDurationMinutes(float64(30)))
	assert.Equal(t, 45, maintenanceDurationMinutes(45))
	assert.Equal(t, 60, maintenanceDurationMinutes("60"))
	assert.Equal(t, 0, maintenanceDurationMinutes(""))
	assert.Equal(t, 0, maintenanceDurationMinutes(nil))
}
