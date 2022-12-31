package leetcode

import "testing"

func TestRunningSum(t *testing.T) {
	var tests = []struct {
		arr  []int
		want int
	}{
		{[]int{232}, 0},
		{[]int{}, 0},
		{[]int{232, 71, 4, 22, 11, 999999, 123}, 4},
		{[]int{232, 71, 4, 22, 11}, 3},
	}
	for _, test := range tests {
		if got := findNumbers(test.arr); got != test.want {
			t.Errorf("findNumbers(%v) = %d - want %d", test.arr, got, test.want)
		}
	}
}
