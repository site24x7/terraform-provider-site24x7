package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/stretchr/testify/assert"
)

type endpointTest struct {
	name         string
	expectedVerb string
	expectedPath string
	expectedBody []byte
	statusCode   int
	responseBody []byte
	fn           func(*testing.T, rest.Client)
}

func runTest(t *testing.T, test *endpointTest) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if test.expectedVerb != "" {
			assert.Equal(t, test.expectedVerb, r.Method)
		}

		if test.expectedPath != "" {
			assert.Equal(t, test.expectedPath, r.RequestURI)
		}

		if test.expectedBody != nil {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			defer r.Body.Close()

			assert.Equal(t, string(indentJSON(t, test.expectedBody)), string(indentJSON(t, body)))
		}

		w.WriteHeader(test.statusCode)
		_, err := w.Write(test.responseBody)
		if err != nil {
			t.Fatal(err)
		}
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	restClient := rest.NewClient(http.DefaultClient, server.URL)

	test.fn(t, restClient)
}

func runTests(t *testing.T, tests []*endpointTest) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			runTest(t, test)
		})
	}
}

func jsonBody(t *testing.T, v interface{}) []byte {
	buf, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}

	return buf
}

func jsonAPIResponseBody(t *testing.T, v interface{}) []byte {
	rawData, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}

	resp := &api.Response{
		Data: rawData,
	}

	return jsonBody(t, resp)
}

func fixture(t *testing.T, filename string) []byte {
	buf, err := ioutil.ReadFile("testdata/fixtures/" + filename)
	if err != nil {
		t.Fatal(err)
	}

	return buf
}

func indentJSON(t *testing.T, buf []byte) []byte {
	var v interface{}

	err := json.Unmarshal(buf, &v)
	if err != nil {
		t.Fatal(err)
	}

	buf2, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	return buf2
}
