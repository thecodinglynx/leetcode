package leetcode

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	var tests = []struct {
		input []string
		want  string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"flower", "flow", "alf"}, ""},
		{[]string{"flight-attendant", "flight"}, "flight"},
		{[]string{"a", "b", "c", "d"}, ""},
	}
	for _, test := range tests {
		if got := longestCommonPrefix(test.input); got != test.want {
			t.Errorf("longestCommonPrefix(%v) = %s - want %s", test.input, got, test.want)
		}
	}
}
