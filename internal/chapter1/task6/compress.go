package task6

import "strconv"

// Compress – O(n) memory, O(n) time
func Compress(inS string) string {
	if len(inS) == 0 {
		return inS
	}

	s := []rune(inS)
	res := make([]rune, 0, len(s))

	lastRune := s[0]
	cnt := 0

	for _, r := range s {
		if r == lastRune {
			cnt++
			continue
		}

		flush(&res, cnt, lastRune)
		lastRune = r
		cnt = 1
	}

	flush(&res, cnt, lastRune)

	if len(s) <= len(res) {
		res = s
	}

	return string(res)
}

func flush(res *[]rune, cnt int, r rune) {
	suffix := append([]rune{r}, []rune(strconv.Itoa(cnt))...)
	*res = append(*res, suffix...)
}
