package apm

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var apmInstancesDataSourceSchema = map[string]*schema.Schema{
	"name_regex": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Regular expression matched against instance names. When omitted, every APM Insight instance is returned.",
	},
	"application_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Only return instances reporting into this application.",
	},
	"ins_type": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Only return instances running this agent type, for example JAVA or PHP. Matched case-insensitively.",
	},
	"managed_state": {
		Type:        schema.TypeString,
		Optional:    true,
		Default:     "any",
		Description: "Filter by managed state. Accepts \"managed\", \"unmanaged\" or \"any\" (the default).",
	},
	"time_window": timeWindowSchema,
	// Computed values
	"ids": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "IDs of the matching instances, sorted.",
	},
	"ids_and_names": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "Matching instances as \"<id>__<name>\", sorted by ID.",
	},
	"instances": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"instance_id": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"instance_name": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"application_id": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"application_name": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"host": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"port": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"ins_type": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"agent_version": {
					Type:     schema.TypeString,
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
				"is_cloud_app": {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"auto_scale": {
					Type:     schema.TypeBool,
					Computed: true,
				},
			},
		},
		Description: "Matching instances, sorted by instance ID.",
	},
}

func DataSourceSite24x7APMInstances() *schema.Resource {
	return &schema.Resource{
		Read:   apmInstancesDataSourceRead,
		Schema: apmInstancesDataSourceSchema,
	}
}

func apmInstancesDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	nameRegex := d.Get("name_regex").(string)
	applicationID := d.Get("application_id").(string)
	insType := d.Get("ins_type").(string)
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

	allInstances, err := client.APMInstances().List(timeWindow)
	if err != nil {
		return err
	}

	matches := make([]*api.APMInstance, 0, len(allInstances))
	for _, instance := range allInstances {
		if !apmInstanceMatches(instance, expression, applicationID, insType, managedState) {
			continue
		}
		matches = append(matches, instance)
	}

	// The API does not promise an ordering, so sort to keep repeated reads
	// from producing a spurious diff.
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].InstanceInfo.InstanceID < matches[j].InstanceInfo.InstanceID
	})

	instanceIDs := make([]string, 0, len(matches))
	instanceIDsAndNames := make([]string, 0, len(matches))
	instances := make([]interface{}, 0, len(matches))

	for _, instance := range matches {
		info := instance.InstanceInfo

		instanceIDs = append(instanceIDs, info.InstanceID)
		instanceIDsAndNames = append(instanceIDsAndNames, info.InstanceID+"__"+info.InstanceName)
		instances = append(instances, map[string]interface{}{
			"instance_id":       info.InstanceID,
			"instance_name":     info.InstanceName,
			"application_id":    info.ApplicationID,
			"application_name":  info.ApplicationName,
			"host":              info.Host,
			"port":              info.Port,
			"ins_type":          info.InstanceType,
			"agent_version":     string(info.AgentVersion),
			"availability":      instance.AvailabilityHealthInfo.Availability,
			"managed_state":     instance.AvailabilityHealthInfo.ManagedState,
			"under_maintenance": instance.AvailabilityHealthInfo.UnderMaintenance,
			"is_cloud_app":      instance.AppConfig.IsCloudApp,
			"auto_scale":        instance.AppConfig.AutoScale,
		})
	}

	d.SetId(fmt.Sprintf("%d", hashcode.String(strings.Join([]string{nameRegex, applicationID, insType, managedState}, "__"))))
	d.Set("ids", instanceIDs)
	d.Set("ids_and_names", instanceIDsAndNames)
	d.Set("instances", instances)

	return nil
}

func apmInstanceMatches(instance *api.APMInstance, expression *regexp.Regexp, applicationID, insType, managedState string) bool {
	info := instance.InstanceInfo

	if expression != nil && (info.InstanceName == "" || !expression.MatchString(info.InstanceName)) {
		return false
	}

	if applicationID != "" && info.ApplicationID != applicationID {
		return false
	}

	if insType != "" && !strings.EqualFold(info.InstanceType, insType) {
		return false
	}

	switch managedState {
	case "managed":
		return instance.AvailabilityHealthInfo.ManagedState
	case "unmanaged":
		return !instance.AvailabilityHealthInfo.ManagedState
	}

	return true
}
