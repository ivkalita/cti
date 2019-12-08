package task8

import "github.com/ivkalita/cti/internal/chapter2/structs"

// FindLoop – O(1) memory, O(n) time
func FindLoop(l structs.LinkedList) *structs.Node {
	s := l.Head
	f := l.Head

	// looking for the collision point
	for {
		if s == nil || f == nil || f.Next == nil {
			return nil
		}
		s = s.Next
		f = f.Next.Next

		if s == f {
			break
		}
	}

	// now slow and fast both point to collide point
	s = l.Head
	for s != f {
		if s == nil || f == nil {
			return nil
		}
		s = s.Next
		f = f.Next
	}

	return s
}
