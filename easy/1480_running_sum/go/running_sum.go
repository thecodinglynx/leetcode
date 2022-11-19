package leetcode

func runningSum(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	res := make([]int, len(nums))
	curSum := nums[0]
	res[0] = curSum
	for i := 1; i < len(nums); i++ {
		curSum = curSum + nums[i]
		res[i] = curSum
	}
	return res
}
