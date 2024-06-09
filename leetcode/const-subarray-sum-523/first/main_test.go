package first

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		i1 []int
		k  int
		o1 bool
	}{
		{[]int{23, 2, 4, 6, 7}, 6, true},
		{[]int{23, 2, 6, 4, 7}, 6, true},
		{[]int{23, 2, 6, 4, 7}, 13, false},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, checkSubarraySum(test.i1, test.k), "should be equal")
		})
	}
}
