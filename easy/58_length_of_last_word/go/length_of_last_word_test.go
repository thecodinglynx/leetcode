package leetcode

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"hello how are you today", 5},
		{"", 0},
		{"hi", 2},
	}
	for _, test := range tests {
		if got := lengthOfLastWord(test.input); got != test.want {
			t.Errorf("lengthOfLastWord(%s) = %d - want %d", test.input, got, test.want)
		}
	}
}
