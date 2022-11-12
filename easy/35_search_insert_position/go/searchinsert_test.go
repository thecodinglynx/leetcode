package leetcode

import "testing"

func TestSearchInsert(t *testing.T) {
	var tests = []struct {
		input  []int
		target int
		want   int
	}{
		{[]int{1, 3}, 4, 2},
		{[]int{1, 3}, 0, 0},
		{[]int{1}, 1, 0},
		{[]int{1}, 0, 0},
		{[]int{1}, 2, 1},
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
	}
	for _, test := range tests {
		if got := searchInsert(test.input, test.target); got != test.want {
			t.Errorf("searchInsert(%v, %d) = %d - want %d", test.input, test.target, got, test.want)
		}
	}
}
