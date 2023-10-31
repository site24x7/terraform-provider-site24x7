package common

import (
	"errors"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var credentialProfileDataSourcceSchema = map[string]*schema.Schema{
	"name_regex": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Regular expression denoting the name of the Credential Profile.",
	},
	"credential_type": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Type of the Credential Profile.",
	},
	"credential_name": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Credential Profile Name.",
	},
	"username": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Username for the Credential Profile.",
	},
	"password": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Password for the Credential Profile.",
	},
}

func DataSourceSite24x7CredentialProfile() *schema.Resource {
	return &schema.Resource{
		Read:   credentialProfileDataSourceRead,
		Schema: credentialProfileDataSourcceSchema,
	}
}

// monitorDataSourceRead fetches all server monitors from Site24x7
func credentialProfileDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	allWebCredentials, err := client.CredentialProfile().ListWebCredentials()
	if err != nil {
		return err
	}

	var genericCredentialProfile *api.CredentialProfile
	if nameRegex, ok := d.GetOk("name_regex"); ok {
		nameRegexPattern := regexp.MustCompile(nameRegex.(string))
		// nameRegexPattern := regexp.MustCompile("(?i)" + nameRegex.(string))
		for _, credentialProfile := range allWebCredentials {
			if len(credentialProfile.CredentialName) > 0 {
				if nameRegexPattern.MatchString(credentialProfile.CredentialName) {
					genericCredentialProfile = new(api.CredentialProfile)
					genericCredentialProfile.ID = credentialProfile.ID
					genericCredentialProfile.CredentialName = credentialProfile.CredentialName
					genericCredentialProfile.CredentialType = credentialProfile.CredentialType
					genericCredentialProfile.UserName = credentialProfile.UserName
					genericCredentialProfile.Password = credentialProfile.Password
				}
			}
		}
	}

	if genericCredentialProfile == nil {
		return errors.New("Unable to find monitor matching the name : \"" + d.Get("name_regex").(string))
	}

	updateResourceData(d, genericCredentialProfile)

	return nil
}

func updateResourceData(d *schema.ResourceData, credentialProfile *api.CredentialProfile) {
	d.SetId(credentialProfile.ID)
	d.Set("credential_name", credentialProfile.CredentialName)
	d.Set("credential_type", credentialProfile.CredentialType)
	d.Set("username", credentialProfile.UserName)
	d.Set("password", credentialProfile.Password)
}
