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

var apmAgentConfigProfilesDataSourceSchema = map[string]*schema.Schema{
	"agent_type": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Return only the profiles of this agent type, for example JAVA. When omitted, the profiles of every agent type are returned.",
	},
	"name_regex": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Regular expression matched against profile names. When omitted, every profile is returned.",
	},
	// Computed values
	"ids": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "IDs of the matching profiles, sorted.",
	},
	"ids_and_names": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "Matching profiles as \"<id>__<name>\", sorted by ID.",
	},
	"profiles": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Resource{Schema: agentConfigProfileComputedSchema()},
		Description: "Matching profiles with their full configuration, sorted by profile ID.",
	},
}

func DataSourceSite24x7APMAgentConfigProfiles() *schema.Resource {
	return &schema.Resource{
		Read:   apmAgentConfigProfilesDataSourceRead,
		Schema: apmAgentConfigProfilesDataSourceSchema,
	}
}

func apmAgentConfigProfilesDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	agentType := d.Get("agent_type").(string)
	nameRegex := d.Get("name_regex").(string)

	var expression *regexp.Regexp
	if nameRegex != "" {
		var err error
		expression, err = regexp.Compile(nameRegex)
		if err != nil {
			return fmt.Errorf("name_regex %q is not a valid regular expression: %s", nameRegex, err)
		}
	}

	allProfiles, err := listAPMAgentConfigProfiles(client, agentType)
	if err != nil {
		return err
	}

	matches := make([]*api.APMAgentConfigProfile, 0, len(allProfiles))
	for _, profile := range allProfiles {
		if expression != nil && (profile.ProfileName == "" || !expression.MatchString(profile.ProfileName)) {
			continue
		}
		matches = append(matches, profile)
	}

	// The API does not promise an ordering, so sort to keep repeated reads from
	// producing a spurious diff.
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].ProfileID < matches[j].ProfileID
	})

	profileIDs := make([]string, 0, len(matches))
	profileIDsAndNames := make([]string, 0, len(matches))
	profiles := make([]interface{}, 0, len(matches))

	for _, profile := range matches {
		profileIDs = append(profileIDs, profile.ProfileID)
		profileIDsAndNames = append(profileIDsAndNames, profile.ProfileID+"__"+profile.ProfileName)
		profiles = append(profiles, flattenAgentConfigProfile(profile))
	}

	// An empty result is a legitimate answer here, unlike the single-profile
	// data source where the caller asked for one specific thing.
	d.SetId(fmt.Sprintf("%d", hashcode.String(agentType+"__"+nameRegex)))
	d.Set("ids", profileIDs)
	d.Set("ids_and_names", profileIDsAndNames)
	d.Set("profiles", profiles)

	return nil
}
