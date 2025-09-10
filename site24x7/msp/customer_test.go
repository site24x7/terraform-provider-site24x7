package msp

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCustomerCreate(t *testing.T) {
	d := customerTestResourceData(t)

	c := fake.NewClient()

	a := &api.Customer{
		CountryCode:     "US",
		Timezone:        "Asia/Kolkata",
		LanguageCode:    "en",
		Industry:        "15",
		RoleTitle:       "4",
		Invite:          false,
		CustomerGroups:  []string{"37152000000043029"},
		Digest:          "s247string",
		Zuids:           []string{"75086549"},
		CustomerCompany: "w3schools",
		DisplayName:     "phillips",
		CustomerWebsite: "https://www.w3schools.com",
		EmailAddress:    "selvalakshmi.m+aug18@zohotest.com",
		PortalName:      "w3schools",
		Captcha:         "D6EF1P",
	}

	c.FakeCustomer.On("Create", a).Return(a, nil).Once()

	require.NoError(t, customerCreate(d, c))

	c.FakeCustomer.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := customerCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestCustomerUpdate(t *testing.T) {
	d := customerTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.Customer{
		UserID:          "123",
		CountryCode:     "US",
		Timezone:        "Asia/Kolkata",
		LanguageCode:    "en",
		Industry:        "15",
		RoleTitle:       "4",
		Invite:          false,
		CustomerGroups:  []string{"37152000000043029"},
		Digest:          "s247string",
		Zuids:           []string{"75086549"},
		CustomerCompany: "w3schools",
		DisplayName:     "phillips",
		CustomerWebsite: "https://www.w3schools.com",
		EmailAddress:    "selvalakshmi.m+aug18@zohotest.com",
		PortalName:      "w3schools",
		Captcha:         "D6EF1P",
	}

	c.FakeCustomer.On("Update", a).Return(a, nil).Once()

	require.NoError(t, customerUpdate(d, c))

	c.FakeCustomer.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := customerUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestCustomerRead(t *testing.T) {
	d := customerTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeCustomer.On("Get", "123").Return(&api.Customer{}, nil).Once()

	require.NoError(t, customerRead(d, c))

	c.FakeCustomer.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := customerRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestCustomerExists(t *testing.T) {
	d := customerTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeCustomer.On("Get", "123").Return(&api.Customer{}, nil).Once()

	exists, err := customerExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeCustomer.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = customerExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeCustomer.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = customerExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func customerTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, CustomerSchema, map[string]interface{}{
		"country_code":     "US",
		"timezone":         "Asia/Kolkata",
		"language_code":    "en",
		"industry":         "15",
		"roletitle":        "4",
		"invite":           false,
		"customer_groups":  []interface{}{"37152000000043029"},
		"digest":           "s247string",
		"zuids":            []interface{}{"75086549"},
		"customer_company": "w3schools",
		"display_name":     "phillips",
		"customer_website": "https://www.w3schools.com",
		"email_address":    "selvalakshmi.m+aug18@zohotest.com",
		"portal_name":      "w3schools",
		"captcha":          "D6EF1P",
	})
}
