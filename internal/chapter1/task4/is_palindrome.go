package task4

// IsPalindrome – O(1) memory, O(n) time
func IsPalindrome(a string) bool {
	runes := make(map[rune]int, 256)
	for _, r := range a {
		if r == ' ' {
			continue
		}

		val, ok := runes[r]
		if !ok {
			val = 0
		}
		val = (val + 1) % 2
		runes[r] = val
	}

	middleFound := false
	for _, v := range runes {
		if v == 0 {
			continue
		}
		if middleFound {
			return false
		}
		middleFound = true
	}

	return true
}
