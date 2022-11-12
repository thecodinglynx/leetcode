package leetcode

func searchInsert(nums []int, target int) int {
	return _searchInsert(nums, 0, len(nums)-1, len(nums)/2, target)
}

func _searchInsert(nums []int, l, h, m, t int) int {
	if nums[m] == t {
		return m
	}
	if l >= h {
		if t > nums[m] {
			return m + 1
		}
		return m
	}
	if t > nums[m] {
		return _searchInsert(nums, m+1, h, m+(h-m+1)/2, t)
	}
	return _searchInsert(nums, l, m-1, l+(m-l-1)/2, t)
}
