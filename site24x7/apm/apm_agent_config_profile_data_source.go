package apm

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

func apmAgentConfigProfileDataSourceSchema() map[string]*schema.Schema {
	dataSourceSchema := map[string]*schema.Schema{}
	for name, attributeSchema := range agentConfigProfileComputedSchema() {
		dataSourceSchema[name] = attributeSchema
	}

	// The three lookup keys are writable here, unlike in the shared computed
	// schema they replace.
	dataSourceSchema["profile_id"] = &schema.Schema{
		Type:          schema.TypeString,
		Optional:      true,
		Computed:      true,
		ConflictsWith: []string{"name_regex", "default_profile"},
		Description:   "ID of the configuration profile to look up. Exactly one of profile_id, name_regex or default_profile must be set.",
	}
	dataSourceSchema["name_regex"] = &schema.Schema{
		Type:          schema.TypeString,
		Optional:      true,
		ConflictsWith: []string{"profile_id", "default_profile"},
		Description: "Regular expression matched against profile names. It is an error for the expression to match more than one profile. " +
			"Set agent_type as well to search only the profiles of one agent type.",
	}
	dataSourceSchema["agent_type"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		Description: "Type of APM Insight agent, for example JAVA, DOTNET, PHP, RUBY, NODEJS or PYTHON. " +
			"Required with default_profile, optional with name_regex to narrow the search, and ignored with profile_id.",
	}
	dataSourceSchema["default_profile"] = &schema.Schema{
		Type:          schema.TypeBool,
		Optional:      true,
		ConflictsWith: []string{"profile_id", "name_regex"},
		Description:   "Set to true to look up the default profile of agent_type, which every new agent of that type picks up.",
	}

	return dataSourceSchema
}

func DataSourceSite24x7APMAgentConfigProfile() *schema.Resource {
	return &schema.Resource{
		Read:   apmAgentConfigProfileDataSourceRead,
		Schema: apmAgentConfigProfileDataSourceSchema(),
	}
}

func apmAgentConfigProfileDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	profileID := d.Get("profile_id").(string)
	nameRegex := d.Get("name_regex").(string)
	agentType := d.Get("agent_type").(string)
	defaultProfile := d.Get("default_profile").(bool)

	var profile *api.APMAgentConfigProfile
	var err error

	switch {
	case profileID != "":
		profile, err = client.APMAgentConfigProfiles().Get(profileID)
	case defaultProfile:
		if agentType == "" {
			return fmt.Errorf("agent_type must be set when default_profile is true")
		}
		profile, err = client.APMAgentConfigProfiles().GetDefault(agentType)
	case nameRegex != "":
		profile, err = findAPMAgentConfigProfileByName(client, nameRegex, agentType)
	default:
		return fmt.Errorf("one of profile_id, name_regex or default_profile must be set")
	}

	if err != nil {
		return err
	}

	d.SetId(profile.ProfileID)
	d.Set("profile_id", profile.ProfileID)
	setAgentConfigProfileData(d, profile)

	return nil
}

// findAPMAgentConfigProfileByName resolves a single profile by name, refusing
// an ambiguous match rather than picking one by API ordering, the same way the
// application data source does.
func findAPMAgentConfigProfileByName(client site24x7.Client, nameRegex, agentType string) (*api.APMAgentConfigProfile, error) {
	expression, err := regexp.Compile(nameRegex)
	if err != nil {
		return nil, fmt.Errorf("name_regex %q is not a valid regular expression: %s", nameRegex, err)
	}

	profiles, err := listAPMAgentConfigProfiles(client, agentType)
	if err != nil {
		return nil, err
	}

	var matches []*api.APMAgentConfigProfile
	var matchedNames []string
	for _, profile := range profiles {
		if profile.ProfileName != "" && expression.MatchString(profile.ProfileName) {
			matches = append(matches, profile)
			matchedNames = append(matchedNames, profile.ProfileName)
		}
	}

	switch len(matches) {
	case 0:
		if agentType != "" {
			return nil, fmt.Errorf(
				"unable to find an APM agent configuration profile matching name_regex %q for agent_type %q",
				nameRegex, agentType,
			)
		}
		return nil, fmt.Errorf("unable to find an APM agent configuration profile matching name_regex %q", nameRegex)
	case 1:
		return matches[0], nil
	default:
		return nil, fmt.Errorf(
			"name_regex %q matched %d APM agent configuration profiles (%s); refine it, or set agent_type, so that it matches exactly one",
			nameRegex, len(matches), strings.Join(matchedNames, ", "),
		)
	}
}

// listAPMAgentConfigProfiles lists every profile, or only those of one agent
// type when agentType is set. The API has a dedicated endpoint for the latter,
// which is cheaper than filtering the full list client-side.
func listAPMAgentConfigProfiles(client site24x7.Client, agentType string) ([]*api.APMAgentConfigProfile, error) {
	if agentType != "" {
		return client.APMAgentConfigProfiles().ListByAgentType(agentType)
	}

	return client.APMAgentConfigProfiles().List()
}
