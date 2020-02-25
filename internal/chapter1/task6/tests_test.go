package task6_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func RunTests(t *testing.T, testable func(string) string) {
	tests := []struct {
		name     string
		in       string
		expected string
	}{
		{
			name:     "example",
			in:       "aabcccccaaa",
			expected: "a2b1c5a3",
		},
		{
			name:     "same length",
			in:       "abc",
			expected: "abc",
		},
		{
			name:     "empty string",
			in:       "",
			expected: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, testable(test.in))
		})
	}
}
