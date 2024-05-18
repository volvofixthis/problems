package first

import "sort"

var evens = []int{}
var odds = []int{}

func sortArrayByParity(nums []int) []int {
	evens = evens[:0]
	odds = odds[:0]
	for _, v := range nums {
		if v%2 == 0 {
			evens = append(evens, v)
		} else {
			odds = append(odds, v)
		}
	}
	sort.Slice(evens, func(i, j int) bool {
		return evens[i] < evens[j]
	})
	sort.Slice(odds, func(i, j int) bool {
		return odds[i] < odds[j]
	})
	copy(nums, evens)
	copy(nums[len(evens):], odds)
	return nums
}
