package leetcode

import "testing"

func TestPivotIndex(t *testing.T) {
	var tests = []struct {
		arr  []int
		want int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{2, 1, -1}, 0},
		{[]int{-1, -1, 0, 0, -1, -1}, 2},
	}
	for _, test := range tests {
		if got := pivotIndex(test.arr); got != test.want {
			t.Errorf("pivotIndex(%v) = %d - want %d", test.arr, got, test.want)
		}
	}
}
