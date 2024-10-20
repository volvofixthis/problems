package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		i1 int
		o1 int
	}{
		{2, 2},
		{3, 3},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			assert.Equal(t, test.o1, climbStairs(test.i1), "Must be equal")
		})
	}
}
