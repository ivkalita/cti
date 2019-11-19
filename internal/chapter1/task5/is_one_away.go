package task5

// IsOneAway – O(1) memory, O(n) time
// based on strings lengths delta uses one of the O(n) checks (inserts/edits)
func IsOneAway(inA string, inB string) bool {
	a := []rune(inA)
	b := []rune(inB)
	// let's assume that a always bigger than b,
	// if not – swap strings
	if len(b) > len(a) {
		a, b = b, a
	}

	delta := len(a) - len(b)

	if delta > 1 {
		return false
	} else if delta == 0 {
		return isOneEditAway(a, b)
	} else {
		return isOneInsertAway(a, b)
	}
}

func isOneEditAway(a []rune, b []rune) bool {
	// assumes that both strings has the same length
	if len(a) != len(b) {
		panic("strings lengths are not the same")
	}

	edits := 0
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}
		if edits > 0 {
			return false
		}
		edits++
	}

	return true
}

func isOneInsertAway(a []rune, b []rune) bool {
	// assumes that len(a) - len(b) == 1
	if len(a)-len(b) != 1 {
		panic("strings lengths delta is not 1")
	}

	inserts := 0

	for i := 0; i < len(a); i++ {
		if len(b) > i-inserts && a[i] == b[i-inserts] {
			continue
		}
		if inserts > 0 {
			return false
		}
		inserts++
	}

	return true
}
