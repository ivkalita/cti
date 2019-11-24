package task4

import "github.com/ivkalita/cti/internal/chapter2/structs"

// Split - O(1) memory, O(n) time
func Split(l *structs.LinkedList, x int) {
	var (
		lesserStart  *structs.Node = nil
		lesser       *structs.Node = nil
		greaterStart *structs.Node = nil
		greater      *structs.Node = nil
	)

	for cur := l.Head; cur != nil; {
		if cur.Data < x {
			if lesser == nil {
				lesser = &structs.Node{Data: cur.Data, Next: nil}
				lesserStart = lesser
				cur = cur.Next
			} else {
				lesser.Next = cur
				lesser = cur
				cur = cur.Next
				lesser.Next = nil
			}
		} else {
			if greater == nil {
				greater = &structs.Node{Data: cur.Data, Next: nil}
				greaterStart = greater
				cur = cur.Next
			} else {
				greater.Next = cur
				greater = cur
				cur = cur.Next
				greater.Next = nil
			}
		}
	}

	if lesserStart == nil {
		l.Head = greater
	} else if greaterStart != nil {
		if lesser == nil {
			panic("can't happen")
		}
		lesser.Next = greaterStart
		l.Head = lesserStart
	} else {
		l.Head = lesserStart
	}
}
