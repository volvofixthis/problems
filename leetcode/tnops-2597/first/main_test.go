package first

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		i1 []int
		i2 int
		o1 int
	}{
		{[]int{2, 4, 6}, 2, 4},
		{[]int{1}, 1, 1},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, beautifulSubsets(test.i1, test.i2), "should be equal")
		})
	}
}
