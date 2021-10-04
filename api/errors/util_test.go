package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNotFound(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "plain error",
			err:      errors.New("foo"),
			expected: false,
		},
		{
			name:     "404 status error",
			err:      NewStatusError(404, "not found"),
			expected: true,
		},
		{
			name:     "5xx status error",
			err:      NewStatusError(503, "service unavailable"),
			expected: false,
		},
		{
			name:     "404 extended status error",
			err:      NewExtendedStatusError(404, "not found", 0, nil),
			expected: true,
		},
		{
			name:     "5xx extended status error",
			err:      NewExtendedStatusError(503, "service unavailable", 0, nil),
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, IsNotFound(test.err))
		})
	}
}
