package task5_test

import (
	"github.com/ivkalita/cti/internal/chapter3/common"
	"github.com/ivkalita/cti/internal/chapter3/task5"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestTwoStacksSort(t *testing.T) {
	tests := map[string][]int{
		"reversed":      {1, 2, 3},
		"empty":         {},
		"almost sorted": {2, 3, 1},
		"sorted":        {3, 2, 1},
		"same in line":  {3, 3, 5, 1},
		"sames":         {5, 1, 5, 2},
	}

	for name, inSlice := range tests {
		t.Run(name, func(t *testing.T) {
			input := task5.NewStack()
			for _, v := range inSlice {
				input.Push(v)
			}
			sort.Sort(sort.Reverse(sort.IntSlice(inSlice)))
			expected := task5.NewStack()
			for _, v := range inSlice {
				expected.Push(v)
			}

			actual := task5.TwoStacksSort(input)
			for {
				expectedEl, expectedErr := expected.Pop()
				actualEl, actualErr := actual.Pop()
				require.Equal(t, expectedErr, actualErr)
				if expectedErr == common.StackIsEmpty {
					break
				}
				require.Equal(t, expectedEl, actualEl)
			}
		})
	}
}
