package leetcode

import "testing"

func TestValidParenthesis(t *testing.T) {
	var tests = []struct {
		input   []int
		val     int
		wantArr []int
		wantIdx int
	}{
		{[]int{3, 2, 2, 3}, 3, []int{2, 2, 3, 3}, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 1, 3, 0, 4, 2, 2, 2}, 5},
	}
	for _, test := range tests {
		orig := make([]int, len(test.input))
		copy(orig, test.input)
		if got := removeElement(test.input, test.val); got != test.wantIdx || !testEq(test.input, test.wantArr) {
			t.Errorf("removeElement(%v, %d) = %d, want: %v", orig, test.val, got, test.wantArr)
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
