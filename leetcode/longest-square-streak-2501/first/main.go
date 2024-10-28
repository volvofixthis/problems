package main

import (
	"math"
	"sort"
)

func longestSquareStreak(nums []int) int {
	longestStreak := 0
	uniqueNumbers := map[int]struct{}{}
	for _, v := range nums {
		uniqueNumbers[v] = struct{}{}
	}
	for _, v := range nums {
		current := v
		currentStreak := 0
		for {
			if _, ok := uniqueNumbers[current]; !ok {
				break
			}
			currentStreak += 1
			ml := current * current
			if ml > int(math.Pow(10, 5)) {
				break
			}
			current = ml
		}
		longestStreak = max(currentStreak, longestStreak)

	}
	if longestStreak < 2 {
		return -1
	}
	return longestStreak
}

func binarySearch(nums []int, target int) bool {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

func longestSquareStreak2(nums []int) int {
	longestStreak := 0
	sort.Ints(nums)
	seen := map[int]struct{}{}
	for _, v := range nums {
		if _, ok := seen[v]; ok {
			continue
		}
		current := v
		currentStreak := 1
		for {
			ml := current * current
			if ml > int(math.Pow(10, 5)) {
				break
			}
			if binarySearch(nums, ml) {
				seen[ml] = struct{}{}
				currentStreak += 1
				current = ml
			} else {
				break
			}
		}
		longestStreak = max(currentStreak, longestStreak)

	}
	if longestStreak < 2 {
		return -1
	}
	return longestStreak
}
