package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volvofixthis/problems/utils"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		i1 []*int
		o1 int
	}{
		{[]*int{utils.TreeVal(1), utils.TreeVal(2), utils.TreeVal(3), utils.TreeVal(4), nil, nil, utils.TreeVal(5)}, 3},
		{[]*int{utils.TreeVal(3), utils.TreeVal(9), utils.TreeVal(20), nil, nil, utils.TreeVal(15), utils.TreeVal(7)}, 2},
		{[]*int{utils.TreeVal(2), nil, utils.TreeVal(3), nil, utils.TreeVal(4), nil, utils.TreeVal(5), nil, utils.TreeVal(6)}, 5},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, minDepth(utils.ArrToBinaryTree2(test.i1)), "should be equal")
		})
	}
}
