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
		{[]int{1, 2, 3}, 2},
		{[]int{1, 1, 3, 3, 5, 5, 7, 7}, 0},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, countElements(test.i1), "should be equal")
		})
	}
}
