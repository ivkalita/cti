package task4_test

import (
	"github.com/ivkalita/cti/internal/chapter2/structs"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func RunTests(t *testing.T, testable func(l *structs.LinkedList, x int)) {
	tests := []struct {
		name string
		in   []int
		x    int
	}{
		{
			name: "from example",
			in:   []int{3, 5, 8, 5, 10, 2, 1},
			x:    5,
		},
		{
			name: "less first",
			in:   []int{1, 2, 3, 5, 6, 7},
			x:    5,
		},
		{
			name: "one element only",
			in:   []int{1},
			x:    2,
		},
		{
			name: "one element only (x2)",
			in:   []int{1},
			x:    0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := structs.NewLinkedList(test.in)

			testable(&l, test.x)

			actual := l.ToSlice()
			lesserPart := true
			for _, v := range actual {
				if lesserPart {
					if v >= test.x {
						lesserPart = false
					}
					continue
				}

				if !lesserPart {
					if v < test.x {
						assert.Failf(t, "order is wrong", "value %v less than %v", v, test.x)
					}
				}
			}

			sort.Ints(actual)
			sort.Ints(test.in)
			assert.Equal(t, test.in, actual)
		})
	}
}
