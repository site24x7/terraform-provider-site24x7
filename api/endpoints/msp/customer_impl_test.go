package msp

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/site24x7/terraform-provider-site24x7/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCustomers(t *testing.T) {
	validation.RunTests(t, []*validation.EndpointTest{
		{
			Name:         "create customer",
			ExpectedVerb: "POST",
			ExpectedPath: "/msp/customers",
			ExpectedBody: validation.Fixture(t, "requests/create_customer.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				customer := &api.Customer{
					RoleTitle:       "4",
					Invite:          false,
					CustomerGroups:  []string{"37152000000043029"},
					Digest:          "1_C_797c6d2644b53cb62763de6ba0980fb01a9a10188ae30dbd313ecfbe1c2417f28f469b5447f67a731fb1359731fcd21259416e5e4025f887b3b9656800f22130",
					CustomerCompany: "w3schools",
					DisplayName:     "phillips",
					EmailAddress:    "selvalakshmi.m+aug18@zohotest.com",
					PortalName:      "w3schools123",
					Captcha:         "D6EF1P",
				}

				_, err := NewCustomers(c).Create(customer)
				require.NoError(t, err)
			},
		},
		{
			Name:         "get customer",
			ExpectedVerb: "GET",
			ExpectedPath: "/msp/customers/1234567890",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/get_customer.json"),
			Fn: func(t *testing.T, c rest.Client) {
				customer, err := NewCustomers(c).Get("1234567890")
				require.NoError(t, err)

				expected := &api.Customer{
					UserID:          "1234567890",
					RoleTitle:       "4",
					Invite:          false,
					CustomerGroups:  []string{"37152000000043029"},
					Digest:          "1_C_797c6d2644b53cb62763de6ba0980fb01a9a10188ae30dbd313ecfbe1c2417f28f469b5447f67a731fb1359731fcd21259416e5e4025f887b3b9656800f22130",
					CustomerCompany: "w3schools",
					DisplayName:     "phillips",
					EmailAddress:    "selvalakshmi.m+aug18@zohotest.com",
					PortalName:      "w3schools123",
				}

				assert.Equal(t, expected, customer)
			},
		},
		{
			Name:         "list customers",
			ExpectedVerb: "GET",
			ExpectedPath: "/msp/customers",
			StatusCode:   200,
			ResponseBody: validation.Fixture(t, "responses/list_customers.json"),
			Fn: func(t *testing.T, c rest.Client) {
				customers, err := NewCustomers(c).List()
				require.NoError(t, err)

				expected := []*api.Customer{
					{
						UserID:          "1234567890",
						RoleTitle:       "4",
						Invite:          false,
						CustomerGroups:  []string{"37152000000043029"},
						DisplayName:     "phillips",
						CustomerCompany: "w3schools",
						EmailAddress:    "selvalakshmi.m+aug18@zohotest.com",
						PortalName:      "w3schools123",
					},
				}

				assert.Equal(t, expected, customers)
			},
		},
		{
			Name:         "update customer",
			ExpectedVerb: "PUT",
			ExpectedPath: "/msp/customers/1234567890",
			ExpectedBody: validation.Fixture(t, "requests/update_customer.json"),
			StatusCode:   200,
			ResponseBody: validation.JsonAPIResponseBody(t, nil),
			Fn: func(t *testing.T, c rest.Client) {
				customer := &api.Customer{
					UserID:          "1234567890",
					RoleTitle:       "4",
					Invite:          false,
					CustomerGroups:  []string{"37152000000043029"},
					DisplayName:     "phillips updated",
					CustomerCompany: "w3schools updated",
					EmailAddress:    "selvalakshmi.m+aug18@zohotest.com",
					PortalName:      "w3schools123",
				}

				_, err := NewCustomers(c).Update(customer)
				require.NoError(t, err)
			},
		},
	})
}
