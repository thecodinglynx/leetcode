package leetcode

import "testing"

func TestSingleNumber(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{1, 1, 4}, 4},
		{[]int{5, 6, 8, 3, 6, 3, 1, 8, 5}, 1},
		{[]int{5}, 5},
		{[]int{-336, 513, -560, -481, -174, 101, -997, 40, -527, -784, -283, -336, 513, -560, -481, -174, 101, -997, 40, -527, -784, -283, 354}, 354},
	}
	for _, test := range tests {
		orig := make([]int, len(test.input))
		copy(orig, test.input)
		if got := singleNumber(test.input); got != test.want {
			t.Errorf("singleNumber(%v) = %d, want: %d", orig, got, test.want)
		}
	}
}
