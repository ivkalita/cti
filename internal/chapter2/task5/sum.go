package task5

import (
	"errors"
	"github.com/ivkalita/cti/internal/chapter2/structs"
)

// SumReverse – O(max(n, m)) memory, O(max(n, m)) time, sums lists in reversed order
func SumReverse(a structs.LinkedList, b structs.LinkedList) (structs.LinkedList, error) {
	c := structs.NewLinkedList([]int{})

	if a.Head == nil || b.Head == nil {
		return c, errors.New("one of the arguments is empty list")
	}

	plusOne := 0
	curA := a.Head
	curB := b.Head
	var curC *structs.Node = nil
	for {
		if curA == nil && curB == nil && plusOne == 0 {
			break
		}

		newVal := plusOne

		if curA != nil {
			if curA.Data < 0 || curA.Data > 9 {
				return c, errors.New("not a valid digit")
			}
			newVal += curA.Data
			curA = curA.Next
		}

		if curB != nil {
			if curB.Data < 0 || curB.Data > 9 {
				return c, errors.New("not a valid digit")
			}
			newVal += curB.Data
			curB = curB.Next
		}

		plusOne = newVal / 10
		newVal = newVal % 10
		n := structs.Node{Data: newVal}

		if curC == nil {
			c.Head = &n
			curC = c.Head
		} else {
			curC.Next = &n
			curC = curC.Next
		}
	}

	return c, nil
}

// Sum – O(max(n, m)) memory, O(max(n, m)) time, sums lists in regular order
func Sum(a structs.LinkedList, b structs.LinkedList) (structs.LinkedList, error) {
	reverse(&a)
	reverse(&b)

	c, err := SumReverse(a, b)
	if err != nil {
		return c, err
	}

	reverse(&c)

	return c, nil
}

// reverse – helper func to reverse a list in place
func reverse(a *structs.LinkedList) {
	if a == nil || a.Head == nil || a.Head.Next == nil {
		return
	}

	cur := a.Head.Next
	prev := a.Head
	prev.Next = nil

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	a.Head = prev
}
