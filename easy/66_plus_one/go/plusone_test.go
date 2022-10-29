package leetcode

import "testing"

func TestPlusOne(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{5, 4, 3, 2}, []int{5, 4, 3, 3}},
		{[]int{5, 4, 3, 9}, []int{5, 4, 4, 0}},
		{[]int{5}, []int{6}},
		{[]int{9}, []int{1, 0}},
		{[]int{9, 9, 9, 9}, []int{1, 0, 0, 0, 0}},
		{[]int{8, 9, 9, 9}, []int{9, 0, 0, 0}},
	}
	for _, test := range tests {
		orig := make([]int, len(test.input))
		copy(orig, test.input)
		if got := plusOne(test.input); !testEq(got, test.want) {
			t.Errorf("twoSum(%v) = %v", orig, got)
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
