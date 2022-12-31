package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := len(nums1) - 1; i >= 0; i-- {
		if m == 0 {
			nums1[i] = nums2[n-1]
			n--
		} else if n == 0 {
			nums1[i] = nums1[m-1]
			m--
		} else if nums2[n-1] > nums1[m-1] {
			nums1[i] = nums2[n-1]
			n--
		} else {
			nums1[i] = nums1[m-1]
			m--
		}
	}
}
