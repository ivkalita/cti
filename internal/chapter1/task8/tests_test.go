package task8_test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func RunTests(t *testing.T, testable func(*[][]int32)) {
	tests := []struct {
		name     string
		in       [][]int32
		expected [][]int32
	}{
		{
			name: "2x2",
			in: [][]int32{
				{1, 1},
				{1, 0},
			},
			expected: [][]int32{
				{1, 0},
				{0, 0},
			},
		},
		{
			name: "3x3",
			in: [][]int32{
				{1, 1, 1},
				{1, 0, 0},
				{1, 1, 1},
			},
			expected: [][]int32{
				{1, 0, 0},
				{0, 0, 0},
				{1, 0, 0},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testable(&test.in)
			require.Equal(t, test.expected, test.in)
		})
	}
}
