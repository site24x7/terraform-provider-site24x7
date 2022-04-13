package validation

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"testing"

	"github.com/site24x7/terraform-provider-site24x7/api"
	"github.com/site24x7/terraform-provider-site24x7/rest"
	"github.com/stretchr/testify/assert"
)

type EndpointTest struct {
	Name         string
	ExpectedVerb string
	ExpectedPath string
	ExpectedBody []byte
	StatusCode   int
	ResponseBody []byte
	Fn           func(*testing.T, rest.Client)
}

func RunTest(t *testing.T, test *EndpointTest) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if test.ExpectedVerb != "" {
			assert.Equal(t, test.ExpectedVerb, r.Method)
		}

		if test.ExpectedPath != "" {
			assert.Equal(t, test.ExpectedPath, r.RequestURI)
		}

		if test.ExpectedBody != nil {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			defer r.Body.Close()

			assert.Equal(t, string(IndentJSON(t, test.ExpectedBody)), string(IndentJSON(t, body)))
		}

		w.WriteHeader(test.StatusCode)
		_, err := w.Write(test.ResponseBody)
		if err != nil {
			t.Fatal(err)
		}
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	clientConfig := rest.ClientConfig{
		APIBaseURL: server.URL,
		TokenURL:   server.URL,
	}

	restClient := rest.NewClient(http.DefaultClient, clientConfig)

	test.Fn(t, restClient)
}

func RunTests(t *testing.T, tests []*EndpointTest) {
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			RunTest(t, test)
		})
	}
}

func JsonBody(t *testing.T, v interface{}) []byte {
	buf, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}

	return buf
}

func JsonAPIResponseBody(t *testing.T, v interface{}) []byte {
	rawData, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}

	resp := &api.Response{
		Data: rawData,
	}

	return JsonBody(t, resp)
}

func Fixture(t *testing.T, filename string) []byte {
	_, dir, _, _ := runtime.Caller(0)
	path := strings.ReplaceAll(dir, "validation/validation.go", "api/endpoints/testdata/fixtures/")
	buf, err := ioutil.ReadFile(path + filename)
	if err != nil {
		t.Fatal(err)
	}

	return buf
}

func IndentJSON(t *testing.T, buf []byte) []byte {
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
