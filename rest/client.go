package rest

import "net/http"

// Client is the interface of a rest client that can build requests for the
// different http verbs.
type Client interface {
	Verb(verb string) *Request
	Post() *Request
	Get() *Request
	Put() *Request
	Delete() *Request
}

// HTTPClient is the interface of an http client that is compatible with
// *http.Client.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type client struct {
	baseURL    string
	httpClient HTTPClient
}

// New Client creates a new REST Client.
func NewClient(httpClient HTTPClient, baseURL string) Client {
	return &client{
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}

// Verb creates a new *Request with given HTTP verb, e.g. 'POST', 'PUT', 'GET'
// or 'DELETE'.
func (c *client) Verb(verb string) *Request {
	return NewRequest(c.httpClient, verb, c.baseURL)
}

// Get creates a new HTTP GET request.
func (c *client) Get() *Request {
	return c.Verb("GET")
}

// Post creates a new HTTP POST request.
func (c *client) Post() *Request {
	return c.Verb("POST")
}

// Put creates a new HTTP PUT request.
func (c *client) Put() *Request {
	return c.Verb("PUT")
}

// Delete creates a new HTTP DELETE request.
func (c *client) Delete() *Request {
	return c.Verb("DELETE")
}
