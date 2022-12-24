package leetcode

import "testing"

func TestValidParenthesis(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"[]", true},
		{"[(){}]", true},
		{"[({{}})]{[]}", true},
		{"[](", false},
		{"[(])", false},
	}
	for _, test := range tests {
		if got := isValid(test.input); got != test.want {
			t.Errorf("isValid(%v) = %t - want %t", test.input, got, test.want)
		}
	}
}
