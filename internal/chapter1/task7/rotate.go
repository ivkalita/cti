package task7

// Rotate – O(1) memory, O(n^2) time
func Rotate(a *[][]int32) {
	// assumes that len(a) == len(a_i)
	n := len(*a)

	top, left := 0, 0

	for side := n; side > 1; side -= 2 {
		right := left + side - 1
		for i := left; i < right; i++ {
			// TODO: explain this formula and refactor
			(*a)[n-top-1][i], (*a)[n-i-1][n-top-1], (*a)[top][n-i-1], (*a)[i][top] = (*a)[i][top], (*a)[n-top-1][i], (*a)[n-i-1][n-top-1], (*a)[top][n-i-1]
		}

		top++
		left++
	}
}
