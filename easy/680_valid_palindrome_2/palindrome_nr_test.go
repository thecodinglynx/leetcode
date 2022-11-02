package validpalindrome2

import "testing"

func TestValidPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{" ", true},
		{"raceacar", true},
		{"abba", true},
		{"abbac", true},
		{"abca", true},
		{"deeee", true},
		{"cuucu", true},
		{"aydmda", true},
	}
	for _, test := range tests {
		if got := validPalindrome(test.input); got != test.want {
			t.Errorf("isPalindrome(%s) = %v", test.input, got)
		}
	}
}
