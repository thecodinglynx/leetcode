package leetcode

import "testing"

func TestClimbingStairs(t *testing.T) {
	var tests = []struct {
		s1   string
		s2   string
		want bool
	}{
		{"ab#c", "ad#c", true},
		{"ab##", "c#d#", true},
		{"a#c", "b", false},
	}
	for _, test := range tests {
		if got := backspaceCompare(test.s1, test.s2); got != test.want {
			t.Errorf("backspaceCompare(%s, %s) = %t", test.s1, test.s2, got)
		}
	}
}
