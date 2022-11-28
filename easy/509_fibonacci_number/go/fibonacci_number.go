package leetcode

func fib(n int) int {
	m := make(map[int]int)
	return _fib(n, m)
}

func _fib(n int, m map[int]int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if _, ok := m[n]; ok {
		return m[n]
	}
	res := fib(n-1) + fib(n-2)
	m[n] = res
	return res
}
