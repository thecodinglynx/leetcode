package leetcode

import "testing"

func TestRotate(t *testing.T) {
	var tests = []struct {
		input []string
		k     int
		want  []string
	}{
		{[]string{"i", "love", "leetcode", "i", "love", "coding"}, 2, []string{"i", "love"}},
		{[]string{"i", "love", "love", "i", "love", "coding"}, 3, []string{"love", "i", "coding"}},
		{[]string{"i", "love", "love", "i", "love", "coding"}, 1, []string{"love"}},
		{[]string{"ac", "ab", "ab", "ac", "ad"}, 2, []string{"ab", "ac"}},
		{[]string{"a", "b", "c", "d", "e", "f", "g", "a", "b", "x", "x", "x"}, 3, []string{"x", "a", "b"}},
		{[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4, []string{"the", "is", "sunny", "day"}},
	}
	for _, test := range tests {
		if got := topKFrequent(test.input, test.k); !ArrayEquals(got, test.want) {
			t.Errorf("topKFrequent(%v, %d) = %v, want %v", test.input, test.k, got, test.want)
		}
	}
}

func ArrayEquals(a1, a2 []string) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}
