package main

import (
	"fmt"
	"testing"
)

// Example 1:
//
// Input: num = 9669
// Output: 9969
// Explanation:
// Changing the first digit results in 6669.
// Changing the second digit results in 9969.
// Changing the third digit results in 9699.
// Changing the fourth digit results in 9666.
// The maximum number is 9969.
//
// Example 2:
//
// Input: num = 9996
// Output: 9999
// Explanation: Changing the last digit 6 to 9 results in the maximum number.
//
// Example 3:
//
// Input: num = 9999
// Output: 9999
// Explanation: It is better not to apply any change.

func TestMaximum69Number(t *testing.T) {
	tests := []struct {
		i1 int
		o1 int
	}{
		{9669, 9969},
		{9996, 9999},
		{9999, 9999},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			if v := maximum69Number(test.i1); v != test.o1 {
				t.Errorf("%d isn't equal to %d", v, test.o1)
			}
		})
	}
}
