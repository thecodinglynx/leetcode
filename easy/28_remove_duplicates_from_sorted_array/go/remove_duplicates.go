package leetcode

func removeDuplicates(nums []int) int {
	res := 0
	if len(nums) <= 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[res] {
			res++
			nums[res] = nums[i]
		}

	}
	return res + 1
}
