package mostwater

import "testing"

func TestMaxArea(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1, 1}, 1},
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{2, 3, 4, 5, 18, 17, 6}, 17},
		{[]int{1, 1000, 1000, 6, 2, 5, 4, 8, 3, 7}, 1000},
	}
	for _, test := range tests {
		if got := maxAreaBF(test.input); got != test.want {
			t.Errorf("maxArea(%v) = %v, want %v", test.input, got, test.want)
		}
		if got := maxArea(test.input); got != test.want {
			t.Errorf("maxArea(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}
