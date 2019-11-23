package task2

import (
	"errors"
	"github.com/ivkalita/cti/internal/chapter2/structs"
)

// KthToLast1 – O(1) memory, O(n) time
func KthToLast1(l structs.LinkedList, k int) (int, error) {
	if k < 0 {
		return 0, errors.New("k less than 0")
	}

	n := 0
	for cur := l.Head; cur != nil; cur = cur.Next {
		n++
	}

	if k > n {
		return 0, errors.New("k greater than total length")
	}

	num := n - k - 1

	cur := l.Head

	for i := 1; i <= num; i++ {
		cur = cur.Next
	}

	return cur.Data, nil
}

// KthToLast2 – O(1) memory, O(n) time
func KthToLast2(l structs.LinkedList, k int) (int, error) {
	p1 := l.Head
	p2 := l.Head

	for i := 0; i < k+1; i++ {
		if p1 == nil {
			return 0, errors.New("out of bounds")
		}
		p1 = p1.Next
	}

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	if p2 == nil {
		return 0, errors.New("out of bounds")
	}

	return p2.Data, nil
}
