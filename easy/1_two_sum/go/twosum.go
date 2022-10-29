package leetcode

// find two sums while allocating minimal memory (time O(n^2), space O(1))
func twoSumSlowLowMem(nums []int, target int) []int {
	var res []int

out:
	for i, v1 := range nums {
		for j, v2 := range nums {
			if i == j {
				continue
			}
			if v1+v2 == target {
				res = []int{i, j}
				break out
			}
		}
	}
	return res
}

// find two sums using memory to speed up processing (time O(n), space O(n))
func twoSumFastHighMem(nums []int, target int) []int {
	var res []int

	// key = nr, value = index
	m := make(map[int]int)
	for i, v := range nums {
		if idx, ok := m[target-v]; ok {
			res = []int{idx, i}
			break
		}
		m[v] = i
	}

	return res
}
