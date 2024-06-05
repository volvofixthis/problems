package first

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		i1 []string
		o1 []string
	}{
		{[]string{"bella", "label", "roller"}, []string{"e", "l", "l"}},
		{[]string{"cool", "lock", "cook"}, []string{"c", "o"}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, commonChars(test.i1), "should be equal")
		})
	}
}
