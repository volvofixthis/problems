package first

import "sort"

func heightChecker(heights []int) int {
	hs := make([]int, len(heights))
	copy(hs, heights)
	sort.Ints(hs)
	result := 0
	for i := 0; i < len(heights); i++ {
		if hs[i] != heights[i] {
			result++
		}
	}
	return result
}
