package main

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		mid := l + (r-l)/2
		if target <= nums[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
