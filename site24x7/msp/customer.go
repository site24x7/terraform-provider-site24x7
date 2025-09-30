package msp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

// Terraform schema for Customer
var CustomerSchema = map[string]*schema.Schema{
	"country_code": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Country code of the customer (e.g., US, IN).",
	},
	"timezone": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Timezone of the customer (e.g., Asia/Kolkata).",
	},
	"language_code": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Language code (e.g., en, fr).",
	},
	"industry": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Industry type identifier.",
	},
	"roletitle": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Role title assigned for customer.",
	},
	"invite": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Whether to invite the customer.",
	},
	"customer_groups": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of customer group IDs.",
	},
	"digest": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Digest token for authentication.",
	},
	"zuids": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "List of Zoho User IDs (zuids).",
	},
	"customer_company": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Customer's company name.",
	},
	"display_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the customer.",
	},
	"customer_website": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Website of the customer.",
	},
	"email_address": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Email address of the customer.",
	},
	"portal_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Portal name for the customer.",
	},
	"captcha": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Captcha validation value.",
	},
}

func ResourceSite24x7Customer() *schema.Resource {
	return &schema.Resource{
		Create: customerCreate,
		Read:   customerRead,
		Update: customerUpdate,
		Delete: customerDelete,
		Exists: customerExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: CustomerSchema,
	}
}

func customerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	customer := resourceDataToCustomer(d)
	customer, err := client.Customers().Create(customer)
	if err != nil {
		return err
	}

	d.SetId(customer.UserID)

	return nil
}

func customerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	customer, err := client.Customers().Get(d.Id())
	if err != nil {
		return err
	}

	updateCustomerResourceData(d, customer)

	return nil
}

func customerUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	customer := resourceDataToCustomer(d)
	// log.Printf("[DEBUG] Customer before update: %+v", customer)
	customer, err := client.Customers().Update(customer)
	if err != nil {
		return err
	}

	d.SetId(customer.UserID)
	// log.Printf("[DEBUG] Customer after update: %+v", customer)
	return nil
}

func customerDelete(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func customerExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(site24x7.Client)

	_, err := client.Customers().Get(d.Id())
	if apierrors.IsNotFound(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

// Mapping Terraform ResourceData to API Customer struct
func resourceDataToCustomer(d *schema.ResourceData) *api.Customer {
	var customerGroups []string
	for _, id := range d.Get("customer_groups").(*schema.Set).List() {
		if id != nil {
			customerGroups = append(customerGroups, id.(string))
		}
	}

	var zuids []string
	for _, id := range d.Get("zuids").(*schema.Set).List() {
		if id != nil {
			zuids = append(zuids, id.(string))
		}
	}

	return &api.Customer{
		UserID:          d.Id(),
		CountryCode:     d.Get("country_code").(string),
		Timezone:        d.Get("timezone").(string),
		LanguageCode:    d.Get("language_code").(string),
		Industry:        d.Get("industry").(string),
		RoleTitle:       d.Get("roletitle").(string),
		Invite:          d.Get("invite").(bool),
		CustomerGroups:  customerGroups,
		Digest:          d.Get("digest").(string),
		Zuids:           zuids,
		CustomerCompany: d.Get("customer_company").(string),
		DisplayName:     d.Get("display_name").(string),
		CustomerWebsite: d.Get("customer_website").(string),
		EmailAddress:    d.Get("email_address").(string),
		PortalName:      d.Get("portal_name").(string),
		Captcha:         d.Get("captcha").(string),
	}
}

// Called during read - populates Terraform state from API response
func updateCustomerResourceData(d *schema.ResourceData, customer *api.Customer) {
	d.Set("country_code", customer.CountryCode)
	d.Set("timezone", customer.Timezone)
	d.Set("language_code", customer.LanguageCode)
	d.Set("industry", customer.Industry)
	d.Set("roletitle", customer.RoleTitle)
	d.Set("invite", customer.Invite)
	d.Set("customer_groups", customer.CustomerGroups)
	d.Set("digest", customer.Digest)
	d.Set("zuids", customer.Zuids)
	d.Set("customer_company", customer.CustomerCompany)
	d.Set("display_name", customer.DisplayName)
	d.Set("customer_website", customer.CustomerWebsite)
	d.Set("email_address", customer.EmailAddress)
	d.Set("portal_name", customer.PortalName)
	d.Set("captcha", customer.Captcha)
}
