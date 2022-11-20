package leetcode

func isSubsequence(s string, t string) bool {
	if len(s) <= 0 {
		return true
	}
	if len(t) <= 0 {
		return false
	}
	if len(s) > len(t) {
		return false
	}
	sp := 0
	for tp := 0; tp < len(t); tp++ {
		if sp == len(s) {
			return true
		}
		if t[tp] == s[sp] {
			sp++
		}
	}
	return sp == len(s)
}
