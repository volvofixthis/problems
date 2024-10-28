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
		{i1: []int{99, 18, 11, 3, 1}, i2: 17, o1: 2},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := searchInsert(tt.i1, tt.i2); got != tt.o1 {
				t.Errorf("searchInsert() = %v, want %v", got, tt.o1)
			}
		})
	}
}
