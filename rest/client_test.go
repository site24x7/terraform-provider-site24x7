package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientVerbs(t *testing.T) {

	clientConfig := ClientConfig{
		APIBaseURL: "http://localhost/api",
		TokenURL:   "http://localhost/api",
	}

	c := NewClient(&fakeHTTPClient{}, clientConfig)

	assert.Equal(t, "POST", c.Post().verb)
	assert.Equal(t, "GET", c.Get().verb)
	assert.Equal(t, "PUT", c.Put().verb)
	assert.Equal(t, "DELETE", c.Delete().verb)
	assert.Equal(t, "OPTIONS", c.Verb("OPTIONS").verb)
}
