package task5_test

import (
	"fmt"
	"github.com/ivkalita/cti/internal/chapter2/structs"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func RunTests(t *testing.T, reversed bool, testable func(a structs.LinkedList, b structs.LinkedList) (structs.LinkedList, error)) {
	tests := []struct {
		a int
		b int
	}{
		{
			a: 0,
			b: 50,
		},
		{
			a: 1,
			b: 100,
		},
		{
			a: 0,
			b: 0,
		},
		{
			a: 999,
			b: 1,
		},
		{
			a: 999,
			b: 9999,
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v %v + %v = %v", reversed, test.a, test.b, test.a+test.b)
		t.Run(name, func(t *testing.T) {
			la := intToList(test.a, reversed)
			lb := intToList(test.b, reversed)
			lc := intToList(test.a+test.b, reversed)

			actual, err := testable(la, lb)

			assert.NoError(t, err)
			assert.Equal(t, lc.ToSlice(), actual.ToSlice())
		})
	}
}

func intToList(val int, reversed bool) structs.LinkedList {
	s := strconv.Itoa(val)
	n := len(s)
	d := make([]int, n)
	for i, r := range s {
		idx := i
		if reversed {
			idx = n - idx - 1
		}
		d[idx] = int(r - '0')
	}

	return structs.NewLinkedList(d)
}
