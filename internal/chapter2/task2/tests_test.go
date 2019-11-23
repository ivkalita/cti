package task2_test

import (
	"github.com/ivkalita/cti/internal/chapter2/structs"
	"github.com/stretchr/testify/require"
	"testing"
)

func RunTests(t *testing.T, testable func(l structs.LinkedList, k int) (int, error)) {
	tests := []struct {
		name        string
		in          []int
		k           int
		expected    int
		expectedErr bool
	}{
		{
			name:        "return the last element",
			in:          []int{1, 2, 3, 4, 5},
			k:           0,
			expected:    5,
			expectedErr: false,
		},
		{
			name:        "return the element before the last",
			in:          []int{1, 2, 3, 4, 5},
			k:           1,
			expected:    4,
			expectedErr: false,
		},
		{
			name:        "out of bounds",
			in:          []int{1, 2, 3},
			k:           -1,
			expected:    0,
			expectedErr: true,
		},
		{
			name:        "out of bounds x2",
			in:          []int{1, 2, 3},
			k:           4,
			expected:    0,
			expectedErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			elem, err := testable(structs.NewLinkedList(test.in), test.k)
			if test.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, test.expected, elem)
			}
		})
	}
}
