package task1

import "github.com/ivkalita/cti/internal/chapter2/structs"

// RemoveDuplicatesWithMap – O(n) memory, O(n) time
func RemoveDuplicatesWithMap(l structs.LinkedList) {
	knownValues := make(map[int]interface{})
	cur := l.Head
	var prev *structs.Node = nil

	for cur != nil {
		_, ok := knownValues[cur.Data]
		if !ok {
			knownValues[cur.Data] = nil
			prev = cur
			cur = cur.Next
			continue
		}

		if prev == nil {
			panic("can't happen")
		}

		prev.Next = cur.Next
		cur = cur.Next
	}
}

// RemoveDuplicatesWithoutMap – O(1) memory, O(n^2) time
func RemoveDuplicatesWithoutMap(l structs.LinkedList) {
	i := l.Head
	for i != nil {
		prev := i
		j := i.Next
		for j != nil {
			if j.Data == i.Data {
				prev.Next = j.Next
			} else {
				prev = j
			}
			j = j.Next
		}
		i = i.Next
	}
}
