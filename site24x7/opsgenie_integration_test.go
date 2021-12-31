package site24x7

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/site24x7/terraform-provider-site24x7/api"
	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/site24x7/terraform-provider-site24x7/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpsgenieIntegrationCreate(t *testing.T) {
	d := opsgenieIntegrationTestResourceData(t)

	c := fake.NewClient()

	a := &api.OpsgenieIntegration{
		Name:          "foo",
		URL:           "www.test.tld",
		SelectionType: 0,
		ManualResolve: true,
		TroubleAlert:  false,
		Monitors:      []string{"234", "567"},
		AlertTagIDs:   []string{"123", "456"},
	}

	c.FakeOpsgenieIntegration.On("Create", a).Return(a, nil).Once()

	require.NoError(t, opsgenieIntegrationCreate(d, c))

	c.FakeOpsgenieIntegration.On("Create", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := opsgenieIntegrationCreate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestOpsgenieIntegrationUpdate(t *testing.T) {
	d := opsgenieIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	a := &api.OpsgenieIntegration{
		ServiceID:     "123",
		ServiceStatus: 0,
		Name:          "foo",
		URL:           "www.test.tld",
		SelectionType: 0,
		ManualResolve: true,
		TroubleAlert:  false,
		Monitors:      []string{"234", "567"},
		AlertTagIDs:   []string{"123", "456"},
	}

	c.FakeOpsgenieIntegration.On("Update", a).Return(a, nil).Once()

	require.NoError(t, opsgenieIntegrationUpdate(d, c))

	c.FakeOpsgenieIntegration.On("Update", a).Return(a, apierrors.NewStatusError(500, "error")).Once()

	err := opsgenieIntegrationUpdate(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestOpsgenieIntegrationRead(t *testing.T) {
	d := opsgenieIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeOpsgenieIntegration.On("Get", "123").Return(&api.OpsgenieIntegration{}, nil).Once()

	require.NoError(t, opsgenieIntegrationRead(d, c))

	c.FakeOpsgenieIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	err := opsgenieIntegrationRead(d, c)

	assert.Equal(t, apierrors.NewStatusError(500, "error"), err)
}

func TestOpsgenieIntegrationDelete(t *testing.T) {
	d := opsgenieIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeThirdPartyIntegrations.On("Delete", "123").Return(nil).Once()

	require.NoError(t, opsgenieIntegrationDelete(d, c))

	c.FakeThirdPartyIntegrations.On("Delete", "123").Return(apierrors.NewStatusError(404, "not found")).Once()

	require.NoError(t, opsgenieIntegrationDelete(d, c))
}

func TestOpsgenieIntegrationExists(t *testing.T) {
	d := opsgenieIntegrationTestResourceData(t)
	d.SetId("123")

	c := fake.NewClient()

	c.FakeOpsgenieIntegration.On("Get", "123").Return(&api.OpsgenieIntegration{}, nil).Once()

	exists, err := opsgenieIntegrationExists(d, c)

	require.NoError(t, err)
	assert.True(t, exists)

	c.FakeOpsgenieIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(404, "not found")).Once()

	exists, err = opsgenieIntegrationExists(d, c)

	require.NoError(t, err)
	assert.False(t, exists)

	c.FakeOpsgenieIntegration.On("Get", "123").Return(nil, apierrors.NewStatusError(500, "error")).Once()

	exists, err = opsgenieIntegrationExists(d, c)

	require.Equal(t, apierrors.NewStatusError(500, "error"), err)
	assert.False(t, exists)
}

func opsgenieIntegrationTestResourceData(t *testing.T) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, OpsgenieIntegrationSchema, map[string]interface{}{
		"name":           "foo",
		"url":            "www.test.tld",
		"selection_type": 0,
		"trouble_alert":  false,
		"manual_resolve": true,
		"monitors": []interface{}{
			"234",
			"567",
		},
		"alert_tags_id": []interface{}{
			"123",
			"456",
		},
	})
}
