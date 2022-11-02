package validpalindrome

import (
	"regexp"
	"strings"
)

func isPalindrome(x string) bool {
	var r = regexp.MustCompile(`[^a-zA-Z0-9]+`)

	x = r.ReplaceAllString(x, "")
	x = strings.ToLower(x)

	for i := 0; i < len(x)/2; i++ {
		if x[i] != x[len(x)-1-i] {
			return false
		}
	}
	return true
}
