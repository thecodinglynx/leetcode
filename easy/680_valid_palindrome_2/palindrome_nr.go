package validpalindrome2

func validPalindrome(x string) bool {
	for i := 0; i < len(x)/2; i++ {
		if x[i] != x[len(x)-1-i] {
			if isPalindrome(x[i+1:len(x)-i]) || isPalindrome(x[i:len(x)-i-1]) {
				return true
			}
			return false
		}
	}
	return true
}

func isPalindrome(z string) bool {
	for i := 0; i < len(z)/2; i++ {
		if z[i] != z[len(z)-1-i] {
			return false
		}
	}
	return true
}
