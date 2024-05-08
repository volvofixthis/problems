package main

import "fmt"

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

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
