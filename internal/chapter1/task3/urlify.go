package task3

import "errors"

// Urlify – O(1) memory, O(n) time
// replaces all spaces in the given string by %20
func Urlify(s *[]rune, realLen int) error {
	if len(*s) < realLen {
		return errors.New("string is too short")
	}

	spaces := 0
	for i := 0; i < realLen; i++ {
		if (*s)[i] == ' ' {
			spaces++
		}
	}

	newLen := spaces*2 + realLen
	if newLen > len(*s) {
		return errors.New("string is too short")
	}

	delta := newLen - realLen

	for i := realLen - 1; i >= 0; i-- {
		if (*s)[i] == ' ' {
			delta -= 2
			(*s)[i+delta], (*s)[i+delta+1], (*s)[i+delta+2] = '%', '2', '0'
		} else {
			(*s)[i+delta] = (*s)[i]
		}
	}

	return nil
}
