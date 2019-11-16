package task1_test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func RunTests(t *testing.T, testable func(a string) bool) {
	tests := []struct {
		name     string
		in       string
		expected bool
	}{
		{
			name:     "empty is unique",
			in:       "",
			expected: true,
		},
		{
			name:     "not unique",
			in:       "abcc",
			expected: false,
		},
		{
			name:     "unique",
			in:       "ace",
			expected: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expected, testable(test.in))
		})
	}
}
