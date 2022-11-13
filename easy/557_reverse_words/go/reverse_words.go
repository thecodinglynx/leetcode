package leetcode

func reverseWords(s string) string {
	p1 := 0
	p2 := 0
	res := ""
	next := 0
	for {
		if p1 >= len(s) {
			return res[:len(res)-1]
		}
		for p2 < len(s) && s[p2] != ' ' {
			p2++
		}
		next = p2 + 1
		for p2-1 >= p1 {
			res = res + string(s[p2-1])
			p2--
		}
		res = res + " "
		p1, p2 = next, next
	}
}
