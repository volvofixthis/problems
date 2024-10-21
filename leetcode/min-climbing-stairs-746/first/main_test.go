package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		i1 []int
		o1 int
	}{
		{[]int{10, 15, 10}, 15},
	}
	for i, test := range tests {
		t.Run(
			fmt.Sprintf("test-%d", i+1),
			func(t *testing.T) {
				assert.Equal(t, test.o1, minCostClimbingStairs(test.i1), "Must be equal")
			},
		)
	}
}
