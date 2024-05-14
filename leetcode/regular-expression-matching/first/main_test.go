package first

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		name string
		i1   string
		i2   string
		o1   bool
	}{
		{"one", "aa", "a", false},
		{"two", "aa", "a*", true},
		{"three", "ab", ".*", true},
		{"four", "aab", "c*a*b", true},
		{"five", "caab", "c*.a*b", true},
		{"six", "caab", "c*..a*b", true},
		{"seven", "caab", ".*", true},
		{"eight", "ca", ".b", false},
		{"nine", "abc", "a.c*d*.*", true},
		{"ten", "ab", ".*c", false},
		{"eleven", "aaa", "ab*a*c*a", true},
		{"twelve", "aaa", "a*a*a*a", true},
		{"thirteen", "aaa", "a*a*a*b*a", true},
		{"fourteen", "ab", ".*..", true},
		{"fithteen", "aabcbcbcaccbcaabc", ".*a*aa*.*b*.c*.*a*", true},
		{"sixteen", "abc", ".*bc*", true},
		{"seventeen", "abbabaaaaaaacaa", "a*.*b.a.*c*b*a*c*", true},
		{"eighteen", "aaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*c", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.o1, isMatch(test.i1, test.i2), "should be equal")
		})
	}
}
