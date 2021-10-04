package rest

import (
	"encoding/json"

	"github.com/site24x7/terraform-provider-site24x7/api"
)

// Response is a holder for the response body and errors that happened during a
// request.
type Response struct {
	err  error
	body []byte
}

// Into unmarshals the response body into v. The passed in value must be a
// pointer. It returns any error that occurred during the request. This is a
// no-op if there were errors before.
func (r Response) Parse(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	resp := &api.Response{}

	err := json.Unmarshal(r.body, resp)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(resp.Data), v)
}

// Err returns the request error if there was one.
func (r Response) Err() error {
	return r.err
}
