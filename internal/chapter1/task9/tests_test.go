package task9_test

import (
	"github.com/stretchr/testify/require"
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
			name:     "is rotation",
			a:        "waterbottle",
			b:        "erbottlewat",
			expected: true,
		},
		{
			name:     "empty string",
			a:        "asdf",
			b:        "",
			expected: false,
		},
		{
			name:     "is not a rotation",
			a:        "aab",
			b:        "aac",
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expected, testable(test.a, test.b))
			require.Equal(t, test.expected, testable(test.b, test.a))
		})
	}
}
