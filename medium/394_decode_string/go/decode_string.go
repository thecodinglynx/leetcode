package leetcode

import (
	"log"
	"strconv"
)

func decodeString(s string) string {
	return decode(0, len(s)-1, s)
}

func decode(start, end int, s string) string {
	var digit string
	var chars string
	for i := start; i <= end; i++ {
		val := rune(s[i])
		if isOpeningBracket(val) {
			var closingBracketIdx int
			findNextNthClosingBracket := 1
			for j := i + 1; j < len(s); j++ {
				if isOpeningBracket(rune(s[j])) {
					findNextNthClosingBracket++
				}
				if isClosingBracket(rune(s[j])) {
					if findNextNthClosingBracket == 1 {
						closingBracketIdx = j
						break
					} else {
						findNextNthClosingBracket--
					}
				}
			}
			res := decode(i+1, closingBracketIdx-1, s)
			chars += resolve(digit, res)
			i = closingBracketIdx
			digit = ""
		} else if isDigit(val) {
			digit += string(val)
		} else {
			chars += string(val)
		}
	}
	return chars
}

func resolve(digit, chars string) string {
	var result string
	if len(digit) <= 0 {
		digit = "1"
	}
	v, err := strconv.Atoi(digit)
	if err != nil {
		log.Fatalf("Unable to convert string %s to int", digit)
	}
	for k := 0; k < v; k++ {
		result += chars
	}
	return result
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isOpeningBracket(r rune) bool {
	return r == '['
}

func isClosingBracket(r rune) bool {
	return r == ']'
}
