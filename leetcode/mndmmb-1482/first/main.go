package first

func getBouquetsNum(bloomDay []int, mid, k int) int {
	flowers := 0
	bouquets := 0
	for _, day := range bloomDay {
		if day <= mid {
			flowers += 1
		} else {
			flowers = 0
		}
		if flowers == k {
			bouquets += 1
			flowers = 0
		}
	}
	return bouquets
}

func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	minDays := -1
	start := 0
	end := bloomDay[0]
	for _, day := range bloomDay {
		if end < day {
			end = day
		}
	}
	for start <= end {
		mid := (end + start) / 2
		if getBouquetsNum(bloomDay, mid, k) >= m {
			minDays = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return minDays
}
