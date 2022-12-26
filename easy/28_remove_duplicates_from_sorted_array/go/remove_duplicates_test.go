package leetcode

import "testing"

func TestValidParenthesis(t *testing.T) {
	var tests = []struct {
		input   []int
		wantArr []int
		wantIdx int
	}{
		{[]int{1, 1, 2}, []int{1, 2, 2}, 2},
		{[]int{1, 2, 2}, []int{1, 2, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4, 2, 2, 3, 3, 4}, 5},
	}
	for _, test := range tests {
		orig := make([]int, len(test.input))
		copy(orig, test.input)
		if got := removeDuplicates(test.input); got != test.wantIdx || !testEq(test.input, test.wantArr) {
			t.Errorf("removeDuplicates(%v) = %d (%v), want: %v & %d", orig, got, test.input, test.wantArr, test.wantIdx)
		}
	}
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
