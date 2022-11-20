package leetcode

import "testing"

func TestIsomorphicStrings(t *testing.T) {
	var tests = []struct {
		s1   string
		s2   string
		want bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"title", "paper", true},
		{"badc", "baba", false},
	}
	for _, test := range tests {
		if got := isIsomorphic(test.s1, test.s2); got != test.want {
			t.Errorf("isIsomorphic(%s, %s) = %t - want %t", test.s1, test.s2, got, test.want)
		}
	}
}
