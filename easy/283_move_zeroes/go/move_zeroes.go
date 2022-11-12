package leetcode

func moveZeroes(nums []int) {
	l := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[l] = nums[i]
			l = l + 1
		}
	}
	for i := 0; i < len(nums)-l; i++ {
		nums[len(nums)-1-i] = 0
	}
}
