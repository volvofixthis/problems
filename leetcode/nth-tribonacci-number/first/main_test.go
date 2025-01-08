package main

import (
	"fmt"
	"testing"
)

// The Tribonacci sequence Tn is defined as follows:
//
// T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.
//
// Given n, return the value of Tn.
//
//
//
// Example 1:
//
// Input: n = 4
// Output: 4
// Explanation:
// T_3 = 0 + 1 + 1 = 2
// T_4 = 1 + 1 + 2 = 4
//
// Example 2:
//
// Input: n = 25
// Output: 1389537
//
//
//
// Constraints:
//
//     0 <= n <= 37
//     The answer is guaranteed to fit within a 32-bit integer, ie. answer <= 2^31 - 1.

func TestTribonacci(t *testing.T) {
	tests := []struct {
		i1 int
		o1 int
	}{
		{4, 4},
		{25, 1389537},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			if v := tribonacci(test.i1); v != test.o1 {
				t.Errorf("%d not equal to %d", v, test.o1)
			}
		})
	}
}
