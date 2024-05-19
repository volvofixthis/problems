package first

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		name string
		i1   string
		o1   int
	}{
		{"one", "42", 42},
		{"two", " -042", -42},
		{"three", "1337c0d3", 1337},
		{"four", "0-1", 0},
		{"five", "words and 987", 0},
		{"six", "01", 1},
		{"seven", "0,0", 0},
		{"eight", "20000000000000000000", 2147483647},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.o1, myAtoi(test.i1), "should be equal")
		})
	}
}
