package task7

// Rotate – O(1) memory, O(n^2) time
// The idea is:
// 1. iterate over each bound (perimeter) from the outermost to innermost
// 2. iterate over one side of each bound and rotate 4 elements at once (one for each side of this bound)
func Rotate(a *[][]int32) {
	// assumes that len(a) == len(a_i)
	n := len(*a)

	top, left := 0, 0

	for side := n; side > 1; side -= 2 {
		right := left + side - 1
		for i := left; i < right; i++ {
			// we a going to rotate 4 elements at once
			x1, y1 := i, top
			x2, y2 := n-top-1, i
			x3, y3 := n-i-1, n-top-1
			x4, y4 := top, n-i-1
			(*a)[x1][y1], (*a)[x2][y2], (*a)[x3][y3], (*a)[x4][y4] = (*a)[x4][y4], (*a)[x1][y1], (*a)[x2][y2], (*a)[x3][y3]
		}

		top++
		left++
	}
}
