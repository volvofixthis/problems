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
		{"one", "abcabcbb", 3},
		{"two", "bbbbb", 1},
		{"three", "pwwkew", 3},
		{"four", "au", 2},
		{"five", "aab", 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.o1, lengthOfLongestSubstring(test.i1), "should be equal")
		})
	}
}
