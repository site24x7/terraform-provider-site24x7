package site24x7

import (
	"errors"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
)

var tagDataSourceSchema = map[string]*schema.Schema{
	"tag_name_regex": {
		Type:     schema.TypeString,
		Required: true,
		// ValidateFunc: validation.StringIsValidRegExp,
	},
	"tag_value_regex": {
		Type:     schema.TypeString,
		Optional: true,
		// ValidateFunc: validation.StringIsValidRegExp,
	},
	"matching_ids": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "List of tag IDs matching the tag_name_regex and tag_value_regex.",
	},
	"matching_ids_and_names": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "List of tag IDs and names matching the tag_name_regex and tag_value_regex.",
	},
	"tag_name": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Display Name for the Tag.",
	},
	"tag_value": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Value for the Tag.",
	},
	"tag_type": {
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "Type of the Tag.",
	},
	"tag_color": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Tag color code.",
	},
}

func DataSourceSite24x7Tag() *schema.Resource {
	return &schema.Resource{
		Read:   tagDataSourceRead,
		Schema: tagDataSourceSchema,
	}
}

// tagDataSourceRead fetches all tags from Site24x7
func tagDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	tagsList, err := client.Tags().List()
	if err != nil {
		return err
	}
	nameRegex := d.Get("tag_name_regex")
	valueRegex := d.Get("tag_value_regex")
	var tag *api.Tag
	var matchingTagIDs []string
	var matchingTagIDsAndNames []string
	if nameRegex != "" {
		// (?i) - Case insensitive match
		nameRegexPattern := regexp.MustCompile("(?i)" + nameRegex.(string))
		var valueRegexPattern *regexp.Regexp
		if valueRegex != "" {
			valueRegexPattern = regexp.MustCompile("(?i)" + valueRegex.(string))
		}
		for _, tagInfo := range tagsList {

			if nameRegex != "" && valueRegex != "" {
				if len(tagInfo.TagName) > 0 && len(tagInfo.TagValue) > 0 {
					if tag == nil && nameRegexPattern.MatchString(tagInfo.TagName) && valueRegexPattern.MatchString(tagInfo.TagValue) {
						tag = new(api.Tag)
						tag.TagID = tagInfo.TagID
						tag.TagName = tagInfo.TagName
						tag.TagValue = tagInfo.TagValue
						tag.TagType = tagInfo.TagType
						tag.TagColor = tagInfo.TagColor
						matchingTagIDs = append(matchingTagIDs, tagInfo.TagID)
						matchingTagIDsAndNames = append(matchingTagIDsAndNames, tagInfo.TagID+"__"+tagInfo.TagName+":"+tagInfo.TagValue)
					} else if tag != nil && nameRegexPattern.MatchString(tagInfo.TagName) && valueRegexPattern.MatchString(tagInfo.TagValue) {
						matchingTagIDs = append(matchingTagIDs, tagInfo.TagID)
						matchingTagIDsAndNames = append(matchingTagIDsAndNames, tagInfo.TagID+"__"+tagInfo.TagName+":"+tagInfo.TagValue)
					}
				}
			} else {
				if len(tagInfo.TagName) > 0 {
					if tag == nil && nameRegexPattern.MatchString(tagInfo.TagName) {
						tag = new(api.Tag)
						tag.TagID = tagInfo.TagID
						tag.TagName = tagInfo.TagName
						tag.TagValue = tagInfo.TagValue
						tag.TagType = tagInfo.TagType
						tag.TagColor = tagInfo.TagColor
						matchingTagIDs = append(matchingTagIDs, tagInfo.TagID)
						matchingTagIDsAndNames = append(matchingTagIDsAndNames, tagInfo.TagID+"__"+tagInfo.TagName+":"+tagInfo.TagValue)
					} else if tag != nil && nameRegexPattern.MatchString(tagInfo.TagName) {
						matchingTagIDs = append(matchingTagIDs, tagInfo.TagID)
						matchingTagIDsAndNames = append(matchingTagIDsAndNames, tagInfo.TagID+"__"+tagInfo.TagName+":"+tagInfo.TagValue)
					}
				}
			}

		}
	} else {
		return errors.New("Please enter a value for the attribute tag_name_regex!")
	}

	if tag == nil {
		return errors.New("Unable to find tag matching the name : \"" + d.Get("tag_name_regex").(string) + "\" and value : \"" + d.Get("tag_value_regex").(string) + "\"")
	}

	updateTagDataSourceResourceData(d, tag, matchingTagIDs, matchingTagIDsAndNames)

	return nil
}

func updateTagDataSourceResourceData(d *schema.ResourceData, tag *api.Tag, matchingTagIDs []string, matchingTagIDsAndNames []string) {
	d.SetId(tag.TagID)
	d.Set("matching_ids", matchingTagIDs)
	d.Set("matching_ids_and_names", matchingTagIDsAndNames)
	d.Set("tag_name", tag.TagName)
	d.Set("tag_value", tag.TagValue)
	d.Set("tag_type", tag.TagType)
	d.Set("tag_color", tag.TagColor)
}
