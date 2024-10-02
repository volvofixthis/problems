package first

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		i1 []int
		o1 int
	}{
		{[]int{-3, 2, -3, 4, 2}, 5},
		{[]int{1, 2}, 1},
		{[]int{1, -2, -3}, 5},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, minStartValue(test.i1), "should be equal")
		})
	}
}
