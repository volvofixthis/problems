package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		name string
		i1   []int
		i2   int
		o1   []int
	}{
		{"first", []int{1, 2, 3, 5}, 3, []int{2, 5}},
		{"second", []int{1, 7}, 1, []int{1, 7}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.o1, kthSmallestPrimeFraction(test.i1, test.i2), "should be equal")
		})
	}
}
