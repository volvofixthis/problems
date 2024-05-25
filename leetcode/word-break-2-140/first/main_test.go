package first

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		i1 string
		i2 []string
		o1 []string
	}{
		{"catsanddog", []string{"cat", "cats", "and", "sand", "dog"}, []string{"cat sand dog", "cats and dog"}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, wordBreak(test.i1, test.i2), "should be equal")
		})
	}
}
