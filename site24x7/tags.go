package site24x7

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
)

var TagSchema = map[string]*schema.Schema{
	"tag_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display Name for the Tag.",
	},
	"tag_value": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Value for the Tag.",
	},
	"tag_color": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Color code for the Tag. Possible values are '#B7DA9E','#73C7A3','#B5DCDF','#D4ABBB','#4895A8','#DFE897','#FCEA8B','#FFC36D','#F79953','#F16B3C','#E55445','#F2E2B6','#DEC57B','#CBBD80','#AAB3D4','#7085BA','#F6BDAE','#EFAB6D','#CA765C','#999','#4A148C','#009688','#00ACC1','#0091EA','#8BC34A','#558B2F'",
	},
}

func ResourceSite24x7Tag() *schema.Resource {
	return &schema.Resource{
		Create: tagCreate,
		Read:   tagRead,
		Update: tagUpdate,
		Delete: tagDelete,
		Exists: tagExists,

		Schema: TagSchema,
	}
}

func tagCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	tag, err := resourceDataToTagCreateOrUpdate(d)

	tag, err = client.Tags().Create(tag)
	if err != nil {
		return err
	}

	d.SetId(tag.TagID)

	// Read is called for updating state after modification
	// https://www.terraform.io/docs/extend/best-practices/detecting-drift.html#update-state-after-modification
	// return tagRead(d, meta)
	return nil
}

func tagRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	tag, err := client.Tags().Get(d.Id())
	if err != nil {
		return err
	}

	updateTagResourceData(d, tag)

	return nil
}

func tagUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	tag, err := resourceDataToTagCreateOrUpdate(d)

	tag, err = client.Tags().Update(tag)
	if err != nil {
		return err
	}

	d.SetId(tag.TagID)

	// Read is called for updating state after modification
	// https://www.terraform.io/docs/extend/best-practices/detecting-drift.html#update-state-after-modification
	// return tagRead(d, meta)
	return nil
}

func tagDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	err := client.Tags().Delete(d.Id())
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}

func tagExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(Client)

	_, err := client.Tags().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceDataToTagCreateOrUpdate(d *schema.ResourceData) (*api.Tag, error) {
	return &api.Tag{
		TagID:    d.Id(),
		TagName:  d.Get("tag_name").(string),
		TagValue: d.Get("tag_value").(string),
		TagColor: d.Get("tag_color").(string),
		TagType:  1,
	}, nil
}

func updateTagResourceData(d *schema.ResourceData, tag *api.Tag) {
	d.Set("tag_name", tag.TagName)
	d.Set("tag_value", tag.TagValue)
	d.Set("tag_color", tag.TagColor)
	d.Set("tag_type", tag.TagType)
}
