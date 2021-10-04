package rest

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResponseIntoReturnsPreviousError(t *testing.T) {
	r := Response{err: errors.New("whoops")}

	var v *string

	err := r.Parse(v)

	require.Error(t, err)
	require.Equal(t, "whoops", err.Error())
	assert.Nil(t, v)
}

func TestResponseIntoReturnsErrorOnInvalidResponseBody(t *testing.T) {
	r := Response{body: []byte("{")}

	var v map[string]string

	err := r.Parse(&v)

	require.Error(t, err)
	assert.Equal(t, "unexpected end of JSON input", err.Error())
	assert.Nil(t, v)
}

func TestResponseParse(t *testing.T) {
	r := Response{body: []byte(`{"code":0,"message":"success","data":{"foo":"bar"}}`)}

	var v map[string]string

	err := r.Parse(&v)

	require.NoError(t, err)
	assert.Equal(t, map[string]string{"foo": "bar"}, v)
}
