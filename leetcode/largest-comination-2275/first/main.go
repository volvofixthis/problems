package main

import "slices"

func largestCombination(candidates []int) int {
	bitCount := make([]int, 24)
	for _, c := range candidates {
		for i := range 24 {
			if c&(1<<i) != 0 {
				bitCount[i]++
			}
		}
	}
	return slices.Max(bitCount)
}
