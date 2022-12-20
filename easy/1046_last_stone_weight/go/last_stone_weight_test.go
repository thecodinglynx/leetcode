package leetcode

import "testing"

func TestRunningSum(t *testing.T) {
	var tests = []struct {
		arr  []int
		want int
	}{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
		{[]int{1}, 1},
		{[]int{8, 8}, 0},
		{[]int{8, 10}, 2},
		{[]int{8, 10, 4, 3, 2}, 1},
	}
	for _, test := range tests {
		if got := lastStoneWeight(test.arr); got != test.want {
			t.Errorf("lastStoneWeight(%v) = %d - want %d", test.arr, got, test.want)
		}
	}
}
