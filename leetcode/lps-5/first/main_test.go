package first

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		i1   string
		o1   bool
	}{
		{"one", "aba", true},
		{"second", "bb", true},
		{"three", "abc", false},
		{"four", "bbc", false},
		{"five", "b", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.o1, isPalindrome(test.i1), "should be equal")
		})
	}
}

func BenchmarkFunc(b *testing.B) {
	for i := 0; i < 100; i++ {
		longestPalindrome("dcddfegcgefkkb")
	}
}

func BenchmarkFunc2(b *testing.B) {
	for i := 0; i < 100; i++ {
		longestPalindrome("abccbabcb")
	}
}

func TestFunc(t *testing.T) {
	tests := []struct {
		name string
		i1   string
		o1   string
	}{
		{"one", "babad", "bab"},
		{"one", "cbbd", "bb"},
		{"three", "a", "a"},
		{"four", "abccbabcb", "abccba"},
		{"five", "dcddfegcgefkkb", "fegcgef"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.o1, longestPalindrome(test.i1), "should be equal")
		})
	}
}
