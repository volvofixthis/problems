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
		k  int
		o1 int
	}{
		{[]int{1, 3, 4}, []int{1, 3, 4}, 1, 5},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, numberOfPairs(test.i1, test.i2, test.k), "should be equal")
		})
	}
}
