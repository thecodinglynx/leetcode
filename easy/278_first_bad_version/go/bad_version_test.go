package firstbad

import "testing"

func TestFirstBadVersion(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{99999999, 5674737},
	}
	for _, test := range tests {
		if got := firstBadVersion(test.input); got != test.want {
			t.Errorf("firstBadVersion(%d) = %d - got %d", test.input, test.want, got)
		}
	}
}
