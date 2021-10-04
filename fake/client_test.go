package fake

import (
	"errors"
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestClientMonitorsCreate(t *testing.T) {
	c := NewClient()

	monitor := &api.WebsiteMonitor{
		DisplayName: "foo",
	}

	c.FakeWebsiteMonitors.On("Create", mock.Anything).Return(monitor, nil).Once()
	c.FakeWebsiteMonitors.On("Create", mock.Anything).Return(nil, errors.New("whoops")).Once()

	result, err := c.WebsiteMonitors().Create(&api.WebsiteMonitor{})

	require.NoError(t, err)

	assert.Equal(t, monitor, result)

	_, err = c.WebsiteMonitors().Create(&api.WebsiteMonitor{})

	require.Error(t, err)

	assert.Equal(t, errors.New("whoops"), err)

	c.FakeWebsiteMonitors.AssertExpectations(t)
}
