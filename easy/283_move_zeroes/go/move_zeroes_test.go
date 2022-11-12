package leetcode

import "testing"

func TestMoveZeroes(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{1, 0, 0, 4, 5, 6, 7}, []int{1, 4, 5, 6, 7, 0, 0}},
		{[]int{1}, []int{1}},
		{[]int{0}, []int{0}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
	}
	for _, test := range tests {
		orig := make([]int, len(test.input))
		copy(orig, test.input)
		if moveZeroes(test.input); !ArrEqual(test.input, test.want) {
			t.Errorf("moveZeroes(%v) = %v, want %v", orig, test.input, test.want)
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
