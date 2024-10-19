package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRob(t *testing.T) {
	tests := []struct {
		i1 []int
		o1 int
	}{
		{[]int{1, 2, 3, 1}, 4},
	}
	for _, test := range tests {
		assert.Equal(t, test.o1, rob(test.i1), "Must be equal")
	}
}
