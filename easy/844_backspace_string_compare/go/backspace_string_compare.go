package leetcode

func backspaceCompare(s string, t string) bool {
	ns := ""
	nt := ""
	backspaces := 0
	for i := len(s) - 1; i >= 0; i-- {
		if rune(s[i]) == rune('#') {
			backspaces++
		} else if backspaces > 0 {
			backspaces--
		} else {
			ns += string(s[i])
		}
	}
	backspaces = 0
	for i := len(t) - 1; i >= 0; i-- {
		if rune(t[i]) == rune('#') {
			backspaces++
		} else if backspaces > 0 {
			backspaces--
		} else {
			nt += string(t[i])
		}
	}
	return ns == nt
}
