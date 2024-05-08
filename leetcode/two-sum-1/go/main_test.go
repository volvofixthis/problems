package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		target   int
		expected []int
	}{
		{"first", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"second", []int{3, 2, 4}, 6, []int{1, 2}},
		{"second", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := twoSum(test.input, test.target)
			assert.Equal(t, test.expected, output, "should be equal")
		})
	}
}
