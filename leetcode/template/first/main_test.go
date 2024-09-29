package first

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		i1 []int
		o1 bool
	}{
		{[]int{2, 6, 4, 1}, false},
		{[]int{1, 2, 34, 3, 4, 5, 7, 23, 12}, true},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, threeConsecutiveOdds(test.i1), "should be equal")
		})
	}
}
