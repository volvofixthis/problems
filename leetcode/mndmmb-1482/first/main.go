package first

import "sort"

func minDays(bloomDay []int, m int, k int) int {
	sort.Ints(bloomDay)
	flowers := 0
	bouquets := 0
	for _, day := range bloomDay {
		flowers += 1
		if flowers == k {
			bouquets += 1
			flowers = 0
		}
		if bouquets == m {
			return day
		}
	}
	return -1
}
