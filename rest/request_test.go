package rest

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	apierrors "github.com/site24x7/terraform-provider-site24x7/api/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequestSetsHeader(t *testing.T) {
	r := NewRequest(nil, "GET", "")

	assert.Empty(t, r.header.Get("Content-Type"))

	r = NewRequest(nil, "GET", "").AddHeader("Content-Type", "application/json")

	assert.Equal(t, "application/json", r.header.Get("Content-Type"))
}

func TestRequestSetsJSONBody(t *testing.T) {
	r := NewRequest(nil, "GET", "").
		Body(map[string]string{"foo": "bar"})

	require.NoError(t, r.err)
	assert.Equal(t, []byte(`{"foo":"bar"}`), r.body)
}

func TestRequestSetsErrorOnInvalidBody(t *testing.T) {
	r := NewRequest(nil, "GET", "").
		Body(func() {})

	require.Error(t, r.err)
	assert.Equal(t, "json: unsupported type: func()", r.err.Error())
}

func TestRequestIsNotSentOnPreviousError(t *testing.T) {
	c := newFakeHTTPClient()

	r := NewRequest(c, "GET", "")
	r.err = errors.New("whoops")

	resp := r.Do()

	assert.Error(t, resp.Err())
	assert.Equal(t, 0, c.called)
}

func TestRequestBuildRequest(t *testing.T) {
	body := map[string]string{
		"foo": "bar",
	}

	r := NewRequest(nil, "DELETE", "").
		Resource("foos").
		ResourceID("123").
		Body(body)

	req, err := r.buildRequest()

	require.NoError(t, err)

	assert.Equal(t, "DELETE", req.Method)
	assert.Equal(t, "/foos/123", req.URL.Path)

	buf, err := ioutil.ReadAll(req.Body)

	require.NoError(t, err)

	assert.Equal(t, `{"foo":"bar"}`, string(buf))
}

func TestRequestDo(t *testing.T) {
	c := newFakeHTTPClient().
		WithStatusCode(200).
		WithResponseBody([]byte(`{"data":{"foo":"bar"}}`))

	resp := NewRequest(c, "POST", "").
		Resource("foos").
		ResourceID("123").
		Do()

	require.NoError(t, resp.Err())
	assert.Equal(t, []byte(`{"data":{"foo":"bar"}}`), resp.body)
}

func TestRequestDoConvertsHTTPErrorsToStatusError(t *testing.T) {
	c := newFakeHTTPClient().
		WithStatusCode(404).
		WithResponseBody([]byte(`{"error_code":456,"message":"not found","error_info":{"foo":"bar"}}`))

	err := NewRequest(c, "PUT", "").
		Resource("foos").
		ResourceID("123").
		Do().
		Err()

	expectedErr := apierrors.NewExtendedStatusError(404, "not found", 456, map[string]interface{}{"foo": "bar"})

	assert.True(t, apierrors.IsExtendedStatusError(err))
	assert.Equal(t, expectedErr, err)
}

func TestRequestDoFallsBackToStatusErrorIfErrorResponseBodyIsInvalid(t *testing.T) {
	c := newFakeHTTPClient().
		WithStatusCode(400).
		WithResponseBody([]byte(`{`))

	err := NewRequest(c, "PUT", "").
		Resource("foos").
		ResourceID("123").
		Body(nil).
		Do().
		Err()

	expectedErr := apierrors.NewStatusError(400, "server replied with: {")

	assert.True(t, apierrors.IsStatusError(err))
	assert.Equal(t, expectedErr, err)
}

type fakeHTTPClient struct {
	resp       *http.Response
	err        error
	called     int
	calledWith []*http.Request
}

func newFakeHTTPClient() *fakeHTTPClient {
	return &fakeHTTPClient{}
}

func (c *fakeHTTPClient) WithStatusCode(code int) *fakeHTTPClient {
	if c.resp == nil {
		c.resp = &http.Response{
			Body: ioutil.NopCloser(bytes.NewReader(nil)),
		}
	}

	c.resp.StatusCode = code
	return c
}

func (c *fakeHTTPClient) WithResponseBody(buf []byte) *fakeHTTPClient {
	if c.resp == nil {
		c.resp = &http.Response{}
	}

	c.resp.Body = ioutil.NopCloser(bytes.NewReader(buf))
	c.resp.ContentLength = int64(len(buf))
	return c
}

func (c *fakeHTTPClient) Do(req *http.Request) (*http.Response, error) {
	c.called++
	if c.calledWith == nil {
		c.calledWith = make([]*http.Request, 0, 1)
	}
	c.calledWith = append(c.calledWith, req)
	return c.resp, c.err
}
