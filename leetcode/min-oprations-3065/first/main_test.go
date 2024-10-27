package main

import "testing"

// Example 1:
//
// Input: nums = [2,11,10,1,3], k = 10
// Output: 3
// Explanation: After one operation, nums becomes equal to [2, 11, 10, 3].
// After two operations, nums becomes equal to [11, 10, 3].
// After three operations, nums becomes equal to [11, 10].
// At this stage, all the elements of nums are greater than or equal to 10 so we can stop.
// It can be shown that 3 is the minimum number of operations needed so that all elements of the array are greater than or equal to 10.
//
// Example 2:
//
// Input: nums = [1,1,2,4,9], k = 1
// Output: 0
// Explanation: All elements of the array are greater than or equal to 1 so we do not need to apply any operations on nums.
//
// Example 3:
//
// Input: nums = [1,1,2,4,9], k = 9
// Output: 4
// Explanation: only a single element of nums is greater than or equal to 9 so we need to apply the operations 4 times on nums.

func TestMinOperations(t *testing.T) {
	tests := []struct {
		i1 []int
		i2 int
		o2 int
	}{
		{i1: []int{2, 11, 10, 1, 3}, i2: 10, o2: 3},
		{i1: []int{1, 1, 2, 4, 9}, i2: 1, o2: 0},
		{i1: []int{1, 1, 2, 4, 9}, i2: 9, o2: 4},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := minOperations(test.i1, test.i2)
			if result != test.o2 {
				t.Errorf("For i1 = %v and i2 = %d, expected %d but got %d", test.i1, test.i2, test.o2, result)
			}
		})
	}
}
