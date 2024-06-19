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
		i3 int
		o1 int
	}{
		{[]int{1, 10, 3, 10, 2}, 3, 1, 3},
		{[]int{1, 10, 3, 10, 2}, 3, 2, -1},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, minDays(test.i1, test.i2, test.i3), "should be equal")
		})
	}
}
