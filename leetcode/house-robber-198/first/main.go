package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func dp(nums []int, memo map[int]int, p int) int {
	if p == 0 {
		return nums[0]
	}
	if p == 1 {
		return max(nums[0], nums[1])
	}
	if v, ok := memo[p]; ok {
		return v
	}
	v := max(dp(nums, memo, p-1), dp(nums, memo, p-2)+nums[p])
	fmt.Println(p)
	memo[p] = v
	return v
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return dp(nums, map[int]int{}, len(nums)-1)
}
