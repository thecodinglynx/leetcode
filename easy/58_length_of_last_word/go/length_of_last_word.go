package leetcode

func lengthOfLastWord(s string) int {
	length := 0
	trimmedIdx := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			break
		}
		trimmedIdx++
	}

	for i := 0; i < len(s)-trimmedIdx; i++ {
		if s[i] == ' ' {
			length = 0
		} else {
			length++
		}
	}
	return length
}
