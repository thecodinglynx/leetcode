package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{" ", true},
		{"raceacar", false},
		{"abba", true},
		{"A man, a plan, a canal: Panama", true},
	}
	for _, test := range tests {
		if got := isPalindrome(test.input); got != test.want {
			t.Errorf("isPalindrome(%s) = %v", test.input, got)
		}
	}
}
