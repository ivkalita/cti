package task6_test

import (
	"github.com/ivkalita/cti/internal/chapter2/structs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func RunTests(t *testing.T, testable func(a structs.LinkedList) bool) {
	tests := []struct {
		name     string
		in       []int
		expected bool
	}{
		{
			name:     "101 -> true",
			in:       []int{1, 0, 1},
			expected: true,
		},
		{
			name:     "102 -> false",
			in:       []int{1, 0, 2},
			expected: false,
		},
		{
			name:     "0 -> true",
			in:       []int{0},
			expected: true,
		},
		{
			name:     "empty -> true",
			in:       []int{},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := structs.NewLinkedList(test.in)

			assert.Equal(t, test.expected, testable(l))
		})
	}
}
