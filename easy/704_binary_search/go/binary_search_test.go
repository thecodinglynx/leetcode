package binarysearch

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		array  []int
		target int
		want   int
	}{
		{[]int{5}, -5, -1},
		{[]int{2, 5}, 0, -1},
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 12, 5},
		{[]int{-1, 0, 3, 5, 9, 12}, 7, -1},
		{[]int{-1, 0, 3, 5, 9, 12}, -1, 0},
		{[]int{-1, 0, 3, 5, 9, 12}, 3, 2},
	}
	for _, test := range tests {
		if got := search(test.array, test.target); got != test.want {
			t.Errorf("search(%v, %d) = %d - want %d", test.array, test.target, got, test.want)
		}
	}
}
