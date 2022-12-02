package leetcode

func minCostClimbingStairs(cost []int) int {
	for i := len(cost) - 3; i >= 0; i-- {
		cost[i] += min(cost[i+2], cost[i+1])
	}
	return min(cost[0], cost[1])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
