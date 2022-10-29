package leetcode

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] = digits[i] + 1
			return digits
		}
	}

	// if we get here, all digits were 9s in which case we simply
	// return prepend a 1 to the already modified digits slice
	return append([]int{1}, digits...)
}
