package task3_test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func RunTests(t *testing.T, testable func(*[]rune, int) error) {
	tests := []struct {
		name        string
		in          []rune
		realLen     int
		expectedErr bool
		expected    string
	}{
		{
			name:        "empty string not touched",
			in:          []rune(""),
			realLen:     0,
			expectedErr: false,
			expected:    "",
		},
		{
			name:        "string with only spaces not touched",
			in:          []rune("   "),
			realLen:     0,
			expectedErr: false,
			expected:    "   ",
		},
		{
			name:        "string with spaces urlified",
			in:          []rune("string with spaces    "),
			realLen:     18,
			expectedErr: false,
			expected:    "string%20with%20spaces",
		},
		{
			name:        "string with a space at the end",
			in:          []rune("string with spaces       "),
			realLen:     19,
			expectedErr: false,
			expected:    "string%20with%20spaces%20",
		},
		{
			name:        "string with a space at the end outside of real len",
			in:          []rune("string with spaces        "),
			realLen:     19,
			expectedErr: false,
			expected:    "string%20with%20spaces%20 ",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testable(&test.in, test.realLen)
			if test.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, []rune(test.expected), test.in)
			}
		})
	}
}
