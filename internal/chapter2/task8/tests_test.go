package task8_test

import (
	"github.com/ivkalita/cti/internal/chapter2/structs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func RunTests(t *testing.T, testable func(structs.LinkedList) *structs.Node) {
	tests := []struct {
		name     string
		generate func() (structs.LinkedList, *structs.Node)
	}{
		{
			name: "regular loop",
			generate: func() (structs.LinkedList, *structs.Node) {
				a := structs.LinkedList{}
				n1 := &structs.Node{Data: 1}
				n2 := &structs.Node{Data: 2}
				n3 := &structs.Node{Data: 3}
				n4 := &structs.Node{Data: 4}
				n5 := &structs.Node{Data: 5}

				a.Head = n1
				n1.Next = n2
				n2.Next = n3
				n3.Next = n4
				n4.Next = n5
				n5.Next = n3

				return a, n3
			},
		},
		{
			name: "loop at the beginning",
			generate: func() (structs.LinkedList, *structs.Node) {
				a := structs.LinkedList{}
				n1 := &structs.Node{Data: 1}
				n2 := &structs.Node{Data: 2}
				n3 := &structs.Node{Data: 3}

				a.Head = n1
				n1.Next = n2
				n2.Next = n3
				n3.Next = n1

				return a, n1
			},
		},
		{
			name: "no loop",
			generate: func() (structs.LinkedList, *structs.Node) {
				a := structs.LinkedList{}
				n1 := &structs.Node{Data: 1}
				n2 := &structs.Node{Data: 2}
				n3 := &structs.Node{Data: 3}

				a.Head = n1
				n1.Next = n2
				n2.Next = n3

				return a, nil
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list, expected := test.generate()

			assert.Equal(t, expected, testable(list))
		})
	}
}
