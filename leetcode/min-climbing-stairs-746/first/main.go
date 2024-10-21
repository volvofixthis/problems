package main

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	up := len(cost) + 1
	memo := make([]int, up)
	for i := 2; i < up; i++ {
		one_step := memo[i-1] + cost[i-1]
		two_steps := memo[i-2] + cost[i-2]
		memo[i] = Min(one_step, two_steps)
	}
	return memo[up-1]
}
