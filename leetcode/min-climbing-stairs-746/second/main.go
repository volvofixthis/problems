package main

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	var dp func(cost []int, i int, memo []int) int
	dp = func(cost []int, i int, memo []int) int {
		if i <= 1 {
			return 0
		}
		if memo[i] > 0 {
			return memo[i]
		}
		oneStep := dp(cost, i-1, memo) + cost[i-1]
		twoSteps := dp(cost, i-2, memo) + cost[i-2]
		memo[i] = min(oneStep, twoSteps)
		return memo[i]
	}
	return dp(cost, len(cost), make([]int, len(cost)+1))
}
