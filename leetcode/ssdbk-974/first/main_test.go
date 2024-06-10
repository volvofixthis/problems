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
		o1 int
	}{
		{[]int{4, 5, 0, -2, -3, 1}, 5, 7},
		{[]int{5}, 9, 0},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, subarraysDivByK(test.i1, test.k), "should be equal")
		})
	}
}
