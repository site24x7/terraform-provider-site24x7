package site24x7

import (
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/fake"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestActionCreate(t *testing.T) {
	d := urlActionTestResourceData(t)

	c := fake.NewClient()

	a := &api.URLAutomation{
		ActionName:             "foobar",
		ActionMethod:           "G",
		AuthMethod:             "B",
		CustomParameters:       "foobarbaz",
		SendCustomParameters:   true,
		SendInJsonFormat:       true,
		SendIncidentParameters: false,
		ActionTimeout:          30,
		ActionUrl:              "https://example.com",
		ActionType:             1,
	}

	c.FakeURLAutomations.On("Create", a).Return(a, nil).Once()

	require.NoError(t, urlActionCreate(d, c))

	c.FakeURLAutomations.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := urlActionCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestActionUpdate(t *testing.T) {
	d := urlActionTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.URLAutomation{
		ActionID:               "123",
		ActionName:             "foobar",
		ActionMethod:           "G",
		AuthMethod:             "B",
		CustomParameters:       "foobarbaz",
		SendCustomParameters:   true,
		SendInJsonFormat:       true,
		SendIncidentParameters: false,
		ActionTimeout:          30,
		ActionUrl:              "https://example.com",
		ActionType:             1,
	}

	c.FakeURLAutomations.On("Update", a).Return(a, nil).Once()

	require.NoError(t, urlActionUpdate(d, c))

	c.FakeURLAutomations.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := urlActionUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestActionRead(t *testing.T) {
	d := urlActionTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeURLAutomations.On("Get", "123").Return(&api.URLAutomation{}, nil).Once()

	require.NoError(t, urlActionRead(d, c))

	c.FakeURLAutomations.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := urlActionRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestActionDelete(t *testing.T) {
	d := urlActionTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeURLAutomations.On("Delete", "123").Return(nil).Once()

	require.NoError(t, urlActionDelete(d, c))

	c.FakeURLAutomations.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, urlActionDelete(d, c))
}

func TestActionExists(t *testing.T) {
	d := urlActionTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeURLAutomations.On("Get", "123").Return(&api.URLAutomation{}, nil).Once()

	exists, err := urlActionExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeURLAutomations.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = urlActionExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeURLAutomations.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = urlActionExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func urlActionTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, URLActionSchema, map[string]interface{}{
		"name":                     "foobar",
		"method":                   "G",
		"auth_method":              "B",
		"custom_parameters":        "foobarbaz",
		"send_custom_parameters":   true,
		"send_in_json_format":      true,
		"send_incident_parameters": false,
		"timeout":                  30,
		"url":                      "https://example.com",
		"type":                     1,
	})
}
