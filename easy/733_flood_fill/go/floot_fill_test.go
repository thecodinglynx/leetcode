package leetcode

import "testing"

func TestFloodFill(t *testing.T) {
	var tests = []struct {
		image [][]int
		sr    int
		sc    int
		color int
		want  [][]int
	}{
		{[][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0, [][]int{{0, 0, 0}, {0, 0, 0}}},
		{[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}},
	}
	for _, test := range tests {
		orig := make([][]int, len(test.image))
		for i := range test.image {
			orig[i] = make([]int, len(test.image[i]))
			copy(orig[i], test.image[i])
		}
		if got := floodFill(test.image, test.sr, test.sc, test.color); !ArrEqual(got, test.want) {
			t.Errorf("floodFill(%v, %d, %d, %d) = %v - want %v", orig, test.sr, test.sc, test.color, got, test.want)
		}
	}
}

func ArrEqual(a1, a2 [][]int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := range a1 {
		if len(a1[i]) != len(a2[i]) {
			return false
		}
		for j := range a1[i] {
			if a1[i][j] != a2[i][j] {
				return false
			}
		}
	}
	return true
}
