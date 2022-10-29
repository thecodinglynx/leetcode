package leetcode

import "testing"

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		arr    []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 22, 33, 44, 55, 66, 77, 88, 99, 999, 888}, 1887, []int{18, 19}},
	}
	for _, test := range tests {
		if got := twoSumSlowLowMem(test.arr, test.target); !testEq(got, test.want) {
			t.Errorf("twoSum(%v, %d) = %v", test.arr, test.target, got)
		}
		if got := twoSumFastHighMem(test.arr, test.target); !testEq(got, test.want) {
			t.Errorf("twoSum(%v, %d) = %v", test.arr, test.target, got)
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
