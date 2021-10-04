package backoff

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultRetryPolicy(t *testing.T) {
	tests := []struct {
		name        string
		ctx         context.Context
		resp        *http.Response
		err         error
		expected    bool
		expectedErr string
	}{
		{
			name: "ctx error should not be retried",
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			}(),
			expected:    false,
			expectedErr: "context canceled",
		},
		{
			name:        "errors should be retried",
			err:         errors.New("something went wrong"),
			expected:    true,
			expectedErr: "something went wrong",
		},
		{
			name:     "empty status code should be retried",
			resp:     newResponse(0),
			expected: true,
		},
		{
			name:     "429 status code should be retried",
			resp:     newResponse(429),
			expected: true,
		},
		{
			name:     "5xx status code should be retried",
			resp:     newResponse(503),
			expected: true,
		},
		{
			name:     "501 status code should not be retried",
			resp:     newResponse(501),
			expected: false,
		},
		{
			name:     "4xx status code should not be retried",
			resp:     newResponse(400),
			expected: false,
		},
		{
			name:     "3xx status code should not be retried",
			resp:     newResponse(302),
			expected: false,
		},
		{
			name:     "2xx status code should not be retried",
			resp:     newResponse(200),
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := test.ctx
			if ctx == nil {
				ctx = context.Background()
			}

			retry, err := DefaultRetryPolicy(ctx, test.resp, test.err)
			if test.expectedErr != "" {
				require.Error(t, err)
				assert.Equal(t, test.expectedErr, err.Error())
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, test.expected, retry)
		})
	}
}

func TestDefaultBackoff(t *testing.T) {
	tests := []struct {
		name       string
		min, max   time.Duration
		attemptNum int
		resp       *http.Response
		expected   time.Duration
	}{
		{
			name:       "increasing exponential backoff by attempt",
			min:        1 * time.Second,
			max:        30 * time.Second,
			attemptNum: 3,
			expected:   8 * time.Second,
		},
		{
			name:       "min exponential backoff on first attempt",
			min:        1 * time.Second,
			max:        30 * time.Second,
			attemptNum: 0,
			expected:   1 * time.Second,
		},
		{
			name:       "exponential backoff capped by max",
			min:        1 * time.Second,
			max:        30 * time.Second,
			attemptNum: 10,
			expected:   30 * time.Second,
		},
		{
			name:     "use Retry-After header value",
			min:      1 * time.Second,
			max:      30 * time.Second,
			expected: 3 * time.Second,
			resp:     withRetryAfter(newResponse(429), "3"),
		},
		{
			name:     "Retry-After capped to max",
			min:      1 * time.Second,
			max:      30 * time.Second,
			expected: 30 * time.Second,
			resp:     withRetryAfter(newResponse(429), "3600"),
		},
		{
			name:       "Invalid Retry-After falls back to exponential backoff",
			min:        1 * time.Second,
			max:        30 * time.Second,
			attemptNum: 3,
			expected:   8 * time.Second,
			resp:       withRetryAfter(newResponse(429), "abc"),
		},
		{
			name:       "Empty Retry-After falls back to exponential backoff",
			min:        1 * time.Second,
			max:        30 * time.Second,
			attemptNum: 2,
			expected:   4 * time.Second,
			resp:       withRetryAfter(newResponse(429), ""),
		},
		{
			name:       "Zero Retry-After falls back to exponential backoff",
			min:        1 * time.Second,
			max:        30 * time.Second,
			attemptNum: 1,
			expected:   2 * time.Second,
			resp:       withRetryAfter(newResponse(429), "0"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			backoff := DefaultBackoff(test.min, test.max, test.attemptNum, test.resp)

			assert.Equal(t, test.expected, backoff)
		})
	}
}

func newResponse(statusCode int) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
	}
}

func withRetryAfter(resp *http.Response, val string) *http.Response {
	if resp.Header == nil {
		resp.Header = http.Header{}
	}
	resp.Header.Add("Retry-After", val)

	return resp
}
