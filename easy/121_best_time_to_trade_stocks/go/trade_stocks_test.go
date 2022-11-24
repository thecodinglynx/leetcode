package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	var tests = []struct {
		arr  []int
		want int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, test := range tests {
		if got := maxProfit(test.arr); got != test.want {
			t.Errorf("maxProfit(%v) = %d - want %d", test.arr, got, test.want)
		}
	}
}
