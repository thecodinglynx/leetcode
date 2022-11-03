package mostwater

import "fmt"

// O(n) time complexity solution with O(1) space complexity
// The idea is as follows:
// - have two pointers starting at opposite sides of height array
// - calculate volume between pointers (min(height[left], height[right]) * (right-left))
// - if volume higher than previous, update
// - move the pointer that has a smaller height closer to the other pointer (increment left pointer, decrement right pointer)
// - repeat until left < right
func maxArea(height []int) int {
	res := 0
	l := 0
	r := len(height) - 1
	for l < r {
		cur := area(height[l], height[r], r-l)
		if cur > res {
			res = cur
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return res
}

// bruteforce solution O(n^2) time complexity with some minor memoization
func maxAreaBF(height []int) int {
	res := 0
	m := make(map[string]int)
	for i := 0; i < len(height); i++ {
		for j := len(height) - 1; j > i; j-- {
			if i == j {
				continue
			}
			dist := abs(j - i)
			if val, ok := m[fmt.Sprintf("%d%d", dist, height[i])]; ok {
				if val == height[j] {
					continue
				}
			}
			if val, ok := m[fmt.Sprintf("%d%d", dist, height[j])]; ok {
				if val == height[i] {
					continue
				}
			}
			m[fmt.Sprintf("%d%d", dist, height[i])] = height[j]
			cur := area(height[i], height[j], abs(j-i))
			if cur > res {
				res = cur
			}
		}
	}
	return res
}

func area(lHeight, rHeight, dist int) int {
	if lHeight > rHeight {
		return rHeight * dist
	}
	return lHeight * dist
}

func abs(nr int) int {
	if nr > 0 {
		return nr
	}
	return -nr
}
