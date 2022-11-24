package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	var tests = []struct {
		arr  string
		want int
	}{
		{"abccccdd", 7},
		{"ad", 1},
		{"aad", 3},
		{"aadd", 4},
		{"abd", 1},
		{"abbd", 3},
		{"aabbcc", 6},
		{"", 0},
	}
	for _, test := range tests {
		if got := longestPalindrome(test.arr); got != test.want {
			t.Errorf("longestPalindrome(%s) = %d - want %d", test.arr, got, test.want)
		}
	}
}
