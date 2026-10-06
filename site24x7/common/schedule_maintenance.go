package common

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

// scheduleMaintenanceType describes one supported maintenance_type.
type scheduleMaintenanceType struct {
	name string
	// attributes maps each type-specific attribute the API reads for this
	// type to whether it is required.
	attributes map[string]bool
}

// scheduleMaintenanceTypes is the single source of which type-specific
// attributes each maintenance_type sends and requires, per the Schedule
// Maintenances API. Yearly (9) is not supported: its request attributes are
// not documented.
var scheduleMaintenanceTypes = map[int]scheduleMaintenanceType{
	api.MaintenanceTypeOnce: {"Once", map[string]bool{
		"start_date": true, "end_date": true, "end_time": true,
	}},
	api.MaintenanceTypeDaily: {"Daily", map[string]bool{
		"end_time": true,
	}},
	api.MaintenanceTypeWeeklyByTime: {"Weekly (By Time)", map[string]bool{
		"start_day": true, "end_day": true, "end_time": true, "execute_every": false,
	}},
	api.MaintenanceTypeWeeklyByDay: {"Weekly (By Day)", map[string]bool{
		"week_days": false, "duration": true, "execute_every": false,
	}},
	api.MaintenanceTypeMonthlyByDate: {"Monthly (By Date)", map[string]bool{
		"monthly_start_date": true, "duration": true,
	}},
	api.MaintenanceTypeMonthlyByDay: {"Monthly (By Day)", map[string]bool{
		"start_week": true, "start_day": true, "duration": true, "start_after": false,
	}},
}

// scheduleMaintenanceTypeAttributes lists every attribute that appears in
// scheduleMaintenanceTypes, in the order validation reports them.
var scheduleMaintenanceTypeAttributes = []string{
	"start_date", "end_date", "end_time", "start_day", "end_day", "week_days",
	"duration", "execute_every", "monthly_start_date", "start_week", "start_after",
}

// Values of maintenance_end_type.
const (
	maintenanceEndNever      = 0
	maintenanceEndAfterTimes = 1
	maintenanceEndOnDate     = 2
)

const scheduleMaintenanceDefaultExecuteEvery = 1

var (
	maintenanceDatePattern = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	maintenanceTimePattern = regexp.MustCompile(`^\d{2}:\d{2}$`)
)

var ScheduleMaintenanceSchema = map[string]*schema.Schema{
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name for the maintenance.",
	},
	"description": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Description for the maintenance.",
	},
	"maintenance_type": {
		Type:     schema.TypeInt,
		Required: true,
		ValidateFunc: validation.IntInSlice([]int{
			api.MaintenanceTypeDaily,
			api.MaintenanceTypeWeeklyByTime,
			api.MaintenanceTypeOnce,
			api.MaintenanceTypeMonthlyByDate,
			api.MaintenanceTypeMonthlyByDay,
			api.MaintenanceTypeWeeklyByDay,
		}),
		Description: "1 = Daily, 2 = Weekly (By Time), 3 = Once, 5 = Monthly (By Date), 6 = Monthly (By Day), 8 = Weekly (By Day).",
	},
	"start_time": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringMatch(maintenanceTimePattern, "must be in hh:mm format"),
		Description:  "Maintenance start time. Format: hh:mm.",
	},
	"end_time": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.StringMatch(maintenanceTimePattern, "must be in hh:mm format"),
		Description:  "Maintenance end time. Format: hh:mm. Required for Once, Daily and Weekly (By Time).",
	},
	"time_zone": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Time zone for the maintenance. Defaults to the account time zone.",
	},
	"perform_monitoring": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Perform uptime monitoring of the resources during the maintenance window.",
	},

	// Once
	"start_date": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.StringMatch(maintenanceDatePattern, "must be in yyyy-mm-dd format"),
		Description:  "Maintenance start date. Format: yyyy-mm-dd. Required for Once.",
	},
	"end_date": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.StringMatch(maintenanceDatePattern, "must be in yyyy-mm-dd format"),
		Description:  "Maintenance end date. Format: yyyy-mm-dd. Required for Once.",
	},

	// Weekly and Monthly
	"start_day": {
		Type:         schema.TypeInt,
		Optional:     true,
		ValidateFunc: validation.IntBetween(0, 6),
		Description:  "Day of the week the maintenance starts on (0 = Sunday ... 6 = Saturday). Required for Weekly (By Time) and Monthly (By Day).",
	},
	"end_day": {
		Type:         schema.TypeInt,
		Optional:     true,
		ValidateFunc: validation.IntBetween(0, 6),
		Description:  "Day of the week the maintenance ends on (0 = Sunday ... 6 = Saturday). Required for Weekly (By Time).",
	},
	"week_days": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeInt,
			ValidateFunc: validation.IntBetween(0, 6),
		},
		Description: "Days of the week the maintenance recurs on (0 = Sunday ... 6 = Saturday). Used by Weekly (By Day).",
	},
	"execute_every": {
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      scheduleMaintenanceDefaultExecuteEvery,
		ValidateFunc: validation.IntBetween(1, 4),
		Description:  "Interval in weeks at which a weekly maintenance recurs, 1 to 4 (2 = bi-weekly). Used by Weekly (By Time) and Weekly (By Day).",
	},
	"duration": {
		Type:         schema.TypeInt,
		Optional:     true,
		ValidateFunc: validation.IntAtLeast(1),
		Description:  "Maintenance duration in minutes, under a day. Required for Weekly (By Day), Monthly (By Date) and Monthly (By Day).",
	},
	"monthly_start_date": {
		Type:         schema.TypeInt,
		Optional:     true,
		ValidateFunc: validation.IntBetween(1, 31),
		Description:  "Date of the month the maintenance recurs on. Required for Monthly (By Date).",
	},
	"start_week": {
		Type:         schema.TypeInt,
		Optional:     true,
		ValidateFunc: validation.IntBetween(1, 5),
		Description:  "Week of the month the maintenance recurs in (1 = First ... 4 = Fourth, 5 = Last). Required for Monthly (By Day).",
	},
	"start_after": {
		Type:         schema.TypeInt,
		Optional:     true,
		ValidateFunc: validation.IntBetween(0, 25),
		Description:  "Number of days after which the maintenance begins every month, 0 to 25. Used by Monthly (By Day).",
	},

	// Recurrence (every type except Once)
	"maintenance_start_on": {
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		ValidateFunc: validation.StringMatch(maintenanceDatePattern, "must be in yyyy-mm-dd format"),
		Description:  "Date the recurring maintenance starts on. Format: yyyy-mm-dd. Defaults to the current date in the maintenance time zone.",
	},
	"maintenance_end_type": {
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      maintenanceEndNever,
		ValidateFunc: validation.IntBetween(maintenanceEndNever, maintenanceEndOnDate),
		Description:  "When the recurring maintenance ends: 0 = Never, 1 = after maintenance_end_after_times runs, 2 = on maintenance_end_on.",
	},
	"maintenance_end_after_times": {
		Type:         schema.TypeInt,
		Optional:     true,
		ValidateFunc: validation.IntAtLeast(1),
		Description:  "Number of runs after which the maintenance ends. Required when maintenance_end_type is 1.",
	},
	"maintenance_end_on": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.StringMatch(maintenanceDatePattern, "must be in yyyy-mm-dd format"),
		Description:  "Date the maintenance ends on. Format: yyyy-mm-dd. Required when maintenance_end_type is 2.",
	},

	// Resource selection
	"selection_type": {
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      int(maintenanceSelectionMonitors),
		ValidateFunc: validation.IntBetween(int(maintenanceSelectionMonitorGroups), int(maintenanceSelectionTags)),
		Description:  "Resources the maintenance applies to: 1 = Monitor Groups, 2 = Monitors, 3 = Tags. All Monitors is not supported.",
	},
	"monitors": {
		Type:        schema.TypeList,
		Optional:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "Monitor IDs. Required when selection_type is 2.",
	},
	"monitor_groups": {
		Type:        schema.TypeList,
		Optional:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "Monitor group IDs. Required when selection_type is 1.",
	},
	"tags": {
		Type:        schema.TypeList,
		Optional:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "Tag IDs. Required when selection_type is 3.",
	},
	"subgroup_monitors": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Include the subgroups of the selected monitor groups. Used when selection_type is 1.",
	},
}

func ResourceSite24x7ScheduleMaintenance() *schema.Resource {
	return &schema.Resource{
		Create: scheduleMaintenanceCreate,
		Read:   scheduleMaintenanceRead,
		Update: scheduleMaintenanceUpdate,
		Delete: scheduleMaintenanceDelete,
		Exists: scheduleMaintenanceExists,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		CustomizeDiff: scheduleMaintenanceCustomizeDiff,

		Schema: ScheduleMaintenanceSchema,
	}
}

func scheduleMaintenanceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	scheduleMaintenance := resourceDataToScheduleMaintenance(d)

	scheduleMaintenance, err := client.ScheduleMaintenance().Create(scheduleMaintenance)
	if err != nil {
		return err
	}

	d.SetId(scheduleMaintenance.MaintenanceID)

	return scheduleMaintenanceRead(d, meta)
}

func scheduleMaintenanceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	scheduleMaintenance, err := client.ScheduleMaintenance().Get(d.Id())
	if err != nil {
		return err
	}

	updateScheduleMaintenanceResourceData(d, scheduleMaintenance)
	return nil
}

func scheduleMaintenanceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	scheduleMaintenance := resourceDataToScheduleMaintenance(d)

	_, err := client.ScheduleMaintenance().Update(scheduleMaintenance)
	if err != nil {
		return err
	}

	return scheduleMaintenanceRead(d, meta)
}

func scheduleMaintenanceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	err := client.ScheduleMaintenance().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}
	return err
}

func scheduleMaintenanceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.ScheduleMaintenance().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// scheduleMaintenanceCustomizeDiff rejects at plan time what the API would
// refuse or silently ignore: a type-specific attribute missing for the chosen
// maintenance_type, or set for a type that does not use it. An ignored
// attribute comes back empty from the API and shows a diff on every plan.
// Values not known until apply are skipped.
func scheduleMaintenanceCustomizeDiff(d *schema.ResourceDiff, meta interface{}) error {
	var problems []string

	if d.NewValueKnown("maintenance_type") {
		problems = append(problems, scheduleMaintenanceTypeProblems(d)...)
	}
	if d.NewValueKnown("selection_type") {
		problems = append(problems, scheduleMaintenanceSelectionProblems(d)...)
	}

	if len(problems) == 0 {
		return nil
	}
	return fmt.Errorf("invalid site24x7_schedule_maintenance configuration:\n  - %s", strings.Join(problems, "\n  - "))
}

func scheduleMaintenanceTypeProblems(d *schema.ResourceDiff) []string {
	maintenanceType := d.Get("maintenance_type").(int)
	spec, ok := scheduleMaintenanceTypes[maintenanceType]
	if !ok {
		// Rejected by the maintenance_type ValidateFunc.
		return nil
	}
	forType := fmt.Sprintf("when maintenance_type is %d (%s)", maintenanceType, spec.name)

	var problems []string
	for _, name := range scheduleMaintenanceTypeAttributes {
		if !d.NewValueKnown(name) {
			continue
		}
		required, used := spec.attributes[name]
		switch {
		case required && !scheduleMaintenanceAttributeGiven(d, name):
			problems = append(problems, fmt.Sprintf("%s is required %s", name, forType))
		case !used && scheduleMaintenanceAttributeSet(d, name):
			problems = append(problems, fmt.Sprintf("%s is not used %s", name, forType))
		}
	}

	if maintenanceType == api.MaintenanceTypeOnce {
		if d.Get("maintenance_end_type").(int) != maintenanceEndNever {
			problems = append(problems, "maintenance_end_type is not used "+forType)
		}
		// maintenance_start_on is Computed, so only a configured change counts.
		if _, ok := d.GetOk("maintenance_start_on"); ok && d.HasChange("maintenance_start_on") {
			problems = append(problems, "maintenance_start_on is not used "+forType)
		}
	}

	endType := d.Get("maintenance_end_type").(int)
	for _, end := range []struct {
		name    string
		endType int
	}{
		{"maintenance_end_after_times", maintenanceEndAfterTimes},
		{"maintenance_end_on", maintenanceEndOnDate},
	} {
		if !d.NewValueKnown(end.name) {
			continue
		}
		_, set := d.GetOk(end.name)
		switch {
		case endType == end.endType && !set:
			problems = append(problems, fmt.Sprintf("%s is required when maintenance_end_type is %d", end.name, end.endType))
		case endType != end.endType && set:
			problems = append(problems, fmt.Sprintf("%s is only used when maintenance_end_type is %d", end.name, end.endType))
		}
	}

	return problems
}

func scheduleMaintenanceSelectionProblems(d *schema.ResourceDiff) []string {
	selectionType := api.ResourceType(d.Get("selection_type").(int))

	// Only a list set for the wrong selection_type is checked. A missing list
	// is left to the API: a list that is unknown until apply, such as a module
	// output, looks unset here, so requiring one would reject valid plans.
	var problems []string
	for _, selection := range []struct {
		name          string
		selectionType api.ResourceType
	}{
		{"monitor_groups", maintenanceSelectionMonitorGroups},
		{"monitors", maintenanceSelectionMonitors},
		{"tags", maintenanceSelectionTags},
	} {
		if _, set := d.GetOk(selection.name); set && selectionType != selection.selectionType {
			problems = append(problems, fmt.Sprintf("%s is not used when selection_type is %d", selection.name, selectionType))
		}
	}

	if selectionType != maintenanceSelectionMonitorGroups && !d.Get("subgroup_monitors").(bool) {
		problems = append(problems, fmt.Sprintf("subgroup_monitors is only used when selection_type is %d", maintenanceSelectionMonitorGroups))
	}

	return problems
}

// scheduleMaintenanceAttributeGiven reports whether a required attribute has
// a value. start_day and end_day are days of the week where 0 (Sunday) is a
// real value, so any assignment counts for them.
func scheduleMaintenanceAttributeGiven(d *schema.ResourceDiff, name string) bool {
	if name == "start_day" || name == "end_day" {
		_, ok := d.GetOkExists(name)
		return ok
	}
	_, ok := d.GetOk(name)
	return ok
}

// scheduleMaintenanceAttributeSet reports whether an attribute holds anything
// other than the value Read stores for a type that does not use it.
func scheduleMaintenanceAttributeSet(d *schema.ResourceDiff, name string) bool {
	if name == "execute_every" {
		return d.Get(name).(int) != scheduleMaintenanceDefaultExecuteEvery
	}
	_, ok := d.GetOk(name)
	return ok
}

func resourceDataToScheduleMaintenance(d *schema.ResourceData) *api.ScheduleMaintenance {
	var monitorsIDs, monitorGroupIDs, tagIDs []string

	for _, id := range d.Get("monitors").([]interface{}) {
		monitorsIDs = append(monitorsIDs, id.(string))
	}

	for _, id := range d.Get("monitor_groups").([]interface{}) {
		monitorGroupIDs = append(monitorGroupIDs, id.(string))
	}

	for _, id := range d.Get("tags").([]interface{}) {
		tagIDs = append(tagIDs, id.(string))
	}

	sm := &api.ScheduleMaintenance{
		MaintenanceID:     d.Id(),
		DisplayName:       d.Get("display_name").(string),
		Description:       d.Get("description").(string),
		MaintenanceType:   d.Get("maintenance_type").(int),
		StartTime:         d.Get("start_time").(string),
		TimeZone:          d.Get("time_zone").(string),
		SelectionType:     api.ResourceType(d.Get("selection_type").(int)),
		Monitors:          monitorsIDs,
		MonitorGroups:     monitorGroupIDs,
		Tags:              tagIDs,
		PerformMonitoring: d.Get("perform_monitoring").(bool),
	}

	// Send only the attributes the API reads for this maintenance type.
	attributes := scheduleMaintenanceTypes[sm.MaintenanceType].attributes
	uses := func(name string) bool {
		_, ok := attributes[name]
		return ok
	}

	if uses("start_date") {
		sm.StartDate = d.Get("start_date").(string)
	}
	if uses("end_date") {
		sm.EndDate = d.Get("end_date").(string)
	}
	if uses("end_time") {
		sm.EndTime = d.Get("end_time").(string)
	}
	if uses("start_day") {
		sm.StartDay = d.Get("start_day").(int)
	}
	if uses("end_day") {
		sm.EndDay = d.Get("end_day").(int)
	}
	if uses("week_days") {
		for _, day := range d.Get("week_days").([]interface{}) {
			sm.WeekDays = append(sm.WeekDays, day.(int))
		}
	}
	if uses("duration") {
		sm.Duration = d.Get("duration").(int)
	}
	if uses("execute_every") {
		sm.ExecuteEvery = d.Get("execute_every").(int)
	}
	if uses("monthly_start_date") {
		sm.MonthlyStartDate = d.Get("monthly_start_date").(int)
	}
	if uses("start_week") {
		sm.StartWeek = d.Get("start_week").(int)
	}
	if uses("start_after") {
		sm.StartAfter = d.Get("start_after").(int)
	}

	if sm.MaintenanceType != api.MaintenanceTypeOnce {
		sm.MaintenanceStartOn = d.Get("maintenance_start_on").(string)
		sm.MaintenanceEndType = d.Get("maintenance_end_type").(int)
		switch sm.MaintenanceEndType {
		case maintenanceEndAfterTimes:
			sm.MaintenanceEndAfterTimes = d.Get("maintenance_end_after_times").(int)
		case maintenanceEndOnDate:
			sm.MaintenanceEndOn = d.Get("maintenance_end_on").(string)
		}
	}

	if sm.SelectionType == maintenanceSelectionMonitorGroups {
		subgroupMonitors := d.Get("subgroup_monitors").(bool)
		sm.SubgroupMonitors = &subgroupMonitors
	}

	return sm
}

// updateScheduleMaintenanceResourceData sets every attribute, so import and a
// change of maintenance_type leave nothing stale.
func updateScheduleMaintenanceResourceData(d *schema.ResourceData, sm *api.ScheduleMaintenance) {
	d.Set("display_name", sm.DisplayName)
	d.Set("description", sm.Description)
	d.Set("maintenance_type", sm.MaintenanceType)
	d.Set("start_time", sm.StartTime)
	d.Set("time_zone", sm.TimeZone)
	d.Set("perform_monitoring", sm.PerformMonitoring)

	executeEvery := sm.ExecuteEvery
	if executeEvery == 0 {
		executeEvery = scheduleMaintenanceDefaultExecuteEvery
	}

	// The API also returns values for attributes a type does not take, such
	// as the duration it derives from start_time and end_time on a Once or
	// Daily maintenance. Those are stored as their zero value or schema
	// default, which matches a configuration that leaves them unset.
	attributes := scheduleMaintenanceTypes[sm.MaintenanceType].attributes
	for name, value := range map[string]interface{}{
		"start_date":         sm.StartDate,
		"end_date":           sm.EndDate,
		"end_time":           sm.EndTime,
		"start_day":          sm.StartDay,
		"end_day":            sm.EndDay,
		"week_days":          sm.WeekDays,
		"execute_every":      executeEvery,
		"duration":           maintenanceDurationMinutes(sm.Duration),
		"monthly_start_date": sm.MonthlyStartDate,
		"start_week":         sm.StartWeek,
		"start_after":        sm.StartAfter,
	} {
		if _, used := attributes[name]; !used {
			attribute := ScheduleMaintenanceSchema[name]
			value = attribute.Default
			if value == nil {
				value = attribute.Type.Zero()
			}
		}
		d.Set(name, value)
	}

	d.Set("maintenance_start_on", sm.MaintenanceStartOn)
	d.Set("maintenance_end_type", sm.MaintenanceEndType)
	d.Set("maintenance_end_after_times", sm.MaintenanceEndAfterTimes)
	d.Set("maintenance_end_on", sm.MaintenanceEndOn)

	d.Set("selection_type", int(sm.SelectionType))
	d.Set("monitors", sm.Monitors)
	d.Set("monitor_groups", sm.MonitorGroups)
	d.Set("tags", sm.Tags)
	// The API documents subgroup_monitors as defaulting to true.
	subgroupMonitors := true
	if sm.SubgroupMonitors != nil {
		subgroupMonitors = *sm.SubgroupMonitors
	}
	d.Set("subgroup_monitors", subgroupMonitors)
}
