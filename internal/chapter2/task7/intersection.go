package task7

import "github.com/ivkalita/cti/internal/chapter2/structs"

// FindIntersection – O(1) memory O(n + m) time
func FindIntersection(a structs.LinkedList, b structs.LinkedList) *structs.Node {
	if a.Head == nil || b.Head == nil {
		return nil
	}

	aLength := 0
	for curA := a.Head; curA != nil; curA = curA.Next {
		aLength++
	}

	bLength := 0
	for curB := b.Head; curB != nil; curB = curB.Next {
		bLength++
	}

	curA := a.Head
	curB := b.Head

	if bLength > aLength {
		for i := 0; i < bLength-aLength; i, curB = i+1, curB.Next {
			curB = curB.Next
		}
	}

	if aLength > bLength {
		for i := 0; i < aLength-bLength; i, curA = i+1, curA.Next {
		}
	}

	for ; curA != nil && curB != nil; curA, curB = curA.Next, curB.Next {
		if curA == curB {
			return curA
		}
	}

	return nil
}
