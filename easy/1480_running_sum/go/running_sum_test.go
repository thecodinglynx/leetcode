package leetcode

import "testing"

func TestRunningSum(t *testing.T) {
	var tests = []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
	}
	for _, test := range tests {
		if got := runningSum(test.arr); !testEq(got, test.want) {
			t.Errorf("runningSum(%v) = %v - want %v", test.arr, got, test.want)
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
