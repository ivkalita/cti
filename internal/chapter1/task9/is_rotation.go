package task9

import "strings"

// IsRotation – O(m) memory, O(n + m) time (because of KMP)
func IsRotation(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	// a = xy
	// b = yx
	// b -> rotation <=> isSubstring(a, bb) = isSubstring(xy, yxyx)
	bb := b + b
	return strings.Index(bb, a) != -1
}
