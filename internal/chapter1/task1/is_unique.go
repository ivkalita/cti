package task1

import (
	"sort"
)

// IsUniqueWithSet – O(1) memory, O(n) time
// uses additional constant (if alphabet size is constant) memory to store set with visited runes
func IsUniqueWithSet(a string) bool {
	set := make(map[rune]interface{}, 256)
	for _, r := range a {
		if _, ok := set[r]; ok {
			return false
		}
		set[r] = true
	}
	return true
}

// IsUniqueWithSort – O(1) memory, O(n log n) time
// sorts string in any order and then check that there are no consecutive equal runes
func IsUniqueWithSort(a string) bool {
	runes := []rune(a)
	if len(runes) == 0 {
		return true
	}

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	prev := runes[0]
	for i, r := range runes {
		if i == 0 {
			continue
		}

		if r == prev {
			return false
		}

		prev = r
	}

	return true
}
