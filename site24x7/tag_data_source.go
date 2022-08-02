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
	// log.Println("Tags list ============================ ", tagsList)
	nameRegex := d.Get("tag_name_regex")
	valueRegex := d.Get("tag_value_regex")
	var tag *api.Tag
	if nameRegex != "" {
		// (?i) - Case insensitive match
		nameRegexPattern := regexp.MustCompile("(?i)" + nameRegex.(string))
		var valueRegexPattern *regexp.Regexp
		if valueRegex != "" {
			valueRegexPattern = regexp.MustCompile("(?i)" + valueRegex.(string))
		}
		for _, tagInfo := range tagsList {
			// log.Println("Matching Tag name ============================ ", tagInfo.TagName)
			if nameRegex != "" && valueRegex != "" {
				if len(tagInfo.TagName) > 0 && len(tagInfo.TagValue) > 0 {
					if nameRegexPattern.MatchString(tagInfo.TagName) && valueRegexPattern.MatchString(tagInfo.TagValue) {
						tag = new(api.Tag)
						tag.TagID = tagInfo.TagID
						tag.TagName = tagInfo.TagName
						tag.TagValue = tagInfo.TagValue
						tag.TagType = tagInfo.TagType
						tag.TagColor = tagInfo.TagColor
						break
					}
				}
			} else {
				if len(tagInfo.TagName) > 0 {
					if nameRegexPattern.MatchString(tagInfo.TagName) {
						tag = new(api.Tag)
						tag.TagID = tagInfo.TagID
						tag.TagName = tagInfo.TagName
						tag.TagValue = tagInfo.TagValue
						tag.TagType = tagInfo.TagType
						tag.TagColor = tagInfo.TagColor
						break
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

	updateResourceData(d, tag)

	return nil
}

func updateResourceData(d *schema.ResourceData, tag *api.Tag) {
	d.SetId(tag.TagID)
	d.Set("tag_name", tag.TagName)
	d.Set("tag_value", tag.TagValue)
	d.Set("tag_type", tag.TagType)
	d.Set("tag_color", tag.TagColor)
}
