package task1_test

import (
	"github.com/ivkalita/cti/internal/chapter2/structs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func RunTests(t *testing.T, testable func(l structs.LinkedList)) {
	tests := []struct {
		name     string
		in       []int
		expected []int
	}{
		{
			name:     "zero nodes",
			in:       []int{},
			expected: []int{},
		},
		{
			name:     "one node",
			in:       []int{1},
			expected: []int{1},
		},
		{
			name:     "list without duplicates",
			in:       []int{3, 1, 2, 10, 4},
			expected: []int{3, 1, 2, 10, 4},
		},
		{
			name:     "list with duplicates",
			in:       []int{3, 1, 3, 10, 5, 15, 5},
			expected: []int{3, 1, 10, 5, 15},
		},
		{
			name:     "list with only duplicates",
			in:       []int{1, 1, 1, 1, 1, 1, 1},
			expected: []int{1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := structs.NewLinkedList(test.in)

			testable(l)

			assert.Equal(t, test.expected, l.ToSlice())
		})
	}
}
