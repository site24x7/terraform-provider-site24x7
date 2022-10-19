package site24x7

import (
	"errors"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
)

var mspDataSourceSchema = map[string]*schema.Schema{
	"customer_name_regex": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Regular expression denoting the name of the MSP customer.",
	},
	"matching_zaaids": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "List of ZAAIDs matching the customer_name_regex.",
	},
	"matching_zaaids_and_names": {
		Type:        schema.TypeList,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: "List of ZAAIDs and names matching the customer_name_regex.",
	},
	"customer_name": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Display name for the MSP customer.",
	},
	"zaaid": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "ZAAID of the MSP customer.",
	},
	"user_id": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "User ID of the MSP customer.",
	},
}

func DataSourceSite24x7MSP() *schema.Resource {
	return &schema.Resource{
		Read:   mspDataSourceRead,
		Schema: mspDataSourceSchema,
	}
}

// mspDataSourceRead fetches all msp from Site24x7
func mspDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(Client)

	mspList, err := client.MSP().List()
	if err != nil {
		return err
	}

	nameRegex := d.Get("customer_name_regex")

	var mspCustomer *api.MSPCustomer
	var matchingZAAIDs []string
	var matchingZAAIDsAndNames []string

	if nameRegex != "" {
		// (?i) - Case insensitive match
		nameRegexPattern := regexp.MustCompile("(?i)" + nameRegex.(string))
		for _, customerInfo := range mspList {
			if len(customerInfo.Name) > 0 {

				if mspCustomer == nil && nameRegexPattern.MatchString(customerInfo.Name) {
					mspCustomer = new(api.MSPCustomer)
					mspCustomer.Name = customerInfo.Name
					mspCustomer.UserID = customerInfo.UserID
					mspCustomer.ZAAID = customerInfo.ZAAID
					matchingZAAIDs = append(matchingZAAIDs, customerInfo.ZAAID)
					matchingZAAIDsAndNames = append(matchingZAAIDsAndNames, customerInfo.ZAAID+"__"+customerInfo.Name)
				} else if mspCustomer != nil && nameRegexPattern.MatchString(customerInfo.Name) {
					matchingZAAIDs = append(matchingZAAIDs, customerInfo.ZAAID)
					matchingZAAIDsAndNames = append(matchingZAAIDsAndNames, customerInfo.ZAAID+"__"+customerInfo.Name)
				}
			}
		}
	} else {
		return errors.New("Please enter a value for the attribute customer_name_regex!")
	}

	if mspCustomer == nil {
		return errors.New("Unable to find MSP customer matching the name : \"" + d.Get("customer_name_regex").(string))
	}
	updateMSPDataSourceResourceData(d, mspCustomer, matchingZAAIDs, matchingZAAIDsAndNames)

	return nil
}

func updateMSPDataSourceResourceData(d *schema.ResourceData, mspCustomer *api.MSPCustomer, matchingZAAIDs []string, matchingZAAIDsAndNames []string) {
	d.SetId(mspCustomer.ZAAID)
	d.Set("matching_zaaids", matchingZAAIDs)
	d.Set("matching_zaaids_and_names", matchingZAAIDsAndNames)
	d.Set("customer_name", mspCustomer.Name)
	d.Set("zaaid", mspCustomer.ZAAID)
	d.Set("user_id", mspCustomer.UserID)
}
