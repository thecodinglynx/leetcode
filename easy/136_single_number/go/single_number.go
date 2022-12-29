package leetcode

import (
	"sort"
)

func singleNumber(nums []int) int {
	sort.Ints(nums)
	prev := nums[0]
	expectDup := true
	for i := 1; i < len(nums); i++ {
		if expectDup {
			if nums[i] != prev {
				return nums[i-1]
			}
			expectDup = false
		} else {
			prev = nums[i]
			expectDup = true
		}
	}

	return nums[len(nums)-1]
}
