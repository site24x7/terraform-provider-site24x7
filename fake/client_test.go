package fake

import (
	"errors"
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/site24x7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// make sure whole interface is implemented
var _ site24x7.Client = &Client{}

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

func TestClientDNSServerMonitorsCreate(t *testing.T) {
	c := NewClient()

	monitor := &api.DNSServerMonitor{
		DisplayName: "foo",
	}

	c.FakeDNSServerMonitors.On("Create", mock.Anything).Return(monitor, nil).Once()
	c.FakeDNSServerMonitors.On("Create", mock.Anything).Return(nil, errors.New("whoops")).Once()

	result, err := c.DNSServerMonitors().Create(&api.DNSServerMonitor{})

	require.NoError(t, err)

	assert.Equal(t, monitor, result)

	_, err = c.DNSServerMonitors().Create(&api.DNSServerMonitor{})

	require.Error(t, err)

	assert.Equal(t, errors.New("whoops"), err)

	c.FakeDNSServerMonitors.AssertExpectations(t)
}
