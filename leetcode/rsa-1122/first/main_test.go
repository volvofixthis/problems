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
		o1 []int
	}{
		{[]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}, []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}},
		{[]int{28, 6, 22, 8, 44, 17}, []int{22, 28, 8, 6}, []int{22, 28, 8, 6, 17, 44}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, relativeSortArray(test.i1, test.i2), "should be equal")
		})
	}
}
