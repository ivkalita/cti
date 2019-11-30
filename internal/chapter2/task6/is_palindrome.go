package task6

import "github.com/ivkalita/cti/internal/chapter2/structs"

// IsPalindrome – O(n) memory, O(n) time
func IsPalindrome(a structs.LinkedList) bool {
	b := reverse(a)

	curA, curB := a.Head, b.Head

	for curA != nil && curB != nil {
		if curA.Data != curB.Data {
			return false
		}

		curA, curB = curA.Next, curB.Next
	}

	if curA != curB {
		return false
	}

	return true
}

// reverse – helper func to reverse a list and return the result (not in place)
func reverse(a structs.LinkedList) structs.LinkedList {
	b := structs.LinkedList{}
	if a.Head == nil {
		return b
	}

	var curB *structs.Node
	curA := a.Head

	for curA != nil {
		n := structs.Node{Data: curA.Data}
		n.Next = curB
		curB = &n
		curA = curA.Next
	}

	b.Head = curB

	return b
}
