package leetcode

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	tr := []rune(t)
	leftMap := make(map[rune]rune)
	rightMap := make(map[rune]rune)
	for i, v := range s {
		rightVal := tr[i]
		if mVal, ok := leftMap[v]; ok && mVal != rightVal {
			return false
		}
		if mVal, ok := rightMap[rightVal]; ok && mVal != v {
			return false
		}
		leftMap[v] = rightVal
		rightMap[rightVal] = v
	}
	return true
}
