package leetcode

func climbStairs(n int) int {
	m := make(map[int]int)
	m[0] = 1
	m[1] = 1
	return climb(n, m)
}

func climb(n int, m map[int]int) int {
	if n < 0 {
		return 0
	}
	if v, ok := m[n]; ok {
		return v
	}
	res := climb(n-1, m) + climb(n-2, m)
	m[n] = res
	return res
}
