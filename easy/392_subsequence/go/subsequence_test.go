package leetcode

import "testing"

func TestIsSubsequence(t *testing.T) {
	var tests = []struct {
		s1   string
		s2   string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"b", "abc", true},
		{"axc", "ahbgdc", false},
	}
	for _, test := range tests {
		if got := isSubsequence(test.s1, test.s2); got != test.want {
			t.Errorf("isSubsequence(%s, %s) = %t - want %t", test.s1, test.s2, got, test.want)
		}
	}
}
