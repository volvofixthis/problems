package first

import (
	"math"
	"sort"
)

func canPlaceBalls(x int, position []int, m int) bool {
	prevPos := position[0]
	placed := 1
	for _, v := range position {
		curPos := v
		if curPos-prevPos >= x {
			prevPos = curPos
			placed += 1
		}
		if placed == m {
			return true
		}
	}
	return false
}

func maxDistance(position []int, m int) int {
	answer := 0
	sort.Ints(position)
	low := position[0]
	high := int(float64(position[len(position)-1])/float64(m-1) + 1)
	for low <= high {
		mid := low + int(math.Floor(float64(high-low)/2.0))
		if canPlaceBalls(mid, position, m) {
			answer = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return answer
}
