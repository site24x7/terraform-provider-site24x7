package apm

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var apmApplicationsDataSourceSchema = map[string]*schema.Schema{
	"name_regex": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Regular expression matched against application names. When omitted, every APM Insight application is returned.",
	},
	"managed_state": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Filter by managed state. Accepts \"managed\", \"unmanaged\" or \"any\" (the default).",
		Default:     "any",
	},
	"time_window": timeWindowSchema,
	// Computed values
	"ids": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "IDs of the matching applications, sorted.",
	},
	"ids_and_names": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "Matching applications as \"<id>__<name>\", sorted by ID.",
	},
	"applications": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"application_id": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"application_name": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"instance_count": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"host_count": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"availability": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"managed_state": {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"under_maintenance": {
					Type:     schema.TypeBool,
					Computed: true,
				},
			},
		},
		Description: "Matching applications, sorted by application ID.",
	},
}

func DataSourceSite24x7APMApplications() *schema.Resource {
	return &schema.Resource{
		Read:   apmApplicationsDataSourceRead,
		Schema: apmApplicationsDataSourceSchema,
	}
}

func apmApplicationsDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	nameRegex := d.Get("name_regex").(string)
	managedState := d.Get("managed_state").(string)
	timeWindow := d.Get("time_window").(string)

	switch managedState {
	case "managed", "unmanaged", "any":
	default:
		return fmt.Errorf("managed_state must be one of \"managed\", \"unmanaged\" or \"any\", got %q", managedState)
	}

	var expression *regexp.Regexp
	if nameRegex != "" {
		var err error
		expression, err = regexp.Compile(nameRegex)
		if err != nil {
			return fmt.Errorf("name_regex %q is not a valid regular expression: %s", nameRegex, err)
		}
	}

	allApplications, err := client.APMApplications().List(timeWindow)
	if err != nil {
		return err
	}

	matches := make([]*api.APMApplication, 0, len(allApplications))
	for _, application := range allApplications {
		if !apmApplicationMatches(application, expression, managedState) {
			continue
		}
		matches = append(matches, application)
	}

	// The API does not promise an ordering, so sort to keep repeated reads
	// from producing a spurious diff.
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].ApplicationInfo.ApplicationID < matches[j].ApplicationInfo.ApplicationID
	})

	applicationIDs := make([]string, 0, len(matches))
	applicationIDsAndNames := make([]string, 0, len(matches))
	applications := make([]interface{}, 0, len(matches))

	for _, application := range matches {
		info := application.ApplicationInfo
		availabilityHealth := application.AvailabilityHealthInfo

		applicationIDs = append(applicationIDs, info.ApplicationID)
		applicationIDsAndNames = append(applicationIDsAndNames, info.ApplicationID+"__"+info.ApplicationName)
		applications = append(applications, map[string]interface{}{
			"application_id":    info.ApplicationID,
			"application_name":  info.ApplicationName,
			"instance_count":    info.InstanceCount,
			"host_count":        info.HostCount,
			"availability":      availabilityHealth.Availability,
			"managed_state":     availabilityHealth.ManagedState,
			"under_maintenance": availabilityHealth.UnderMaintenance,
		})
	}

	// An empty result is a legitimate answer here, unlike the single-application
	// data source where the caller asked for one specific thing.
	d.SetId(fmt.Sprintf("%d", hashcode.String(nameRegex+"__"+managedState)))
	d.Set("ids", applicationIDs)
	d.Set("ids_and_names", applicationIDsAndNames)
	d.Set("applications", applications)

	return nil
}

func apmApplicationMatches(application *api.APMApplication, expression *regexp.Regexp, managedState string) bool {
	name := application.ApplicationInfo.ApplicationName
	if expression != nil && (name == "" || !expression.MatchString(name)) {
		return false
	}

	switch managedState {
	case "managed":
		return application.AvailabilityHealthInfo.ManagedState
	case "unmanaged":
		return !application.AvailabilityHealthInfo.ManagedState
	}

	return true
}
