package integration

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConnectwiseIntegrationCreate(t *testing.T) {
	d := connectwiseIntegrationTestResourceData(t)

	c := fake.NewClient()

	a := &api.ConnectwiseIntegration{
		Name:          "foo",
		URL:           "https://staging.connectwisedev.com/",
		Company: 	   "zylker_c",
		PublicKey: 	   "KaxKPKiP88i6rmAb",
    	PrivateKey:    "Fkb7dlqwhQGIxcc5",
    	CompanyId: 	   "GreenInc",
		CloseStatus:   "Closed (resolved)",
		SelectionType: 0,
		TroubleAlert:  true,
		CriticalAlert: true,
		DownAlert:     true,
		AlertTagIDs:   []string{"123", "456"},
	}

	c.FakeConnectwiseIntegration.On("Create", a).Return(a, nil).Once()

	require.NoError(t, connectwiseIntegrationCreate(d, c))

	c.FakeConnectwiseIntegration.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := connectwiseIntegrationCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestConnectwiseDutyIntegrationUpdate(t *testing.T) {
	d := connectwiseIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.ConnectwiseIntegration{
		ServiceID:     "123",
		Name:          "foo",
		URL:           "https://staging.connectwisedev.com/",
		Company: 	   "zylker_c",
		PublicKey: 	   "KaxKPKiP88i6rmAb",
    	PrivateKey:    "Fkb7dlqwhQGIxcc5",
    	CompanyId: 	   "GreenInc",
		CloseStatus:   "Closed (resolved)",
		SelectionType: 0,
		TroubleAlert:  true,
		CriticalAlert: true,
		DownAlert:     true,
		AlertTagIDs:   []string{"123", "456"},
	}

	c.FakeConnectwiseIntegration.On("Update", a).Return(a, nil).Once()

	require.NoError(t, connectwiseIntegrationUpdate(d, c))

	c.FakeConnectwiseIntegration.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := connectwiseIntegrationUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestConnectwiseIntegrationRead(t *testing.T) {
	d := connectwiseIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeConnectwiseIntegration.On("Get", "123").Return(&api.ConnectwiseIntegration{}, nil).Once()

	require.NoError(t, connectwiseIntegrationRead(d, c))

	c.FakeConnectwiseIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := connectwiseIntegrationRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestConnectwiseIntegrationDelete(t *testing.T) {
	d := connectwiseIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeThirdPartyIntegrations.On("Delete", "123").Return(nil).Once()

	require.NoError(t, connectwiseIntegrationDelete(d, c))

	c.FakeThirdPartyIntegrations.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, connectwiseIntegrationDelete(d, c))
}

func TestConnectwiseIntegrationExists(t *testing.T) {
	d := connectwiseIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeConnectwiseIntegration.On("Get", "123").Return(&api.ConnectwiseIntegration{}, nil).Once()

	exists, err := connectwiseIntegrationExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeConnectwiseIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = connectwiseIntegrationExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeConnectwiseIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = connectwiseIntegrationExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func connectwiseIntegrationTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, ConnectwiseIntegrationSchema, map[string]interface{}{
		"name":           "foo",
		"url":            "https://staging.connectwisedev.com/",
		"selection_type": 0,
		"company": 		  "zylker_c",
		"public_key": 	  "KaxKPKiP88i6rmAb",
    	"private_key": 	  "Fkb7dlqwhQGIxcc5",
    	"company_id": 	  "GreenInc",
		"close_status":   "Closed (resolved)",
		"trouble_alert":  true,
		"critical_alert": true,
		"down_alert":     true,
		"alert_tags_id": []interface{}{
			"123",
			"456",
		},
	})
}