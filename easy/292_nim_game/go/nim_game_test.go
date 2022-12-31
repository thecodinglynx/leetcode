package leetcode

import "testing"

func TestMoveZeroes(t *testing.T) {
	var tests = []struct {
		input int
		want  bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, true},
		{7, true},
		{8, false},
	}
	for _, test := range tests {
		if got := canWinNim(test.input); got != test.want {
			t.Errorf("canWinNim(%d) = %t, want %t", test.input, got, test.want)
		}
	}
}
