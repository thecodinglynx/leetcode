package leetcode

import "testing"

func TestRotate(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"3[a]", "aaa"},
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"2[b4[c]d3[a]]", "bccccdaaabccccdaaa"},
	}
	for _, test := range tests {
		if got := decodeString(test.input); got != test.want {
			t.Errorf("decodeString(%s) = %s, want %s", test.input, got, test.want)
		}
	}
}

func TestIsDigit(t *testing.T) {
	var tests = []struct {
		input rune
		want  bool
	}{
		{'0', true},
		{'1', true},
		{'2', true},
		{'3', true},
		{'4', true},
		{'5', true},
		{'6', true},
		{'7', true},
		{'8', true},
		{'9', true},
		{'[', false},
		{']', false},
		{'a', false},
	}
	for _, test := range tests {
		if got := isDigit(test.input); got != test.want {
			t.Errorf("isDigit(%v) = %t, want %t", test.input, got, test.want)
		}
	}
}

func TestOpeningBracket(t *testing.T) {
	var tests = []struct {
		input rune
		want  bool
	}{
		{'0', false},
		{'1', false},
		{'2', false},
		{'3', false},
		{'4', false},
		{'5', false},
		{'6', false},
		{'7', false},
		{'8', false},
		{'9', false},
		{'[', true},
		{']', false},
		{'a', false},
	}
	for _, test := range tests {
		if got := isOpeningBracket(test.input); got != test.want {
			t.Errorf("isOpeningBracket(%v) = %t, want %t", test.input, got, test.want)
		}
	}
}

func TestClosingBracket(t *testing.T) {
	var tests = []struct {
		input rune
		want  bool
	}{
		{'0', false},
		{'1', false},
		{'2', false},
		{'3', false},
		{'4', false},
		{'5', false},
		{'6', false},
		{'7', false},
		{'8', false},
		{'9', false},
		{'[', false},
		{']', true},
		{'a', false},
	}
	for _, test := range tests {
		if got := isClosingBracket(test.input); got != test.want {
			t.Errorf("isClosingBracket(%v) = %t, want %t", test.input, got, test.want)
		}
	}
}
