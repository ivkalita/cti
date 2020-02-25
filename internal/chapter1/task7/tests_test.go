package task7_test

import (
	"github.com/stretchr/testify/assert"
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
				{1, 2},
				{3, 4},
			},
			expected: [][]int32{
				{2, 4},
				{1, 3},
			},
		},
		{
			name: "3x3",
			in: [][]int32{
				{1, 2, 3},
				{0, 0, 0},
				{1, 2, 3},
			},
			expected: [][]int32{
				{3, 0, 3},
				{2, 0, 2},
				{1, 0, 1},
			},
		},
		{
			name: "4x4",
			in: [][]int32{
				{11, 12, 13, 44},
				{21, 22, 23, 24},
				{31, 32, 33, 34},
				{41, 42, 43, 44},
			},
			expected: [][]int32{
				{44, 24, 34, 44},
				{13, 23, 33, 43},
				{12, 22, 32, 42},
				{11, 21, 31, 41},
			},
		},
		{
			name: "5x5",
			in: [][]int32{
				{11, 12, 13, 14, 15},
				{21, 22, 23, 24, 25},
				{31, 32, 33, 34, 35},
				{41, 42, 43, 44, 45},
				{51, 52, 53, 54, 55},
			},
			expected: [][]int32{
				{15, 25, 35, 45, 55},
				{14, 24, 34, 44, 54},
				{13, 23, 33, 43, 53},
				{12, 22, 32, 42, 52},
				{11, 21, 31, 41, 51},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testable(&test.in)
			assert.Equal(t, test.expected, test.in)
		})
	}
}
