package apm

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

func apmApplicationDataSourceSchema() map[string]*schema.Schema {
	dataSourceSchema := map[string]*schema.Schema{
		"application_id": {
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{"name_regex"},
			Description:   "ID of the APM Insight application to look up. Exactly one of application_id or name_regex must be set.",
		},
		"name_regex": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"application_id"},
			Description:   "Regular expression matched against application names. It is an error for the expression to match more than one application.",
		},
		"time_window": timeWindowSchema,
		"managed_state": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "True when the application is managed rather than suspended.",
		},
	}

	for name, attributeSchema := range applicationComputedSchema() {
		dataSourceSchema[name] = attributeSchema
	}

	return dataSourceSchema
}

func DataSourceSite24x7APMApplication() *schema.Resource {
	return &schema.Resource{
		Read:   apmApplicationDataSourceRead,
		Schema: apmApplicationDataSourceSchema(),
	}
}

func apmApplicationDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)
	timeWindow := d.Get("time_window").(string)

	applicationID := d.Get("application_id").(string)
	nameRegex := d.Get("name_regex").(string)

	if applicationID == "" && nameRegex == "" {
		return fmt.Errorf("one of application_id or name_regex must be set")
	}

	var application *api.APMApplication
	var err error

	if applicationID != "" {
		application, err = client.APMApplications().Get(applicationID, timeWindow)
		if err != nil {
			return err
		}
	} else {
		application, err = findAPMApplicationByName(client, nameRegex, timeWindow)
		if err != nil {
			return err
		}
	}

	d.SetId(application.ApplicationInfo.ApplicationID)
	d.Set("application_id", application.ApplicationInfo.ApplicationID)
	d.Set("managed_state", application.AvailabilityHealthInfo.ManagedState)
	setApplicationData(d, application)

	return nil
}

// findAPMApplicationByName resolves a single application by name. Unlike some
// of the older data sources in this provider it refuses an ambiguous match
// rather than silently picking one, since which one you got would otherwise
// depend on API ordering.
func findAPMApplicationByName(client site24x7.Client, nameRegex, timeWindow string) (*api.APMApplication, error) {
	expression, err := regexp.Compile(nameRegex)
	if err != nil {
		return nil, fmt.Errorf("name_regex %q is not a valid regular expression: %s", nameRegex, err)
	}

	applications, err := client.APMApplications().List(timeWindow)
	if err != nil {
		return nil, err
	}

	var matches []*api.APMApplication
	var matchedNames []string
	for _, application := range applications {
		name := application.ApplicationInfo.ApplicationName
		if name != "" && expression.MatchString(name) {
			matches = append(matches, application)
			matchedNames = append(matchedNames, name)
		}
	}

	switch len(matches) {
	case 0:
		return nil, fmt.Errorf("unable to find an APM application matching name_regex %q", nameRegex)
	case 1:
		return matches[0], nil
	default:
		return nil, fmt.Errorf(
			"name_regex %q matched %d APM applications (%s); refine it so that it matches exactly one",
			nameRegex, len(matches), strings.Join(matchedNames, ", "),
		)
	}
}
