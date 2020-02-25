package task4_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func RunTests(t *testing.T, testable func(string) bool) {
	tests := []struct {
		name     string
		in       string
		expected bool
	}{
		{
			name:     "empty string is a palindrome",
			in:       "",
			expected: true,
		},
		{
			name:     "not a palindrome",
			in:       "abc",
			expected: false,
		},
		{
			name:     "palindrome without spaces",
			in:       "aab",
			expected: true,
		},
		{
			name:     "palindrome with spaces",
			in:       "a ab",
			expected: true,
		},
		{
			name:     "palindrome with multiple spaces",
			in:       "a a b",
			expected: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, testable(test.in))
		})
	}
}
