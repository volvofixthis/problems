package first

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volvofixthis/problems/utils/binarytree"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		i1 []*int
		o1 [][]int
	}{
		{[]*int{binarytree.TreeVal(3), binarytree.TreeVal(9), binarytree.TreeVal(20), nil, nil, binarytree.TreeVal(15), binarytree.TreeVal(7)}, [][]int{{15, 7}, {9, 20}, {3}}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, levelOrder(binarytree.ArrToBinaryTree2(test.i1)), "should be equal")
		})
	}
}
