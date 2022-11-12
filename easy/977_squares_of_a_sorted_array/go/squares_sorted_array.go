package leetcode

func sortedSquares(nums []int) []int {
	l := 0
	h := len(nums) - 1
	i := h
	res := make([]int, len(nums))
	for {
		if l >= h {
			return res
		}
		if abs(nums[h]) > abs(nums[l]) {
			res[i] = nums[h] * nums[h]
			h = h - 1

		} else {
			res[i] = nums[l] * nums[l]
			l = l + 1
		}
		i = i - 1
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
