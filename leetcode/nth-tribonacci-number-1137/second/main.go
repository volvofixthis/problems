package main

func dp(n int, m []int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	if v := m[n]; v != 0 {
		return v
	}
	f := dp(n-3, m) + dp(n-2, m) + dp(n-1, m)
	m[n] = f
	return f
}

func tribonacci(n int) int {
	m := make([]int, n+1)
	return dp(n, m)
}
