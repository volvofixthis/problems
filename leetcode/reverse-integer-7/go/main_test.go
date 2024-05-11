package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		name string
		i1   int
		o1   int
	}{
		{"first", 123, 321},
		{"second", -123, -321},
		{"third", 120, 21},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.o1, reverse(test.i1), "should be equal")
		})
	}
}
