package first

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volvofixthis/problems/utils"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		i1 []int
		i2 int
		i3 int
		o1 []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 4, []int{1, 4, 3, 2, 5}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, utils.ListNodeToArr(reverseBetween(utils.ArrToListNode(test.i1), test.i2, test.i3)), "should be equal")
		})
	}
}
