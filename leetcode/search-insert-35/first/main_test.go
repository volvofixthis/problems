package main

import "testing"

// Example 1:
//
// Input: nums = [1,3,5,6], target = 5
// Output: 2
//
// Example 2:
//
// Input: nums = [1,3,5,6], target = 2
// Output: 1
//
// Example 3:
//
// Input: nums = [1,3,5,6], target = 7
// Output: 4

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		i1 []int
		i2 int
		o1 int
	}{
		{i1: []int{1, 3, 5, 6}, i2: 5, o1: 2},
		{i1: []int{1, 3, 5, 6}, i2: 2, o1: 1},
		{i1: []int{1, 3, 5, 6}, i2: 7, o1: 4},
		{i1: []int{1, 3, 5, 6}, i2: 0, o1: 0},
		{i1: []int{1, 3, 11, 18, 99}, i2: 17, o1: 3},
		{i1: []int{1}, i2: 0, o1: 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := searchInsert(tt.i1, tt.i2); got != tt.o1 {
				t.Errorf("searchInsert() = %v, want %v", got, tt.o1)
			}
		})
	}
}
