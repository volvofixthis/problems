package main

import (
	"fmt"
	"testing"
)

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
		{[]int{1, 3, 5}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			if v := searchInsert(test.i1, test.i2); v != test.o1 {
				t.Errorf("%d not equal to %d", v, test.o1)
			}
		})
	}
}
