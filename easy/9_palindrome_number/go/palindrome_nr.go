package palindromenumber

import "fmt"

func isPalindrome(x int) bool {
	str := fmt.Sprint(x)
	len := len(str)
	for i := 0; i < len/2; i++ {
		if str[i] != str[len-1-i] {
			return false
		}
	}
	return true
}
