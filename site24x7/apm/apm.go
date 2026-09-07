package apm

import (
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apmendpoint "github.com/site24x7/terraform-provider-site24x7/api/endpoints/apm"
)

// timeWindowSchema is shared by every APM data source and resource.
//
// The APM Insight endpoints require a time window in the request path, but
// this provider only reads identity and configuration attributes, which do not
// vary by window. It is exposed so that callers who need a specific window can
// set one, and defaults to the cheapest.
var timeWindowSchema = &schema.Schema{
	Type:        schema.TypeString,
	Optional:    true,
	Default:     apmendpoint.DefaultTimeWindow,
	Description: "Time window used when querying the APM Insight API, for example \"H\" for the last hour. Only affects the request path; the attributes read by this provider are the same for every window.",
}

// instanceElem is the nested shape used for the instances of an application.
func instanceElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique ID of the agent instance.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the agent instance, usually host:port.",
			},
			"host": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Host address the instance reports from.",
			},
			"port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Port number of the instance.",
			},
			"ins_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of APM Insight agent, for example JAVA, PHP, RUBY, DOTNET or NODEJS.",
			},
			"agent_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version of the APM Insight agent deployed at this instance.",
			},
		},
	}
}

// applicationComputedSchema is the set of attributes shared by the application
// data source and the application resource.
//
// Performance metrics returned by the API - response times, apdex, throughput,
// error counts, cpu time - are intentionally absent. They are recomputed on
// every request, so storing them would report drift on every refresh that no
// apply could resolve.
func applicationComputedSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the APM Insight application.",
		},
		"instance_ids": {
			Type:        schema.TypeList,
			Computed:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "IDs of all instances reporting into this application, sorted.",
		},
		"instance_count": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Number of instances reporting into this application.",
		},
		"host_count": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Number of hosts running this application.",
		},
		"hosts": {
			Type:        schema.TypeList,
			Computed:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Host addresses running this application, sorted.",
		},
		"instances": {
			Type:        schema.TypeList,
			Computed:    true,
			Elem:        instanceElem(),
			Description: "Instances reporting into this application, sorted by instance ID.",
		},
		"rum_app_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "ID of the linked Real User Monitoring application, empty when none is linked.",
		},
		"rum_app_key": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Key of the linked Real User Monitoring application. This key is embedded in browser-side JavaScript and is not a secret.",
		},
		"rum_app_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the linked Real User Monitoring application.",
		},
		"availability": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Current availability status, for example AVAILABLE or NOTAVAILABLE.",
		},
		"under_maintenance": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "True when the application is currently under maintenance.",
		},
	}
}

// flattenHosts returns the host addresses in a stable order. The API returns
// them as a JSON object, and Go map iteration order is randomised, so sorting
// is what keeps repeated reads from producing a spurious diff.
func flattenHosts(hosts map[string][]string) []string {
	flattened := make([]string, 0, len(hosts))
	for host := range hosts {
		flattened = append(flattened, host)
	}

	sort.Strings(flattened)
	return flattened
}

// flattenInstances returns the instances sorted by instance ID, for the same
// reason as flattenHosts.
func flattenInstances(instances map[string]api.APMInstanceInfo) []interface{} {
	instanceIDs := make([]string, 0, len(instances))
	for instanceID := range instances {
		instanceIDs = append(instanceIDs, instanceID)
	}
	sort.Strings(instanceIDs)

	flattened := make([]interface{}, 0, len(instances))
	for _, instanceID := range instanceIDs {
		flattened = append(flattened, flattenInstanceInfo(instances[instanceID]))
	}

	return flattened
}

func flattenInstanceInfo(info api.APMInstanceInfo) map[string]interface{} {
	return map[string]interface{}{
		"instance_id":   info.InstanceID,
		"instance_name": info.InstanceName,
		"host":          info.Host,
		"port":          info.Port,
		"ins_type":      info.InstanceType,
		"agent_version": string(info.AgentVersion),
	}
}

// sortedInstanceIDs copies and sorts the instance ID list so that the ordering
// the API happens to return does not leak into state.
func sortedInstanceIDs(instanceIDs []string) []string {
	sorted := make([]string, len(instanceIDs))
	copy(sorted, instanceIDs)
	sort.Strings(sorted)
	return sorted
}

// setApplicationData writes the shared application attributes into state.
func setApplicationData(d *schema.ResourceData, application *api.APMApplication) {
	info := application.ApplicationInfo

	d.Set("application_name", info.ApplicationName)
	d.Set("instance_ids", sortedInstanceIDs(info.InstanceIDs))
	d.Set("instance_count", info.InstanceCount)
	d.Set("host_count", info.HostCount)
	d.Set("hosts", flattenHosts(info.Hosts))
	d.Set("instances", flattenInstances(info.Instances))
	d.Set("rum_app_id", info.RUMInfo.RUMAppID)
	d.Set("rum_app_key", info.RUMInfo.RUMAppKey)
	d.Set("rum_app_name", info.RUMInfo.RUMAppName)
	d.Set("availability", application.AvailabilityHealthInfo.Availability)
	d.Set("under_maintenance", application.AvailabilityHealthInfo.UnderMaintenance)
}
