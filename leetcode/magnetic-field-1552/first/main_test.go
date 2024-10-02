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
		{[]int{1, 2, 3, 4, 7}, 3, 3},
		{[]int{5, 4, 3, 2, 1, 1000000000}, 2, 999999999},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, maxDistance(test.i1, test.i2), "should be equal")
		})
	}
}
