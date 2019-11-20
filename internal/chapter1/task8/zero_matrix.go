package task8

// Zeroes – O(1) memory, O(nm) time
// works in 2 steps:
// 1. mark rows and columns that should be zeroed
// 2. zeroes rows and columns
func Zeroes(a *[][]int32) {
	for i := 0; i < len(*a); i++ {
		for j := 0; j < len((*a)[i]); j++ {
			if (*a)[i][j] == 0 {
				(*a)[i][0] = 0
				(*a)[0][j] = 0
			}
		}
	}

	for i := 0; i < len(*a); i++ {
		for j := 0; j < len((*a)[i]); j++ {
			if (*a)[i][0] == 0 || (*a)[0][j] == 0 {
				(*a)[i][j] = 0
			}
		}
	}
}
