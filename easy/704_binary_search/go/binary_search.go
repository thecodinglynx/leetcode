package binarysearch

func search(nums []int, target int) int {
	return binsearch(nums, 0, len(nums)-1, target)
}

func binsearch(nums []int, l, h, target int) int {
	if l > h {
		return -1
	}
	mid := l + (h-l)/2
	if nums[mid] == target {
		return mid
	}
	if target > nums[mid] {
		return binsearch(nums, mid+1, h, target)
	}
	return binsearch(nums, l, mid-1, target)
}
