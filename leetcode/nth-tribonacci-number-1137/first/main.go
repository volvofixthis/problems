package main

type memo map[int]int

func dp(n int, m memo) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	if v, ok := m[n]; ok {
		return v
	}
	f := dp(n-3, m) + dp(n-2, m) + dp(n-1, m)
	m[n] = f
	return f
}

func tribonacci(n int) int {
	m := memo{}
	return dp(n, m)
}
