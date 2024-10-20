package main

func climbStairsRecursive(i, n int, memo []int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	if memo[i] > 0 {
		return memo[i]
	}
	memo[i] = climbStairsRecursive(i+1, n, memo) + climbStairsRecursive(i+2, n, memo)
	return memo[i]
}

func climbStairs(n int) int {
	memo := make([]int, n)
	return climbStairsRecursive(0, n, memo)
}
