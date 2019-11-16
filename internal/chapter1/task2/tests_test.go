package task2_test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func RunTests(t *testing.T, testable func(a string, b string) bool) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected bool
	}{
		{
			name:     "two empty strings equal",
			a:        "",
			b:        "",
			expected: true,
		},
		{
			name:     "first empty not equal",
			a:        "a",
			b:        "",
			expected: false,
		},
		{
			name:     "second empty not equal",
			a:        "",
			b:        "b",
			expected: false,
		},
		{
			name:     "same strings",
			a:        "abcd",
			b:        "abcd",
			expected: true,
		},
		{
			name:     "real permutation",
			a:        "abcd",
			b:        "dcba",
			expected: true,
		},
		{
			name:     "permutation with duplicates",
			a:        "abcda",
			b:        "bcaad",
			expected: true,
		},
		{
			name:     "is not permutation",
			a:        "abcd",
			b:        "cbae",
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expected, testable(test.a, test.b))
		})
	}
}
