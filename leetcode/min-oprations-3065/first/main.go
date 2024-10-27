package main

func minOperations(nums []int, k int) int {
	c := 0
	for _, v := range nums {
		if k > v {
			c++
		}
	}
	return c
}
