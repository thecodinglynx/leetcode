package leetcode

func longestPalindrome(s string) int {
	m := make(map[rune]int)
	total := 0
	length := 0
	for i := 0; i < len(s); i++ {
		m[rune(s[i])] += 1
		total++
	}

	for _, v := range m {
		if v%2 == 0 {
			length += v
		} else if (v - 1) > 0 {
			length += (v - 1)
		}
	}

	if length%2 == 0 && total > length {
		length += 1
	}

	return length
}
