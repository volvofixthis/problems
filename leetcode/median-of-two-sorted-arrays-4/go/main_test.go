package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		name string
		i1   []int
		i2   []int
		o1   float64
	}{
		{"first", []int{1, 3}, []int{2}, 2.0},
		{"second", []int{1, 2}, []int{3, 4}, 2.5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.o1, findMedianSortedArrays(test.i1, test.i2), "should be equal")
		})
	}
}
