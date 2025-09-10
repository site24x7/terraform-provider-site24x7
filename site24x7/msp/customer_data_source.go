package msp

import (
	"errors"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
)

var customerDataSourceSchema = map[string]*schema.Schema{
	"name_regex": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Regular expression denoting the display name of the customer.",
	},
	"matching_ids": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "List of customer IDs matching the name_regex.",
	},
	"matching_ids_and_names": {
		Type:        schema.TypeMap,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "Map of customer IDs and display names matching the name_regex.",
	},
	"display_name": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Display name of the customer.",
	},
	"customer_company": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Company name of the customer.",
	},
	"portal_name": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Unique portal name of the customer.",
	},
	"email_address": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Email address of the customer.",
	},
}

func DataSourceSite24x7Customer() *schema.Resource {
	return &schema.Resource{
		Read:   customerDataSourceRead,
		Schema: customerDataSourceSchema,
	}
}

// customerDataSourceRead fetches customers from Site24x7 based on name_regex
func customerDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(site24x7.Client)

	customerList, err := client.Customers().List()
	if err != nil {
		return err
	}

	nameRegex := d.Get("name_regex").(string)
	var matchingCustomerIDs []string
	matchingCustomerIDsAndNames := make(map[string]string)
	var customer *api.Customer

	if nameRegex != "" {
		nameRegexPattern := regexp.MustCompile("(?i)" + nameRegex)
		for _, custInfo := range customerList {
			if nameRegexPattern.MatchString(custInfo.DisplayName) {
				customer = custInfo
				matchingCustomerIDs = append(matchingCustomerIDs, custInfo.UserID)
				matchingCustomerIDsAndNames[custInfo.UserID] = custInfo.DisplayName
			}
		}
	} else {
		return errors.New("Please enter a value for the attribute name_regex!")
	}

	if customer == nil {
		return errors.New("Unable to find customer matching the name: \"" + nameRegex + "\"")
	}

	updateCustomerDataSourceResourceData(d, customer, matchingCustomerIDs, matchingCustomerIDsAndNames)

	return nil
}

func updateCustomerDataSourceResourceData(d *schema.ResourceData, customer *api.Customer, matchingCustomerIDs []string, matchingCustomerIDsAndNames map[string]string) {
	d.SetId(customer.UserID) // Set the ID to the matched customer's ID
	d.Set("matching_ids", matchingCustomerIDs)
	d.Set("matching_ids_and_names", matchingCustomerIDsAndNames)
	d.Set("display_name", customer.DisplayName)
	d.Set("customer_company", customer.CustomerCompany)
	d.Set("portal_name", customer.PortalName)
	d.Set("email_address", customer.EmailAddress)
}
