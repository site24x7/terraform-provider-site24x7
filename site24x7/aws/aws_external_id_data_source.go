package aws

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var awsExternalIDDataSourceSchema = map[string]*schema.Schema{}

func DataSourceSite24x7AWSExternalID() *schema.Resource {
	return &schema.Resource{
		Read:   awsExternalIDDataSourceRead,
		Schema: awsExternalIDDataSourceSchema,
	}
}

// awsExternalIDDataSourceRead fetches AWS External ID from Site24x7
func awsExternalIDDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	awsExternalID, err := client.AWSExternalID().Get()
	if err != nil {
		return err
	}

	updateAWSExternalIDDataSourceResourceData(d, awsExternalID)

	return nil
}

func updateAWSExternalIDDataSourceResourceData(d *schema.ResourceData, awsExternalID *api.AWSExternalID) {
	d.SetId(awsExternalID.ID)
}
