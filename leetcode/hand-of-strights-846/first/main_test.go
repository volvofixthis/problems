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
		o1 bool
	}{
		{[]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3, true},
		{[]int{1, 2, 3, 4, 5}, 4, false},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, isNStraightHand(test.i1, test.i2), "should be equal")
		})
	}
}
