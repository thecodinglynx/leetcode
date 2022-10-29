package leetcode

func romanToInt(val string) int {
	m := make(map[rune]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	var res int
	for i := range val {
		cur := rune(val[i])
		res = res + m[cur]
		if i != 0 {
			prev := rune(val[i-1])
			switch prev {
			case 'I':
				if cur == 'V' || cur == 'X' {
					res = res - (2 * m['I'])
				}
			case 'X':
				if cur == 'L' || cur == 'C' {
					res = res - (2 * m['X'])
				}
			case 'C':
				if cur == 'D' || cur == 'M' {
					res = res - (2 * m['C'])
				}
			}
		}
	}
	return res
}
