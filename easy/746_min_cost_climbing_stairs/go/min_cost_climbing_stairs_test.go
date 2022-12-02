package leetcode

import "testing"

func TestClimbingStairs(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
		{[]int{0, 2, 2, 1}, 2},
		{[]int{0, 1, 2, 2}, 2},
	}
	for _, test := range tests {
		if got := minCostClimbingStairs(test.input); got != test.want {
			t.Errorf("minCostClimbingStairs(%v) = %d - want %d", test.input, got, test.want)
		}
	}
}
