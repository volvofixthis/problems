package main

import (
	"fmt"
	"testing"
)

func TestCalcNeg(t *testing.T) {
	tests := []struct {
		i1 []int
		i2 []int
		o1 int
	}{
		{[]int{1, 4, 5, 12, 17, 100, 500, 505, 600, 1000}, []int{6, 700, 505}, 101},
		{[]int{1000}, []int{500}, 500},
		{[]int{}, []int{}, 0},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d-test", i+1), func(t *testing.T) {
			if v := calcNeg(test.i1, test.i2); v != test.o1 {
				t.Errorf("%d not equal to %d", test.o1, v)
			}
		})
	}
}
