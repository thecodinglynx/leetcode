package leetcode

import "testing"

func TestReverseString(t *testing.T) {
	var tests = []struct {
		input []byte
		want  []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
	}
	for _, test := range tests {
		orig := make([]byte, len(test.input))
		copy(orig, test.input)
		if reverseString(test.input); !ArrEqual(test.input, test.want) {
			t.Errorf("reverseString(%v), got %v - want %v", orig, test.input, test.want)
		}
	}
}

func ArrEqual(a1, a2 []byte) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}
