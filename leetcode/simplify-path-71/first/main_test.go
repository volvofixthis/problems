package first

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		i1 string
		o1 string
	}{
		{"/home/", "/home"},
		{"/home//foo/", "/home/foo"},
		{"/home/user/Documents/../Pictures", "/home/user/Pictures"},
		{"/../", "/"},
		{"/", "/"},
		{"/./", "/"},
		{"/.../a/../b/c/../d/./", "/.../b/d"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, simplifyPath(test.i1), "should be equal")
		})
	}
}
