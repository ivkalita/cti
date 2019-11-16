package task2

// CheckPermutation – O(1) memory, O(n) time
// uses additional set to store total number of each visited rune
func CheckPermutation(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	aRunes := []rune(a)
	bRunes := []rune(b)

	n := len(aRunes)
	diff := make(map[rune]int, n)
	for i := 0; i < n; i++ {
		diff[aRunes[i]] += 1
		diff[bRunes[i]] -= 1
	}

	for _, count := range diff {
		if count > 0 {
			return false
		}
	}

	return true
}
