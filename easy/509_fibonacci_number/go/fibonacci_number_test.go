package leetcode

import "testing"

func TestFibonacciNumber(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{20, 6765},
	}
	for _, test := range tests {
		if got := fib(test.input); got != test.want {
			t.Errorf("fib(%d) = %d - want %d", test.input, got, test.want)
		}
	}
}
