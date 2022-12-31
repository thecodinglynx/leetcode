package leetcode

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	var tests = []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1, 2}, 2, []int{}, 0, []int{1, 2}},
	}
	for _, test := range tests {
		orig := make([]int, len(test.nums1))
		copy(orig, test.nums1)
		if merge(test.nums1, test.m, test.nums2, test.n); !testEq(test.nums1, test.want) {
			t.Errorf("merge(%v, %d, %v, %d) = %v - want %v", orig, test.m, test.nums2, test.n, test.nums1, test.want)
		}
	}
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
