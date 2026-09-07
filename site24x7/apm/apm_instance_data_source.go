package apm

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

func apmInstanceDataSourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"instance_id": {
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{"name_regex"},
			Description:   "ID of the APM Insight instance to look up. Exactly one of instance_id or name_regex must be set.",
		},
		"name_regex": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"instance_id"},
			Description:   "Regular expression matched against instance names. It is an error for the expression to match more than one instance.",
		},
		"time_window": timeWindowSchema,
		// Computed values
		"instance_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the instance, usually host:port.",
		},
		"application_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "ID of the application this instance reports into.",
		},
		"application_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the application this instance reports into.",
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
		"availability": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Current availability status, for example AVAILABLE or NOTAVAILABLE.",
		},
		"managed_state": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "True when the instance is managed rather than suspended.",
		},
		"under_maintenance": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "True when the instance is currently under maintenance.",
		},
		"is_cloud_app": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "True when the instance is hosted in a cloud environment such as AWS or Azure.",
		},
		"auto_scale": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "True when autoscaling is enabled for the instance.",
		},
	}
}

func DataSourceSite24x7APMInstance() *schema.Resource {
	return &schema.Resource{
		Read:   apmInstanceDataSourceRead,
		Schema: apmInstanceDataSourceSchema(),
	}
}

func apmInstanceDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)
	timeWindow := d.Get("time_window").(string)

	instanceID := d.Get("instance_id").(string)
	nameRegex := d.Get("name_regex").(string)

	if instanceID == "" && nameRegex == "" {
		return fmt.Errorf("one of instance_id or name_regex must be set")
	}

	var instance *api.APMInstance
	var err error

	if instanceID != "" {
		instance, err = client.APMInstances().Get(instanceID, timeWindow)
		if err != nil {
			return err
		}
	} else {
		instance, err = findAPMInstanceByName(client, nameRegex, timeWindow)
		if err != nil {
			return err
		}
	}

	info := instance.InstanceInfo

	d.SetId(info.InstanceID)
	d.Set("instance_id", info.InstanceID)
	d.Set("instance_name", info.InstanceName)
	d.Set("application_id", info.ApplicationID)
	d.Set("application_name", info.ApplicationName)
	d.Set("host", info.Host)
	d.Set("port", info.Port)
	d.Set("ins_type", info.InstanceType)
	d.Set("agent_version", string(info.AgentVersion))
	d.Set("availability", instance.AvailabilityHealthInfo.Availability)
	d.Set("managed_state", instance.AvailabilityHealthInfo.ManagedState)
	d.Set("under_maintenance", instance.AvailabilityHealthInfo.UnderMaintenance)
	d.Set("is_cloud_app", instance.AppConfig.IsCloudApp)
	d.Set("auto_scale", instance.AppConfig.AutoScale)

	return nil
}

// findAPMInstanceByName resolves a single instance by name, refusing an
// ambiguous match for the same reason as findAPMApplicationByName.
func findAPMInstanceByName(client site24x7.Client, nameRegex, timeWindow string) (*api.APMInstance, error) {
	expression, err := regexp.Compile(nameRegex)
	if err != nil {
		return nil, fmt.Errorf("name_regex %q is not a valid regular expression: %s", nameRegex, err)
	}

	instances, err := client.APMInstances().List(timeWindow)
	if err != nil {
		return nil, err
	}

	var matches []*api.APMInstance
	var matchedNames []string
	for _, instance := range instances {
		name := instance.InstanceInfo.InstanceName
		if name != "" && expression.MatchString(name) {
			matches = append(matches, instance)
			matchedNames = append(matchedNames, name)
		}
	}

	switch len(matches) {
	case 0:
		return nil, fmt.Errorf("unable to find an APM instance matching name_regex %q", nameRegex)
	case 1:
		return matches[0], nil
	default:
		return nil, fmt.Errorf(
			"name_regex %q matched %d APM instances (%s); refine it so that it matches exactly one",
			nameRegex, len(matches), strings.Join(matchedNames, ", "),
		)
	}
}
