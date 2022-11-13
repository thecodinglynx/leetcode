package leetcode

import "testing"

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		input  []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 5, []int{1, 2}},
		{[]int{2, 3, 4}, 7, []int{2, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
		{[]int{-1, 2, 4, 6, 8, 10, 13, 15}, 15, []int{2, 7}},
	}
	for _, test := range tests {
		if got := twoSum(test.input, test.target); !ArrEqual(got, test.want) {
			t.Errorf("twoSum(%v, %d) = %v, want %v", test.input, test.target, got, test.want)
		}
	}
}

func ArrEqual(a1, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}
