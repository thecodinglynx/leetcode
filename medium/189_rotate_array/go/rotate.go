package leetcode

// uses extra array which duplicates the values in the provided array
// E.g. rotate(84721,3) -> 8472184721 -> 84_72184_721 with 72184 being the solution
func rotate(nums []int, k int) {
	d := append(nums, nums...)
	for i := range nums {
		nums[i] = d[(len(nums) - (k % len(nums)) + i)]
	}
}

// another better solution is likely to use two pointers:
// l = 0
// h = len(nums) - k // won't work for all cases though
// and keep switching numbers while incrementing l and h
