package leetcode

func pivotIndex(nums []int) int {
	pivLeft := make([]int, len(nums))
	pivRight := make([]int, len(nums))

	rSum := 0
	for i := len(nums) - 2; i >= 0; i-- {
		rSum += nums[i+1]
		pivRight[i] = rSum
	}

	lSum := 0
	for i := 1; i < len(nums); i++ {
		lSum += nums[i-1]
		pivLeft[i] = lSum
	}

	for i := 0; i < len(pivLeft); i++ {
		if pivRight[i] == pivLeft[i] {
			return i
		}
	}

	return -1
}
