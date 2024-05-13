package common

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var credentialProfileSchema = map[string]*schema.Schema{
	"credential_type": {
		Type:        schema.TypeInt,
		Description: "Type of the Credential Profile.",
		Required:    true,
	},
	"credential_name": {
		Type:        schema.TypeString,
		Description: "Credential Profile Name.",
		Required:    true,
	},
	"username": {
		Type:        schema.TypeString,
		Description: "Username for the Credential Profile.",
		Required:    true,
	},
	"password": {
		Type:        schema.TypeString,
		Description: "Password for the Credential Profile.",
		Required:    true,
	},
}

func ResourceSite24x7CredentialProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSite24x7CredentialProfileCreate,
		Read:   resourceSite24x7CredentialProfileRead,
		Update: resourceSite24x7CredentialProfileUpdate,
		Delete: resourceSite24x7CredentialProfileDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: credentialProfileSchema,
	}
}

func resourceSite24x7CredentialProfileCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client) // Replace with the actual Site24x7 client initialization

	credentilProfile, err := resourceDataToCredentialProfile(d)

	if err != nil {
		return err
	}

	credentialProfile, err := client.CredentialProfile().Create(credentilProfile)
	if err != nil {
		return err
	}

	d.SetId(credentialProfile.ID)

	return nil
}

func resourceSite24x7CredentialProfileRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client) // Replace with the actual Site24x7 client initialization

	// Read the resource ID from the Terraform state
	credentilProfileID := d.Id()

	credentilProfile, err := client.CredentialProfile().Get(credentilProfileID)
	if err != nil {
		return err
	}

	d.Set("credential_type", credentilProfile.CredentialType)
	d.Set("credential_name", credentilProfile.CredentialName)
	d.Set("username", credentilProfile.UserName)
	d.Set("password", credentilProfile.Password)

	return nil
}

func resourceSite24x7CredentialProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client) // Replace with the actual Site24x7 client initialization

	credentilProfile, err := resourceDataToCredentialProfile(d)

	if err != nil {
		return err
	}

	credentilProfile, err = client.CredentialProfile().Update(credentilProfile)

	if err != nil {
		return err
	}
	d.SetId(credentilProfile.ID)

	return nil
}

func resourceSite24x7CredentialProfileDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client) // Replace with the actual Site24x7 client initialization

	// Read the resource ID from the Terraform state
	credentilProfileID := d.Id()

	// Implement the logic to delete the credential profile using the Site24x7 API
	// Example:
	err := client.CredentialProfile().Delete(credentilProfileID)
	if err != nil {
		return err
	}

	// Mark the resource as deleted
	d.SetId("")

	return nil
}

func resourceDataToCredentialProfile(d *schema.ResourceData) (*api.CredentialProfile, error) {

	credentialType := d.Get("credential_type").(int)
	credentialName := d.Get("credential_name").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	credentilProfile := &api.CredentialProfile{
		ID:             d.Id(),
		CredentialType: credentialType,
		CredentialName: credentialName,
		UserName:       username,
		Password:       password,
	}
	return credentilProfile, nil
}
