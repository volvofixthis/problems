package main

import "testing"

// Example 1:
//
// Input: nums = [4,3,6,16,8,2]
// Output: 3
// Explanation: Choose the subsequence [4,16,2]. After sorting it, it becomes [2,4,16].
// - 4 = 2 * 2.
// - 16 = 4 * 4.
// Therefore, [4,16,2] is a square streak.
// It can be shown that every subsequence of length 4 is not a square streak.
//
// Example 2:
//
// Input: nums = [2,3,5,6,7]
// Output: -1
// Explanation: There is no square streak in nums so return -1.

func TestLongestSquareStreak(t *testing.T) {
	tests := []struct {
		i1 []int
		o1 int
	}{
		{i1: []int{4, 3, 6, 16, 8, 2}, o1: 3},
		{i1: []int{2, 3, 5, 6, 7}, o1: -1},
	}

	for _, test := range tests {
		if got := longestSquareStreak(test.i1); got != test.o1 {
			t.Errorf("longestSquareStreak(%v) = %v; want %v", test.i1, got, test.o1)
		}
	}
}

func TestLongestSquareStreak2(t *testing.T) {
	tests := []struct {
		i1 []int
		o1 int
	}{
		{i1: []int{4, 3, 6, 16, 8, 2}, o1: 3},
		{i1: []int{2, 3, 5, 6, 7}, o1: -1},
	}

	for _, test := range tests {
		if got := longestSquareStreak2(test.i1); got != test.o1 {
			t.Errorf("longestSquareStreak(%v) = %v; want %v", test.i1, got, test.o1)
		}
	}
}
