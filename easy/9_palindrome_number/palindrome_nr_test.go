package palindromenumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input int
		want  bool
	}{
		{121, true},
		{5, true},
		{5544334455, true},
		{55655, true},
		{112, false},
		{56, false},
		{9876543, false},
	}
	for _, test := range tests {
		if got := isPalindrome(test.input); got != test.want {
			t.Errorf("isPalindrome(%d) = %v", test.input, got)
		}
	}
}
