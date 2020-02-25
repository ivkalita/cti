package task5_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func RunTests(t *testing.T, testable func(string, string) bool) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected bool
	}{
		{
			name:     "two empty strings are one away",
			a:        "",
			b:        "",
			expected: true,
		},
		{
			name:     "empty string and 1-char string are one away",
			a:        "",
			b:        "a",
			expected: true,
		},
		{
			name:     "one insert away",
			a:        "ab",
			b:        "abc",
			expected: true,
		},
		{
			name:     "one change away",
			a:        "abc",
			b:        "abd",
			expected: true,
		},
		{
			name:     "multiple inserts",
			a:        "a",
			b:        "abc",
			expected: false,
		},
		{
			name:     "multiple changes",
			a:        "abc",
			b:        "acd",
			expected: false,
		},
		{
			name:     "insert at the first position",
			a:        "abc",
			b:        "bc",
			expected: true,
		},
		{
			name:     "example1",
			a:        "pale",
			b:        "ple",
			expected: true,
		},
		{
			name:     "example2",
			a:        "pales",
			b:        "ples",
			expected: true,
		},
		{
			name:     "example3",
			a:        "pale",
			b:        "bale",
			expected: true,
		},
		{
			name:     "example4",
			a:        "pale",
			b:        "bake",
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, testable(test.a, test.b))
			assert.Equal(t, test.expected, testable(test.b, test.a))
		})
	}
}
