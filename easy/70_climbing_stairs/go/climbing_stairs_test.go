package leetcode

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
		{7, 21},
		{8, 34},
	}
	for _, test := range tests {
		if got := climbStairs(test.input); got != test.want {
			t.Errorf("climbStairs(%d) = %d - want %d", test.input, got, test.want)
		}
	}
}
