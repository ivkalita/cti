package task3_test

import (
	"github.com/ivkalita/cti/internal/chapter2/structs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func RunTests(t *testing.T, testable func(n *structs.Node)) {
	tests := []struct {
		name     string
		in       []int
		n        int
		expected []int
	}{
		{
			name:     "success",
			in:       []int{1, 2, 3, 4},
			n:        1,
			expected: []int{1, 3, 4},
		},
		{
			name:     "with duplicates",
			in:       []int{1, 2, 2, 2, 3},
			n:        2,
			expected: []int{1, 2, 2, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := structs.NewLinkedList(test.in)
			cur := l.Head
			for i := 0; i < test.n; i++ {
				cur = cur.Next
			}

			testable(cur)

			assert.Equal(t, test.expected, l.ToSlice())
		})
	}
}
