package task7_test

import (
	"github.com/ivkalita/cti/internal/chapter2/structs"
	"github.com/stretchr/testify/require"
	"testing"
)

func RunTests(t *testing.T, testable func(a structs.LinkedList, b structs.LinkedList) *structs.Node) {
	tests := []struct {
		name     string
		generate func() (structs.LinkedList, structs.LinkedList, *structs.Node)
	}{
		{
			name: "no intersection",
			generate: func() (a structs.LinkedList, b structs.LinkedList, expected *structs.Node) {
				a = structs.LinkedList{}
				a.Head = &structs.Node{Data: 0}

				b = structs.LinkedList{}
				b.Head = &structs.Node{Data: 0}

				return a, b, nil
			},
		},
		{
			name: "last element is an intersection",
			generate: func() (a structs.LinkedList, b structs.LinkedList, expected *structs.Node) {
				a = structs.LinkedList{}
				a.Head = &structs.Node{Data: 0}

				b = structs.LinkedList{}
				b.Head = &structs.Node{Data: 0}

				expected = &structs.Node{Data: 1}

				a.Head.Next = expected
				b.Head.Next = expected

				return a, b, expected
			},
		},
		{
			name: "intersection in the middle",
			generate: func() (a structs.LinkedList, b structs.LinkedList, expected *structs.Node) {
				a = structs.LinkedList{}
				a.Head = &structs.Node{Data: 0}

				b = structs.LinkedList{}
				b.Head = &structs.Node{Data: 0}

				expected = &structs.Node{Data: 1}
				expected.Next = &structs.Node{Data: 2}

				a.Head.Next = expected
				b.Head.Next = expected

				return a, b, expected
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a, b, expected := test.generate()

			require.Equal(t, expected, testable(a, b))
		})
	}
}
