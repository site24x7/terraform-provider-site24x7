package common

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

// Selection types of a maintenance window, as documented for the
// selection_type attribute of the Schedule Maintenances API.
const (
	maintenanceSelectionMonitorGroups api.ResourceType = 1
	maintenanceSelectionMonitors      api.ResourceType = 2
	maintenanceSelectionTags          api.ResourceType = 3
)

var scheduleMaintenancesDataSourceSchema = map[string]*schema.Schema{
	"name_regex": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Regular expression matched against the maintenance display name.",
	},
	"monitor_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Return only maintenances that list this monitor directly (selection_type 2). Maintenances that reach the monitor through a monitor group or tag are not matched.",
	},
	"monitor_group_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Return only maintenances that list this monitor group (selection_type 1).",
	},
	"tag_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Return only maintenances that list this tag (selection_type 3).",
	},
	"maintenance_type": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Return only maintenances of this type (for example 3 = Once). Omit or set 0 to match every type.",
	},
	"maintenance_status": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Return only maintenances with this status code, as returned by the API (for example \"V\").",
	},
	// Computed values
	"ids": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "IDs of the matching maintenances, sorted.",
	},
	"ids_and_names": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "Matching maintenances as \"<id>__<display_name>\", sorted by ID.",
	},
	"maintenances": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Resource{Schema: scheduleMaintenanceDataSourceElemSchema()},
		Description: "Matching maintenances, sorted by maintenance ID.",
	},
}

func scheduleMaintenanceDataSourceElemSchema() map[string]*schema.Schema {
	computed := func(t schema.ValueType) *schema.Schema {
		return &schema.Schema{Type: t, Computed: true}
	}
	computedList := func(t schema.ValueType) *schema.Schema {
		return &schema.Schema{Type: schema.TypeList, Computed: true, Elem: &schema.Schema{Type: t}}
	}

	return map[string]*schema.Schema{
		"maintenance_id":              computed(schema.TypeString),
		"display_name":                computed(schema.TypeString),
		"description":                 computed(schema.TypeString),
		"maintenance_type":            computed(schema.TypeInt),
		"maintenance_status":          computed(schema.TypeString),
		"selection_type":              computed(schema.TypeInt),
		"monitors":                    computedList(schema.TypeString),
		"monitor_groups":              computedList(schema.TypeString),
		"tags":                        computedList(schema.TypeString),
		"subgroup_monitors":           computed(schema.TypeBool),
		"perform_monitoring":          computed(schema.TypeBool),
		"time_zone":                   computed(schema.TypeString),
		"start_date":                  computed(schema.TypeString),
		"end_date":                    computed(schema.TypeString),
		"start_time":                  computed(schema.TypeString),
		"end_time":                    computed(schema.TypeString),
		"duration":                    computed(schema.TypeInt),
		"start_day":                   computed(schema.TypeInt),
		"end_day":                     computed(schema.TypeInt),
		"week_days":                   computedList(schema.TypeInt),
		"execute_every":               computed(schema.TypeInt),
		"start_week":                  computed(schema.TypeInt),
		"start_after":                 computed(schema.TypeInt),
		"monthly_start_date":          computed(schema.TypeInt),
		"maintenance_start_on":        computed(schema.TypeString),
		"maintenance_end_type":        computed(schema.TypeInt),
		"maintenance_end_after_times": computed(schema.TypeInt),
		"maintenance_end_on":          computed(schema.TypeString),
		"maintenance_start_time":      computed(schema.TypeString),
		"zuid":                        computed(schema.TypeString),
	}
}

func DataSourceSite24x7ScheduleMaintenances() *schema.Resource {
	return &schema.Resource{
		Read:   scheduleMaintenancesDataSourceRead,
		Schema: scheduleMaintenancesDataSourceSchema,
	}
}

type scheduleMaintenanceFilter struct {
	nameRegex         *regexp.Regexp
	monitorID         string
	monitorGroupID    string
	tagID             string
	maintenanceType   int
	maintenanceStatus string
}

// scheduleMaintenancesDataSourceRead lists every maintenance and filters on the
// client side: GET /maintenance takes no query parameters.
func scheduleMaintenancesDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	filter := scheduleMaintenanceFilter{
		monitorID:         d.Get("monitor_id").(string),
		monitorGroupID:    d.Get("monitor_group_id").(string),
		tagID:             d.Get("tag_id").(string),
		maintenanceType:   d.Get("maintenance_type").(int),
		maintenanceStatus: d.Get("maintenance_status").(string),
	}

	nameRegex := d.Get("name_regex").(string)
	if nameRegex != "" {
		expression, err := regexp.Compile(nameRegex)
		if err != nil {
			return fmt.Errorf("name_regex %q is not a valid regular expression: %s", nameRegex, err)
		}
		filter.nameRegex = expression
	}

	allMaintenances, err := client.ScheduleMaintenance().List()
	if err != nil {
		return err
	}

	matches := make([]*api.ScheduleMaintenance, 0, len(allMaintenances))
	for _, maintenance := range allMaintenances {
		if scheduleMaintenanceMatches(maintenance, filter) {
			matches = append(matches, maintenance)
		}
	}

	// The API does not promise an ordering, so sort to keep repeated reads
	// from producing a spurious diff.
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].MaintenanceID < matches[j].MaintenanceID
	})

	ids := make([]string, 0, len(matches))
	idsAndNames := make([]string, 0, len(matches))
	maintenances := make([]interface{}, 0, len(matches))
	for _, maintenance := range matches {
		ids = append(ids, maintenance.MaintenanceID)
		idsAndNames = append(idsAndNames, maintenance.MaintenanceID+"__"+maintenance.DisplayName)
		maintenances = append(maintenances, flattenScheduleMaintenance(maintenance))
	}

	// An empty result is a legitimate answer: a monitor with no maintenance
	// window is not an error.
	d.SetId(fmt.Sprintf("%d", hashcode.String(strings.Join([]string{
		nameRegex,
		filter.monitorID,
		filter.monitorGroupID,
		filter.tagID,
		strconv.Itoa(filter.maintenanceType),
		filter.maintenanceStatus,
	}, "__"))))
	d.Set("ids", ids)
	d.Set("ids_and_names", idsAndNames)
	d.Set("maintenances", maintenances)

	return nil
}

// scheduleMaintenanceMatches reports whether maintenance passes every filter
// that is set. Unset filters match everything.
func scheduleMaintenanceMatches(maintenance *api.ScheduleMaintenance, filter scheduleMaintenanceFilter) bool {
	if filter.nameRegex != nil && !filter.nameRegex.MatchString(maintenance.DisplayName) {
		return false
	}
	if filter.maintenanceType != 0 && maintenance.MaintenanceType != filter.maintenanceType {
		return false
	}
	if filter.maintenanceStatus != "" && maintenance.MaintenanceStatus != filter.maintenanceStatus {
		return false
	}
	if filter.monitorID != "" &&
		!(maintenance.SelectionType == maintenanceSelectionMonitors && containsString(maintenance.Monitors, filter.monitorID)) {
		return false
	}
	if filter.monitorGroupID != "" &&
		!(maintenance.SelectionType == maintenanceSelectionMonitorGroups && containsString(maintenance.MonitorGroups, filter.monitorGroupID)) {
		return false
	}
	if filter.tagID != "" &&
		!(maintenance.SelectionType == maintenanceSelectionTags && containsString(maintenance.Tags, filter.tagID)) {
		return false
	}
	return true
}

func flattenScheduleMaintenance(maintenance *api.ScheduleMaintenance) map[string]interface{} {
	// The API documents subgroup_monitors as defaulting to true.
	subgroupMonitors := true
	if maintenance.SubgroupMonitors != nil {
		subgroupMonitors = *maintenance.SubgroupMonitors
	}

	return map[string]interface{}{
		"maintenance_id":              maintenance.MaintenanceID,
		"display_name":                maintenance.DisplayName,
		"description":                 maintenance.Description,
		"maintenance_type":            maintenance.MaintenanceType,
		"maintenance_status":          maintenance.MaintenanceStatus,
		"selection_type":              int(maintenance.SelectionType),
		"monitors":                    maintenance.Monitors,
		"monitor_groups":              maintenance.MonitorGroups,
		"tags":                        maintenance.Tags,
		"subgroup_monitors":           subgroupMonitors,
		"perform_monitoring":          maintenance.PerformMonitoring,
		"time_zone":                   maintenance.TimeZone,
		"start_date":                  maintenance.StartDate,
		"end_date":                    maintenance.EndDate,
		"start_time":                  maintenance.StartTime,
		"end_time":                    maintenance.EndTime,
		"duration":                    maintenanceDurationMinutes(maintenance.Duration),
		"start_day":                   maintenance.StartDay,
		"end_day":                     maintenance.EndDay,
		"week_days":                   maintenance.WeekDays,
		"execute_every":               maintenance.ExecuteEvery,
		"start_week":                  maintenance.StartWeek,
		"start_after":                 maintenance.StartAfter,
		"monthly_start_date":          maintenance.MonthlyStartDate,
		"maintenance_start_on":        maintenance.MaintenanceStartOn,
		"maintenance_end_type":        maintenance.MaintenanceEndType,
		"maintenance_end_after_times": maintenance.MaintenanceEndAfterTimes,
		"maintenance_end_on":          maintenance.MaintenanceEndOn,
		"maintenance_start_time":      maintenance.MaintenanceStartTime,
		"zuid":                        maintenance.ZUID,
	}
}

// maintenanceDurationMinutes normalises duration, which the API returns as a
// number, an empty string or null depending on the maintenance type.
func maintenanceDurationMinutes(duration interface{}) int {
	switch v := duration.(type) {
	case float64:
		return int(v)
	case int:
		return v
	case string:
		if minutes, err := strconv.Atoi(v); err == nil {
			return minutes
		}
	}
	return 0
}

func containsString(values []string, value string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}
