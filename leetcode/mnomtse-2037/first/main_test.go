package first

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		i1 []int
		i2 []int
		o1 int
	}{
		{[]int{3, 1, 5}, []int{2, 7, 4}, 4},
		{[]int{4, 1, 5, 9}, []int{1, 3, 2, 6}, 7},
		{[]int{2, 2, 6, 6}, []int{1, 3, 2, 6}, 4},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, minMovesToSeat(test.i1, test.i2), "should be equal")
		})
	}
}
