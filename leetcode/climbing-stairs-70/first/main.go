package main

func climbStairsRecursive(i, n int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	return climbStairsRecursive(i+1, n) + climbStairsRecursive(i+2, n)
}

func climbStairs(n int) int {
	return climbStairsRecursive(0, n)
}
