package first

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	i1 [][]int
	i2 int
	o1 int
}{
	{[][]int{{5, 7}, {1, 3}, {9, 10}}, 10, 2},
	{[][]int{{2, 4}, {1, 3}}, 5, 1},
	{[][]int{{1, 6}}, 6, 0},
	{[][]int{{1, 500000}, {400000, 599999}, {800000, 1500000}}, 1500000, 200000},
	{[][]int{{3, 4}, {4, 8}, {2, 5}, {3, 8}}, 8, 1},
	{[][]int{}, 2, 2},
	{[][]int{{6, 11}, {7, 13}, {8, 9}, {5, 8}, {3, 13}, {11, 13}, {1, 3}, {5, 10}, {8, 13}, {3, 9}}, 14, 1},
}

func BenchmarkCase1(b *testing.B) {
	countDays(tests[3].i2, tests[3].i1)
}

func TestFunc(t *testing.T) {
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, countDays(test.i2, test.i1), "should be equal")
		})
	}
}
