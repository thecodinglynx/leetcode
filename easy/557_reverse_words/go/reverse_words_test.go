package leetcode

import "testing"

func TestReverseWords(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{

		{"Let's", "s'teL"},
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
		{"God Ding", "doG gniD"},
		{"I love u", "I evol u"},
	}
	for _, test := range tests {
		if got := reverseWords(test.input); got != test.want {
			t.Errorf("reverseWords(%v) = \"%v\", want \"%v\"", test.input, got, test.want)
		}
	}
}
