package leetcode

import "testing"

func TestRotate(t *testing.T) {
	var tests = []struct {
		input []int
		k     int
		want  []int
	}{

		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{1, 2}, 3, []int{2, 1}},
		{[]int{8, 4, 7, 2, 1}, 3, []int{7, 2, 1, 8, 4}},
		{[]int{-1}, 2, []int{-1}},
		{[]int{1, 2}, 0, []int{1, 2}},
		{[]int{1, 2}, 2, []int{1, 2}},
	}
	for _, test := range tests {
		orig := make([]int, len(test.input))
		copy(orig, test.input)
		if rotate(test.input, test.k); !ArrEqual(test.input, test.want) {
			t.Errorf("rotate(%v, %d) = %v, want %v", orig, test.k, test.input, test.want)
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
