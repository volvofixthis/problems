package main

import (
	"fmt"
	"testing"
)

// Example 1:
//
// Input: matrix =
// [
//
//	[0,1,1,1],
//	[1,1,1,1],
//	[0,1,1,1]
//
// ]
// Output: 15
// Explanation:
// There are 10 squares of side 1.
// There are 4 squares of side 2.
// There is  1 square of side 3.
// Total number of squares = 10 + 4 + 1 = 15.
//
// Example 2:
//
// Input: matrix =
// [
//
//	[1,0,1],
//	[1,1,0],
//	[1,1,0]
//
// ]
// Output: 7
// Explanation:
// There are 6 squares of side 1.
// There is 1 square of side 2.
// Total number of squares = 6 + 1 = 7.

func TestCountSquares(t *testing.T) {
	tests := []struct {
		i1 [][]int
		o1 int
	}{
		{
			i1: [][]int{
				{0, 1, 1, 1},
				{1, 1, 1, 1},
				{0, 1, 1, 1},
			},
			o1: 15,
		},
		{
			i1: [][]int{
				{1, 0, 1},
				{1, 1, 0},
				{1, 1, 0},
			},
			o1: 7,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.i1), func(t *testing.T) {
			if got := countSquares(tt.i1); got != tt.o1 {
				t.Errorf("countSquares() = %v, want %v", got, tt.o1)
			}
		})
	}
}
