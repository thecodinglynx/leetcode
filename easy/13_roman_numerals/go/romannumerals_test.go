package leetcode

import "testing"

func TestRomanToInt(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
	for _, test := range tests {
		if got := romanToInt(test.input); got != test.want {
			t.Errorf("romanToInt(%q) = %d", test.input, got)
		}
	}
}
