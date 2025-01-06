package main

import (
	"fmt"
	"testing"
)

// Example 1:
//
// Input: x = 4
// Output: 2
// Explanation: The square root of 4 is 2, so we return 2.
//
// Example 2:
//
// Input: x = 8
// Output: 2
// Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.

func TestMySqrt(t *testing.T) {
	tests := []struct {
		i1 int
		o1 int
	}{
		{4, 2},
		{8, 2},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			if v := mySqrt(test.i1); test.o1 != v {
				t.Errorf("Target output %d not equal to result %d", test.o1, v)
			}
		})
	}
}
