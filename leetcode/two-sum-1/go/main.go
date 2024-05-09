package main

func twoSum(nums []int, target int) []int {
	delta := map[int]int{}
	for i, v := range nums {
		d := target - v
		if j, ok := delta[v]; ok {
			return []int{j, i}
		}
		delta[d] = i
	}
	return []int{}
}
